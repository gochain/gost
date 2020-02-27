const AuthMock = artifacts.require("AuthMock");
const AddressSet = artifacts.require("AddressSet");
const truffleAssert = require('truffle-assertions');

contract("AuthMock", async accounts => {
    let auth;
    const firstVoter = accounts[0];
    let addrSetLib;

    beforeEach(async () => {
        addrSetLib = await AddressSet.new();
        await AuthMock.link("AddressSet", addrSetLib.address);
        auth = await AuthMock.new(firstVoter);
    })

    it("initialized properly", async () => {
        assert.equal(await auth.signersLength(), 1, "Expected 1 signer");
        assert.equal(await auth.getSigner(0), firstVoter, "Unexpected signer");

        assert.equal(await auth.votersLength(), 1, "Expected 1 signer");
        assert.equal(await auth.getVoter(0), firstVoter, "Unexpected voter");
    })

    it("reject non-voter votes", async () => {
        try {
            await auth.setVote(accounts[1],true,true, {from: accounts[1]});
            assert.fail("Non-voter voted");
        } catch (error) {
            assert(error.toString().includes("Sender not a voter"), "Expected different error");
        }
    })

    it("add and remove signers and voters", async() => {
        // Add 2nd signer.
        await auth.setVote(accounts[1], false, true);
        assert.equal(await auth.votersLength(), 1, "Expected 1 voter");
        assert.equal(await auth.signersLength(), 2, "Expected 2 signers");

        // Allow 2nd signer to vote.
        await auth.setVote(accounts[1], true, true);
        assert.equal(await auth.signersLength(), 2, "Expected 2 signers");
        assert.equal(await auth.votersLength(), 2, "Expected 2 voters");

        // Add 3rd signer (requires consensus).
        await auth.setVote(accounts[2], false, true);
        await auth.setVote(accounts[2], false, true, {from: accounts[1]});
        assert.equal(await auth.signersLength(), 3, "Expected 3 signers");
        assert.equal(await auth.votersLength(), 2, "Expected 2 voters");

        // Remove 3rd signer.
        await auth.setVote(accounts[2], false, false);
        let tx = await auth.setVote(accounts[2], false, false, {from: accounts[1]});
        truffleAssert.eventEmitted(tx, 'SignerRemoved', (ev) => {
            return ev.signer === accounts[2];
        });
        assert.equal(await auth.signersLength(), 2, "Expected 2 signers");
        assert.equal(await auth.votersLength(), 2, "Expected 2 voters");

        // Remove 2nd signer (self-vote).
        await auth.setVote(accounts[1], false, false);
        tx = await auth.setVote(accounts[1], false, false, {from: accounts[1]});
        truffleAssert.eventEmitted(tx, 'SignerRemoved', (ev) => {
            return ev.signer === accounts[1];
        });
        assert.equal(await auth.signersLength(), 1, "Expected 1 signer");
        assert.equal(await auth.votersLength(), 1, "Expected 1 voter");

        // Add back as voter.
        await auth.setVote(accounts[1], true, true);
        assert.equal(await auth.signersLength(), 2, "Expected 2 signers");
        assert.equal(await auth.votersLength(), 2, "Expected 2 voters");

        // Revoke 2nd signer's voting permissions.
        await auth.setVote(accounts[1], true, false);
        tx = await auth.setVote(accounts[1], true, false, {from: accounts[1]});
        assert.equal(await auth.signersLength(), 2, "Expected 2 signers");
        assert.equal(await auth.votersLength(), 1, "Expected 1 voter");
        truffleAssert.eventNotEmitted(tx, 'SignerRemoved');

        // Remove 2nd signer.
        tx = await auth.setVote(accounts[1], false, false);
        truffleAssert.eventEmitted(tx, 'SignerRemoved', (ev) => {
            return ev.signer === accounts[1];
        });
        assert.equal(await auth.signersLength(), 1, "Expected 1 signer");
        assert.equal(await auth.votersLength(), 1, "Expected 1 voter");
    })

    it("post-process on voter remove", async () => {
        // Vote to add 2nd and 3rd voter.
        await auth.setVote(accounts[1], true, true);
        await auth.setVote(accounts[2], true, true);
        await auth.setVote(accounts[2], true, true, {from: accounts[1]});
        assert.equal(await auth.votersLength(), 3, "Voters not added");

        // Cast vote from 3rd.
        await auth.setVote(accounts[3], true, true, {from: accounts[2]});
        let vote = await auth.votes(accounts[2]);
        assert.equal(vote.addr, accounts[3], "Wrong vote address");
        assert.equal(vote.voter, true, "Expected voter vote");
        assert.equal(vote.add, true, "Expected add vote");

        // Vote to remove 3rd voter.
        await auth.setVote(accounts[2], true, false);
        await auth.setVote(accounts[2], true, false, {from: accounts[1]});

        assert.equal(await auth.votersLength(), 2, "Voter not removed");
        assert.equal(await auth.signersLength(), 3, "Unexpected signer count");
        vote = await auth.votes(accounts[2]);
        assert.equal(vote.addr, 0, "Vote not removed");
    })

    it("onlySigners", async () => {
        assert(await auth.testOnlySigners(), "Signer unable to call");
        try {
            await auth.testOnlySigners({from: accounts[1]});
            fail("Non-signer able to call")
        } catch(error) {
            assert(error.toString().includes("Sender not a signer"), "Unexpected error")
        }
    })
})
