package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math"
	"net"
	"os"
	"strconv"
	"syscall"
	"time"

	"github.com/fatih/color"
	"github.com/nikola43/panda141035/config"
	"github.com/nikola43/panda141035/crypto"
	"github.com/nikola43/panda141035/router"

	"golang.org/x/sys/unix"

	"github.com/songgao/water"
	"github.com/vishvananda/netlink"
)

const (
	maxBufSize  = 1518
	maxChanSize = 1000
)

var sendBytes = 0
var receivedBytes = 0

// Server represents vpn server
type Server struct {
	Cipher crypto.Cipher
	Router router.Gateway
	Config *config.Config
	Notify chan struct{}

	irb map[string][]string

	read  chan []byte
	write chan []byte
}

type tun struct {
	maxWorkers int

	read  chan []byte
	write chan []byte
}

type header struct {
	version int
	src     net.IP
	dst     net.IP
}

// Run stars workers
func (s Server) Run(ctx context.Context) {
	node, err := s.Config.Whoami()
	if err != nil {
		log.Fatal(err)
	}

	if !s.Config.Server.Insecure {
		if err := s.initCrypto(); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("address:", s.Config.Server.Address)

	err = setupTunInterface(node.PrivateAddresses, s.Config.Server.Mtu)
	if err != nil {
		fmt.Println("setupTunInterface")
		log.Fatal(err)
	}

	s.updateRoutes()
	s.Router.Table().Dump()

	s.read = make(chan []byte, maxChanSize)
	s.write = make(chan []byte, maxChanSize)

	t := &tun{
		maxWorkers: s.Config.Server.MaxWorkers,
	}

	t.read = make(chan []byte, maxChanSize)
	t.write = make(chan []byte, maxChanSize)

	go t.run(ctx)
	go s.run(ctx)

	s.cross(ctx, t)

	s.watcher(ctx)
}

func (s *Server) initCrypto() error {
	switch s.Config.Crypto.Type {
	case "gcm":
		s.Cipher = &crypto.GCM{
			Passphrase: s.Config.Crypto.Key,
		}
		s.Cipher.Init()
	case "cbc":
		s.Cipher = &crypto.CBC{
			Passphrase: s.Config.Crypto.Key,
		}
		s.Cipher.Init()
	default:
		return errors.New("crypto not support")
	}
	return nil
}

func (s *Server) watcher(ctx context.Context) {
	go func() {
		for {
			select {
			case <-s.Notify:
			case <-ctx.Done():
				return
			}

			fmt.Println("updating routes and crypto key ...")

			s.updateRoutes()
			if !s.Config.Server.Insecure {
				s.initCrypto()
			}
		}
	}()
}

func (s *Server) run(ctx context.Context) {
	for i := 0; i < s.Config.Server.MaxWorkers; i++ {
		conn, err := s.listenPacket(ctx)
		if err != nil {
			fmt.Println("listenPacket")
			log.Fatal(err)
		}

		go s.reader(ctx, conn)
		go s.writer(ctx, conn)
	}
}

func (s *Server) cross(ctx context.Context, t *tun) {
	go func() {
		for {
			b := <-s.read

			select {
			case t.write <- b:
			case <-ctx.Done():
				return
			}
		}
	}()

	go func() {
		for {
			b := <-t.read

			select {
			case s.write <- b:
			case <-ctx.Done():
				return
			}
		}
	}()
}

func (s Server) listenPacket(ctx context.Context) (net.PacketConn, error) {
	lc := net.ListenConfig{
		Control: func(network, address string, c syscall.RawConn) error {
			var sockoptErr error
			err := c.Control(func(fd uintptr) {
				sockoptErr = unix.SetsockoptInt(
					int(fd),
					unix.SOL_SOCKET,
					unix.SO_REUSEPORT,
					1,
				)
			})

			if err != nil {
				return err
			}
			return sockoptErr
		},
		KeepAlive: time.Duration(s.Config.Server.Keepalive) * time.Second,
	}

	return lc.ListenPacket(ctx, "udp", s.Config.Server.Address)
}

func (s *Server) reader(ctx context.Context, conn net.PacketConn) {
	for {
		b := make([]byte, maxBufSize)
		n, _, err := conn.ReadFrom(b)
		if err != nil {
			panic(err)
			continue
		}

		if !s.Config.Server.Insecure {
			b, err = s.Cipher.Decrypt(b[:n])
			if err != nil {
				panic(err)
				continue
			}
		}

		select {
		case s.read <- b:
		case <-ctx.Done():
			return
		default:
		}
	}
}

func (s *Server) writer(ctx context.Context, conn net.PacketConn) {
	_, port, err := net.SplitHostPort(s.Config.Server.Address)
	if err != nil {
		fmt.Println("SplitHostPort")
		log.Fatal(err)
	}

	for {
		select {
		case b := <-s.write:
			h, err := parseHeader(b)
			if err != nil {
				fmt.Println("parseHeader")
				panic(err)
				continue
			}

			nexthop := s.Router.Table().Get(h.dst)
			if nexthop != nil {
				rAddr, _ := net.ResolveUDPAddr("udp",
					net.JoinHostPort(nexthop.String(), port))

				if !s.Config.Server.Insecure {
					b, err = s.Cipher.Encrypt(b)
					if err != nil {
						fmt.Println("Encrypt")
						panic(err)
						continue
					}
				}

				_, err = conn.WriteTo(b, rAddr)
				if err != nil {
					fmt.Println("WriteTo")
					panic(err)
				}
			}

		case <-ctx.Done():
			return
		}
	}
}

// updateRoutes updates routes
func (s *Server) updateRoutes() {
	irb := s.Config.GetIRB()

	// add routes
	for nexthop, subnets := range irb {
		for _, subnet := range subnets {
			_, dst, _ := net.ParseCIDR(subnet)
			nexthop := net.ParseIP(nexthop)
			err := s.Router.Table().Add(dst, nexthop)
			if err != nil && !errors.Is(err, os.ErrExist) {
				panic(err)
			}
		}
	}

	// rm routes
	if len(s.irb) > 0 {
		for nexthop := range irb {
			// new host / peer / node
			if _, ok := s.irb[nexthop]; !ok {
				continue
			}
			diff := diffStrSlice(irb[nexthop], s.irb[nexthop])
			for _, subnet := range diff {
				_, dst, _ := net.ParseCIDR(subnet)
				nexthop := net.ParseIP(nexthop)
				err := s.Router.Table().Delete(dst, nexthop)
				if err != nil {
					panic(err)
				}
			}
		}
	}

	s.irb = irb
}

// run stars workers to read/write from tunnel
func (t tun) run(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	for i := 0; i < t.maxWorkers; i++ {
		ifce, err := createTunInterface()
		if err != nil {
			log.Fatal(err)
		}
		go t.reader(ctx, ifce)
		go t.writer(ctx, ifce)
	}

	<-ctx.Done()
}

// reader reads from tun interface
func (t *tun) reader(ctx context.Context, ifce *water.Interface) {
	for {
		b := make([]byte, maxBufSize)
		n, err := ifce.Read(b)
		if err != nil {
			panic(err)
		}
		sendBytes += n
		//fmt.Println(color.CyanString(" Received: "), HumanFileSize(float64(receivedBytes)))
		fmt.Println(color.CyanString(" Send: "), HumanFileSize(float64(sendBytes)))
		select {
		case t.read <- b[:n]:
		case <-ctx.Done():
			return

		default:
		}
	}
}

// writer writes to tun interface
func (t *tun) writer(ctx context.Context, ifce *water.Interface) {
	var b = make([]byte, 0)

	for {
		select {
		case b = <-t.write:
			//fmt.Println("b", b)
			n, err := ifce.Write(b)
			if err != nil {
				panic(err)
			}

			receivedBytes += n
			fmt.Println(color.CyanString(" Received: "), HumanFileSize(float64(receivedBytes)))
			//fmt.Println(color.CyanString(" Send: "), HumanFileSize(float64(sendBytes)))
		case <-ctx.Done():
			return
		}
	}
}

// setupTunInterface creates and sets tun interface attributes
func setupTunInterface(ipaddrs []string, mtu int) error {

	ifname := "stardust"
	config := water.Config{
		DeviceType: water.TUN,
		PlatformSpecificParams: water.PlatformSpecificParams{
			Name:       ifname,
			MultiQueue: true,
		},
	}

	_, err := water.New(config)
	if err != nil {
		return err
	}

	ifce, err := netlink.LinkByName(ifname)
	if err != nil {
		return err
	}

	netlink.LinkSetMTU(ifce, mtu)
	netlink.LinkSetTxQLen(ifce, 1000)
	netlink.LinkSetUp(ifce)

	for _, ipnet := range ipaddrs {
		addr, err := netlink.ParseAddr(ipnet)
		if err != nil {
			return err
		}
		err = netlink.AddrAdd(ifce, addr)
		if err != nil {
			return err
		}
	}

	return nil
}

// createTunInterface creates a cloned tun interface
func createTunInterface() (*water.Interface, error) {
	ifname := "stardust"
	config := water.Config{
		DeviceType: water.TUN,
		PlatformSpecificParams: water.PlatformSpecificParams{
			Name:       ifname,
			MultiQueue: true,
		},
	}

	return water.New(config)
}

func parseHeader(b []byte) (*header, error) {
	if len(b) < net.IPv4len {
		return nil, errors.New("small packet")
	}

	h := new(header)

	h.version = int(b[0] >> 4)

	if h.version == 4 {
		h.src = make(net.IP, net.IPv4len)
		copy(h.src, b[12:16])
		h.dst = make(net.IP, net.IPv4len)
		copy(h.dst, b[16:20])

		return h, nil
	}

	h.src = make(net.IP, net.IPv6len)
	copy(h.src, b[8:24])
	h.dst = make(net.IP, net.IPv6len)
	copy(h.dst, b[24:40])

	return h, nil
}

func diffStrSlice(n, o []string) []string {
	check := make(map[string]bool)
	diff := []string{}

	for _, k := range n {
		check[k] = true
	}

	for _, k := range o {
		if _, ok := check[k]; !ok {
			diff = append(diff, k)
		}
	}

	return diff
}

var (
	suffixes [5]string
)

func Round(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}

func HumanFileSize(size float64) string {
	//fmt.Println(size)
	suffixes[0] = "B"
	suffixes[1] = "KB"
	suffixes[2] = "MB"
	suffixes[3] = "GB"
	suffixes[4] = "TB"

	base := math.Log(size) / math.Log(1024)
	getSize := Round(math.Pow(1024, base-math.Floor(base)), .5, 2)
	//fmt.Println(int(math.Floor(base)))
	getSuffix := suffixes[int(math.Floor(base))]
	return strconv.FormatFloat(getSize, 'f', -1, 64) + " " + string(getSuffix)
}
