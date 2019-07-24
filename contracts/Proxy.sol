pragma solidity ^0.5.8;

import "../lib/oz/contracts/token/ERC20/ERC20Detailed.sol";
import "../lib/oz/contracts/token/ERC20/ERC20Burnable.sol";
import "../lib/oz/contracts/token/ERC20/ERC20Mintable.sol";
import "./Transfers.sol";

// A Proxy contract is an ERC20 which also manages token transfers on a proxy network.
// Tokens are minted when transfers from the original network are confirmed, and burned when returned.
// Every token corresponds to a frozen token on the original network.
contract Proxy is Transfers, ERC20Burnable, ERC20Detailed {
    constructor(address confirmationsContract, string memory name, string memory symbol, uint8 decimals) ERC20Detailed(name, symbol, decimals) Transfers(confirmationsContract) public { }

    // Initiates a transfer to the origin network by burning tokens
    // and emitting a confirmable GOSTTransfer event.
    function originTransferTo(uint amount, address to) public {
        _burn(msg.sender, amount);
        emit GOSTTransfer(to, amount);
    }

    function originTransfer(uint amount) public {
        originTransferTo(amount, msg.sender);
    }

    // Claims an incoming transfer and mints tokens for the caller.
    // Event must be confirmed and not yet claimed.
    function claimTransfer(uint blockNum, uint logIndex, uint amount) public {
        _claimTransfer(blockNum, logIndex, msg.sender, amount);
        _mint(msg.sender, amount);
    }
}
