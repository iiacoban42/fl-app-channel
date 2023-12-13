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
	Bin: "0x608060405234801561001057600080fd5b50611ae0806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c80630d1feb4f14610030575b600080fd5b61004361003e36600461136f565b610045565b005b60026100546040860186611888565b90501461007c5760405162461bcd60e51b8152600401610073906116dc565b60405180910390fd5b600061008b60608501856118d6565b60008161009457fe5b919091013560f81c9150508181146100be5760405162461bcd60e51b8152600401610073906115dc565b6100cb60608401846118d6565b6000816100d457fe5b9091013560f81c6001838101161490506101005760405162461bcd60e51b81526004016100739061165d565b61010d60608501856118d6565b600481811061011857fe5b919091013560f81c1590506102095761013460608401846118d6565b600181811061013f57fe5b909101356001600160f81b031916905061015c60608601866118d6565b600181811061016757fe5b9050013560f81c60f81b6001600160f81b031916146101985760405162461bcd60e51b815260040161007390611861565b6101a560608401846118d6565b60028181106101b057fe5b909101356001600160f81b03191690506101cd60608601866118d6565b60028181106101d857fe5b9050013560f81c60f81b6001600160f81b031916146102095760405162461bcd60e51b815260040161007390611636565b600061021860608501856118d6565b600281811061022357fe5b919091013560f81c915050600580820190828001808201919084010160ff811661025060608901896118d6565b90501461026f5760405162461bcd60e51b81526004016100739061170c565b61027c60608801886118d6565b600281811061028757fe5b919091013560f81c905061029e60608901896118d6565b60038181106102a957fe5b9050013560f81c60f81b60f81c60ff1611156102d75760405162461bcd60e51b81526004016100739061152f565b60ff851661069a576103936102ef60608a018a6118d6565b61030391600588810160ff1692909161197c565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506103459250505060608a018a6118d6565b61035991600589810160ff1692909161197c565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610cb792505050565b6103af5760405162461bcd60e51b815260040161007390611439565b6104876103bf60608a018a6118d6565b6103d49160ff8789018116929088169161197c565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506104169250505060608a018a6118d6565b61042b9160ff888a018116929089169161197c565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061046d9250505060608c018c6118d6565b600381811061047857fe5b919091013560f81c9050610d35565b6104a35760405162461bcd60e51b815260040161007390611768565b61051f6104b360608a018a6118d6565b6104c89160ff8689018116929087169161197c565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061050a9250505060608a018a6118d6565b61042b9160ff878a018116929088169161197c565b61053b5760405162461bcd60e51b8152600401610073906114de565b61054860608901896118d6565b600481811061055357fe5b919091013560f81c15905061069a57600161057160608a018a6118d6565b600381811061057c57fe5b60ff92013560f81c9290920116905061059860608901896118d6565b60038181106105a357fe5b9050013560f81c60f81b60f81c60ff16146105d05760405162461bcd60e51b815260040161007390611402565b6105dd60608801886118d6565b6105ea60608b018b6118d6565b60038181106105f557fe5b919091013560f81c860160ff16905081811061060d57fe5b919091013560f81c151590506106355760405162461bcd60e51b8152600401610073906117bc565b61064260608801886118d6565b61064f60608b018b6118d6565b600381811061065a57fe5b919091013560f81c850160ff16905081811061067257fe5b919091013560f81c1515905061069a5760405162461bcd60e51b815260040161007390611606565b8460ff166001141561097c576106b360608801886118d6565b60038181106106be57fe5b919091013560f81c90506106d560608a018a6118d6565b60038181106106e057fe5b9050013560f81c60f81b60f81c60ff161461070d5760405162461bcd60e51b815260040161007390611470565b61078961071d60608a018a6118d6565b6107329160ff8789018116929088169161197c565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506107749250505060608a018a6118d6565b6103599160ff888a018116929089169161197c565b6107a55760405162461bcd60e51b81526004016100739061182a565b6108216107b560608a018a6118d6565b6107ca9160ff8689018116929087169161197c565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061080c9250505060608a018a6118d6565b6103599160ff878a018116929088169161197c565b61083d5760405162461bcd60e51b8152600401610073906114a7565b6108f961084d60608a018a6118d6565b61086191600588810160ff1692909161197c565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506108a39250505060608a018a6118d6565b6108b791600589810160ff1692909161197c565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061046d9250505060608b018b6118d6565b6109155760405162461bcd60e51b81526004016100739061155c565b61092260608801886118d6565b61092f60608b018b6118d6565b600381811061093a57fe5b600592013560f81c9190910160ff16905081811061095457fe5b919091013560f81c1515905061097c5760405162461bcd60e51b815260040161007390611731565b60006109c861098e60608a018a6118d6565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610dc492505050565b90508015156109dd60a08a0160808b0161134f565b1515146109fc5760405162461bcd60e51b815260040161007390611681565b610a96610a0c60408a018a61191b565b610a169080611888565b80806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250610a559250505060408c018c61191b565b610a5f9080611888565b80806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250610e4092505050565b610ae6610aa660408a018a61191b565b610ab4906040810190611888565b610abd916119f2565b610aca60408c018c61191b565b610ad8906040810190611888565b610ae1916119f2565b610f3b565b6000610af560408b018b61191b565b610b03906020810190611888565b610b0c916119a4565b9050610b58610b1e60608b018b6118d6565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610fa092505050565b15610c7d57670de0b6b3a764000060016000805b8451811015610c7857838e8060400190610b86919061191b565b610b94906020810190611888565b83818110610b9e57fe5b9050602002810190610bb09190611888565b8560ff16818110610bbd57fe5b9050602002013501858281518110610bd157fe5b60200260200101518460ff1681518110610be757fe5b602002602001018181525050838e8060400190610c04919061191b565b610c12906020810190611888565b83818110610c1c57fe5b9050602002810190610c2e9190611888565b8460ff16818110610c3b57fe5b9050602002013503858281518110610c4f57fe5b60200260200101518360ff1681518110610c6557fe5b6020908102919091010152600101610b6c565b505050505b610caa610c8d60408b018b61191b565b610c9b906020810190611888565b610ca4916119a4565b82610fec565b5050505050505050505050565b60008151835114610cca57506000610d2f565b60005b8351811015610d2957828181518110610ce257fe5b602001015160f81c60f81b6001600160f81b031916848281518110610d0357fe5b01602001516001600160f81b03191614610d21576000915050610d2f565b600101610ccd565b50600190505b92915050565b60008251845114610d4857506000610dbd565b60005b8451811015610db7578260ff16811415610d6457610daf565b838181518110610d7057fe5b602001015160f81c60f81b6001600160f81b031916858281518110610d9157fe5b01602001516001600160f81b03191614610daf576000915050610dbd565b600101610d4b565b50600190505b9392505050565b805160009082906003908110610dd657fe5b016020015182516001600160f81b03199091169083906002908110610df757fe5b01602001516001600160f81b031916148015610e2a5750815182906002908110610e1d57fe5b60209101015160f81c6003145b15610e3757506001610e3b565b5060005b919050565b8051825114610e96576040805162461bcd60e51b815260206004820152601960248201527f616464726573735b5d3a20756e657175616c206c656e67746800000000000000604482015290519081900360640190fd5b60005b8251811015610f3657818181518110610eae57fe5b60200260200101516001600160a01b0316838281518110610ecb57fe5b60200260200101516001600160a01b031614610f2e576040805162461bcd60e51b815260206004820152601760248201527f616464726573735b5d3a20756e657175616c206974656d000000000000000000604482015290519081900360640190fd5b600101610e99565b505050565b8051825114610f5c5760405162461bcd60e51b8152600401610073906117f3565b60005b8251811015610f3657610f98838281518110610f7757fe5b6020026020010151838381518110610f8b57fe5b6020026020010151611051565b600101610f5f565b805160009060029083906004908110610fb557fe5b016020015160f81c148015610e2a5750815160009083906003908110610fd757fe5b016020015160f81c14610e3757506001610e3b565b805182511461100d5760405162461bcd60e51b8152600401610073906116a5565b60005b8251811015610f365761104983828151811061102857fe5b602002602001015183838151811061103c57fe5b602002602001015161109a565b600101611010565b80518251146110725760405162461bcd60e51b8152600401610073906115ae565b6110848260200151826020015161109a565b6110968260400151826040015161117e565b5050565b80518251146110f0576040805162461bcd60e51b815260206004820152601960248201527f75696e743235365b5d3a20756e657175616c206c656e67746800000000000000604482015290519081900360640190fd5b60005b8251811015610f365781818151811061110857fe5b602002602001015183828151811061111c57fe5b602002602001015114611176576040805162461bcd60e51b815260206004820152601760248201527f75696e743235365b5d3a20756e657175616c206974656d000000000000000000604482015290519081900360640190fd5b6001016110f3565b80518251146111d4576040805162461bcd60e51b815260206004820152601860248201527f75696e7431365b5d3a20756e657175616c206c656e6774680000000000000000604482015290519081900360640190fd5b60005b8251811015610f36578181815181106111ec57fe5b602002602001015161ffff1683828151811061120457fe5b602002602001015161ffff161461125b576040805162461bcd60e51b815260206004820152601660248201527575696e7431365b5d3a20756e657175616c206974656d60501b604482015290519081900360640190fd5b6001016111d7565b600082601f830112611273578081fd5b813560206112886112838361195e565b61193a565b82815281810190858301838502870184018810156112a4578586fd5b855b858110156112d157813561ffff811681146112bf578788fd5b845292840192908401906001016112a6565b5090979650505050505050565b600082601f8301126112ee578081fd5b813560206112fe6112838361195e565b828152818101908583018385028701840188101561131a578586fd5b855b858110156112d15781358452928401929084019060010161131c565b600060a08284031215611349578081fd5b50919050565b600060208284031215611360578081fd5b81358015158114610dbd578182fd5b60008060008060808587031215611384578283fd5b843567ffffffffffffffff8082111561139b578485fd5b9086019060c082890312156113ae578485fd5b909450602086013590808211156113c3578485fd5b6113cf88838901611338565b945060408701359150808211156113e4578384fd5b506113f187828801611338565b949793965093946060013593505050565b6020808252601a908201527f6163746f72206d75737420696e6372656d656e7420726f756e64000000000000604082015260600190565b6020808252601d908201527f6163746f722063616e6e6f74206f766572726964652077656967687473000000604082015260600190565b6020808252601c908201527f6163746f722063616e6e6f7420696e6372656d656e7420726f756e6400000000604082015260600190565b6020808252601a908201527f6163746f722063616e6e6f74206f76657272696465206c6f7373000000000000604082015260600190565b60208082526031908201527f6163746f722063616e6e6f74206f76657272696465206c6f7373206f75747369604082015270032329031bab93932b73a103937bab7321607d1b606082015260800190565b602080825260139082015272726f756e64206f7574206f6620626f756e647360681b604082015260600190565b60208082526032908201527f6163746f722063616e6e6f74206f7665727269646520776569676874206f75746040820152711cda59194818dd5c9c995b9d081c9bdd5b9960721b606082015260800190565b60208082526014908201527314dd58905b1b1bd8ce881d5b995c5d585b08125160621b604082015260600190565b60208082526010908201526f30b1ba37b9103737ba1039b4b3b732b960811b604082015260600190565b6020808252601690820152756163746f722063616e6e6f7420736b6970206c6f737360501b604082015260600190565b6020808252600d908201526c1c9bdd5b990818da185b99d959609a1b604082015260600190565b6020808252600a90820152693732bc3a1030b1ba37b960b11b604082015260600190565b6020808252600a908201526966696e616c20666c616760b01b604082015260600190565b6020808252601b908201527f75696e743235365b5d5b5d3a20756e657175616c206c656e6774680000000000604082015260600190565b6020808252601690820152756e756d626572206f66207061727469636970616e747360501b604082015260600190565b6020808252600b908201526a0c8c2e8c240d8cadccee8d60ab1b604082015260600190565b60208082526018908201527f6163746f722063616e6e6f7420736b6970207765696768740000000000000000604082015260600190565b60208082526034908201527f6163746f722063616e6e6f74206f76657272696465206163637572616379206f6040820152731d5d1cda59194818dd5c9c995b9d081c9bdd5b9960621b606082015260800190565b6020808252601a908201527f6163746f722063616e6e6f7420736b6970206163637572616379000000000000604082015260600190565b6020808252601a908201527f537562416c6c6f635b5d3a20756e657175616c206c656e677468000000000000604082015260600190565b6020808252601e908201527f6163746f722063616e6e6f74206f766572726964652061636375726163790000604082015260600190565b6020808252600d908201526c1b5bd9195b0818da185b99d959609a1b604082015260600190565b6000808335601e1984360301811261189e578283fd5b83018035915067ffffffffffffffff8211156118b8578283fd5b60209081019250810236038213156118cf57600080fd5b9250929050565b6000808335601e198436030181126118ec578283fd5b83018035915067ffffffffffffffff821115611906578283fd5b6020019150368190038213156118cf57600080fd5b60008235605e19833603018112611930578182fd5b9190910192915050565b60405181810167ffffffffffffffff8111828210171561195657fe5b604052919050565b600067ffffffffffffffff82111561197257fe5b5060209081020190565b6000808585111561198b578182fd5b83861115611997578182fd5b5050820193919092039150565b60006119b26112838461195e565b8381526020808201919084845b878110156119e6576119d436833589016112de565b855293820193908201906001016119bf565b50919695505050505050565b6000611a006112838461195e565b8381526020808201919084845b878110156119e657813587016060808236031215611a29578788fd5b604080519182019167ffffffffffffffff8084118285101715611a4857fe5b92825283358152868401359280841115611a60578a8bfd5b611a6c368587016112de565b8883015282850135935080841115611a82578a8bfd5b50611a8f36848601611263565b91810191909152875250509382019390820190600101611a0d56fea2646970667358221220283ebfed06162f4d11a88a12b3435e16e202ff54a59e96b7c6f7b50ca716884864736f6c63430007060033",
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
