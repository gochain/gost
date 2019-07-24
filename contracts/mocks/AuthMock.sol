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

    function signersLength() public view returns (uint) {
        return signers.list.length;
    }

    function getSigner(uint i) public view returns (address) {
        return signers.list[i];
    }

    function votersLength() public view returns (uint) {
        return voters.list.length;
    }

    function getVoter(uint i) public view returns (address) {
        return voters.list[i];
    }
}
