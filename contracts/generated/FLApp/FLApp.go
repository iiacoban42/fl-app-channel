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
	Bin: "0x608060405234801561001057600080fd5b5061193b806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c80630d1feb4f14610030575b600080fd5b61004361003e3660046111e3565b610045565b005b6002610054604086018661170b565b90501461007c5760405162461bcd60e51b8152600401610073906115ea565b60405180910390fd5b600061008b6060850185611759565b60008161009457fe5b919091013560f81c9150508181146100be5760405162461bcd60e51b8152600401610073906114ea565b6100cb6060840184611759565b6000816100d457fe5b9091013560f81c6001838101161490506101005760405162461bcd60e51b81526004016100739061156b565b61010d6060850185611759565b600481811061011857fe5b919091013560f81c159050610209576101346060840184611759565b600181811061013f57fe5b909101356001600160f81b031916905061015c6060860186611759565b600181811061016757fe5b9050013560f81c60f81b6001600160f81b031916146101985760405162461bcd60e51b8152600401610073906116e4565b6101a56060840184611759565b60028181106101b057fe5b909101356001600160f81b03191690506101cd6060860186611759565b60028181106101d857fe5b9050013560f81c60f81b6001600160f81b031916146102095760405162461bcd60e51b815260040161007390611544565b60006102186060850185611759565b600281811061022357fe5b919091013560f81c915050600580820190828001808201919084010160ff81166102506060890189611759565b90501461026f5760405162461bcd60e51b81526004016100739061161a565b61027c6060880188611759565b600281811061028757fe5b919091013560f81c905061029e60608a018a611759565b60028181106102a957fe5b919091013560f81c919091111590506102c56060890189611759565b60028181106102d057fe5b9050013560f81c60f81b6040516020016102ea9190611333565b604051602081830303815290604052906103175760405162461bcd60e51b81526004016100739190611432565b5060ff851661060c5761032d6060890189611759565b600481811061033857fe5b909101356001600160f81b03191615801591506103a45750600161035f60608a018a611759565b600381811061036a57fe5b60ff92013560f81c929092011690506103866060890189611759565b600381811061039157fe5b9050013560f81c60f81b60f81c60ff1614155b156103b260608a018a611759565b60038181106103bd57fe5b9091013560f81c60010190506103d660608a018a611759565b60038181106103e157fe5b9050013560f81c60f81b6040516020016103fc9291906113cf565b604051602081830303815290604052906104295760405162461bcd60e51b81526004016100739190611432565b506104376060880188611759565b60ff600587011681811061044757fe5b909101356001600160f81b031916905061046460608a018a611759565b60ff600588011681811061047457fe5b909101356001600160f81b03191691909114905061049560608a018a611759565b60ff60058801168181106104a557fe5b909101356001600160f81b03191690506104c260608a018a611759565b60ff60058901168181106104d257fe5b9050013560f81c60f81b6040516020016104ed929190611276565b6040516020818303038152906040529061051a5760405162461bcd60e51b81526004016100739190611432565b506105286060890189611759565b600481811061053357fe5b919091013560f81c15905061060c5761054f6060880188611759565b61055c60608b018b611759565b600381811061056757fe5b919091013560f81c860160ff16905081811061057f57fe5b919091013560f81c151590506105a75760405162461bcd60e51b815260040161007390611676565b6105b46060880188611759565b6105c160608b018b611759565b60038181106105cc57fe5b919091013560f81c850160ff1690508181106105e457fe5b919091013560f81c1515905061060c5760405162461bcd60e51b815260040161007390611514565b8460ff16600114156108ca576106256060880188611759565b600381811061063057fe5b909101356001600160f81b031916905061064d60608a018a611759565b600381811061065857fe5b9050013560f81c60f81b6001600160f81b031916146106895760405162461bcd60e51b815260040161007390611485565b6106966060880188611759565b85850160ff168181106106a557fe5b909101356001600160f81b03191690506106c260608a018a611759565b86860160ff168181106106d157fe5b909101356001600160f81b0319169190911490506106f260608a018a611759565b86860160ff1681811061070157fe5b909101356001600160f81b031916905061071e60608a018a611759565b87870160ff1681811061072d57fe5b9050013560f81c60f81b60405160200161074892919061136e565b604051602081830303815290604052906107755760405162461bcd60e51b81526004016100739190611432565b506107836060880188611759565b85840160ff1681811061079257fe5b909101356001600160f81b03191690506107af60608a018a611759565b86850160ff168181106107be57fe5b909101356001600160f81b0319169190911490506107df60608a018a611759565b86850160ff168181106107ee57fe5b909101356001600160f81b031916905061080b60608a018a611759565b87860160ff1681811061081a57fe5b9050013560f81c60f81b6040516020016108359291906112d6565b604051602081830303815290604052906108625760405162461bcd60e51b81526004016100739190611432565b506108706060880188611759565b61087d60608b018b611759565b600381811061088857fe5b600592013560f81c9190910160ff1690508181106108a257fe5b919091013560f81c151590506108ca5760405162461bcd60e51b81526004016100739061163f565b6000808061091a6108de60608c018c611759565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508a9250610c3b915050565b9194509250905082151561093460a08c0160808d016111bc565b1515146109535760405162461bcd60e51b81526004016100739061158f565b6109ed61096360408c018c61179e565b61096d908061170b565b808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152506109ac9250505060408e018e61179e565b6109b6908061170b565b80806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250610cf992505050565b610a3d6109fd60408c018c61179e565b610a0b90604081019061170b565b610a149161184d565b610a2160408e018e61179e565b610a2f90604081019061170b565b610a389161184d565b610df4565b6000610a4c60408d018d61179e565b610a5a90602081019061170b565b610a63916117ff565b90508215610bff57805160018390039067ffffffffffffffff81118015610a8957600080fd5b50604051908082528060200260200182016040528015610abd57816020015b6060815260200190600190039081610aa85790505b50915060005b8251811015610bfc57604080516002808252606082018352909160208301908036833701905050838281518110610af657fe5b6020908102919091010152610b0e60408f018f61179e565b610b1c90602081019061170b565b82818110610b2657fe5b9050602002810190610b38919061170b565b6001818110610b4357fe5b905060200201358e8060400190610b5a919061179e565b610b6890602081019061170b565b83818110610b7257fe5b9050602002810190610b84919061170b565b6000818110610b8f57fe5b9050602002013501838281518110610ba357fe5b60200260200101518560ff1681518110610bb957fe5b6020026020010181815250506000838281518110610bd357fe5b60200260200101518360ff1681518110610be957fe5b6020908102919091010152600101610ac3565b50505b610c2c610c0f60408d018d61179e565b610c1d90602081019061170b565b610c26916117ff565b82610e59565b50505050505050505050505050565b81516000908190819085906003908110610c5157fe5b016020015185516001600160f81b03199091169086906002908110610c7257fe5b01602001516001600160f81b031916148015610ca55750845185906002908110610c9857fe5b60209101015160f81c6003145b15610ce857603c60ff16858560ff1681518110610cbe57fe5b016020015160f81c10610cd957506001915081905080610cf2565b50600191508190506000610cf2565b5060009150819050805b9250925092565b8051825114610d4f576040805162461bcd60e51b815260206004820152601960248201527f616464726573735b5d3a20756e657175616c206c656e67746800000000000000604482015290519081900360640190fd5b60005b8251811015610def57818181518110610d6757fe5b60200260200101516001600160a01b0316838281518110610d8457fe5b60200260200101516001600160a01b031614610de7576040805162461bcd60e51b815260206004820152601760248201527f616464726573735b5d3a20756e657175616c206974656d000000000000000000604482015290519081900360640190fd5b600101610d52565b505050565b8051825114610e155760405162461bcd60e51b8152600401610073906116ad565b60005b8251811015610def57610e51838281518110610e3057fe5b6020026020010151838381518110610e4457fe5b6020026020010151610ebe565b600101610e18565b8051825114610e7a5760405162461bcd60e51b8152600401610073906115b3565b60005b8251811015610def57610eb6838281518110610e9557fe5b6020026020010151838381518110610ea957fe5b6020026020010151610f07565b600101610e7d565b8051825114610edf5760405162461bcd60e51b8152600401610073906114bc565b610ef182602001518260200151610f07565b610f0382604001518260400151610feb565b5050565b8051825114610f5d576040805162461bcd60e51b815260206004820152601960248201527f75696e743235365b5d3a20756e657175616c206c656e67746800000000000000604482015290519081900360640190fd5b60005b8251811015610def57818181518110610f7557fe5b6020026020010151838281518110610f8957fe5b602002602001015114610fe3576040805162461bcd60e51b815260206004820152601760248201527f75696e743235365b5d3a20756e657175616c206974656d000000000000000000604482015290519081900360640190fd5b600101610f60565b8051825114611041576040805162461bcd60e51b815260206004820152601860248201527f75696e7431365b5d3a20756e657175616c206c656e6774680000000000000000604482015290519081900360640190fd5b60005b8251811015610def5781818151811061105957fe5b602002602001015161ffff1683828151811061107157fe5b602002602001015161ffff16146110c8576040805162461bcd60e51b815260206004820152601660248201527575696e7431365b5d3a20756e657175616c206974656d60501b604482015290519081900360640190fd5b600101611044565b600082601f8301126110e0578081fd5b813560206110f56110f0836117e1565b6117bd565b8281528181019085830183850287018401881015611111578586fd5b855b8581101561113e57813561ffff8116811461112c578788fd5b84529284019290840190600101611113565b5090979650505050505050565b600082601f83011261115b578081fd5b8135602061116b6110f0836117e1565b8281528181019085830183850287018401881015611187578586fd5b855b8581101561113e57813584529284019290840190600101611189565b600060a082840312156111b6578081fd5b50919050565b6000602082840312156111cd578081fd5b813580151581146111dc578182fd5b9392505050565b600080600080608085870312156111f8578283fd5b843567ffffffffffffffff8082111561120f578485fd5b9086019060c08289031215611222578485fd5b90945060208601359080821115611237578485fd5b611243888389016111a5565b94506040870135915080821115611258578384fd5b50611265878288016111a5565b949793965093946060013593505050565b7f6163746f722063616e6e6f74206f7665727269646520776569676874733a20658152661e1c1958dd195960ca1b60208201526001600160f81b03199283166027820152650161033b7ba160d51b60288201529116602e820152602f0190565b7f6163746f722063616e6e6f74206f76657272696465206c6f73733a206578706581526318dd195960e21b60208201526001600160f81b03199283166024820152650161033b7ba160d51b60258201529116602b820152602c0190565b7f726f756e64496e646578206f7574206f6620626f756e64733a2000000000000081526001600160f81b031991909116601a820152601b0190565b7f6163746f722063616e6e6f74206f766572726964652061636375726163793a20815267195e1c1958dd195960c21b60208201526001600160f81b03199283166028820152650161033b7ba160d51b60298201529116602f82015260300190565b7f6163746f72206d75737420696e6372656d656e7420726f756e643a2065787065815264031ba32b2160dd1b602082015260f89290921b6001600160f81b03199081166025840152650161033b7ba160d51b602684015216602c820152602d0190565b6000602080835283518082850152825b8181101561145e57858101830151858201604001528201611442565b8181111561146f5783604083870101525b50601f01601f1916929092016040019392505050565b6020808252601c908201527f6163746f722063616e6e6f7420696e6372656d656e7420726f756e6400000000604082015260600190565b60208082526014908201527314dd58905b1b1bd8ce881d5b995c5d585b08125160621b604082015260600190565b60208082526010908201526f30b1ba37b9103737ba1039b4b3b732b960811b604082015260600190565b6020808252601690820152756163746f722063616e6e6f7420736b6970206c6f737360501b604082015260600190565b6020808252600d908201526c1c9bdd5b990818da185b99d959609a1b604082015260600190565b6020808252600a90820152693732bc3a1030b1ba37b960b11b604082015260600190565b6020808252600a908201526966696e616c20666c616760b01b604082015260600190565b6020808252601b908201527f75696e743235365b5d5b5d3a20756e657175616c206c656e6774680000000000604082015260600190565b6020808252601690820152756e756d626572206f66207061727469636970616e747360501b604082015260600190565b6020808252600b908201526a0c8c2e8c240d8cadccee8d60ab1b604082015260600190565b60208082526018908201527f6163746f722063616e6e6f7420736b6970207765696768740000000000000000604082015260600190565b6020808252601a908201527f6163746f722063616e6e6f7420736b6970206163637572616379000000000000604082015260600190565b6020808252601a908201527f537562416c6c6f635b5d3a20756e657175616c206c656e677468000000000000604082015260600190565b6020808252600d908201526c1b5bd9195b0818da185b99d959609a1b604082015260600190565b6000808335601e19843603018112611721578283fd5b83018035915067ffffffffffffffff82111561173b578283fd5b602090810192508102360382131561175257600080fd5b9250929050565b6000808335601e1984360301811261176f578283fd5b83018035915067ffffffffffffffff821115611789578283fd5b60200191503681900382131561175257600080fd5b60008235605e198336030181126117b3578182fd5b9190910192915050565b60405181810167ffffffffffffffff811182821017156117d957fe5b604052919050565b600067ffffffffffffffff8211156117f557fe5b5060209081020190565b600061180d6110f0846117e1565b8381526020808201919084845b878110156118415761182f368335890161114b565b8552938201939082019060010161181a565b50919695505050505050565b600061185b6110f0846117e1565b8381526020808201919084845b8781101561184157813587016060808236031215611884578788fd5b604080519182019167ffffffffffffffff80841182851017156118a357fe5b928252833581528684013592808411156118bb578a8bfd5b6118c73685870161114b565b88830152828501359350808411156118dd578a8bfd5b506118ea368486016110d0565b9181019190915287525050938201939082019060010161186856fea264697066735822122039d693368a110c369c88dda6e398ed62cdd8e7c6a2f0b0a312c99c85cde2e69164736f6c63430007060033",
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
