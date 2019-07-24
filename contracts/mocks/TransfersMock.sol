pragma solidity ^0.5.8;

import "../Transfers.sol";

contract TransfersMock is Transfers {
    constructor(address confirmationsContract) Transfers(confirmationsContract) public { }

    function claimTransfer(uint blockNum, uint logIndex, address addr, uint amount) public {
        _claimTransfer(blockNum, logIndex, addr, amount);
    }

    function emitEvent(address addr, uint amount) public {
        emit GOSTTransfer(addr, amount);
    }
}