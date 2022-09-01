package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/btcsuite/btcd/btcutil"
	secp256k1 "github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	skein "github.com/nikola43/panda141035/crypto"
	"golang.org/x/crypto/sha3"
)

type Wallet struct {
	PublicKey  string `json:"public_key"`
	PrivateKey string `json:"private_key"`
}

// SysInfo saves the basic system information
type MasterWallet struct {
	PublicKey        string `json:"public_key"`
	PrivateKey       string `json:"private_key"`
	BtcWallet        Wallet `json:"btc_wallet"`
	EthWallet        Wallet `json:"eth_wallet"`
	MasterPrivateKey string `json:"master_private_key"`
}

func GenRandomBytes(size int) (blk []byte, err error) {
	blk = make([]byte, size)
	_, err = rand.Read(blk)
	return
}

func Public(privateKey string) (publicKey string) {
	var e ecdsa.PrivateKey
	e.D, _ = new(big.Int).SetString(privateKey, 16)
	e.PublicKey.Curve = secp256k1.S256()
	e.PublicKey.X, e.PublicKey.Y = e.PublicKey.Curve.ScalarBaseMult(e.D.Bytes())
	pu := elliptic.Marshal(secp256k1.S256(), e.X, e.Y)
	return fmt.Sprintf("%x", pu)
}

func Private(privateKey string) *ecdsa.PrivateKey {
	var e ecdsa.PrivateKey
	e.D, _ = new(big.Int).SetString(privateKey, 16)
	e.PublicKey.Curve = secp256k1.S256()
	e.PublicKey.X, e.PublicKey.Y = e.PublicKey.Curve.ScalarBaseMult(e.D.Bytes())
	return &e
}

func NewMasterWallet() *MasterWallet {

	token := []byte("E83385AF76B2B1997326B567461FB73DD9C27EAB9E1E86D26779F4650C5F2B75E83385AF76B2B1997326B567461FB73DD9C27EAB9E1E86D26779F4650C5F2B75E83385AF76B2B1997326B567461FB73DD9C27EAB9E1E86D26779F4650C5F2B75E83385AF76B2B1997326B567461FB73DD9C27EAB9E1E86D26779F4650C5F2B75E83385AF76B2B1997326B567461FB73DD9C27EAB9E1E86D26779F4650C5F2B75E83385AF76B2B1997326B567461FB73DD9C27EAB9E1E86D26779F4650C5F2B75E83385AF76B2B1997326B567461FB73DD9C27EAB9E1E86D26779F4650C5F2B75E83385AF76B2B1997326B567461FB73DD9C27EAB9E1E86D26779F4650C5F2B75E83385AF76B2B1997326B567461FB73DD9C27EAB9E1E86D26779F4650C5F2B75E83385AF76B2B1997326B567461FB73DD9C27EAB9E1E86D26779F4650C5F2B75E83385AF76B2B1997326B567461FB73DD9C27EAB9E1E86D26779F4650C5F2B75E83385AF76B2B1997326B567461FB73DD9C27EAB9E1E86D26779F4650C5F2B75E83385AF76B2B1997326B567461FB73DD9C27EAB9E1E86D26779F4650C5F2B75E83385AF76B2B1997326B567461FB73DD9C27EAB9E1E86D26779F4650C5F2B75E83385AF76B2B1997326B567461FB73DD9C27EAB9E1E86D26779F4650C5F2B75E83385AF76B2B1997326B567461FB73DD9C27EAB9E1E86D26779F4650C5F2B75")
	//rand.Read(token)

	log.Println(strings.ToUpper(Public(string(token))))

	//pk := Private(string(token))

	pk, err := GenerateEcdsaPrivateKey()
	if err != nil {
		log.Fatal(err)
	}

	//s := GenerateETHWalletFromPrivateKEy(pk)
	//fmt.Println(s)

	masterWallet := new(MasterWallet)
	masterWallet.BtcWallet = GenerateBTCWalletFromPrivateKey(pk)
	masterWallet.EthWallet = GenerateETHWalletFromPrivateKey(pk)
	masterWallet.MasterPrivateKey = pk.X.String()

	btcPkSkeinHash := HashSkein1024([]byte(masterWallet.BtcWallet.PrivateKey))
	ethPkSkeinHash := HashSkein1024([]byte(masterWallet.EthWallet.PrivateKey))
	l := make([]byte, len(ethPkSkeinHash))
	for i := 0; i < len(l); i++ {
		l[i] = btcPkSkeinHash[i]
		if i > 512 {
			l[i] = ethPkSkeinHash[i]
		}
	}

	masterPrivate := HashSkein1024(l[:128])
	masterPublicKey := HashSkein1024([]byte(masterPrivate[:128]))
	masterWallet.PrivateKey = hex.EncodeToString(masterPrivate)[:256]
	masterWallet.PublicKey = hex.EncodeToString(masterPublicKey)[:43]

	return masterWallet
}

func (mw MasterWallet) ToString() {
	fmt.Println("BTC Public Key", mw.BtcWallet.PublicKey)
	fmt.Println("BTC Private Key", mw.BtcWallet.PrivateKey)
	fmt.Println("ETH Public Key", mw.EthWallet.PublicKey)
	fmt.Println("ETH Private Key", mw.EthWallet.PrivateKey)
	fmt.Println("KAS Public Key", mw.PublicKey)
	fmt.Println("KAS Private Key", mw.PrivateKey)
	fmt.Println("MASTER Private Key", mw.MasterPrivateKey)
}

func (mw MasterWallet) MasterAddressFromBtcEthPrivateKey(btcPk, ethPk string) *MasterWallet {

	masterWallet := new(MasterWallet)
	btcPkSkeinHash := HashSkein1024([]byte(btcPk))
	ethPkSkeinHash := HashSkein1024([]byte(ethPk))

	l := make([]byte, len(ethPkSkeinHash))
	for i := 0; i < len(l); i++ {
		l[i] = btcPkSkeinHash[i]
		if i > 512 {
			l[i] = ethPkSkeinHash[i]
		}
	}

	masterPrivate := HashSkein1024(l[:128])
	masterPublicKey := HashSkein1024([]byte(masterPrivate[:128]))
	masterWallet.PrivateKey = hex.EncodeToString(masterPrivate)[:256]
	masterWallet.PublicKey = hex.EncodeToString(masterPublicKey)[:43]

	return masterWallet
}

func (mw MasterWallet) MasterAddressFromPrivateKey(masterPrivate []byte) *MasterWallet {
	masterWallet := new(MasterWallet)
	masterPublicKey := HashSkein1024([]byte(masterPrivate[:64]))
	masterWallet.PrivateKey = hex.EncodeToString(masterPrivate)[:256]
	masterWallet.PublicKey = hex.EncodeToString(masterPublicKey)[:43]
	return masterWallet
}

func (mw MasterWallet) MasterAddress() string {
	return mw.PublicKey
}

func (mw MasterWallet) EthAddress() string {
	return mw.EthWallet.PublicKey
}

func (mw MasterWallet) BtcAddress() string {
	return mw.BtcWallet.PublicKey
}

func (network Network) GetAddress(wif *btcutil.WIF) (*btcutil.AddressPubKey, error) {
	return btcutil.NewAddressPubKey(wif.PrivKey.PubKey().SerializeCompressed(), network.GetNetworkParams())
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

func ImportWIF(wifStr string) (*btcutil.WIF, error) {
	wif, err := btcutil.DecodeWIF(wifStr)
	if err != nil {
		return nil, err
	}
	/*
		if !wif.IsForNet(network.GetNetworkParams()) {
			return nil, errors.New("The WIF string is not valid for the `" + network.name + "` network")
		}
	*/
	return wif, nil
}

// GeneratePrivateKey returns a private key that is suitable for use with
// secp256k1.
func GenerateEcdsaPrivateKey() (*ecdsa.PrivateKey, error) {
	key, err := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
	if err != nil {
		return nil, err
	}
	return key, nil
}

// GeneratePrivateKey returns a private key that is suitable for use with
// secp256k1.
func GenerateSecp256k1PrivateKey(pk *ecdsa.PrivateKey) (*secp256k1.PrivateKey, error) {
	return secp256k1.PrivKeyFromBytes(pk.D.Bytes()), nil
}

func CreateBTCWifFromPk(pk *secp256k1.PrivateKey) *btcutil.WIF {
	wif, err := networks["btc"].CreateWifFromPk(pk)
	if err != nil {
		log.Fatal(err)
	}
	return wif
}

func GenerateBTCWalletFromWIF(wif *btcutil.WIF) Wallet {

	pk := wif.String()

	address, err := networks["btc"].GetAddress(wif)
	if err != nil {
		log.Fatal(err)
	}
	wallet := Wallet{
		PublicKey:  address.EncodeAddress(),
		PrivateKey: pk,
	}
	return wallet
}

func GenerateBTCWallet() Wallet {

	wif, err := networks["btc"].CreateWif()
	if err != nil {
		log.Fatal(err)
	}
	pk := wif.String()

	address, err := networks["btc"].GetAddress(wif)
	if err != nil {
		log.Fatal(err)
	}
	wallet := Wallet{
		PublicKey:  address.EncodeAddress(),
		PrivateKey: pk,
	}
	return wallet
}

func GenerateBTCWalletFromPrivateKey(privateKey *ecdsa.PrivateKey) Wallet {
	secp256k1PrivateKey, err := GenerateSecp256k1PrivateKey(privateKey)
	if err != nil {
		log.Fatal(err)
	}
	wif, err := networks["btc"].CreateWifFromPk(secp256k1PrivateKey)
	if err != nil {
		log.Fatal(err)
	}
	pk := wif.String()

	address, err := networks["btc"].GetAddress(wif)
	if err != nil {
		log.Fatal(err)
	}
	wallet := Wallet{
		PublicKey:  address.EncodeAddress(),
		PrivateKey: pk,
	}
	return wallet
}

func GenerateETHPrivateKey() string {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)
	publicKey := privateKey.Public()
	_, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	return hexutil.Encode(privateKeyBytes)[2:]
}

func GenerateETHWalletFromPrivateKey(privateKey *ecdsa.PrivateKey) Wallet {
	privateKeyBytes := crypto.FromECDSA(privateKey)
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])

	wallet := Wallet{
		PublicKey:  address,
		PrivateKey: hexutil.Encode(privateKeyBytes)[2:],
	}
	return wallet
}

func GenerateETHWallet() Wallet {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])

	wallet := Wallet{
		PublicKey:  address,
		PrivateKey: hexutil.Encode(privateKeyBytes)[2:],
	}
	return wallet
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

func GenerateETHWalletFromPlainPrivateKey(pk string) (*Wallet, error) {

	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		return nil, err
	}

	publicKeyECDSA, ok := privateKey.Public().(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("error casting public key to ECDSA")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])

	wallet := &Wallet{
		PublicKey:  address,
		PrivateKey: pk,
	}
	return wallet, nil

}
