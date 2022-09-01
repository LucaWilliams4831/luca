package cli

import (
	"context"
	"fmt"
)

type InfoCommand struct {
	args *Args
}

const (
	infoCommand     = "info"
	infoDescription = "get info about system, wallets, networks, nodes"
	walletsInfo     = "wallets"
	networkInfo     = "network"
	nodeInfo        = "node"
)

func newInfoCommand() Command {
	return Command{
		Name:        infoCommand,
		Description: infoDescription,
		Exec:        &InfoCommand{},
	}
}

func (c *InfoCommand) walletsInfo() error {
	fmt.Println("wallets info")

	return nil
}

func (c *InfoCommand) networkInfo() error {
	networkID := c.args.pop()
	if networkID == "" {
		return ErrorFromString(fmt.Sprintf("%s-network: no network id passed", infoCommand))
	}

	fmt.Println("network info")

	return nil
}

func (c *InfoCommand) nodeInfo() error {
	nodeHash := c.args.pop()
	if nodeHash == "" {
		return ErrorFromString(fmt.Sprintf("%s-node: no node hash passed", infoCommand))
	}

	fmt.Println("node info")
	return nil
}

func (c *InfoCommand) ExecCommand(ctx context.Context, args []string) error {
	if len(args) == 0 {
		return ErrorFromString(fmt.Sprintf("%s: no subcommand passed", infoCommand))
	}
	c.args = &Args{args}

	subcommand := c.args.pop()
	switch subcommand {
	case walletsInfo:
		return c.walletsInfo()
	case networkInfo:
		return c.networkInfo()
	case nodeInfo:
		return c.nodeInfo()
	}

	return ErrorFromString(fmt.Sprintf("%s: invalid subcommand passed", infoCommand))
}
