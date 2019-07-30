// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gost

import (
	"math/big"
	"strings"

	"github.com/gochain-io/gochain/v3"
	"github.com/gochain-io/gochain/v3/accounts/abi"
	"github.com/gochain-io/gochain/v3/accounts/abi/bind"
	"github.com/gochain-io/gochain/v3/common"
	"github.com/gochain-io/gochain/v3/core/types"
	"github.com/gochain-io/gochain/v3/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = gochain.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ConfirmationsABI is the input ABI used to generate the binding from.
const ConfirmationsABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"_confirmedGas\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"blockNum\",\"type\":\"uint256\"},{\"name\":\"logIndex\",\"type\":\"uint256\"},{\"name\":\"eventHash\",\"type\":\"bytes32\"}],\"name\":\"request\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"confirmGas\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"voter\",\"type\":\"bool\"},{\"name\":\"add\",\"type\":\"bool\"}],\"name\":\"setVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalConfirmGas\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"bool\"}],\"name\":\"count\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"blockNum\",\"type\":\"uint256\"},{\"name\":\"logIndex\",\"type\":\"uint256\"},{\"name\":\"eventHash\",\"type\":\"bytes32\"},{\"name\":\"valid\",\"type\":\"bool\"}],\"name\":\"confirm\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"votes\",\"outputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"voter\",\"type\":\"bool\"},{\"name\":\"add\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"status\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"voter\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"blockNum\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"logIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"eventHash\",\"type\":\"bytes32\"}],\"name\":\"ConfirmationRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"blockNum\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"logIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"eventHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"valid\",\"type\":\"bool\"}],\"name\":\"Confirmed\",\"type\":\"event\"}]"

// ConfirmationsBin is the compiled bytecode used for deploying new contracts.
const ConfirmationsBin = `0x608060405234801561001057600080fd5b506040516020806116a68339810180604052602081101561003057600080fd5b505180610049600082610065602090811b610d3517901c565b61005e60028261006560201b610d351760201c565b505061012e565b6001600160a01b0381166000908152600183016020526040902054156100ec57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600e60248201527f416c726561647920696e20736574000000000000000000000000000000000000604482015290519081900360640190fd5b815460018181018085556000858152602080822090940180546001600160a01b039096166001600160a01b031990961686179055938452930190526040902055565b6115698061013d6000396000f3fe6080604052600436106100865760003560e01c80635c5f6b39116100595780635c5f6b39146101355780637849ed5a1461014a5780638f7fd0d61461018d578063d8bff5a5146101cb578063fad3ffd61461022857610086565b80632ea6ece81461008b57806332030803146100b257806342715aec146100dd57806348c5000d146100f2575b600080fd5b34801561009757600080fd5b506100a0610282565b60408051918252519081900360200190f35b6100db600480360360608110156100c857600080fd5b5080359060208101359060400135610289565b005b3480156100e957600080fd5b506100a06104de565b3480156100fe57600080fd5b506100db6004803603606081101561011557600080fd5b506001600160a01b038135169060208101351515906040013515156104e5565b34801561014157600080fd5b506100a0610749565b34801561015657600080fd5b506100a06004803603606081101561016d57600080fd5b506001600160a01b03813516906020810135151590604001351515610750565b34801561019957600080fd5b506100db600480360360808110156101b057600080fd5b50803590602081013590604081013590606001351515610773565b3480156101d757600080fd5b506101fe600480360360208110156101ee57600080fd5b50356001600160a01b0316610990565b604080516001600160a01b0390941684529115156020840152151582820152519081900360600190f35b34801561023457600080fd5b5061025e6004803603606081101561024b57600080fd5b50803590602081013590604001356109c1565b6040518082600381111561026e57fe5b60ff16815260200191505060405180910390f35b6202710081565b6000838152600660209081526040808320858452825280832084845290915281205460ff1660038111156102b957fe5b1461030e5760408051600160e51b62461bcd02815260206004820152601360248201527f537461747573206d757374206265204e6f6e6500000000000000000000000000604482015290519081900360640190fd5b600083815260066020908152604080832085845282528083208484529091529020805460ff191660011790553a622625a0023481111561038257604051600160e51b62461bcd0281526004018080602001828103825260278152602001806115176027913960400191505060405180910390fd5b803411156103bb5760405133903483900380156108fc02916000818181858888f193505050501580156103b9573d6000803e3d6000fd5b505b6000848152600760209081526040808320868452825280832085845282528083203a6004909101558051606081018252878152808301878152818301878152600880546001810180835591885293517ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee360039095029485015591517ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee4840155517ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee59092019190915587845260098352818420878552835281842086855283529281902083905580518581529051869288927f2382b1fef9adb7c2a096ef31e2e402e08d3b5f9eb5fc18ea9a56835e44627b03929081900390910190a35050505050565b62015f9081565b336000908152600160205260409020546105495760408051600160e51b62461bcd02815260206004820152601260248201527f53656e646572206e6f74206120766f7465720000000000000000000000000000604482015290519081900360640190fd5b33600090815260046020526040902080546001600160a01b0316156105e55780546001600160a01b03808216600090815260056020908152604080832060ff600160a01b8704811615158552908352818420600160a81b909604161515835293905291909120805460001901905584166105e0575033600090815260046020526040902080546001600160b01b0319169055610744565b6105f9565b6001600160a01b0384166105f95750610744565b604080516060810182526001600160a01b03808716808352861515602080850182815288151586880181815233600090815260048086528a822099518a54955193511515600160a81b02600160a81b60ff0219941515600160a01b0274ff00000000000000000000000000000000000000001992909b166001600160a01b03199097169690961716989098179190911692909217909655928352600581528583209183529081528482209382529283528381208054600101908190558451600160e21b633881de0f02815292830191909152925173__$1c537defa8d8abb880144579cecc12eeda$__9263e207783c9260248082019391829003018186803b15801561070457600080fd5b505af4158015610718573d6000803e3d6000fd5b505050506040513d602081101561072e57600080fd5b50518110610741576107418585856109e7565b50505b505050565b622625a081565b600560209081526000938452604080852082529284528284209052825290205481565b336000908152600360205260409020546107d75760408051600160e51b62461bcd02815260206004820152601360248201527f53656e646572206e6f742061207369676e657200000000000000000000000000604482015290519081900360640190fd5b60016000858152600660209081526040808320878452825280832086845290915290205460ff16600381111561080957fe5b146108135761098a565b60008481526007602090815260408083208684528252808320858452909152902060048101543a1461087957604051600160e51b62461bcd0281526004018080602001828103825260228152602001806114f56022913960400191505060405180910390fd5b8082156108a957336000908152600383016020526040902054156108a4576108a48260020133610c16565b6108d0565b503360009081526001820160205260409020546002820190156108d0576108d08233610c16565b336000908152600182016020526040902054610987576108f08133610d35565b73__$1c537defa8d8abb880144579cecc12eeda$__63e207783c60026040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b15801561094057600080fd5b505af4158015610954573d6000803e3d6000fd5b505050506040513d602081101561096a57600080fd5b50518154101561097b57505061098a565b61098786868686610de7565b50505b50505050565b6004602052600090815260409020546001600160a01b0381169060ff600160a01b8204811691600160a81b90041683565b600660209081526000938452604080852082529284528284209052825290205460ff1681565b60005b600054811015610aa4576000806000018281548110610a0557fe5b60009182526020808320909101546001600160a01b0390811680845260049092526040909220805491935091908116908716148015610a535750805460ff600160a01b909104161515851515145b8015610a6e5750805460ff600160a81b909104161515841515145b15610a9a576001600160a01b038216600090815260046020526040902080546001600160b01b03191690555b50506001016109ea565b506001600160a01b038316600090815260056020908152604080832085158015855290835281842085151585529092528220919091558290610ae35750805b15610b3d576001600160a01b038316600090815260036020526040902054610b1057610b10600284610d35565b6001600160a01b038316600090815260016020526040902054610b3857610b38600084610d35565b610744565b818015610b48575080155b15610b74576001600160a01b03831660009081526001602052604090205415610b3857610b38836110c5565b81158015610b7f5750805b15610bac576001600160a01b038316600090815260036020526040902054610b3857610b38600284610d35565b81158015610bb8575080155b15610744576001600160a01b03831660009081526001602052604090205415610be457610be4836110c5565b6001600160a01b0383166000908152600360205260409020541561074457610c0d600284610c16565b610744836112b6565b6001600160a01b038116600090815260018301602052604090205480610c765760408051600160e51b62461bcd02815260206004820152600a6024820152600160b21b69139bdd081a5b881cd95d02604482015290519081900360640190fd5b6001600160a01b03821660009081526001840160205260408120558254836000198201828110610ca257fe5b60009182526020909120015484546001600160a01b039091169085906000198501908110610ccc57fe5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550836000016001820381548110610d0c57fe5b600091825260209091200180546001600160a01b0319169055835461074185600019830161143e565b6001600160a01b038116600090815260018301602052604090205415610da55760408051600160e51b62461bcd02815260206004820152600e60248201527f416c726561647920696e20736574000000000000000000000000000000000000604482015290519081900360640190fd5b815460018181018085556000858152602080822090940180546001600160a01b039096166001600160a01b031990961686179055938452930190526040902055565b60008115610df757506003610dfb565b5060025b600085815260066020908152604080832087845282528083208684529091529020805482919060ff19166001836003811115610e3357fe5b02179055506000858152600760209081526040808320878452825280832086845290915290208083610e655750600281015b805460048301546223b4a0026000828281610e7c57fe5b04905060005b83811015610ee8576000856000018281548110610e9b57fe5b60009182526020822001546040516001600160a01b039091169250829185156108fc02918691818181858888f19350505050158015610ede573d6000803e3d6000fd5b5050600101610e82565b506004850154604051848302840391620271000290339082840180156108fc02916000818181858888f19350505050158015610f28573d6000803e3d6000fd5b5060008c81526007602090815260408083208e845282528083208d84529091528120908181610f578282611462565b5050600282016000610f698282611462565b50506000600492909201829055508c81526009602090815260408083208e845282528083208d845290915281208054919055600854600181111561102f57600060086001830381548110610fb957fe5b906000526020600020906003020190508060086001850381548110610fda57fe5b6000918252602080832084546003909302019182556001808501548184015560029485015492850192909255845483526009815260408084209286015484529181528183209490930154825292909152208290555b6008600182038154811061103f57fe5b6000918252602082206003909102018181556001810182905560020155600880549061106f906000198301611480565b508c8e7fdf1051063b9bab79d715a3919f387eb9ee4291be1e9241dffb4500694141f25c8e8e60405180838152602001821515151581526020019250505060405180910390a35050505050505050505050505050565b6110d0600082610c16565b6001600160a01b03808216600090815260046020526040902080549091161561115b5780546001600160a01b03808216600090815260056020908152604080832060ff600160a01b8704811615158552908352818420600160a81b90960416151583529381528382208054600019019055918516815260049091522080546001600160b01b03191690555b600073__$1c537defa8d8abb880144579cecc12eeda$__63e207783c60006040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b1580156111ad57600080fd5b505af41580156111c1573d6000803e3d6000fd5b505050506040513d60208110156111d757600080fd5b5051905060005b60005481101561098a57600460008060000183815481106111fb57fe5b60009182526020808320909101546001600160a01b0390811684529083019390935260409091019020805490945016156112ab5782546001600160a01b038116600090815260056020908152604080832060ff600160a01b8604811615158552908352818420600160a81b9095041615158352929052205482116112ab5782546112a3906001600160a01b0381169060ff600160a01b8204811691600160a81b9004166109e7565b5050506112b3565b6001016111de565b50565b600073__$1c537defa8d8abb880144579cecc12eeda$__63e207783c60026040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b15801561130857600080fd5b505af415801561131c573d6000803e3d6000fd5b505050506040513d602081101561133257600080fd5b5051905060005b6008548110156107445760006008828154811061135257fe5b6000918252602080832060039092029091018054835260078252604080842060018084015486529084528185206002840154865284528185206001600160a01b038a1686529081019093529092205491925090156113b9576113b48186610c16565b6113e6565b6001600160a01b0385166000908152600382016020526040902054156113e6576113e68160020186610c16565b8054841161140c576114078260000154836001015484600201546001610de7565b611437565b60028101548411611430576114078260000154836001015484600201546000610de7565b6001909201915b5050611339565b815481835581811115610744576000838152602090206107449181019083016114ac565b50805460008255906000526020600020908101906112b391906114ac565b8154818355818111156107445760030281600302836000526020600020918201910161074491906114cd565b6114ca91905b808211156114c657600081556001016114b2565b5090565b90565b6114ca91905b808211156114c65760008082556001820181905560028201556003016114d356fe47617320707269636520646f65736e2774206d61746368207265717565737465642e54782076616c756520646f65736e277420636f76657220636f6e6669726d6174696f6e20666565a165627a7a723058207b35dcef7f25604c0b1a1f4281d31faea428f3ce9a8f58bd06c2421a7203a7a00029`

// DeployConfirmations deploys a new GoChain contract, binding an instance of Confirmations to it.
func DeployConfirmations(auth *bind.TransactOpts, backend bind.ContractBackend, voter common.Address) (common.Address, *types.Transaction, *Confirmations, error) {
	parsed, err := abi.JSON(strings.NewReader(ConfirmationsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ConfirmationsBin), backend, voter)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Confirmations{ConfirmationsCaller: ConfirmationsCaller{contract: contract}, ConfirmationsTransactor: ConfirmationsTransactor{contract: contract}, ConfirmationsFilterer: ConfirmationsFilterer{contract: contract}}, nil
}

// Confirmations is an auto generated Go binding around an GoChain contract.
type Confirmations struct {
	ConfirmationsCaller     // Read-only binding to the contract
	ConfirmationsTransactor // Write-only binding to the contract
	ConfirmationsFilterer   // Log filterer for contract events
}

// ConfirmationsCaller is an auto generated read-only Go binding around an GoChain contract.
type ConfirmationsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConfirmationsTransactor is an auto generated write-only Go binding around an GoChain contract.
type ConfirmationsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConfirmationsFilterer is an auto generated log filtering Go binding around an GoChain contract events.
type ConfirmationsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConfirmationsSession is an auto generated Go binding around an GoChain contract,
// with pre-set call and transact options.
type ConfirmationsSession struct {
	Contract     *Confirmations    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ConfirmationsCallerSession is an auto generated read-only Go binding around an GoChain contract,
// with pre-set call options.
type ConfirmationsCallerSession struct {
	Contract *ConfirmationsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ConfirmationsTransactorSession is an auto generated write-only Go binding around an GoChain contract,
// with pre-set transact options.
type ConfirmationsTransactorSession struct {
	Contract     *ConfirmationsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ConfirmationsRaw is an auto generated low-level Go binding around an GoChain contract.
type ConfirmationsRaw struct {
	Contract *Confirmations // Generic contract binding to access the raw methods on
}

// ConfirmationsCallerRaw is an auto generated low-level read-only Go binding around an GoChain contract.
type ConfirmationsCallerRaw struct {
	Contract *ConfirmationsCaller // Generic read-only contract binding to access the raw methods on
}

// ConfirmationsTransactorRaw is an auto generated low-level write-only Go binding around an GoChain contract.
type ConfirmationsTransactorRaw struct {
	Contract *ConfirmationsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewConfirmations creates a new instance of Confirmations, bound to a specific deployed contract.
func NewConfirmations(address common.Address, backend bind.ContractBackend) (*Confirmations, error) {
	contract, err := bindConfirmations(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Confirmations{ConfirmationsCaller: ConfirmationsCaller{contract: contract}, ConfirmationsTransactor: ConfirmationsTransactor{contract: contract}, ConfirmationsFilterer: ConfirmationsFilterer{contract: contract}}, nil
}

// NewConfirmationsCaller creates a new read-only instance of Confirmations, bound to a specific deployed contract.
func NewConfirmationsCaller(address common.Address, caller bind.ContractCaller) (*ConfirmationsCaller, error) {
	contract, err := bindConfirmations(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConfirmationsCaller{contract: contract}, nil
}

// NewConfirmationsTransactor creates a new write-only instance of Confirmations, bound to a specific deployed contract.
func NewConfirmationsTransactor(address common.Address, transactor bind.ContractTransactor) (*ConfirmationsTransactor, error) {
	contract, err := bindConfirmations(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConfirmationsTransactor{contract: contract}, nil
}

// NewConfirmationsFilterer creates a new log filterer instance of Confirmations, bound to a specific deployed contract.
func NewConfirmationsFilterer(address common.Address, filterer bind.ContractFilterer) (*ConfirmationsFilterer, error) {
	contract, err := bindConfirmations(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConfirmationsFilterer{contract: contract}, nil
}

// bindConfirmations binds a generic wrapper to an already deployed contract.
func bindConfirmations(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ConfirmationsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Confirmations *ConfirmationsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Confirmations.Contract.ConfirmationsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Confirmations *ConfirmationsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Confirmations.Contract.ConfirmationsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Confirmations *ConfirmationsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Confirmations.Contract.ConfirmationsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Confirmations *ConfirmationsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Confirmations.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Confirmations *ConfirmationsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Confirmations.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Confirmations *ConfirmationsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Confirmations.Contract.contract.Transact(opts, method, params...)
}

// ConfirmedGas is a free data retrieval call binding the contract method 0x2ea6ece8.
//
// Solidity: function _confirmedGas() constant returns(uint256)
func (_Confirmations *ConfirmationsCaller) ConfirmedGas(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Confirmations.contract.Call(opts, out, "_confirmedGas")
	return *ret0, err
}

// ConfirmedGas is a free data retrieval call binding the contract method 0x2ea6ece8.
//
// Solidity: function _confirmedGas() constant returns(uint256)
func (_Confirmations *ConfirmationsSession) ConfirmedGas() (*big.Int, error) {
	return _Confirmations.Contract.ConfirmedGas(&_Confirmations.CallOpts)
}

// ConfirmedGas is a free data retrieval call binding the contract method 0x2ea6ece8.
//
// Solidity: function _confirmedGas() constant returns(uint256)
func (_Confirmations *ConfirmationsCallerSession) ConfirmedGas() (*big.Int, error) {
	return _Confirmations.Contract.ConfirmedGas(&_Confirmations.CallOpts)
}

// ConfirmGas is a free data retrieval call binding the contract method 0x42715aec.
//
// Solidity: function confirmGas() constant returns(uint256)
func (_Confirmations *ConfirmationsCaller) ConfirmGas(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Confirmations.contract.Call(opts, out, "confirmGas")
	return *ret0, err
}

// ConfirmGas is a free data retrieval call binding the contract method 0x42715aec.
//
// Solidity: function confirmGas() constant returns(uint256)
func (_Confirmations *ConfirmationsSession) ConfirmGas() (*big.Int, error) {
	return _Confirmations.Contract.ConfirmGas(&_Confirmations.CallOpts)
}

// ConfirmGas is a free data retrieval call binding the contract method 0x42715aec.
//
// Solidity: function confirmGas() constant returns(uint256)
func (_Confirmations *ConfirmationsCallerSession) ConfirmGas() (*big.Int, error) {
	return _Confirmations.Contract.ConfirmGas(&_Confirmations.CallOpts)
}

// Count is a free data retrieval call binding the contract method 0x7849ed5a.
//
// Solidity: function count(address , bool , bool ) constant returns(uint256)
func (_Confirmations *ConfirmationsCaller) Count(opts *bind.CallOpts, arg0 common.Address, arg1 bool, arg2 bool) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Confirmations.contract.Call(opts, out, "count", arg0, arg1, arg2)
	return *ret0, err
}

// Count is a free data retrieval call binding the contract method 0x7849ed5a.
//
// Solidity: function count(address , bool , bool ) constant returns(uint256)
func (_Confirmations *ConfirmationsSession) Count(arg0 common.Address, arg1 bool, arg2 bool) (*big.Int, error) {
	return _Confirmations.Contract.Count(&_Confirmations.CallOpts, arg0, arg1, arg2)
}

// Count is a free data retrieval call binding the contract method 0x7849ed5a.
//
// Solidity: function count(address , bool , bool ) constant returns(uint256)
func (_Confirmations *ConfirmationsCallerSession) Count(arg0 common.Address, arg1 bool, arg2 bool) (*big.Int, error) {
	return _Confirmations.Contract.Count(&_Confirmations.CallOpts, arg0, arg1, arg2)
}

// Status is a free data retrieval call binding the contract method 0xfad3ffd6.
//
// Solidity: function status(uint256 , uint256 , bytes32 ) constant returns(uint8)
func (_Confirmations *ConfirmationsCaller) Status(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int, arg2 [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Confirmations.contract.Call(opts, out, "status", arg0, arg1, arg2)
	return *ret0, err
}

// Status is a free data retrieval call binding the contract method 0xfad3ffd6.
//
// Solidity: function status(uint256 , uint256 , bytes32 ) constant returns(uint8)
func (_Confirmations *ConfirmationsSession) Status(arg0 *big.Int, arg1 *big.Int, arg2 [32]byte) (uint8, error) {
	return _Confirmations.Contract.Status(&_Confirmations.CallOpts, arg0, arg1, arg2)
}

// Status is a free data retrieval call binding the contract method 0xfad3ffd6.
//
// Solidity: function status(uint256 , uint256 , bytes32 ) constant returns(uint8)
func (_Confirmations *ConfirmationsCallerSession) Status(arg0 *big.Int, arg1 *big.Int, arg2 [32]byte) (uint8, error) {
	return _Confirmations.Contract.Status(&_Confirmations.CallOpts, arg0, arg1, arg2)
}

// TotalConfirmGas is a free data retrieval call binding the contract method 0x5c5f6b39.
//
// Solidity: function totalConfirmGas() constant returns(uint256)
func (_Confirmations *ConfirmationsCaller) TotalConfirmGas(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Confirmations.contract.Call(opts, out, "totalConfirmGas")
	return *ret0, err
}

// TotalConfirmGas is a free data retrieval call binding the contract method 0x5c5f6b39.
//
// Solidity: function totalConfirmGas() constant returns(uint256)
func (_Confirmations *ConfirmationsSession) TotalConfirmGas() (*big.Int, error) {
	return _Confirmations.Contract.TotalConfirmGas(&_Confirmations.CallOpts)
}

// TotalConfirmGas is a free data retrieval call binding the contract method 0x5c5f6b39.
//
// Solidity: function totalConfirmGas() constant returns(uint256)
func (_Confirmations *ConfirmationsCallerSession) TotalConfirmGas() (*big.Int, error) {
	return _Confirmations.Contract.TotalConfirmGas(&_Confirmations.CallOpts)
}

// Votes is a free data retrieval call binding the contract method 0xd8bff5a5.
//
// Solidity: function votes(address ) constant returns(address addr, bool voter, bool add)
func (_Confirmations *ConfirmationsCaller) Votes(opts *bind.CallOpts, arg0 common.Address) (struct {
	Addr  common.Address
	Voter bool
	Add   bool
}, error) {
	ret := new(struct {
		Addr  common.Address
		Voter bool
		Add   bool
	})
	out := ret
	err := _Confirmations.contract.Call(opts, out, "votes", arg0)
	return *ret, err
}

// Votes is a free data retrieval call binding the contract method 0xd8bff5a5.
//
// Solidity: function votes(address ) constant returns(address addr, bool voter, bool add)
func (_Confirmations *ConfirmationsSession) Votes(arg0 common.Address) (struct {
	Addr  common.Address
	Voter bool
	Add   bool
}, error) {
	return _Confirmations.Contract.Votes(&_Confirmations.CallOpts, arg0)
}

// Votes is a free data retrieval call binding the contract method 0xd8bff5a5.
//
// Solidity: function votes(address ) constant returns(address addr, bool voter, bool add)
func (_Confirmations *ConfirmationsCallerSession) Votes(arg0 common.Address) (struct {
	Addr  common.Address
	Voter bool
	Add   bool
}, error) {
	return _Confirmations.Contract.Votes(&_Confirmations.CallOpts, arg0)
}

// Confirm is a paid mutator transaction binding the contract method 0x8f7fd0d6.
//
// Solidity: function confirm(uint256 blockNum, uint256 logIndex, bytes32 eventHash, bool valid) returns()
func (_Confirmations *ConfirmationsTransactor) Confirm(opts *bind.TransactOpts, blockNum *big.Int, logIndex *big.Int, eventHash [32]byte, valid bool) (*types.Transaction, error) {
	return _Confirmations.contract.Transact(opts, "confirm", blockNum, logIndex, eventHash, valid)
}

// Confirm is a paid mutator transaction binding the contract method 0x8f7fd0d6.
//
// Solidity: function confirm(uint256 blockNum, uint256 logIndex, bytes32 eventHash, bool valid) returns()
func (_Confirmations *ConfirmationsSession) Confirm(blockNum *big.Int, logIndex *big.Int, eventHash [32]byte, valid bool) (*types.Transaction, error) {
	return _Confirmations.Contract.Confirm(&_Confirmations.TransactOpts, blockNum, logIndex, eventHash, valid)
}

// Confirm is a paid mutator transaction binding the contract method 0x8f7fd0d6.
//
// Solidity: function confirm(uint256 blockNum, uint256 logIndex, bytes32 eventHash, bool valid) returns()
func (_Confirmations *ConfirmationsTransactorSession) Confirm(blockNum *big.Int, logIndex *big.Int, eventHash [32]byte, valid bool) (*types.Transaction, error) {
	return _Confirmations.Contract.Confirm(&_Confirmations.TransactOpts, blockNum, logIndex, eventHash, valid)
}

// Request is a paid mutator transaction binding the contract method 0x32030803.
//
// Solidity: function request(uint256 blockNum, uint256 logIndex, bytes32 eventHash) returns()
func (_Confirmations *ConfirmationsTransactor) Request(opts *bind.TransactOpts, blockNum *big.Int, logIndex *big.Int, eventHash [32]byte) (*types.Transaction, error) {
	return _Confirmations.contract.Transact(opts, "request", blockNum, logIndex, eventHash)
}

// Request is a paid mutator transaction binding the contract method 0x32030803.
//
// Solidity: function request(uint256 blockNum, uint256 logIndex, bytes32 eventHash) returns()
func (_Confirmations *ConfirmationsSession) Request(blockNum *big.Int, logIndex *big.Int, eventHash [32]byte) (*types.Transaction, error) {
	return _Confirmations.Contract.Request(&_Confirmations.TransactOpts, blockNum, logIndex, eventHash)
}

// Request is a paid mutator transaction binding the contract method 0x32030803.
//
// Solidity: function request(uint256 blockNum, uint256 logIndex, bytes32 eventHash) returns()
func (_Confirmations *ConfirmationsTransactorSession) Request(blockNum *big.Int, logIndex *big.Int, eventHash [32]byte) (*types.Transaction, error) {
	return _Confirmations.Contract.Request(&_Confirmations.TransactOpts, blockNum, logIndex, eventHash)
}

// SetVote is a paid mutator transaction binding the contract method 0x48c5000d.
//
// Solidity: function setVote(address addr, bool voter, bool add) returns()
func (_Confirmations *ConfirmationsTransactor) SetVote(opts *bind.TransactOpts, addr common.Address, voter bool, add bool) (*types.Transaction, error) {
	return _Confirmations.contract.Transact(opts, "setVote", addr, voter, add)
}

// SetVote is a paid mutator transaction binding the contract method 0x48c5000d.
//
// Solidity: function setVote(address addr, bool voter, bool add) returns()
func (_Confirmations *ConfirmationsSession) SetVote(addr common.Address, voter bool, add bool) (*types.Transaction, error) {
	return _Confirmations.Contract.SetVote(&_Confirmations.TransactOpts, addr, voter, add)
}

// SetVote is a paid mutator transaction binding the contract method 0x48c5000d.
//
// Solidity: function setVote(address addr, bool voter, bool add) returns()
func (_Confirmations *ConfirmationsTransactorSession) SetVote(addr common.Address, voter bool, add bool) (*types.Transaction, error) {
	return _Confirmations.Contract.SetVote(&_Confirmations.TransactOpts, addr, voter, add)
}

// ConfirmationsConfirmationRequestedIterator is returned from FilterConfirmationRequested and is used to iterate over the raw logs and unpacked data for ConfirmationRequested events raised by the Confirmations contract.
type ConfirmationsConfirmationRequestedIterator struct {
	Event *ConfirmationsConfirmationRequested // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log       // Log channel receiving the found contract events
	sub  gochain.Subscription // Subscription for errors, completion and termination
	done bool                 // Whether the subscription completed delivering logs
	fail error                // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ConfirmationsConfirmationRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfirmationsConfirmationRequested)
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
		it.Event = new(ConfirmationsConfirmationRequested)
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
func (it *ConfirmationsConfirmationRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfirmationsConfirmationRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfirmationsConfirmationRequested represents a ConfirmationRequested event raised by the Confirmations contract.
type ConfirmationsConfirmationRequested struct {
	BlockNum  *big.Int
	LogIndex  *big.Int
	EventHash [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterConfirmationRequested is a free log retrieval operation binding the contract event 0x2382b1fef9adb7c2a096ef31e2e402e08d3b5f9eb5fc18ea9a56835e44627b03.
//
// Solidity: event ConfirmationRequested(uint256 indexed blockNum, uint256 indexed logIndex, bytes32 eventHash)
func (_Confirmations *ConfirmationsFilterer) FilterConfirmationRequested(opts *bind.FilterOpts, blockNum []*big.Int, logIndex []*big.Int) (*ConfirmationsConfirmationRequestedIterator, error) {

	var blockNumRule []interface{}
	for _, blockNumItem := range blockNum {
		blockNumRule = append(blockNumRule, blockNumItem)
	}
	var logIndexRule []interface{}
	for _, logIndexItem := range logIndex {
		logIndexRule = append(logIndexRule, logIndexItem)
	}

	logs, sub, err := _Confirmations.contract.FilterLogs(opts, "ConfirmationRequested", blockNumRule, logIndexRule)
	if err != nil {
		return nil, err
	}
	return &ConfirmationsConfirmationRequestedIterator{contract: _Confirmations.contract, event: "ConfirmationRequested", logs: logs, sub: sub}, nil
}

// WatchConfirmationRequested is a free log subscription operation binding the contract event 0x2382b1fef9adb7c2a096ef31e2e402e08d3b5f9eb5fc18ea9a56835e44627b03.
//
// Solidity: event ConfirmationRequested(uint256 indexed blockNum, uint256 indexed logIndex, bytes32 eventHash)
func (_Confirmations *ConfirmationsFilterer) WatchConfirmationRequested(opts *bind.WatchOpts, sink chan<- *ConfirmationsConfirmationRequested, blockNum []*big.Int, logIndex []*big.Int) (event.Subscription, error) {

	var blockNumRule []interface{}
	for _, blockNumItem := range blockNum {
		blockNumRule = append(blockNumRule, blockNumItem)
	}
	var logIndexRule []interface{}
	for _, logIndexItem := range logIndex {
		logIndexRule = append(logIndexRule, logIndexItem)
	}

	logs, sub, err := _Confirmations.contract.WatchLogs(opts, "ConfirmationRequested", blockNumRule, logIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfirmationsConfirmationRequested)
				if err := _Confirmations.contract.UnpackLog(event, "ConfirmationRequested", log); err != nil {
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

// ConfirmationsConfirmedIterator is returned from FilterConfirmed and is used to iterate over the raw logs and unpacked data for Confirmed events raised by the Confirmations contract.
type ConfirmationsConfirmedIterator struct {
	Event *ConfirmationsConfirmed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log       // Log channel receiving the found contract events
	sub  gochain.Subscription // Subscription for errors, completion and termination
	done bool                 // Whether the subscription completed delivering logs
	fail error                // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ConfirmationsConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfirmationsConfirmed)
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
		it.Event = new(ConfirmationsConfirmed)
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
func (it *ConfirmationsConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfirmationsConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfirmationsConfirmed represents a Confirmed event raised by the Confirmations contract.
type ConfirmationsConfirmed struct {
	BlockNum  *big.Int
	LogIndex  *big.Int
	EventHash [32]byte
	Valid     bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterConfirmed is a free log retrieval operation binding the contract event 0xdf1051063b9bab79d715a3919f387eb9ee4291be1e9241dffb4500694141f25c.
//
// Solidity: event Confirmed(uint256 indexed blockNum, uint256 indexed logIndex, bytes32 eventHash, bool valid)
func (_Confirmations *ConfirmationsFilterer) FilterConfirmed(opts *bind.FilterOpts, blockNum []*big.Int, logIndex []*big.Int) (*ConfirmationsConfirmedIterator, error) {

	var blockNumRule []interface{}
	for _, blockNumItem := range blockNum {
		blockNumRule = append(blockNumRule, blockNumItem)
	}
	var logIndexRule []interface{}
	for _, logIndexItem := range logIndex {
		logIndexRule = append(logIndexRule, logIndexItem)
	}

	logs, sub, err := _Confirmations.contract.FilterLogs(opts, "Confirmed", blockNumRule, logIndexRule)
	if err != nil {
		return nil, err
	}
	return &ConfirmationsConfirmedIterator{contract: _Confirmations.contract, event: "Confirmed", logs: logs, sub: sub}, nil
}

// WatchConfirmed is a free log subscription operation binding the contract event 0xdf1051063b9bab79d715a3919f387eb9ee4291be1e9241dffb4500694141f25c.
//
// Solidity: event Confirmed(uint256 indexed blockNum, uint256 indexed logIndex, bytes32 eventHash, bool valid)
func (_Confirmations *ConfirmationsFilterer) WatchConfirmed(opts *bind.WatchOpts, sink chan<- *ConfirmationsConfirmed, blockNum []*big.Int, logIndex []*big.Int) (event.Subscription, error) {

	var blockNumRule []interface{}
	for _, blockNumItem := range blockNum {
		blockNumRule = append(blockNumRule, blockNumItem)
	}
	var logIndexRule []interface{}
	for _, logIndexItem := range logIndex {
		logIndexRule = append(logIndexRule, logIndexItem)
	}

	logs, sub, err := _Confirmations.contract.WatchLogs(opts, "Confirmed", blockNumRule, logIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfirmationsConfirmed)
				if err := _Confirmations.contract.UnpackLog(event, "Confirmed", log); err != nil {
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
