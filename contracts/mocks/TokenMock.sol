pragma solidity ^0.5.8;

import "../../lib/oz/contracts/token/ERC20/ERC20.sol";
import "../../lib/oz/contracts/token/ERC20/ERC20Detailed.sol";

contract TokenMock is ERC20, ERC20Detailed {
    constructor () public ERC20Detailed("TokenMock", "MOCK", 18) {
        _mint(msg.sender, 10000 * (10 ** uint256(decimals())));
    }
}
