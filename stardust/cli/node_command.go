package cli

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/fatih/color"
	"github.com/nikola43/stardust/NodeManagerV83"
	"github.com/nikola43/stardust/config"
	"github.com/nikola43/stardust/router"
	"github.com/nikola43/stardust/server"
	"github.com/nikola43/stardust/sysinfo"
	"github.com/nikola43/stardust/wallet"
	"github.com/nikola43/web3golanghelper/web3helper"
)

type NodeCommand struct {
	args *Args
}

const (
	nodeCommand     = "node"
	nodeDescription = "node wallets, networks, nodes"

	Run        = "run"
	UpdateEtcd = "update-config"
)

func newNodeCommand() Command {
	return Command{
		Name:        nodeCommand,
		Description: nodeDescription,
		Exec:        &NodeCommand{},
	}
}

func (c *NodeCommand) Run() error {

	fmt.Println("run server")
	RunServer()
	return nil
}

func (c *NodeCommand) UpdateEtcd() {
	fmt.Println("UpdateConf")
	err := config.UpdateConf(update, configFile)
	if err != nil {
		fmt.Println("UpdateConf")
		panic(err)
	}
	os.Exit(0)
}

func (c *NodeCommand) ExecCommand(ctx context.Context, args []string) error {
	if len(args) == 0 {
		return ErrorFromString(fmt.Sprintf("%s: no subcommand passed", nodeCommand))
	}
	c.args = &Args{args}

	subcommand := c.args.pop()
	switch subcommand {
	case Run:
		c.Run()
	case UpdateEtcd:
		c.UpdateEtcd()
	}

	return nil
}

func RunServer() {

	mw := wallet.NewMasterWallet()

	// system config
	numCpu := runtime.NumCPU()
	usedCpu := numCpu
	runtime.GOMAXPROCS(usedCpu)

	PrintSystemInfo(numCpu, usedCpu)
	PrintNetworkStatus()
	PrintUserBalance(mw.PublicKey, 932)
	PrintUserBalance2(mw.PublicKey, 923)

	mw.ToString()

	//sysHash := GetSysInfo()

	//key := make([]byte, 32)
	//rand.Read(key)
	//fmt.Println(key)
	//crypto.EncryptSysData([]byte(sysHash), []byte(key))
	//crypto.DecryptFile("sysdata.txt.bin", []byte(key))
	//os.Exit(0)

	// create unix syscall
	sig := make(chan os.Signal, 1)
	notify := make(chan struct{}, 1)
	signal.Notify(sig, os.Interrupt, os.Kill)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	UpdateEtcdConf()

	// get etcd config
	cfg := GetEtcdConfig()
	cfg.Watcher(ctx, notify)

	// init server
	r := router.New(ctx)
	s := server.Server{
		Config: cfg,
		Router: r,
		Notify: notify,
	}
	s.Run(ctx)
	<-sig
}

func InitServer(octx context.Context, notify *chan struct{}) {
	r := router.New(octx)
	s := server.Server{
		Config: cfg,
		Router: r,
		Notify: *notify,
	}

	s.Run(octx)
}

func UpdateEtcdConf() {
	// check if we need update nodes config file
	if update != "" {
		err := config.UpdateConf(update, configFile)
		if err != nil {
			fmt.Println("UpdateConf")
			panic(err)
		}
		os.Exit(0)
	}
}

func GetEtcdConfig() *config.Config {
	var cfg *config.Config

	// get etcd config
	if etcd {
		cfg = config.New().FromEtcd(configFile)
	} else {
		cfg = config.New().FromFile(configFile)
	}
	err := cfg.Load()
	if err != nil {
		log.Fatal(err)
	}

	return cfg
}

func GetSysInfo() string {

	info := sysinfo.NewSysInfo()
	fmt.Printf("%+v\n", info)
	fmt.Printf("%+s\n", info.ToString())
	fmt.Printf("%+s\n", info.ToHash())
	return info.ToHash()
}

func InitWeb3() {
	pk := "b366406bc0b4883b9b4b3b41117d6c62839174b7d21ec32a5ad0cc76cb3496bd"
	rpcUrl := "https://speedy-nodes-nyc.moralis.io/84a2745d907034e6d388f8d6/avalanche/testnet"
	wsUrl := "wss://speedy-nodes-nyc.moralis.io/84a2745d907034e6d388f8d6/avalanche/testnet/ws"
	web3GolangHelper := web3helper.NewWeb3GolangHelper(rpcUrl, wsUrl, pk)

	chainID, err := web3GolangHelper.HttpClient().NetworkID(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Chain Id: " + chainID.String())

	proccessEvents(web3GolangHelper)
}

func proccessEvents(web3GolangHelper *web3helper.Web3GolangHelper) {
	nodeAddress := "0x2Fcd73952e53aAd026c378F378812E5bb069eF6E"
	nodeAbi, _ := abi.JSON(strings.NewReader(string(NodeManagerV83.NodeManagerV83ABI)))
	fmt.Println(color.YellowString("  ----------------- Blockchain Events -----------------"))
	fmt.Println(color.CyanString("\tListen node manager address: "), color.GreenString(nodeAddress))
	logs := make(chan types.Log)
	sub := BuildContractEventSubscription(web3GolangHelper, nodeAddress, logs)

	for {
		select {
		case err := <-sub.Err():
			fmt.Println(err)
			//out <- err.Error()

		case vLog := <-logs:
			fmt.Println("paco")
			fmt.Println("vLog.TxHash: " + vLog.TxHash.Hex())
			res, err := nodeAbi.Unpack("GiftCardPayed", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(res)
			//services.SetGiftCardIntentPayment(res[2].(string))

		}
	}
}

func BuildContractEventSubscription(web3GolangHelper *web3helper.Web3GolangHelper, contractAddress string, logs chan types.Log) ethereum.Subscription {

	query := ethereum.FilterQuery{
		Addresses: []common.Address{common.HexToAddress(contractAddress)},
	}

	sub, err := web3GolangHelper.WebSocketClient().SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		fmt.Println(sub)
	}
	return sub
}

func PrintSystemInfo(numCpu, usedCpu int) {
	fmt.Println("")
	fmt.Println(color.YellowString("  ----------------- System Info -----------------"))
	fmt.Println(color.CyanString("\t    Number CPU cores available: "), color.GreenString(strconv.Itoa(numCpu)))
	fmt.Println(color.CyanString("\t    Used of CPU cores: "), color.YellowString(strconv.Itoa(usedCpu)))
	fmt.Println()
}

func PrintNetworkStatus() {
	fmt.Println(color.YellowString("  ----------------- Network Info -----------------"))
	fmt.Println(color.CyanString("\t    Number Nodes: "), color.YellowString(strconv.Itoa(3)))
	fmt.Println(color.CyanString("\t    Prague: "), color.YellowString(strconv.Itoa(1)))
	fmt.Println(color.CyanString("\t    Kiev: "), color.YellowString(strconv.Itoa(1)))
	fmt.Println(color.CyanString("\t    Singapour: "), color.YellowString(strconv.Itoa(1)))
	fmt.Println()
}

func PrintUserBalance(address string, balance int) {
	fmt.Println(color.YellowString("  ----------------- Node Owner -----------------"))
	fmt.Println(color.CyanString("  "), color.GreenString(address))
	fmt.Println(color.CyanString("\t    Balance: "), color.YellowString(strconv.Itoa(balance)), color.YellowString(" $ZOE"))
	fmt.Println()
}

func PrintUserBalance2(address string, balance int) {
	fmt.Println(color.YellowString("  ----------------- Network Info -----------------"))
	fmt.Println(color.CyanString("\t    Send: "), color.YellowString(strconv.Itoa(1732)), color.YellowString("MB"))
	fmt.Println(color.CyanString("\t    Received: "), color.YellowString(strconv.Itoa(1343)), color.YellowString("MB"))
	fmt.Println(color.CyanString("\t    Duration: "), color.YellowString("19:20:04"))
	fmt.Println(color.RedString("\t    Paid: "), color.YellowString(strconv.Itoa(1)), color.YellowString("$ZOE = 2.52$"))
	fmt.Println()
}
