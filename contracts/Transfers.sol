pragma solidity ^0.5.8;

import "./Confirmations.sol";

// Transfers manages confirmations and claims for GOSTTransfer events.
contract Transfers {
    // Confirmations from another chain.
    IConfirmations confirmations;
    // The contract on the other network which emits the GOSTTransfer events.
    address public transferSourceContract;
    event GOSTTransfer(address indexed addr, uint amount);
    bytes32 public eventID = 0xda2543cfac858777ecd10b22ad9db15a5d6fc0eb2abd01e4dbbb30f012bb168c;
    // GOSTTransfer events which have been claimed.
    // blockNum=>logIndex=>eventHash
    mapping(uint=>mapping(uint=>mapping(bytes32=>bool))) public claimed;

    constructor(address confirmationsContract) public {
        confirmations = IConfirmations(confirmationsContract);
    }

    //TODO onlyOwner
    function setTransferSourceContract(address _transferContract) public {
        require(transferSourceContract == address(0), "transferContract already set");
        transferSourceContract = _transferContract;
    }

    // Claim a confirmed GOSTTransfer event.
    // Event must be confirmed and not already claimed.
    function _claimTransfer(uint blockNum, uint logIndex, address addr, uint amount) internal {
        bytes32 eh = transferEventHash(addr, amount);
        require(!claimed[blockNum][logIndex][eh], "Already claimed");
        require(confirmations.status(blockNum, logIndex, eh) == IConfirmations.Status.Confirmed,
            "Transfer must be confirmed");
        claimed[blockNum][logIndex][eh] = true;
    }

    // Compute the event hash for a GOSTTransfer event from the source contract.
    function transferEventHash(address addr, uint amount) public view returns (bytes32) {
        return keccak256(abi.encode(transferSourceContract, eventID, addr, amount));
    }
}
