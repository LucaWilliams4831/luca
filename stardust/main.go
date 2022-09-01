package stardust

import (
	"fmt"
	"os"

	"github.com/fatih/color"
<<<<<<< HEAD
	"github.com/nikola43/panda141035/cli"
	"github.com/nikola43/panda141035/config"
	qrterminal "github.com/mdp/qrterminal/v3"
=======
	qrterminal "github.com/mdp/qrterminal/v3"
	"github.com/nikola43/stardust/cli"
	"github.com/nikola43/stardust/config"
>>>>>>> ba7fba0 (built with no error)
)

var (
	configFile string
	update     string
	etcd       bool
	cfg        *config.Config
)

func main() {
	ShowPaymentQr()
	if err := cli.New().Run(); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}

func ShowPaymentQr() {

	fmt.Println()
	fmt.Println(color.CyanString("Send 1 ETH to: "), color.YellowString("0x6d5F00aE01F715D3082Ad40dfB5c18A1a35d3A17"))
	fmt.Println()

	// Generate a 'dense' qrcode with the 'Low' level error correction and write it to Stdout
	qrterminal.Generate("0x6d5F00aE01F715D3082Ad40dfB5c18A1a35d3A17", qrterminal.H, os.Stdout)
}
