// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package FLApp

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
	_ = abi.ConvertType
)

// ChannelAllocation is an auto generated low-level Go binding around an user-defined struct.
type ChannelAllocation struct {
	Assets   []common.Address
	Balances [][]*big.Int
	Locked   []ChannelSubAlloc
}

// ChannelParams is an auto generated low-level Go binding around an user-defined struct.
type ChannelParams struct {
	ChallengeDuration *big.Int
	Nonce             *big.Int
	Participants      []common.Address
	App               common.Address
	LedgerChannel     bool
	VirtualChannel    bool
}

// ChannelState is an auto generated low-level Go binding around an user-defined struct.
type ChannelState struct {
	ChannelID [32]byte
	Version   uint64
	Outcome   ChannelAllocation
	AppData   []byte
	IsFinal   bool
}

// ChannelSubAlloc is an auto generated low-level Go binding around an user-defined struct.
type ChannelSubAlloc struct {
	ID       [32]byte
	Balances []*big.Int
	IndexMap []uint16
}

// FLAppMetaData contains all meta data concerning the FLApp contract.
var FLAppMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"challengeDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"participants\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"app\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"ledgerChannel\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"virtualChannel\",\"type\":\"bool\"}],\"internalType\":\"structChannel.Params\",\"name\":\"params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"channelID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[][]\",\"name\":\"balances\",\"type\":\"uint256[][]\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"ID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint16[]\",\"name\":\"indexMap\",\"type\":\"uint16[]\"}],\"internalType\":\"structChannel.SubAlloc[]\",\"name\":\"locked\",\"type\":\"tuple[]\"}],\"internalType\":\"structChannel.Allocation\",\"name\":\"outcome\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structChannel.State\",\"name\":\"from\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"channelID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[][]\",\"name\":\"balances\",\"type\":\"uint256[][]\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"ID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint16[]\",\"name\":\"indexMap\",\"type\":\"uint16[]\"}],\"internalType\":\"structChannel.SubAlloc[]\",\"name\":\"locked\",\"type\":\"tuple[]\"}],\"internalType\":\"structChannel.Allocation\",\"name\":\"outcome\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structChannel.State\",\"name\":\"to\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"signerIdx\",\"type\":\"uint256\"}],\"name\":\"validTransition\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610e74806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c80630d1feb4f14610030575b600080fd5b61004361003e366004610a4e565b610045565b005b60026100546040860186610c44565b90501461007c5760405162461bcd60e51b815260040161007390610bb8565b60405180910390fd5b600061008b6060850185610c92565b60008161009457fe5b919091013560f81c9150600690506100af6060850185610c92565b9050146100ce5760405162461bcd60e51b815260040161007390610be8565b818160ff16146100f05760405162461bcd60e51b815260040161007390610b0f565b6100fd6060840184610c92565b60008161010657fe5b9091013560f81c6001838101161490506101325760405162461bcd60e51b815260040161007390610b39565b600080806101806101466060880188610c92565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061049d92505050565b9194509250905082151561019a60a0880160808901610a27565b1515146101b95760405162461bcd60e51b815260040161007390610b5d565b6102536101c96040880188610cd7565b6101d39080610c44565b808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152506102129250505060408a018a610cd7565b61021c9080610c44565b8080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525061056492505050565b6102a36102636040880188610cd7565b610271906040810190610c44565b61027a91610d86565b61028760408a018a610cd7565b610295906040810190610c44565b61029e91610d86565b61065f565b60006102b26040890189610cd7565b6102c0906020810190610c44565b6102c991610d38565b9050821561046557805160018390039067ffffffffffffffff811180156102ef57600080fd5b5060405190808252806020026020018201604052801561032357816020015b606081526020019060019003908161030e5790505b50915060005b82518110156104625760408051600280825260608201835290916020830190803683370190505083828151811061035c57fe5b602090810291909101015261037460408b018b610cd7565b610382906020810190610c44565b8281811061038c57fe5b905060200281019061039e9190610c44565b60018181106103a957fe5b905060200201358a80604001906103c09190610cd7565b6103ce906020810190610c44565b838181106103d857fe5b90506020028101906103ea9190610c44565b60008181106103f557fe5b905060200201350183828151811061040957fe5b60200260200101518560ff168151811061041f57fe5b602002602001018181525050600083828151811061043957fe5b60200260200101518360ff168151811061044f57fe5b6020908102919091010152600101610329565b50505b6104926104756040890189610cd7565b610483906020810190610c44565b61048c91610d38565b826106c4565b505050505050505050565b805160009081908190849060049081106104b357fe5b01602001516001600160f81b0319161580159061051257508351849060059081106104da57fe5b01602001516001600160f81b031916151580610512575083518490600690811061050057fe5b01602001516001600160f81b03191615155b15610553578351603c908590600590811061052957fe5b016020015160f81c106105445750600191508190508061055d565b5060019150819050600061055d565b5060009150819050805b9193909250565b80518251146105ba576040805162461bcd60e51b815260206004820152601960248201527f616464726573735b5d3a20756e657175616c206c656e67746800000000000000604482015290519081900360640190fd5b60005b825181101561065a578181815181106105d257fe5b60200260200101516001600160a01b03168382815181106105ef57fe5b60200260200101516001600160a01b031614610652576040805162461bcd60e51b815260206004820152601760248201527f616464726573735b5d3a20756e657175616c206974656d000000000000000000604482015290519081900360640190fd5b6001016105bd565b505050565b80518251146106805760405162461bcd60e51b815260040161007390610c0d565b60005b825181101561065a576106bc83828151811061069b57fe5b60200260200101518383815181106106af57fe5b6020026020010151610729565b600101610683565b80518251146106e55760405162461bcd60e51b815260040161007390610b81565b60005b825181101561065a5761072183828151811061070057fe5b602002602001015183838151811061071457fe5b6020026020010151610772565b6001016106e8565b805182511461074a5760405162461bcd60e51b815260040161007390610ae1565b61075c82602001518260200151610772565b61076e82604001518260400151610856565b5050565b80518251146107c8576040805162461bcd60e51b815260206004820152601960248201527f75696e743235365b5d3a20756e657175616c206c656e67746800000000000000604482015290519081900360640190fd5b60005b825181101561065a578181815181106107e057fe5b60200260200101518382815181106107f457fe5b60200260200101511461084e576040805162461bcd60e51b815260206004820152601760248201527f75696e743235365b5d3a20756e657175616c206974656d000000000000000000604482015290519081900360640190fd5b6001016107cb565b80518251146108ac576040805162461bcd60e51b815260206004820152601860248201527f75696e7431365b5d3a20756e657175616c206c656e6774680000000000000000604482015290519081900360640190fd5b60005b825181101561065a578181815181106108c457fe5b602002602001015161ffff168382815181106108dc57fe5b602002602001015161ffff1614610933576040805162461bcd60e51b815260206004820152601660248201527575696e7431365b5d3a20756e657175616c206974656d60501b604482015290519081900360640190fd5b6001016108af565b600082601f83011261094b578081fd5b8135602061096061095b83610d1a565b610cf6565b828152818101908583018385028701840188101561097c578586fd5b855b858110156109a957813561ffff81168114610997578788fd5b8452928401929084019060010161097e565b5090979650505050505050565b600082601f8301126109c6578081fd5b813560206109d661095b83610d1a565b82815281810190858301838502870184018810156109f2578586fd5b855b858110156109a9578135845292840192908401906001016109f4565b600060a08284031215610a21578081fd5b50919050565b600060208284031215610a38578081fd5b81358015158114610a47578182fd5b9392505050565b60008060008060808587031215610a63578283fd5b843567ffffffffffffffff80821115610a7a578485fd5b9086019060c08289031215610a8d578485fd5b90945060208601359080821115610aa2578485fd5b610aae88838901610a10565b94506040870135915080821115610ac3578384fd5b50610ad087828801610a10565b949793965093946060013593505050565b60208082526014908201527314dd58905b1b1bd8ce881d5b995c5d585b08125160621b604082015260600190565b60208082526010908201526f30b1ba37b9103737ba1039b4b3b732b960811b604082015260600190565b6020808252600a90820152693732bc3a1030b1ba37b960b11b604082015260600190565b6020808252600a908201526966696e616c20666c616760b01b604082015260600190565b6020808252601b908201527f75696e743235365b5d5b5d3a20756e657175616c206c656e6774680000000000604082015260600190565b6020808252601690820152756e756d626572206f66207061727469636970616e747360501b604082015260600190565b6020808252600b908201526a0c8c2e8c240d8cadccee8d60ab1b604082015260600190565b6020808252601a908201527f537562416c6c6f635b5d3a20756e657175616c206c656e677468000000000000604082015260600190565b6000808335601e19843603018112610c5a578283fd5b83018035915067ffffffffffffffff821115610c74578283fd5b6020908101925081023603821315610c8b57600080fd5b9250929050565b6000808335601e19843603018112610ca8578283fd5b83018035915067ffffffffffffffff821115610cc2578283fd5b602001915036819003821315610c8b57600080fd5b60008235605e19833603018112610cec578182fd5b9190910192915050565b60405181810167ffffffffffffffff81118282101715610d1257fe5b604052919050565b600067ffffffffffffffff821115610d2e57fe5b5060209081020190565b6000610d4661095b84610d1a565b8381526020808201919084845b87811015610d7a57610d6836833589016109b6565b85529382019390820190600101610d53565b50919695505050505050565b6000610d9461095b84610d1a565b8381526020808201919084845b87811015610d7a57813587016060808236031215610dbd578788fd5b604080519182019167ffffffffffffffff8084118285101715610ddc57fe5b92825283358152868401359280841115610df4578a8bfd5b610e00368587016109b6565b8883015282850135935080841115610e16578a8bfd5b50610e233684860161093b565b91810191909152875250509382019390820190600101610da156fea264697066735822122009c329e7ce9a2b3bd98fcf10b7b9651fc096102d28ae3c7ff7e6eb51c76fd17164736f6c63430007060033",
}

// FLAppABI is the input ABI used to generate the binding from.
// Deprecated: Use FLAppMetaData.ABI instead.
var FLAppABI = FLAppMetaData.ABI

// FLAppBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FLAppMetaData.Bin instead.
var FLAppBin = FLAppMetaData.Bin

// DeployFLApp deploys a new Ethereum contract, binding an instance of FLApp to it.
func DeployFLApp(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *FLApp, error) {
	parsed, err := FLAppMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FLAppBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FLApp{FLAppCaller: FLAppCaller{contract: contract}, FLAppTransactor: FLAppTransactor{contract: contract}, FLAppFilterer: FLAppFilterer{contract: contract}}, nil
}

// FLApp is an auto generated Go binding around an Ethereum contract.
type FLApp struct {
	FLAppCaller     // Read-only binding to the contract
	FLAppTransactor // Write-only binding to the contract
	FLAppFilterer   // Log filterer for contract events
}

// FLAppCaller is an auto generated read-only Go binding around an Ethereum contract.
type FLAppCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FLAppTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FLAppTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FLAppFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FLAppFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FLAppSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FLAppSession struct {
	Contract     *FLApp            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FLAppCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FLAppCallerSession struct {
	Contract *FLAppCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// FLAppTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FLAppTransactorSession struct {
	Contract     *FLAppTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FLAppRaw is an auto generated low-level Go binding around an Ethereum contract.
type FLAppRaw struct {
	Contract *FLApp // Generic contract binding to access the raw methods on
}

// FLAppCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FLAppCallerRaw struct {
	Contract *FLAppCaller // Generic read-only contract binding to access the raw methods on
}

// FLAppTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FLAppTransactorRaw struct {
	Contract *FLAppTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFLApp creates a new instance of FLApp, bound to a specific deployed contract.
func NewFLApp(address common.Address, backend bind.ContractBackend) (*FLApp, error) {
	contract, err := bindFLApp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FLApp{FLAppCaller: FLAppCaller{contract: contract}, FLAppTransactor: FLAppTransactor{contract: contract}, FLAppFilterer: FLAppFilterer{contract: contract}}, nil
}

// NewFLAppCaller creates a new read-only instance of FLApp, bound to a specific deployed contract.
func NewFLAppCaller(address common.Address, caller bind.ContractCaller) (*FLAppCaller, error) {
	contract, err := bindFLApp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FLAppCaller{contract: contract}, nil
}

// NewFLAppTransactor creates a new write-only instance of FLApp, bound to a specific deployed contract.
func NewFLAppTransactor(address common.Address, transactor bind.ContractTransactor) (*FLAppTransactor, error) {
	contract, err := bindFLApp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FLAppTransactor{contract: contract}, nil
}

// NewFLAppFilterer creates a new log filterer instance of FLApp, bound to a specific deployed contract.
func NewFLAppFilterer(address common.Address, filterer bind.ContractFilterer) (*FLAppFilterer, error) {
	contract, err := bindFLApp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FLAppFilterer{contract: contract}, nil
}

// bindFLApp binds a generic wrapper to an already deployed contract.
func bindFLApp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FLAppMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FLApp *FLAppRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FLApp.Contract.FLAppCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FLApp *FLAppRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FLApp.Contract.FLAppTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FLApp *FLAppRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FLApp.Contract.FLAppTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FLApp *FLAppCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FLApp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FLApp *FLAppTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FLApp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FLApp *FLAppTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FLApp.Contract.contract.Transact(opts, method, params...)
}

// ValidTransition is a free data retrieval call binding the contract method 0x0d1feb4f.
//
// Solidity: function validTransition((uint256,uint256,address[],address,bool,bool) params, (bytes32,uint64,(address[],uint256[][],(bytes32,uint256[],uint16[])[]),bytes,bool) from, (bytes32,uint64,(address[],uint256[][],(bytes32,uint256[],uint16[])[]),bytes,bool) to, uint256 signerIdx) pure returns()
func (_FLApp *FLAppCaller) ValidTransition(opts *bind.CallOpts, params ChannelParams, from ChannelState, to ChannelState, signerIdx *big.Int) error {
	var out []interface{}
	err := _FLApp.contract.Call(opts, &out, "validTransition", params, from, to, signerIdx)

	if err != nil {
		return err
	}

	return err

}

// ValidTransition is a free data retrieval call binding the contract method 0x0d1feb4f.
//
// Solidity: function validTransition((uint256,uint256,address[],address,bool,bool) params, (bytes32,uint64,(address[],uint256[][],(bytes32,uint256[],uint16[])[]),bytes,bool) from, (bytes32,uint64,(address[],uint256[][],(bytes32,uint256[],uint16[])[]),bytes,bool) to, uint256 signerIdx) pure returns()
func (_FLApp *FLAppSession) ValidTransition(params ChannelParams, from ChannelState, to ChannelState, signerIdx *big.Int) error {
	return _FLApp.Contract.ValidTransition(&_FLApp.CallOpts, params, from, to, signerIdx)
}

// ValidTransition is a free data retrieval call binding the contract method 0x0d1feb4f.
//
// Solidity: function validTransition((uint256,uint256,address[],address,bool,bool) params, (bytes32,uint64,(address[],uint256[][],(bytes32,uint256[],uint16[])[]),bytes,bool) from, (bytes32,uint64,(address[],uint256[][],(bytes32,uint256[],uint16[])[]),bytes,bool) to, uint256 signerIdx) pure returns()
func (_FLApp *FLAppCallerSession) ValidTransition(params ChannelParams, from ChannelState, to ChannelState, signerIdx *big.Int) error {
	return _FLApp.Contract.ValidTransition(&_FLApp.CallOpts, params, from, to, signerIdx)
}
