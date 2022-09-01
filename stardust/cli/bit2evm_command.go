package cli

import (
	"context"
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/fatih/color"
	wallet "github.com/nikola43/panda141035/wallet"
)

type Bit2EvmCommand struct {
	args *Args
}

const (
	bit2evmCommand     = "bit2evm"
	bit2evmDescription = "generate ETH public and private key from BTC private key"
)

func newBit2EvmCommand() Command {
	return Command{
		Name:        bit2evmCommand,
		Description: bit2evmDescription,
		Exec:        &Bit2EvmCommand{},
	}
}

func (c *Bit2EvmCommand) ExecCommand(ctx context.Context, args []string) error {
	c.args = &Args{args}
	if len(args) == 0 {
		return ErrorFromString(fmt.Sprintf("file not found"))
	}
	wif, err := wallet.ImportWIF(args[0])
	if err != nil {
		log.Fatal(err)
	}
	w := wallet.GenerateBTCWalletFromWIF(wif)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(color.CyanString("BTC Public Key: "), color.YellowString(w.PublicKey))
	fmt.Println(color.CyanString("BTC Private Key: "), color.YellowString(w.PrivateKey))
	fmt.Println()

	ethDerivedPrivateKey := HashValue(w.PrivateKey)
	ethDerivedPublicKey, err := GenerateAddressFromPlainPrivateKey(ethDerivedPrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(color.CyanString("ETH Derived Public Key: "), color.YellowString(ethDerivedPublicKey.Hex()))
	fmt.Println(color.CyanString("ETH Derived Private Key: "), color.YellowString(ethDerivedPrivateKey))
	fmt.Println()

	return nil
}

func GenerateAddressFromPlainPrivateKey(pk string) (common.Address, error) {

	var address common.Address
	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		return address, err
	}

	publicKeyECDSA, ok := privateKey.Public().(*ecdsa.PublicKey)
	if !ok {
		return address, errors.New("error casting public key to ECDSA")
	}

	return crypto.PubkeyToAddress(*publicKeyECDSA), nil
}

func HashValue(value string) string {
	hash := sha256.New()
	hash.Write([]byte(value))
	return hex.EncodeToString(hash.Sum(nil))
}
