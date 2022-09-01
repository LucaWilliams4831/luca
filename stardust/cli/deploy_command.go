package cli

import (
	"context"
	"fmt"
)

type DeployCommand struct {
	args *Args
}

const (
	deployCommand     = "deploy"
	deployDescription = "deploy web project on chain"
)

func newDeployCommand() Command {
	return Command{
		Name:        deployCommand,
		Description: deployDescription,
		Exec:        &DeployCommand{},
	}
}

func (c *DeployCommand) deployMasterWallet() error {
	fmt.Println("deploy master wallet")
	return nil
}

func (c *DeployCommand) deployBTCWallet() error {
	fmt.Println("deploy BTC wallet")
	return nil
}

func (c *DeployCommand) ExecCommand(ctx context.Context, args []string) error {
	if len(args) == 0 {
		return ErrorFromString(fmt.Sprintf("%s: no subcommand passed", deployCommand))
	}
	c.args = &Args{args}

	return ErrorFromString(fmt.Sprintf("%s: invalid subcommand passed", deployCommand))
}
