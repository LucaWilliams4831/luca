package cli

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/nikola43/panda141035/config"
	"github.com/nikola43/panda141035/sysinfo"
	"gopkg.in/yaml.v2"
)

var (
	configFile string
	update     string
	etcd       bool
	cfg        *config.Config
)

func init() {
	flag.StringVar(&configFile, "config", "", "configuration file")
	flag.StringVar(&update, "update", "", "update etc / file")
	flag.BoolVar(&etcd, "etcd", true, "enable etcd")
	flag.Parse()
}

type CreateCommand struct {
	args *Args
}

type Conf struct {
	Revision uint32 `yaml:"revision"`
	Etcd     Etcd   `yaml:"etcd"`
	Server   Server `yaml:"server"`
	Crypto   Crypto `yaml:"crypto"`
	Nodes    []Node `yaml:"nodes"`
}

type Etcd struct {
	Endpoints []string `yaml:"endpoints"`
	Timeout   uint32   `yaml:"timeout"`
}

type Server struct {
	Keepalive uint32 `yaml:"keepalive"`
	Insecure  bool   `yaml:"insecure"`
	Mtu       uint32 `yaml:"mtu"`
}

type Crypto struct {
	Type string `yaml:"type"`
	Key  string `yaml:"key"`
}

type Node struct {
	Node NodeInfo `yaml:"node"`
}

type NodeInfo struct {
	Name             string   `yaml:"name"`
	Address          string   `yaml:"address"`
	PrivateAddresses []string `yaml:"privateAddresses"`
	PrivateSubnets   []string `yaml:"privateSubnets"`
}

/*

master-wallet
$ stardust create btc-wallet
$ stardust create eth-wallet
$ stardust create network
$ stardust create network 3
$ stardust create node 2cfd7f563b
*/

const (
	createCommand     = "create"
	createDescription = "create wallets, networks, nodes"

	createMasterWallet = "master-wallet"
	createBTCWallet    = "btc-wallet"
	createEthWallet    = "eth-wallet"
	createNetwork      = "network"
	createNode         = "node"
)

func newCreateCommand() Command {
	return Command{
		Name:        createCommand,
		Description: createDescription,
		Exec:        &CreateCommand{},
	}
}

func (c *CreateCommand) createMasterWallet() error {
	fmt.Println("create master wallet")
	return nil
}

func (c *CreateCommand) createBTCWallet() error {
	fmt.Println("create BTC wallet")
	return nil
}

func (c *CreateCommand) createEthWallet() error {
	fmt.Println("create ETH wallet")
	return nil
}

func (c *CreateCommand) createNetwork() error {
	args := c.args
	fmt.Println(args)

	var config Conf
	config.Revision = 1
	config.Etcd.Endpoints = append(config.Etcd.Endpoints, "localhost:2379")
	config.Etcd.Timeout = 5
	config.Server.Keepalive = 10
	config.Server.Insecure = false
	config.Server.Mtu = 1300
	config.Crypto.Type = "gcm"
	config.Crypto.Key = "6368616e676520746869732070617373776f726420746f206120736563726574"

	var a []string
	var b []string
	a = append(a, "10.110.0.1/24")
	b = append(b, "10.110.0.0/24")

	name := "node1"
	nodes := Node{Node: NodeInfo{name, sysinfo.GeLocalIp().String(), a, b}}
	config.Nodes = append(config.Nodes, nodes)

	data, err := yaml.Marshal(&config)

	if err != nil {

		log.Fatal(err)
	}

	err2 := ioutil.WriteFile("stardust.yaml", data, 0755)

	if err2 != nil {

		log.Fatal(err2)
	}

	return nil
}

func GetConf() *Conf {
	c := new(Conf)

	yamlFile, err := ioutil.ReadFile("stardust.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

func (c *CreateCommand) createNode() error {
	config := GetConf()
	var a []string
	var b []string
	nodeNumber := strconv.Itoa(len(config.Nodes) + 1)
	a = append(a, "10.110.0."+nodeNumber+"/24")
	b = append(b, "10.110.0.0/24")

	name := "node" + nodeNumber
	nodes := Node{Node: NodeInfo{name, sysinfo.GeLocalIp().String(), a, b}}

	config.Nodes = append(config.Nodes, nodes)

	data, err := yaml.Marshal(&config)

	if err != nil {

		log.Fatal(err)
	}

	err2 := ioutil.WriteFile("stardust.yaml", data, 0)

	if err2 != nil {

		log.Fatal(err2)
	}

	return nil
}

func (c *CreateCommand) ExecCommand(ctx context.Context, args []string) error {
	if len(args) == 0 {
		return ErrorFromString(fmt.Sprintf("%s: no subcommand passed", createCommand))
	}
	c.args = &Args{args}

	subcommand := c.args.pop()
	switch subcommand {
	case createMasterWallet:
		return c.createMasterWallet()
	case createBTCWallet:
		return c.createBTCWallet()
	case createEthWallet:
		return c.createEthWallet()
	case createNode:
		return c.createNode()
	case createNetwork:
		return c.createNetwork()
	}

	return ErrorFromString(fmt.Sprintf("%s: invalid subcommand passed", createCommand))
}
