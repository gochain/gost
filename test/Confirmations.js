const Confirmations = artifacts.require("Confirmations");
const AddressSet = artifacts.require("AddressSet");
const truffleAssert = require('truffle-assertions')

contract("Confirmations", async accounts => {
    let hash = web3.utils.keccak256("test hash");
    let confs;
    let confGas;
    let defaultPrice = web3.utils.toBN(web3.utils.toWei("10", "gwei"));

    beforeEach(async () => {
        let addrSetLib = await AddressSet.new();
        await Confirmations.link("AddressSet", addrSetLib.address);
        confs = await Confirmations.new(accounts[0]);
        confGas = await confs.totalConfirmGas();
    })

    it("confirm", async () => {
        let tx = await confs.request(1, 1, hash, {value: confGas.mul(defaultPrice), gasPrice: defaultPrice});
        truffleAssert.eventEmitted(tx, 'ConfirmationRequested', (ev) => {
            return ev.blockNum == 1 && ev.logIndex == 1 && ev.eventHash == hash;
        });
        let s = await confs.status(1, 1, hash);
        assert.equal(s, 1, "Status not pending");

        tx = await confs.confirm(1, 1, hash, true, {gasPrice: defaultPrice});
        truffleAssert.eventEmitted(tx, 'Confirmed', (ev) => {
            return ev.blockNum == 1 && ev.logIndex == 1 && ev.eventHash == hash && ev.valid == true;
        });
        s = await confs.status(1, 1, hash);
        assert.equal(s, 3, "Status not Confirmed");

        // Ensure extra confirm ignored.
        tx = await confs.confirm(1, 1, hash, true);
        truffleAssert.eventNotEmitted(tx, 'Confirmed');
    })

    it("confirm requires pending", async () => {
        let tx = await confs.confirm(1, 1, hash, true);
        truffleAssert.eventNotEmitted(tx, 'Confirmed');
        let s = await confs.status(1, 1, hash);
        assert.equal(s, 0, "Status not None");
    })

    it("confirm requires signer", async () => {
        await confs.request(1, 1, hash, {value: confGas.mul(defaultPrice), gasPrice: defaultPrice});
        try {
            await confs.confirm(1, 1, hash, true, {from: accounts[1]});
            fail("Non-signer confirmed");
        } catch(error) {
            assert(error.toString().includes("Sender not a signer"), "Unexpected error:" + error);
        }
    })

    it("post-process on signer remove", async () => {
        // Add another signer.
        await confs.setVote(accounts[1],false,true);

        // Request and confirm 2 events from 1.
        await confs.request(1, 1, hash, {value: confGas.mul(defaultPrice), gasPrice: defaultPrice});
        await confs.confirm(1, 1, hash, true, {gasPrice: defaultPrice});
        await confs.request(2, 2, hash, {value: confGas.mul(defaultPrice), gasPrice: defaultPrice});
        await confs.confirm(2, 2, hash, false, {gasPrice: defaultPrice});

        // Verify pending.
        let s = await confs.status(1, 1, hash);
        assert.equal(s, 1, "Status not pending");
        s = await confs.status(2, 2, hash);
        assert.equal(s, 1, "Status not pending");

        // Verify other signer is allowed while 0 is not.
        let a, reqPrice = await confs.shouldConfirm(1, 1, hash);
        assert.equal(a, false, "Confirm still allowed");
        assert.equal(reqPrice, defaultPrice, "Request price mismatch");
        a, reqPrice = await confs.shouldConfirm(1, 1, hash, {from: accounts[1]});
        assert.equal(a, true, "Confirm not allowed");
        assert.equal(reqPrice, defaultPrice, "Request price mismatch");

        // Remove other signer.
        let tx = await confs.setVote(accounts[1], false, false);

        // Verify post-removal consensus processed for pending.
        truffleAssert.eventEmitted(tx, 'Confirmed', (ev) => {
            return ev.blockNum == 1 && ev.logIndex == 1 && ev.eventHash == hash && ev.valid == true;
        });
        truffleAssert.eventEmitted(tx, 'Confirmed', (ev) => {
            return ev.blockNum == 2 && ev.logIndex == 2 && ev.eventHash == hash && ev.valid == false;
        });
        s = await confs.status(1, 1, hash);
        assert.equal(s, 3, "Status not Confirmed");
        s = await confs.status(2, 2, hash);
        assert.equal(s, 2, "Status not Invalid");
    })

    it("confirmation gas used", async () => {
        // Need 50 signging accounts.
        let seeded = [accounts[0]];
        for (let i=0;i<49;i++) {
            let s = await web3.eth.personal.newAccount("password");
            await web3.eth.personal.unlockAccount(s, "password", 600);
            seeded.push(s);
            await web3.eth.sendTransaction({value: web3.utils.toWei("1", "ether"), from: accounts[0], to: s});
        }

        let block = 0;
        let test = async function(signers) {
            block++;
            let majority = Math.floor(signers.length/2)+1;
            // Vote all signers in (noop if already a signer).
            for (let i=0;i<signers.length;i++) {
                await confs.setVote(signers[i], false, true);
            }

            await confs.request(block, 1, hash,
                {value: confGas.mul(defaultPrice), gasPrice: defaultPrice}); //TODO vary gas price?

            var balsBefore = [];
            for (let i=0;i<majority;i++) {
                let bal = await web3.eth.getBalance(signers[i]);
                balsBefore.push(web3.utils.toBN(bal));
            }

            for (let i=0;i<majority;i++) {
                await confs.confirm(block, 1, hash, true, {from: signers[i], gasPrice: defaultPrice});
            }

            var balsAfter = [];
            for (let i=0;i<majority;i++) {
                let bal = await web3.eth.getBalance(signers[i]);
                balsAfter.push(web3.utils.toBN(bal));
            }

            for (let i=0;i<majority;i++) {
                let a = balsAfter[i];
                let b = balsBefore[i];
                if (a.lt(b)) {
                    assert.fail("Confirming was net negative for signer " + (i+1) + "/" + signers.length + ": " +
                        b.toString() + "->" + a.toString() + "=" + a.sub(b));
                }
            }
            assert.equal(await web3.eth.getBalance(confs.address), 0, "Fee should payout completely");
        }

        for (let i=1;i<=50;i++) {
            await test(seeded.slice(0, i));
        }
    })

    //TODO under/overpriced claim rejected, matching accepted (and vary the price: 1, 10, 100 gwei?)

    //TODO test underpayment rejected

    //TODO test overpayment refunded

    //TODO test more complex events than just GOSTTransfer
})
