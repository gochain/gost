pragma solidity ^0.5.8;

// AddressSet manages a set of addresses which can be iterated over.
library AddressSet {
    struct Set {
        address[] list;
        mapping(address=>uint) map;
    }

    function addAll(Set storage set, address[] memory init) internal {
        for (uint i=0;i<init.length;i++){
            add(set, init[i]);
        }
    }

    function add(Set storage set, address addr) internal {
        require(set.map[addr] == 0, "Already in set");
        set.map[addr] = set.list.push(addr);
    }

    function remove(Set storage set, address addr) internal {
        uint r = set.map[addr];
        require(r > 0, "Not in set");
        delete set.map[addr];
        uint l = set.list.length;
        set.list[r-1] = set.list[l-1];
        delete set.list[l-1];
        set.list.length--;
    }

    function majority(Set storage set) public view returns (uint) {
        return (set.list.length/2)+1;
    }
}
