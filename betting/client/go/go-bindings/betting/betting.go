// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// BettingMetaData contains all meta data concerning the Betting contract.
var BettingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"initOutcomes\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"gambler\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"outcome\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BetMade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOracle\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOracle\",\"type\":\"address\"}],\"name\":\"OracleChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"wins\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalPrize\",\"type\":\"uint256\"}],\"name\":\"Winners\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"gambler\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"bets\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"outcome\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"outcome\",\"type\":\"bytes32\"}],\"name\":\"checkOutcome\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"outcomeString\",\"type\":\"string\"}],\"name\":\"checkOutcomeString\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkWinnings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOracle\",\"type\":\"address\"}],\"name\":\"chooseOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractReset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decisionMade\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"gamblers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGamblers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOutcomes\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWinners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isGambler\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"}],\"name\":\"isOracle\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"outcome\",\"type\":\"bytes32\"}],\"name\":\"makeBet\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"decidedOutcome\",\"type\":\"bytes32\"}],\"name\":\"makeDecision\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"outcomeBets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"outcomes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_outcomes\",\"type\":\"bytes32[]\"}],\"name\":\"setOutcomes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"validOutcomes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"winners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wins\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"89a78f1a": "bets(address)",
		"52311c14": "checkOutcome(bytes32)",
		"8efe1833": "checkOutcomeString(string)",
		"5ee715d6": "checkWinnings()",
		"d0aa7e25": "chooseOracle(address)",
		"a7a5ce88": "contractReset()",
		"0ba83c6d": "decisionMade()",
		"2aaf9c66": "gamblers(uint256)",
		"09eeebb0": "getGamblers()",
		"3d3db8eb": "getOutcomes()",
		"df15c37e": "getWinners()",
		"a7fe401d": "isGambler(address)",
		"a97e5c93": "isOracle(address)",
		"284152dc": "makeBet(bytes32)",
		"b8ecdb24": "makeDecision(bytes32)",
		"7dc0d1d0": "oracle()",
		"27edd6cf": "outcomeBets(bytes32)",
		"eed2a147": "outcomes(uint256)",
		"8da5cb5b": "owner()",
		"6646586f": "setOutcomes(bytes32[])",
		"914cf4f5": "validOutcomes(bytes32)",
		"a2fb1175": "winners(uint256)",
		"8047a97a": "wins(address)",
		"2e1a7d4d": "withdraw(uint256)",
	},
	Bin: "0x60806040523480156200001157600080fd5b506040516200190d3803806200190d8339810160408190526200003491620001d2565b600281511015620000965760405162461bcd60e51b815260206004820152602160248201527f6d757374207265676973746572206174206c656173742032206f7574636f6d656044820152607360f81b60648201526084015b60405180910390fd5b600080546001600160a01b03191633179055620000b381620000ba565b50620002db565b6000546001600160a01b03163314620001165760405162461bcd60e51b815260206004820152601660248201527f73656e6465722069736e277420746865206f776e65720000000000000000000060448201526064016200008d565b60005b8151811015620001b85760088282815181106200013a576200013a6200029b565b60209081029190910181015182546001818101855560009485529284200155835190916009918590859081106200017557620001756200029b565b6020026020010151815260200190815260200160002060006101000a81548160ff0219169083151502179055508080620001af90620002b1565b91505062000119565b5050565b634e487b7160e01b600052604160045260246000fd5b60006020808385031215620001e657600080fd5b82516001600160401b0380821115620001fe57600080fd5b818501915085601f8301126200021357600080fd5b815181811115620002285762000228620001bc565b8060051b604051601f19603f83011681018181108582111715620002505762000250620001bc565b6040529182528482019250838101850191888311156200026f57600080fd5b938501935b828510156200028f5784518452938501939285019262000274565b98975050505050505050565b634e487b7160e01b600052603260045260246000fd5b6000600019821415620002d457634e487b7160e01b600052601160045260246000fd5b5060010190565b61162280620002eb6000396000f3fe60806040526004361061014b5760003560e01c806389a78f1a116100b6578063a7fe401d1161006f578063a7fe401d146103ff578063a97e5c931461042f578063b8ecdb241461045e578063d0aa7e251461047e578063df15c37e1461049e578063eed2a147146104b357600080fd5b806389a78f1a146103115780638da5cb5b1461035a5780638efe18331461037a578063914cf4f51461039a578063a2fb1175146103ca578063a7a5ce88146103ea57600080fd5b80633d3db8eb116101085780633d3db8eb1461024d57806352311c141461026f5780635ee715d61461028f5780636646586f146102a45780637dc0d1d0146102c45780638047a97a146102e457600080fd5b806309eeebb0146101505780630ba83c6d1461017b57806327edd6cf146101a5578063284152dc146101e05780632aaf9c66146101f55780632e1a7d4d1461022d575b600080fd5b34801561015c57600080fd5b506101656104d3565b6040516101729190611233565b60405180910390f35b34801561018757600080fd5b506005546101959060ff1681565b6040519015158152602001610172565b3480156101b157600080fd5b506101d26101c0366004611280565b600a6020526000908152604090205481565b604051908152602001610172565b6101f36101ee366004611280565b610535565b005b34801561020157600080fd5b50610215610210366004611280565b6107ff565b6040516001600160a01b039091168152602001610172565b34801561023957600080fd5b506101f3610248366004611280565b610829565b34801561025957600080fd5b50610262610994565b6040516101729190611299565b34801561027b57600080fd5b506101d261028a366004611280565b6109eb565b34801561029b57600080fd5b506101d2610a2c565b3480156102b057600080fd5b506101f36102bf366004611318565b610a5b565b3480156102d057600080fd5b50600154610215906001600160a01b031681565b3480156102f057600080fd5b506101d26102ff3660046113be565b60076020526000908152604090205481565b34801561031d57600080fd5b5061034561032c3660046113be565b6004602052600090815260409020805460019091015482565b60408051928352602083019190915201610172565b34801561036657600080fd5b50600054610215906001600160a01b031681565b34801561038657600080fd5b506101d26103953660046113e7565b610b21565b3480156103a657600080fd5b506101956103b5366004611280565b60096020526000908152604090205460ff1681565b3480156103d657600080fd5b506102156103e5366004611280565b610b3b565b3480156103f657600080fd5b506101f3610b4b565b34801561040b57600080fd5b5061019561041a3660046113be565b60036020526000908152604090205460ff1681565b34801561043b57600080fd5b5061019561044a3660046113be565b6001546001600160a01b0390811691161490565b34801561046a57600080fd5b506101f3610479366004611280565b610c51565b34801561048a57600080fd5b506101f36104993660046113be565b61102b565b3480156104aa57600080fd5b50610165611178565b3480156104bf57600080fd5b506101d26104ce366004611280565b6111d8565b6060600280548060200260200160405190810160405280929190818152602001828054801561052b57602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161050d575b5050505050905090565b6001546001600160a01b03166105845760405162461bcd60e51b815260206004820152600f60248201526e1b9bc81bdc9858db1948199bdd5b99608a1b60448201526064015b60405180910390fd5b6000546001600160a01b03163314156105d65760405162461bcd60e51b81526020600482015260146024820152731d1a19481bdddb995c8818d85b9b9bdd0818995d60621b604482015260640161057b565b6001546001600160a01b031633141561063d5760405162461bcd60e51b8152602060048201526024808201527f746865206f7261636c65206f66207468652062657474696e672063616e6e6f746044820152630818995d60e21b606482015260840161057b565b60055460ff161561069b5760405162461bcd60e51b815260206004820152602260248201527f63616e6e6f7420626574206166746572206465636973696f6e20776173206d61604482015261646560f01b606482015260840161057b565b3360009081526003602052604090205460ff16156106fb5760405162461bcd60e51b815260206004820152601e60248201527f656163682067616d626c65722063616e206f6e6c7920626574206f6e63650000604482015260640161057b565b60008181526009602052604090205460ff166107295760405162461bcd60e51b815260040161057b9061147c565b3360008181526004602090815260408083208581553460019182018190556002805492830190557f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace90910180546001600160a01b031916909517909455848352600a909152812080549091906107a09084906114c2565b90915550503360008181526003602052604090819020805460ff19166001179055518291907ff8e19b8ac4065a7305ef78373f54d2b12c631e1c04c0918be0342ac63f61c033906107f49034815260200190565b60405180910390a350565b6002818154811061080f57600080fd5b6000918252602090912001546001600160a01b0316905081565b336000908152600760205260409020546108ab5760405162461bcd60e51b815260206004820152603b60248201527f73656e6465722073686f756c6420626520612077696e6e6572206f7220696e7360448201527f756666696369656e742072657175657374656420616d6f756e74200000000000606482015260840161057b565b3360009081526007602052604090205481111561090a5760405162461bcd60e51b815260206004820152601d60248201527f696e73756666696369656e742072657175657374656420616d6f756e74000000604482015260640161057b565b33600090815260076020526040812080548392906109299084906114da565b9091555050604051339082156108fc029083906000818181858888f1935050505015801561095b573d6000803e3d6000fd5b5060405181815233907f7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d59060200160405180910390a250565b6060600880548060200260200160405190810160405280929190818152602001828054801561052b57602002820191906000526020600020905b8154815260200190600101908083116109ce575050505050905090565b60008181526009602052604081205460ff16610a195760405162461bcd60e51b815260040161057b9061147c565b506000908152600a602052604090205490565b3360009081526007602052604081205415610a5557503360009081526007602052604090205490565b50600090565b6000546001600160a01b03163314610a855760405162461bcd60e51b815260040161057b906114f1565b60005b8151811015610b1d576008828281518110610aa557610aa5611521565b6020908102919091018101518254600181810185556000948552928420015583519091600991859085908110610add57610add611521565b6020026020010151815260200190815260200160002060006101000a81548160ff0219169083151502179055508080610b1590611537565b915050610a88565b5050565b80516020820120600090610b34816109eb565b9392505050565b6006818154811061080f57600080fd5b6000546001600160a01b03163314610b755760405162461bcd60e51b815260040161057b906114f1565b60055460ff16610bc75760405162461bcd60e51b815260206004820152601c60248201527f63616e6e6f74207265736574206265666f7265206465636973696f6e00000000604482015260640161057b565b6005805460ff1916905560005b600854811015610c2a576009600060088381548110610bf557610bf5611521565b600091825260208083209091015483528201929092526040019020805460ff1916905580610c2281611537565b915050610bd4565b50610c37600860006111f9565b610c43600660006111f9565b610c4f600260006111f9565b565b6001546001600160a01b03163314610cab5760405162461bcd60e51b815260206004820152601760248201527f73656e6465722069736e277420746865206f7261636c65000000000000000000604482015260640161057b565b600081815260096020526040902054819060ff16610cdb5760405162461bcd60e51b815260040161057b9061147c565b60055460ff1615610d2e5760405162461bcd60e51b815260206004820152601b60248201527f63616e206d616b65206465636973696f6e206f6e6c79206f6e63650000000000604482015260640161057b565b60008281526009602052604090205460ff16610d5c5760405162461bcd60e51b815260040161057b9061147c565b600254610d6857600080fd5b600854610d7457600080fd5b6005805460ff19166001179055600080805b600254811015610ed057846004600060028481548110610da857610da8611521565b60009182526020808320909101546001600160a01b031683528201929092526040019020541415610e7357600660028281548110610de857610de8611521565b6000918252602080832090910154835460018101855593835290822090920180546001600160a01b0319166001600160a01b039093169290921790915560028054600492919084908110610e3e57610e3e611521565b60009182526020808320909101546001600160a01b03168352820192909252604001902060010154610e7090846114c2565b92505b6004600060028381548110610e8a57610e8a611521565b60009182526020808320909101546001600160a01b03168352820192909252604001902060010154610ebc90836114c2565b915080610ec881611537565b915050610d86565b5081610f395760018054600680548084019091557ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f0180546001600160a01b0319166001600160a01b039283161790559054166000908152600760205260409020819055610feb565b60005b600654811015610fe95782826004600060068581548110610f5f57610f5f611521565b60009182526020808320909101546001600160a01b03168352820192909252604001902060010154610f919190611552565b610f9b9190611571565b6007600060068481548110610fb257610fb2611521565b60009182526020808320909101546001600160a01b0316835282019290925260400190205580610fe181611537565b915050610f3c565b505b7f413cf1e17d70833c017de9b3c64d0104e3e37a7191d95b3d0d6276de992a0b4d60068260405161101d929190611593565b60405180910390a150505050565b6000546001600160a01b031633146110555760405162461bcd60e51b815260040161057b906114f1565b6001600160a01b03811660009081526003602052604090205460ff16156110be5760405162461bcd60e51b815260206004820152601e60248201527f746865206f7261636c652063616e6e6f7420626520612067616d626c65720000604482015260640161057b565b6000546001600160a01b038281169116141561111c5760405162461bcd60e51b815260206004820152601d60248201527f746865206f776e65722063616e6e6f7420626520616e206f7261636c65000000604482015260640161057b565b6001546040516001600160a01b038084169216907f05cd89403c6bdeac21c2ff33de395121a31fa1bc2bf3adf4825f1f86e79969dd90600090a3600180546001600160a01b0319166001600160a01b0392909216919091179055565b6060600680548060200260200160405190810160405280929190818152602001828054801561052b576020028201919060005260206000209081546001600160a01b0316815260019091019060200180831161050d575050505050905090565b600881815481106111e857600080fd5b600091825260209091200154905081565b5080546000825590600052602060002090810190611217919061121a565b50565b5b8082111561122f576000815560010161121b565b5090565b6020808252825182820181905260009190848201906040850190845b818110156112745783516001600160a01b03168352928401929184019160010161124f565b50909695505050505050565b60006020828403121561129257600080fd5b5035919050565b6020808252825182820181905260009190848201906040850190845b81811015611274578351835292840192918401916001016112b5565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff81118282101715611310576113106112d1565b604052919050565b6000602080838503121561132b57600080fd5b823567ffffffffffffffff8082111561134357600080fd5b818501915085601f83011261135757600080fd5b813581811115611369576113696112d1565b8060051b915061137a8483016112e7565b818152918301840191848101908884111561139457600080fd5b938501935b838510156113b257843582529385019390850190611399565b98975050505050505050565b6000602082840312156113d057600080fd5b81356001600160a01b0381168114610b3457600080fd5b600060208083850312156113fa57600080fd5b823567ffffffffffffffff8082111561141257600080fd5b818501915085601f83011261142657600080fd5b813581811115611438576114386112d1565b61144a601f8201601f191685016112e7565b9150808252868482850101111561146057600080fd5b8084840185840137600090820190930192909252509392505050565b6020808252601690820152751bdd5d18dbdb59481b9bdd081c9959da5cdd195c995960521b604082015260600190565b634e487b7160e01b600052601160045260246000fd5b600082198211156114d5576114d56114ac565b500190565b6000828210156114ec576114ec6114ac565b500390565b60208082526016908201527539b2b73232b91034b9b713ba103a34329037bbb732b960511b604082015260600190565b634e487b7160e01b600052603260045260246000fd5b600060001982141561154b5761154b6114ac565b5060010190565b600081600019048311821515161561156c5761156c6114ac565b500290565b60008261158e57634e487b7160e01b600052601260045260246000fd5b500490565b6000604082016040835280855480835260608501915086600052602092508260002060005b828110156115dd5781546001600160a01b0316845292840192600191820191016115b8565b5050509201929092529291505056fea26469706673582212207210ef846a4949ab7a861434acef5b9cca48069b6a88780c02ef03e6d698c1ab64736f6c63430008090033",
}

// BettingABI is the input ABI used to generate the binding from.
// Deprecated: Use BettingMetaData.ABI instead.
var BettingABI = BettingMetaData.ABI

// Deprecated: Use BettingMetaData.Sigs instead.
// BettingFuncSigs maps the 4-byte function signature to its string representation.
var BettingFuncSigs = BettingMetaData.Sigs

// BettingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BettingMetaData.Bin instead.
var BettingBin = BettingMetaData.Bin

// DeployBetting deploys a new Ethereum contract, binding an instance of Betting to it.
func DeployBetting(auth *bind.TransactOpts, backend bind.ContractBackend, initOutcomes [][32]byte) (common.Address, *types.Transaction, *Betting, error) {
	parsed, err := BettingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BettingBin), backend, initOutcomes)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Betting{BettingCaller: BettingCaller{contract: contract}, BettingTransactor: BettingTransactor{contract: contract}, BettingFilterer: BettingFilterer{contract: contract}}, nil
}

// Betting is an auto generated Go binding around an Ethereum contract.
type Betting struct {
	BettingCaller     // Read-only binding to the contract
	BettingTransactor // Write-only binding to the contract
	BettingFilterer   // Log filterer for contract events
}

// BettingCaller is an auto generated read-only Go binding around an Ethereum contract.
type BettingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BettingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BettingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BettingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BettingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BettingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BettingSession struct {
	Contract     *Betting          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BettingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BettingCallerSession struct {
	Contract *BettingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// BettingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BettingTransactorSession struct {
	Contract     *BettingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// BettingRaw is an auto generated low-level Go binding around an Ethereum contract.
type BettingRaw struct {
	Contract *Betting // Generic contract binding to access the raw methods on
}

// BettingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BettingCallerRaw struct {
	Contract *BettingCaller // Generic read-only contract binding to access the raw methods on
}

// BettingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BettingTransactorRaw struct {
	Contract *BettingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBetting creates a new instance of Betting, bound to a specific deployed contract.
func NewBetting(address common.Address, backend bind.ContractBackend) (*Betting, error) {
	contract, err := bindBetting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Betting{BettingCaller: BettingCaller{contract: contract}, BettingTransactor: BettingTransactor{contract: contract}, BettingFilterer: BettingFilterer{contract: contract}}, nil
}

// NewBettingCaller creates a new read-only instance of Betting, bound to a specific deployed contract.
func NewBettingCaller(address common.Address, caller bind.ContractCaller) (*BettingCaller, error) {
	contract, err := bindBetting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BettingCaller{contract: contract}, nil
}

// NewBettingTransactor creates a new write-only instance of Betting, bound to a specific deployed contract.
func NewBettingTransactor(address common.Address, transactor bind.ContractTransactor) (*BettingTransactor, error) {
	contract, err := bindBetting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BettingTransactor{contract: contract}, nil
}

// NewBettingFilterer creates a new log filterer instance of Betting, bound to a specific deployed contract.
func NewBettingFilterer(address common.Address, filterer bind.ContractFilterer) (*BettingFilterer, error) {
	contract, err := bindBetting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BettingFilterer{contract: contract}, nil
}

// bindBetting binds a generic wrapper to an already deployed contract.
func bindBetting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BettingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Betting *BettingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Betting.Contract.BettingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Betting *BettingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Betting.Contract.BettingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Betting *BettingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Betting.Contract.BettingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Betting *BettingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Betting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Betting *BettingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Betting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Betting *BettingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Betting.Contract.contract.Transact(opts, method, params...)
}

// Bets is a free data retrieval call binding the contract method 0x89a78f1a.
//
// Solidity: function bets(address ) view returns(bytes32 outcome, uint256 amount)
func (_Betting *BettingCaller) Bets(opts *bind.CallOpts, arg0 common.Address) (struct {
	Outcome [32]byte
	Amount  *big.Int
}, error) {
	var out []interface{}
	err := _Betting.contract.Call(opts, &out, "bets", arg0)

	outstruct := new(struct {
		Outcome [32]byte
		Amount  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Outcome = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Bets is a free data retrieval call binding the contract method 0x89a78f1a.
//
// Solidity: function bets(address ) view returns(bytes32 outcome, uint256 amount)
func (_Betting *BettingSession) Bets(arg0 common.Address) (struct {
	Outcome [32]byte
	Amount  *big.Int
}, error) {
	return _Betting.Contract.Bets(&_Betting.CallOpts, arg0)
}

// Bets is a free data retrieval call binding the contract method 0x89a78f1a.
//
// Solidity: function bets(address ) view returns(bytes32 outcome, uint256 amount)
func (_Betting *BettingCallerSession) Bets(arg0 common.Address) (struct {
	Outcome [32]byte
	Amount  *big.Int
}, error) {
	return _Betting.Contract.Bets(&_Betting.CallOpts, arg0)
}

// CheckOutcome is a free data retrieval call binding the contract method 0x52311c14.
//
// Solidity: function checkOutcome(bytes32 outcome) view returns(uint256)
func (_Betting *BettingCaller) CheckOutcome(opts *bind.CallOpts, outcome [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Betting.contract.Call(opts, &out, "checkOutcome", outcome)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CheckOutcome is a free data retrieval call binding the contract method 0x52311c14.
//
// Solidity: function checkOutcome(bytes32 outcome) view returns(uint256)
func (_Betting *BettingSession) CheckOutcome(outcome [32]byte) (*big.Int, error) {
	return _Betting.Contract.CheckOutcome(&_Betting.CallOpts, outcome)
}

// CheckOutcome is a free data retrieval call binding the contract method 0x52311c14.
//
// Solidity: function checkOutcome(bytes32 outcome) view returns(uint256)
func (_Betting *BettingCallerSession) CheckOutcome(outcome [32]byte) (*big.Int, error) {
	return _Betting.Contract.CheckOutcome(&_Betting.CallOpts, outcome)
}

// CheckOutcomeString is a free data retrieval call binding the contract method 0x8efe1833.
//
// Solidity: function checkOutcomeString(string outcomeString) view returns(uint256)
func (_Betting *BettingCaller) CheckOutcomeString(opts *bind.CallOpts, outcomeString string) (*big.Int, error) {
	var out []interface{}
	err := _Betting.contract.Call(opts, &out, "checkOutcomeString", outcomeString)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CheckOutcomeString is a free data retrieval call binding the contract method 0x8efe1833.
//
// Solidity: function checkOutcomeString(string outcomeString) view returns(uint256)
func (_Betting *BettingSession) CheckOutcomeString(outcomeString string) (*big.Int, error) {
	return _Betting.Contract.CheckOutcomeString(&_Betting.CallOpts, outcomeString)
}

// CheckOutcomeString is a free data retrieval call binding the contract method 0x8efe1833.
//
// Solidity: function checkOutcomeString(string outcomeString) view returns(uint256)
func (_Betting *BettingCallerSession) CheckOutcomeString(outcomeString string) (*big.Int, error) {
	return _Betting.Contract.CheckOutcomeString(&_Betting.CallOpts, outcomeString)
}

// CheckWinnings is a free data retrieval call binding the contract method 0x5ee715d6.
//
// Solidity: function checkWinnings() view returns(uint256)
func (_Betting *BettingCaller) CheckWinnings(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Betting.contract.Call(opts, &out, "checkWinnings")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CheckWinnings is a free data retrieval call binding the contract method 0x5ee715d6.
//
// Solidity: function checkWinnings() view returns(uint256)
func (_Betting *BettingSession) CheckWinnings() (*big.Int, error) {
	return _Betting.Contract.CheckWinnings(&_Betting.CallOpts)
}

// CheckWinnings is a free data retrieval call binding the contract method 0x5ee715d6.
//
// Solidity: function checkWinnings() view returns(uint256)
func (_Betting *BettingCallerSession) CheckWinnings() (*big.Int, error) {
	return _Betting.Contract.CheckWinnings(&_Betting.CallOpts)
}

// DecisionMade is a free data retrieval call binding the contract method 0x0ba83c6d.
//
// Solidity: function decisionMade() view returns(bool)
func (_Betting *BettingCaller) DecisionMade(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Betting.contract.Call(opts, &out, "decisionMade")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DecisionMade is a free data retrieval call binding the contract method 0x0ba83c6d.
//
// Solidity: function decisionMade() view returns(bool)
func (_Betting *BettingSession) DecisionMade() (bool, error) {
	return _Betting.Contract.DecisionMade(&_Betting.CallOpts)
}

// DecisionMade is a free data retrieval call binding the contract method 0x0ba83c6d.
//
// Solidity: function decisionMade() view returns(bool)
func (_Betting *BettingCallerSession) DecisionMade() (bool, error) {
	return _Betting.Contract.DecisionMade(&_Betting.CallOpts)
}

// Gamblers is a free data retrieval call binding the contract method 0x2aaf9c66.
//
// Solidity: function gamblers(uint256 ) view returns(address)
func (_Betting *BettingCaller) Gamblers(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Betting.contract.Call(opts, &out, "gamblers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gamblers is a free data retrieval call binding the contract method 0x2aaf9c66.
//
// Solidity: function gamblers(uint256 ) view returns(address)
func (_Betting *BettingSession) Gamblers(arg0 *big.Int) (common.Address, error) {
	return _Betting.Contract.Gamblers(&_Betting.CallOpts, arg0)
}

// Gamblers is a free data retrieval call binding the contract method 0x2aaf9c66.
//
// Solidity: function gamblers(uint256 ) view returns(address)
func (_Betting *BettingCallerSession) Gamblers(arg0 *big.Int) (common.Address, error) {
	return _Betting.Contract.Gamblers(&_Betting.CallOpts, arg0)
}

// GetGamblers is a free data retrieval call binding the contract method 0x09eeebb0.
//
// Solidity: function getGamblers() view returns(address[])
func (_Betting *BettingCaller) GetGamblers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Betting.contract.Call(opts, &out, "getGamblers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetGamblers is a free data retrieval call binding the contract method 0x09eeebb0.
//
// Solidity: function getGamblers() view returns(address[])
func (_Betting *BettingSession) GetGamblers() ([]common.Address, error) {
	return _Betting.Contract.GetGamblers(&_Betting.CallOpts)
}

// GetGamblers is a free data retrieval call binding the contract method 0x09eeebb0.
//
// Solidity: function getGamblers() view returns(address[])
func (_Betting *BettingCallerSession) GetGamblers() ([]common.Address, error) {
	return _Betting.Contract.GetGamblers(&_Betting.CallOpts)
}

// GetOutcomes is a free data retrieval call binding the contract method 0x3d3db8eb.
//
// Solidity: function getOutcomes() view returns(bytes32[])
func (_Betting *BettingCaller) GetOutcomes(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _Betting.contract.Call(opts, &out, "getOutcomes")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetOutcomes is a free data retrieval call binding the contract method 0x3d3db8eb.
//
// Solidity: function getOutcomes() view returns(bytes32[])
func (_Betting *BettingSession) GetOutcomes() ([][32]byte, error) {
	return _Betting.Contract.GetOutcomes(&_Betting.CallOpts)
}

// GetOutcomes is a free data retrieval call binding the contract method 0x3d3db8eb.
//
// Solidity: function getOutcomes() view returns(bytes32[])
func (_Betting *BettingCallerSession) GetOutcomes() ([][32]byte, error) {
	return _Betting.Contract.GetOutcomes(&_Betting.CallOpts)
}

// GetWinners is a free data retrieval call binding the contract method 0xdf15c37e.
//
// Solidity: function getWinners() view returns(address[])
func (_Betting *BettingCaller) GetWinners(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Betting.contract.Call(opts, &out, "getWinners")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetWinners is a free data retrieval call binding the contract method 0xdf15c37e.
//
// Solidity: function getWinners() view returns(address[])
func (_Betting *BettingSession) GetWinners() ([]common.Address, error) {
	return _Betting.Contract.GetWinners(&_Betting.CallOpts)
}

// GetWinners is a free data retrieval call binding the contract method 0xdf15c37e.
//
// Solidity: function getWinners() view returns(address[])
func (_Betting *BettingCallerSession) GetWinners() ([]common.Address, error) {
	return _Betting.Contract.GetWinners(&_Betting.CallOpts)
}

// IsGambler is a free data retrieval call binding the contract method 0xa7fe401d.
//
// Solidity: function isGambler(address ) view returns(bool)
func (_Betting *BettingCaller) IsGambler(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Betting.contract.Call(opts, &out, "isGambler", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsGambler is a free data retrieval call binding the contract method 0xa7fe401d.
//
// Solidity: function isGambler(address ) view returns(bool)
func (_Betting *BettingSession) IsGambler(arg0 common.Address) (bool, error) {
	return _Betting.Contract.IsGambler(&_Betting.CallOpts, arg0)
}

// IsGambler is a free data retrieval call binding the contract method 0xa7fe401d.
//
// Solidity: function isGambler(address ) view returns(bool)
func (_Betting *BettingCallerSession) IsGambler(arg0 common.Address) (bool, error) {
	return _Betting.Contract.IsGambler(&_Betting.CallOpts, arg0)
}

// IsOracle is a free data retrieval call binding the contract method 0xa97e5c93.
//
// Solidity: function isOracle(address _oracle) view returns(bool)
func (_Betting *BettingCaller) IsOracle(opts *bind.CallOpts, _oracle common.Address) (bool, error) {
	var out []interface{}
	err := _Betting.contract.Call(opts, &out, "isOracle", _oracle)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOracle is a free data retrieval call binding the contract method 0xa97e5c93.
//
// Solidity: function isOracle(address _oracle) view returns(bool)
func (_Betting *BettingSession) IsOracle(_oracle common.Address) (bool, error) {
	return _Betting.Contract.IsOracle(&_Betting.CallOpts, _oracle)
}

// IsOracle is a free data retrieval call binding the contract method 0xa97e5c93.
//
// Solidity: function isOracle(address _oracle) view returns(bool)
func (_Betting *BettingCallerSession) IsOracle(_oracle common.Address) (bool, error) {
	return _Betting.Contract.IsOracle(&_Betting.CallOpts, _oracle)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Betting *BettingCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Betting.contract.Call(opts, &out, "oracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Betting *BettingSession) Oracle() (common.Address, error) {
	return _Betting.Contract.Oracle(&_Betting.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Betting *BettingCallerSession) Oracle() (common.Address, error) {
	return _Betting.Contract.Oracle(&_Betting.CallOpts)
}

// OutcomeBets is a free data retrieval call binding the contract method 0x27edd6cf.
//
// Solidity: function outcomeBets(bytes32 ) view returns(uint256)
func (_Betting *BettingCaller) OutcomeBets(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Betting.contract.Call(opts, &out, "outcomeBets", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OutcomeBets is a free data retrieval call binding the contract method 0x27edd6cf.
//
// Solidity: function outcomeBets(bytes32 ) view returns(uint256)
func (_Betting *BettingSession) OutcomeBets(arg0 [32]byte) (*big.Int, error) {
	return _Betting.Contract.OutcomeBets(&_Betting.CallOpts, arg0)
}

// OutcomeBets is a free data retrieval call binding the contract method 0x27edd6cf.
//
// Solidity: function outcomeBets(bytes32 ) view returns(uint256)
func (_Betting *BettingCallerSession) OutcomeBets(arg0 [32]byte) (*big.Int, error) {
	return _Betting.Contract.OutcomeBets(&_Betting.CallOpts, arg0)
}

// Outcomes is a free data retrieval call binding the contract method 0xeed2a147.
//
// Solidity: function outcomes(uint256 ) view returns(bytes32)
func (_Betting *BettingCaller) Outcomes(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Betting.contract.Call(opts, &out, "outcomes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Outcomes is a free data retrieval call binding the contract method 0xeed2a147.
//
// Solidity: function outcomes(uint256 ) view returns(bytes32)
func (_Betting *BettingSession) Outcomes(arg0 *big.Int) ([32]byte, error) {
	return _Betting.Contract.Outcomes(&_Betting.CallOpts, arg0)
}

// Outcomes is a free data retrieval call binding the contract method 0xeed2a147.
//
// Solidity: function outcomes(uint256 ) view returns(bytes32)
func (_Betting *BettingCallerSession) Outcomes(arg0 *big.Int) ([32]byte, error) {
	return _Betting.Contract.Outcomes(&_Betting.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Betting *BettingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Betting.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Betting *BettingSession) Owner() (common.Address, error) {
	return _Betting.Contract.Owner(&_Betting.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Betting *BettingCallerSession) Owner() (common.Address, error) {
	return _Betting.Contract.Owner(&_Betting.CallOpts)
}

// ValidOutcomes is a free data retrieval call binding the contract method 0x914cf4f5.
//
// Solidity: function validOutcomes(bytes32 ) view returns(bool)
func (_Betting *BettingCaller) ValidOutcomes(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Betting.contract.Call(opts, &out, "validOutcomes", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidOutcomes is a free data retrieval call binding the contract method 0x914cf4f5.
//
// Solidity: function validOutcomes(bytes32 ) view returns(bool)
func (_Betting *BettingSession) ValidOutcomes(arg0 [32]byte) (bool, error) {
	return _Betting.Contract.ValidOutcomes(&_Betting.CallOpts, arg0)
}

// ValidOutcomes is a free data retrieval call binding the contract method 0x914cf4f5.
//
// Solidity: function validOutcomes(bytes32 ) view returns(bool)
func (_Betting *BettingCallerSession) ValidOutcomes(arg0 [32]byte) (bool, error) {
	return _Betting.Contract.ValidOutcomes(&_Betting.CallOpts, arg0)
}

// Winners is a free data retrieval call binding the contract method 0xa2fb1175.
//
// Solidity: function winners(uint256 ) view returns(address)
func (_Betting *BettingCaller) Winners(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Betting.contract.Call(opts, &out, "winners", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Winners is a free data retrieval call binding the contract method 0xa2fb1175.
//
// Solidity: function winners(uint256 ) view returns(address)
func (_Betting *BettingSession) Winners(arg0 *big.Int) (common.Address, error) {
	return _Betting.Contract.Winners(&_Betting.CallOpts, arg0)
}

// Winners is a free data retrieval call binding the contract method 0xa2fb1175.
//
// Solidity: function winners(uint256 ) view returns(address)
func (_Betting *BettingCallerSession) Winners(arg0 *big.Int) (common.Address, error) {
	return _Betting.Contract.Winners(&_Betting.CallOpts, arg0)
}

// Wins is a free data retrieval call binding the contract method 0x8047a97a.
//
// Solidity: function wins(address ) view returns(uint256)
func (_Betting *BettingCaller) Wins(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Betting.contract.Call(opts, &out, "wins", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wins is a free data retrieval call binding the contract method 0x8047a97a.
//
// Solidity: function wins(address ) view returns(uint256)
func (_Betting *BettingSession) Wins(arg0 common.Address) (*big.Int, error) {
	return _Betting.Contract.Wins(&_Betting.CallOpts, arg0)
}

// Wins is a free data retrieval call binding the contract method 0x8047a97a.
//
// Solidity: function wins(address ) view returns(uint256)
func (_Betting *BettingCallerSession) Wins(arg0 common.Address) (*big.Int, error) {
	return _Betting.Contract.Wins(&_Betting.CallOpts, arg0)
}

// ChooseOracle is a paid mutator transaction binding the contract method 0xd0aa7e25.
//
// Solidity: function chooseOracle(address newOracle) returns()
func (_Betting *BettingTransactor) ChooseOracle(opts *bind.TransactOpts, newOracle common.Address) (*types.Transaction, error) {
	return _Betting.contract.Transact(opts, "chooseOracle", newOracle)
}

// ChooseOracle is a paid mutator transaction binding the contract method 0xd0aa7e25.
//
// Solidity: function chooseOracle(address newOracle) returns()
func (_Betting *BettingSession) ChooseOracle(newOracle common.Address) (*types.Transaction, error) {
	return _Betting.Contract.ChooseOracle(&_Betting.TransactOpts, newOracle)
}

// ChooseOracle is a paid mutator transaction binding the contract method 0xd0aa7e25.
//
// Solidity: function chooseOracle(address newOracle) returns()
func (_Betting *BettingTransactorSession) ChooseOracle(newOracle common.Address) (*types.Transaction, error) {
	return _Betting.Contract.ChooseOracle(&_Betting.TransactOpts, newOracle)
}

// ContractReset is a paid mutator transaction binding the contract method 0xa7a5ce88.
//
// Solidity: function contractReset() returns()
func (_Betting *BettingTransactor) ContractReset(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Betting.contract.Transact(opts, "contractReset")
}

// ContractReset is a paid mutator transaction binding the contract method 0xa7a5ce88.
//
// Solidity: function contractReset() returns()
func (_Betting *BettingSession) ContractReset() (*types.Transaction, error) {
	return _Betting.Contract.ContractReset(&_Betting.TransactOpts)
}

// ContractReset is a paid mutator transaction binding the contract method 0xa7a5ce88.
//
// Solidity: function contractReset() returns()
func (_Betting *BettingTransactorSession) ContractReset() (*types.Transaction, error) {
	return _Betting.Contract.ContractReset(&_Betting.TransactOpts)
}

// MakeBet is a paid mutator transaction binding the contract method 0x284152dc.
//
// Solidity: function makeBet(bytes32 outcome) payable returns()
func (_Betting *BettingTransactor) MakeBet(opts *bind.TransactOpts, outcome [32]byte) (*types.Transaction, error) {
	return _Betting.contract.Transact(opts, "makeBet", outcome)
}

// MakeBet is a paid mutator transaction binding the contract method 0x284152dc.
//
// Solidity: function makeBet(bytes32 outcome) payable returns()
func (_Betting *BettingSession) MakeBet(outcome [32]byte) (*types.Transaction, error) {
	return _Betting.Contract.MakeBet(&_Betting.TransactOpts, outcome)
}

// MakeBet is a paid mutator transaction binding the contract method 0x284152dc.
//
// Solidity: function makeBet(bytes32 outcome) payable returns()
func (_Betting *BettingTransactorSession) MakeBet(outcome [32]byte) (*types.Transaction, error) {
	return _Betting.Contract.MakeBet(&_Betting.TransactOpts, outcome)
}

// MakeDecision is a paid mutator transaction binding the contract method 0xb8ecdb24.
//
// Solidity: function makeDecision(bytes32 decidedOutcome) returns()
func (_Betting *BettingTransactor) MakeDecision(opts *bind.TransactOpts, decidedOutcome [32]byte) (*types.Transaction, error) {
	return _Betting.contract.Transact(opts, "makeDecision", decidedOutcome)
}

// MakeDecision is a paid mutator transaction binding the contract method 0xb8ecdb24.
//
// Solidity: function makeDecision(bytes32 decidedOutcome) returns()
func (_Betting *BettingSession) MakeDecision(decidedOutcome [32]byte) (*types.Transaction, error) {
	return _Betting.Contract.MakeDecision(&_Betting.TransactOpts, decidedOutcome)
}

// MakeDecision is a paid mutator transaction binding the contract method 0xb8ecdb24.
//
// Solidity: function makeDecision(bytes32 decidedOutcome) returns()
func (_Betting *BettingTransactorSession) MakeDecision(decidedOutcome [32]byte) (*types.Transaction, error) {
	return _Betting.Contract.MakeDecision(&_Betting.TransactOpts, decidedOutcome)
}

// SetOutcomes is a paid mutator transaction binding the contract method 0x6646586f.
//
// Solidity: function setOutcomes(bytes32[] _outcomes) returns()
func (_Betting *BettingTransactor) SetOutcomes(opts *bind.TransactOpts, _outcomes [][32]byte) (*types.Transaction, error) {
	return _Betting.contract.Transact(opts, "setOutcomes", _outcomes)
}

// SetOutcomes is a paid mutator transaction binding the contract method 0x6646586f.
//
// Solidity: function setOutcomes(bytes32[] _outcomes) returns()
func (_Betting *BettingSession) SetOutcomes(_outcomes [][32]byte) (*types.Transaction, error) {
	return _Betting.Contract.SetOutcomes(&_Betting.TransactOpts, _outcomes)
}

// SetOutcomes is a paid mutator transaction binding the contract method 0x6646586f.
//
// Solidity: function setOutcomes(bytes32[] _outcomes) returns()
func (_Betting *BettingTransactorSession) SetOutcomes(_outcomes [][32]byte) (*types.Transaction, error) {
	return _Betting.Contract.SetOutcomes(&_Betting.TransactOpts, _outcomes)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Betting *BettingTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Betting.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Betting *BettingSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Betting.Contract.Withdraw(&_Betting.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Betting *BettingTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Betting.Contract.Withdraw(&_Betting.TransactOpts, amount)
}

// BettingBetMadeIterator is returned from FilterBetMade and is used to iterate over the raw logs and unpacked data for BetMade events raised by the Betting contract.
type BettingBetMadeIterator struct {
	Event *BettingBetMade // Event containing the contract specifics and raw log

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
func (it *BettingBetMadeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BettingBetMade)
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
		it.Event = new(BettingBetMade)
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
func (it *BettingBetMadeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BettingBetMadeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BettingBetMade represents a BetMade event raised by the Betting contract.
type BettingBetMade struct {
	Gambler common.Address
	Outcome [32]byte
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBetMade is a free log retrieval operation binding the contract event 0xf8e19b8ac4065a7305ef78373f54d2b12c631e1c04c0918be0342ac63f61c033.
//
// Solidity: event BetMade(address indexed gambler, bytes32 indexed outcome, uint256 amount)
func (_Betting *BettingFilterer) FilterBetMade(opts *bind.FilterOpts, gambler []common.Address, outcome [][32]byte) (*BettingBetMadeIterator, error) {

	var gamblerRule []interface{}
	for _, gamblerItem := range gambler {
		gamblerRule = append(gamblerRule, gamblerItem)
	}
	var outcomeRule []interface{}
	for _, outcomeItem := range outcome {
		outcomeRule = append(outcomeRule, outcomeItem)
	}

	logs, sub, err := _Betting.contract.FilterLogs(opts, "BetMade", gamblerRule, outcomeRule)
	if err != nil {
		return nil, err
	}
	return &BettingBetMadeIterator{contract: _Betting.contract, event: "BetMade", logs: logs, sub: sub}, nil
}

// WatchBetMade is a free log subscription operation binding the contract event 0xf8e19b8ac4065a7305ef78373f54d2b12c631e1c04c0918be0342ac63f61c033.
//
// Solidity: event BetMade(address indexed gambler, bytes32 indexed outcome, uint256 amount)
func (_Betting *BettingFilterer) WatchBetMade(opts *bind.WatchOpts, sink chan<- *BettingBetMade, gambler []common.Address, outcome [][32]byte) (event.Subscription, error) {

	var gamblerRule []interface{}
	for _, gamblerItem := range gambler {
		gamblerRule = append(gamblerRule, gamblerItem)
	}
	var outcomeRule []interface{}
	for _, outcomeItem := range outcome {
		outcomeRule = append(outcomeRule, outcomeItem)
	}

	logs, sub, err := _Betting.contract.WatchLogs(opts, "BetMade", gamblerRule, outcomeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BettingBetMade)
				if err := _Betting.contract.UnpackLog(event, "BetMade", log); err != nil {
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

// ParseBetMade is a log parse operation binding the contract event 0xf8e19b8ac4065a7305ef78373f54d2b12c631e1c04c0918be0342ac63f61c033.
//
// Solidity: event BetMade(address indexed gambler, bytes32 indexed outcome, uint256 amount)
func (_Betting *BettingFilterer) ParseBetMade(log types.Log) (*BettingBetMade, error) {
	event := new(BettingBetMade)
	if err := _Betting.contract.UnpackLog(event, "BetMade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BettingOracleChangedIterator is returned from FilterOracleChanged and is used to iterate over the raw logs and unpacked data for OracleChanged events raised by the Betting contract.
type BettingOracleChangedIterator struct {
	Event *BettingOracleChanged // Event containing the contract specifics and raw log

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
func (it *BettingOracleChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BettingOracleChanged)
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
		it.Event = new(BettingOracleChanged)
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
func (it *BettingOracleChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BettingOracleChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BettingOracleChanged represents a OracleChanged event raised by the Betting contract.
type BettingOracleChanged struct {
	PreviousOracle common.Address
	NewOracle      common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOracleChanged is a free log retrieval operation binding the contract event 0x05cd89403c6bdeac21c2ff33de395121a31fa1bc2bf3adf4825f1f86e79969dd.
//
// Solidity: event OracleChanged(address indexed previousOracle, address indexed newOracle)
func (_Betting *BettingFilterer) FilterOracleChanged(opts *bind.FilterOpts, previousOracle []common.Address, newOracle []common.Address) (*BettingOracleChangedIterator, error) {

	var previousOracleRule []interface{}
	for _, previousOracleItem := range previousOracle {
		previousOracleRule = append(previousOracleRule, previousOracleItem)
	}
	var newOracleRule []interface{}
	for _, newOracleItem := range newOracle {
		newOracleRule = append(newOracleRule, newOracleItem)
	}

	logs, sub, err := _Betting.contract.FilterLogs(opts, "OracleChanged", previousOracleRule, newOracleRule)
	if err != nil {
		return nil, err
	}
	return &BettingOracleChangedIterator{contract: _Betting.contract, event: "OracleChanged", logs: logs, sub: sub}, nil
}

// WatchOracleChanged is a free log subscription operation binding the contract event 0x05cd89403c6bdeac21c2ff33de395121a31fa1bc2bf3adf4825f1f86e79969dd.
//
// Solidity: event OracleChanged(address indexed previousOracle, address indexed newOracle)
func (_Betting *BettingFilterer) WatchOracleChanged(opts *bind.WatchOpts, sink chan<- *BettingOracleChanged, previousOracle []common.Address, newOracle []common.Address) (event.Subscription, error) {

	var previousOracleRule []interface{}
	for _, previousOracleItem := range previousOracle {
		previousOracleRule = append(previousOracleRule, previousOracleItem)
	}
	var newOracleRule []interface{}
	for _, newOracleItem := range newOracle {
		newOracleRule = append(newOracleRule, newOracleItem)
	}

	logs, sub, err := _Betting.contract.WatchLogs(opts, "OracleChanged", previousOracleRule, newOracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BettingOracleChanged)
				if err := _Betting.contract.UnpackLog(event, "OracleChanged", log); err != nil {
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

// ParseOracleChanged is a log parse operation binding the contract event 0x05cd89403c6bdeac21c2ff33de395121a31fa1bc2bf3adf4825f1f86e79969dd.
//
// Solidity: event OracleChanged(address indexed previousOracle, address indexed newOracle)
func (_Betting *BettingFilterer) ParseOracleChanged(log types.Log) (*BettingOracleChanged, error) {
	event := new(BettingOracleChanged)
	if err := _Betting.contract.UnpackLog(event, "OracleChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BettingWinnersIterator is returned from FilterWinners and is used to iterate over the raw logs and unpacked data for Winners events raised by the Betting contract.
type BettingWinnersIterator struct {
	Event *BettingWinners // Event containing the contract specifics and raw log

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
func (it *BettingWinnersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BettingWinners)
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
		it.Event = new(BettingWinners)
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
func (it *BettingWinnersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BettingWinnersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BettingWinners represents a Winners event raised by the Betting contract.
type BettingWinners struct {
	Wins       []common.Address
	TotalPrize *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWinners is a free log retrieval operation binding the contract event 0x413cf1e17d70833c017de9b3c64d0104e3e37a7191d95b3d0d6276de992a0b4d.
//
// Solidity: event Winners(address[] wins, uint256 totalPrize)
func (_Betting *BettingFilterer) FilterWinners(opts *bind.FilterOpts) (*BettingWinnersIterator, error) {

	logs, sub, err := _Betting.contract.FilterLogs(opts, "Winners")
	if err != nil {
		return nil, err
	}
	return &BettingWinnersIterator{contract: _Betting.contract, event: "Winners", logs: logs, sub: sub}, nil
}

// WatchWinners is a free log subscription operation binding the contract event 0x413cf1e17d70833c017de9b3c64d0104e3e37a7191d95b3d0d6276de992a0b4d.
//
// Solidity: event Winners(address[] wins, uint256 totalPrize)
func (_Betting *BettingFilterer) WatchWinners(opts *bind.WatchOpts, sink chan<- *BettingWinners) (event.Subscription, error) {

	logs, sub, err := _Betting.contract.WatchLogs(opts, "Winners")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BettingWinners)
				if err := _Betting.contract.UnpackLog(event, "Winners", log); err != nil {
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

// ParseWinners is a log parse operation binding the contract event 0x413cf1e17d70833c017de9b3c64d0104e3e37a7191d95b3d0d6276de992a0b4d.
//
// Solidity: event Winners(address[] wins, uint256 totalPrize)
func (_Betting *BettingFilterer) ParseWinners(log types.Log) (*BettingWinners, error) {
	event := new(BettingWinners)
	if err := _Betting.contract.UnpackLog(event, "Winners", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BettingWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the Betting contract.
type BettingWithdrawnIterator struct {
	Event *BettingWithdrawn // Event containing the contract specifics and raw log

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
func (it *BettingWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BettingWithdrawn)
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
		it.Event = new(BettingWithdrawn)
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
func (it *BettingWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BettingWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BettingWithdrawn represents a Withdrawn event raised by the Betting contract.
type BettingWithdrawn struct {
	Gambler common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed gambler, uint256 amount)
func (_Betting *BettingFilterer) FilterWithdrawn(opts *bind.FilterOpts, gambler []common.Address) (*BettingWithdrawnIterator, error) {

	var gamblerRule []interface{}
	for _, gamblerItem := range gambler {
		gamblerRule = append(gamblerRule, gamblerItem)
	}

	logs, sub, err := _Betting.contract.FilterLogs(opts, "Withdrawn", gamblerRule)
	if err != nil {
		return nil, err
	}
	return &BettingWithdrawnIterator{contract: _Betting.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed gambler, uint256 amount)
func (_Betting *BettingFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *BettingWithdrawn, gambler []common.Address) (event.Subscription, error) {

	var gamblerRule []interface{}
	for _, gamblerItem := range gambler {
		gamblerRule = append(gamblerRule, gamblerItem)
	}

	logs, sub, err := _Betting.contract.WatchLogs(opts, "Withdrawn", gamblerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BettingWithdrawn)
				if err := _Betting.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed gambler, uint256 amount)
func (_Betting *BettingFilterer) ParseWithdrawn(log types.Log) (*BettingWithdrawn, error) {
	event := new(BettingWithdrawn)
	if err := _Betting.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
