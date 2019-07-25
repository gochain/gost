const TransfersMock = artifacts.require("TransfersMock");
const truffleAssert = require('truffle-assertions');

contract("Transfers", async accounts => {
    const source = "0x0000000000000000000000000000000000001234";
    let transfers;
    let eventID;

    beforeEach(async () => {
        transfers = await TransfersMock.new(accounts[0]);
        await transfers.setTransferSourceContract(source);
        eventID = await transfers.eventID();
    })

    it("javascript and solidity hashes match", async () => {
        let test = async function(addr, amount) {
            let sol = await transfers.transferEventHash(addr, amount);
            let js = web3.utils.keccak256(web3.eth.abi.encodeParameters(
                ['address', 'bytes32', 'address', 'uint256'],
                [source, eventID, addr, amount]
            ));
            assert.equal(sol, js, "Hashes should match");
        };

        await test(accounts[0], 1);
        await test(accounts[1], 10);
        await test(accounts[2], 1000);
    })

    it("eventID correct", async () => {
        let tx = await transfers.emitEvent(accounts[0], 1);

        truffleAssert.eventEmitted(tx, 'GOSTTransfer', (ev) => {
            return ev.addr === accounts[0] && ev.amount == 1;
        });
        assert.equal(tx.receipt.rawLogs[0].topics[0], eventID, "EventID constant mismatch");
    })

    //TODO test request/verify confirmation

    //TODO test claim, then reject second claim

    //TODO reject multiple setTransferSourceContract
})
