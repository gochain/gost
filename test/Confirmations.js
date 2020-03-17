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

    it("request and confirm", async () => {
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
        await truffleAssert.reverts(confs.confirm(1, 1, hash, true), "Status not Pending",
          "Confirmed non-pending");
    })

    it("request twice and confirm", async () => {
        let tx = await confs.request(1, 1, hash, {value: confGas.mul(defaultPrice), gasPrice: defaultPrice});
        truffleAssert.eventEmitted(tx, 'ConfirmationRequested', (ev) => {
            return ev.blockNum == 1 && ev.logIndex == 1 && ev.eventHash == hash;
        });
        let s = await confs.status(1, 1, hash);
        assert.equal(s, 1, "Status not pending");

        // Second request.
        tx = await confs.request(1, 1, hash, {value: confGas.mul(defaultPrice), gasPrice: defaultPrice});
        truffleAssert.eventEmitted(tx, 'ConfirmationRequested', (ev) => {
            return ev.blockNum == 1 && ev.logIndex == 1 && ev.eventHash == hash;
        });
        s = await confs.status(1, 1, hash);
        assert.equal(s, 1, "Status not pending");

        // Check price and votes.
        let ps = await confs.pendingStatus(1, 1, hash);
        expect(ps.gasPrice).to.eql(web3.utils.toBN(2).mul(defaultPrice));
        expect(ps.total).to.eql(web3.utils.toBN(1));
        expect(ps.valid).to.eql(web3.utils.toBN(0));
        expect(ps.invalid).to.eql(web3.utils.toBN(0));

        // Confirm at original gas price.
        tx = await confs.confirm(1, 1, hash, true, {gasPrice: defaultPrice});
        truffleAssert.eventEmitted(tx, 'Confirmed', (ev) => {
            return ev.blockNum == 1 && ev.logIndex == 1 && ev.eventHash == hash && ev.valid == true;
        });
        s = await confs.status(1, 1, hash);
        assert.equal(s, 3, "Status not Confirmed");

        assert.equal(await web3.eth.getBalance(confs.address), 0, "Fee should payout completely");
    })

    it("confirm requires pending", async () => {
        await truffleAssert.reverts(confs.confirm(1, 1, hash, true),"Status not Pending",
          "should fail since not pending");
    })

    it("confirm requires signer", async () => {
        await confs.request(1, 1, hash, {value: confGas.mul(defaultPrice), gasPrice: defaultPrice});
        await truffleAssert.reverts(confs.confirm(1, 1, hash, true, {from: accounts[1]}),
          "Sender not a signer",
          "Non-signer confirmed");
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
        let should = await confs.shouldConfirm(1, 1, hash);
        assert.equal(should[0], false, "Confirm still allowed");
        should = await confs.shouldConfirm(1, 1, hash, {from: accounts[1]});
        assert.equal(should[0], true, "Confirm not allowed");
        assert.equal(should[1].toString(), defaultPrice.toString(), "Request price mismatch");

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
        // Need 50 signing accounts.
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

    it("request rejects underpayment", async () => {
        await truffleAssert.reverts(confs.request(1, 1, hash, {gasPrice: defaultPrice}),
          "Tx value doesn't cover confirmation fee",
          "Underpriced request");
    })

    it("request refunds overpayment", async () => {
        const fee = confGas.mul(defaultPrice);
        const over = fee.add(web3.utils.toBN(10));
        await confs.request(1, 1, hash, {value: over, gasPrice: defaultPrice});
        assert.equal(await web3.eth.getBalance(confs.address), fee, "Overpayment not refunded");
    })

    it("confirm ignores gas price", async () => {
        let tx = await confs.request(1, 1, hash, {value: confGas.mul(defaultPrice), gasPrice: defaultPrice});
        truffleAssert.eventEmitted(tx, 'ConfirmationRequested', (ev) => {
            return ev.blockNum == 1 && ev.logIndex == 1 && ev.eventHash == hash;
        });

        const x10th = defaultPrice.divRound(web3.utils.toBN(10));
        tx = await confs.request(1, 1, hash, {value: confGas.mul(x10th), gasPrice: x10th});
        truffleAssert.eventEmitted(tx, 'ConfirmationRequested', (ev) => {
            return ev.blockNum == 1 && ev.logIndex == 1 && ev.eventHash == hash;
        });

        const x10 = defaultPrice.mul(web3.utils.toBN(10));
        tx = await confs.request(1, 1, hash, {value: confGas.mul(x10), gasPrice: x10});
        truffleAssert.eventEmitted(tx, 'ConfirmationRequested', (ev) => {
            return ev.blockNum == 1 && ev.logIndex == 1 && ev.eventHash == hash;
        });

        // Check price.
        let ps = await confs.pendingStatus(1, 1, hash);
        assert.equal(ps.gasPrice.toString(), defaultPrice.add(x10th).add(x10).toString(), "Pending price not summed");
    })
})
