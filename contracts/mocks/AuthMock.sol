pragma solidity ^0.5.8;

import "../Auth.sol";

// AuthMock is a test contract which completes the abstract Auth contract.
contract AuthMock is Auth {
    constructor(address voter) Auth(voter) public {}

    event SignerRemoved(address signer);

    function onRemoveSigner(address signer) internal {
        emit SignerRemoved(signer);
    }

    function testOnlySigners() public view onlySigners returns (bool) {
        return true;
    }
}
