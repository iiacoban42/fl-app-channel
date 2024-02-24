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
	Bin: "0x608060405234801561001057600080fd5b50611aa7806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c80630d1feb4f14610030575b600080fd5b61004361003e36600461135a565b610045565b005b6002610054604086018661184f565b90501461007c5760405162461bcd60e51b815260040161007390611697565b60405180910390fd5b600061008b606085018561189d565b60008161009457fe5b919091013560f81c9150508181146100be5760405162461bcd60e51b8152600401610073906115c7565b6100cb606084018461189d565b6000816100d457fe5b9091013560f81c6001838101161490506101005760405162461bcd60e51b815260040161007390611618565b61010d606085018561189d565b600481811061011857fe5b919091013560f81c15905061020957610134606084018461189d565b600181811061013f57fe5b909101356001600160f81b031916905061015c606086018661189d565b600181811061016757fe5b9050013560f81c60f81b6001600160f81b031916146101985760405162461bcd60e51b8152600401610073906117e5565b6101a5606084018461189d565b60028181106101b057fe5b909101356001600160f81b03191690506101cd606086018661189d565b60028181106101d857fe5b9050013560f81c60f81b6001600160f81b031916146102095760405162461bcd60e51b8152600401610073906115f1565b6000610218606085018561189d565b600281811061022357fe5b919091013560f81c915050600580820190828001808201919084010160ff8116610250606089018961189d565b90501461026f5760405162461bcd60e51b8152600401610073906116c7565b61027c606088018861189d565b600281811061028757fe5b919091013560f81c905061029e606089018961189d565b60038181106102a957fe5b9050013560f81c60f81b60f81c60ff1611156102d75760405162461bcd60e51b81526004016100739061151a565b60ff8516610685576103936102ef60608a018a61189d565b61030391600588810160ff16929091611943565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506103459250505060608a018a61189d565b61035991600589810160ff16929091611943565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610ca292505050565b6103af5760405162461bcd60e51b815260040161007390611424565b6104876103bf60608a018a61189d565b6103d49160ff87890181169290881691611943565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506104169250505060608a018a61189d565b61042b9160ff888a0181169290891691611943565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061046d9250505060608c018c61189d565b600381811061047857fe5b919091013560f81c9050610d20565b6104a35760405162461bcd60e51b815260040161007390611723565b61051f6104b360608a018a61189d565b6104c89160ff86890181169290871691611943565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061050a9250505060608a018a61189d565b61042b9160ff878a0181169290881691611943565b61053b5760405162461bcd60e51b8152600401610073906114c9565b610548606089018961189d565b600481811061055357fe5b919091013560f81c15905061068557600161057160608a018a61189d565b600381811061057c57fe5b60ff92013560f81c92909201169050610598606089018961189d565b60038181106105a357fe5b9050013560f81c60f81b60f81c60ff16146105d05760405162461bcd60e51b8152600401610073906113ed565b6105dd606088018861189d565b6105ea60608b018b61189d565b60038181106105f557fe5b919091013560f81c860160ff16905081811061060d57fe5b919091013560f81c1515905080610669575061062c606088018861189d565b61063960608b018b61189d565b600381811061064457fe5b919091013560f81c850160ff16905081811061065c57fe5b919091013560f81c151590505b6106855760405162461bcd60e51b81526004016100739061180c565b8460ff16600114156109675761069e606088018861189d565b60038181106106a957fe5b919091013560f81c90506106c060608a018a61189d565b60038181106106cb57fe5b9050013560f81c60f81b60f81c60ff16146106f85760405162461bcd60e51b81526004016100739061145b565b61077461070860608a018a61189d565b61071d9160ff87890181169290881691611943565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061075f9250505060608a018a61189d565b6103599160ff888a0181169290891691611943565b6107905760405162461bcd60e51b8152600401610073906117ae565b61080c6107a060608a018a61189d565b6107b59160ff86890181169290871691611943565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506107f79250505060608a018a61189d565b6103599160ff878a0181169290881691611943565b6108285760405162461bcd60e51b815260040161007390611492565b6108e461083860608a018a61189d565b61084c91600588810160ff16929091611943565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061088e9250505060608a018a61189d565b6108a291600589810160ff16929091611943565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061046d9250505060608b018b61189d565b6109005760405162461bcd60e51b815260040161007390611547565b61090d606088018861189d565b61091a60608b018b61189d565b600381811061092557fe5b600592013560f81c9190910160ff16905081811061093f57fe5b919091013560f81c151590506109675760405162461bcd60e51b8152600401610073906116ec565b60006109b361097960608a018a61189d565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610daf92505050565b90508015156109c860a08a0160808b0161133a565b1515146109e75760405162461bcd60e51b81526004016100739061163c565b610a816109f760408a018a6118e2565b610a01908061184f565b80806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250610a409250505060408c018c6118e2565b610a4a908061184f565b80806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250610e2b92505050565b610ad1610a9160408a018a6118e2565b610a9f90604081019061184f565b610aa8916119b9565b610ab560408c018c6118e2565b610ac390604081019061184f565b610acc916119b9565b610f26565b6000610ae060408b018b6118e2565b610aee90602081019061184f565b610af79161196b565b9050610b43610b0960608b018b61189d565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610f8b92505050565b15610c6857670de0b6b3a764000060016000805b8451811015610c6357838e8060400190610b7191906118e2565b610b7f90602081019061184f565b83818110610b8957fe5b9050602002810190610b9b919061184f565b8560ff16818110610ba857fe5b9050602002013501858281518110610bbc57fe5b60200260200101518460ff1681518110610bd257fe5b602002602001018181525050838e8060400190610bef91906118e2565b610bfd90602081019061184f565b83818110610c0757fe5b9050602002810190610c19919061184f565b8460ff16818110610c2657fe5b9050602002013503858281518110610c3a57fe5b60200260200101518360ff1681518110610c5057fe5b6020908102919091010152600101610b57565b505050505b610c95610c7860408b018b6118e2565b610c8690602081019061184f565b610c8f9161196b565b82610fd7565b5050505050505050505050565b60008151835114610cb557506000610d1a565b60005b8351811015610d1457828181518110610ccd57fe5b602001015160f81c60f81b6001600160f81b031916848281518110610cee57fe5b01602001516001600160f81b03191614610d0c576000915050610d1a565b600101610cb8565b50600190505b92915050565b60008251845114610d3357506000610da8565b60005b8451811015610da2578260ff16811415610d4f57610d9a565b838181518110610d5b57fe5b602001015160f81c60f81b6001600160f81b031916858281518110610d7c57fe5b01602001516001600160f81b03191614610d9a576000915050610da8565b600101610d36565b50600190505b9392505050565b805160009082906003908110610dc157fe5b016020015182516001600160f81b03199091169083906002908110610de257fe5b01602001516001600160f81b031916148015610e155750815182906004908110610e0857fe5b60209101015160f81c6003145b15610e2257506001610e26565b5060005b919050565b8051825114610e81576040805162461bcd60e51b815260206004820152601960248201527f616464726573735b5d3a20756e657175616c206c656e67746800000000000000604482015290519081900360640190fd5b60005b8251811015610f2157818181518110610e9957fe5b60200260200101516001600160a01b0316838281518110610eb657fe5b60200260200101516001600160a01b031614610f19576040805162461bcd60e51b815260206004820152601760248201527f616464726573735b5d3a20756e657175616c206974656d000000000000000000604482015290519081900360640190fd5b600101610e84565b505050565b8051825114610f475760405162461bcd60e51b815260040161007390611777565b60005b8251811015610f2157610f83838281518110610f6257fe5b6020026020010151838381518110610f7657fe5b602002602001015161103c565b600101610f4a565b805160009060029083906004908110610fa057fe5b016020015160f81c148015610e155750815160009083906005908110610fc257fe5b016020015160f81c14610e2257506001610e26565b8051825114610ff85760405162461bcd60e51b815260040161007390611660565b60005b8251811015610f215761103483828151811061101357fe5b602002602001015183838151811061102757fe5b6020026020010151611085565b600101610ffb565b805182511461105d5760405162461bcd60e51b815260040161007390611599565b61106f82602001518260200151611085565b61108182604001518260400151611169565b5050565b80518251146110db576040805162461bcd60e51b815260206004820152601960248201527f75696e743235365b5d3a20756e657175616c206c656e67746800000000000000604482015290519081900360640190fd5b60005b8251811015610f21578181815181106110f357fe5b602002602001015183828151811061110757fe5b602002602001015114611161576040805162461bcd60e51b815260206004820152601760248201527f75696e743235365b5d3a20756e657175616c206974656d000000000000000000604482015290519081900360640190fd5b6001016110de565b80518251146111bf576040805162461bcd60e51b815260206004820152601860248201527f75696e7431365b5d3a20756e657175616c206c656e6774680000000000000000604482015290519081900360640190fd5b60005b8251811015610f21578181815181106111d757fe5b602002602001015161ffff168382815181106111ef57fe5b602002602001015161ffff1614611246576040805162461bcd60e51b815260206004820152601660248201527575696e7431365b5d3a20756e657175616c206974656d60501b604482015290519081900360640190fd5b6001016111c2565b600082601f83011261125e578081fd5b8135602061127361126e83611925565b611901565b828152818101908583018385028701840188101561128f578586fd5b855b858110156112bc57813561ffff811681146112aa578788fd5b84529284019290840190600101611291565b5090979650505050505050565b600082601f8301126112d9578081fd5b813560206112e961126e83611925565b8281528181019085830183850287018401881015611305578586fd5b855b858110156112bc57813584529284019290840190600101611307565b600060a08284031215611334578081fd5b50919050565b60006020828403121561134b578081fd5b81358015158114610da8578182fd5b6000806000806080858703121561136f578283fd5b843567ffffffffffffffff80821115611386578485fd5b9086019060c08289031215611399578485fd5b909450602086013590808211156113ae578485fd5b6113ba88838901611323565b945060408701359150808211156113cf578384fd5b506113dc87828801611323565b949793965093946060013593505050565b6020808252601a908201527f6163746f72206d75737420696e6372656d656e7420726f756e64000000000000604082015260600190565b6020808252601d908201527f6163746f722063616e6e6f74206f766572726964652077656967687473000000604082015260600190565b6020808252601c908201527f6163746f722063616e6e6f7420696e6372656d656e7420726f756e6400000000604082015260600190565b6020808252601a908201527f6163746f722063616e6e6f74206f76657272696465206c6f7373000000000000604082015260600190565b60208082526031908201527f6163746f722063616e6e6f74206f76657272696465206c6f7373206f75747369604082015270032329031bab93932b73a103937bab7321607d1b606082015260800190565b602080825260139082015272726f756e64206f7574206f6620626f756e647360681b604082015260600190565b60208082526032908201527f6163746f722063616e6e6f74206f7665727269646520776569676874206f75746040820152711cda59194818dd5c9c995b9d081c9bdd5b9960721b606082015260800190565b60208082526014908201527314dd58905b1b1bd8ce881d5b995c5d585b08125160621b604082015260600190565b60208082526010908201526f30b1ba37b9103737ba1039b4b3b732b960811b604082015260600190565b6020808252600d908201526c1c9bdd5b990818da185b99d959609a1b604082015260600190565b6020808252600a90820152693732bc3a1030b1ba37b960b11b604082015260600190565b6020808252600a908201526966696e616c20666c616760b01b604082015260600190565b6020808252601b908201527f75696e743235365b5d5b5d3a20756e657175616c206c656e6774680000000000604082015260600190565b6020808252601690820152756e756d626572206f66207061727469636970616e747360501b604082015260600190565b6020808252600b908201526a0c8c2e8c240d8cadccee8d60ab1b604082015260600190565b60208082526018908201527f6163746f722063616e6e6f7420736b6970207765696768740000000000000000604082015260600190565b60208082526034908201527f6163746f722063616e6e6f74206f76657272696465206163637572616379206f6040820152731d5d1cda59194818dd5c9c995b9d081c9bdd5b9960621b606082015260800190565b6020808252601a908201527f537562416c6c6f635b5d3a20756e657175616c206c656e677468000000000000604082015260600190565b6020808252601e908201527f6163746f722063616e6e6f74206f766572726964652061636375726163790000604082015260600190565b6020808252600d908201526c1b5bd9195b0818da185b99d959609a1b604082015260600190565b60208082526023908201527f6163746f722063616e6e6f7420736b697020616363757261637920616e64206c6040820152626f737360e81b606082015260800190565b6000808335601e19843603018112611865578283fd5b83018035915067ffffffffffffffff82111561187f578283fd5b602090810192508102360382131561189657600080fd5b9250929050565b6000808335601e198436030181126118b3578283fd5b83018035915067ffffffffffffffff8211156118cd578283fd5b60200191503681900382131561189657600080fd5b60008235605e198336030181126118f7578182fd5b9190910192915050565b60405181810167ffffffffffffffff8111828210171561191d57fe5b604052919050565b600067ffffffffffffffff82111561193957fe5b5060209081020190565b60008085851115611952578182fd5b8386111561195e578182fd5b5050820193919092039150565b600061197961126e84611925565b8381526020808201919084845b878110156119ad5761199b36833589016112c9565b85529382019390820190600101611986565b50919695505050505050565b60006119c761126e84611925565b8381526020808201919084845b878110156119ad578135870160608082360312156119f0578788fd5b604080519182019167ffffffffffffffff8084118285101715611a0f57fe5b92825283358152868401359280841115611a27578a8bfd5b611a33368587016112c9565b8883015282850135935080841115611a49578a8bfd5b50611a563684860161124e565b918101919091528752505093820193908201906001016119d456fea264697066735822122079f9769056fe7034519d3e9d6a46a9e886246841c6a25a066bafe27806686ca564736f6c63430007060033",
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
