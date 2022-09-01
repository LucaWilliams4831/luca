package cli

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"runtime/debug"
	"strings"
	"syscall"
	"time"

	"github.com/sipt/shuttle"
	"github.com/sipt/shuttle/config"
	. "github.com/sipt/shuttle/constant"
	"github.com/sipt/shuttle/controller"
	"github.com/sipt/shuttle/dns"
	"github.com/sipt/shuttle/extension/network"
	"github.com/sipt/shuttle/log"
	"github.com/sipt/shuttle/proxy"
	"github.com/sipt/shuttle/rule"

	_ "github.com/sipt/shuttle/ciphers"
	_ "github.com/sipt/shuttle/proxy/protocol"
	_ "github.com/sipt/shuttle/proxy/selector"
)

type ShuttleCommand struct {
	args *Args
}

const (
	shuttleCommand     = "shuttle"
	shuttleDescription = "execute shuttle module"
)

var (
	StopSocksSignal = make(chan bool, 1)
	StopHTTPSignal  = make(chan bool, 1)
)

var eventChan chan *EventObj

func ListenEvent() {
	eventChan = make(chan *EventObj, 1)
	go dealEvent(eventChan)
}

func dealEvent(c chan *EventObj) {
	for {
		t := <-c
		switch t.Type {
		case EventShutdown.Type:
			log.Logger.Info("[Shuttle] is shutdown, see you later!")
			shutdown(config.CurrentConfig().General.SetAsSystemProxy)
			os.Exit(0)
			return
		case EventReloadConfig.Type:
			_, err := reloadConfig(config.CurrentConfigFile(), StopSocksSignal, StopHTTPSignal)
			if err != nil {
				log.Logger.Error("Reload Config failed: ", err)
				fmt.Println(err.Error())
				os.Exit(1)
			}
		case EventRestartHttpProxy.Type:
			StopHTTPSignal <- true
			go HandleHTTP(config.CurrentConfig(), StopHTTPSignal)
		case EventRestartSocksProxy.Type:
			StopSocksSignal <- true
			go HandleSocks5(config.CurrentConfig(), StopSocksSignal)
		case EventRestartController.Type:
			controller.ShutdownController()
			go controller.StartController(config.CurrentConfig(), eventChan)
		case EventUpgrade.Type:
			//todo
			fileName := t.GetData().(string)
			shutdown(config.CurrentConfig().General.SetAsSystemProxy)
			log.Logger.Info("[Shuttle] is shutdown, for upgrade!")
			var name string
			if runtime.GOOS == "windows" {
				name = "upgrade"
			} else {
				name = "./upgrade"
			}
			cmd := exec.Command(name, "-f="+fileName)
			err := cmd.Start()
			if err != nil {
				fmt.Println(err.Error())
			}
			os.Exit(0)
		}
	}
}

func newShuttleEvmCommand() Command {
	return Command{
		Name:        shuttleCommand,
		Description: shuttleDescription,
		Exec:        &ShuttleCommand{},
	}
}

func (c *ShuttleCommand) ExecCommand(ctx context.Context, args []string) error {
	c.args = &Args{args}
	if len(args) == 0 {
		//
	}
	configPath := flag.String("c", "shuttle.yaml", "configuration file path")
	logMode := flag.String("l", "file", "logMode: off | console | file")
	logPath := flag.String("lp", "logs", "logs path")
	flag.Parse()
	var (
		conf *config.Config
		err  error
	)
	//init Logger
	if err = log.InitLogger(*logMode, *logPath); err != nil {
		fmt.Println(err.Error())
		return nil
	}
	if conf, err = loadConfig(*configPath); err != nil {
		fmt.Println(err.Error())
		return nil
	}

	//event listen
	ListenEvent()

	// 启动api控制
	go controller.StartController(conf, eventChan)
	//go HandleUDP()
	go HandleHTTP(conf, StopHTTPSignal)
	go HandleSocks5(conf, StopSocksSignal)

	// Catch "Ctrl + C"
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	if conf.General.SetAsSystemProxy == "" || conf.General.SetAsSystemProxy == config.SetAsSystemProxyAuto {
		//enable system proxy
		EnableSystemProxy(conf)
	}
	fmt.Println("success")

	<-signalChan
	log.Logger.Info("[Shuttle] is shutdown, see you later!")
	shutdown(conf.General.SetAsSystemProxy)

	return nil
}

//load config
func loadConfig(configPath string) (conf *config.Config, err error) {
	//init Config
	conf, err = config.LoadConfig(configPath)
	if err != nil {
		return
	}
	//init Config Value
	shuttle.InitConfigValue(conf)
	//init DNS & GeoIP
	if err = dns.ApplyConfig(conf); err != nil {
		return
	}
	//init Logger
	if err = log.ApplyConfig(conf); err != nil {
		return
	}
	//init Proxy & ProxyGroup
	if err = proxy.ApplyConfig(conf); err != nil {
		return
	}
	//init Rule
	if err = rule.ApplyConfig(conf); err != nil {
		return
	}
	//init HttpMap
	if err = shuttle.ApplyHTTPModifyConfig(conf); err != nil {
		return
	}
	//init MITM
	if err = shuttle.ApplyMITMConfig(conf); err != nil {
		return
	}
	return
}

//reload config
func reloadConfig(configPath string, StopSocksSignal, StopHTTPSignal chan bool) (conf *config.Config, err error) {
	oldConf := config.CurrentConfig()
	conf, err = loadConfig(configPath)
	if err != nil {
		return
	}
	// controller
	if oldConf.GetControllerInterface() != conf.GetControllerInterface() ||
		oldConf.GetControllerPort() != conf.GetControllerPort() {
		//restart controller
		controller.ShutdownController()
		// 启动api控制
		go controller.StartController(conf, eventChan)
	}

	// http proxy
	if oldConf.GetHTTPInterface() != conf.GetHTTPInterface() ||
		oldConf.GetHTTPPort() != conf.GetHTTPPort() {
		//restart http proxy
		StopHTTPSignal <- true
		go HandleHTTP(conf, StopHTTPSignal)
	}

	// socks5 proxy
	if oldConf.GetSOCKSInterface() != conf.GetSOCKSInterface() ||
		oldConf.GetSOCKSPort() != conf.GetSOCKSPort() {
		//restart http proxy
		StopSocksSignal <- true
		go HandleSocks5(conf, StopSocksSignal)
	}
	return
}

func shutdown(setAsSystemProxy string) {
	controller.ShutdownController()
	StopSocksSignal <- true
	StopHTTPSignal <- true
	if setAsSystemProxy == "" || setAsSystemProxy == config.SetAsSystemProxyAuto {
		//disable system proxy
		DisableSystemProxy()
	}
	log.Logger.Close()
	dns.CloseGeoDB()
	time.Sleep(time.Second)
}

func EnableSystemProxy(config IProxyConfig) {
	network.WebProxySwitch(true, "127.0.0.1", config.GetHTTPPort())
	network.SecureWebProxySwitch(true, "127.0.0.1", config.GetHTTPPort())
	network.SocksProxySwitch(true, "127.0.0.1", config.GetSOCKSPort())
}

func DisableSystemProxy() {
	network.WebProxySwitch(false)
	network.SecureWebProxySwitch(false)
	network.SocksProxySwitch(false)
}

type IProxyConfig interface {
	ISOCKSProxyConfig
	IHTTPProxyConfig
}

//SOCKS5 Proxy
type ISOCKSProxyConfig interface {
	GetSOCKSInterface() string
	SetSOCKSInterface(string)
	GetSOCKSPort() string
	SetSOCKSPort(string)
}

func HandleSocks5(config ISOCKSProxyConfig, stopHandle chan bool) {
	addr := net.JoinHostPort(config.GetSOCKSInterface(), config.GetSOCKSPort())
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	log.Logger.Info("Listen to [SOCKS]: ", addr)
	var shutdown = false
	go func() {
		if shutdown = <-stopHandle; shutdown {
			listener.Close()
			log.Logger.Infof("close socks listener!")
		}
	}()
	for {
		conn, err := listener.Accept()
		if err != nil {
			if shutdown && strings.Contains(err.Error(), "use of closed network connection") {
				log.Logger.Info("Stopped HTTP/HTTPS Proxy goroutine...")
				return
			} else {
				log.Logger.Error(err)
			}
			continue
		}
		go func() {
			defer func() {
				if err := recover(); err != nil {
					log.Logger.Errorf("[HTTP/HTTPS]panic :%v", err)
					log.Logger.Errorf("[HTTP/HTTPS]stack :%s", debug.Stack())
					conn.Close()
				}
			}()
			log.Logger.Debug("[SOCKS]Accept tcp connection")
			shuttle.SocksHandle(conn)
		}()
	}
}

//HTTP Proxy
type IHTTPProxyConfig interface {
	GetHTTPInterface() string
	SetHTTPInterface(string)
	GetHTTPPort() string
	SetHTTPPort(string)
}

func HandleHTTP(config IHTTPProxyConfig, stopHandle chan bool) {
	addr := net.JoinHostPort(config.GetHTTPInterface(), config.GetHTTPPort())
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	log.Logger.Info("Listen to [HTTP/HTTPS]: ", addr)

	var shutdown = false
	go func() {
		if shutdown = <-stopHandle; shutdown {
			listener.Close()
			log.Logger.Infof("close HTTP/HTTPS listener!")
		}
	}()
	for {
		conn, err := listener.Accept()
		if err != nil {
			if shutdown && strings.Contains(err.Error(), "use of closed network connection") {
				log.Logger.Info("Stopped HTTP/HTTPS Proxy goroutine...")
				return
			} else {
				log.Logger.Error(err)
			}
			continue
		}
		go func() {
			defer func() {
				conn.Close()
				if err := recover(); err != nil {
					log.Logger.Errorf("[HTTP/HTTPS]panic :%v", err)
					log.Logger.Errorf("[HTTP/HTTPS]stack :%s", debug.Stack())
				}
			}()
			log.Logger.Debug("[HTTP/HTTPS]Accept tcp connection")
			shuttle.HandleHTTP(conn)
		}()
	}
}
