package cli

import (
	"context"
	"os"

	"github.com/cristalhq/acmd"
)

type Command = acmd.Command

type CLI struct {
	commands []Command
	cfg      acmd.Config
}

const (
	appName        = "stardust"
	appDescription = "stardust description"
	appVersion     = "0.0.1"
)

var nopFunc = func(context.Context, []string) error { return nil }

func New() *CLI {
	cfg := acmd.Config{
		AppName:        appName,
		AppDescription: appDescription,
		Version:        appVersion,
		Output:         os.Stdout,
		Args:           os.Args,
	}

	// print help if at least one command-line argument is not passed
	if len(cfg.Args) == 1 {
		cfg.Args = append(cfg.Args, "help")
	}

	c := &CLI{
		cfg: cfg,
	}
	c.registerCommands()

	return c

}

func (c *CLI) registerCommands() {
	cmds := make([]Command, 0)
	cmds = append(cmds, newCreateCommand())
	cmds = append(cmds, newInfoCommand())
	cmds = append(cmds, newBit2EvmCommand())
	cmds = append(cmds, newEvm2BitCommand())
	cmds = append(cmds, newUploadCommand())
	cmds = append(cmds, newDeployCommand())
	cmds = append(cmds, newNodeCommand())
	cmds = append(cmds, newShuttleEvmCommand())
	c.commands = cmds
}

func (c *CLI) Run() error {
	return acmd.RunnerOf(c.commands, c.cfg).Run()
}
