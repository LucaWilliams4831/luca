package cli

import (
	"context"
	"fmt"
	"log"

	"github.com/fatih/color"
	skein "github.com/nikola43/panda141035/crypto"
	"github.com/nikola43/panda141035/wallet"
)

type Evm2BitCommand struct {
	args *Args
}

const (
	evm2bitCommand     = "evm2bit"
	evm2bitDescription = "generate BTC public and private key from ETH private key"
)

func newEvm2BitCommand() Command {
	return Command{
		Name:        evm2bitCommand,
		Description: evm2bitDescription,
		Exec:        &Evm2BitCommand{},
	}
}

func (c *Evm2BitCommand) ExecCommand(ctx context.Context, args []string) error {
	c.args = &Args{args}
	if len(args) == 0 {
		return ErrorFromString(fmt.Sprintf("file not found"))
	}
	fmt.Println(color.YellowString("BTC to ETH"))

	ethWallet, err := wallet.GenerateETHWalletFromPlainPrivateKey(args[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(color.CyanString("ETH Public Key: "), color.YellowString(ethWallet.PublicKey))
	fmt.Println(color.CyanString("ETH Private Key: "), color.YellowString(ethWallet.PrivateKey))
	fmt.Println()

	btcDerivedPrivateKey := HashValue(ethWallet.PrivateKey)
	btcDerivedPublicKey, err := GenerateAddressFromPlainPrivateKey(btcDerivedPrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(color.CyanString("BTC Derived Public Key: "), color.YellowString(btcDerivedPublicKey.Hex()))
	fmt.Println(color.CyanString("BTC Derived Private Key: "), color.YellowString(btcDerivedPrivateKey))
	fmt.Println()

	return nil
}
func HashSkein1024(data []byte) []byte {
	sk := new(skein.Skein1024)
	sk.Init(1024)
	sk.Update(data)
	outputBuffer := make([]byte, 1024)
	sk.Final(outputBuffer)
	//return hex.EncodeToString(outputBuffer)
	return outputBuffer
}
