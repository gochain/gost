pragma solidity ^0.5.8;

import "./Auth.sol";

// The IConfirmations interface contains public methods for confirming events from other networks.
interface IConfirmations {
    // The status of an event is None before confirmation is requested, and Pending after.
    // Once a majority consensus is reached, it will transition to Invalid or Confirmed, both of which are final.
    enum Status { None, Pending, Invalid, Confirmed }

    // Get the status for an event.
    function status(uint blockNum, uint logIndex, bytes32 eventHash) external view returns (Status);

    // Request confirmation of an event. Status must be None or Pending.
    // Emits a ConfirmationRequested event and results in a Pending status.
    // Must include payment of 2,500,000 gas @this tx's gas price to cover signers' confirmation gas fees.
    function request(uint blockNum, uint logIndex, bytes32 eventHash) external payable;

    // Check the status of a pending request.
    // Returns total signer count, valid/invalid vote counts, and the gasPrice incentive.
    function pendingStatus(uint blockNum, uint logIndex, bytes32 eventHash) external view returns (uint gasPrice, uint total, uint valid, uint invalid);

    // Emitted when confirmation of an event is requested.
    event ConfirmationRequested(uint indexed blockNum, uint indexed logIndex, bytes32 eventHash);

    // Emitted when an event is confirmed by a majority of signers to be either valid or invalid.
    event Confirmed(uint indexed blockNum, uint indexed logIndex, bytes32 eventHash, bool valid);
}

// The Confirmations contracts implements IConfirmations, and includes administrative methods
// for GoChain signers to manage the signer and voter sets, and to confirm pending requests.
contract Confirmations is Auth, IConfirmations {
    // Casting a confirmation vote with confirm() uses a little less than 90000 gas.
    uint constant public confirmGas = 90000;
    // Once consensus is reached, _confirmed() uses a little less than 152000 gas to
    // update the state and issue payouts to the confirmers.
    uint constant public _confirmedGas = 160000;
    // Roughly the total gas used to confirm by a complete set of signers (26/50).
    // Used in combination with the gas price to determine the required fee, which is
    // paid by the requester and distributed to confirming signers.
    uint constant public totalConfirmGas = 26 * confirmGas + _confirmedGas;

    // Uniquely identifies an event. Corresponds to map keys as well.
    struct Key {
        uint blockNum;
        uint logIndex;
        bytes32 eventHash;
    }

    // Event statuses.
    mapping(uint=>mapping(uint=>mapping(bytes32=>IConfirmations.Status))) public status;

    // Pending confirmation.
    struct Pending {
        AddressSet.Set validVotes;
        AddressSet.Set invalidVotes;
        // Set by confirmation requester, used by signers to confirm.
        uint gasPrice;
    }

    // Events pending confirmation.
    mapping(uint=>mapping(uint=>mapping(bytes32=>Pending))) pending;
    // Iterable list of keys from pending.
    Key[] public pendingList;
    // Index + 1 from pendingList.
    mapping(uint=>mapping(uint=>mapping(bytes32=>uint))) pendingIdx;

    constructor(address voter) Auth(voter) public {}

    function request(uint blockNum, uint logIndex, bytes32 eventHash) public payable {
        Status s = status[blockNum][logIndex][eventHash];
        require(s == Status.None || s == Status.Pending, "Status must be None or Pending");

        uint fee = totalConfirmGas * tx.gasprice;
        require(msg.value >= fee, "Tx value doesn't cover confirmation fee");
        if (msg.value > fee) {
            // Refund overpayment.
            msg.sender.transfer(msg.value - fee);
        }

        if (s == Status.None) {
            status[blockNum][logIndex][eventHash] = Status.Pending;
            uint l = pendingList.push(Key(blockNum, logIndex, eventHash));
            pendingIdx[blockNum][logIndex][eventHash] = l;
        }

        pending[blockNum][logIndex][eventHash].gasPrice += tx.gasprice;

        emit ConfirmationRequested(blockNum, logIndex, eventHash);
    }

    function pendingStatus(uint blockNum, uint logIndex, bytes32 eventHash) public view returns (uint gasPrice, uint total, uint valid, uint invalid) {
        if (status[blockNum][logIndex][eventHash] != Status.Pending) {
            return (0, 0, 0, 0);
        }
        Pending storage p = pending[blockNum][logIndex][eventHash];
        return (p.gasPrice, signers.list.length, p.validVotes.list.length, p.invalidVotes.list.length);
    }

    // Returns true if the request is pending and the calling signer has not yet confirmed.
    function shouldConfirm(uint blockNum, uint logIndex, bytes32 eventHash) onlySigners public view returns (bool, uint) {
        if (status[blockNum][logIndex][eventHash] != Status.Pending) {
            return (false, 0);
        }
        Pending storage p = pending[blockNum][logIndex][eventHash];
        if (p.validVotes.map[msg.sender] > 0) {
            return (false, 0);
        }
        if (p.invalidVotes.map[msg.sender] > 0) {
            return (false, 0);
        }
        return (true, p.gasPrice);
    }

    // Vote to confirm an event. Only callable by GoChain signers. Status must be Pending. Gas price must match requested.
    // If this is the final vote, emits a Confirmed event and distributes fees.
    function confirm(uint blockNum, uint logIndex, bytes32 eventHash, bool valid) onlySigners public  {
        require(status[blockNum][logIndex][eventHash] == Status.Pending, "Status not Pending");

        Pending storage p = pending[blockNum][logIndex][eventHash];
        AddressSet.Set storage votes = p.validVotes;
        if (valid) {
            if (p.invalidVotes.map[msg.sender] > 0) {
                AddressSet.remove(p.invalidVotes, msg.sender);
            }
        } else {
            votes = p.invalidVotes;
            if (p.validVotes.map[msg.sender] > 0) {
                AddressSet.remove(p.validVotes, msg.sender);
            }
        }
        if (votes.map[msg.sender] == 0) {
            AddressSet.add(votes, msg.sender);
            if (votes.list.length < AddressSet.majority(signers)) {
                return;
            }
            _confirmed(blockNum, logIndex, eventHash, valid);
        }
    }

    // Executes a confirmation by updating state, distributing the fee, and emitting a Confirmed event.
    function _confirmed(uint blockNum, uint logIndex, bytes32 eventHash, bool valid) internal {
        // Update status.
        Status s;
        if (valid) {
            s = Status.Confirmed;
        } else {
            s = Status.Invalid;
        }
        status[blockNum][logIndex][eventHash] = s;

        // Distribute reward.
        Pending storage p = pending[blockNum][logIndex][eventHash];
        AddressSet.Set storage confirmers = p.validVotes;
        if (!valid) {
            confirmers = p.invalidVotes;
        }
        uint count = confirmers.list.length;
        uint totalConfirmFee = p.gasPrice * 26 * confirmGas;
        uint confirmFee = totalConfirmFee / count;
        for (uint i = 0;i<count;i++) {
            address payable c = address(uint160(confirmers.list[i]));
            c.transfer(confirmFee);
        }
        uint remainder = (totalConfirmFee - (count*confirmFee));
        uint _confirmedFee = _confirmedGas * p.gasPrice;
        msg.sender.transfer(_confirmedFee + remainder);

        // Remove pending state.
        delete pending[blockNum][logIndex][eventHash];
        uint i = pendingIdx[blockNum][logIndex][eventHash];
        delete pendingIdx[blockNum][logIndex][eventHash];
        uint l = pendingList.length;
        if (l > 1) {
            // Move last Key to removed's position.
            Key storage last = pendingList[l-1];
            pendingList[i-1] = last;
            pendingIdx[last.blockNum][last.logIndex][last.eventHash] = i;
        }
        // Trim last list entry.
        delete pendingList[l-1];
        pendingList.length--;

        emit Confirmed(blockNum, logIndex, eventHash, valid);
    }

    // When a signer is removed, its votes are also removed, and the remaining pending votes
    // are rechecked for consensus, since the threshold for majority may have decreased.
    function onRemoveSigner(address signer) internal {
        uint majority = AddressSet.majority(signers);
        for (uint i = 0; i < pendingList.length;) {
            Key storage key = pendingList[i];
            Pending storage p = pending[key.blockNum][key.logIndex][key.eventHash];
            if (p.validVotes.map[signer] > 0) {
                AddressSet.remove(p.validVotes, signer);
            } else if (p.invalidVotes.map[signer] > 0) {
                AddressSet.remove(p.invalidVotes, signer);
            }
            if (p.validVotes.list.length >= majority) {
                _confirmed(key.blockNum, key.logIndex, key.eventHash, true);
            } else if (p.invalidVotes.list.length >= majority) {
                _confirmed(key.blockNum, key.logIndex, key.eventHash, false);
            } else {
                // _confirmed() deletes the entry, so only increment when it's not called.
                i++;
            }
        }
    }

    function pendingListLength() public view returns (uint) {
        return pendingList.length;
    }
}
