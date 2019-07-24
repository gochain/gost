pragma solidity ^0.5.8;

import "../lib/oz/contracts/token/ERC20/IERC20.sol";
import "./Transfers.sol";

// A Custody contract manages token transfers on the original network.
// Tokens transferred to the proxy network are held frozen, and unlocked when return transfers are confirmed.
contract Custody is Transfers {
    IERC20 token;

    constructor(address tokenContract, address confirmationsContract) Transfers(confirmationsContract) public {
        token = IERC20(tokenContract);
    }

    // Initiates a transfer to the proxy network by freezing tokens in this
    // contract and emitting a confirmable GOSTTransfer event.
    // The caller must first allow the contract to transferFrom amount.
    function proxyTransferTo(uint amount, address to) public {
        require(token.transferFrom(msg.sender, address(this), amount));
        emit GOSTTransfer(to, amount);
    }

    function proxyTransfer(uint amount) public {
        proxyTransferTo(amount, msg.sender);
    }

    // Claims an incoming transfer from the proxy network and unlocks the tokens for the caller.
    // Event must be confirmed and not already claimed.
    function claimTransfer(uint blockNum, uint logIndex, uint amount) public {
        _claimTransfer(blockNum, logIndex, msg.sender, amount);
        token.transfer(msg.sender, amount);
    }
}
