// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package esd

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// EsdABI is the input ABI used to generate the binding from.
const EsdABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Advance\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"valueUnderlying\",\"type\":\"uint256\"}],\"name\":\"Bond\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"Commit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"CouponApproval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"couponsExpired\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lessRedeemable\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lessDebt\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBonded\",\"type\":\"uint256\"}],\"name\":\"CouponExpiration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dollarAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"couponAmount\",\"type\":\"uint256\"}],\"name\":\"CouponPurchase\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"couponAmount\",\"type\":\"uint256\"}],\"name\":\"CouponRedemption\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"CouponTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Incentivization\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"candidate\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"Proposal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDebt\",\"type\":\"uint256\"}],\"name\":\"SupplyDecrease\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newRedeemable\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lessDebt\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBonded\",\"type\":\"uint256\"}],\"name\":\"SupplyIncrease\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"SupplyNeutral\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"valueUnderlying\",\"type\":\"uint256\"}],\"name\":\"Unbond\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"candidate\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumCandidate.Vote\",\"name\":\"vote\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bonded\",\"type\":\"uint256\"}],\"name\":\"Vote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowanceCoupons\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approveCoupons\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"approveFor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOfBonded\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"balanceOfCoupons\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOfStaged\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"bond\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"bootstrappingAt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"commit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"couponPremium\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"couponsExpiration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"dollar\",\"outputs\":[{\"internalType\":\"contractIDollar\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"emergencyCommit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"epoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"epochTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"expiringCoupons\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"expiringCouponsAtIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"impl\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"isInitialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"isNominated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"contractIOracle\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"outstandingCoupons\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"periodFor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dollarAmount\",\"type\":\"uint256\"}],\"name\":\"purchaseCoupons\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"recordedVote\",\"outputs\":[{\"internalType\":\"enumCandidate.Vote\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"couponEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"couponAmount\",\"type\":\"uint256\"}],\"name\":\"redeemCoupons\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"rejectFor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"startFor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"statusOf\",\"outputs\":[{\"internalType\":\"enumAccount.Status\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalBonded\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"totalBondedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalCoupons\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalDebt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalNet\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalRedeemable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalStaged\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferCoupons\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"unbond\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"unbondUnderlying\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"candidate\",\"type\":\"address\"},{\"internalType\":\"enumCandidate.Vote\",\"name\":\"vote\",\"type\":\"uint8\"}],\"name\":\"vote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"votesFor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"advance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Esd is an auto generated Go binding around an Ethereum contract.
type Esd struct {
	EsdCaller     // Read-only binding to the contract
	EsdTransactor // Write-only binding to the contract
	EsdFilterer   // Log filterer for contract events
}

// EsdCaller is an auto generated read-only Go binding around an Ethereum contract.
type EsdCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EsdTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EsdTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EsdFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EsdFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EsdSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EsdSession struct {
	Contract     *Esd              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EsdCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EsdCallerSession struct {
	Contract *EsdCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// EsdTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EsdTransactorSession struct {
	Contract     *EsdTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EsdRaw is an auto generated low-level Go binding around an Ethereum contract.
type EsdRaw struct {
	Contract *Esd // Generic contract binding to access the raw methods on
}

// EsdCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EsdCallerRaw struct {
	Contract *EsdCaller // Generic read-only contract binding to access the raw methods on
}

// EsdTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EsdTransactorRaw struct {
	Contract *EsdTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEsd creates a new instance of Esd, bound to a specific deployed contract.
func NewEsd(address common.Address, backend bind.ContractBackend) (*Esd, error) {
	contract, err := bindEsd(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Esd{EsdCaller: EsdCaller{contract: contract}, EsdTransactor: EsdTransactor{contract: contract}, EsdFilterer: EsdFilterer{contract: contract}}, nil
}

// NewEsdCaller creates a new read-only instance of Esd, bound to a specific deployed contract.
func NewEsdCaller(address common.Address, caller bind.ContractCaller) (*EsdCaller, error) {
	contract, err := bindEsd(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EsdCaller{contract: contract}, nil
}

// NewEsdTransactor creates a new write-only instance of Esd, bound to a specific deployed contract.
func NewEsdTransactor(address common.Address, transactor bind.ContractTransactor) (*EsdTransactor, error) {
	contract, err := bindEsd(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EsdTransactor{contract: contract}, nil
}

// NewEsdFilterer creates a new log filterer instance of Esd, bound to a specific deployed contract.
func NewEsdFilterer(address common.Address, filterer bind.ContractFilterer) (*EsdFilterer, error) {
	contract, err := bindEsd(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EsdFilterer{contract: contract}, nil
}

// bindEsd binds a generic wrapper to an already deployed contract.
func bindEsd(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EsdABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Esd *EsdRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Esd.Contract.EsdCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Esd *EsdRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Esd.Contract.EsdTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Esd *EsdRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Esd.Contract.EsdTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Esd *EsdCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Esd.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Esd *EsdTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Esd.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Esd *EsdTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Esd.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Esd *EsdCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Esd.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Esd *EsdSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Esd.Contract.Allowance(&_Esd.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Esd *EsdCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Esd.Contract.Allowance(&_Esd.CallOpts, owner, spender)
}

// AllowanceCoupons is a free data retrieval call binding the contract method 0x9f6e1b26.
//
// Solidity: function allowanceCoupons(address owner, address spender) view returns(uint256)
func (_Esd *EsdCaller) AllowanceCoupons(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Esd.contract.Call(opts, &out, "allowanceCoupons", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllowanceCoupons is a free data retrieval call binding the contract method 0x9f6e1b26.
//
// Solidity: function allowanceCoupons(address owner, address spender) view returns(uint256)
func (_Esd *EsdSession) AllowanceCoupons(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Esd.Contract.AllowanceCoupons(&_Esd.CallOpts, owner, spender)
}

// AllowanceCoupons is a free data retrieval call binding the contract method 0x9f6e1b26.
//
// Solidity: function allowanceCoupons(address owner, address spender) view returns(uint256)
func (_Esd *EsdCallerSession) AllowanceCoupons(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Esd.Contract.AllowanceCoupons(&_Esd.CallOpts, owner, spender)
}

// ApproveFor is a free data retrieval call binding the contract method 0x64668022.
//
// Solidity: function approveFor(address candidate) view returns(uint256)
func (_Esd *EsdCaller) ApproveFor(opts *bind.CallOpts, candidate common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Esd.contract.Call(opts, &out, "approveFor", candidate)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ApproveFor is a free data retrieval call binding the contract method 0x64668022.
//
// Solidity: function approveFor(address candidate) view returns(uint256)
func (_Esd *EsdSession) ApproveFor(candidate common.Address) (*big.Int, error) {
	return _Esd.Contract.ApproveFor(&_Esd.CallOpts, candidate)
}

// ApproveFor is a free data retrieval call binding the contract method 0x64668022.
//
// Solidity: function approveFor(address candidate) view returns(uint256)
func (_Esd *EsdCallerSession) ApproveFor(candidate common.Address) (*big.Int, error) {
	return _Esd.Contract.ApproveFor(&_Esd.CallOpts, candidate)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Esd *EsdCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Esd.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Esd *EsdSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Esd.Contract.BalanceOf(&_Esd.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Esd *EsdCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Esd.Contract.BalanceOf(&_Esd.CallOpts, account)
}

// BalanceOfBonded is a free data retrieval call binding the contract method 0x825ad607.
//
// Solidity: function balanceOfBonded(address account) view returns(uint256)
func (_Esd *EsdCaller) BalanceOfBonded(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Esd.contract.Call(opts, &out, "balanceOfBonded", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOfBonded is a free data retrieval call binding the contract method 0x825ad607.
//
// Solidity: function balanceOfBonded(address account) view returns(uint256)
func (_Esd *EsdSession) BalanceOfBonded(account common.Address) (*big.Int, error) {
	return _Esd.Contract.BalanceOfBonded(&_Esd.CallOpts, account)
}

// BalanceOfBonded is a free data retrieval call binding the contract method 0x825ad607.
//
// Solidity: function balanceOfBonded(address account) view returns(uint256)
func (_Esd *EsdCallerSession) BalanceOfBonded(account common.Address) (*big.Int, error) {
	return _Esd.Contract.BalanceOfBonded(&_Esd.CallOpts, account)
}

// BalanceOfCoupons is a free data retrieval call binding the contract method 0xbc7513e2.
//
// Solidity: function balanceOfCoupons(address account, uint256 epoch) view returns(uint256)
func (_Esd *EsdCaller) BalanceOfCoupons(opts *bind.CallOpts, account common.Address, epoch *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Esd.contract.Call(opts, &out, "balanceOfCoupons", account, epoch)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOfCoupons is a free data retrieval call binding the contract method 0xbc7513e2.
//
// Solidity: function balanceOfCoupons(address account, uint256 epoch) view returns(uint256)
func (_Esd *EsdSession) BalanceOfCoupons(account common.Address, epoch *big.Int) (*big.Int, error) {
	return _Esd.Contract.BalanceOfCoupons(&_Esd.CallOpts, account, epoch)
}

// BalanceOfCoupons is a free data retrieval call binding the contract method 0xbc7513e2.
//
// Solidity: function balanceOfCoupons(address account, uint256 epoch) view returns(uint256)
func (_Esd *EsdCallerSession) BalanceOfCoupons(account common.Address, epoch *big.Int) (*big.Int, error) {
	return _Esd.Contract.BalanceOfCoupons(&_Esd.CallOpts, account, epoch)
}

// BalanceOfStaged is a free data retrieval call binding the contract method 0x86cf9f14.
//
// Solidity: function balanceOfStaged(address account) view returns(uint256)
func (_Esd *EsdCaller) BalanceOfStaged(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Esd.contract.Call(opts, &out, "balanceOfStaged", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOfStaged is a free data retrieval call binding the contract method 0x86cf9f14.
//
// Solidity: function balanceOfStaged(address account) view returns(uint256)
func (_Esd *EsdSession) BalanceOfStaged(account common.Address) (*big.Int, error) {
	return _Esd.Contract.BalanceOfStaged(&_Esd.CallOpts, account)
}

// BalanceOfStaged is a free data retrieval call binding the contract method 0x86cf9f14.
//
// Solidity: function balanceOfStaged(address account) view returns(uint256)
func (_Esd *EsdCallerSession) BalanceOfStaged(account common.Address) (*big.Int, error) {
	return _Esd.Contract.BalanceOfStaged(&_Esd.CallOpts, account)
}

// BootstrappingAt is a free data retrieval call binding the contract method 0x75d5024b.
//
// Solidity: function bootstrappingAt(uint256 epoch) view returns(bool)
func (_Esd *EsdCaller) BootstrappingAt(opts *bind.CallOpts, epoch *big.Int) (bool, error) {
	var out []interface{}
	err := _Esd.contract.Call(opts, &out, "bootstrappingAt", epoch)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BootstrappingAt is a free data retrieval call binding the contract method 0x75d5024b.
//
// Solidity: function bootstrappingAt(uint256 epoch) view returns(bool)
func (_Esd *EsdSession) BootstrappingAt(epoch *big.Int) (bool, error) {
	return _Esd.Contract.BootstrappingAt(&_Esd.CallOpts, epoch)
}

// BootstrappingAt is a free data retrieval call binding the contract method 0x75d5024b.
//
// Solidity: function bootstrappingAt(uint256 epoch) view returns(bool)
func (_Esd *EsdCallerSession) BootstrappingAt(epoch *big.Int) (bool, error) {
	return _Esd.Contract.BootstrappingAt(&_Esd.CallOpts, epoch)
}

// CouponPremium is a free data retrieval call binding the contract method 0xd8f54138.
//
// Solidity: function couponPremium(uint256 amount) view returns(uint256)
func (_Esd *EsdCaller) CouponPremium(opts *bind.CallOpts, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Esd.contract.Call(opts, &out, "couponPremium", amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CouponPremium is a free data retrieval call binding the contract method 0xd8f54138.
//
// Solidity: function couponPremium(uint256 amount) view returns(uint256)
func (_Esd *EsdSession) CouponPremium(amount *big.Int) (*big.Int, error) {
	return _Esd.Contract.CouponPremium(&_Esd.CallOpts, amount)
}

// CouponPremium is a free data retrieval call binding the contract method 0xd8f54138.
//
// Solidity: function couponPremium(uint256 amount) view returns(uint256)
func (_Esd *EsdCallerSession) CouponPremium(amount *big.Int) (*big.Int, error) {
	return _Esd.Contract.CouponPremium(&_Esd.CallOpts, amount)
}

// CouponsExpiration is a free data retrieval call binding the contract method 0x10e95b6c.
//
// Solidity: function couponsExpiration(uint256 epoch) view returns(uint256)
func (_Esd *EsdCaller) CouponsExpiration(opts *bind.CallOpts, epoch *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Esd.contract.Call(opts, &out, "couponsExpiration", epoch)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CouponsExpiration is a free data retrieval call binding the contract method 0x10e95b6c.
//
// Solidity: function couponsExpiration(uint256 epoch) view returns(uint256)
func (_Esd *EsdSession) CouponsExpiration(epoch *big.Int) (*big.Int, error) {
	return _Esd.Contract.CouponsExpiration(&_Esd.CallOpts, epoch)
}

// CouponsExpiration is a free data retrieval call binding the contract method 0x10e95b6c.
//
// Solidity: function couponsExpiration(uint256 epoch) view returns(uint256)
func (_Esd *EsdCallerSession) CouponsExpiration(epoch *big.Int) (*big.Int, error) {
	return _Esd.Contract.CouponsExpiration(&_Esd.CallOpts, epoch)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Esd *EsdCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Esd.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Esd *EsdSession) Decimals() (uint8, error) {
	return _Esd.Contract.Decimals(&_Esd.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Esd *EsdCallerSession) Decimals() (uint8, error) {
	return _Esd.Contract.Decimals(&_Esd.CallOpts)
}

// Dollar is a free data retrieval call binding the contract method 0x51adeb57.
//
// Solidity: function dollar() view returns(address)
func (_Esd *EsdCaller) Dollar(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Esd.contract.Call(opts, &out, "dollar")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dollar is a free data retrieval call binding the contract method 0x51adeb57.
//
// Solidity: function dollar() view returns(address)
func (_Esd *EsdSession) Dollar() (common.Address, error) {
	return _Esd.Contract.Dollar(&_Esd.CallOpts)
}

// Dollar is a free data retrieval call binding the contract method 0x51adeb57.
//
// Solidity: function dollar() view returns(address)
func (_Esd *EsdCallerSession) Dollar() (common.Address, error) {
	return _Esd.Contract.Dollar(&_Esd.CallOpts)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint256)
func (_Esd *EsdCaller) Epoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Esd.contract.Call(opts, &out, "epoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint256)
func (_Esd *EsdSession) Epoch() (*big.Int, error) {
	return _Esd.Contract.Epoch(&_Esd.CallOpts)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint256)
func (_Esd *EsdCallerSession) Epoch() (*big.Int, error) {
	return _Esd.Contract.Epoch(&_Esd.CallOpts)
}

// EpochTime is a free data retrieval call binding the contract method 0x5053e461.
//
// Solidity: function epochTime() view returns(uint256)
func (_Esd *EsdCaller) EpochTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Esd.contract.Call(opts, &out, "epochTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochTime is a free data retrieval call binding the contract method 0x5053e461.
//
// Solidity: function epochTime() view returns(uint256)
func (_Esd *EsdSession) EpochTime() (*big.Int, error) {
	return _Esd.Contract.EpochTime(&_Esd.CallOpts)
}

// EpochTime is a free data retrieval call binding the contract method 0x5053e461.
//
// Solidity: function epochTime() view returns(uint256)
func (_Esd *EsdCallerSession) EpochTime() (*big.Int, error) {
	return _Esd.Contract.EpochTime(&_Esd.CallOpts)
}

// ExpiringCoupons is a free data retrieval call binding the contract method 0x6a39e328.
//
// Solidity: function expiringCoupons(uint256 epoch) view returns(uint256)
func (_Esd *EsdCaller) ExpiringCoupons(opts *bind.CallOpts, epoch *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Esd.contract.Call(opts, &out, "expiringCoupons", epoch)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExpiringCoupons is a free data retrieval call binding the contract method 0x6a39e328.
//
// Solidity: function expiringCoupons(uint256 epoch) view returns(uint256)
func (_Esd *EsdSession) ExpiringCoupons(epoch *big.Int) (*big.Int, error) {
	return _Esd.Contract.ExpiringCoupons(&_Esd.CallOpts, epoch)
}

// ExpiringCoupons is a free data retrieval call binding the contract method 0x6a39e328.
//
// Solidity: function expiringCoupons(uint256 epoch) view returns(uint256)
func (_Esd *EsdCallerSession) ExpiringCoupons(epoch *big.Int) (*big.Int, error) {
	return _Esd.Contract.ExpiringCoupons(&_Esd.CallOpts, epoch)
}

// ExpiringCouponsAtIndex is a free data retrieval call binding the contract method 0x4c736099.
//
// Solidity: function expiringCouponsAtIndex(uint256 epoch, uint256 i) view returns(uint256)
func (_Esd *EsdCaller) ExpiringCouponsAtIndex(opts *bind.CallOpts, epoch *big.Int, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Esd.contract.Call(opts, &out, "expiringCouponsAtIndex", epoch, i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExpiringCouponsAtIndex is a free data retrieval call binding the contract method 0x4c736099.
//
// Solidity: function expiringCouponsAtIndex(uint256 epoch, uint256 i) view returns(uint256)
func (_Esd *EsdSession) ExpiringCouponsAtIndex(epoch *big.Int, i *big.Int) (*big.Int, error) {
	return _Esd.Contract.ExpiringCouponsAtIndex(&_Esd.CallOpts, epoch, i)
}

// ExpiringCouponsAtIndex is a free data retrieval call binding the contract method 0x4c736099.
//
// Solidity: function expiringCouponsAtIndex(uint256 epoch, uint256 i) view returns(uint256)
func (_Esd *EsdCallerSession) ExpiringCouponsAtIndex(epoch *big.Int, i *big.Int) (*big.Int, error) {
	return _Esd.Contract.ExpiringCouponsAtIndex(&_Esd.CallOpts, epoch, i)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address impl)
func (_Esd *EsdCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Esd.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address impl)
func (_Esd *EsdSession) Implementation() (common.Address, error) {
	return _Esd.Contract.Implementation(&_Esd.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address impl)
func (_Esd *EsdCallerSession) Implementation() (common.Address, error) {
	return _Esd.Contract.Implementation(&_Esd.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0xd60b347f.
//
// Solidity: function isInitialized(address candidate) view returns(bool)
func (_Esd *EsdCaller) IsInitialized(opts *bind.CallOpts, candidate common.Address) (bool, error) {
	var out []interface{}
	err := _Esd.contract.Call(opts, &out, "isInitialized", candidate)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInitialized is a free data retrieval call binding the contract method 0xd60b347f.
//
// Solidity: function isInitialized(address candidate) view returns(bool)
func (_Esd *EsdSession) IsInitialized(candidate common.Address) (bool, error) {
	return _Esd.Contract.IsInitialized(&_Esd.CallOpts, candidate)
}

// IsInitialized is a free data retrieval call binding the contract method 0xd60b347f.
//
// Solidity: function isInitialized(address candidate) view returns(bool)
func (_Esd *EsdCallerSession) IsInitialized(candidate common.Address) (bool, error) {
	return _Esd.Contract.IsInitialized(&_Esd.CallOpts, candidate)
}

// IsNominated is a free data retrieval call binding the contract method 0x3a3e6c81.
//
// Solidity: function isNominated(address candidate) view returns(bool)
func (_Esd *EsdCaller) IsNominated(opts *bind.CallOpts, candidate common.Address) (bool, error) {
	var out []interface{}
	err := _Esd.contract.Call(opts, &out, "isNominated", candidate)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsNominated is a free data retrieval call binding the contract method 0x3a3e6c81.
//
// Solidity: function isNominated(address candidate) view returns(bool)
func (_Esd *EsdSession) IsNominated(candidate common.Address) (bool, error) {
	return _Esd.Contract.IsNominated(&_Esd.CallOpts, candidate)
}

// IsNominated is a free data retrieval call binding the contract method 0x3a3e6c81.
//
// Solidity: function isNominated(address candidate) view returns(bool)
func (_Esd *EsdCallerSession) IsNominated(candidate common.Address) (bool, error) {
	return _Esd.Contract.IsNominated(&_Esd.CallOpts, candidate)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Esd *EsdCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Esd.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Esd *EsdSession) Name() (string, error) {
	return _Esd.Contract.Name(&_Esd.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Esd *EsdCallerSession) Name() (string, error) {
	return _Esd.Contract.Name(&_Esd.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Esd *EsdCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Esd.contract.Call(opts, &out, "oracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Esd *EsdSession) Oracle() (common.Address, error) {
	return _Esd.Contract.Oracle(&_Esd.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Esd *EsdCallerSession) Oracle() (common.Address, error) {
	return _Esd.Contract.Oracle(&_Esd.CallOpts)
}

// OutstandingCoupons is a free data retrieval call binding the contract method 0xc9aff70c.
//
// Solidity: function outstandingCoupons(uint256 epoch) view returns(uint256)
func (_Esd *EsdCaller) OutstandingCoupons(opts *bind.CallOpts, epoch *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Esd.contract.Call(opts, &out, "outstandingCoupons", epoch)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OutstandingCoupons is a free data retrieval call binding the contract method 0xc9aff70c.
//
// Solidity: function outstandingCoupons(uint256 epoch) view returns(uint256)
func (_Esd *EsdSession) OutstandingCoupons(epoch *big.Int) (*big.Int, error) {
	return _Esd.Contract.OutstandingCoupons(&_Esd.CallOpts, epoch)
}

// OutstandingCoupons is a free data retrieval call binding the contract method 0xc9aff70c.
//
// Solidity: function outstandingCoupons(uint256 epoch) view returns(uint256)
func (_Esd *EsdCallerSession) OutstandingCoupons(epoch *big.Int) (*big.Int, error) {
	return _Esd.Contract.OutstandingCoupons(&_Esd.CallOpts, epoch)
}

// PeriodFor is a free data retrieval call binding the contract method 0x15e14bf6.
//
// Solidity: function periodFor(address candidate) view returns(uint256)
func (_Esd *EsdCaller) PeriodFor(opts *bind.CallOpts, candidate common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Esd.contract.Call(opts, &out, "periodFor", candidate)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PeriodFor is a free data retrieval call binding the contract method 0x15e14bf6.
//
// Solidity: function periodFor(address candidate) view returns(uint256)
func (_Esd *EsdSession) PeriodFor(candidate common.Address) (*big.Int, error) {
	return _Esd.Contract.PeriodFor(&_Esd.CallOpts, candidate)
}

// PeriodFor is a free data retrieval call binding the contract method 0x15e14bf6.
//
// Solidity: function periodFor(address candidate) view returns(uint256)
func (_Esd *EsdCallerSession) PeriodFor(candidate common.Address) (*big.Int, error) {
	return _Esd.Contract.PeriodFor(&_Esd.CallOpts, candidate)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_Esd *EsdCaller) Pool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Esd.contract.Call(opts, &out, "pool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_Esd *EsdSession) Pool() (common.Address, error) {
	return _Esd.Contract.Pool(&_Esd.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_Esd *EsdCallerSession) Pool() (common.Address, error) {
	return _Esd.Contract.Pool(&_Esd.CallOpts)
}

// RecordedVote is a free data retrieval call binding the contract method 0x3fbba9a6.
//
// Solidity: function recordedVote(address account, address candidate) view returns(uint8)
func (_Esd *EsdCaller) RecordedVote(opts *bind.CallOpts, account common.Address, candidate common.Address) (uint8, error) {
	var out []interface{}
	err := _Esd.contract.Call(opts, &out, "recordedVote", account, candidate)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// RecordedVote is a free data retrieval call binding the contract method 0x3fbba9a6.
//
// Solidity: function recordedVote(address account, address candidate) view returns(uint8)
func (_Esd *EsdSession) RecordedVote(account common.Address, candidate common.Address) (uint8, error) {
	return _Esd.Contract.RecordedVote(&_Esd.CallOpts, account, candidate)
}

// RecordedVote is a free data retrieval call binding the contract method 0x3fbba9a6.
//
// Solidity: function recordedVote(address account, address candidate) view returns(uint8)
func (_Esd *EsdCallerSession) RecordedVote(account common.Address, candidate common.Address) (uint8, error) {
	return _Esd.Contract.RecordedVote(&_Esd.CallOpts, account, candidate)
}

// RejectFor is a free data retrieval call binding the contract method 0x353a420c.
//
// Solidity: function rejectFor(address candidate) view returns(uint256)
func (_Esd *EsdCaller) RejectFor(opts *bind.CallOpts, candidate common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Esd.contract.Call(opts, &out, "rejectFor", candidate)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RejectFor is a free data retrieval call binding the contract method 0x353a420c.
//
// Solidity: function rejectFor(address candidate) view returns(uint256)
func (_Esd *EsdSession) RejectFor(candidate common.Address) (*big.Int, error) {
	return _Esd.Contract.RejectFor(&_Esd.CallOpts, candidate)
}

// RejectFor is a free data retrieval call binding the contract method 0x353a420c.
//
// Solidity: function rejectFor(address candidate) view returns(uint256)
func (_Esd *EsdCallerSession) RejectFor(candidate common.Address) (*big.Int, error) {
	return _Esd.Contract.RejectFor(&_Esd.CallOpts, candidate)
}

// StartFor is a free data retrieval call binding the contract method 0xf1b7cf49.
//
// Solidity: function startFor(address candidate) view returns(uint256)
func (_Esd *EsdCaller) StartFor(opts *bind.CallOpts, candidate common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Esd.contract.Call(opts, &out, "startFor", candidate)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartFor is a free data retrieval call binding the contract method 0xf1b7cf49.
//
// Solidity: function startFor(address candidate) view returns(uint256)
func (_Esd *EsdSession) StartFor(candidate common.Address) (*big.Int, error) {
	return _Esd.Contract.StartFor(&_Esd.CallOpts, candidate)
}

// StartFor is a free data retrieval call binding the contract method 0xf1b7cf49.
//
// Solidity: function startFor(address candidate) view returns(uint256)
func (_Esd *EsdCallerSession) StartFor(candidate common.Address) (*big.Int, error) {
	return _Esd.Contract.StartFor(&_Esd.CallOpts, candidate)
}

// StatusOf is a free data retrieval call binding the contract method 0x97a5d5b5.
//
// Solidity: function statusOf(address account) view returns(uint8)
func (_Esd *EsdCaller) StatusOf(opts *bind.CallOpts, account common.Address) (uint8, error) {
	var out []interface{}
	err := _Esd.contract.Call(opts, &out, "statusOf", account)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// StatusOf is a free data retrieval call binding the contract method 0x97a5d5b5.
//
// Solidity: function statusOf(address account) view returns(uint8)
func (_Esd *EsdSession) StatusOf(account common.Address) (uint8, error) {
	return _Esd.Contract.StatusOf(&_Esd.CallOpts, account)
}

// StatusOf is a free data retrieval call binding the contract method 0x97a5d5b5.
//
// Solidity: function statusOf(address account) view returns(uint8)
func (_Esd *EsdCallerSession) StatusOf(account common.Address) (uint8, error) {
	return _Esd.Contract.StatusOf(&_Esd.CallOpts, account)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Esd *EsdCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Esd.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Esd *EsdSession) Symbol() (string, error) {
	return _Esd.Contract.Symbol(&_Esd.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Esd *EsdCallerSession) Symbol() (string, error) {
	return _Esd.Contract.Symbol(&_Esd.CallOpts)
}

// TotalBonded is a free data retrieval call binding the contract method 0x44d96e95.
//
// Solidity: function totalBonded() view returns(uint256)
func (_Esd *EsdCaller) TotalBonded(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Esd.contract.Call(opts, &out, "totalBonded")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBonded is a free data retrieval call binding the contract method 0x44d96e95.
//
// Solidity: function totalBonded() view returns(uint256)
func (_Esd *EsdSession) TotalBonded() (*big.Int, error) {
	return _Esd.Contract.TotalBonded(&_Esd.CallOpts)
}

// TotalBonded is a free data retrieval call binding the contract method 0x44d96e95.
//
// Solidity: function totalBonded() view returns(uint256)
func (_Esd *EsdCallerSession) TotalBonded() (*big.Int, error) {
	return _Esd.Contract.TotalBonded(&_Esd.CallOpts)
}

// TotalBondedAt is a free data retrieval call binding the contract method 0xffbe3b73.
//
// Solidity: function totalBondedAt(uint256 epoch) view returns(uint256)
func (_Esd *EsdCaller) TotalBondedAt(opts *bind.CallOpts, epoch *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Esd.contract.Call(opts, &out, "totalBondedAt", epoch)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBondedAt is a free data retrieval call binding the contract method 0xffbe3b73.
//
// Solidity: function totalBondedAt(uint256 epoch) view returns(uint256)
func (_Esd *EsdSession) TotalBondedAt(epoch *big.Int) (*big.Int, error) {
	return _Esd.Contract.TotalBondedAt(&_Esd.CallOpts, epoch)
}

// TotalBondedAt is a free data retrieval call binding the contract method 0xffbe3b73.
//
// Solidity: function totalBondedAt(uint256 epoch) view returns(uint256)
func (_Esd *EsdCallerSession) TotalBondedAt(epoch *big.Int) (*big.Int, error) {
	return _Esd.Contract.TotalBondedAt(&_Esd.CallOpts, epoch)
}

// TotalCoupons is a free data retrieval call binding the contract method 0x9a649edc.
//
// Solidity: function totalCoupons() view returns(uint256)
func (_Esd *EsdCaller) TotalCoupons(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Esd.contract.Call(opts, &out, "totalCoupons")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalCoupons is a free data retrieval call binding the contract method 0x9a649edc.
//
// Solidity: function totalCoupons() view returns(uint256)
func (_Esd *EsdSession) TotalCoupons() (*big.Int, error) {
	return _Esd.Contract.TotalCoupons(&_Esd.CallOpts)
}

// TotalCoupons is a free data retrieval call binding the contract method 0x9a649edc.
//
// Solidity: function totalCoupons() view returns(uint256)
func (_Esd *EsdCallerSession) TotalCoupons() (*big.Int, error) {
	return _Esd.Contract.TotalCoupons(&_Esd.CallOpts)
}

// TotalDebt is a free data retrieval call binding the contract method 0xfc7b9c18.
//
// Solidity: function totalDebt() view returns(uint256)
func (_Esd *EsdCaller) TotalDebt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Esd.contract.Call(opts, &out, "totalDebt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalDebt is a free data retrieval call binding the contract method 0xfc7b9c18.
//
// Solidity: function totalDebt() view returns(uint256)
func (_Esd *EsdSession) TotalDebt() (*big.Int, error) {
	return _Esd.Contract.TotalDebt(&_Esd.CallOpts)
}

// TotalDebt is a free data retrieval call binding the contract method 0xfc7b9c18.
//
// Solidity: function totalDebt() view returns(uint256)
func (_Esd *EsdCallerSession) TotalDebt() (*big.Int, error) {
	return _Esd.Contract.TotalDebt(&_Esd.CallOpts)
}

// TotalNet is a free data retrieval call binding the contract method 0xa6c409f1.
//
// Solidity: function totalNet() view returns(uint256)
func (_Esd *EsdCaller) TotalNet(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Esd.contract.Call(opts, &out, "totalNet")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalNet is a free data retrieval call binding the contract method 0xa6c409f1.
//
// Solidity: function totalNet() view returns(uint256)
func (_Esd *EsdSession) TotalNet() (*big.Int, error) {
	return _Esd.Contract.TotalNet(&_Esd.CallOpts)
}

// TotalNet is a free data retrieval call binding the contract method 0xa6c409f1.
//
// Solidity: function totalNet() view returns(uint256)
func (_Esd *EsdCallerSession) TotalNet() (*big.Int, error) {
	return _Esd.Contract.TotalNet(&_Esd.CallOpts)
}

// TotalRedeemable is a free data retrieval call binding the contract method 0x1edbcf6c.
//
// Solidity: function totalRedeemable() view returns(uint256)
func (_Esd *EsdCaller) TotalRedeemable(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Esd.contract.Call(opts, &out, "totalRedeemable")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalRedeemable is a free data retrieval call binding the contract method 0x1edbcf6c.
//
// Solidity: function totalRedeemable() view returns(uint256)
func (_Esd *EsdSession) TotalRedeemable() (*big.Int, error) {
	return _Esd.Contract.TotalRedeemable(&_Esd.CallOpts)
}

// TotalRedeemable is a free data retrieval call binding the contract method 0x1edbcf6c.
//
// Solidity: function totalRedeemable() view returns(uint256)
func (_Esd *EsdCallerSession) TotalRedeemable() (*big.Int, error) {
	return _Esd.Contract.TotalRedeemable(&_Esd.CallOpts)
}

// TotalStaged is a free data retrieval call binding the contract method 0xcf023779.
//
// Solidity: function totalStaged() view returns(uint256)
func (_Esd *EsdCaller) TotalStaged(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Esd.contract.Call(opts, &out, "totalStaged")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStaged is a free data retrieval call binding the contract method 0xcf023779.
//
// Solidity: function totalStaged() view returns(uint256)
func (_Esd *EsdSession) TotalStaged() (*big.Int, error) {
	return _Esd.Contract.TotalStaged(&_Esd.CallOpts)
}

// TotalStaged is a free data retrieval call binding the contract method 0xcf023779.
//
// Solidity: function totalStaged() view returns(uint256)
func (_Esd *EsdCallerSession) TotalStaged() (*big.Int, error) {
	return _Esd.Contract.TotalStaged(&_Esd.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Esd *EsdCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Esd.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Esd *EsdSession) TotalSupply() (*big.Int, error) {
	return _Esd.Contract.TotalSupply(&_Esd.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Esd *EsdCallerSession) TotalSupply() (*big.Int, error) {
	return _Esd.Contract.TotalSupply(&_Esd.CallOpts)
}

// VotesFor is a free data retrieval call binding the contract method 0xa50cd8e7.
//
// Solidity: function votesFor(address candidate) view returns(uint256)
func (_Esd *EsdCaller) VotesFor(opts *bind.CallOpts, candidate common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Esd.contract.Call(opts, &out, "votesFor", candidate)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotesFor is a free data retrieval call binding the contract method 0xa50cd8e7.
//
// Solidity: function votesFor(address candidate) view returns(uint256)
func (_Esd *EsdSession) VotesFor(candidate common.Address) (*big.Int, error) {
	return _Esd.Contract.VotesFor(&_Esd.CallOpts, candidate)
}

// VotesFor is a free data retrieval call binding the contract method 0xa50cd8e7.
//
// Solidity: function votesFor(address candidate) view returns(uint256)
func (_Esd *EsdCallerSession) VotesFor(candidate common.Address) (*big.Int, error) {
	return _Esd.Contract.VotesFor(&_Esd.CallOpts, candidate)
}

// Advance is a paid mutator transaction binding the contract method 0xea105ac7.
//
// Solidity: function advance() returns()
func (_Esd *EsdTransactor) Advance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Esd.contract.Transact(opts, "advance")
}

// Advance is a paid mutator transaction binding the contract method 0xea105ac7.
//
// Solidity: function advance() returns()
func (_Esd *EsdSession) Advance() (*types.Transaction, error) {
	return _Esd.Contract.Advance(&_Esd.TransactOpts)
}

// Advance is a paid mutator transaction binding the contract method 0xea105ac7.
//
// Solidity: function advance() returns()
func (_Esd *EsdTransactorSession) Advance() (*types.Transaction, error) {
	return _Esd.Contract.Advance(&_Esd.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Esd *EsdTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Esd.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Esd *EsdSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Esd.Contract.Approve(&_Esd.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Esd *EsdTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Esd.Contract.Approve(&_Esd.TransactOpts, spender, amount)
}

// ApproveCoupons is a paid mutator transaction binding the contract method 0x2f7f889e.
//
// Solidity: function approveCoupons(address spender, uint256 amount) returns()
func (_Esd *EsdTransactor) ApproveCoupons(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Esd.contract.Transact(opts, "approveCoupons", spender, amount)
}

// ApproveCoupons is a paid mutator transaction binding the contract method 0x2f7f889e.
//
// Solidity: function approveCoupons(address spender, uint256 amount) returns()
func (_Esd *EsdSession) ApproveCoupons(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Esd.Contract.ApproveCoupons(&_Esd.TransactOpts, spender, amount)
}

// ApproveCoupons is a paid mutator transaction binding the contract method 0x2f7f889e.
//
// Solidity: function approveCoupons(address spender, uint256 amount) returns()
func (_Esd *EsdTransactorSession) ApproveCoupons(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Esd.Contract.ApproveCoupons(&_Esd.TransactOpts, spender, amount)
}

// Bond is a paid mutator transaction binding the contract method 0x9940686e.
//
// Solidity: function bond(uint256 value) returns()
func (_Esd *EsdTransactor) Bond(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Esd.contract.Transact(opts, "bond", value)
}

// Bond is a paid mutator transaction binding the contract method 0x9940686e.
//
// Solidity: function bond(uint256 value) returns()
func (_Esd *EsdSession) Bond(value *big.Int) (*types.Transaction, error) {
	return _Esd.Contract.Bond(&_Esd.TransactOpts, value)
}

// Bond is a paid mutator transaction binding the contract method 0x9940686e.
//
// Solidity: function bond(uint256 value) returns()
func (_Esd *EsdTransactorSession) Bond(value *big.Int) (*types.Transaction, error) {
	return _Esd.Contract.Bond(&_Esd.TransactOpts, value)
}

// Commit is a paid mutator transaction binding the contract method 0x369e8c1d.
//
// Solidity: function commit(address candidate) returns()
func (_Esd *EsdTransactor) Commit(opts *bind.TransactOpts, candidate common.Address) (*types.Transaction, error) {
	return _Esd.contract.Transact(opts, "commit", candidate)
}

// Commit is a paid mutator transaction binding the contract method 0x369e8c1d.
//
// Solidity: function commit(address candidate) returns()
func (_Esd *EsdSession) Commit(candidate common.Address) (*types.Transaction, error) {
	return _Esd.Contract.Commit(&_Esd.TransactOpts, candidate)
}

// Commit is a paid mutator transaction binding the contract method 0x369e8c1d.
//
// Solidity: function commit(address candidate) returns()
func (_Esd *EsdTransactorSession) Commit(candidate common.Address) (*types.Transaction, error) {
	return _Esd.Contract.Commit(&_Esd.TransactOpts, candidate)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 value) returns()
func (_Esd *EsdTransactor) Deposit(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Esd.contract.Transact(opts, "deposit", value)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 value) returns()
func (_Esd *EsdSession) Deposit(value *big.Int) (*types.Transaction, error) {
	return _Esd.Contract.Deposit(&_Esd.TransactOpts, value)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 value) returns()
func (_Esd *EsdTransactorSession) Deposit(value *big.Int) (*types.Transaction, error) {
	return _Esd.Contract.Deposit(&_Esd.TransactOpts, value)
}

// EmergencyCommit is a paid mutator transaction binding the contract method 0x47c05069.
//
// Solidity: function emergencyCommit(address candidate) returns()
func (_Esd *EsdTransactor) EmergencyCommit(opts *bind.TransactOpts, candidate common.Address) (*types.Transaction, error) {
	return _Esd.contract.Transact(opts, "emergencyCommit", candidate)
}

// EmergencyCommit is a paid mutator transaction binding the contract method 0x47c05069.
//
// Solidity: function emergencyCommit(address candidate) returns()
func (_Esd *EsdSession) EmergencyCommit(candidate common.Address) (*types.Transaction, error) {
	return _Esd.Contract.EmergencyCommit(&_Esd.TransactOpts, candidate)
}

// EmergencyCommit is a paid mutator transaction binding the contract method 0x47c05069.
//
// Solidity: function emergencyCommit(address candidate) returns()
func (_Esd *EsdTransactorSession) EmergencyCommit(candidate common.Address) (*types.Transaction, error) {
	return _Esd.Contract.EmergencyCommit(&_Esd.TransactOpts, candidate)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Esd *EsdTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Esd.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Esd *EsdSession) Initialize() (*types.Transaction, error) {
	return _Esd.Contract.Initialize(&_Esd.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Esd *EsdTransactorSession) Initialize() (*types.Transaction, error) {
	return _Esd.Contract.Initialize(&_Esd.TransactOpts)
}

// PurchaseCoupons is a paid mutator transaction binding the contract method 0xe5f55c7e.
//
// Solidity: function purchaseCoupons(uint256 dollarAmount) returns(uint256)
func (_Esd *EsdTransactor) PurchaseCoupons(opts *bind.TransactOpts, dollarAmount *big.Int) (*types.Transaction, error) {
	return _Esd.contract.Transact(opts, "purchaseCoupons", dollarAmount)
}

// PurchaseCoupons is a paid mutator transaction binding the contract method 0xe5f55c7e.
//
// Solidity: function purchaseCoupons(uint256 dollarAmount) returns(uint256)
func (_Esd *EsdSession) PurchaseCoupons(dollarAmount *big.Int) (*types.Transaction, error) {
	return _Esd.Contract.PurchaseCoupons(&_Esd.TransactOpts, dollarAmount)
}

// PurchaseCoupons is a paid mutator transaction binding the contract method 0xe5f55c7e.
//
// Solidity: function purchaseCoupons(uint256 dollarAmount) returns(uint256)
func (_Esd *EsdTransactorSession) PurchaseCoupons(dollarAmount *big.Int) (*types.Transaction, error) {
	return _Esd.Contract.PurchaseCoupons(&_Esd.TransactOpts, dollarAmount)
}

// RedeemCoupons is a paid mutator transaction binding the contract method 0xd6a9cf08.
//
// Solidity: function redeemCoupons(uint256 couponEpoch, uint256 couponAmount) returns()
func (_Esd *EsdTransactor) RedeemCoupons(opts *bind.TransactOpts, couponEpoch *big.Int, couponAmount *big.Int) (*types.Transaction, error) {
	return _Esd.contract.Transact(opts, "redeemCoupons", couponEpoch, couponAmount)
}

// RedeemCoupons is a paid mutator transaction binding the contract method 0xd6a9cf08.
//
// Solidity: function redeemCoupons(uint256 couponEpoch, uint256 couponAmount) returns()
func (_Esd *EsdSession) RedeemCoupons(couponEpoch *big.Int, couponAmount *big.Int) (*types.Transaction, error) {
	return _Esd.Contract.RedeemCoupons(&_Esd.TransactOpts, couponEpoch, couponAmount)
}

// RedeemCoupons is a paid mutator transaction binding the contract method 0xd6a9cf08.
//
// Solidity: function redeemCoupons(uint256 couponEpoch, uint256 couponAmount) returns()
func (_Esd *EsdTransactorSession) RedeemCoupons(couponEpoch *big.Int, couponAmount *big.Int) (*types.Transaction, error) {
	return _Esd.Contract.RedeemCoupons(&_Esd.TransactOpts, couponEpoch, couponAmount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Esd *EsdTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Esd.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Esd *EsdSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Esd.Contract.Transfer(&_Esd.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Esd *EsdTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Esd.Contract.Transfer(&_Esd.TransactOpts, recipient, amount)
}

// TransferCoupons is a paid mutator transaction binding the contract method 0x005edd37.
//
// Solidity: function transferCoupons(address sender, address recipient, uint256 epoch, uint256 amount) returns()
func (_Esd *EsdTransactor) TransferCoupons(opts *bind.TransactOpts, sender common.Address, recipient common.Address, epoch *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Esd.contract.Transact(opts, "transferCoupons", sender, recipient, epoch, amount)
}

// TransferCoupons is a paid mutator transaction binding the contract method 0x005edd37.
//
// Solidity: function transferCoupons(address sender, address recipient, uint256 epoch, uint256 amount) returns()
func (_Esd *EsdSession) TransferCoupons(sender common.Address, recipient common.Address, epoch *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Esd.Contract.TransferCoupons(&_Esd.TransactOpts, sender, recipient, epoch, amount)
}

// TransferCoupons is a paid mutator transaction binding the contract method 0x005edd37.
//
// Solidity: function transferCoupons(address sender, address recipient, uint256 epoch, uint256 amount) returns()
func (_Esd *EsdTransactorSession) TransferCoupons(sender common.Address, recipient common.Address, epoch *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Esd.Contract.TransferCoupons(&_Esd.TransactOpts, sender, recipient, epoch, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Esd *EsdTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Esd.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Esd *EsdSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Esd.Contract.TransferFrom(&_Esd.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Esd *EsdTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Esd.Contract.TransferFrom(&_Esd.TransactOpts, sender, recipient, amount)
}

// Unbond is a paid mutator transaction binding the contract method 0x27de9e32.
//
// Solidity: function unbond(uint256 value) returns()
func (_Esd *EsdTransactor) Unbond(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Esd.contract.Transact(opts, "unbond", value)
}

// Unbond is a paid mutator transaction binding the contract method 0x27de9e32.
//
// Solidity: function unbond(uint256 value) returns()
func (_Esd *EsdSession) Unbond(value *big.Int) (*types.Transaction, error) {
	return _Esd.Contract.Unbond(&_Esd.TransactOpts, value)
}

// Unbond is a paid mutator transaction binding the contract method 0x27de9e32.
//
// Solidity: function unbond(uint256 value) returns()
func (_Esd *EsdTransactorSession) Unbond(value *big.Int) (*types.Transaction, error) {
	return _Esd.Contract.Unbond(&_Esd.TransactOpts, value)
}

// UnbondUnderlying is a paid mutator transaction binding the contract method 0xdf9a2b1c.
//
// Solidity: function unbondUnderlying(uint256 value) returns()
func (_Esd *EsdTransactor) UnbondUnderlying(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Esd.contract.Transact(opts, "unbondUnderlying", value)
}

// UnbondUnderlying is a paid mutator transaction binding the contract method 0xdf9a2b1c.
//
// Solidity: function unbondUnderlying(uint256 value) returns()
func (_Esd *EsdSession) UnbondUnderlying(value *big.Int) (*types.Transaction, error) {
	return _Esd.Contract.UnbondUnderlying(&_Esd.TransactOpts, value)
}

// UnbondUnderlying is a paid mutator transaction binding the contract method 0xdf9a2b1c.
//
// Solidity: function unbondUnderlying(uint256 value) returns()
func (_Esd *EsdTransactorSession) UnbondUnderlying(value *big.Int) (*types.Transaction, error) {
	return _Esd.Contract.UnbondUnderlying(&_Esd.TransactOpts, value)
}

// Vote is a paid mutator transaction binding the contract method 0xbceb514d.
//
// Solidity: function vote(address candidate, uint8 vote) returns()
func (_Esd *EsdTransactor) Vote(opts *bind.TransactOpts, candidate common.Address, vote uint8) (*types.Transaction, error) {
	return _Esd.contract.Transact(opts, "vote", candidate, vote)
}

// Vote is a paid mutator transaction binding the contract method 0xbceb514d.
//
// Solidity: function vote(address candidate, uint8 vote) returns()
func (_Esd *EsdSession) Vote(candidate common.Address, vote uint8) (*types.Transaction, error) {
	return _Esd.Contract.Vote(&_Esd.TransactOpts, candidate, vote)
}

// Vote is a paid mutator transaction binding the contract method 0xbceb514d.
//
// Solidity: function vote(address candidate, uint8 vote) returns()
func (_Esd *EsdTransactorSession) Vote(candidate common.Address, vote uint8) (*types.Transaction, error) {
	return _Esd.Contract.Vote(&_Esd.TransactOpts, candidate, vote)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 value) returns()
func (_Esd *EsdTransactor) Withdraw(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Esd.contract.Transact(opts, "withdraw", value)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 value) returns()
func (_Esd *EsdSession) Withdraw(value *big.Int) (*types.Transaction, error) {
	return _Esd.Contract.Withdraw(&_Esd.TransactOpts, value)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 value) returns()
func (_Esd *EsdTransactorSession) Withdraw(value *big.Int) (*types.Transaction, error) {
	return _Esd.Contract.Withdraw(&_Esd.TransactOpts, value)
}

// EsdAdvanceIterator is returned from FilterAdvance and is used to iterate over the raw logs and unpacked data for Advance events raised by the Esd contract.
type EsdAdvanceIterator struct {
	Event *EsdAdvance // Event containing the contract specifics and raw log

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
func (it *EsdAdvanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EsdAdvance)
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
		it.Event = new(EsdAdvance)
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
func (it *EsdAdvanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EsdAdvanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EsdAdvance represents a Advance event raised by the Esd contract.
type EsdAdvance struct {
	Epoch     *big.Int
	Block     *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAdvance is a free log retrieval operation binding the contract event 0xc30b728d1c19e5db3678b8ea9e9a063a5655071e1a325c2f7fdbca48baa90600.
//
// Solidity: event Advance(uint256 indexed epoch, uint256 block, uint256 timestamp)
func (_Esd *EsdFilterer) FilterAdvance(opts *bind.FilterOpts, epoch []*big.Int) (*EsdAdvanceIterator, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Esd.contract.FilterLogs(opts, "Advance", epochRule)
	if err != nil {
		return nil, err
	}
	return &EsdAdvanceIterator{contract: _Esd.contract, event: "Advance", logs: logs, sub: sub}, nil
}

// WatchAdvance is a free log subscription operation binding the contract event 0xc30b728d1c19e5db3678b8ea9e9a063a5655071e1a325c2f7fdbca48baa90600.
//
// Solidity: event Advance(uint256 indexed epoch, uint256 block, uint256 timestamp)
func (_Esd *EsdFilterer) WatchAdvance(opts *bind.WatchOpts, sink chan<- *EsdAdvance, epoch []*big.Int) (event.Subscription, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Esd.contract.WatchLogs(opts, "Advance", epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EsdAdvance)
				if err := _Esd.contract.UnpackLog(event, "Advance", log); err != nil {
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

// ParseAdvance is a log parse operation binding the contract event 0xc30b728d1c19e5db3678b8ea9e9a063a5655071e1a325c2f7fdbca48baa90600.
//
// Solidity: event Advance(uint256 indexed epoch, uint256 block, uint256 timestamp)
func (_Esd *EsdFilterer) ParseAdvance(log types.Log) (*EsdAdvance, error) {
	event := new(EsdAdvance)
	if err := _Esd.contract.UnpackLog(event, "Advance", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EsdBondIterator is returned from FilterBond and is used to iterate over the raw logs and unpacked data for Bond events raised by the Esd contract.
type EsdBondIterator struct {
	Event *EsdBond // Event containing the contract specifics and raw log

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
func (it *EsdBondIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EsdBond)
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
		it.Event = new(EsdBond)
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
func (it *EsdBondIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EsdBondIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EsdBond represents a Bond event raised by the Esd contract.
type EsdBond struct {
	Account         common.Address
	Start           *big.Int
	Value           *big.Int
	ValueUnderlying *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBond is a free log retrieval operation binding the contract event 0x44002fdef5a0c2d2e4e05572e9780b95aef97e0e93ffd7cc076b09fa78ff2b46.
//
// Solidity: event Bond(address indexed account, uint256 start, uint256 value, uint256 valueUnderlying)
func (_Esd *EsdFilterer) FilterBond(opts *bind.FilterOpts, account []common.Address) (*EsdBondIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Esd.contract.FilterLogs(opts, "Bond", accountRule)
	if err != nil {
		return nil, err
	}
	return &EsdBondIterator{contract: _Esd.contract, event: "Bond", logs: logs, sub: sub}, nil
}

// WatchBond is a free log subscription operation binding the contract event 0x44002fdef5a0c2d2e4e05572e9780b95aef97e0e93ffd7cc076b09fa78ff2b46.
//
// Solidity: event Bond(address indexed account, uint256 start, uint256 value, uint256 valueUnderlying)
func (_Esd *EsdFilterer) WatchBond(opts *bind.WatchOpts, sink chan<- *EsdBond, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Esd.contract.WatchLogs(opts, "Bond", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EsdBond)
				if err := _Esd.contract.UnpackLog(event, "Bond", log); err != nil {
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

// ParseBond is a log parse operation binding the contract event 0x44002fdef5a0c2d2e4e05572e9780b95aef97e0e93ffd7cc076b09fa78ff2b46.
//
// Solidity: event Bond(address indexed account, uint256 start, uint256 value, uint256 valueUnderlying)
func (_Esd *EsdFilterer) ParseBond(log types.Log) (*EsdBond, error) {
	event := new(EsdBond)
	if err := _Esd.contract.UnpackLog(event, "Bond", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EsdCommitIterator is returned from FilterCommit and is used to iterate over the raw logs and unpacked data for Commit events raised by the Esd contract.
type EsdCommitIterator struct {
	Event *EsdCommit // Event containing the contract specifics and raw log

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
func (it *EsdCommitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EsdCommit)
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
		it.Event = new(EsdCommit)
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
func (it *EsdCommitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EsdCommitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EsdCommit represents a Commit event raised by the Esd contract.
type EsdCommit struct {
	Account   common.Address
	Candidate common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCommit is a free log retrieval operation binding the contract event 0x815ca4497ab9fc80c76f210e6e842a5e198e195aa136034557eee144f790e7bb.
//
// Solidity: event Commit(address indexed account, address indexed candidate)
func (_Esd *EsdFilterer) FilterCommit(opts *bind.FilterOpts, account []common.Address, candidate []common.Address) (*EsdCommitIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var candidateRule []interface{}
	for _, candidateItem := range candidate {
		candidateRule = append(candidateRule, candidateItem)
	}

	logs, sub, err := _Esd.contract.FilterLogs(opts, "Commit", accountRule, candidateRule)
	if err != nil {
		return nil, err
	}
	return &EsdCommitIterator{contract: _Esd.contract, event: "Commit", logs: logs, sub: sub}, nil
}

// WatchCommit is a free log subscription operation binding the contract event 0x815ca4497ab9fc80c76f210e6e842a5e198e195aa136034557eee144f790e7bb.
//
// Solidity: event Commit(address indexed account, address indexed candidate)
func (_Esd *EsdFilterer) WatchCommit(opts *bind.WatchOpts, sink chan<- *EsdCommit, account []common.Address, candidate []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var candidateRule []interface{}
	for _, candidateItem := range candidate {
		candidateRule = append(candidateRule, candidateItem)
	}

	logs, sub, err := _Esd.contract.WatchLogs(opts, "Commit", accountRule, candidateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EsdCommit)
				if err := _Esd.contract.UnpackLog(event, "Commit", log); err != nil {
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

// ParseCommit is a log parse operation binding the contract event 0x815ca4497ab9fc80c76f210e6e842a5e198e195aa136034557eee144f790e7bb.
//
// Solidity: event Commit(address indexed account, address indexed candidate)
func (_Esd *EsdFilterer) ParseCommit(log types.Log) (*EsdCommit, error) {
	event := new(EsdCommit)
	if err := _Esd.contract.UnpackLog(event, "Commit", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EsdCouponApprovalIterator is returned from FilterCouponApproval and is used to iterate over the raw logs and unpacked data for CouponApproval events raised by the Esd contract.
type EsdCouponApprovalIterator struct {
	Event *EsdCouponApproval // Event containing the contract specifics and raw log

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
func (it *EsdCouponApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EsdCouponApproval)
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
		it.Event = new(EsdCouponApproval)
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
func (it *EsdCouponApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EsdCouponApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EsdCouponApproval represents a CouponApproval event raised by the Esd contract.
type EsdCouponApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCouponApproval is a free log retrieval operation binding the contract event 0x8ff27e6b95060c1ca851e7c2c28af8b413eb1a8bcb637b0290da9543a709cce3.
//
// Solidity: event CouponApproval(address indexed owner, address indexed spender, uint256 value)
func (_Esd *EsdFilterer) FilterCouponApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*EsdCouponApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Esd.contract.FilterLogs(opts, "CouponApproval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &EsdCouponApprovalIterator{contract: _Esd.contract, event: "CouponApproval", logs: logs, sub: sub}, nil
}

// WatchCouponApproval is a free log subscription operation binding the contract event 0x8ff27e6b95060c1ca851e7c2c28af8b413eb1a8bcb637b0290da9543a709cce3.
//
// Solidity: event CouponApproval(address indexed owner, address indexed spender, uint256 value)
func (_Esd *EsdFilterer) WatchCouponApproval(opts *bind.WatchOpts, sink chan<- *EsdCouponApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Esd.contract.WatchLogs(opts, "CouponApproval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EsdCouponApproval)
				if err := _Esd.contract.UnpackLog(event, "CouponApproval", log); err != nil {
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

// ParseCouponApproval is a log parse operation binding the contract event 0x8ff27e6b95060c1ca851e7c2c28af8b413eb1a8bcb637b0290da9543a709cce3.
//
// Solidity: event CouponApproval(address indexed owner, address indexed spender, uint256 value)
func (_Esd *EsdFilterer) ParseCouponApproval(log types.Log) (*EsdCouponApproval, error) {
	event := new(EsdCouponApproval)
	if err := _Esd.contract.UnpackLog(event, "CouponApproval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EsdCouponExpirationIterator is returned from FilterCouponExpiration and is used to iterate over the raw logs and unpacked data for CouponExpiration events raised by the Esd contract.
type EsdCouponExpirationIterator struct {
	Event *EsdCouponExpiration // Event containing the contract specifics and raw log

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
func (it *EsdCouponExpirationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EsdCouponExpiration)
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
		it.Event = new(EsdCouponExpiration)
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
func (it *EsdCouponExpirationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EsdCouponExpirationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EsdCouponExpiration represents a CouponExpiration event raised by the Esd contract.
type EsdCouponExpiration struct {
	Epoch          *big.Int
	CouponsExpired *big.Int
	LessRedeemable *big.Int
	LessDebt       *big.Int
	NewBonded      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterCouponExpiration is a free log retrieval operation binding the contract event 0x753df65b37159bf237ae1fca97ba1bd57cf83bc9498f271a514a4d7bafe87bda.
//
// Solidity: event CouponExpiration(uint256 indexed epoch, uint256 couponsExpired, uint256 lessRedeemable, uint256 lessDebt, uint256 newBonded)
func (_Esd *EsdFilterer) FilterCouponExpiration(opts *bind.FilterOpts, epoch []*big.Int) (*EsdCouponExpirationIterator, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Esd.contract.FilterLogs(opts, "CouponExpiration", epochRule)
	if err != nil {
		return nil, err
	}
	return &EsdCouponExpirationIterator{contract: _Esd.contract, event: "CouponExpiration", logs: logs, sub: sub}, nil
}

// WatchCouponExpiration is a free log subscription operation binding the contract event 0x753df65b37159bf237ae1fca97ba1bd57cf83bc9498f271a514a4d7bafe87bda.
//
// Solidity: event CouponExpiration(uint256 indexed epoch, uint256 couponsExpired, uint256 lessRedeemable, uint256 lessDebt, uint256 newBonded)
func (_Esd *EsdFilterer) WatchCouponExpiration(opts *bind.WatchOpts, sink chan<- *EsdCouponExpiration, epoch []*big.Int) (event.Subscription, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Esd.contract.WatchLogs(opts, "CouponExpiration", epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EsdCouponExpiration)
				if err := _Esd.contract.UnpackLog(event, "CouponExpiration", log); err != nil {
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

// ParseCouponExpiration is a log parse operation binding the contract event 0x753df65b37159bf237ae1fca97ba1bd57cf83bc9498f271a514a4d7bafe87bda.
//
// Solidity: event CouponExpiration(uint256 indexed epoch, uint256 couponsExpired, uint256 lessRedeemable, uint256 lessDebt, uint256 newBonded)
func (_Esd *EsdFilterer) ParseCouponExpiration(log types.Log) (*EsdCouponExpiration, error) {
	event := new(EsdCouponExpiration)
	if err := _Esd.contract.UnpackLog(event, "CouponExpiration", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EsdCouponPurchaseIterator is returned from FilterCouponPurchase and is used to iterate over the raw logs and unpacked data for CouponPurchase events raised by the Esd contract.
type EsdCouponPurchaseIterator struct {
	Event *EsdCouponPurchase // Event containing the contract specifics and raw log

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
func (it *EsdCouponPurchaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EsdCouponPurchase)
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
		it.Event = new(EsdCouponPurchase)
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
func (it *EsdCouponPurchaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EsdCouponPurchaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EsdCouponPurchase represents a CouponPurchase event raised by the Esd contract.
type EsdCouponPurchase struct {
	Account      common.Address
	Epoch        *big.Int
	DollarAmount *big.Int
	CouponAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCouponPurchase is a free log retrieval operation binding the contract event 0xbce252db29f761f815dc2e9ea60005af065efba6eb619d2a0b2a113fdeb61414.
//
// Solidity: event CouponPurchase(address indexed account, uint256 indexed epoch, uint256 dollarAmount, uint256 couponAmount)
func (_Esd *EsdFilterer) FilterCouponPurchase(opts *bind.FilterOpts, account []common.Address, epoch []*big.Int) (*EsdCouponPurchaseIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Esd.contract.FilterLogs(opts, "CouponPurchase", accountRule, epochRule)
	if err != nil {
		return nil, err
	}
	return &EsdCouponPurchaseIterator{contract: _Esd.contract, event: "CouponPurchase", logs: logs, sub: sub}, nil
}

// WatchCouponPurchase is a free log subscription operation binding the contract event 0xbce252db29f761f815dc2e9ea60005af065efba6eb619d2a0b2a113fdeb61414.
//
// Solidity: event CouponPurchase(address indexed account, uint256 indexed epoch, uint256 dollarAmount, uint256 couponAmount)
func (_Esd *EsdFilterer) WatchCouponPurchase(opts *bind.WatchOpts, sink chan<- *EsdCouponPurchase, account []common.Address, epoch []*big.Int) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Esd.contract.WatchLogs(opts, "CouponPurchase", accountRule, epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EsdCouponPurchase)
				if err := _Esd.contract.UnpackLog(event, "CouponPurchase", log); err != nil {
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

// ParseCouponPurchase is a log parse operation binding the contract event 0xbce252db29f761f815dc2e9ea60005af065efba6eb619d2a0b2a113fdeb61414.
//
// Solidity: event CouponPurchase(address indexed account, uint256 indexed epoch, uint256 dollarAmount, uint256 couponAmount)
func (_Esd *EsdFilterer) ParseCouponPurchase(log types.Log) (*EsdCouponPurchase, error) {
	event := new(EsdCouponPurchase)
	if err := _Esd.contract.UnpackLog(event, "CouponPurchase", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EsdCouponRedemptionIterator is returned from FilterCouponRedemption and is used to iterate over the raw logs and unpacked data for CouponRedemption events raised by the Esd contract.
type EsdCouponRedemptionIterator struct {
	Event *EsdCouponRedemption // Event containing the contract specifics and raw log

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
func (it *EsdCouponRedemptionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EsdCouponRedemption)
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
		it.Event = new(EsdCouponRedemption)
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
func (it *EsdCouponRedemptionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EsdCouponRedemptionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EsdCouponRedemption represents a CouponRedemption event raised by the Esd contract.
type EsdCouponRedemption struct {
	Account      common.Address
	Epoch        *big.Int
	CouponAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCouponRedemption is a free log retrieval operation binding the contract event 0x46e9903ae8ac9e9f0c9bc321b05965c1c036e7d4783758703f5cdfc4133c51b6.
//
// Solidity: event CouponRedemption(address indexed account, uint256 indexed epoch, uint256 couponAmount)
func (_Esd *EsdFilterer) FilterCouponRedemption(opts *bind.FilterOpts, account []common.Address, epoch []*big.Int) (*EsdCouponRedemptionIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Esd.contract.FilterLogs(opts, "CouponRedemption", accountRule, epochRule)
	if err != nil {
		return nil, err
	}
	return &EsdCouponRedemptionIterator{contract: _Esd.contract, event: "CouponRedemption", logs: logs, sub: sub}, nil
}

// WatchCouponRedemption is a free log subscription operation binding the contract event 0x46e9903ae8ac9e9f0c9bc321b05965c1c036e7d4783758703f5cdfc4133c51b6.
//
// Solidity: event CouponRedemption(address indexed account, uint256 indexed epoch, uint256 couponAmount)
func (_Esd *EsdFilterer) WatchCouponRedemption(opts *bind.WatchOpts, sink chan<- *EsdCouponRedemption, account []common.Address, epoch []*big.Int) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Esd.contract.WatchLogs(opts, "CouponRedemption", accountRule, epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EsdCouponRedemption)
				if err := _Esd.contract.UnpackLog(event, "CouponRedemption", log); err != nil {
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

// ParseCouponRedemption is a log parse operation binding the contract event 0x46e9903ae8ac9e9f0c9bc321b05965c1c036e7d4783758703f5cdfc4133c51b6.
//
// Solidity: event CouponRedemption(address indexed account, uint256 indexed epoch, uint256 couponAmount)
func (_Esd *EsdFilterer) ParseCouponRedemption(log types.Log) (*EsdCouponRedemption, error) {
	event := new(EsdCouponRedemption)
	if err := _Esd.contract.UnpackLog(event, "CouponRedemption", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EsdCouponTransferIterator is returned from FilterCouponTransfer and is used to iterate over the raw logs and unpacked data for CouponTransfer events raised by the Esd contract.
type EsdCouponTransferIterator struct {
	Event *EsdCouponTransfer // Event containing the contract specifics and raw log

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
func (it *EsdCouponTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EsdCouponTransfer)
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
		it.Event = new(EsdCouponTransfer)
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
func (it *EsdCouponTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EsdCouponTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EsdCouponTransfer represents a CouponTransfer event raised by the Esd contract.
type EsdCouponTransfer struct {
	From  common.Address
	To    common.Address
	Epoch *big.Int
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterCouponTransfer is a free log retrieval operation binding the contract event 0x0f1dbb1ccbe57a1590c7baad7b01d581b730c9ebc535dcde4345e6db424063d8.
//
// Solidity: event CouponTransfer(address indexed from, address indexed to, uint256 indexed epoch, uint256 value)
func (_Esd *EsdFilterer) FilterCouponTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, epoch []*big.Int) (*EsdCouponTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Esd.contract.FilterLogs(opts, "CouponTransfer", fromRule, toRule, epochRule)
	if err != nil {
		return nil, err
	}
	return &EsdCouponTransferIterator{contract: _Esd.contract, event: "CouponTransfer", logs: logs, sub: sub}, nil
}

// WatchCouponTransfer is a free log subscription operation binding the contract event 0x0f1dbb1ccbe57a1590c7baad7b01d581b730c9ebc535dcde4345e6db424063d8.
//
// Solidity: event CouponTransfer(address indexed from, address indexed to, uint256 indexed epoch, uint256 value)
func (_Esd *EsdFilterer) WatchCouponTransfer(opts *bind.WatchOpts, sink chan<- *EsdCouponTransfer, from []common.Address, to []common.Address, epoch []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Esd.contract.WatchLogs(opts, "CouponTransfer", fromRule, toRule, epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EsdCouponTransfer)
				if err := _Esd.contract.UnpackLog(event, "CouponTransfer", log); err != nil {
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

// ParseCouponTransfer is a log parse operation binding the contract event 0x0f1dbb1ccbe57a1590c7baad7b01d581b730c9ebc535dcde4345e6db424063d8.
//
// Solidity: event CouponTransfer(address indexed from, address indexed to, uint256 indexed epoch, uint256 value)
func (_Esd *EsdFilterer) ParseCouponTransfer(log types.Log) (*EsdCouponTransfer, error) {
	event := new(EsdCouponTransfer)
	if err := _Esd.contract.UnpackLog(event, "CouponTransfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EsdDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Esd contract.
type EsdDepositIterator struct {
	Event *EsdDeposit // Event containing the contract specifics and raw log

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
func (it *EsdDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EsdDeposit)
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
		it.Event = new(EsdDeposit)
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
func (it *EsdDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EsdDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EsdDeposit represents a Deposit event raised by the Esd contract.
type EsdDeposit struct {
	Account common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed account, uint256 value)
func (_Esd *EsdFilterer) FilterDeposit(opts *bind.FilterOpts, account []common.Address) (*EsdDepositIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Esd.contract.FilterLogs(opts, "Deposit", accountRule)
	if err != nil {
		return nil, err
	}
	return &EsdDepositIterator{contract: _Esd.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed account, uint256 value)
func (_Esd *EsdFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *EsdDeposit, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Esd.contract.WatchLogs(opts, "Deposit", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EsdDeposit)
				if err := _Esd.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed account, uint256 value)
func (_Esd *EsdFilterer) ParseDeposit(log types.Log) (*EsdDeposit, error) {
	event := new(EsdDeposit)
	if err := _Esd.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EsdIncentivizationIterator is returned from FilterIncentivization and is used to iterate over the raw logs and unpacked data for Incentivization events raised by the Esd contract.
type EsdIncentivizationIterator struct {
	Event *EsdIncentivization // Event containing the contract specifics and raw log

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
func (it *EsdIncentivizationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EsdIncentivization)
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
		it.Event = new(EsdIncentivization)
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
func (it *EsdIncentivizationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EsdIncentivizationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EsdIncentivization represents a Incentivization event raised by the Esd contract.
type EsdIncentivization struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterIncentivization is a free log retrieval operation binding the contract event 0xbb4f656853bc420ad6e4321622c07eefb4ed40e3f91b35553ce14a6dff4c0981.
//
// Solidity: event Incentivization(address indexed account, uint256 amount)
func (_Esd *EsdFilterer) FilterIncentivization(opts *bind.FilterOpts, account []common.Address) (*EsdIncentivizationIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Esd.contract.FilterLogs(opts, "Incentivization", accountRule)
	if err != nil {
		return nil, err
	}
	return &EsdIncentivizationIterator{contract: _Esd.contract, event: "Incentivization", logs: logs, sub: sub}, nil
}

// WatchIncentivization is a free log subscription operation binding the contract event 0xbb4f656853bc420ad6e4321622c07eefb4ed40e3f91b35553ce14a6dff4c0981.
//
// Solidity: event Incentivization(address indexed account, uint256 amount)
func (_Esd *EsdFilterer) WatchIncentivization(opts *bind.WatchOpts, sink chan<- *EsdIncentivization, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Esd.contract.WatchLogs(opts, "Incentivization", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EsdIncentivization)
				if err := _Esd.contract.UnpackLog(event, "Incentivization", log); err != nil {
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

// ParseIncentivization is a log parse operation binding the contract event 0xbb4f656853bc420ad6e4321622c07eefb4ed40e3f91b35553ce14a6dff4c0981.
//
// Solidity: event Incentivization(address indexed account, uint256 amount)
func (_Esd *EsdFilterer) ParseIncentivization(log types.Log) (*EsdIncentivization, error) {
	event := new(EsdIncentivization)
	if err := _Esd.contract.UnpackLog(event, "Incentivization", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EsdProposalIterator is returned from FilterProposal and is used to iterate over the raw logs and unpacked data for Proposal events raised by the Esd contract.
type EsdProposalIterator struct {
	Event *EsdProposal // Event containing the contract specifics and raw log

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
func (it *EsdProposalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EsdProposal)
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
		it.Event = new(EsdProposal)
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
func (it *EsdProposalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EsdProposalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EsdProposal represents a Proposal event raised by the Esd contract.
type EsdProposal struct {
	Candidate common.Address
	Account   common.Address
	Start     *big.Int
	Period    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProposal is a free log retrieval operation binding the contract event 0xd15e38a680a427478883cd2d32eb664cb6bb2090b0126a23ebaf3e3784b8c56b.
//
// Solidity: event Proposal(address indexed candidate, address indexed account, uint256 indexed start, uint256 period)
func (_Esd *EsdFilterer) FilterProposal(opts *bind.FilterOpts, candidate []common.Address, account []common.Address, start []*big.Int) (*EsdProposalIterator, error) {

	var candidateRule []interface{}
	for _, candidateItem := range candidate {
		candidateRule = append(candidateRule, candidateItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var startRule []interface{}
	for _, startItem := range start {
		startRule = append(startRule, startItem)
	}

	logs, sub, err := _Esd.contract.FilterLogs(opts, "Proposal", candidateRule, accountRule, startRule)
	if err != nil {
		return nil, err
	}
	return &EsdProposalIterator{contract: _Esd.contract, event: "Proposal", logs: logs, sub: sub}, nil
}

// WatchProposal is a free log subscription operation binding the contract event 0xd15e38a680a427478883cd2d32eb664cb6bb2090b0126a23ebaf3e3784b8c56b.
//
// Solidity: event Proposal(address indexed candidate, address indexed account, uint256 indexed start, uint256 period)
func (_Esd *EsdFilterer) WatchProposal(opts *bind.WatchOpts, sink chan<- *EsdProposal, candidate []common.Address, account []common.Address, start []*big.Int) (event.Subscription, error) {

	var candidateRule []interface{}
	for _, candidateItem := range candidate {
		candidateRule = append(candidateRule, candidateItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var startRule []interface{}
	for _, startItem := range start {
		startRule = append(startRule, startItem)
	}

	logs, sub, err := _Esd.contract.WatchLogs(opts, "Proposal", candidateRule, accountRule, startRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EsdProposal)
				if err := _Esd.contract.UnpackLog(event, "Proposal", log); err != nil {
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

// ParseProposal is a log parse operation binding the contract event 0xd15e38a680a427478883cd2d32eb664cb6bb2090b0126a23ebaf3e3784b8c56b.
//
// Solidity: event Proposal(address indexed candidate, address indexed account, uint256 indexed start, uint256 period)
func (_Esd *EsdFilterer) ParseProposal(log types.Log) (*EsdProposal, error) {
	event := new(EsdProposal)
	if err := _Esd.contract.UnpackLog(event, "Proposal", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EsdSupplyDecreaseIterator is returned from FilterSupplyDecrease and is used to iterate over the raw logs and unpacked data for SupplyDecrease events raised by the Esd contract.
type EsdSupplyDecreaseIterator struct {
	Event *EsdSupplyDecrease // Event containing the contract specifics and raw log

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
func (it *EsdSupplyDecreaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EsdSupplyDecrease)
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
		it.Event = new(EsdSupplyDecrease)
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
func (it *EsdSupplyDecreaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EsdSupplyDecreaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EsdSupplyDecrease represents a SupplyDecrease event raised by the Esd contract.
type EsdSupplyDecrease struct {
	Epoch   *big.Int
	Price   *big.Int
	NewDebt *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSupplyDecrease is a free log retrieval operation binding the contract event 0x5e139d4b8080a4a00dcc151e8217694aeebae893936326aa22096924a9906677.
//
// Solidity: event SupplyDecrease(uint256 indexed epoch, uint256 price, uint256 newDebt)
func (_Esd *EsdFilterer) FilterSupplyDecrease(opts *bind.FilterOpts, epoch []*big.Int) (*EsdSupplyDecreaseIterator, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Esd.contract.FilterLogs(opts, "SupplyDecrease", epochRule)
	if err != nil {
		return nil, err
	}
	return &EsdSupplyDecreaseIterator{contract: _Esd.contract, event: "SupplyDecrease", logs: logs, sub: sub}, nil
}

// WatchSupplyDecrease is a free log subscription operation binding the contract event 0x5e139d4b8080a4a00dcc151e8217694aeebae893936326aa22096924a9906677.
//
// Solidity: event SupplyDecrease(uint256 indexed epoch, uint256 price, uint256 newDebt)
func (_Esd *EsdFilterer) WatchSupplyDecrease(opts *bind.WatchOpts, sink chan<- *EsdSupplyDecrease, epoch []*big.Int) (event.Subscription, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Esd.contract.WatchLogs(opts, "SupplyDecrease", epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EsdSupplyDecrease)
				if err := _Esd.contract.UnpackLog(event, "SupplyDecrease", log); err != nil {
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

// ParseSupplyDecrease is a log parse operation binding the contract event 0x5e139d4b8080a4a00dcc151e8217694aeebae893936326aa22096924a9906677.
//
// Solidity: event SupplyDecrease(uint256 indexed epoch, uint256 price, uint256 newDebt)
func (_Esd *EsdFilterer) ParseSupplyDecrease(log types.Log) (*EsdSupplyDecrease, error) {
	event := new(EsdSupplyDecrease)
	if err := _Esd.contract.UnpackLog(event, "SupplyDecrease", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EsdSupplyIncreaseIterator is returned from FilterSupplyIncrease and is used to iterate over the raw logs and unpacked data for SupplyIncrease events raised by the Esd contract.
type EsdSupplyIncreaseIterator struct {
	Event *EsdSupplyIncrease // Event containing the contract specifics and raw log

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
func (it *EsdSupplyIncreaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EsdSupplyIncrease)
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
		it.Event = new(EsdSupplyIncrease)
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
func (it *EsdSupplyIncreaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EsdSupplyIncreaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EsdSupplyIncrease represents a SupplyIncrease event raised by the Esd contract.
type EsdSupplyIncrease struct {
	Epoch         *big.Int
	Price         *big.Int
	NewRedeemable *big.Int
	LessDebt      *big.Int
	NewBonded     *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSupplyIncrease is a free log retrieval operation binding the contract event 0x32fcaa1e76ed9517f4749d8ec9a77dd5e7329456d740b9bf9665d900eef5e283.
//
// Solidity: event SupplyIncrease(uint256 indexed epoch, uint256 price, uint256 newRedeemable, uint256 lessDebt, uint256 newBonded)
func (_Esd *EsdFilterer) FilterSupplyIncrease(opts *bind.FilterOpts, epoch []*big.Int) (*EsdSupplyIncreaseIterator, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Esd.contract.FilterLogs(opts, "SupplyIncrease", epochRule)
	if err != nil {
		return nil, err
	}
	return &EsdSupplyIncreaseIterator{contract: _Esd.contract, event: "SupplyIncrease", logs: logs, sub: sub}, nil
}

// WatchSupplyIncrease is a free log subscription operation binding the contract event 0x32fcaa1e76ed9517f4749d8ec9a77dd5e7329456d740b9bf9665d900eef5e283.
//
// Solidity: event SupplyIncrease(uint256 indexed epoch, uint256 price, uint256 newRedeemable, uint256 lessDebt, uint256 newBonded)
func (_Esd *EsdFilterer) WatchSupplyIncrease(opts *bind.WatchOpts, sink chan<- *EsdSupplyIncrease, epoch []*big.Int) (event.Subscription, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Esd.contract.WatchLogs(opts, "SupplyIncrease", epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EsdSupplyIncrease)
				if err := _Esd.contract.UnpackLog(event, "SupplyIncrease", log); err != nil {
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

// ParseSupplyIncrease is a log parse operation binding the contract event 0x32fcaa1e76ed9517f4749d8ec9a77dd5e7329456d740b9bf9665d900eef5e283.
//
// Solidity: event SupplyIncrease(uint256 indexed epoch, uint256 price, uint256 newRedeemable, uint256 lessDebt, uint256 newBonded)
func (_Esd *EsdFilterer) ParseSupplyIncrease(log types.Log) (*EsdSupplyIncrease, error) {
	event := new(EsdSupplyIncrease)
	if err := _Esd.contract.UnpackLog(event, "SupplyIncrease", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EsdSupplyNeutralIterator is returned from FilterSupplyNeutral and is used to iterate over the raw logs and unpacked data for SupplyNeutral events raised by the Esd contract.
type EsdSupplyNeutralIterator struct {
	Event *EsdSupplyNeutral // Event containing the contract specifics and raw log

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
func (it *EsdSupplyNeutralIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EsdSupplyNeutral)
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
		it.Event = new(EsdSupplyNeutral)
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
func (it *EsdSupplyNeutralIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EsdSupplyNeutralIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EsdSupplyNeutral represents a SupplyNeutral event raised by the Esd contract.
type EsdSupplyNeutral struct {
	Epoch *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSupplyNeutral is a free log retrieval operation binding the contract event 0xff7db5a0dc69b02c191ba632db46961b7d0daa1bd30709ddba9b80ad0a15d2c0.
//
// Solidity: event SupplyNeutral(uint256 indexed epoch)
func (_Esd *EsdFilterer) FilterSupplyNeutral(opts *bind.FilterOpts, epoch []*big.Int) (*EsdSupplyNeutralIterator, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Esd.contract.FilterLogs(opts, "SupplyNeutral", epochRule)
	if err != nil {
		return nil, err
	}
	return &EsdSupplyNeutralIterator{contract: _Esd.contract, event: "SupplyNeutral", logs: logs, sub: sub}, nil
}

// WatchSupplyNeutral is a free log subscription operation binding the contract event 0xff7db5a0dc69b02c191ba632db46961b7d0daa1bd30709ddba9b80ad0a15d2c0.
//
// Solidity: event SupplyNeutral(uint256 indexed epoch)
func (_Esd *EsdFilterer) WatchSupplyNeutral(opts *bind.WatchOpts, sink chan<- *EsdSupplyNeutral, epoch []*big.Int) (event.Subscription, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Esd.contract.WatchLogs(opts, "SupplyNeutral", epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EsdSupplyNeutral)
				if err := _Esd.contract.UnpackLog(event, "SupplyNeutral", log); err != nil {
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

// ParseSupplyNeutral is a log parse operation binding the contract event 0xff7db5a0dc69b02c191ba632db46961b7d0daa1bd30709ddba9b80ad0a15d2c0.
//
// Solidity: event SupplyNeutral(uint256 indexed epoch)
func (_Esd *EsdFilterer) ParseSupplyNeutral(log types.Log) (*EsdSupplyNeutral, error) {
	event := new(EsdSupplyNeutral)
	if err := _Esd.contract.UnpackLog(event, "SupplyNeutral", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EsdTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Esd contract.
type EsdTransferIterator struct {
	Event *EsdTransfer // Event containing the contract specifics and raw log

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
func (it *EsdTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EsdTransfer)
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
		it.Event = new(EsdTransfer)
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
func (it *EsdTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EsdTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EsdTransfer represents a Transfer event raised by the Esd contract.
type EsdTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Esd *EsdFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EsdTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Esd.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EsdTransferIterator{contract: _Esd.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Esd *EsdFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *EsdTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Esd.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EsdTransfer)
				if err := _Esd.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Esd *EsdFilterer) ParseTransfer(log types.Log) (*EsdTransfer, error) {
	event := new(EsdTransfer)
	if err := _Esd.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EsdUnbondIterator is returned from FilterUnbond and is used to iterate over the raw logs and unpacked data for Unbond events raised by the Esd contract.
type EsdUnbondIterator struct {
	Event *EsdUnbond // Event containing the contract specifics and raw log

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
func (it *EsdUnbondIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EsdUnbond)
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
		it.Event = new(EsdUnbond)
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
func (it *EsdUnbondIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EsdUnbondIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EsdUnbond represents a Unbond event raised by the Esd contract.
type EsdUnbond struct {
	Account         common.Address
	Start           *big.Int
	Value           *big.Int
	ValueUnderlying *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnbond is a free log retrieval operation binding the contract event 0x93530ac0ee8c50e696e13c5ac62355d0c0ba4bd943620d5bda1eb08b64ae7512.
//
// Solidity: event Unbond(address indexed account, uint256 start, uint256 value, uint256 valueUnderlying)
func (_Esd *EsdFilterer) FilterUnbond(opts *bind.FilterOpts, account []common.Address) (*EsdUnbondIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Esd.contract.FilterLogs(opts, "Unbond", accountRule)
	if err != nil {
		return nil, err
	}
	return &EsdUnbondIterator{contract: _Esd.contract, event: "Unbond", logs: logs, sub: sub}, nil
}

// WatchUnbond is a free log subscription operation binding the contract event 0x93530ac0ee8c50e696e13c5ac62355d0c0ba4bd943620d5bda1eb08b64ae7512.
//
// Solidity: event Unbond(address indexed account, uint256 start, uint256 value, uint256 valueUnderlying)
func (_Esd *EsdFilterer) WatchUnbond(opts *bind.WatchOpts, sink chan<- *EsdUnbond, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Esd.contract.WatchLogs(opts, "Unbond", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EsdUnbond)
				if err := _Esd.contract.UnpackLog(event, "Unbond", log); err != nil {
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

// ParseUnbond is a log parse operation binding the contract event 0x93530ac0ee8c50e696e13c5ac62355d0c0ba4bd943620d5bda1eb08b64ae7512.
//
// Solidity: event Unbond(address indexed account, uint256 start, uint256 value, uint256 valueUnderlying)
func (_Esd *EsdFilterer) ParseUnbond(log types.Log) (*EsdUnbond, error) {
	event := new(EsdUnbond)
	if err := _Esd.contract.UnpackLog(event, "Unbond", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EsdUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Esd contract.
type EsdUpgradedIterator struct {
	Event *EsdUpgraded // Event containing the contract specifics and raw log

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
func (it *EsdUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EsdUpgraded)
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
		it.Event = new(EsdUpgraded)
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
func (it *EsdUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EsdUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EsdUpgraded represents a Upgraded event raised by the Esd contract.
type EsdUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Esd *EsdFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*EsdUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Esd.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &EsdUpgradedIterator{contract: _Esd.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Esd *EsdFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *EsdUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Esd.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EsdUpgraded)
				if err := _Esd.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Esd *EsdFilterer) ParseUpgraded(log types.Log) (*EsdUpgraded, error) {
	event := new(EsdUpgraded)
	if err := _Esd.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EsdVoteIterator is returned from FilterVote and is used to iterate over the raw logs and unpacked data for Vote events raised by the Esd contract.
type EsdVoteIterator struct {
	Event *EsdVote // Event containing the contract specifics and raw log

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
func (it *EsdVoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EsdVote)
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
		it.Event = new(EsdVote)
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
func (it *EsdVoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EsdVoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EsdVote represents a Vote event raised by the Esd contract.
type EsdVote struct {
	Account   common.Address
	Candidate common.Address
	Vote      uint8
	Bonded    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterVote is a free log retrieval operation binding the contract event 0xbe50c78cbc15b0864819aadea36c6499da421b33c38e2ef19bebda583c708c78.
//
// Solidity: event Vote(address indexed account, address indexed candidate, uint8 vote, uint256 bonded)
func (_Esd *EsdFilterer) FilterVote(opts *bind.FilterOpts, account []common.Address, candidate []common.Address) (*EsdVoteIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var candidateRule []interface{}
	for _, candidateItem := range candidate {
		candidateRule = append(candidateRule, candidateItem)
	}

	logs, sub, err := _Esd.contract.FilterLogs(opts, "Vote", accountRule, candidateRule)
	if err != nil {
		return nil, err
	}
	return &EsdVoteIterator{contract: _Esd.contract, event: "Vote", logs: logs, sub: sub}, nil
}

// WatchVote is a free log subscription operation binding the contract event 0xbe50c78cbc15b0864819aadea36c6499da421b33c38e2ef19bebda583c708c78.
//
// Solidity: event Vote(address indexed account, address indexed candidate, uint8 vote, uint256 bonded)
func (_Esd *EsdFilterer) WatchVote(opts *bind.WatchOpts, sink chan<- *EsdVote, account []common.Address, candidate []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var candidateRule []interface{}
	for _, candidateItem := range candidate {
		candidateRule = append(candidateRule, candidateItem)
	}

	logs, sub, err := _Esd.contract.WatchLogs(opts, "Vote", accountRule, candidateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EsdVote)
				if err := _Esd.contract.UnpackLog(event, "Vote", log); err != nil {
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

// ParseVote is a log parse operation binding the contract event 0xbe50c78cbc15b0864819aadea36c6499da421b33c38e2ef19bebda583c708c78.
//
// Solidity: event Vote(address indexed account, address indexed candidate, uint8 vote, uint256 bonded)
func (_Esd *EsdFilterer) ParseVote(log types.Log) (*EsdVote, error) {
	event := new(EsdVote)
	if err := _Esd.contract.UnpackLog(event, "Vote", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EsdWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Esd contract.
type EsdWithdrawIterator struct {
	Event *EsdWithdraw // Event containing the contract specifics and raw log

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
func (it *EsdWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EsdWithdraw)
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
		it.Event = new(EsdWithdraw)
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
func (it *EsdWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EsdWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EsdWithdraw represents a Withdraw event raised by the Esd contract.
type EsdWithdraw struct {
	Account common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed account, uint256 value)
func (_Esd *EsdFilterer) FilterWithdraw(opts *bind.FilterOpts, account []common.Address) (*EsdWithdrawIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Esd.contract.FilterLogs(opts, "Withdraw", accountRule)
	if err != nil {
		return nil, err
	}
	return &EsdWithdrawIterator{contract: _Esd.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed account, uint256 value)
func (_Esd *EsdFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *EsdWithdraw, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Esd.contract.WatchLogs(opts, "Withdraw", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EsdWithdraw)
				if err := _Esd.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed account, uint256 value)
func (_Esd *EsdFilterer) ParseWithdraw(log types.Log) (*EsdWithdraw, error) {
	event := new(EsdWithdraw)
	if err := _Esd.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	return event, nil
}
