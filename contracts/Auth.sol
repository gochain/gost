pragma solidity ^0.5.8;

import "./AddressSet.sol";

// Auth manages sets of addresses with voting and signing authorization, intended to mirror
// the state of GoChain.
contract Auth {
    AddressSet.Set voters;
    AddressSet.Set signers;
    struct Vote{
        address addr;
        bool voter; // or signer
        bool add; // or remove
    }
    //TODO need to establish a canonical order to update, since we only vote on one thing at a time...
    mapping(address=>Vote) public votes;
    mapping(address=>mapping(bool=>mapping(bool=>uint))) public count;

    constructor(address voter) public {
        AddressSet.add(voters, voter);
        AddressSet.add(signers, voter);
    }

    modifier onlyVoters() {
        require(
            voters.map[msg.sender] != 0,
            "Sender not a voter"
        );
        _;
    }

    modifier onlySigners() {
        require(
            signers.map[msg.sender] != 0,
            "Sender not a signer"
        );
        _;
    }

    // Cast a vote to update the signer or voter set. Overwrites any previous pending vote.
    // Passing an empty address will remove an existing vote.
    function setVote(address addr, bool voter, bool add) public onlyVoters {
        Vote storage vote = votes[msg.sender];
        if (vote.addr != address(0)) {
            count[vote.addr][vote.voter][vote.add]--;
            if (addr == address(0)) {
                // Delete existing.
                delete votes[msg.sender];
                return;
            }
        } else if (addr == address(0)) {
            // No existing vote to delete.
            return;
        }

        votes[msg.sender] = Vote(addr, voter, add);
        uint c = ++count[addr][voter][add];
        if (c >= AddressSet.majority(voters)) {
            _process(addr, voter, add);
        }
    }

    // Removes a voter and its vote, then checks for consensus.
    function removeVoter(address voter) internal {
        AddressSet.remove(voters, voter);
        Vote storage vote = votes[voter];
        if (vote.addr != address(0)) {
            count[vote.addr][vote.voter][vote.add]--;
            delete votes[voter];
        }

        // Check for consensus.
        uint majority = AddressSet.majority(voters);
        for (uint i=0;i<voters.list.length;i++) {
            vote = votes[voters.list[i]];
            if (vote.addr != address(0)) {
                if (count[vote.addr][vote.voter][vote.add] >= majority) {
                    // Consensus.
                    _process(vote.addr, vote.voter, vote.add);
                    return;
                }
            }
        }
    }

    // Process a vote which has majority consensus.
    function _process(address addr, bool voter, bool add) internal {
        // Remove matching votes.
        for (uint i=0;i<voters.list.length;i++) {
            address v = voters.list[i];
            Vote storage vote = votes[v];
            if (vote.addr == addr && vote.voter == voter && vote.add == add) {
                delete votes[v];
            }
        }
        // Delete count.
        delete count[addr][voter][add];

        // Execute change.
        if (voter && add) {
            // Ensure signer.
            if (signers.map[addr] == 0) {
                AddressSet.add(signers, addr);
            }
            // Ensure voter.
            if (voters.map[addr] == 0) {
                AddressSet.add(voters, addr);
            }
        } else if (voter && !add) {
            // Ensure not voter.
            if (voters.map[addr] != 0) {
                removeVoter(addr);
            }
        } else if (!voter && add) {
            // Ensure signer.
            if (signers.map[addr] == 0) {
                AddressSet.add(signers, addr);
            }
        } else if (!voter && !add) {
            // Ensure not voter.
            if (voters.map[addr] != 0) {
                removeVoter(addr);
            }
            // Ensure not signer.
            if (signers.map[addr] != 0) {
                AddressSet.remove(signers, addr);
                onRemoveSigner(addr);
            }
        }
    }

    // Called when a signer is removed
    function onRemoveSigner(address signer) internal;
}
