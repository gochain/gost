pragma solidity ^0.5.8;

import "../AddressSet.sol";

// AddressSetMock is a test contract.
contract AddressSetMock {
    AddressSet.Set set;

    function addAll(address[] memory init) public {
        AddressSet.addAll(set, init);
    }

    function add(address addr) public {
        AddressSet.add(set, addr);
    }

    function remove(address addr) public {
        AddressSet.remove(set, addr);
    }

    function majority() public view returns (uint) {
        return AddressSet.majority(set);
    }

    function has(address addr) public view returns (bool) {
        return set.map[addr] > 0;
    }

    function length() public view returns (uint) {
        return set.list.length;
    }
}