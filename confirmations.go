// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gost

import (
	"math/big"
	"strings"

	"github.com/gochain/gochain/v3"
	"github.com/gochain/gochain/v3/accounts/abi"
	"github.com/gochain/gochain/v3/accounts/abi/bind"
	"github.com/gochain/gochain/v3/common"
	"github.com/gochain/gochain/v3/core/types"
	"github.com/gochain/gochain/v3/event"
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
const ConfirmationsABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"_confirmedGas\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"blockNum\",\"type\":\"uint256\"},{\"name\":\"logIndex\",\"type\":\"uint256\"},{\"name\":\"eventHash\",\"type\":\"bytes32\"}],\"name\":\"request\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"getSigner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"signersLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"confirmGas\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"voter\",\"type\":\"bool\"},{\"name\":\"add\",\"type\":\"bool\"}],\"name\":\"setVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalConfirmGas\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"votersLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"bool\"}],\"name\":\"count\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"isSigner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"blockNum\",\"type\":\"uint256\"},{\"name\":\"logIndex\",\"type\":\"uint256\"},{\"name\":\"eventHash\",\"type\":\"bytes32\"},{\"name\":\"valid\",\"type\":\"bool\"}],\"name\":\"confirm\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"isVoter\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"getVoter\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"votes\",\"outputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"voter\",\"type\":\"bool\"},{\"name\":\"add\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"status\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"voter\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"blockNum\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"logIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"eventHash\",\"type\":\"bytes32\"}],\"name\":\"ConfirmationRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"blockNum\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"logIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"eventHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"valid\",\"type\":\"bool\"}],\"name\":\"Confirmed\",\"type\":\"event\"}]"

// ConfirmationsBin is the compiled bytecode used for deploying new contracts.
var ConfirmationsBin = "0x608060405234801561001057600080fd5b50604051602080620018a08339810180604052602081101561003157600080fd5b50518061004a600082610066602090811b610f3017901c565b61005f60028261006660201b610f301760201c565b505061012f565b6001600160a01b0381166000908152600183016020526040902054156100ed57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600e60248201527f416c726561647920696e20736574000000000000000000000000000000000000604482015290519081900360640190fd5b815460018181018085556000858152602080822090940180546001600160a01b039096166001600160a01b031990961686179055938452930190526040902055565b611761806200013f6000396000f3fe6080604052600436106100e85760003560e01c80636c6c39fb1161008a578063a7771ee311610059578063a7771ee3146102e4578063d07bff0c14610317578063d8bff5a514610341578063fad3ffd61461039e576100e8565b80636c6c39fb146102075780637849ed5a1461021c5780637df73e271461025f5780638f7fd0d6146102a6576100e8565b806341f684f3116100c657806341f684f31461018557806342715aec1461019a57806348c5000d146101af5780635c5f6b39146101f2576100e8565b80632ea6ece8146100ed57806332030803146101145780633ffefe4e1461013f575b600080fd5b3480156100f957600080fd5b506101026103f8565b60408051918252519081900360200190f35b61013d6004803603606081101561012a57600080fd5b50803590602081013590604001356103ff565b005b34801561014b57600080fd5b506101696004803603602081101561016257600080fd5b5035610654565b604080516001600160a01b039092168252519081900360200190f35b34801561019157600080fd5b50610102610681565b3480156101a657600080fd5b50610102610688565b3480156101bb57600080fd5b5061013d600480360360608110156101d257600080fd5b506001600160a01b0381351690602081013515159060400135151561068f565b3480156101fe57600080fd5b506101026108f3565b34801561021357600080fd5b506101026108fa565b34801561022857600080fd5b506101026004803603606081101561023f57600080fd5b506001600160a01b03813516906020810135151590604001351515610900565b34801561026b57600080fd5b506102926004803603602081101561028257600080fd5b50356001600160a01b0316610923565b604080519115158252519081900360200190f35b3480156102b257600080fd5b5061013d600480360360808110156102c957600080fd5b50803590602081013590604081013590606001351515610940565b3480156102f057600080fd5b506102926004803603602081101561030757600080fd5b50356001600160a01b0316610b5d565b34801561032357600080fd5b506101696004803603602081101561033a57600080fd5b5035610b7a565b34801561034d57600080fd5b506103746004803603602081101561036457600080fd5b50356001600160a01b0316610b8b565b604080516001600160a01b0390941684529115156020840152151582820152519081900360600190f35b3480156103aa57600080fd5b506103d4600480360360608110156103c157600080fd5b5080359060208101359060400135610bbc565b604051808260038111156103e457fe5b60ff16815260200191505060405180910390f35b6202710081565b6000838152600660209081526040808320858452825280832084845290915281205460ff16600381111561042f57fe5b146104845760408051600160e51b62461bcd02815260206004820152601360248201527f537461747573206d757374206265204e6f6e6500000000000000000000000000604482015290519081900360640190fd5b600083815260066020908152604080832085845282528083208484529091529020805460ff191660011790553a622625a002348111156104f857604051600160e51b62461bcd02815260040180806020018281038252602781526020018061170f6027913960400191505060405180910390fd5b803411156105315760405133903483900380156108fc02916000818181858888f1935050505015801561052f573d6000803e3d6000fd5b505b6000848152600760209081526040808320868452825280832085845282528083203a6004909101558051606081018252878152808301878152818301878152600880546001810180835591885293517ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee360039095029485015591517ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee4840155517ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee59092019190915587845260098352818420878552835281842086855283529281902083905580518581529051869288927f2382b1fef9adb7c2a096ef31e2e402e08d3b5f9eb5fc18ea9a56835e44627b03929081900390910190a35050505050565b60006002600001828154811061066657fe5b6000918252602090912001546001600160a01b031692915050565b6002545b90565b62015f9081565b336000908152600160205260409020546106f35760408051600160e51b62461bcd02815260206004820152601260248201527f53656e646572206e6f74206120766f7465720000000000000000000000000000604482015290519081900360640190fd5b33600090815260046020526040902080546001600160a01b03161561078f5780546001600160a01b03808216600090815260056020908152604080832060ff600160a01b8704811615158552908352818420600160a81b9096041615158352939052919091208054600019019055841661078a575033600090815260046020526040902080546001600160b01b03191690556108ee565b6107a3565b6001600160a01b0384166107a357506108ee565b604080516060810182526001600160a01b03808716808352861515602080850182815288151586880181815233600090815260048086528a822099518a54955193511515600160a81b02600160a81b60ff0219941515600160a01b0274ff00000000000000000000000000000000000000001992909b166001600160a01b03199097169690961716989098179190911692909217909655928352600581528583209183529081528482209382529283528381208054600101908190558451600160e21b633881de0f02815292830191909152925173__$1c537defa8d8abb880144579cecc12eeda$__9263e207783c9260248082019391829003018186803b1580156108ae57600080fd5b505af41580156108c2573d6000803e3d6000fd5b505050506040513d60208110156108d857600080fd5b505181106108eb576108eb858585610be2565b50505b505050565b622625a081565b60005490565b600560209081526000938452604080852082529284528284209052825290205481565b6001600160a01b0316600090815260036020526040902054151590565b336000908152600360205260409020546109a45760408051600160e51b62461bcd02815260206004820152601360248201527f53656e646572206e6f742061207369676e657200000000000000000000000000604482015290519081900360640190fd5b60016000858152600660209081526040808320878452825280832086845290915290205460ff1660038111156109d657fe5b146109e057610b57565b60008481526007602090815260408083208684528252808320858452909152902060048101543a14610a4657604051600160e51b62461bcd0281526004018080602001828103825260228152602001806116ed6022913960400191505060405180910390fd5b808215610a765733600090815260038301602052604090205415610a7157610a718260020133610e11565b610a9d565b50336000908152600182016020526040902054600282019015610a9d57610a9d8233610e11565b336000908152600182016020526040902054610b5457610abd8133610f30565b73__$1c537defa8d8abb880144579cecc12eeda$__63e207783c60026040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b158015610b0d57600080fd5b505af4158015610b21573d6000803e3d6000fd5b505050506040513d6020811015610b3757600080fd5b505181541015610b48575050610b57565b610b5486868686610fe2565b50505b50505050565b6001600160a01b0316600090815260016020526040902054151590565b600080600001828154811061066657fe5b6004602052600090815260409020546001600160a01b0381169060ff600160a01b8204811691600160a81b90041683565b600660209081526000938452604080852082529284528284209052825290205460ff1681565b60005b600054811015610c9f576000806000018281548110610c0057fe5b60009182526020808320909101546001600160a01b0390811680845260049092526040909220805491935091908116908716148015610c4e5750805460ff600160a01b909104161515851515145b8015610c695750805460ff600160a81b909104161515841515145b15610c95576001600160a01b038216600090815260046020526040902080546001600160b01b03191690555b5050600101610be5565b506001600160a01b038316600090815260056020908152604080832085158015855290835281842085151585529092528220919091558290610cde5750805b15610d38576001600160a01b038316600090815260036020526040902054610d0b57610d0b600284610f30565b6001600160a01b038316600090815260016020526040902054610d3357610d33600084610f30565b6108ee565b818015610d43575080155b15610d6f576001600160a01b03831660009081526001602052604090205415610d3357610d33836112c0565b81158015610d7a5750805b15610da7576001600160a01b038316600090815260036020526040902054610d3357610d33600284610f30565b81158015610db3575080155b156108ee576001600160a01b03831660009081526001602052604090205415610ddf57610ddf836112c0565b6001600160a01b038316600090815260036020526040902054156108ee57610e08600284610e11565b6108ee836114b1565b6001600160a01b038116600090815260018301602052604090205480610e715760408051600160e51b62461bcd02815260206004820152600a6024820152600160b21b69139bdd081a5b881cd95d02604482015290519081900360640190fd5b6001600160a01b03821660009081526001840160205260408120558254836000198201828110610e9d57fe5b60009182526020909120015484546001600160a01b039091169085906000198501908110610ec757fe5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550836000016001820381548110610f0757fe5b600091825260209091200180546001600160a01b031916905583546108eb856000198301611639565b6001600160a01b038116600090815260018301602052604090205415610fa05760408051600160e51b62461bcd02815260206004820152600e60248201527f416c726561647920696e20736574000000000000000000000000000000000000604482015290519081900360640190fd5b815460018181018085556000858152602080822090940180546001600160a01b039096166001600160a01b031990961686179055938452930190526040902055565b60008115610ff257506003610ff6565b5060025b600085815260066020908152604080832087845282528083208684529091529020805482919060ff1916600183600381111561102e57fe5b021790555060008581526007602090815260408083208784528252808320868452909152902080836110605750600281015b805460048301546223b4a002600082828161107757fe5b04905060005b838110156110e357600085600001828154811061109657fe5b60009182526020822001546040516001600160a01b039091169250829185156108fc02918691818181858888f193505050501580156110d9573d6000803e3d6000fd5b505060010161107d565b506004850154604051848302840391620271000290339082840180156108fc02916000818181858888f19350505050158015611123573d6000803e3d6000fd5b5060008c81526007602090815260408083208e845282528083208d84529091528120908181611152828261165d565b5050600282016000611164828261165d565b50506000600492909201829055508c81526009602090815260408083208e845282528083208d845290915281208054919055600854600181111561122a576000600860018303815481106111b457fe5b9060005260206000209060030201905080600860018503815481106111d557fe5b6000918252602080832084546003909302019182556001808501548184015560029485015492850192909255845483526009815260408084209286015484529181528183209490930154825292909152208290555b6008600182038154811061123a57fe5b6000918252602082206003909102018181556001810182905560020155600880549061126a90600019830161167b565b508c8e7fdf1051063b9bab79d715a3919f387eb9ee4291be1e9241dffb4500694141f25c8e8e60405180838152602001821515151581526020019250505060405180910390a35050505050505050505050505050565b6112cb600082610e11565b6001600160a01b0380821660009081526004602052604090208054909116156113565780546001600160a01b03808216600090815260056020908152604080832060ff600160a01b8704811615158552908352818420600160a81b90960416151583529381528382208054600019019055918516815260049091522080546001600160b01b03191690555b600073__$1c537defa8d8abb880144579cecc12eeda$__63e207783c60006040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b1580156113a857600080fd5b505af41580156113bc573d6000803e3d6000fd5b505050506040513d60208110156113d257600080fd5b5051905060005b600054811015610b5757600460008060000183815481106113f657fe5b60009182526020808320909101546001600160a01b0390811684529083019390935260409091019020805490945016156114a65782546001600160a01b038116600090815260056020908152604080832060ff600160a01b8604811615158552908352818420600160a81b9095041615158352929052205482116114a657825461149e906001600160a01b0381169060ff600160a01b8204811691600160a81b900416610be2565b5050506114ae565b6001016113d9565b50565b600073__$1c537defa8d8abb880144579cecc12eeda$__63e207783c60026040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b15801561150357600080fd5b505af4158015611517573d6000803e3d6000fd5b505050506040513d602081101561152d57600080fd5b5051905060005b6008548110156108ee5760006008828154811061154d57fe5b6000918252602080832060039092029091018054835260078252604080842060018084015486529084528185206002840154865284528185206001600160a01b038a1686529081019093529092205491925090156115b4576115af8186610e11565b6115e1565b6001600160a01b0385166000908152600382016020526040902054156115e1576115e18160020186610e11565b80548411611607576116028260000154836001015484600201546001610fe2565b611632565b6002810154841161162b576116028260000154836001015484600201546000610fe2565b6001909201915b5050611534565b8154818355818111156108ee576000838152602090206108ee9181019083016116a7565b50805460008255906000526020600020908101906114ae91906116a7565b8154818355818111156108ee576003028160030283600052602060002091820191016108ee91906116c5565b61068591905b808211156116c157600081556001016116ad565b5090565b61068591905b808211156116c15760008082556001820181905560028201556003016116cb56fe47617320707269636520646f65736e2774206d61746368207265717565737465642e54782076616c756520646f65736e277420636f76657220636f6e6669726d6174696f6e20666565a165627a7a72305820a3ed023288ae2a50cdd06ce8b62d5fd1521d278032eccaecb0a33ca5b50d0d290029"

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

// GetSigner is a free data retrieval call binding the contract method 0x3ffefe4e.
//
// Solidity: function getSigner(uint256 i) constant returns(address)
func (_Confirmations *ConfirmationsCaller) GetSigner(opts *bind.CallOpts, i *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Confirmations.contract.Call(opts, out, "getSigner", i)
	return *ret0, err
}

// GetSigner is a free data retrieval call binding the contract method 0x3ffefe4e.
//
// Solidity: function getSigner(uint256 i) constant returns(address)
func (_Confirmations *ConfirmationsSession) GetSigner(i *big.Int) (common.Address, error) {
	return _Confirmations.Contract.GetSigner(&_Confirmations.CallOpts, i)
}

// GetSigner is a free data retrieval call binding the contract method 0x3ffefe4e.
//
// Solidity: function getSigner(uint256 i) constant returns(address)
func (_Confirmations *ConfirmationsCallerSession) GetSigner(i *big.Int) (common.Address, error) {
	return _Confirmations.Contract.GetSigner(&_Confirmations.CallOpts, i)
}

// GetVoter is a free data retrieval call binding the contract method 0xd07bff0c.
//
// Solidity: function getVoter(uint256 i) constant returns(address)
func (_Confirmations *ConfirmationsCaller) GetVoter(opts *bind.CallOpts, i *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Confirmations.contract.Call(opts, out, "getVoter", i)
	return *ret0, err
}

// GetVoter is a free data retrieval call binding the contract method 0xd07bff0c.
//
// Solidity: function getVoter(uint256 i) constant returns(address)
func (_Confirmations *ConfirmationsSession) GetVoter(i *big.Int) (common.Address, error) {
	return _Confirmations.Contract.GetVoter(&_Confirmations.CallOpts, i)
}

// GetVoter is a free data retrieval call binding the contract method 0xd07bff0c.
//
// Solidity: function getVoter(uint256 i) constant returns(address)
func (_Confirmations *ConfirmationsCallerSession) GetVoter(i *big.Int) (common.Address, error) {
	return _Confirmations.Contract.GetVoter(&_Confirmations.CallOpts, i)
}

// IsSigner is a free data retrieval call binding the contract method 0x7df73e27.
//
// Solidity: function isSigner(address voter) constant returns(bool)
func (_Confirmations *ConfirmationsCaller) IsSigner(opts *bind.CallOpts, voter common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Confirmations.contract.Call(opts, out, "isSigner", voter)
	return *ret0, err
}

// IsSigner is a free data retrieval call binding the contract method 0x7df73e27.
//
// Solidity: function isSigner(address voter) constant returns(bool)
func (_Confirmations *ConfirmationsSession) IsSigner(voter common.Address) (bool, error) {
	return _Confirmations.Contract.IsSigner(&_Confirmations.CallOpts, voter)
}

// IsSigner is a free data retrieval call binding the contract method 0x7df73e27.
//
// Solidity: function isSigner(address voter) constant returns(bool)
func (_Confirmations *ConfirmationsCallerSession) IsSigner(voter common.Address) (bool, error) {
	return _Confirmations.Contract.IsSigner(&_Confirmations.CallOpts, voter)
}

// IsVoter is a free data retrieval call binding the contract method 0xa7771ee3.
//
// Solidity: function isVoter(address voter) constant returns(bool)
func (_Confirmations *ConfirmationsCaller) IsVoter(opts *bind.CallOpts, voter common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Confirmations.contract.Call(opts, out, "isVoter", voter)
	return *ret0, err
}

// IsVoter is a free data retrieval call binding the contract method 0xa7771ee3.
//
// Solidity: function isVoter(address voter) constant returns(bool)
func (_Confirmations *ConfirmationsSession) IsVoter(voter common.Address) (bool, error) {
	return _Confirmations.Contract.IsVoter(&_Confirmations.CallOpts, voter)
}

// IsVoter is a free data retrieval call binding the contract method 0xa7771ee3.
//
// Solidity: function isVoter(address voter) constant returns(bool)
func (_Confirmations *ConfirmationsCallerSession) IsVoter(voter common.Address) (bool, error) {
	return _Confirmations.Contract.IsVoter(&_Confirmations.CallOpts, voter)
}

// SignersLength is a free data retrieval call binding the contract method 0x41f684f3.
//
// Solidity: function signersLength() constant returns(uint256)
func (_Confirmations *ConfirmationsCaller) SignersLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Confirmations.contract.Call(opts, out, "signersLength")
	return *ret0, err
}

// SignersLength is a free data retrieval call binding the contract method 0x41f684f3.
//
// Solidity: function signersLength() constant returns(uint256)
func (_Confirmations *ConfirmationsSession) SignersLength() (*big.Int, error) {
	return _Confirmations.Contract.SignersLength(&_Confirmations.CallOpts)
}

// SignersLength is a free data retrieval call binding the contract method 0x41f684f3.
//
// Solidity: function signersLength() constant returns(uint256)
func (_Confirmations *ConfirmationsCallerSession) SignersLength() (*big.Int, error) {
	return _Confirmations.Contract.SignersLength(&_Confirmations.CallOpts)
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

// VotersLength is a free data retrieval call binding the contract method 0x6c6c39fb.
//
// Solidity: function votersLength() constant returns(uint256)
func (_Confirmations *ConfirmationsCaller) VotersLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Confirmations.contract.Call(opts, out, "votersLength")
	return *ret0, err
}

// VotersLength is a free data retrieval call binding the contract method 0x6c6c39fb.
//
// Solidity: function votersLength() constant returns(uint256)
func (_Confirmations *ConfirmationsSession) VotersLength() (*big.Int, error) {
	return _Confirmations.Contract.VotersLength(&_Confirmations.CallOpts)
}

// VotersLength is a free data retrieval call binding the contract method 0x6c6c39fb.
//
// Solidity: function votersLength() constant returns(uint256)
func (_Confirmations *ConfirmationsCallerSession) VotersLength() (*big.Int, error) {
	return _Confirmations.Contract.VotersLength(&_Confirmations.CallOpts)
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

// ParseConfirmationRequested is a log parse operation binding the contract event 0x2382b1fef9adb7c2a096ef31e2e402e08d3b5f9eb5fc18ea9a56835e44627b03.
//
// Solidity: event ConfirmationRequested(uint256 indexed blockNum, uint256 indexed logIndex, bytes32 eventHash)
func (_Confirmations *ConfirmationsFilterer) ParseConfirmationRequested(log types.Log) (*ConfirmationsConfirmationRequested, error) {
	event := new(ConfirmationsConfirmationRequested)
	if err := _Confirmations.contract.UnpackLog(event, "ConfirmationRequested", log); err != nil {
		return nil, err
	}
	return event, nil
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

// ParseConfirmed is a log parse operation binding the contract event 0xdf1051063b9bab79d715a3919f387eb9ee4291be1e9241dffb4500694141f25c.
//
// Solidity: event Confirmed(uint256 indexed blockNum, uint256 indexed logIndex, bytes32 eventHash, bool valid)
func (_Confirmations *ConfirmationsFilterer) ParseConfirmed(log types.Log) (*ConfirmationsConfirmed, error) {
	event := new(ConfirmationsConfirmed)
	if err := _Confirmations.contract.UnpackLog(event, "Confirmed", log); err != nil {
		return nil, err
	}
	return event, nil
}
