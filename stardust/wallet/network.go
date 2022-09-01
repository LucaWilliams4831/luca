package wallet

import (
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	secp256k1 "github.com/decred/dcrd/dcrec/secp256k1/v4"
)

type Network struct {
	name        string
	symbol      string
	xpubkey     byte
	xprivatekey byte
}

var networks = map[string]Network{
	"rdd": {name: "reddcoin", symbol: "rdd", xpubkey: 0x3d, xprivatekey: 0xbd},
	"dgb": {name: "digibyte", symbol: "dgb", xpubkey: 0x1e, xprivatekey: 0x80},
	"btc": {name: "bitcoin", symbol: "btc", xpubkey: 0x00, xprivatekey: 0x80},
	"ltc": {name: "litecoin", symbol: "ltc", xpubkey: 0x30, xprivatekey: 0xb0},
}

func (network Network) GetNetworkParams() *chaincfg.Params {
	networkParams := &chaincfg.MainNetParams
	networkParams.PubKeyHashAddrID = network.xpubkey
	networkParams.PrivateKeyID = network.xprivatekey
	return networkParams
}

func CreatePrivateKey() (*secp256k1.PrivateKey, error) {
	secret, err := btcec.NewPrivateKey()
	if err != nil {
		return nil, err
	}
	return secret, err
}

func (network Network) CreateWifFromPk(secret *secp256k1.PrivateKey) (*btcutil.WIF, error) {

	return btcutil.NewWIF(secret, network.GetNetworkParams(), true)
}

func (network Network) CreateWif() (*btcutil.WIF, error) {
	secret, err := CreatePrivateKey()
	if err != nil {
		return nil, err
	}
	return btcutil.NewWIF(secret, network.GetNetworkParams(), true)
}
