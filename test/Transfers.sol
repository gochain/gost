pragma solidity ^0.5.8;

import "truffle/Assert.sol";
import "truffle/DeployedAddresses.sol";
import "../contracts/Transfers.sol";

contract TestTransfers {
    address constant source = 0x0000000000000000000000000000000000001234;
    Transfers transfers = new Transfers(0x0000000000000000000000000000000000000000);

    function beforeEach() public {
        transfers.setTransferSourceContract(source);
    }

    function testEventHash() public {
        testEventHash(0x0000000000000000000000000000000000000001, 1,
            0x2ed084d268563f6ebd1105b7cc07acf35eddd49baa175c1672a1cd0b20e5d1b7);
        testEventHash(0x000000000000000000000000000000000000000F, 10,
            0x8f61e6180ce2c962b15c7d01e7dd21a66da0f17d8b8251af39b25a969e8cca88);
        testEventHash(0x000000000000000000000000000000000000000A, 1000,
            0x907a2100da50a80ba06717a00d452d097fdb5d831b1b0117f923574edcc8c33c);
        testEventHash(0x0000000000000000000000000000000000000009, 1000000000000000000,
            0xaa1e87375db1dcda2293abe5db1dec2335610f93cd58258509681551de6c3da0);
    }

    event EventHash(bytes32 got);
    function testEventHash(address addr, uint amount, bytes32 expect) internal {
        bytes32 hash = transfers.transferEventHash(addr, amount);
        emit EventHash(hash);
        Assert.equal(hash, expect, "Hashes should match");
    }
}
