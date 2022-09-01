// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package NodeManagerV83

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// Node is an auto generated low-level Go binding around an user-defined struct.
type Node struct {
	Id          uint32
	TierIndex   uint8
	Owner       common.Address
	CreatedTime uint32
	ClaimedTime uint32
	LimitedTime uint32
	Multiplier  *big.Int
	Leftover    *big.Int
}

// Tier is an auto generated low-level Go binding around an user-defined struct.
type Tier struct {
	Id             uint8
	Name           string
	Price          *big.Int
	RewardsPerTime *big.Int
	ClaimInterval  uint32
	MaintenanceFee *big.Int
}

// NodeManagerV83MetaData contains all meta data concerning the NodeManagerV83 contract.
var NodeManagerV83MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"GiftCardPayed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"NodeCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"NodeTransfered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"NodeUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_isBlacklisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardsPerTime\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"claimInterval\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"maintenanceFee\",\"type\":\"uint256\"}],\"name\":\"addTier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"}],\"name\":\"bindBoostNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"boostNFT\",\"outputs\":[{\"internalType\":\"contractIBoostNFT\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32[]\",\"name\":\"indice\",\"type\":\"uint32[]\"}],\"name\":\"burnNodes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"burnUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"orderID\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"mode\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"buyGiftCard\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ownerAddress\",\"type\":\"address\"}],\"name\":\"checkNodes\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"tierIndex\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"createdTime\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"claimedTime\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"limitedTime\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"multiplier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"leftover\",\"type\":\"uint256\"}],\"internalType\":\"structNode[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimFee\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"claimable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tierName\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"count\",\"type\":\"uint32\"}],\"name\":\"compound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"countOfTier\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"countOfUser\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"countTotal\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tierName\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"count\",\"type\":\"uint32\"}],\"name\":\"create\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"discountPer10\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"getAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastTimestampClaim\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"managers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mantPercent\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxCountOfUser\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxMonthValue\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"migrateNodesFromOldVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"migrateRewardsFromOldVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"string\",\"name\":\"tierName\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"count\",\"type\":\"uint32\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nftAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"nodes\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"tierIndex\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"createdTime\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"claimedTime\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"limitedTime\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"multiplier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"leftover\",\"type\":\"uint256\"}],\"internalType\":\"structNode[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nodesOfUser\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nodesTotal\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"tierIndex\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"createdTime\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"claimedTime\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"limitedTime\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"multiplier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"leftover\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oldNodeManager\",\"outputs\":[{\"internalType\":\"contractINodeReward\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"oldRewardsOfUser\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operationsPoolAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operationsPoolFee\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"count\",\"type\":\"uint8\"},{\"internalType\":\"uint256[]\",\"name\":\"selected\",\"type\":\"uint256[]\"}],\"name\":\"pay\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"payInterval\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"rateClaimFee\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tierName\",\"type\":\"string\"}],\"name\":\"removeTier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewardMigrated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewardsOfUser\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardsPoolAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardsPoolFee\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardsTotal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sellPricePercent\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"setAddressInBlacklist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"value\",\"type\":\"uint32\"}],\"name\":\"setClaimFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"value\",\"type\":\"uint32\"}],\"name\":\"setDiscountPer10\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"val\",\"type\":\"bool\"}],\"name\":\"setDynamicClaimFeeEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"value\",\"type\":\"uint32\"}],\"name\":\"setMantPercent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_count\",\"type\":\"uint32\"}],\"name\":\"setMaxCountOfUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"value\",\"type\":\"uint8\"}],\"name\":\"setMaxMonthValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"}],\"name\":\"setNFTAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"setOperationsPoolAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"value\",\"type\":\"uint32\"}],\"name\":\"setOperationsPoolFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"value\",\"type\":\"uint32\"}],\"name\":\"setPayInterval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"setRewardsPoolAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"value\",\"type\":\"uint32\"}],\"name\":\"setRewardsPoolFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"setRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"setTokenAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"value\",\"type\":\"uint32\"}],\"name\":\"setTransferFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"value\",\"type\":\"uint32\"}],\"name\":\"setsellPricePercent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"tierMap\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tierTotal\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tiers\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardsPerTime\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"claimInterval\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"maintenanceFee\",\"type\":\"uint256\"}],\"internalType\":\"structTier[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tierName\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"count\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transferFee\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniswapV2Router\",\"outputs\":[{\"internalType\":\"contractIJoeRouter02\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpaidNodes\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"tierIndex\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"createdTime\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"claimedTime\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"limitedTime\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"multiplier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"leftover\",\"type\":\"uint256\"}],\"internalType\":\"structNode[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tierName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardsPerTime\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"claimInterval\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"maintenanceFee\",\"type\":\"uint256\"}],\"name\":\"updateTier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userMigrated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"anyToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawRate\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5061a2df80620000226000396000f3fe6080604052600436106104525760003560e01c8063721ffa941161023f578063acb2ad6f11610139578063d33953bb116100b6578063f2fde38b1161007a578063f2fde38b14611106578063f5278a8b1461112f578063f940e38514611158578063fba6c21f14611181578063fdff9b4d146111ac57610452565b8063d33953bb14611037578063dbe2b8ce14611062578063e0d128731461108b578063e985f283146110b4578063f2d049b6146110dd57610452565b8063c057e4f1116100fd578063c057e4f114610f54578063c0d7865514610f7d578063c63d63c014610fa6578063cce8505514610fcf578063d1b4359314610ffa57610452565b8063acb2ad6f14610e6d578063aea5c39414610e98578063b1ee5a8114610ec3578063b2dd3b3714610eec578063bb2d7f3a14610f2957610452565b80638c04a2e4116101c757806399d32fc41161018b57806399d32fc414610da95780639c43b6e214610dd45780639d76ea5814610dfd5780639ea164b314610e28578063a224cee714610e4457610452565b80638c04a2e414610cd65780638da5cb5b14610cff5780638e6a1bc214610d2a578063919df19814610d555780639746f9e814610d8057610452565b806379420fb11161020e57806379420fb114610c055780637e7b0d5714610c3057806381c31cf914610c5957806385b8eb4414610c82578063864bbf5914610cab57610452565b8063721ffa9414610b4457806374bf71c514610b8157806377ffb52b14610bc5578063784eceef14610bee57610452565b80633ba28fcb116103505780634ee571ce116102d85780635a64affb1161029c5780635a64affb14610a5f5780635bcb11bb14610a8a5780635bf8633a14610ab35780635c19521714610ade57806369d0373814610b1b57610452565b80634ee571ce1461096a578063501f947714610993578063530ee54b146109d057806355953d2514610a0d578063563e497c14610a3657610452565b8063450970b51161031f578063450970b5146108855780634a95d9d5146108c25780634c3f427d146108ed5780634cdae7121461092a5780634e71d92d1461095357610452565b80633ba28fcb146107c95780633fb53751146107f4578063402914f51461081d5780634506598d1461085a57610452565b806320b242ec116103de5780632d233503116103a25780632d233503146106f65780632dbc78e8146107215780632e9ef9761461074a5780633747bd451461077357806338cc48311461079e57610452565b806320b242ec146105fa57806322b4582214610616578063253ad3ba146106535780632655c1b91461069057806326a4e8d2146106cd57610452565b80630f81e539116104255780630f81e539146105015780631694505e1461052a578063189a5a17146105555780631c87cdb0146105925780631cdd3be3146105bd57610452565b806305ec8cec146104575780630a3eccf21461046e5780630c6afded146104995780630dcf1417146104d6575b600080fd5b34801561046357600080fd5b5061046c6111e9565b005b34801561047a57600080fd5b5061048361143b565b6040516104909190617e2d565b60405180910390f35b3480156104a557600080fd5b506104c060048036038101906104bb9190617eba565b611451565b6040516104cd9190617f02565b60405180910390f35b3480156104e257600080fd5b506104eb611471565b6040516104f89190617e2d565b60405180910390f35b34801561050d57600080fd5b5061052860048036038101906105239190617f49565b611487565b005b34801561053657600080fd5b5061053f61153b565b60405161054c9190617fd5565b60405180910390f35b34801561056157600080fd5b5061057c60048036038101906105779190617eba565b611561565b6040516105899190618195565b60405180910390f35b34801561059e57600080fd5b506105a7611896565b6040516105b49190617e2d565b60405180910390f35b3480156105c957600080fd5b506105e460048036038101906105df9190617eba565b6118ac565b6040516105f19190617f02565b60405180910390f35b610614600480360381019061060f9190618329565b6118cc565b005b34801561062257600080fd5b5061063d600480360381019061063891906183ac565b611cab565b60405161064a9190617e2d565b60405180910390f35b34801561065f57600080fd5b5061067a60048036038101906106759190617eba565b611ce4565b6040516106879190617f02565b60405180910390f35b34801561069c57600080fd5b506106b760048036038101906106b291906183f5565b611d04565b6040516106c49190618444565b60405180910390f35b3480156106d957600080fd5b506106f460048036038101906106ef9190617eba565b611d35565b005b34801561070257600080fd5b5061070b611e09565b604051610718919061846e565b60405180910390f35b34801561072d57600080fd5b5061074860048036038101906107439190618551565b611e1c565b005b34801561075657600080fd5b50610771600480360381019061076c9190617f49565b612274565b005b34801561077f57600080fd5b50610788612328565b6040516107959190617e2d565b60405180910390f35b3480156107aa57600080fd5b506107b361233e565b6040516107c091906185a9565b60405180910390f35b3480156107d557600080fd5b506107de612346565b6040516107eb9190617e2d565b60405180910390f35b34801561080057600080fd5b5061081b60048036038101906108169190617f49565b61235c565b005b34801561082957600080fd5b50610844600480360381019061083f9190617eba565b612410565b6040516108519190618444565b60405180910390f35b34801561086657600080fd5b5061086f612476565b60405161087c91906185e5565b60405180910390f35b34801561089157600080fd5b506108ac60048036038101906108a79190617eba565b61249c565b6040516108b99190618195565b60405180910390f35b3480156108ce57600080fd5b506108d76128e3565b6040516108e491906187d3565b60405180910390f35b3480156108f957600080fd5b50610914600480360381019061090f9190617eba565b612b10565b6040516109219190618444565b60405180910390f35b34801561093657600080fd5b50610951600480360381019061094c9190618821565b612b28565b005b34801561095f57600080fd5b50610968612bd5565b005b34801561097657600080fd5b50610991600480360381019061098c9190618911565b613113565b005b34801561099f57600080fd5b506109ba60048036038101906109b59190617eba565b613232565b6040516109c79190617e2d565b60405180910390f35b3480156109dc57600080fd5b506109f760048036038101906109f291906183ac565b61339b565b604051610a04919061846e565b60405180910390f35b348015610a1957600080fd5b50610a346004803603810190610a2f9190617eba565b6133d1565b005b348015610a4257600080fd5b50610a5d6004803603810190610a58919061899c565b6134a5565b005b348015610a6b57600080fd5b50610a74613d01565b604051610a819190617e2d565b60405180910390f35b348015610a9657600080fd5b50610ab16004803603810190610aac9190617eba565b613d17565b005b348015610abf57600080fd5b50610ac8614162565b604051610ad591906185a9565b60405180910390f35b348015610aea57600080fd5b50610b056004803603810190610b009190618a1f565b614188565b604051610b129190618444565b60405180910390f35b348015610b2757600080fd5b50610b426004803603810190610b3d9190617eba565b614453565b005b348015610b5057600080fd5b50610b6b6004803603810190610b669190617eba565b614527565b604051610b789190617e2d565b60405180910390f35b348015610b8d57600080fd5b50610ba86004803603810190610ba39190618a1f565b61454a565b604051610bbc989796959493929190618a4c565b60405180910390f35b348015610bd157600080fd5b50610bec6004803603810190610be79190618af6565b61460f565b005b348015610bfa57600080fd5b50610c036146bd565b005b348015610c1157600080fd5b50610c1a6148ed565b604051610c27919061846e565b60405180910390f35b348015610c3c57600080fd5b50610c576004803603810190610c529190618b23565b614900565b005b348015610c6557600080fd5b50610c806004803603810190610c7b9190617f49565b614ac4565b005b348015610c8e57600080fd5b50610ca96004803603810190610ca49190618b7f565b614b78565b005b348015610cb757600080fd5b50610cc0614e3d565b604051610ccd9190617e2d565b60405180910390f35b348015610ce257600080fd5b50610cfd6004803603810190610cf89190617f49565b614e53565b005b348015610d0b57600080fd5b50610d14614f07565b604051610d2191906185a9565b60405180910390f35b348015610d3657600080fd5b50610d3f614f2d565b604051610d4c9190617e2d565b60405180910390f35b348015610d6157600080fd5b50610d6a614f43565b604051610d7791906185a9565b60405180910390f35b348015610d8c57600080fd5b50610da76004803603810190610da29190617eba565b614f69565b005b348015610db557600080fd5b50610dbe61503d565b604051610dcb9190617e2d565b60405180910390f35b348015610de057600080fd5b50610dfb6004803603810190610df69190618b23565b615053565b005b348015610e0957600080fd5b50610e12615140565b604051610e1f91906185a9565b60405180910390f35b610e426004803603810190610e3d9190618cd9565b615166565b005b348015610e5057600080fd5b50610e6b6004803603810190610e669190618d35565b615796565b005b348015610e7957600080fd5b50610e82615c35565b604051610e8f9190617e2d565b60405180910390f35b348015610ea457600080fd5b50610ead615c4b565b604051610eba9190618d9f565b60405180910390f35b348015610ecf57600080fd5b50610eea6004803603810190610ee59190618dba565b615c71565b005b348015610ef857600080fd5b50610f136004803603810190610f0e9190617eba565b615d5c565b604051610f209190618444565b60405180910390f35b348015610f3557600080fd5b50610f3e615d74565b604051610f4b9190618444565b60405180910390f35b348015610f6057600080fd5b50610f7b6004803603810190610f769190617f49565b615d7a565b005b348015610f8957600080fd5b50610fa46004803603810190610f9f9190617eba565b615e2e565b005b348015610fb257600080fd5b50610fcd6004803603810190610fc89190617f49565b615f02565b005b348015610fdb57600080fd5b50610fe4615fb6565b604051610ff191906185a9565b60405180910390f35b34801561100657600080fd5b50611021600480360381019061101c9190617eba565b615fdc565b60405161102e9190617e2d565b60405180910390f35b34801561104357600080fd5b5061104c615fff565b6040516110599190617e2d565b60405180910390f35b34801561106e57600080fd5b5061108960048036038101906110849190618dfa565b616015565b005b34801561109757600080fd5b506110b260048036038101906110ad9190617f49565b6162f6565b005b3480156110c057600080fd5b506110db60048036038101906110d69190617eba565b6163aa565b005b3480156110e957600080fd5b5061110460048036038101906110ff91906183ac565b61647e565b005b34801561111257600080fd5b5061112d60048036038101906111289190617eba565b6165f2565b005b34801561113b57600080fd5b5061115660048036038101906111519190617f49565b616735565b005b34801561116457600080fd5b5061117f600480360381019061117a9190618ebf565b6167e9565b005b34801561118d57600080fd5b50611196616975565b6040516111a39190618195565b60405180910390f35b3480156111b857600080fd5b506111d360048036038101906111ce9190617eba565b616dce565b6040516111e09190617f02565b60405180910390f35b601060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16611275576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161126c90618f5c565b60405180910390fd5b601660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615611302576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112f990618fc8565b60405180910390fd5b601460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663619a635e336040518263ffffffff1660e01b815260040161135d91906185a9565b602060405180830381865afa15801561137a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061139e9190618ffd565b600e60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506001601660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550565b601160189054906101000a900463ffffffff1681565b60106020528060005260406000206000915054906101000a900460ff1681565b601160049054906101000a900463ffffffff1681565b3373ffffffffffffffffffffffffffffffffffffffff16601560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614611517576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161150e90619076565b60405180910390fd5b80601160006101000a81548163ffffffff021916908363ffffffff16021790555050565b601260049054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60606000600a60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900463ffffffff1663ffffffff1667ffffffffffffffff8111156115d4576115d36181d2565b5b60405190808252806020026020018201604052801561160d57816020015b6115fa617c69565b8152602001906001900390816115f25790505b5090506000600860008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090506000805b82805490508163ffffffff16101561188a576000838263ffffffff168154811061168457611683619096565b5b90600052602060002001549050600081111561187657600060076001836116ab91906190f4565b815481106116bc576116bb619096565b5b906000526020600020906004020190508773ffffffffffffffffffffffffffffffffffffffff168160000160059054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16036118745780604051806101000160405290816000820160009054906101000a900463ffffffff1663ffffffff1663ffffffff1681526020016000820160049054906101000a900460ff1660ff1660ff1681526020016000820160059054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016000820160199054906101000a900463ffffffff1663ffffffff1663ffffffff1681526020016001820160009054906101000a900463ffffffff1663ffffffff1663ffffffff1681526020016001820160049054906101000a900463ffffffff1663ffffffff1663ffffffff16815260200160028201548152602001600382015481525050868563ffffffff168151811061186857611867619096565b5b60200260200101819052505b505b50808061188290619128565b915050611657565b50829350505050919050565b601160009054906101000a900463ffffffff1681565b600f6020528060005260406000206000915054906101000a900460ff1681565b73d00ae08403b9bbb9124bb305c09058e32c39a48c73ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161415801561191c5750600082145b156119c9578373ffffffffffffffffffffffffffffffffffffffff166323b872dd33600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16846040518463ffffffff1660e01b815260040161198093929190619154565b6020604051808303816000875af115801561199f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119c391906191a0565b50611c68565b73d00ae08403b9bbb9124bb305c09058e32c39a48c73ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1603611ac15780341015611a53576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a4a90619219565b60405180910390fd5b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc349081150290604051600060405180830381858888f19350505050158015611abb573d6000803e3d6000fd5b50611c67565b611ac9616dee565b80600e60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015611b4b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611b4290619285565b60405180910390fd5b80600e60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254611b9a91906190f4565b925050819055508373ffffffffffffffffffffffffffffffffffffffff166323b872dd600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16846040518463ffffffff1660e01b8152600401611c2293929190619154565b6020604051808303816000875af1158015611c41573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c6591906191a0565b505b5b7f34618306376b53644b3c2e2e35fc793e80e74a56c2c29194492397a49ef0227733858584604051611c9d94939291906192de565b60405180910390a150505050565b600b818051602081018201805184825260208301602085012081835280955050505050506000915054906101000a900463ffffffff1681565b60166020528060005260406000206000915054906101000a900460ff1681565b60086020528160005260406000208181548110611d2057600080fd5b90600052602060002001600091509150505481565b3373ffffffffffffffffffffffffffffffffffffffff16601560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614611dc5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611dbc90619076565b60405180910390fd5b80600060026101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600660019054906101000a900460ff1681565b3373ffffffffffffffffffffffffffffffffffffffff16601560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614611eac576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611ea390619076565b60405180910390fd5b6000805b82518163ffffffff161015612230576000838263ffffffff1681518110611eda57611ed9619096565b5b602002602001015163ffffffff169050600081111561221c5760006007600183611f0491906190f4565b81548110611f1557611f14619096565b5b90600052602060002090600402019050600073ffffffffffffffffffffffffffffffffffffffff168160000160059054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461221a576000600860008360000160059054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905060005b81805490508163ffffffff16101561206457818163ffffffff168154811061201257612011619096565b5b90600052602060002001548403612051576000828263ffffffff168154811061203e5761203d619096565b5b9060005260206000200181905550612064565b808061205c90619128565b915050611fe7565b50600a60008360000160059054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600081819054906101000a900463ffffffff16809291906120e89061932a565b91906101000a81548163ffffffff021916908363ffffffff1602179055505060008260000160056101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060008260010160006101000a81548163ffffffff021916908363ffffffff160217905550600060048360000160049054906101000a900460ff1660ff168154811061219957612198619096565b5b90600052602060002090600602019050600b816001016040516121bc9190619452565b9081526020016040518091039020600081819054906101000a900463ffffffff16809291906121ea9061932a565b91906101000a81548163ffffffff021916908363ffffffff16021790555050858061221490619128565b96505050505b505b50808061222890619128565b915050611eb0565b5080600960008282829054906101000a900463ffffffff166122529190619469565b92506101000a81548163ffffffff021916908363ffffffff1602179055505050565b3373ffffffffffffffffffffffffffffffffffffffff16601560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614612304576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016122fb90619076565b60405180910390fd5b80601160086101000a81548163ffffffff021916908363ffffffff16021790555050565b601760009054906101000a900463ffffffff1681565b600030905090565b601160149054906101000a900463ffffffff1681565b3373ffffffffffffffffffffffffffffffffffffffff16601560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146123ec576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016123e390619076565b60405180910390fd5b806011600c6101000a81548163ffffffff021916908363ffffffff16021790555050565b60008061241f83600080616fb3565b50509050600e60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548161246e919061949d565b915050919050565b601360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60606000805b6007805490508163ffffffff1610156125e457600060078263ffffffff16815481106124d1576124d0619096565b5b90600052602060002090600402019050600073ffffffffffffffffffffffffffffffffffffffff168160000160059054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415801561259157508473ffffffffffffffffffffffffffffffffffffffff168160000160059054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16145b80156125bc57504263ffffffff168160010160049054906101000a900463ffffffff1663ffffffff16105b156125d05782806125cc90619128565b9350505b5080806125dc90619128565b9150506124a2565b5060008163ffffffff1667ffffffffffffffff811115612607576126066181d2565b5b60405190808252806020026020018201604052801561264057816020015b61262d617c69565b8152602001906001900390816126255790505b5090506000805b6007805490508163ffffffff1610156128d757600060078263ffffffff168154811061267657612675619096565b5b90600052602060002090600402019050600073ffffffffffffffffffffffffffffffffffffffff168160000160059054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415801561273657508673ffffffffffffffffffffffffffffffffffffffff168160000160059054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16145b801561276157504263ffffffff168160010160049054906101000a900463ffffffff1663ffffffff16105b156128c35780604051806101000160405290816000820160009054906101000a900463ffffffff1663ffffffff1663ffffffff1681526020016000820160049054906101000a900460ff1660ff1660ff1681526020016000820160059054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016000820160199054906101000a900463ffffffff1663ffffffff1663ffffffff1681526020016001820160009054906101000a900463ffffffff1663ffffffff1663ffffffff1681526020016001820160049054906101000a900463ffffffff1663ffffffff1663ffffffff1681526020016002820154815260200160038201548152505084848061289e90619128565b955063ffffffff16815181106128b7576128b6619096565b5b60200260200101819052505b5080806128cf90619128565b915050612647565b50819350505050919050565b60606000600660009054906101000a900460ff1660ff1667ffffffffffffffff811115612913576129126181d2565b5b60405190808252806020026020018201604052801561294c57816020015b612939617cdf565b8152602001906001900390816129315790505b5090506000805b6004805490508160ff161015612b0757600060048260ff168154811061297c5761297b619096565b5b9060005260206000209060060201905060006005826001016040516129a19190619452565b908152602001604051809103902060009054906101000a900460ff1660ff161115612af357806040518060c00160405290816000820160009054906101000a900460ff1660ff1660ff1681526020016001820180546129ff90619382565b80601f0160208091040260200160405190810160405280929190818152602001828054612a2b90619382565b8015612a785780601f10612a4d57610100808354040283529160200191612a78565b820191906000526020600020905b815481529060010190602001808311612a5b57829003601f168201915b5050505050815260200160028201548152602001600382015481526020016004820160009054906101000a900463ffffffff1663ffffffff1663ffffffff168152602001600582015481525050848480612ad1906194f3565b955060ff1681518110612ae757612ae6619096565b5b60200260200101819052505b508080612aff906194f3565b915050612953565b50819250505090565b600e6020528060005260406000206000915090505481565b3373ffffffffffffffffffffffffffffffffffffffff16601560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614612bb8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612baf90619076565b60405180910390fd5b80601a60006101000a81548160ff02191690831515021790555050565b612bdd616dee565b6000612be833613232565b9050601a60009054906101000a900460ff168015612c0c575060008163ffffffff16115b15612db65760006127108263ffffffff16600e60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054612c67919061951c565b612c7191906195a5565b905080600e60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254612cc291906190f4565b92505081905550600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16846040518463ffffffff1660e01b8152600401612d6c93929190619154565b6020604051808303816000875af1158015612d8b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612daf91906191a0565b5050612f69565b6000612710601160109054906101000a900463ffffffff1663ffffffff16600e60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054612e1e919061951c565b612e2891906195a5565b905080600e60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254612e7991906190f4565b92505081905550600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16846040518463ffffffff1660e01b8152600401612f2393929190619154565b6020604051808303816000875af1158015612f42573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612f6691906191a0565b50505b600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1633600e60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546040518463ffffffff1660e01b815260040161302993929190619154565b6020604051808303816000875af1158015613048573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061306c91906191a0565b506000600e60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555042601960003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548163ffffffff021916908363ffffffff16021790555050565b3373ffffffffffffffffffffffffffffffffffffffff16601560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146131a3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161319a90619076565b60405180910390fd5b60008351116131e7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016131de90619622565b60405180910390fd5b60005b835181101561322c5761321884828151811061320957613208619096565b5b60200260200101518484617366565b50808061322490619642565b9150506131ea565b50505050565b600080601960008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900463ffffffff1663ffffffff16036132aa57601160109054906101000a900463ffffffff169050613396565b6000601960008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900463ffffffff16426133079190619469565b9050620697808163ffffffff161161332457611388915050613396565b620e80808163ffffffff161161333f57610fa0915050613396565b621669808163ffffffff161161335a57610bb8915050613396565b621fa4008163ffffffff1611613375576107d0915050613396565b62278d008163ffffffff1611613390576103e8915050613396565b60649150505b919050565b6005818051602081018201805184825260208301602085012081835280955050505050506000915054906101000a900460ff1681565b3373ffffffffffffffffffffffffffffffffffffffff16601560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614613461576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161345890619076565b60405180910390fd5b80601360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b3373ffffffffffffffffffffffffffffffffffffffff16601560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614613535576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161352c90619076565b60405180910390fd5b600f60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615806135d85750600f60008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16155b613617576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161360e906196d6565b60405180910390fd5b60006005856040516136299190619727565b908152602001604051809103902060009054906101000a900460ff169050601160189054906101000a900463ffffffff1663ffffffff1684600a60008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900463ffffffff166136bb919061973e565b63ffffffff161115613702576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016136f9906197c4565b60405180910390fd5b6000600460018361371391906197e4565b60ff168154811061372757613726619096565b5b906000526020600020906006020190506000600860008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090506000600860008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905060008060005b84805490508163ffffffff161015613a35576000858263ffffffff16815481106137f0576137ef619096565b5b906000526020600020015490506000811115613a21576000600760018361381791906190f4565b8154811061382857613827619096565b5b906000526020600020906004020190508a73ffffffffffffffffffffffffffffffffffffffff168160000160059054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161480156138b957508060000160049054906101000a900460ff1660ff1660018a6138b491906197e4565b60ff16145b15613a1f57898160000160056101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060008860040160009054906101000a900463ffffffff1663ffffffff1689600301548360010160009054906101000a900463ffffffff1663ffffffff164261394991906190f4565b613953919061951c565b61395d91906195a5565b9050670de0b6b3a76400008161397391906195a5565b8561397e919061949d565b9450428260010160006101000a81548163ffffffff021916908363ffffffff16021790555085806139ae90619128565b965050868390806001815401808255809150506001900390600052602060002001600090919091909150556000888563ffffffff16815481106139f4576139f3619096565b5b90600052602060002001819055508c63ffffffff168663ffffffff1603613a1d57505050613a35565b505b505b508080613a2d90619128565b9150506137c3565b508863ffffffff168263ffffffff1614613a84576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613a7b90619864565b60405180910390fd5b88600a60008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282829054906101000a900463ffffffff16613ae29190619469565b92506101000a81548163ffffffff021916908363ffffffff16021790555088600a60008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282829054906101000a900463ffffffff16613b5e919061973e565b92506101000a81548163ffffffff021916908363ffffffff1602179055506000811115613cba5780600d60008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254613bd4919061949d565b9250508190555080600c6000828254613bed919061949d565b92505081905550600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168a846040518463ffffffff1660e01b8152600401613c7593929190619154565b6020604051808303816000875af1158015613c94573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613cb891906191a0565b505b7f8a091ff4fd9357c3a846a822cac198fd60c380e70afd0af7e2451a3953727e0388888b604051613ced93929190619884565b60405180910390a150505050505050505050565b601260009054906101000a900463ffffffff1681565b3373ffffffffffffffffffffffffffffffffffffffff16601560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614613da7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613d9e90619076565b60405180910390fd5b6000600860008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905060005b81805490508163ffffffff161015613fd3576000828263ffffffff1681548110613e1a57613e19619096565b5b906000526020600020015490506000811115613fbf5760006007600183613e4191906190f4565b81548110613e5257613e51619096565b5b906000526020600020906004020190508473ffffffffffffffffffffffffffffffffffffffff168160000160059054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603613fbd5760008160000160056101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060008160010160006101000a81548163ffffffff021916908363ffffffff160217905550600060048260000160049054906101000a900460ff1660ff1681548110613f4b57613f4a619096565b5b90600052602060002090600602019050600b81600101604051613f6e9190619452565b9081526020016040518091039020600081819054906101000a900463ffffffff1680929190613f9c9061932a565b91906101000a81548163ffffffff021916908363ffffffff16021790555050505b505b508080613fcb90619128565b915050613ded565b50600067ffffffffffffffff811115613fef57613fee6181d2565b5b60405190808252806020026020018201604052801561401d5781602001602082028036833780820191505090505b50600860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209080519060200190614070929190617d1e565b50600a60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900463ffffffff16600960008282829054906101000a900463ffffffff166140e19190619469565b92506101000a81548163ffffffff021916908363ffffffff1602179055506000600a60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548163ffffffff021916908363ffffffff1602179055505050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008073ffffffffffffffffffffffffffffffffffffffff16601260049054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16036141e8576000905061444e565b6000600267ffffffffffffffff811115614205576142046181d2565b5b6040519080825280602002602001820160405280156142335781602001602082028036833780820191505090505b509050600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff168160008151811061426d5761426c619096565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050601260049054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166373b295c26040518163ffffffff1660e01b8152600401602060405180830381865afa158015614314573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061433891906198d0565b8160018151811061434c5761434b619096565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250506000601260049054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d06ca61f85846040518363ffffffff1660e01b81526004016143e59291906199ac565b600060405180830381865afa158015614402573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019061442b9190619a73565b90508060018151811061444157614440619096565b5b6020026020010151925050505b919050565b3373ffffffffffffffffffffffffffffffffffffffff16601560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146144e3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016144da90619076565b60405180910390fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60196020528060005260406000206000915054906101000a900463ffffffff1681565b6007818154811061455a57600080fd5b90600052602060002090600402016000915090508060000160009054906101000a900463ffffffff16908060000160049054906101000a900460ff16908060000160059054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060000160199054906101000a900463ffffffff16908060010160009054906101000a900463ffffffff16908060010160049054906101000a900463ffffffff16908060020154908060030154905088565b3373ffffffffffffffffffffffffffffffffffffffff16601560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461469f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161469690619076565b60405180910390fd5b80600660016101000a81548160ff021916908360ff16021790555050565b601060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161561474a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161474190619b08565b60405180910390fd5b6000601460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634491a7e4336040518263ffffffff1660e01b81526004016147a791906185a9565b602060405180830381865afa1580156147c4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906147e89190618ffd565b905060008163ffffffff161415801561484b5750601060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16155b156148ea576001601060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506148e8336040518060400160405280600581526020017f626173696300000000000000000000000000000000000000000000000000000081525083617366565b505b50565b600660009054906101000a900460ff1681565b600061490d338484617366565b9050614917616dee565b80600e60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015614999576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161499090619285565b60405180910390fd5b80600e60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546149e891906190f4565b925050819055507f3a823590161fd78e3988b3da59eb891df9f43272cf652f586d2e79f65f6f3571338484600960009054906101000a900463ffffffff16600a60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900463ffffffff16600b89604051614a869190619727565b908152602001604051809103902060009054906101000a900463ffffffff16604051614ab796959493929190619b28565b60405180910390a1505050565b3373ffffffffffffffffffffffffffffffffffffffff16601560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614614b54576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614b4b90619076565b60405180910390fd5b80601160106101000a81548163ffffffff021916908363ffffffff16021790555050565b3373ffffffffffffffffffffffffffffffffffffffff16601560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614614c08576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614bff90619076565b60405180910390fd5b60008411614c4b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614c4290619bdc565b60405180910390fd5b60008311614c8e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614c8590619c48565b60405180910390fd5b60008263ffffffff1611614cd7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614cce90619cb4565b60405180910390fd5b60046040518060c0016040528060048054905060ff1681526020018781526020018681526020018581526020018463ffffffff16815260200183815250908060018154018082558091505060019003906000526020600020906006020160009091909190915060008201518160000160006101000a81548160ff021916908360ff1602179055506020820151816001019080519060200190614d7a929190617d6b565b50604082015181600201556060820151816003015560808201518160040160006101000a81548163ffffffff021916908363ffffffff16021790555060a082015181600501555050600480549050600586604051614dd89190619727565b908152602001604051809103902060006101000a81548160ff021916908360ff1602179055506006600081819054906101000a900460ff1680929190614e1d906194f3565b91906101000a81548160ff021916908360ff160217905550505050505050565b6011600c9054906101000a900463ffffffff1681565b3373ffffffffffffffffffffffffffffffffffffffff16601560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614614ee3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614eda90619076565b60405180910390fd5b806011601c6101000a81548163ffffffff021916908363ffffffff16021790555050565b601560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6011601c9054906101000a900463ffffffff1681565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b3373ffffffffffffffffffffffffffffffffffffffff16601560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614614ff9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614ff090619076565b60405180910390fd5b80600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b601160109054906101000a900463ffffffff1681565b6000615060338484617366565b905061506b81617969565b7f3a823590161fd78e3988b3da59eb891df9f43272cf652f586d2e79f65f6f3571338484600960009054906101000a900463ffffffff16600a60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900463ffffffff16600b896040516151029190619727565b908152602001604051809103902060009054906101000a900463ffffffff1660405161513396959493929190619b28565b60405180910390a1505050565b600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008260ff1611801561518e5750600660019054906101000a900460ff1660ff168260ff1611155b6151cd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016151c490619d20565b60405180910390fd5b600080825103615490576000600860003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905060005b81805490508163ffffffff161015615489576000828263ffffffff168154811061524a57615249619096565b5b906000526020600020015490506000811115615475576000600760018361527191906190f4565b8154811061528257615281619096565b5b906000526020600020906004020190503373ffffffffffffffffffffffffffffffffffffffff168160000160059054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff160361547357600060048260000160049054906101000a900460ff1660ff168154811061531357615312619096565b5b906000526020600020906006020190506011601c9054906101000a900463ffffffff168860ff166153449190619d40565b8260010160048282829054906101000a900463ffffffff16615366919061973e565b92506101000a81548163ffffffff021916908363ffffffff16021790555060006011601c9054906101000a900463ffffffff1663ffffffff16428460010160049054906101000a900463ffffffff1663ffffffff166153c591906190f4565b6153cf91906195a5565b905060038160ff161115615418576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161540f90619dca565b60405180910390fd5b6154638960ff16612710601760009054906101000a900463ffffffff1663ffffffff16856002015461544a919061951c565b61545491906195a5565b61545e919061951c565b614188565b8761546e919061949d565b965050505b505b50808061548190619128565b91505061521d565b50506156e5565b60005b82518163ffffffff1610156156e3576000838263ffffffff16815181106154bd576154bc619096565b5b602002602001015190506000600782815481106154dd576154dc619096565b5b906000526020600020906004020190503373ffffffffffffffffffffffffffffffffffffffff168160000160059054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16036156ce57600060048260000160049054906101000a900460ff1660ff168154811061556e5761556d619096565b5b906000526020600020906006020190506011601c9054906101000a900463ffffffff168760ff1661559f9190619d40565b8260010160048282829054906101000a900463ffffffff166155c1919061973e565b92506101000a81548163ffffffff021916908363ffffffff16021790555060006011601c9054906101000a900463ffffffff1663ffffffff16428460010160049054906101000a900463ffffffff1663ffffffff1661562091906190f4565b61562a91906195a5565b905060038160ff161115615673576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161566a90619dca565b60405180910390fd5b6156be8860ff16612710601760009054906101000a900463ffffffff1663ffffffff1685600201546156a5919061951c565b6156af91906195a5565b6156b9919061951c565b614188565b866156c9919061949d565b955050505b505080806156db90619128565b915050615493565b505b34811115615728576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161571f90619e36565b60405180910390fd5b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f19350505050158015615790573d6000803e3d6000fd5b50505050565b60006157a26001617b56565b905080156157c6576001600060016101000a81548160ff0219169083151502179055505b816000815181106157da576157d9619096565b5b6020026020010151600060026101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508160018151811061583657615835619096565b5b6020026020010151600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508160028151811061589257615891619096565b5b6020026020010151600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555033601560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506159776040518060400160405280600581526020017f6261736963000000000000000000000000000000000000000000000000000000815250678ac7230489e8000067025bf6196bd100006201518066038d7ea4c68000614b78565b6159d46040518060400160405280600581526020017f6c696768740000000000000000000000000000000000000000000000000000008152506802b5e3af16b1880000670b1a2bc2ec500000620151806601c6bf52634000614b78565b615a306040518060400160405280600381526020017f70726f000000000000000000000000000000000000000000000000000000000081525068056bc75e2d63100000671bc16d674ec8000062015180655af3107a4000614b78565b600a601160006101000a81548163ffffffff021916908363ffffffff1602179055506000601160046101000a81548163ffffffff021916908363ffffffff1602179055506000601160086101000a81548163ffffffff021916908363ffffffff160217905550611f406011600c6101000a81548163ffffffff021916908363ffffffff1602179055506107d0601160146101000a81548163ffffffff021916908363ffffffff1602179055506103e8601160106101000a81548163ffffffff021916908363ffffffff1602179055506064601160186101000a81548163ffffffff021916908363ffffffff1602179055506019601260006101000a81548163ffffffff021916908363ffffffff16021790555062278d006011601c6101000a81548163ffffffff021916908363ffffffff1602179055506003600660016101000a81548160ff021916908360ff1602179055507305c88f67fa0711b3a76ada2b6f0a2d3a54fc775c601460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508015615c315760008060016101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024986001604051615c289190619e91565b60405180910390a15b5050565b601160089054906101000a900463ffffffff1681565b601460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b3373ffffffffffffffffffffffffffffffffffffffff16601560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614615d01576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401615cf890619076565b60405180910390fd5b80600f60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b600d6020528060005260406000206000915090505481565b600c5481565b3373ffffffffffffffffffffffffffffffffffffffff16601560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614615e0a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401615e0190619076565b60405180910390fd5b80601760006101000a81548163ffffffff021916908363ffffffff16021790555050565b3373ffffffffffffffffffffffffffffffffffffffff16601560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614615ebe576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401615eb590619076565b60405180910390fd5b80601260046101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b3373ffffffffffffffffffffffffffffffffffffffff16601560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614615f92576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401615f8990619076565b60405180910390fd5b80601160186101000a81548163ffffffff021916908363ffffffff16021790555050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600a6020528060005260406000206000915054906101000a900463ffffffff1681565b600960009054906101000a900463ffffffff1681565b3373ffffffffffffffffffffffffffffffffffffffff16601560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146160a5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161609c90619076565b60405180910390fd5b60006005876040516160b79190619727565b908152602001604051809103902060009054906101000a900460ff16905060008160ff161161611b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161611290619ef8565b60405180910390fd5b858051906020012087805190602001200361616b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161616290619f64565b60405180910390fd5b600085116161ae576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016161a590619bdc565b60405180910390fd5b600084116161f1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016161e890619fd0565b60405180910390fd5b6000600460018361620291906197e4565b60ff168154811061621657616215619096565b5b906000526020600020906006020190508681600101908051906020019061623e929190617d6b565b50858160020181905550848160030181905550838160040160006101000a81548163ffffffff021916908363ffffffff160217905550828160050181905550600060058960405161628f9190619727565b908152602001604051809103902060006101000a81548160ff021916908360ff160217905550816005886040516162c69190619727565b908152602001604051809103902060006101000a81548160ff021916908360ff1602179055505050505050505050565b3373ffffffffffffffffffffffffffffffffffffffff16601560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614616386576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161637d90619076565b60405180910390fd5b80601260006101000a81548163ffffffff021916908363ffffffff16021790555050565b3373ffffffffffffffffffffffffffffffffffffffff16601560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461643a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161643190619076565b60405180910390fd5b80600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b3373ffffffffffffffffffffffffffffffffffffffff16601560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461650e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161650590619076565b60405180910390fd5b60006005826040516165209190619727565b908152602001604051809103902060009054906101000a900460ff1660ff161161657f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016165769061a03c565b60405180910390fd5b60006005826040516165919190619727565b908152602001604051809103902060006101000a81548160ff021916908360ff1602179055506006600081819054906101000a900460ff16809291906165d69061a05c565b91906101000a81548160ff021916908360ff1602179055505050565b3373ffffffffffffffffffffffffffffffffffffffff16601560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614616682576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161667990619076565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036166f1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016166e89061a0d1565b60405180910390fd5b80601560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b3373ffffffffffffffffffffffffffffffffffffffff16601560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146167c5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016167bc90619076565b60405180910390fd5b80601160146101000a81548163ffffffff021916908363ffffffff16021790555050565b3373ffffffffffffffffffffffffffffffffffffffff16601560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614616879576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161687090619076565b60405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff1663a9059cbb828473ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b81526004016168cf91906185a9565b602060405180830381865afa1580156168ec573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906169109190618ffd565b6040518363ffffffff1660e01b815260040161692d92919061a0f1565b6020604051808303816000875af115801561694c573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061697091906191a0565b505050565b60603373ffffffffffffffffffffffffffffffffffffffff16601560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614616a07576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016169fe90619076565b60405180910390fd5b6000805b6007805490508163ffffffff161015616b0f57600060078263ffffffff1681548110616a3a57616a39619096565b5b90600052602060002090600402019050600073ffffffffffffffffffffffffffffffffffffffff168160000160059054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614158015616ae757504263ffffffff166011601c9054906101000a900463ffffffff168260010160049054906101000a900463ffffffff16616adf919061973e565b63ffffffff16105b15616afb578280616af790619128565b9350505b508080616b0790619128565b915050616a0b565b5060008163ffffffff1667ffffffffffffffff811115616b3257616b316181d2565b5b604051908082528060200260200182016040528015616b6b57816020015b616b58617c69565b815260200190600190039081616b505790505b5090506000805b6007805490508163ffffffff161015616dc457600060078263ffffffff1681548110616ba157616ba0619096565b5b90600052602060002090600402019050600073ffffffffffffffffffffffffffffffffffffffff168160000160059054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614158015616c4e57504263ffffffff166011601c9054906101000a900463ffffffff168260010160049054906101000a900463ffffffff16616c46919061973e565b63ffffffff16105b15616db05780604051806101000160405290816000820160009054906101000a900463ffffffff1663ffffffff1663ffffffff1681526020016000820160049054906101000a900460ff1660ff1660ff1681526020016000820160059054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016000820160199054906101000a900463ffffffff1663ffffffff1663ffffffff1681526020016001820160009054906101000a900463ffffffff1663ffffffff1663ffffffff1681526020016001820160049054906101000a900463ffffffff1663ffffffff1663ffffffff16815260200160028201548152602001600382015481525050848480616d8b90619128565b955063ffffffff1681518110616da457616da3619096565b5b60200260200101819052505b508080616dbc90619128565b915050616b72565b5081935050505090565b60186020528060005260406000206000915054906101000a900460ff1681565b6000806000616dff33600080616fb3565b9250925092506000831115616f075782600d60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054616e59919061949d565b600d60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555082600c54616eaa919061949d565b600c8190555082600e60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254616eff919061949d565b925050819055505b60005b8263ffffffff168163ffffffff161015616fad576000828263ffffffff1681518110616f3957616f38619096565b5b6020026020010151905060006007600183616f5491906190f4565b81548110616f6557616f64619096565b5b90600052602060002090600402019050428160010160006101000a81548163ffffffff021916908363ffffffff16021790555050508080616fa590619128565b915050616f0a565b50505050565b60008060606000806000600860008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002080549050905060008167ffffffffffffffff81111561701e5761701d6181d2565b5b60405190808252806020026020018201604052801561704c5781602001602082028036833780820191505090505b50905060005b828163ffffffff16101561734f576000600860008d73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208263ffffffff16815481106170b9576170b8619096565b5b90600052602060002001549050600081111561733a57600060076001836170e091906190f4565b815481106170f1576170f0619096565b5b906000526020600020906004020160000160059054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690506000600760018461713691906190f4565b8154811061714757617146619096565b5b906000526020600020906004020160000160049054906101000a900460ff1690506000600760018561717991906190f4565b8154811061718a57617189619096565b5b906000526020600020906004020160010160009054906101000a900463ffffffff16905060008e60ff16141580156171d4575060018e6171ca91906197e4565b60ff168260ff1614155b156171e2575050505061733c565b8e73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff160361733657600060048360ff168154811061722e5761722d619096565b5b9060005260206000209060060201600301549050600060048460ff168154811061725b5761725a619096565b5b906000526020600020906006020160040160009054906101000a900463ffffffff1663ffffffff1690508981670de0b6b3a7640000848663ffffffff16426172a391906190f4565b6172ad919061951c565b6172b791906195a5565b6172c191906195a5565b6172cb919061949d565b995085888c63ffffffff16815181106172e7576172e6619096565b5b6020026020010181815250508a806172fe90619128565b9b505060008f63ffffffff161415801561732357508e63ffffffff168b63ffffffff16145b156173335750505050505061734f565b50505b5050505b505b808061734790619128565b915050617052565b508284829650965096505050505093509350939050565b6000600f60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16156173f5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016173ec9061a166565b60405180910390fd5b601560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614806174c95750601160189054906101000a900463ffffffff1663ffffffff1682600a60008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900463ffffffff166174c0919061973e565b63ffffffff1611155b617508576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016174ff9061a1d2565b60405180910390fd5b6000600160058560405161751c9190619727565b908152602001604051809103902060009054906101000a900460ff1661754291906197e4565b9050600060048260ff168154811061755d5761755c619096565b5b9060005260206000209060060201600201549050600042905060005b8563ffffffff168163ffffffff1610156177df57600760405180610100016040528060078054905063ffffffff1681526020018660ff1681526020018a73ffffffffffffffffffffffffffffffffffffffff1681526020018463ffffffff1681526020018463ffffffff1681526020016011601c9054906101000a900463ffffffff1685617607919061973e565b63ffffffff168152602001600081526020016000815250908060018154018082558091505060019003906000526020600020906004020160009091909190915060008201518160000160006101000a81548163ffffffff021916908363ffffffff16021790555060208201518160000160046101000a81548160ff021916908360ff16021790555060408201518160000160056101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160000160196101000a81548163ffffffff021916908363ffffffff16021790555060808201518160010160006101000a81548163ffffffff021916908363ffffffff16021790555060a08201518160010160046101000a81548163ffffffff021916908363ffffffff16021790555060c0820151816002015560e082015181600301555050600860008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600780549050908060018154018082558091505060019003906000526020600020016000909190919091505580806177d790619128565b915050617579565b5084600a60008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282829054906101000a900463ffffffff1661783e919061973e565b92506101000a81548163ffffffff021916908363ffffffff16021790555084600b8760405161786d9190619727565b908152602001604051809103902060008282829054906101000a900463ffffffff16617899919061973e565b92506101000a81548163ffffffff021916908363ffffffff16021790555084600960008282829054906101000a900463ffffffff166178d8919061973e565b92506101000a81548163ffffffff021916908363ffffffff16021790555060008563ffffffff168361790a919061951c565b9050600a8663ffffffff161061795b57612710601160009054906101000a900463ffffffff1661271061793d9190619469565b63ffffffff168261794e919061951c565b61795891906195a5565b90505b809450505050509392505050565b600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd33600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166127106011600c9054906101000a900463ffffffff1663ffffffff16866179f2919061951c565b6179fc91906195a5565b6040518463ffffffff1660e01b8152600401617a1a93929190619154565b6020604051808303816000875af1158015617a39573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190617a5d91906191a0565b50600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd33600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16612710601160149054906101000a900463ffffffff1663ffffffff1686617ae7919061951c565b617af191906195a5565b6040518463ffffffff1660e01b8152600401617b0f93929190619154565b6020604051808303816000875af1158015617b2e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190617b5291906191a0565b5050565b60008060019054906101000a900460ff1615617bcd5760018260ff16148015617b855750617b8330617c46565b155b617bc4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401617bbb9061a264565b60405180910390fd5b60009050617c41565b8160ff1660008054906101000a900460ff1660ff1610617c22576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401617c199061a264565b60405180910390fd5b816000806101000a81548160ff021916908360ff160217905550600190505b919050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b604051806101000160405280600063ffffffff168152602001600060ff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600063ffffffff168152602001600063ffffffff168152602001600063ffffffff16815260200160008152602001600081525090565b6040518060c00160405280600060ff168152602001606081526020016000815260200160008152602001600063ffffffff168152602001600081525090565b828054828255906000526020600020908101928215617d5a579160200282015b82811115617d59578251825591602001919060010190617d3e565b5b509050617d679190617df1565b5090565b828054617d7790619382565b90600052602060002090601f016020900481019282617d995760008555617de0565b82601f10617db257805160ff1916838001178555617de0565b82800160010185558215617de0579182015b82811115617ddf578251825591602001919060010190617dc4565b5b509050617ded9190617df1565b5090565b5b80821115617e0a576000816000905550600101617df2565b5090565b600063ffffffff82169050919050565b617e2781617e0e565b82525050565b6000602082019050617e426000830184617e1e565b92915050565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000617e8782617e5c565b9050919050565b617e9781617e7c565b8114617ea257600080fd5b50565b600081359050617eb481617e8e565b92915050565b600060208284031215617ed057617ecf617e52565b5b6000617ede84828501617ea5565b91505092915050565b60008115159050919050565b617efc81617ee7565b82525050565b6000602082019050617f176000830184617ef3565b92915050565b617f2681617e0e565b8114617f3157600080fd5b50565b600081359050617f4381617f1d565b92915050565b600060208284031215617f5f57617f5e617e52565b5b6000617f6d84828501617f34565b91505092915050565b6000819050919050565b6000617f9b617f96617f9184617e5c565b617f76565b617e5c565b9050919050565b6000617fad82617f80565b9050919050565b6000617fbf82617fa2565b9050919050565b617fcf81617fb4565b82525050565b6000602082019050617fea6000830184617fc6565b92915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b61802581617e0e565b82525050565b600060ff82169050919050565b6180418161802b565b82525050565b61805081617e7c565b82525050565b6000819050919050565b61806981618056565b82525050565b61010082016000820151618086600085018261801c565b5060208201516180996020850182618038565b5060408201516180ac6040850182618047565b5060608201516180bf606085018261801c565b5060808201516180d2608085018261801c565b5060a08201516180e560a085018261801c565b5060c08201516180f860c0850182618060565b5060e082015161810b60e0850182618060565b50505050565b600061811d838361806f565b6101008301905092915050565b6000602082019050919050565b600061814282617ff0565b61814c8185617ffb565b93506181578361800c565b8060005b8381101561818857815161816f8882618111565b975061817a8361812a565b92505060018101905061815b565b5085935050505092915050565b600060208201905081810360008301526181af8184618137565b905092915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61820a826181c1565b810181811067ffffffffffffffff82111715618229576182286181d2565b5b80604052505050565b600061823c617e48565b90506182488282618201565b919050565b600067ffffffffffffffff821115618268576182676181d2565b5b618271826181c1565b9050602081019050919050565b82818337600083830152505050565b60006182a061829b8461824d565b618232565b9050828152602081018484840111156182bc576182bb6181bc565b5b6182c784828561827e565b509392505050565b600082601f8301126182e4576182e36181b7565b5b81356182f484826020860161828d565b91505092915050565b61830681618056565b811461831157600080fd5b50565b600081359050618323816182fd565b92915050565b6000806000806080858703121561834357618342617e52565b5b600061835187828801617ea5565b945050602085013567ffffffffffffffff81111561837257618371617e57565b5b61837e878288016182cf565b935050604061838f87828801618314565b92505060606183a087828801618314565b91505092959194509250565b6000602082840312156183c2576183c1617e52565b5b600082013567ffffffffffffffff8111156183e0576183df617e57565b5b6183ec848285016182cf565b91505092915050565b6000806040838503121561840c5761840b617e52565b5b600061841a85828601617ea5565b925050602061842b85828601618314565b9150509250929050565b61843e81618056565b82525050565b60006020820190506184596000830184618435565b92915050565b6184688161802b565b82525050565b6000602082019050618483600083018461845f565b92915050565b600067ffffffffffffffff8211156184a4576184a36181d2565b5b602082029050602081019050919050565b600080fd5b60006184cd6184c884618489565b618232565b905080838252602082019050602084028301858111156184f0576184ef6184b5565b5b835b8181101561851957806185058882617f34565b8452602084019350506020810190506184f2565b5050509392505050565b600082601f830112618538576185376181b7565b5b81356185488482602086016184ba565b91505092915050565b60006020828403121561856757618566617e52565b5b600082013567ffffffffffffffff81111561858557618584617e57565b5b61859184828501618523565b91505092915050565b6185a381617e7c565b82525050565b60006020820190506185be600083018461859a565b92915050565b60006185cf82617fa2565b9050919050565b6185df816185c4565b82525050565b60006020820190506185fa60008301846185d6565b92915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561866657808201518184015260208101905061864b565b83811115618675576000848401525b50505050565b60006186868261862c565b6186908185618637565b93506186a0818560208601618648565b6186a9816181c1565b840191505092915050565b600060c0830160008301516186cc6000860182618038565b50602083015184820360208601526186e4828261867b565b91505060408301516186f96040860182618060565b50606083015161870c6060860182618060565b50608083015161871f608086018261801c565b5060a083015161873260a0860182618060565b508091505092915050565b600061874983836186b4565b905092915050565b6000602082019050919050565b600061876982618600565b618773818561860b565b9350836020820285016187858561861c565b8060005b858110156187c157848403895281516187a2858261873d565b94506187ad83618751565b925060208a01995050600181019050618789565b50829750879550505050505092915050565b600060208201905081810360008301526187ed818461875e565b905092915050565b6187fe81617ee7565b811461880957600080fd5b50565b60008135905061881b816187f5565b92915050565b60006020828403121561883757618836617e52565b5b60006188458482850161880c565b91505092915050565b600067ffffffffffffffff821115618869576188686181d2565b5b602082029050602081019050919050565b600061888d6188888461884e565b618232565b905080838252602082019050602084028301858111156188b0576188af6184b5565b5b835b818110156188d957806188c58882617ea5565b8452602084019350506020810190506188b2565b5050509392505050565b600082601f8301126188f8576188f76181b7565b5b813561890884826020860161887a565b91505092915050565b60008060006060848603121561892a57618929617e52565b5b600084013567ffffffffffffffff81111561894857618947617e57565b5b618954868287016188e3565b935050602084013567ffffffffffffffff81111561897557618974617e57565b5b618981868287016182cf565b925050604061899286828701617f34565b9150509250925092565b600080600080608085870312156189b6576189b5617e52565b5b600085013567ffffffffffffffff8111156189d4576189d3617e57565b5b6189e0878288016182cf565b94505060206189f187828801617f34565b9350506040618a0287828801617ea5565b9250506060618a1387828801617ea5565b91505092959194509250565b600060208284031215618a3557618a34617e52565b5b6000618a4384828501618314565b91505092915050565b600061010082019050618a62600083018b617e1e565b618a6f602083018a61845f565b618a7c604083018961859a565b618a896060830188617e1e565b618a966080830187617e1e565b618aa360a0830186617e1e565b618ab060c0830185618435565b618abd60e0830184618435565b9998505050505050505050565b618ad38161802b565b8114618ade57600080fd5b50565b600081359050618af081618aca565b92915050565b600060208284031215618b0c57618b0b617e52565b5b6000618b1a84828501618ae1565b91505092915050565b60008060408385031215618b3a57618b39617e52565b5b600083013567ffffffffffffffff811115618b5857618b57617e57565b5b618b64858286016182cf565b9250506020618b7585828601617f34565b9150509250929050565b600080600080600060a08688031215618b9b57618b9a617e52565b5b600086013567ffffffffffffffff811115618bb957618bb8617e57565b5b618bc5888289016182cf565b9550506020618bd688828901618314565b9450506040618be788828901618314565b9350506060618bf888828901617f34565b9250506080618c0988828901618314565b9150509295509295909350565b600067ffffffffffffffff821115618c3157618c306181d2565b5b602082029050602081019050919050565b6000618c55618c5084618c16565b618232565b90508083825260208201905060208402830185811115618c7857618c776184b5565b5b835b81811015618ca15780618c8d8882618314565b845260208401935050602081019050618c7a565b5050509392505050565b600082601f830112618cc057618cbf6181b7565b5b8135618cd0848260208601618c42565b91505092915050565b60008060408385031215618cf057618cef617e52565b5b6000618cfe85828601618ae1565b925050602083013567ffffffffffffffff811115618d1f57618d1e617e57565b5b618d2b85828601618cab565b9150509250929050565b600060208284031215618d4b57618d4a617e52565b5b600082013567ffffffffffffffff811115618d6957618d68617e57565b5b618d75848285016188e3565b91505092915050565b6000618d8982617fa2565b9050919050565b618d9981618d7e565b82525050565b6000602082019050618db46000830184618d90565b92915050565b60008060408385031215618dd157618dd0617e52565b5b6000618ddf85828601617ea5565b9250506020618df08582860161880c565b9150509250929050565b60008060008060008060c08789031215618e1757618e16617e52565b5b600087013567ffffffffffffffff811115618e3557618e34617e57565b5b618e4189828a016182cf565b965050602087013567ffffffffffffffff811115618e6257618e61617e57565b5b618e6e89828a016182cf565b9550506040618e7f89828a01618314565b9450506060618e9089828a01618314565b9350506080618ea189828a01617f34565b92505060a0618eb289828a01618314565b9150509295509295509295565b60008060408385031215618ed657618ed5617e52565b5b6000618ee485828601617ea5565b9250506020618ef585828601617ea5565b9150509250929050565b600082825260208201905092915050565b7f4e6f646573206d69677261746564000000000000000000000000000000000000600082015250565b6000618f46600e83618eff565b9150618f5182618f10565b602082019050919050565b60006020820190508181036000830152618f7581618f39565b9050919050565b7f526577617264206d696772617465640000000000000000000000000000000000600082015250565b6000618fb2600f83618eff565b9150618fbd82618f7c565b602082019050919050565b60006020820190508181036000830152618fe181618fa5565b9050919050565b600081519050618ff7816182fd565b92915050565b60006020828403121561901357619012617e52565b5b600061902184828501618fe8565b91505092915050565b7f6f6e6c794f776e65720000000000000000000000000000000000000000000000600082015250565b6000619060600983618eff565b915061906b8261902a565b602082019050919050565b6000602082019050818103600083015261908f81619053565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006190ff82618056565b915061910a83618056565b92508282101561911d5761911c6190c5565b5b828203905092915050565b600061913382617e0e565b915063ffffffff8203619149576191486190c5565b5b600182019050919050565b6000606082019050619169600083018661859a565b619176602083018561859a565b6191836040830184618435565b949350505050565b60008151905061919a816187f5565b92915050565b6000602082840312156191b6576191b5617e52565b5b60006191c48482850161918b565b91505092915050565b7f6c6f772076616c75650000000000000000000000000000000000000000000000600082015250565b6000619203600983618eff565b915061920e826191cd565b602082019050919050565b60006020820190508181036000830152619232816191f6565b9050919050565b7f496e737566660000000000000000000000000000000000000000000000000000600082015250565b600061926f600683618eff565b915061927a82619239565b602082019050919050565b6000602082019050818103600083015261929e81619262565b9050919050565b60006192b08261862c565b6192ba8185618eff565b93506192ca818560208601618648565b6192d3816181c1565b840191505092915050565b60006080820190506192f3600083018761859a565b619300602083018661859a565b818103604083015261931281856192a5565b90506193216060830184618435565b95945050505050565b600061933582617e0e565b915060008203619348576193476190c5565b5b600182039050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061939a57607f821691505b6020821081036193ad576193ac619353565b5b50919050565b600081905092915050565b60008190508160005260206000209050919050565b600081546193e081619382565b6193ea81866193b3565b94506001821660008114619405576001811461941657619449565b60ff19831686528186019350619449565b61941f856193be565b60005b8381101561944157815481890152600182019150602081019050619422565b838801955050505b50505092915050565b600061945e82846193d3565b915081905092915050565b600061947482617e0e565b915061947f83617e0e565b925082821015619492576194916190c5565b5b828203905092915050565b60006194a882618056565b91506194b383618056565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156194e8576194e76190c5565b5b828201905092915050565b60006194fe8261802b565b915060ff8203619511576195106190c5565b5b600182019050919050565b600061952782618056565b915061953283618056565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561956b5761956a6190c5565b5b828202905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b60006195b082618056565b91506195bb83618056565b9250826195cb576195ca619576565b5b828204905092915050565b7f456d707479000000000000000000000000000000000000000000000000000000600082015250565b600061960c600583618eff565b9150619617826195d6565b602082019050919050565b6000602082019050818103600083015261963b816195ff565b9050919050565b600061964d82618056565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361967f5761967e6190c5565b5b600182019050919050565b7f426c61636b000000000000000000000000000000000000000000000000000000600082015250565b60006196c0600583618eff565b91506196cb8261968a565b602082019050919050565b600060208201905081810360008301526196ef816196b3565b9050919050565b60006197018261862c565b61970b81856193b3565b935061971b818560208601618648565b80840191505092915050565b600061973382846196f6565b915081905092915050565b600061974982617e0e565b915061975483617e0e565b92508263ffffffff0382111561976d5761976c6190c5565b5b828201905092915050565b7f6d61780000000000000000000000000000000000000000000000000000000000600082015250565b60006197ae600383618eff565b91506197b982619778565b602082019050919050565b600060208201905081810360008301526197dd816197a1565b9050919050565b60006197ef8261802b565b91506197fa8361802b565b92508282101561980d5761980c6190c5565b5b828203905092915050565b7f4e20656e6f756768000000000000000000000000000000000000000000000000600082015250565b600061984e600883618eff565b915061985982619818565b602082019050919050565b6000602082019050818103600083015261987d81619841565b9050919050565b6000606082019050619899600083018661859a565b6198a6602083018561859a565b6198b36040830184617e1e565b949350505050565b6000815190506198ca81617e8e565b92915050565b6000602082840312156198e6576198e5617e52565b5b60006198f4848285016198bb565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b60006199358383618047565b60208301905092915050565b6000602082019050919050565b6000619959826198fd565b6199638185619908565b935061996e83619919565b8060005b8381101561999f5781516199868882619929565b975061999183619941565b925050600181019050619972565b5085935050505092915050565b60006040820190506199c16000830185618435565b81810360208301526199d3818461994e565b90509392505050565b60006199ef6199ea84618c16565b618232565b90508083825260208201905060208402830185811115619a1257619a116184b5565b5b835b81811015619a3b5780619a278882618fe8565b845260208401935050602081019050619a14565b5050509392505050565b600082601f830112619a5a57619a596181b7565b5b8151619a6a8482602086016199dc565b91505092915050565b600060208284031215619a8957619a88617e52565b5b600082015167ffffffffffffffff811115619aa757619aa6617e57565b5b619ab384828501619a45565b91505092915050565b7f6d69677261746564000000000000000000000000000000000000000000000000600082015250565b6000619af2600883618eff565b9150619afd82619abc565b602082019050919050565b60006020820190508181036000830152619b2181619ae5565b9050919050565b600060c082019050619b3d600083018961859a565b8181036020830152619b4f81886192a5565b9050619b5e6040830187617e1e565b619b6b6060830186617e1e565b619b786080830185617e1e565b619b8560a0830184617e1e565b979650505050505050565b7f7072696365000000000000000000000000000000000000000000000000000000600082015250565b6000619bc6600583618eff565b9150619bd182619b90565b602082019050919050565b60006020820190508181036000830152619bf581619bb9565b9050919050565b7f7265776172647300000000000000000000000000000000000000000000000000600082015250565b6000619c32600783618eff565b9150619c3d82619bfc565b602082019050919050565b60006020820190508181036000830152619c6181619c25565b9050919050565b7f636c61696d000000000000000000000000000000000000000000000000000000600082015250565b6000619c9e600583618eff565b9150619ca982619c68565b602082019050919050565b60006020820190508181036000830152619ccd81619c91565b9050919050565b7f496e76616c696400000000000000000000000000000000000000000000000000600082015250565b6000619d0a600783618eff565b9150619d1582619cd4565b602082019050919050565b60006020820190508181036000830152619d3981619cfd565b9050919050565b6000619d4b82617e0e565b9150619d5683617e0e565b92508163ffffffff0483118215151615619d7357619d726190c5565b5b828202905092915050565b7f33206d0000000000000000000000000000000000000000000000000000000000600082015250565b6000619db4600383618eff565b9150619dbf82619d7e565b602082019050919050565b60006020820190508181036000830152619de381619da7565b9050919050565b7f4665650000000000000000000000000000000000000000000000000000000000600082015250565b6000619e20600383618eff565b9150619e2b82619dea565b602082019050919050565b60006020820190508181036000830152619e4f81619e13565b9050919050565b6000819050919050565b6000619e7b619e76619e7184619e56565b617f76565b61802b565b9050919050565b619e8b81619e60565b82525050565b6000602082019050619ea66000830184619e82565b92915050565b7f4f6c640000000000000000000000000000000000000000000000000000000000600082015250565b6000619ee2600383618eff565b9150619eed82619eac565b602082019050919050565b60006020820190508181036000830152619f1181619ed5565b9050919050565b7f6e616d6520696e636f7272656374000000000000000000000000000000000000600082015250565b6000619f4e600e83618eff565b9150619f5982619f18565b602082019050919050565b60006020820190508181036000830152619f7d81619f41565b9050919050565b7f7265776172647350657254696d65000000000000000000000000000000000000600082015250565b6000619fba600e83618eff565b9150619fc582619f84565b602082019050919050565b60006020820190508181036000830152619fe981619fad565b9050919050565b7f72656d6f76656400000000000000000000000000000000000000000000000000600082015250565b600061a026600783618eff565b915061a03182619ff0565b602082019050919050565b6000602082019050818103600083015261a0558161a019565b9050919050565b600061a0678261802b565b91506000820361a07a5761a0796190c5565b5b600182039050919050565b7f7a65726f00000000000000000000000000000000000000000000000000000000600082015250565b600061a0bb600483618eff565b915061a0c68261a085565b602082019050919050565b6000602082019050818103600083015261a0ea8161a0ae565b9050919050565b600060408201905061a106600083018561859a565b61a1136020830184618435565b9392505050565b7f426c61636b6c6973746564000000000000000000000000000000000000000000600082015250565b600061a150600b83618eff565b915061a15b8261a11a565b602082019050919050565b6000602082019050818103600083015261a17f8161a143565b9050919050565b7f4d61780000000000000000000000000000000000000000000000000000000000600082015250565b600061a1bc600383618eff565b915061a1c78261a186565b602082019050919050565b6000602082019050818103600083015261a1eb8161a1af565b9050919050565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b600061a24e602e83618eff565b915061a2598261a1f2565b604082019050919050565b6000602082019050818103600083015261a27d8161a241565b905091905056fea2646970667358221220f62dc06fe229084d1ee4f6e9d58b002b227d62411ddff0b9e1ebd57e0767d68564736f6c637827302e382e31342d646576656c6f702e323032322e352e392b636f6d6d69742e34343135376161360058",
}

// NodeManagerV83ABI is the input ABI used to generate the binding from.
// Deprecated: Use NodeManagerV83MetaData.ABI instead.
var NodeManagerV83ABI = NodeManagerV83MetaData.ABI

// NodeManagerV83Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use NodeManagerV83MetaData.Bin instead.
var NodeManagerV83Bin = NodeManagerV83MetaData.Bin

// DeployNodeManagerV83 deploys a new Ethereum contract, binding an instance of NodeManagerV83 to it.
func DeployNodeManagerV83(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *NodeManagerV83, error) {
	parsed, err := NodeManagerV83MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NodeManagerV83Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NodeManagerV83{NodeManagerV83Caller: NodeManagerV83Caller{contract: contract}, NodeManagerV83Transactor: NodeManagerV83Transactor{contract: contract}, NodeManagerV83Filterer: NodeManagerV83Filterer{contract: contract}}, nil
}

// NodeManagerV83 is an auto generated Go binding around an Ethereum contract.
type NodeManagerV83 struct {
	NodeManagerV83Caller     // Read-only binding to the contract
	NodeManagerV83Transactor // Write-only binding to the contract
	NodeManagerV83Filterer   // Log filterer for contract events
}

// NodeManagerV83Caller is an auto generated read-only Go binding around an Ethereum contract.
type NodeManagerV83Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeManagerV83Transactor is an auto generated write-only Go binding around an Ethereum contract.
type NodeManagerV83Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeManagerV83Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NodeManagerV83Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeManagerV83Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NodeManagerV83Session struct {
	Contract     *NodeManagerV83   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NodeManagerV83CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NodeManagerV83CallerSession struct {
	Contract *NodeManagerV83Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// NodeManagerV83TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NodeManagerV83TransactorSession struct {
	Contract     *NodeManagerV83Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// NodeManagerV83Raw is an auto generated low-level Go binding around an Ethereum contract.
type NodeManagerV83Raw struct {
	Contract *NodeManagerV83 // Generic contract binding to access the raw methods on
}

// NodeManagerV83CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NodeManagerV83CallerRaw struct {
	Contract *NodeManagerV83Caller // Generic read-only contract binding to access the raw methods on
}

// NodeManagerV83TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NodeManagerV83TransactorRaw struct {
	Contract *NodeManagerV83Transactor // Generic write-only contract binding to access the raw methods on
}

// NewNodeManagerV83 creates a new instance of NodeManagerV83, bound to a specific deployed contract.
func NewNodeManagerV83(address common.Address, backend bind.ContractBackend) (*NodeManagerV83, error) {
	contract, err := bindNodeManagerV83(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NodeManagerV83{NodeManagerV83Caller: NodeManagerV83Caller{contract: contract}, NodeManagerV83Transactor: NodeManagerV83Transactor{contract: contract}, NodeManagerV83Filterer: NodeManagerV83Filterer{contract: contract}}, nil
}

// NewNodeManagerV83Caller creates a new read-only instance of NodeManagerV83, bound to a specific deployed contract.
func NewNodeManagerV83Caller(address common.Address, caller bind.ContractCaller) (*NodeManagerV83Caller, error) {
	contract, err := bindNodeManagerV83(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NodeManagerV83Caller{contract: contract}, nil
}

// NewNodeManagerV83Transactor creates a new write-only instance of NodeManagerV83, bound to a specific deployed contract.
func NewNodeManagerV83Transactor(address common.Address, transactor bind.ContractTransactor) (*NodeManagerV83Transactor, error) {
	contract, err := bindNodeManagerV83(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NodeManagerV83Transactor{contract: contract}, nil
}

// NewNodeManagerV83Filterer creates a new log filterer instance of NodeManagerV83, bound to a specific deployed contract.
func NewNodeManagerV83Filterer(address common.Address, filterer bind.ContractFilterer) (*NodeManagerV83Filterer, error) {
	contract, err := bindNodeManagerV83(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NodeManagerV83Filterer{contract: contract}, nil
}

// bindNodeManagerV83 binds a generic wrapper to an already deployed contract.
func bindNodeManagerV83(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NodeManagerV83ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeManagerV83 *NodeManagerV83Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NodeManagerV83.Contract.NodeManagerV83Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeManagerV83 *NodeManagerV83Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeManagerV83.Contract.NodeManagerV83Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeManagerV83 *NodeManagerV83Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeManagerV83.Contract.NodeManagerV83Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeManagerV83 *NodeManagerV83CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NodeManagerV83.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeManagerV83 *NodeManagerV83TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeManagerV83.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeManagerV83 *NodeManagerV83TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeManagerV83.Contract.contract.Transact(opts, method, params...)
}

// IsBlacklisted is a free data retrieval call binding the contract method 0x1cdd3be3.
//
// Solidity: function _isBlacklisted(address ) view returns(bool)
func (_NodeManagerV83 *NodeManagerV83Caller) IsBlacklisted(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _NodeManagerV83.contract.Call(opts, &out, "_isBlacklisted", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBlacklisted is a free data retrieval call binding the contract method 0x1cdd3be3.
//
// Solidity: function _isBlacklisted(address ) view returns(bool)
func (_NodeManagerV83 *NodeManagerV83Session) IsBlacklisted(arg0 common.Address) (bool, error) {
	return _NodeManagerV83.Contract.IsBlacklisted(&_NodeManagerV83.CallOpts, arg0)
}

// IsBlacklisted is a free data retrieval call binding the contract method 0x1cdd3be3.
//
// Solidity: function _isBlacklisted(address ) view returns(bool)
func (_NodeManagerV83 *NodeManagerV83CallerSession) IsBlacklisted(arg0 common.Address) (bool, error) {
	return _NodeManagerV83.Contract.IsBlacklisted(&_NodeManagerV83.CallOpts, arg0)
}

// BoostNFT is a free data retrieval call binding the contract method 0x4506598d.
//
// Solidity: function boostNFT() view returns(address)
func (_NodeManagerV83 *NodeManagerV83Caller) BoostNFT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NodeManagerV83.contract.Call(opts, &out, "boostNFT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BoostNFT is a free data retrieval call binding the contract method 0x4506598d.
//
// Solidity: function boostNFT() view returns(address)
func (_NodeManagerV83 *NodeManagerV83Session) BoostNFT() (common.Address, error) {
	return _NodeManagerV83.Contract.BoostNFT(&_NodeManagerV83.CallOpts)
}

// BoostNFT is a free data retrieval call binding the contract method 0x4506598d.
//
// Solidity: function boostNFT() view returns(address)
func (_NodeManagerV83 *NodeManagerV83CallerSession) BoostNFT() (common.Address, error) {
	return _NodeManagerV83.Contract.BoostNFT(&_NodeManagerV83.CallOpts)
}

// CheckNodes is a free data retrieval call binding the contract method 0x450970b5.
//
// Solidity: function checkNodes(address ownerAddress) view returns((uint32,uint8,address,uint32,uint32,uint32,uint256,uint256)[])
func (_NodeManagerV83 *NodeManagerV83Caller) CheckNodes(opts *bind.CallOpts, ownerAddress common.Address) ([]Node, error) {
	var out []interface{}
	err := _NodeManagerV83.contract.Call(opts, &out, "checkNodes", ownerAddress)

	if err != nil {
		return *new([]Node), err
	}

	out0 := *abi.ConvertType(out[0], new([]Node)).(*[]Node)

	return out0, err

}

// CheckNodes is a free data retrieval call binding the contract method 0x450970b5.
//
// Solidity: function checkNodes(address ownerAddress) view returns((uint32,uint8,address,uint32,uint32,uint32,uint256,uint256)[])
func (_NodeManagerV83 *NodeManagerV83Session) CheckNodes(ownerAddress common.Address) ([]Node, error) {
	return _NodeManagerV83.Contract.CheckNodes(&_NodeManagerV83.CallOpts, ownerAddress)
}

// CheckNodes is a free data retrieval call binding the contract method 0x450970b5.
//
// Solidity: function checkNodes(address ownerAddress) view returns((uint32,uint8,address,uint32,uint32,uint32,uint256,uint256)[])
func (_NodeManagerV83 *NodeManagerV83CallerSession) CheckNodes(ownerAddress common.Address) ([]Node, error) {
	return _NodeManagerV83.Contract.CheckNodes(&_NodeManagerV83.CallOpts, ownerAddress)
}

// ClaimFee is a free data retrieval call binding the contract method 0x99d32fc4.
//
// Solidity: function claimFee() view returns(uint32)
func (_NodeManagerV83 *NodeManagerV83Caller) ClaimFee(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _NodeManagerV83.contract.Call(opts, &out, "claimFee")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ClaimFee is a free data retrieval call binding the contract method 0x99d32fc4.
//
// Solidity: function claimFee() view returns(uint32)
func (_NodeManagerV83 *NodeManagerV83Session) ClaimFee() (uint32, error) {
	return _NodeManagerV83.Contract.ClaimFee(&_NodeManagerV83.CallOpts)
}

// ClaimFee is a free data retrieval call binding the contract method 0x99d32fc4.
//
// Solidity: function claimFee() view returns(uint32)
func (_NodeManagerV83 *NodeManagerV83CallerSession) ClaimFee() (uint32, error) {
	return _NodeManagerV83.Contract.ClaimFee(&_NodeManagerV83.CallOpts)
}

// Claimable is a free data retrieval call binding the contract method 0x402914f5.
//
// Solidity: function claimable(address _account) view returns(uint256)
func (_NodeManagerV83 *NodeManagerV83Caller) Claimable(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NodeManagerV83.contract.Call(opts, &out, "claimable", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Claimable is a free data retrieval call binding the contract method 0x402914f5.
//
// Solidity: function claimable(address _account) view returns(uint256)
func (_NodeManagerV83 *NodeManagerV83Session) Claimable(_account common.Address) (*big.Int, error) {
	return _NodeManagerV83.Contract.Claimable(&_NodeManagerV83.CallOpts, _account)
}

// Claimable is a free data retrieval call binding the contract method 0x402914f5.
//
// Solidity: function claimable(address _account) view returns(uint256)
func (_NodeManagerV83 *NodeManagerV83CallerSession) Claimable(_account common.Address) (*big.Int, error) {
	return _NodeManagerV83.Contract.Claimable(&_NodeManagerV83.CallOpts, _account)
}

// CountOfTier is a free data retrieval call binding the contract method 0x22b45822.
//
// Solidity: function countOfTier(string ) view returns(uint32)
func (_NodeManagerV83 *NodeManagerV83Caller) CountOfTier(opts *bind.CallOpts, arg0 string) (uint32, error) {
	var out []interface{}
	err := _NodeManagerV83.contract.Call(opts, &out, "countOfTier", arg0)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// CountOfTier is a free data retrieval call binding the contract method 0x22b45822.
//
// Solidity: function countOfTier(string ) view returns(uint32)
func (_NodeManagerV83 *NodeManagerV83Session) CountOfTier(arg0 string) (uint32, error) {
	return _NodeManagerV83.Contract.CountOfTier(&_NodeManagerV83.CallOpts, arg0)
}

// CountOfTier is a free data retrieval call binding the contract method 0x22b45822.
//
// Solidity: function countOfTier(string ) view returns(uint32)
func (_NodeManagerV83 *NodeManagerV83CallerSession) CountOfTier(arg0 string) (uint32, error) {
	return _NodeManagerV83.Contract.CountOfTier(&_NodeManagerV83.CallOpts, arg0)
}

// CountOfUser is a free data retrieval call binding the contract method 0xd1b43593.
//
// Solidity: function countOfUser(address ) view returns(uint32)
func (_NodeManagerV83 *NodeManagerV83Caller) CountOfUser(opts *bind.CallOpts, arg0 common.Address) (uint32, error) {
	var out []interface{}
	err := _NodeManagerV83.contract.Call(opts, &out, "countOfUser", arg0)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// CountOfUser is a free data retrieval call binding the contract method 0xd1b43593.
//
// Solidity: function countOfUser(address ) view returns(uint32)
func (_NodeManagerV83 *NodeManagerV83Session) CountOfUser(arg0 common.Address) (uint32, error) {
	return _NodeManagerV83.Contract.CountOfUser(&_NodeManagerV83.CallOpts, arg0)
}

// CountOfUser is a free data retrieval call binding the contract method 0xd1b43593.
//
// Solidity: function countOfUser(address ) view returns(uint32)
func (_NodeManagerV83 *NodeManagerV83CallerSession) CountOfUser(arg0 common.Address) (uint32, error) {
	return _NodeManagerV83.Contract.CountOfUser(&_NodeManagerV83.CallOpts, arg0)
}

// CountTotal is a free data retrieval call binding the contract method 0xd33953bb.
//
// Solidity: function countTotal() view returns(uint32)
func (_NodeManagerV83 *NodeManagerV83Caller) CountTotal(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _NodeManagerV83.contract.Call(opts, &out, "countTotal")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// CountTotal is a free data retrieval call binding the contract method 0xd33953bb.
//
// Solidity: function countTotal() view returns(uint32)
func (_NodeManagerV83 *NodeManagerV83Session) CountTotal() (uint32, error) {
	return _NodeManagerV83.Contract.CountTotal(&_NodeManagerV83.CallOpts)
}

// CountTotal is a free data retrieval call binding the contract method 0xd33953bb.
//
// Solidity: function countTotal() view returns(uint32)
func (_NodeManagerV83 *NodeManagerV83CallerSession) CountTotal() (uint32, error) {
	return _NodeManagerV83.Contract.CountTotal(&_NodeManagerV83.CallOpts)
}

// DiscountPer10 is a free data retrieval call binding the contract method 0x1c87cdb0.
//
// Solidity: function discountPer10() view returns(uint32)
func (_NodeManagerV83 *NodeManagerV83Caller) DiscountPer10(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _NodeManagerV83.contract.Call(opts, &out, "discountPer10")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// DiscountPer10 is a free data retrieval call binding the contract method 0x1c87cdb0.
//
// Solidity: function discountPer10() view returns(uint32)
func (_NodeManagerV83 *NodeManagerV83Session) DiscountPer10() (uint32, error) {
	return _NodeManagerV83.Contract.DiscountPer10(&_NodeManagerV83.CallOpts)
}

// DiscountPer10 is a free data retrieval call binding the contract method 0x1c87cdb0.
//
// Solidity: function discountPer10() view returns(uint32)
func (_NodeManagerV83 *NodeManagerV83CallerSession) DiscountPer10() (uint32, error) {
	return _NodeManagerV83.Contract.DiscountPer10(&_NodeManagerV83.CallOpts)
}

// GetAddress is a free data retrieval call binding the contract method 0x38cc4831.
//
// Solidity: function getAddress() view returns(address)
func (_NodeManagerV83 *NodeManagerV83Caller) GetAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NodeManagerV83.contract.Call(opts, &out, "getAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddress is a free data retrieval call binding the contract method 0x38cc4831.
//
// Solidity: function getAddress() view returns(address)
func (_NodeManagerV83 *NodeManagerV83Session) GetAddress() (common.Address, error) {
	return _NodeManagerV83.Contract.GetAddress(&_NodeManagerV83.CallOpts)
}

// GetAddress is a free data retrieval call binding the contract method 0x38cc4831.
//
// Solidity: function getAddress() view returns(address)
func (_NodeManagerV83 *NodeManagerV83CallerSession) GetAddress() (common.Address, error) {
	return _NodeManagerV83.Contract.GetAddress(&_NodeManagerV83.CallOpts)
}

// GetAmountOut is a free data retrieval call binding the contract method 0x5c195217.
//
// Solidity: function getAmountOut(uint256 _amount) view returns(uint256)
func (_NodeManagerV83 *NodeManagerV83Caller) GetAmountOut(opts *bind.CallOpts, _amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _NodeManagerV83.contract.Call(opts, &out, "getAmountOut", _amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountOut is a free data retrieval call binding the contract method 0x5c195217.
//
// Solidity: function getAmountOut(uint256 _amount) view returns(uint256)
func (_NodeManagerV83 *NodeManagerV83Session) GetAmountOut(_amount *big.Int) (*big.Int, error) {
	return _NodeManagerV83.Contract.GetAmountOut(&_NodeManagerV83.CallOpts, _amount)
}

// GetAmountOut is a free data retrieval call binding the contract method 0x5c195217.
//
// Solidity: function getAmountOut(uint256 _amount) view returns(uint256)
func (_NodeManagerV83 *NodeManagerV83CallerSession) GetAmountOut(_amount *big.Int) (*big.Int, error) {
	return _NodeManagerV83.Contract.GetAmountOut(&_NodeManagerV83.CallOpts, _amount)
}

// LastTimestampClaim is a free data retrieval call binding the contract method 0x721ffa94.
//
// Solidity: function lastTimestampClaim(address ) view returns(uint32)
func (_NodeManagerV83 *NodeManagerV83Caller) LastTimestampClaim(opts *bind.CallOpts, arg0 common.Address) (uint32, error) {
	var out []interface{}
	err := _NodeManagerV83.contract.Call(opts, &out, "lastTimestampClaim", arg0)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LastTimestampClaim is a free data retrieval call binding the contract method 0x721ffa94.
//
// Solidity: function lastTimestampClaim(address ) view returns(uint32)
func (_NodeManagerV83 *NodeManagerV83Session) LastTimestampClaim(arg0 common.Address) (uint32, error) {
	return _NodeManagerV83.Contract.LastTimestampClaim(&_NodeManagerV83.CallOpts, arg0)
}

// LastTimestampClaim is a free data retrieval call binding the contract method 0x721ffa94.
//
// Solidity: function lastTimestampClaim(address ) view returns(uint32)
func (_NodeManagerV83 *NodeManagerV83CallerSession) LastTimestampClaim(arg0 common.Address) (uint32, error) {
	return _NodeManagerV83.Contract.LastTimestampClaim(&_NodeManagerV83.CallOpts, arg0)
}

// Managers is a free data retrieval call binding the contract method 0xfdff9b4d.
//
// Solidity: function managers(address ) view returns(bool)
func (_NodeManagerV83 *NodeManagerV83Caller) Managers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _NodeManagerV83.contract.Call(opts, &out, "managers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Managers is a free data retrieval call binding the contract method 0xfdff9b4d.
//
// Solidity: function managers(address ) view returns(bool)
func (_NodeManagerV83 *NodeManagerV83Session) Managers(arg0 common.Address) (bool, error) {
	return _NodeManagerV83.Contract.Managers(&_NodeManagerV83.CallOpts, arg0)
}

// Managers is a free data retrieval call binding the contract method 0xfdff9b4d.
//
// Solidity: function managers(address ) view returns(bool)
func (_NodeManagerV83 *NodeManagerV83CallerSession) Managers(arg0 common.Address) (bool, error) {
	return _NodeManagerV83.Contract.Managers(&_NodeManagerV83.CallOpts, arg0)
}

// MantPercent is a free data retrieval call binding the contract method 0x3747bd45.
//
// Solidity: function mantPercent() view returns(uint32)
func (_NodeManagerV83 *NodeManagerV83Caller) MantPercent(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _NodeManagerV83.contract.Call(opts, &out, "mantPercent")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MantPercent is a free data retrieval call binding the contract method 0x3747bd45.
//
// Solidity: function mantPercent() view returns(uint32)
func (_NodeManagerV83 *NodeManagerV83Session) MantPercent() (uint32, error) {
	return _NodeManagerV83.Contract.MantPercent(&_NodeManagerV83.CallOpts)
}

// MantPercent is a free data retrieval call binding the contract method 0x3747bd45.
//
// Solidity: function mantPercent() view returns(uint32)
func (_NodeManagerV83 *NodeManagerV83CallerSession) MantPercent() (uint32, error) {
	return _NodeManagerV83.Contract.MantPercent(&_NodeManagerV83.CallOpts)
}

// MaxCountOfUser is a free data retrieval call binding the contract method 0x0a3eccf2.
//
// Solidity: function maxCountOfUser() view returns(uint32)
func (_NodeManagerV83 *NodeManagerV83Caller) MaxCountOfUser(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _NodeManagerV83.contract.Call(opts, &out, "maxCountOfUser")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MaxCountOfUser is a free data retrieval call binding the contract method 0x0a3eccf2.
//
// Solidity: function maxCountOfUser() view returns(uint32)
func (_NodeManagerV83 *NodeManagerV83Session) MaxCountOfUser() (uint32, error) {
	return _NodeManagerV83.Contract.MaxCountOfUser(&_NodeManagerV83.CallOpts)
}

// MaxCountOfUser is a free data retrieval call binding the contract method 0x0a3eccf2.
//
// Solidity: function maxCountOfUser() view returns(uint32)
func (_NodeManagerV83 *NodeManagerV83CallerSession) MaxCountOfUser() (uint32, error) {
	return _NodeManagerV83.Contract.MaxCountOfUser(&_NodeManagerV83.CallOpts)
}

// MaxMonthValue is a free data retrieval call binding the contract method 0x2d233503.
//
// Solidity: function maxMonthValue() view returns(uint8)
func (_NodeManagerV83 *NodeManagerV83Caller) MaxMonthValue(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _NodeManagerV83.contract.Call(opts, &out, "maxMonthValue")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MaxMonthValue is a free data retrieval call binding the contract method 0x2d233503.
//
// Solidity: function maxMonthValue() view returns(uint8)
func (_NodeManagerV83 *NodeManagerV83Session) MaxMonthValue() (uint8, error) {
	return _NodeManagerV83.Contract.MaxMonthValue(&_NodeManagerV83.CallOpts)
}

// MaxMonthValue is a free data retrieval call binding the contract method 0x2d233503.
//
// Solidity: function maxMonthValue() view returns(uint8)
func (_NodeManagerV83 *NodeManagerV83CallerSession) MaxMonthValue() (uint8, error) {
	return _NodeManagerV83.Contract.MaxMonthValue(&_NodeManagerV83.CallOpts)
}

// NftAddress is a free data retrieval call binding the contract method 0x5bf8633a.
//
// Solidity: function nftAddress() view returns(address)
func (_NodeManagerV83 *NodeManagerV83Caller) NftAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NodeManagerV83.contract.Call(opts, &out, "nftAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NftAddress is a free data retrieval call binding the contract method 0x5bf8633a.
//
// Solidity: function nftAddress() view returns(address)
func (_NodeManagerV83 *NodeManagerV83Session) NftAddress() (common.Address, error) {
	return _NodeManagerV83.Contract.NftAddress(&_NodeManagerV83.CallOpts)
}

// NftAddress is a free data retrieval call binding the contract method 0x5bf8633a.
//
// Solidity: function nftAddress() view returns(address)
func (_NodeManagerV83 *NodeManagerV83CallerSession) NftAddress() (common.Address, error) {
	return _NodeManagerV83.Contract.NftAddress(&_NodeManagerV83.CallOpts)
}

// Nodes is a free data retrieval call binding the contract method 0x189a5a17.
//
// Solidity: function nodes(address account) view returns((uint32,uint8,address,uint32,uint32,uint32,uint256,uint256)[])
func (_NodeManagerV83 *NodeManagerV83Caller) Nodes(opts *bind.CallOpts, account common.Address) ([]Node, error) {
	var out []interface{}
	err := _NodeManagerV83.contract.Call(opts, &out, "nodes", account)

	if err != nil {
		return *new([]Node), err
	}

	out0 := *abi.ConvertType(out[0], new([]Node)).(*[]Node)

	return out0, err

}

// Nodes is a free data retrieval call binding the contract method 0x189a5a17.
//
// Solidity: function nodes(address account) view returns((uint32,uint8,address,uint32,uint32,uint32,uint256,uint256)[])
func (_NodeManagerV83 *NodeManagerV83Session) Nodes(account common.Address) ([]Node, error) {
	return _NodeManagerV83.Contract.Nodes(&_NodeManagerV83.CallOpts, account)
}

// Nodes is a free data retrieval call binding the contract method 0x189a5a17.
//
// Solidity: function nodes(address account) view returns((uint32,uint8,address,uint32,uint32,uint32,uint256,uint256)[])
func (_NodeManagerV83 *NodeManagerV83CallerSession) Nodes(account common.Address) ([]Node, error) {
	return _NodeManagerV83.Contract.Nodes(&_NodeManagerV83.CallOpts, account)
}

// NodesOfUser is a free data retrieval call binding the contract method 0x2655c1b9.
//
// Solidity: function nodesOfUser(address , uint256 ) view returns(uint256)
func (_NodeManagerV83 *NodeManagerV83Caller) NodesOfUser(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _NodeManagerV83.contract.Call(opts, &out, "nodesOfUser", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NodesOfUser is a free data retrieval call binding the contract method 0x2655c1b9.
//
// Solidity: function nodesOfUser(address , uint256 ) view returns(uint256)
func (_NodeManagerV83 *NodeManagerV83Session) NodesOfUser(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _NodeManagerV83.Contract.NodesOfUser(&_NodeManagerV83.CallOpts, arg0, arg1)
}

// NodesOfUser is a free data retrieval call binding the contract method 0x2655c1b9.
//
// Solidity: function nodesOfUser(address , uint256 ) view returns(uint256)
func (_NodeManagerV83 *NodeManagerV83CallerSession) NodesOfUser(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _NodeManagerV83.Contract.NodesOfUser(&_NodeManagerV83.CallOpts, arg0, arg1)
}

// NodesTotal is a free data retrieval call binding the contract method 0x74bf71c5.
//
// Solidity: function nodesTotal(uint256 ) view returns(uint32 id, uint8 tierIndex, address owner, uint32 createdTime, uint32 claimedTime, uint32 limitedTime, uint256 multiplier, uint256 leftover)
func (_NodeManagerV83 *NodeManagerV83Caller) NodesTotal(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id          uint32
	TierIndex   uint8
	Owner       common.Address
	CreatedTime uint32
	ClaimedTime uint32
	LimitedTime uint32
	Multiplier  *big.Int
	Leftover    *big.Int
}, error) {
	var out []interface{}
	err := _NodeManagerV83.contract.Call(opts, &out, "nodesTotal", arg0)

	outstruct := new(struct {
		Id          uint32
		TierIndex   uint8
		Owner       common.Address
		CreatedTime uint32
		ClaimedTime uint32
		LimitedTime uint32
		Multiplier  *big.Int
		Leftover    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.TierIndex = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.Owner = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.CreatedTime = *abi.ConvertType(out[3], new(uint32)).(*uint32)
	outstruct.ClaimedTime = *abi.ConvertType(out[4], new(uint32)).(*uint32)
	outstruct.LimitedTime = *abi.ConvertType(out[5], new(uint32)).(*uint32)
	outstruct.Multiplier = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Leftover = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// NodesTotal is a free data retrieval call binding the contract method 0x74bf71c5.
//
// Solidity: function nodesTotal(uint256 ) view returns(uint32 id, uint8 tierIndex, address owner, uint32 createdTime, uint32 claimedTime, uint32 limitedTime, uint256 multiplier, uint256 leftover)
func (_NodeManagerV83 *NodeManagerV83Session) NodesTotal(arg0 *big.Int) (struct {
	Id          uint32
	TierIndex   uint8
	Owner       common.Address
	CreatedTime uint32
	ClaimedTime uint32
	LimitedTime uint32
	Multiplier  *big.Int
	Leftover    *big.Int
}, error) {
	return _NodeManagerV83.Contract.NodesTotal(&_NodeManagerV83.CallOpts, arg0)
}

// NodesTotal is a free data retrieval call binding the contract method 0x74bf71c5.
//
// Solidity: function nodesTotal(uint256 ) view returns(uint32 id, uint8 tierIndex, address owner, uint32 createdTime, uint32 claimedTime, uint32 limitedTime, uint256 multiplier, uint256 leftover)
func (_NodeManagerV83 *NodeManagerV83CallerSession) NodesTotal(arg0 *big.Int) (struct {
	Id          uint32
	TierIndex   uint8
	Owner       common.Address
	CreatedTime uint32
	ClaimedTime uint32
	LimitedTime uint32
	Multiplier  *big.Int
	Leftover    *big.Int
}, error) {
	return _NodeManagerV83.Contract.NodesTotal(&_NodeManagerV83.CallOpts, arg0)
}

// OldNodeManager is a free data retrieval call binding the contract method 0xaea5c394.
//
// Solidity: function oldNodeManager() view returns(address)
func (_NodeManagerV83 *NodeManagerV83Caller) OldNodeManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NodeManagerV83.contract.Call(opts, &out, "oldNodeManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OldNodeManager is a free data retrieval call binding the contract method 0xaea5c394.
//
// Solidity: function oldNodeManager() view returns(address)
func (_NodeManagerV83 *NodeManagerV83Session) OldNodeManager() (common.Address, error) {
	return _NodeManagerV83.Contract.OldNodeManager(&_NodeManagerV83.CallOpts)
}

// OldNodeManager is a free data retrieval call binding the contract method 0xaea5c394.
//
// Solidity: function oldNodeManager() view returns(address)
func (_NodeManagerV83 *NodeManagerV83CallerSession) OldNodeManager() (common.Address, error) {
	return _NodeManagerV83.Contract.OldNodeManager(&_NodeManagerV83.CallOpts)
}

// OldRewardsOfUser is a free data retrieval call binding the contract method 0x4c3f427d.
//
// Solidity: function oldRewardsOfUser(address ) view returns(uint256)
func (_NodeManagerV83 *NodeManagerV83Caller) OldRewardsOfUser(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NodeManagerV83.contract.Call(opts, &out, "oldRewardsOfUser", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OldRewardsOfUser is a free data retrieval call binding the contract method 0x4c3f427d.
//
// Solidity: function oldRewardsOfUser(address ) view returns(uint256)
func (_NodeManagerV83 *NodeManagerV83Session) OldRewardsOfUser(arg0 common.Address) (*big.Int, error) {
	return _NodeManagerV83.Contract.OldRewardsOfUser(&_NodeManagerV83.CallOpts, arg0)
}

// OldRewardsOfUser is a free data retrieval call binding the contract method 0x4c3f427d.
//
// Solidity: function oldRewardsOfUser(address ) view returns(uint256)
func (_NodeManagerV83 *NodeManagerV83CallerSession) OldRewardsOfUser(arg0 common.Address) (*big.Int, error) {
	return _NodeManagerV83.Contract.OldRewardsOfUser(&_NodeManagerV83.CallOpts, arg0)
}

// OperationsPoolAddress is a free data retrieval call binding the contract method 0x919df198.
//
// Solidity: function operationsPoolAddress() view returns(address)
func (_NodeManagerV83 *NodeManagerV83Caller) OperationsPoolAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NodeManagerV83.contract.Call(opts, &out, "operationsPoolAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OperationsPoolAddress is a free data retrieval call binding the contract method 0x919df198.
//
// Solidity: function operationsPoolAddress() view returns(address)
func (_NodeManagerV83 *NodeManagerV83Session) OperationsPoolAddress() (common.Address, error) {
	return _NodeManagerV83.Contract.OperationsPoolAddress(&_NodeManagerV83.CallOpts)
}

// OperationsPoolAddress is a free data retrieval call binding the contract method 0x919df198.
//
// Solidity: function operationsPoolAddress() view returns(address)
func (_NodeManagerV83 *NodeManagerV83CallerSession) OperationsPoolAddress() (common.Address, error) {
	return _NodeManagerV83.Contract.OperationsPoolAddress(&_NodeManagerV83.CallOpts)
}

// OperationsPoolFee is a free data retrieval call binding the contract method 0x3ba28fcb.
//
// Solidity: function operationsPoolFee() view returns(uint32)
func (_NodeManagerV83 *NodeManagerV83Caller) OperationsPoolFee(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _NodeManagerV83.contract.Call(opts, &out, "operationsPoolFee")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// OperationsPoolFee is a free data retrieval call binding the contract method 0x3ba28fcb.
//
// Solidity: function operationsPoolFee() view returns(uint32)
func (_NodeManagerV83 *NodeManagerV83Session) OperationsPoolFee() (uint32, error) {
	return _NodeManagerV83.Contract.OperationsPoolFee(&_NodeManagerV83.CallOpts)
}

// OperationsPoolFee is a free data retrieval call binding the contract method 0x3ba28fcb.
//
// Solidity: function operationsPoolFee() view returns(uint32)
func (_NodeManagerV83 *NodeManagerV83CallerSession) OperationsPoolFee() (uint32, error) {
	return _NodeManagerV83.Contract.OperationsPoolFee(&_NodeManagerV83.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NodeManagerV83 *NodeManagerV83Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NodeManagerV83.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NodeManagerV83 *NodeManagerV83Session) Owner() (common.Address, error) {
	return _NodeManagerV83.Contract.Owner(&_NodeManagerV83.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NodeManagerV83 *NodeManagerV83CallerSession) Owner() (common.Address, error) {
	return _NodeManagerV83.Contract.Owner(&_NodeManagerV83.CallOpts)
}

// PayInterval is a free data retrieval call binding the contract method 0x8e6a1bc2.
//
// Solidity: function payInterval() view returns(uint32)
func (_NodeManagerV83 *NodeManagerV83Caller) PayInterval(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _NodeManagerV83.contract.Call(opts, &out, "payInterval")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// PayInterval is a free data retrieval call binding the contract method 0x8e6a1bc2.
//
// Solidity: function payInterval() view returns(uint32)
func (_NodeManagerV83 *NodeManagerV83Session) PayInterval() (uint32, error) {
	return _NodeManagerV83.Contract.PayInterval(&_NodeManagerV83.CallOpts)
}

// PayInterval is a free data retrieval call binding the contract method 0x8e6a1bc2.
//
// Solidity: function payInterval() view returns(uint32)
func (_NodeManagerV83 *NodeManagerV83CallerSession) PayInterval() (uint32, error) {
	return _NodeManagerV83.Contract.PayInterval(&_NodeManagerV83.CallOpts)
}

// RateClaimFee is a free data retrieval call binding the contract method 0x501f9477.
//
// Solidity: function rateClaimFee(address account) view returns(uint32)
func (_NodeManagerV83 *NodeManagerV83Caller) RateClaimFee(opts *bind.CallOpts, account common.Address) (uint32, error) {
	var out []interface{}
	err := _NodeManagerV83.contract.Call(opts, &out, "rateClaimFee", account)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// RateClaimFee is a free data retrieval call binding the contract method 0x501f9477.
//
// Solidity: function rateClaimFee(address account) view returns(uint32)
func (_NodeManagerV83 *NodeManagerV83Session) RateClaimFee(account common.Address) (uint32, error) {
	return _NodeManagerV83.Contract.RateClaimFee(&_NodeManagerV83.CallOpts, account)
}

// RateClaimFee is a free data retrieval call binding the contract method 0x501f9477.
//
// Solidity: function rateClaimFee(address account) view returns(uint32)
func (_NodeManagerV83 *NodeManagerV83CallerSession) RateClaimFee(account common.Address) (uint32, error) {
	return _NodeManagerV83.Contract.RateClaimFee(&_NodeManagerV83.CallOpts, account)
}

// RewardMigrated is a free data retrieval call binding the contract method 0x253ad3ba.
//
// Solidity: function rewardMigrated(address ) view returns(bool)
func (_NodeManagerV83 *NodeManagerV83Caller) RewardMigrated(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _NodeManagerV83.contract.Call(opts, &out, "rewardMigrated", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RewardMigrated is a free data retrieval call binding the contract method 0x253ad3ba.
//
// Solidity: function rewardMigrated(address ) view returns(bool)
func (_NodeManagerV83 *NodeManagerV83Session) RewardMigrated(arg0 common.Address) (bool, error) {
	return _NodeManagerV83.Contract.RewardMigrated(&_NodeManagerV83.CallOpts, arg0)
}

// RewardMigrated is a free data retrieval call binding the contract method 0x253ad3ba.
//
// Solidity: function rewardMigrated(address ) view returns(bool)
func (_NodeManagerV83 *NodeManagerV83CallerSession) RewardMigrated(arg0 common.Address) (bool, error) {
	return _NodeManagerV83.Contract.RewardMigrated(&_NodeManagerV83.CallOpts, arg0)
}

// RewardsOfUser is a free data retrieval call binding the contract method 0xb2dd3b37.
//
// Solidity: function rewardsOfUser(address ) view returns(uint256)
func (_NodeManagerV83 *NodeManagerV83Caller) RewardsOfUser(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NodeManagerV83.contract.Call(opts, &out, "rewardsOfUser", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardsOfUser is a free data retrieval call binding the contract method 0xb2dd3b37.
//
// Solidity: function rewardsOfUser(address ) view returns(uint256)
func (_NodeManagerV83 *NodeManagerV83Session) RewardsOfUser(arg0 common.Address) (*big.Int, error) {
	return _NodeManagerV83.Contract.RewardsOfUser(&_NodeManagerV83.CallOpts, arg0)
}

// RewardsOfUser is a free data retrieval call binding the contract method 0xb2dd3b37.
//
// Solidity: function rewardsOfUser(address ) view returns(uint256)
func (_NodeManagerV83 *NodeManagerV83CallerSession) RewardsOfUser(arg0 common.Address) (*big.Int, error) {
	return _NodeManagerV83.Contract.RewardsOfUser(&_NodeManagerV83.CallOpts, arg0)
}

// RewardsPoolAddress is a free data retrieval call binding the contract method 0xcce85055.
//
// Solidity: function rewardsPoolAddress() view returns(address)
func (_NodeManagerV83 *NodeManagerV83Caller) RewardsPoolAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NodeManagerV83.contract.Call(opts, &out, "rewardsPoolAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardsPoolAddress is a free data retrieval call binding the contract method 0xcce85055.
//
// Solidity: function rewardsPoolAddress() view returns(address)
func (_NodeManagerV83 *NodeManagerV83Session) RewardsPoolAddress() (common.Address, error) {
	return _NodeManagerV83.Contract.RewardsPoolAddress(&_NodeManagerV83.CallOpts)
}

// RewardsPoolAddress is a free data retrieval call binding the contract method 0xcce85055.
//
// Solidity: function rewardsPoolAddress() view returns(address)
func (_NodeManagerV83 *NodeManagerV83CallerSession) RewardsPoolAddress() (common.Address, error) {
	return _NodeManagerV83.Contract.RewardsPoolAddress(&_NodeManagerV83.CallOpts)
}

// RewardsPoolFee is a free data retrieval call binding the contract method 0x864bbf59.
//
// Solidity: function rewardsPoolFee() view returns(uint32)
func (_NodeManagerV83 *NodeManagerV83Caller) RewardsPoolFee(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _NodeManagerV83.contract.Call(opts, &out, "rewardsPoolFee")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// RewardsPoolFee is a free data retrieval call binding the contract method 0x864bbf59.
//
// Solidity: function rewardsPoolFee() view returns(uint32)
func (_NodeManagerV83 *NodeManagerV83Session) RewardsPoolFee() (uint32, error) {
	return _NodeManagerV83.Contract.RewardsPoolFee(&_NodeManagerV83.CallOpts)
}

// RewardsPoolFee is a free data retrieval call binding the contract method 0x864bbf59.
//
// Solidity: function rewardsPoolFee() view returns(uint32)
func (_NodeManagerV83 *NodeManagerV83CallerSession) RewardsPoolFee() (uint32, error) {
	return _NodeManagerV83.Contract.RewardsPoolFee(&_NodeManagerV83.CallOpts)
}

// RewardsTotal is a free data retrieval call binding the contract method 0xbb2d7f3a.
//
// Solidity: function rewardsTotal() view returns(uint256)
func (_NodeManagerV83 *NodeManagerV83Caller) RewardsTotal(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NodeManagerV83.contract.Call(opts, &out, "rewardsTotal")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardsTotal is a free data retrieval call binding the contract method 0xbb2d7f3a.
//
// Solidity: function rewardsTotal() view returns(uint256)
func (_NodeManagerV83 *NodeManagerV83Session) RewardsTotal() (*big.Int, error) {
	return _NodeManagerV83.Contract.RewardsTotal(&_NodeManagerV83.CallOpts)
}

// RewardsTotal is a free data retrieval call binding the contract method 0xbb2d7f3a.
//
// Solidity: function rewardsTotal() view returns(uint256)
func (_NodeManagerV83 *NodeManagerV83CallerSession) RewardsTotal() (*big.Int, error) {
	return _NodeManagerV83.Contract.RewardsTotal(&_NodeManagerV83.CallOpts)
}

// SellPricePercent is a free data retrieval call binding the contract method 0x5a64affb.
//
// Solidity: function sellPricePercent() view returns(uint32)
func (_NodeManagerV83 *NodeManagerV83Caller) SellPricePercent(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _NodeManagerV83.contract.Call(opts, &out, "sellPricePercent")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// SellPricePercent is a free data retrieval call binding the contract method 0x5a64affb.
//
// Solidity: function sellPricePercent() view returns(uint32)
func (_NodeManagerV83 *NodeManagerV83Session) SellPricePercent() (uint32, error) {
	return _NodeManagerV83.Contract.SellPricePercent(&_NodeManagerV83.CallOpts)
}

// SellPricePercent is a free data retrieval call binding the contract method 0x5a64affb.
//
// Solidity: function sellPricePercent() view returns(uint32)
func (_NodeManagerV83 *NodeManagerV83CallerSession) SellPricePercent() (uint32, error) {
	return _NodeManagerV83.Contract.SellPricePercent(&_NodeManagerV83.CallOpts)
}

// TierMap is a free data retrieval call binding the contract method 0x530ee54b.
//
// Solidity: function tierMap(string ) view returns(uint8)
func (_NodeManagerV83 *NodeManagerV83Caller) TierMap(opts *bind.CallOpts, arg0 string) (uint8, error) {
	var out []interface{}
	err := _NodeManagerV83.contract.Call(opts, &out, "tierMap", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TierMap is a free data retrieval call binding the contract method 0x530ee54b.
//
// Solidity: function tierMap(string ) view returns(uint8)
func (_NodeManagerV83 *NodeManagerV83Session) TierMap(arg0 string) (uint8, error) {
	return _NodeManagerV83.Contract.TierMap(&_NodeManagerV83.CallOpts, arg0)
}

// TierMap is a free data retrieval call binding the contract method 0x530ee54b.
//
// Solidity: function tierMap(string ) view returns(uint8)
func (_NodeManagerV83 *NodeManagerV83CallerSession) TierMap(arg0 string) (uint8, error) {
	return _NodeManagerV83.Contract.TierMap(&_NodeManagerV83.CallOpts, arg0)
}

// TierTotal is a free data retrieval call binding the contract method 0x79420fb1.
//
// Solidity: function tierTotal() view returns(uint8)
func (_NodeManagerV83 *NodeManagerV83Caller) TierTotal(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _NodeManagerV83.contract.Call(opts, &out, "tierTotal")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TierTotal is a free data retrieval call binding the contract method 0x79420fb1.
//
// Solidity: function tierTotal() view returns(uint8)
func (_NodeManagerV83 *NodeManagerV83Session) TierTotal() (uint8, error) {
	return _NodeManagerV83.Contract.TierTotal(&_NodeManagerV83.CallOpts)
}

// TierTotal is a free data retrieval call binding the contract method 0x79420fb1.
//
// Solidity: function tierTotal() view returns(uint8)
func (_NodeManagerV83 *NodeManagerV83CallerSession) TierTotal() (uint8, error) {
	return _NodeManagerV83.Contract.TierTotal(&_NodeManagerV83.CallOpts)
}

// Tiers is a free data retrieval call binding the contract method 0x4a95d9d5.
//
// Solidity: function tiers() view returns((uint8,string,uint256,uint256,uint32,uint256)[])
func (_NodeManagerV83 *NodeManagerV83Caller) Tiers(opts *bind.CallOpts) ([]Tier, error) {
	var out []interface{}
	err := _NodeManagerV83.contract.Call(opts, &out, "tiers")

	if err != nil {
		return *new([]Tier), err
	}

	out0 := *abi.ConvertType(out[0], new([]Tier)).(*[]Tier)

	return out0, err

}

// Tiers is a free data retrieval call binding the contract method 0x4a95d9d5.
//
// Solidity: function tiers() view returns((uint8,string,uint256,uint256,uint32,uint256)[])
func (_NodeManagerV83 *NodeManagerV83Session) Tiers() ([]Tier, error) {
	return _NodeManagerV83.Contract.Tiers(&_NodeManagerV83.CallOpts)
}

// Tiers is a free data retrieval call binding the contract method 0x4a95d9d5.
//
// Solidity: function tiers() view returns((uint8,string,uint256,uint256,uint32,uint256)[])
func (_NodeManagerV83 *NodeManagerV83CallerSession) Tiers() ([]Tier, error) {
	return _NodeManagerV83.Contract.Tiers(&_NodeManagerV83.CallOpts)
}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() view returns(address)
func (_NodeManagerV83 *NodeManagerV83Caller) TokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NodeManagerV83.contract.Call(opts, &out, "tokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() view returns(address)
func (_NodeManagerV83 *NodeManagerV83Session) TokenAddress() (common.Address, error) {
	return _NodeManagerV83.Contract.TokenAddress(&_NodeManagerV83.CallOpts)
}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() view returns(address)
func (_NodeManagerV83 *NodeManagerV83CallerSession) TokenAddress() (common.Address, error) {
	return _NodeManagerV83.Contract.TokenAddress(&_NodeManagerV83.CallOpts)
}

// TransferFee is a free data retrieval call binding the contract method 0xacb2ad6f.
//
// Solidity: function transferFee() view returns(uint32)
func (_NodeManagerV83 *NodeManagerV83Caller) TransferFee(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _NodeManagerV83.contract.Call(opts, &out, "transferFee")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TransferFee is a free data retrieval call binding the contract method 0xacb2ad6f.
//
// Solidity: function transferFee() view returns(uint32)
func (_NodeManagerV83 *NodeManagerV83Session) TransferFee() (uint32, error) {
	return _NodeManagerV83.Contract.TransferFee(&_NodeManagerV83.CallOpts)
}

// TransferFee is a free data retrieval call binding the contract method 0xacb2ad6f.
//
// Solidity: function transferFee() view returns(uint32)
func (_NodeManagerV83 *NodeManagerV83CallerSession) TransferFee() (uint32, error) {
	return _NodeManagerV83.Contract.TransferFee(&_NodeManagerV83.CallOpts)
}

// UniswapV2Router is a free data retrieval call binding the contract method 0x1694505e.
//
// Solidity: function uniswapV2Router() view returns(address)
func (_NodeManagerV83 *NodeManagerV83Caller) UniswapV2Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NodeManagerV83.contract.Call(opts, &out, "uniswapV2Router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UniswapV2Router is a free data retrieval call binding the contract method 0x1694505e.
//
// Solidity: function uniswapV2Router() view returns(address)
func (_NodeManagerV83 *NodeManagerV83Session) UniswapV2Router() (common.Address, error) {
	return _NodeManagerV83.Contract.UniswapV2Router(&_NodeManagerV83.CallOpts)
}

// UniswapV2Router is a free data retrieval call binding the contract method 0x1694505e.
//
// Solidity: function uniswapV2Router() view returns(address)
func (_NodeManagerV83 *NodeManagerV83CallerSession) UniswapV2Router() (common.Address, error) {
	return _NodeManagerV83.Contract.UniswapV2Router(&_NodeManagerV83.CallOpts)
}

// UnpaidNodes is a free data retrieval call binding the contract method 0xfba6c21f.
//
// Solidity: function unpaidNodes() view returns((uint32,uint8,address,uint32,uint32,uint32,uint256,uint256)[])
func (_NodeManagerV83 *NodeManagerV83Caller) UnpaidNodes(opts *bind.CallOpts) ([]Node, error) {
	var out []interface{}
	err := _NodeManagerV83.contract.Call(opts, &out, "unpaidNodes")

	if err != nil {
		return *new([]Node), err
	}

	out0 := *abi.ConvertType(out[0], new([]Node)).(*[]Node)

	return out0, err

}

// UnpaidNodes is a free data retrieval call binding the contract method 0xfba6c21f.
//
// Solidity: function unpaidNodes() view returns((uint32,uint8,address,uint32,uint32,uint32,uint256,uint256)[])
func (_NodeManagerV83 *NodeManagerV83Session) UnpaidNodes() ([]Node, error) {
	return _NodeManagerV83.Contract.UnpaidNodes(&_NodeManagerV83.CallOpts)
}

// UnpaidNodes is a free data retrieval call binding the contract method 0xfba6c21f.
//
// Solidity: function unpaidNodes() view returns((uint32,uint8,address,uint32,uint32,uint32,uint256,uint256)[])
func (_NodeManagerV83 *NodeManagerV83CallerSession) UnpaidNodes() ([]Node, error) {
	return _NodeManagerV83.Contract.UnpaidNodes(&_NodeManagerV83.CallOpts)
}

// UserMigrated is a free data retrieval call binding the contract method 0x0c6afded.
//
// Solidity: function userMigrated(address ) view returns(bool)
func (_NodeManagerV83 *NodeManagerV83Caller) UserMigrated(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _NodeManagerV83.contract.Call(opts, &out, "userMigrated", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UserMigrated is a free data retrieval call binding the contract method 0x0c6afded.
//
// Solidity: function userMigrated(address ) view returns(bool)
func (_NodeManagerV83 *NodeManagerV83Session) UserMigrated(arg0 common.Address) (bool, error) {
	return _NodeManagerV83.Contract.UserMigrated(&_NodeManagerV83.CallOpts, arg0)
}

// UserMigrated is a free data retrieval call binding the contract method 0x0c6afded.
//
// Solidity: function userMigrated(address ) view returns(bool)
func (_NodeManagerV83 *NodeManagerV83CallerSession) UserMigrated(arg0 common.Address) (bool, error) {
	return _NodeManagerV83.Contract.UserMigrated(&_NodeManagerV83.CallOpts, arg0)
}

// WithdrawRate is a free data retrieval call binding the contract method 0x0dcf1417.
//
// Solidity: function withdrawRate() view returns(uint32)
func (_NodeManagerV83 *NodeManagerV83Caller) WithdrawRate(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _NodeManagerV83.contract.Call(opts, &out, "withdrawRate")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// WithdrawRate is a free data retrieval call binding the contract method 0x0dcf1417.
//
// Solidity: function withdrawRate() view returns(uint32)
func (_NodeManagerV83 *NodeManagerV83Session) WithdrawRate() (uint32, error) {
	return _NodeManagerV83.Contract.WithdrawRate(&_NodeManagerV83.CallOpts)
}

// WithdrawRate is a free data retrieval call binding the contract method 0x0dcf1417.
//
// Solidity: function withdrawRate() view returns(uint32)
func (_NodeManagerV83 *NodeManagerV83CallerSession) WithdrawRate() (uint32, error) {
	return _NodeManagerV83.Contract.WithdrawRate(&_NodeManagerV83.CallOpts)
}

// AddTier is a paid mutator transaction binding the contract method 0x85b8eb44.
//
// Solidity: function addTier(string name, uint256 price, uint256 rewardsPerTime, uint32 claimInterval, uint256 maintenanceFee) returns()
func (_NodeManagerV83 *NodeManagerV83Transactor) AddTier(opts *bind.TransactOpts, name string, price *big.Int, rewardsPerTime *big.Int, claimInterval uint32, maintenanceFee *big.Int) (*types.Transaction, error) {
	return _NodeManagerV83.contract.Transact(opts, "addTier", name, price, rewardsPerTime, claimInterval, maintenanceFee)
}

// AddTier is a paid mutator transaction binding the contract method 0x85b8eb44.
//
// Solidity: function addTier(string name, uint256 price, uint256 rewardsPerTime, uint32 claimInterval, uint256 maintenanceFee) returns()
func (_NodeManagerV83 *NodeManagerV83Session) AddTier(name string, price *big.Int, rewardsPerTime *big.Int, claimInterval uint32, maintenanceFee *big.Int) (*types.Transaction, error) {
	return _NodeManagerV83.Contract.AddTier(&_NodeManagerV83.TransactOpts, name, price, rewardsPerTime, claimInterval, maintenanceFee)
}

// AddTier is a paid mutator transaction binding the contract method 0x85b8eb44.
//
// Solidity: function addTier(string name, uint256 price, uint256 rewardsPerTime, uint32 claimInterval, uint256 maintenanceFee) returns()
func (_NodeManagerV83 *NodeManagerV83TransactorSession) AddTier(name string, price *big.Int, rewardsPerTime *big.Int, claimInterval uint32, maintenanceFee *big.Int) (*types.Transaction, error) {
	return _NodeManagerV83.Contract.AddTier(&_NodeManagerV83.TransactOpts, name, price, rewardsPerTime, claimInterval, maintenanceFee)
}

// BindBoostNFT is a paid mutator transaction binding the contract method 0x55953d25.
//
// Solidity: function bindBoostNFT(address _nftAddress) returns()
func (_NodeManagerV83 *NodeManagerV83Transactor) BindBoostNFT(opts *bind.TransactOpts, _nftAddress common.Address) (*types.Transaction, error) {
	return _NodeManagerV83.contract.Transact(opts, "bindBoostNFT", _nftAddress)
}

// BindBoostNFT is a paid mutator transaction binding the contract method 0x55953d25.
//
// Solidity: function bindBoostNFT(address _nftAddress) returns()
func (_NodeManagerV83 *NodeManagerV83Session) BindBoostNFT(_nftAddress common.Address) (*types.Transaction, error) {
	return _NodeManagerV83.Contract.BindBoostNFT(&_NodeManagerV83.TransactOpts, _nftAddress)
}

// BindBoostNFT is a paid mutator transaction binding the contract method 0x55953d25.
//
// Solidity: function bindBoostNFT(address _nftAddress) returns()
func (_NodeManagerV83 *NodeManagerV83TransactorSession) BindBoostNFT(_nftAddress common.Address) (*types.Transaction, error) {
	return _NodeManagerV83.Contract.BindBoostNFT(&_NodeManagerV83.TransactOpts, _nftAddress)
}

// BurnNodes is a paid mutator transaction binding the contract method 0x2dbc78e8.
//
// Solidity: function burnNodes(uint32[] indice) returns()
func (_NodeManagerV83 *NodeManagerV83Transactor) BurnNodes(opts *bind.TransactOpts, indice []uint32) (*types.Transaction, error) {
	return _NodeManagerV83.contract.Transact(opts, "burnNodes", indice)
}

// BurnNodes is a paid mutator transaction binding the contract method 0x2dbc78e8.
//
// Solidity: function burnNodes(uint32[] indice) returns()
func (_NodeManagerV83 *NodeManagerV83Session) BurnNodes(indice []uint32) (*types.Transaction, error) {
	return _NodeManagerV83.Contract.BurnNodes(&_NodeManagerV83.TransactOpts, indice)
}

// BurnNodes is a paid mutator transaction binding the contract method 0x2dbc78e8.
//
// Solidity: function burnNodes(uint32[] indice) returns()
func (_NodeManagerV83 *NodeManagerV83TransactorSession) BurnNodes(indice []uint32) (*types.Transaction, error) {
	return _NodeManagerV83.Contract.BurnNodes(&_NodeManagerV83.TransactOpts, indice)
}

// BurnUser is a paid mutator transaction binding the contract method 0x5bcb11bb.
//
// Solidity: function burnUser(address account) returns()
func (_NodeManagerV83 *NodeManagerV83Transactor) BurnUser(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _NodeManagerV83.contract.Transact(opts, "burnUser", account)
}

// BurnUser is a paid mutator transaction binding the contract method 0x5bcb11bb.
//
// Solidity: function burnUser(address account) returns()
func (_NodeManagerV83 *NodeManagerV83Session) BurnUser(account common.Address) (*types.Transaction, error) {
	return _NodeManagerV83.Contract.BurnUser(&_NodeManagerV83.TransactOpts, account)
}

// BurnUser is a paid mutator transaction binding the contract method 0x5bcb11bb.
//
// Solidity: function burnUser(address account) returns()
func (_NodeManagerV83 *NodeManagerV83TransactorSession) BurnUser(account common.Address) (*types.Transaction, error) {
	return _NodeManagerV83.Contract.BurnUser(&_NodeManagerV83.TransactOpts, account)
}

// BuyGiftCard is a paid mutator transaction binding the contract method 0x20b242ec.
//
// Solidity: function buyGiftCard(address token, string orderID, uint256 mode, uint256 amount) payable returns()
func (_NodeManagerV83 *NodeManagerV83Transactor) BuyGiftCard(opts *bind.TransactOpts, token common.Address, orderID string, mode *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _NodeManagerV83.contract.Transact(opts, "buyGiftCard", token, orderID, mode, amount)
}

// BuyGiftCard is a paid mutator transaction binding the contract method 0x20b242ec.
//
// Solidity: function buyGiftCard(address token, string orderID, uint256 mode, uint256 amount) payable returns()
func (_NodeManagerV83 *NodeManagerV83Session) BuyGiftCard(token common.Address, orderID string, mode *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _NodeManagerV83.Contract.BuyGiftCard(&_NodeManagerV83.TransactOpts, token, orderID, mode, amount)
}

// BuyGiftCard is a paid mutator transaction binding the contract method 0x20b242ec.
//
// Solidity: function buyGiftCard(address token, string orderID, uint256 mode, uint256 amount) payable returns()
func (_NodeManagerV83 *NodeManagerV83TransactorSession) BuyGiftCard(token common.Address, orderID string, mode *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _NodeManagerV83.Contract.BuyGiftCard(&_NodeManagerV83.TransactOpts, token, orderID, mode, amount)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_NodeManagerV83 *NodeManagerV83Transactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeManagerV83.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_NodeManagerV83 *NodeManagerV83Session) Claim() (*types.Transaction, error) {
	return _NodeManagerV83.Contract.Claim(&_NodeManagerV83.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_NodeManagerV83 *NodeManagerV83TransactorSession) Claim() (*types.Transaction, error) {
	return _NodeManagerV83.Contract.Claim(&_NodeManagerV83.TransactOpts)
}

// Compound is a paid mutator transaction binding the contract method 0x7e7b0d57.
//
// Solidity: function compound(string tierName, uint32 count) returns()
func (_NodeManagerV83 *NodeManagerV83Transactor) Compound(opts *bind.TransactOpts, tierName string, count uint32) (*types.Transaction, error) {
	return _NodeManagerV83.contract.Transact(opts, "compound", tierName, count)
}

// Compound is a paid mutator transaction binding the contract method 0x7e7b0d57.
//
// Solidity: function compound(string tierName, uint32 count) returns()
func (_NodeManagerV83 *NodeManagerV83Session) Compound(tierName string, count uint32) (*types.Transaction, error) {
	return _NodeManagerV83.Contract.Compound(&_NodeManagerV83.TransactOpts, tierName, count)
}

// Compound is a paid mutator transaction binding the contract method 0x7e7b0d57.
//
// Solidity: function compound(string tierName, uint32 count) returns()
func (_NodeManagerV83 *NodeManagerV83TransactorSession) Compound(tierName string, count uint32) (*types.Transaction, error) {
	return _NodeManagerV83.Contract.Compound(&_NodeManagerV83.TransactOpts, tierName, count)
}

// Create is a paid mutator transaction binding the contract method 0x9c43b6e2.
//
// Solidity: function create(string tierName, uint32 count) returns()
func (_NodeManagerV83 *NodeManagerV83Transactor) Create(opts *bind.TransactOpts, tierName string, count uint32) (*types.Transaction, error) {
	return _NodeManagerV83.contract.Transact(opts, "create", tierName, count)
}

// Create is a paid mutator transaction binding the contract method 0x9c43b6e2.
//
// Solidity: function create(string tierName, uint32 count) returns()
func (_NodeManagerV83 *NodeManagerV83Session) Create(tierName string, count uint32) (*types.Transaction, error) {
	return _NodeManagerV83.Contract.Create(&_NodeManagerV83.TransactOpts, tierName, count)
}

// Create is a paid mutator transaction binding the contract method 0x9c43b6e2.
//
// Solidity: function create(string tierName, uint32 count) returns()
func (_NodeManagerV83 *NodeManagerV83TransactorSession) Create(tierName string, count uint32) (*types.Transaction, error) {
	return _NodeManagerV83.Contract.Create(&_NodeManagerV83.TransactOpts, tierName, count)
}

// Initialize is a paid mutator transaction binding the contract method 0xa224cee7.
//
// Solidity: function initialize(address[] addresses) returns()
func (_NodeManagerV83 *NodeManagerV83Transactor) Initialize(opts *bind.TransactOpts, addresses []common.Address) (*types.Transaction, error) {
	return _NodeManagerV83.contract.Transact(opts, "initialize", addresses)
}

// Initialize is a paid mutator transaction binding the contract method 0xa224cee7.
//
// Solidity: function initialize(address[] addresses) returns()
func (_NodeManagerV83 *NodeManagerV83Session) Initialize(addresses []common.Address) (*types.Transaction, error) {
	return _NodeManagerV83.Contract.Initialize(&_NodeManagerV83.TransactOpts, addresses)
}

// Initialize is a paid mutator transaction binding the contract method 0xa224cee7.
//
// Solidity: function initialize(address[] addresses) returns()
func (_NodeManagerV83 *NodeManagerV83TransactorSession) Initialize(addresses []common.Address) (*types.Transaction, error) {
	return _NodeManagerV83.Contract.Initialize(&_NodeManagerV83.TransactOpts, addresses)
}

// MigrateNodesFromOldVersion is a paid mutator transaction binding the contract method 0x784eceef.
//
// Solidity: function migrateNodesFromOldVersion() returns()
func (_NodeManagerV83 *NodeManagerV83Transactor) MigrateNodesFromOldVersion(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeManagerV83.contract.Transact(opts, "migrateNodesFromOldVersion")
}

// MigrateNodesFromOldVersion is a paid mutator transaction binding the contract method 0x784eceef.
//
// Solidity: function migrateNodesFromOldVersion() returns()
func (_NodeManagerV83 *NodeManagerV83Session) MigrateNodesFromOldVersion() (*types.Transaction, error) {
	return _NodeManagerV83.Contract.MigrateNodesFromOldVersion(&_NodeManagerV83.TransactOpts)
}

// MigrateNodesFromOldVersion is a paid mutator transaction binding the contract method 0x784eceef.
//
// Solidity: function migrateNodesFromOldVersion() returns()
func (_NodeManagerV83 *NodeManagerV83TransactorSession) MigrateNodesFromOldVersion() (*types.Transaction, error) {
	return _NodeManagerV83.Contract.MigrateNodesFromOldVersion(&_NodeManagerV83.TransactOpts)
}

// MigrateRewardsFromOldVersion is a paid mutator transaction binding the contract method 0x05ec8cec.
//
// Solidity: function migrateRewardsFromOldVersion() returns()
func (_NodeManagerV83 *NodeManagerV83Transactor) MigrateRewardsFromOldVersion(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeManagerV83.contract.Transact(opts, "migrateRewardsFromOldVersion")
}

// MigrateRewardsFromOldVersion is a paid mutator transaction binding the contract method 0x05ec8cec.
//
// Solidity: function migrateRewardsFromOldVersion() returns()
func (_NodeManagerV83 *NodeManagerV83Session) MigrateRewardsFromOldVersion() (*types.Transaction, error) {
	return _NodeManagerV83.Contract.MigrateRewardsFromOldVersion(&_NodeManagerV83.TransactOpts)
}

// MigrateRewardsFromOldVersion is a paid mutator transaction binding the contract method 0x05ec8cec.
//
// Solidity: function migrateRewardsFromOldVersion() returns()
func (_NodeManagerV83 *NodeManagerV83TransactorSession) MigrateRewardsFromOldVersion() (*types.Transaction, error) {
	return _NodeManagerV83.Contract.MigrateRewardsFromOldVersion(&_NodeManagerV83.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x4ee571ce.
//
// Solidity: function mint(address[] accounts, string tierName, uint32 count) returns()
func (_NodeManagerV83 *NodeManagerV83Transactor) Mint(opts *bind.TransactOpts, accounts []common.Address, tierName string, count uint32) (*types.Transaction, error) {
	return _NodeManagerV83.contract.Transact(opts, "mint", accounts, tierName, count)
}

// Mint is a paid mutator transaction binding the contract method 0x4ee571ce.
//
// Solidity: function mint(address[] accounts, string tierName, uint32 count) returns()
func (_NodeManagerV83 *NodeManagerV83Session) Mint(accounts []common.Address, tierName string, count uint32) (*types.Transaction, error) {
	return _NodeManagerV83.Contract.Mint(&_NodeManagerV83.TransactOpts, accounts, tierName, count)
}

// Mint is a paid mutator transaction binding the contract method 0x4ee571ce.
//
// Solidity: function mint(address[] accounts, string tierName, uint32 count) returns()
func (_NodeManagerV83 *NodeManagerV83TransactorSession) Mint(accounts []common.Address, tierName string, count uint32) (*types.Transaction, error) {
	return _NodeManagerV83.Contract.Mint(&_NodeManagerV83.TransactOpts, accounts, tierName, count)
}

// Pay is a paid mutator transaction binding the contract method 0x9ea164b3.
//
// Solidity: function pay(uint8 count, uint256[] selected) payable returns()
func (_NodeManagerV83 *NodeManagerV83Transactor) Pay(opts *bind.TransactOpts, count uint8, selected []*big.Int) (*types.Transaction, error) {
	return _NodeManagerV83.contract.Transact(opts, "pay", count, selected)
}

// Pay is a paid mutator transaction binding the contract method 0x9ea164b3.
//
// Solidity: function pay(uint8 count, uint256[] selected) payable returns()
func (_NodeManagerV83 *NodeManagerV83Session) Pay(count uint8, selected []*big.Int) (*types.Transaction, error) {
	return _NodeManagerV83.Contract.Pay(&_NodeManagerV83.TransactOpts, count, selected)
}

// Pay is a paid mutator transaction binding the contract method 0x9ea164b3.
//
// Solidity: function pay(uint8 count, uint256[] selected) payable returns()
func (_NodeManagerV83 *NodeManagerV83TransactorSession) Pay(count uint8, selected []*big.Int) (*types.Transaction, error) {
	return _NodeManagerV83.Contract.Pay(&_NodeManagerV83.TransactOpts, count, selected)
}

// RemoveTier is a paid mutator transaction binding the contract method 0xf2d049b6.
//
// Solidity: function removeTier(string tierName) returns()
func (_NodeManagerV83 *NodeManagerV83Transactor) RemoveTier(opts *bind.TransactOpts, tierName string) (*types.Transaction, error) {
	return _NodeManagerV83.contract.Transact(opts, "removeTier", tierName)
}

// RemoveTier is a paid mutator transaction binding the contract method 0xf2d049b6.
//
// Solidity: function removeTier(string tierName) returns()
func (_NodeManagerV83 *NodeManagerV83Session) RemoveTier(tierName string) (*types.Transaction, error) {
	return _NodeManagerV83.Contract.RemoveTier(&_NodeManagerV83.TransactOpts, tierName)
}

// RemoveTier is a paid mutator transaction binding the contract method 0xf2d049b6.
//
// Solidity: function removeTier(string tierName) returns()
func (_NodeManagerV83 *NodeManagerV83TransactorSession) RemoveTier(tierName string) (*types.Transaction, error) {
	return _NodeManagerV83.Contract.RemoveTier(&_NodeManagerV83.TransactOpts, tierName)
}

// SetAddressInBlacklist is a paid mutator transaction binding the contract method 0xb1ee5a81.
//
// Solidity: function setAddressInBlacklist(address walletAddress, bool value) returns()
func (_NodeManagerV83 *NodeManagerV83Transactor) SetAddressInBlacklist(opts *bind.TransactOpts, walletAddress common.Address, value bool) (*types.Transaction, error) {
	return _NodeManagerV83.contract.Transact(opts, "setAddressInBlacklist", walletAddress, value)
}

// SetAddressInBlacklist is a paid mutator transaction binding the contract method 0xb1ee5a81.
//
// Solidity: function setAddressInBlacklist(address walletAddress, bool value) returns()
func (_NodeManagerV83 *NodeManagerV83Session) SetAddressInBlacklist(walletAddress common.Address, value bool) (*types.Transaction, error) {
	return _NodeManagerV83.Contract.SetAddressInBlacklist(&_NodeManagerV83.TransactOpts, walletAddress, value)
}

// SetAddressInBlacklist is a paid mutator transaction binding the contract method 0xb1ee5a81.
//
// Solidity: function setAddressInBlacklist(address walletAddress, bool value) returns()
func (_NodeManagerV83 *NodeManagerV83TransactorSession) SetAddressInBlacklist(walletAddress common.Address, value bool) (*types.Transaction, error) {
	return _NodeManagerV83.Contract.SetAddressInBlacklist(&_NodeManagerV83.TransactOpts, walletAddress, value)
}

// SetClaimFee is a paid mutator transaction binding the contract method 0x81c31cf9.
//
// Solidity: function setClaimFee(uint32 value) returns()
func (_NodeManagerV83 *NodeManagerV83Transactor) SetClaimFee(opts *bind.TransactOpts, value uint32) (*types.Transaction, error) {
	return _NodeManagerV83.contract.Transact(opts, "setClaimFee", value)
}

// SetClaimFee is a paid mutator transaction binding the contract method 0x81c31cf9.
//
// Solidity: function setClaimFee(uint32 value) returns()
func (_NodeManagerV83 *NodeManagerV83Session) SetClaimFee(value uint32) (*types.Transaction, error) {
	return _NodeManagerV83.Contract.SetClaimFee(&_NodeManagerV83.TransactOpts, value)
}

// SetClaimFee is a paid mutator transaction binding the contract method 0x81c31cf9.
//
// Solidity: function setClaimFee(uint32 value) returns()
func (_NodeManagerV83 *NodeManagerV83TransactorSession) SetClaimFee(value uint32) (*types.Transaction, error) {
	return _NodeManagerV83.Contract.SetClaimFee(&_NodeManagerV83.TransactOpts, value)
}

// SetDiscountPer10 is a paid mutator transaction binding the contract method 0x0f81e539.
//
// Solidity: function setDiscountPer10(uint32 value) returns()
func (_NodeManagerV83 *NodeManagerV83Transactor) SetDiscountPer10(opts *bind.TransactOpts, value uint32) (*types.Transaction, error) {
	return _NodeManagerV83.contract.Transact(opts, "setDiscountPer10", value)
}

// SetDiscountPer10 is a paid mutator transaction binding the contract method 0x0f81e539.
//
// Solidity: function setDiscountPer10(uint32 value) returns()
func (_NodeManagerV83 *NodeManagerV83Session) SetDiscountPer10(value uint32) (*types.Transaction, error) {
	return _NodeManagerV83.Contract.SetDiscountPer10(&_NodeManagerV83.TransactOpts, value)
}

// SetDiscountPer10 is a paid mutator transaction binding the contract method 0x0f81e539.
//
// Solidity: function setDiscountPer10(uint32 value) returns()
func (_NodeManagerV83 *NodeManagerV83TransactorSession) SetDiscountPer10(value uint32) (*types.Transaction, error) {
	return _NodeManagerV83.Contract.SetDiscountPer10(&_NodeManagerV83.TransactOpts, value)
}

// SetDynamicClaimFeeEnabled is a paid mutator transaction binding the contract method 0x4cdae712.
//
// Solidity: function setDynamicClaimFeeEnabled(bool val) returns()
func (_NodeManagerV83 *NodeManagerV83Transactor) SetDynamicClaimFeeEnabled(opts *bind.TransactOpts, val bool) (*types.Transaction, error) {
	return _NodeManagerV83.contract.Transact(opts, "setDynamicClaimFeeEnabled", val)
}

// SetDynamicClaimFeeEnabled is a paid mutator transaction binding the contract method 0x4cdae712.
//
// Solidity: function setDynamicClaimFeeEnabled(bool val) returns()
func (_NodeManagerV83 *NodeManagerV83Session) SetDynamicClaimFeeEnabled(val bool) (*types.Transaction, error) {
	return _NodeManagerV83.Contract.SetDynamicClaimFeeEnabled(&_NodeManagerV83.TransactOpts, val)
}

// SetDynamicClaimFeeEnabled is a paid mutator transaction binding the contract method 0x4cdae712.
//
// Solidity: function setDynamicClaimFeeEnabled(bool val) returns()
func (_NodeManagerV83 *NodeManagerV83TransactorSession) SetDynamicClaimFeeEnabled(val bool) (*types.Transaction, error) {
	return _NodeManagerV83.Contract.SetDynamicClaimFeeEnabled(&_NodeManagerV83.TransactOpts, val)
}

// SetMantPercent is a paid mutator transaction binding the contract method 0xc057e4f1.
//
// Solidity: function setMantPercent(uint32 value) returns()
func (_NodeManagerV83 *NodeManagerV83Transactor) SetMantPercent(opts *bind.TransactOpts, value uint32) (*types.Transaction, error) {
	return _NodeManagerV83.contract.Transact(opts, "setMantPercent", value)
}

// SetMantPercent is a paid mutator transaction binding the contract method 0xc057e4f1.
//
// Solidity: function setMantPercent(uint32 value) returns()
func (_NodeManagerV83 *NodeManagerV83Session) SetMantPercent(value uint32) (*types.Transaction, error) {
	return _NodeManagerV83.Contract.SetMantPercent(&_NodeManagerV83.TransactOpts, value)
}

// SetMantPercent is a paid mutator transaction binding the contract method 0xc057e4f1.
//
// Solidity: function setMantPercent(uint32 value) returns()
func (_NodeManagerV83 *NodeManagerV83TransactorSession) SetMantPercent(value uint32) (*types.Transaction, error) {
	return _NodeManagerV83.Contract.SetMantPercent(&_NodeManagerV83.TransactOpts, value)
}

// SetMaxCountOfUser is a paid mutator transaction binding the contract method 0xc63d63c0.
//
// Solidity: function setMaxCountOfUser(uint32 _count) returns()
func (_NodeManagerV83 *NodeManagerV83Transactor) SetMaxCountOfUser(opts *bind.TransactOpts, _count uint32) (*types.Transaction, error) {
	return _NodeManagerV83.contract.Transact(opts, "setMaxCountOfUser", _count)
}

// SetMaxCountOfUser is a paid mutator transaction binding the contract method 0xc63d63c0.
//
// Solidity: function setMaxCountOfUser(uint32 _count) returns()
func (_NodeManagerV83 *NodeManagerV83Session) SetMaxCountOfUser(_count uint32) (*types.Transaction, error) {
	return _NodeManagerV83.Contract.SetMaxCountOfUser(&_NodeManagerV83.TransactOpts, _count)
}

// SetMaxCountOfUser is a paid mutator transaction binding the contract method 0xc63d63c0.
//
// Solidity: function setMaxCountOfUser(uint32 _count) returns()
func (_NodeManagerV83 *NodeManagerV83TransactorSession) SetMaxCountOfUser(_count uint32) (*types.Transaction, error) {
	return _NodeManagerV83.Contract.SetMaxCountOfUser(&_NodeManagerV83.TransactOpts, _count)
}

// SetMaxMonthValue is a paid mutator transaction binding the contract method 0x77ffb52b.
//
// Solidity: function setMaxMonthValue(uint8 value) returns()
func (_NodeManagerV83 *NodeManagerV83Transactor) SetMaxMonthValue(opts *bind.TransactOpts, value uint8) (*types.Transaction, error) {
	return _NodeManagerV83.contract.Transact(opts, "setMaxMonthValue", value)
}

// SetMaxMonthValue is a paid mutator transaction binding the contract method 0x77ffb52b.
//
// Solidity: function setMaxMonthValue(uint8 value) returns()
func (_NodeManagerV83 *NodeManagerV83Session) SetMaxMonthValue(value uint8) (*types.Transaction, error) {
	return _NodeManagerV83.Contract.SetMaxMonthValue(&_NodeManagerV83.TransactOpts, value)
}

// SetMaxMonthValue is a paid mutator transaction binding the contract method 0x77ffb52b.
//
// Solidity: function setMaxMonthValue(uint8 value) returns()
func (_NodeManagerV83 *NodeManagerV83TransactorSession) SetMaxMonthValue(value uint8) (*types.Transaction, error) {
	return _NodeManagerV83.Contract.SetMaxMonthValue(&_NodeManagerV83.TransactOpts, value)
}

// SetNFTAddress is a paid mutator transaction binding the contract method 0x69d03738.
//
// Solidity: function setNFTAddress(address _nftAddress) returns()
func (_NodeManagerV83 *NodeManagerV83Transactor) SetNFTAddress(opts *bind.TransactOpts, _nftAddress common.Address) (*types.Transaction, error) {
	return _NodeManagerV83.contract.Transact(opts, "setNFTAddress", _nftAddress)
}

// SetNFTAddress is a paid mutator transaction binding the contract method 0x69d03738.
//
// Solidity: function setNFTAddress(address _nftAddress) returns()
func (_NodeManagerV83 *NodeManagerV83Session) SetNFTAddress(_nftAddress common.Address) (*types.Transaction, error) {
	return _NodeManagerV83.Contract.SetNFTAddress(&_NodeManagerV83.TransactOpts, _nftAddress)
}

// SetNFTAddress is a paid mutator transaction binding the contract method 0x69d03738.
//
// Solidity: function setNFTAddress(address _nftAddress) returns()
func (_NodeManagerV83 *NodeManagerV83TransactorSession) SetNFTAddress(_nftAddress common.Address) (*types.Transaction, error) {
	return _NodeManagerV83.Contract.SetNFTAddress(&_NodeManagerV83.TransactOpts, _nftAddress)
}

// SetOperationsPoolAddress is a paid mutator transaction binding the contract method 0xe985f283.
//
// Solidity: function setOperationsPoolAddress(address account) returns()
func (_NodeManagerV83 *NodeManagerV83Transactor) SetOperationsPoolAddress(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _NodeManagerV83.contract.Transact(opts, "setOperationsPoolAddress", account)
}

// SetOperationsPoolAddress is a paid mutator transaction binding the contract method 0xe985f283.
//
// Solidity: function setOperationsPoolAddress(address account) returns()
func (_NodeManagerV83 *NodeManagerV83Session) SetOperationsPoolAddress(account common.Address) (*types.Transaction, error) {
	return _NodeManagerV83.Contract.SetOperationsPoolAddress(&_NodeManagerV83.TransactOpts, account)
}

// SetOperationsPoolAddress is a paid mutator transaction binding the contract method 0xe985f283.
//
// Solidity: function setOperationsPoolAddress(address account) returns()
func (_NodeManagerV83 *NodeManagerV83TransactorSession) SetOperationsPoolAddress(account common.Address) (*types.Transaction, error) {
	return _NodeManagerV83.Contract.SetOperationsPoolAddress(&_NodeManagerV83.TransactOpts, account)
}

// SetOperationsPoolFee is a paid mutator transaction binding the contract method 0xf5278a8b.
//
// Solidity: function setOperationsPoolFee(uint32 value) returns()
func (_NodeManagerV83 *NodeManagerV83Transactor) SetOperationsPoolFee(opts *bind.TransactOpts, value uint32) (*types.Transaction, error) {
	return _NodeManagerV83.contract.Transact(opts, "setOperationsPoolFee", value)
}

// SetOperationsPoolFee is a paid mutator transaction binding the contract method 0xf5278a8b.
//
// Solidity: function setOperationsPoolFee(uint32 value) returns()
func (_NodeManagerV83 *NodeManagerV83Session) SetOperationsPoolFee(value uint32) (*types.Transaction, error) {
	return _NodeManagerV83.Contract.SetOperationsPoolFee(&_NodeManagerV83.TransactOpts, value)
}

// SetOperationsPoolFee is a paid mutator transaction binding the contract method 0xf5278a8b.
//
// Solidity: function setOperationsPoolFee(uint32 value) returns()
func (_NodeManagerV83 *NodeManagerV83TransactorSession) SetOperationsPoolFee(value uint32) (*types.Transaction, error) {
	return _NodeManagerV83.Contract.SetOperationsPoolFee(&_NodeManagerV83.TransactOpts, value)
}

// SetPayInterval is a paid mutator transaction binding the contract method 0x8c04a2e4.
//
// Solidity: function setPayInterval(uint32 value) returns()
func (_NodeManagerV83 *NodeManagerV83Transactor) SetPayInterval(opts *bind.TransactOpts, value uint32) (*types.Transaction, error) {
	return _NodeManagerV83.contract.Transact(opts, "setPayInterval", value)
}

// SetPayInterval is a paid mutator transaction binding the contract method 0x8c04a2e4.
//
// Solidity: function setPayInterval(uint32 value) returns()
func (_NodeManagerV83 *NodeManagerV83Session) SetPayInterval(value uint32) (*types.Transaction, error) {
	return _NodeManagerV83.Contract.SetPayInterval(&_NodeManagerV83.TransactOpts, value)
}

// SetPayInterval is a paid mutator transaction binding the contract method 0x8c04a2e4.
//
// Solidity: function setPayInterval(uint32 value) returns()
func (_NodeManagerV83 *NodeManagerV83TransactorSession) SetPayInterval(value uint32) (*types.Transaction, error) {
	return _NodeManagerV83.Contract.SetPayInterval(&_NodeManagerV83.TransactOpts, value)
}

// SetRewardsPoolAddress is a paid mutator transaction binding the contract method 0x9746f9e8.
//
// Solidity: function setRewardsPoolAddress(address account) returns()
func (_NodeManagerV83 *NodeManagerV83Transactor) SetRewardsPoolAddress(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _NodeManagerV83.contract.Transact(opts, "setRewardsPoolAddress", account)
}

// SetRewardsPoolAddress is a paid mutator transaction binding the contract method 0x9746f9e8.
//
// Solidity: function setRewardsPoolAddress(address account) returns()
func (_NodeManagerV83 *NodeManagerV83Session) SetRewardsPoolAddress(account common.Address) (*types.Transaction, error) {
	return _NodeManagerV83.Contract.SetRewardsPoolAddress(&_NodeManagerV83.TransactOpts, account)
}

// SetRewardsPoolAddress is a paid mutator transaction binding the contract method 0x9746f9e8.
//
// Solidity: function setRewardsPoolAddress(address account) returns()
func (_NodeManagerV83 *NodeManagerV83TransactorSession) SetRewardsPoolAddress(account common.Address) (*types.Transaction, error) {
	return _NodeManagerV83.Contract.SetRewardsPoolAddress(&_NodeManagerV83.TransactOpts, account)
}

// SetRewardsPoolFee is a paid mutator transaction binding the contract method 0x3fb53751.
//
// Solidity: function setRewardsPoolFee(uint32 value) returns()
func (_NodeManagerV83 *NodeManagerV83Transactor) SetRewardsPoolFee(opts *bind.TransactOpts, value uint32) (*types.Transaction, error) {
	return _NodeManagerV83.contract.Transact(opts, "setRewardsPoolFee", value)
}

// SetRewardsPoolFee is a paid mutator transaction binding the contract method 0x3fb53751.
//
// Solidity: function setRewardsPoolFee(uint32 value) returns()
func (_NodeManagerV83 *NodeManagerV83Session) SetRewardsPoolFee(value uint32) (*types.Transaction, error) {
	return _NodeManagerV83.Contract.SetRewardsPoolFee(&_NodeManagerV83.TransactOpts, value)
}

// SetRewardsPoolFee is a paid mutator transaction binding the contract method 0x3fb53751.
//
// Solidity: function setRewardsPoolFee(uint32 value) returns()
func (_NodeManagerV83 *NodeManagerV83TransactorSession) SetRewardsPoolFee(value uint32) (*types.Transaction, error) {
	return _NodeManagerV83.Contract.SetRewardsPoolFee(&_NodeManagerV83.TransactOpts, value)
}

// SetRouter is a paid mutator transaction binding the contract method 0xc0d78655.
//
// Solidity: function setRouter(address router) returns()
func (_NodeManagerV83 *NodeManagerV83Transactor) SetRouter(opts *bind.TransactOpts, router common.Address) (*types.Transaction, error) {
	return _NodeManagerV83.contract.Transact(opts, "setRouter", router)
}

// SetRouter is a paid mutator transaction binding the contract method 0xc0d78655.
//
// Solidity: function setRouter(address router) returns()
func (_NodeManagerV83 *NodeManagerV83Session) SetRouter(router common.Address) (*types.Transaction, error) {
	return _NodeManagerV83.Contract.SetRouter(&_NodeManagerV83.TransactOpts, router)
}

// SetRouter is a paid mutator transaction binding the contract method 0xc0d78655.
//
// Solidity: function setRouter(address router) returns()
func (_NodeManagerV83 *NodeManagerV83TransactorSession) SetRouter(router common.Address) (*types.Transaction, error) {
	return _NodeManagerV83.Contract.SetRouter(&_NodeManagerV83.TransactOpts, router)
}

// SetTokenAddress is a paid mutator transaction binding the contract method 0x26a4e8d2.
//
// Solidity: function setTokenAddress(address token) returns()
func (_NodeManagerV83 *NodeManagerV83Transactor) SetTokenAddress(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _NodeManagerV83.contract.Transact(opts, "setTokenAddress", token)
}

// SetTokenAddress is a paid mutator transaction binding the contract method 0x26a4e8d2.
//
// Solidity: function setTokenAddress(address token) returns()
func (_NodeManagerV83 *NodeManagerV83Session) SetTokenAddress(token common.Address) (*types.Transaction, error) {
	return _NodeManagerV83.Contract.SetTokenAddress(&_NodeManagerV83.TransactOpts, token)
}

// SetTokenAddress is a paid mutator transaction binding the contract method 0x26a4e8d2.
//
// Solidity: function setTokenAddress(address token) returns()
func (_NodeManagerV83 *NodeManagerV83TransactorSession) SetTokenAddress(token common.Address) (*types.Transaction, error) {
	return _NodeManagerV83.Contract.SetTokenAddress(&_NodeManagerV83.TransactOpts, token)
}

// SetTransferFee is a paid mutator transaction binding the contract method 0x2e9ef976.
//
// Solidity: function setTransferFee(uint32 value) returns()
func (_NodeManagerV83 *NodeManagerV83Transactor) SetTransferFee(opts *bind.TransactOpts, value uint32) (*types.Transaction, error) {
	return _NodeManagerV83.contract.Transact(opts, "setTransferFee", value)
}

// SetTransferFee is a paid mutator transaction binding the contract method 0x2e9ef976.
//
// Solidity: function setTransferFee(uint32 value) returns()
func (_NodeManagerV83 *NodeManagerV83Session) SetTransferFee(value uint32) (*types.Transaction, error) {
	return _NodeManagerV83.Contract.SetTransferFee(&_NodeManagerV83.TransactOpts, value)
}

// SetTransferFee is a paid mutator transaction binding the contract method 0x2e9ef976.
//
// Solidity: function setTransferFee(uint32 value) returns()
func (_NodeManagerV83 *NodeManagerV83TransactorSession) SetTransferFee(value uint32) (*types.Transaction, error) {
	return _NodeManagerV83.Contract.SetTransferFee(&_NodeManagerV83.TransactOpts, value)
}

// SetsellPricePercent is a paid mutator transaction binding the contract method 0xe0d12873.
//
// Solidity: function setsellPricePercent(uint32 value) returns()
func (_NodeManagerV83 *NodeManagerV83Transactor) SetsellPricePercent(opts *bind.TransactOpts, value uint32) (*types.Transaction, error) {
	return _NodeManagerV83.contract.Transact(opts, "setsellPricePercent", value)
}

// SetsellPricePercent is a paid mutator transaction binding the contract method 0xe0d12873.
//
// Solidity: function setsellPricePercent(uint32 value) returns()
func (_NodeManagerV83 *NodeManagerV83Session) SetsellPricePercent(value uint32) (*types.Transaction, error) {
	return _NodeManagerV83.Contract.SetsellPricePercent(&_NodeManagerV83.TransactOpts, value)
}

// SetsellPricePercent is a paid mutator transaction binding the contract method 0xe0d12873.
//
// Solidity: function setsellPricePercent(uint32 value) returns()
func (_NodeManagerV83 *NodeManagerV83TransactorSession) SetsellPricePercent(value uint32) (*types.Transaction, error) {
	return _NodeManagerV83.Contract.SetsellPricePercent(&_NodeManagerV83.TransactOpts, value)
}

// Transfer is a paid mutator transaction binding the contract method 0x563e497c.
//
// Solidity: function transfer(string tierName, uint32 count, address from, address to) returns()
func (_NodeManagerV83 *NodeManagerV83Transactor) Transfer(opts *bind.TransactOpts, tierName string, count uint32, from common.Address, to common.Address) (*types.Transaction, error) {
	return _NodeManagerV83.contract.Transact(opts, "transfer", tierName, count, from, to)
}

// Transfer is a paid mutator transaction binding the contract method 0x563e497c.
//
// Solidity: function transfer(string tierName, uint32 count, address from, address to) returns()
func (_NodeManagerV83 *NodeManagerV83Session) Transfer(tierName string, count uint32, from common.Address, to common.Address) (*types.Transaction, error) {
	return _NodeManagerV83.Contract.Transfer(&_NodeManagerV83.TransactOpts, tierName, count, from, to)
}

// Transfer is a paid mutator transaction binding the contract method 0x563e497c.
//
// Solidity: function transfer(string tierName, uint32 count, address from, address to) returns()
func (_NodeManagerV83 *NodeManagerV83TransactorSession) Transfer(tierName string, count uint32, from common.Address, to common.Address) (*types.Transaction, error) {
	return _NodeManagerV83.Contract.Transfer(&_NodeManagerV83.TransactOpts, tierName, count, from, to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NodeManagerV83 *NodeManagerV83Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _NodeManagerV83.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NodeManagerV83 *NodeManagerV83Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NodeManagerV83.Contract.TransferOwnership(&_NodeManagerV83.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NodeManagerV83 *NodeManagerV83TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NodeManagerV83.Contract.TransferOwnership(&_NodeManagerV83.TransactOpts, newOwner)
}

// UpdateTier is a paid mutator transaction binding the contract method 0xdbe2b8ce.
//
// Solidity: function updateTier(string tierName, string name, uint256 price, uint256 rewardsPerTime, uint32 claimInterval, uint256 maintenanceFee) returns()
func (_NodeManagerV83 *NodeManagerV83Transactor) UpdateTier(opts *bind.TransactOpts, tierName string, name string, price *big.Int, rewardsPerTime *big.Int, claimInterval uint32, maintenanceFee *big.Int) (*types.Transaction, error) {
	return _NodeManagerV83.contract.Transact(opts, "updateTier", tierName, name, price, rewardsPerTime, claimInterval, maintenanceFee)
}

// UpdateTier is a paid mutator transaction binding the contract method 0xdbe2b8ce.
//
// Solidity: function updateTier(string tierName, string name, uint256 price, uint256 rewardsPerTime, uint32 claimInterval, uint256 maintenanceFee) returns()
func (_NodeManagerV83 *NodeManagerV83Session) UpdateTier(tierName string, name string, price *big.Int, rewardsPerTime *big.Int, claimInterval uint32, maintenanceFee *big.Int) (*types.Transaction, error) {
	return _NodeManagerV83.Contract.UpdateTier(&_NodeManagerV83.TransactOpts, tierName, name, price, rewardsPerTime, claimInterval, maintenanceFee)
}

// UpdateTier is a paid mutator transaction binding the contract method 0xdbe2b8ce.
//
// Solidity: function updateTier(string tierName, string name, uint256 price, uint256 rewardsPerTime, uint32 claimInterval, uint256 maintenanceFee) returns()
func (_NodeManagerV83 *NodeManagerV83TransactorSession) UpdateTier(tierName string, name string, price *big.Int, rewardsPerTime *big.Int, claimInterval uint32, maintenanceFee *big.Int) (*types.Transaction, error) {
	return _NodeManagerV83.Contract.UpdateTier(&_NodeManagerV83.TransactOpts, tierName, name, price, rewardsPerTime, claimInterval, maintenanceFee)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf940e385.
//
// Solidity: function withdraw(address anyToken, address recipient) returns()
func (_NodeManagerV83 *NodeManagerV83Transactor) Withdraw(opts *bind.TransactOpts, anyToken common.Address, recipient common.Address) (*types.Transaction, error) {
	return _NodeManagerV83.contract.Transact(opts, "withdraw", anyToken, recipient)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf940e385.
//
// Solidity: function withdraw(address anyToken, address recipient) returns()
func (_NodeManagerV83 *NodeManagerV83Session) Withdraw(anyToken common.Address, recipient common.Address) (*types.Transaction, error) {
	return _NodeManagerV83.Contract.Withdraw(&_NodeManagerV83.TransactOpts, anyToken, recipient)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf940e385.
//
// Solidity: function withdraw(address anyToken, address recipient) returns()
func (_NodeManagerV83 *NodeManagerV83TransactorSession) Withdraw(anyToken common.Address, recipient common.Address) (*types.Transaction, error) {
	return _NodeManagerV83.Contract.Withdraw(&_NodeManagerV83.TransactOpts, anyToken, recipient)
}

// NodeManagerV83GiftCardPayedIterator is returned from FilterGiftCardPayed and is used to iterate over the raw logs and unpacked data for GiftCardPayed events raised by the NodeManagerV83 contract.
type NodeManagerV83GiftCardPayedIterator struct {
	Event *NodeManagerV83GiftCardPayed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NodeManagerV83GiftCardPayedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeManagerV83GiftCardPayed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NodeManagerV83GiftCardPayed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NodeManagerV83GiftCardPayedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeManagerV83GiftCardPayedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeManagerV83GiftCardPayed represents a GiftCardPayed event raised by the NodeManagerV83 contract.
type NodeManagerV83GiftCardPayed struct {
	Arg0 common.Address
	Arg1 common.Address
	Arg2 string
	Arg3 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterGiftCardPayed is a free log retrieval operation binding the contract event 0x34618306376b53644b3c2e2e35fc793e80e74a56c2c29194492397a49ef02277.
//
// Solidity: event GiftCardPayed(address arg0, address arg1, string arg2, uint256 arg3)
func (_NodeManagerV83 *NodeManagerV83Filterer) FilterGiftCardPayed(opts *bind.FilterOpts) (*NodeManagerV83GiftCardPayedIterator, error) {

	logs, sub, err := _NodeManagerV83.contract.FilterLogs(opts, "GiftCardPayed")
	if err != nil {
		return nil, err
	}
	return &NodeManagerV83GiftCardPayedIterator{contract: _NodeManagerV83.contract, event: "GiftCardPayed", logs: logs, sub: sub}, nil
}

// WatchGiftCardPayed is a free log subscription operation binding the contract event 0x34618306376b53644b3c2e2e35fc793e80e74a56c2c29194492397a49ef02277.
//
// Solidity: event GiftCardPayed(address arg0, address arg1, string arg2, uint256 arg3)
func (_NodeManagerV83 *NodeManagerV83Filterer) WatchGiftCardPayed(opts *bind.WatchOpts, sink chan<- *NodeManagerV83GiftCardPayed) (event.Subscription, error) {

	logs, sub, err := _NodeManagerV83.contract.WatchLogs(opts, "GiftCardPayed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeManagerV83GiftCardPayed)
				if err := _NodeManagerV83.contract.UnpackLog(event, "GiftCardPayed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGiftCardPayed is a log parse operation binding the contract event 0x34618306376b53644b3c2e2e35fc793e80e74a56c2c29194492397a49ef02277.
//
// Solidity: event GiftCardPayed(address arg0, address arg1, string arg2, uint256 arg3)
func (_NodeManagerV83 *NodeManagerV83Filterer) ParseGiftCardPayed(log types.Log) (*NodeManagerV83GiftCardPayed, error) {
	event := new(NodeManagerV83GiftCardPayed)
	if err := _NodeManagerV83.contract.UnpackLog(event, "GiftCardPayed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeManagerV83InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the NodeManagerV83 contract.
type NodeManagerV83InitializedIterator struct {
	Event *NodeManagerV83Initialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NodeManagerV83InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeManagerV83Initialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NodeManagerV83Initialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NodeManagerV83InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeManagerV83InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeManagerV83Initialized represents a Initialized event raised by the NodeManagerV83 contract.
type NodeManagerV83Initialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_NodeManagerV83 *NodeManagerV83Filterer) FilterInitialized(opts *bind.FilterOpts) (*NodeManagerV83InitializedIterator, error) {

	logs, sub, err := _NodeManagerV83.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &NodeManagerV83InitializedIterator{contract: _NodeManagerV83.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_NodeManagerV83 *NodeManagerV83Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *NodeManagerV83Initialized) (event.Subscription, error) {

	logs, sub, err := _NodeManagerV83.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeManagerV83Initialized)
				if err := _NodeManagerV83.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_NodeManagerV83 *NodeManagerV83Filterer) ParseInitialized(log types.Log) (*NodeManagerV83Initialized, error) {
	event := new(NodeManagerV83Initialized)
	if err := _NodeManagerV83.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeManagerV83NodeCreatedIterator is returned from FilterNodeCreated and is used to iterate over the raw logs and unpacked data for NodeCreated events raised by the NodeManagerV83 contract.
type NodeManagerV83NodeCreatedIterator struct {
	Event *NodeManagerV83NodeCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NodeManagerV83NodeCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeManagerV83NodeCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NodeManagerV83NodeCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NodeManagerV83NodeCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeManagerV83NodeCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeManagerV83NodeCreated represents a NodeCreated event raised by the NodeManagerV83 contract.
type NodeManagerV83NodeCreated struct {
	Arg0 common.Address
	Arg1 string
	Arg2 uint32
	Arg3 uint32
	Arg4 uint32
	Arg5 uint32
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNodeCreated is a free log retrieval operation binding the contract event 0x3a823590161fd78e3988b3da59eb891df9f43272cf652f586d2e79f65f6f3571.
//
// Solidity: event NodeCreated(address arg0, string arg1, uint32 arg2, uint32 arg3, uint32 arg4, uint32 arg5)
func (_NodeManagerV83 *NodeManagerV83Filterer) FilterNodeCreated(opts *bind.FilterOpts) (*NodeManagerV83NodeCreatedIterator, error) {

	logs, sub, err := _NodeManagerV83.contract.FilterLogs(opts, "NodeCreated")
	if err != nil {
		return nil, err
	}
	return &NodeManagerV83NodeCreatedIterator{contract: _NodeManagerV83.contract, event: "NodeCreated", logs: logs, sub: sub}, nil
}

// WatchNodeCreated is a free log subscription operation binding the contract event 0x3a823590161fd78e3988b3da59eb891df9f43272cf652f586d2e79f65f6f3571.
//
// Solidity: event NodeCreated(address arg0, string arg1, uint32 arg2, uint32 arg3, uint32 arg4, uint32 arg5)
func (_NodeManagerV83 *NodeManagerV83Filterer) WatchNodeCreated(opts *bind.WatchOpts, sink chan<- *NodeManagerV83NodeCreated) (event.Subscription, error) {

	logs, sub, err := _NodeManagerV83.contract.WatchLogs(opts, "NodeCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeManagerV83NodeCreated)
				if err := _NodeManagerV83.contract.UnpackLog(event, "NodeCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNodeCreated is a log parse operation binding the contract event 0x3a823590161fd78e3988b3da59eb891df9f43272cf652f586d2e79f65f6f3571.
//
// Solidity: event NodeCreated(address arg0, string arg1, uint32 arg2, uint32 arg3, uint32 arg4, uint32 arg5)
func (_NodeManagerV83 *NodeManagerV83Filterer) ParseNodeCreated(log types.Log) (*NodeManagerV83NodeCreated, error) {
	event := new(NodeManagerV83NodeCreated)
	if err := _NodeManagerV83.contract.UnpackLog(event, "NodeCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeManagerV83NodeTransferedIterator is returned from FilterNodeTransfered and is used to iterate over the raw logs and unpacked data for NodeTransfered events raised by the NodeManagerV83 contract.
type NodeManagerV83NodeTransferedIterator struct {
	Event *NodeManagerV83NodeTransfered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NodeManagerV83NodeTransferedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeManagerV83NodeTransfered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NodeManagerV83NodeTransfered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NodeManagerV83NodeTransferedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeManagerV83NodeTransferedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeManagerV83NodeTransfered represents a NodeTransfered event raised by the NodeManagerV83 contract.
type NodeManagerV83NodeTransfered struct {
	Arg0 common.Address
	Arg1 common.Address
	Arg2 uint32
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNodeTransfered is a free log retrieval operation binding the contract event 0x8a091ff4fd9357c3a846a822cac198fd60c380e70afd0af7e2451a3953727e03.
//
// Solidity: event NodeTransfered(address arg0, address arg1, uint32 arg2)
func (_NodeManagerV83 *NodeManagerV83Filterer) FilterNodeTransfered(opts *bind.FilterOpts) (*NodeManagerV83NodeTransferedIterator, error) {

	logs, sub, err := _NodeManagerV83.contract.FilterLogs(opts, "NodeTransfered")
	if err != nil {
		return nil, err
	}
	return &NodeManagerV83NodeTransferedIterator{contract: _NodeManagerV83.contract, event: "NodeTransfered", logs: logs, sub: sub}, nil
}

// WatchNodeTransfered is a free log subscription operation binding the contract event 0x8a091ff4fd9357c3a846a822cac198fd60c380e70afd0af7e2451a3953727e03.
//
// Solidity: event NodeTransfered(address arg0, address arg1, uint32 arg2)
func (_NodeManagerV83 *NodeManagerV83Filterer) WatchNodeTransfered(opts *bind.WatchOpts, sink chan<- *NodeManagerV83NodeTransfered) (event.Subscription, error) {

	logs, sub, err := _NodeManagerV83.contract.WatchLogs(opts, "NodeTransfered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeManagerV83NodeTransfered)
				if err := _NodeManagerV83.contract.UnpackLog(event, "NodeTransfered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNodeTransfered is a log parse operation binding the contract event 0x8a091ff4fd9357c3a846a822cac198fd60c380e70afd0af7e2451a3953727e03.
//
// Solidity: event NodeTransfered(address arg0, address arg1, uint32 arg2)
func (_NodeManagerV83 *NodeManagerV83Filterer) ParseNodeTransfered(log types.Log) (*NodeManagerV83NodeTransfered, error) {
	event := new(NodeManagerV83NodeTransfered)
	if err := _NodeManagerV83.contract.UnpackLog(event, "NodeTransfered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeManagerV83NodeUpdatedIterator is returned from FilterNodeUpdated and is used to iterate over the raw logs and unpacked data for NodeUpdated events raised by the NodeManagerV83 contract.
type NodeManagerV83NodeUpdatedIterator struct {
	Event *NodeManagerV83NodeUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NodeManagerV83NodeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeManagerV83NodeUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NodeManagerV83NodeUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NodeManagerV83NodeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeManagerV83NodeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeManagerV83NodeUpdated represents a NodeUpdated event raised by the NodeManagerV83 contract.
type NodeManagerV83NodeUpdated struct {
	Arg0 common.Address
	Arg1 string
	Arg2 string
	Arg3 uint32
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNodeUpdated is a free log retrieval operation binding the contract event 0x003fb356fcd327584ecb5dba27616f7e2085b7b87a5750df071877c0e6b5a840.
//
// Solidity: event NodeUpdated(address arg0, string arg1, string arg2, uint32 arg3)
func (_NodeManagerV83 *NodeManagerV83Filterer) FilterNodeUpdated(opts *bind.FilterOpts) (*NodeManagerV83NodeUpdatedIterator, error) {

	logs, sub, err := _NodeManagerV83.contract.FilterLogs(opts, "NodeUpdated")
	if err != nil {
		return nil, err
	}
	return &NodeManagerV83NodeUpdatedIterator{contract: _NodeManagerV83.contract, event: "NodeUpdated", logs: logs, sub: sub}, nil
}

// WatchNodeUpdated is a free log subscription operation binding the contract event 0x003fb356fcd327584ecb5dba27616f7e2085b7b87a5750df071877c0e6b5a840.
//
// Solidity: event NodeUpdated(address arg0, string arg1, string arg2, uint32 arg3)
func (_NodeManagerV83 *NodeManagerV83Filterer) WatchNodeUpdated(opts *bind.WatchOpts, sink chan<- *NodeManagerV83NodeUpdated) (event.Subscription, error) {

	logs, sub, err := _NodeManagerV83.contract.WatchLogs(opts, "NodeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeManagerV83NodeUpdated)
				if err := _NodeManagerV83.contract.UnpackLog(event, "NodeUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNodeUpdated is a log parse operation binding the contract event 0x003fb356fcd327584ecb5dba27616f7e2085b7b87a5750df071877c0e6b5a840.
//
// Solidity: event NodeUpdated(address arg0, string arg1, string arg2, uint32 arg3)
func (_NodeManagerV83 *NodeManagerV83Filterer) ParseNodeUpdated(log types.Log) (*NodeManagerV83NodeUpdated, error) {
	event := new(NodeManagerV83NodeUpdated)
	if err := _NodeManagerV83.contract.UnpackLog(event, "NodeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
