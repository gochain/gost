const Confirmations = artifacts.require("Confirmations");
const AddressSet = artifacts.require("AddressSet");
const Custody = artifacts.require("Custody");
const Proxy = artifacts.require("Proxy");
const TokenMock = artifacts.require("TokenMock");
const truffleAssert = require('truffle-assertions')

contract("Integration", async accounts => {
    let token;
    let custody;
    let originConfs;
    let proxy;
    let proxyConfs;

    let confGas;
    let defaultPrice = web3.utils.toBN(web3.utils.toWei("10", "gwei"));

    beforeEach(async () => {
        let addrSetLib = await AddressSet.new();

        await Confirmations.link("AddressSet", addrSetLib.address);
        originConfs = await Confirmations.new(accounts[0]);
        proxyConfs = await Confirmations.new(accounts[0]);

        token = await TokenMock.new();
        await Custody.link("AddressSet", addrSetLib.address);
        custody = await Custody.new(token.address, originConfs.address);

        await Proxy.link("AddressSet", addrSetLib.address);
        proxy = await Proxy.new(proxyConfs.address, "Proxy Token", "PRXY", 18);

        await custody.setTransferSourceContract(proxy.address);
        await proxy.setTransferSourceContract(custody.address);

        confGas = await proxyConfs.totalConfirmGas();
    })

    it("transfer origin to proxy and back", async () => {
        const amount = 10;
        let logHash = function(rawLog) {
            let hexes = ['0x000000000000000000000000',rawLog.address].concat(rawLog.topics);
            hexes.push(rawLog.data);
            let bytes = hexes.map(t => Buffer.from(web3.utils.hexToBytes(t, 'hex')));
            let toHash = '0x' + Buffer.concat(bytes).toString('hex');
            return web3.utils.keccak256(toHash);
        };

        // Initiate transfer to proxy.
        await token.approve(custody.address, amount);
        let tx = await custody.proxyTransfer(amount);
        truffleAssert.eventEmitted(tx, 'GOSTTransfer', (ev) => {
            return ev.addr == accounts[0] && ev.amount == amount;
        });

        let blockNum = tx.receipt.blockNumber;

        // Request confirmation of the transfer.
        let reqHash = await proxy.transferEventHash(accounts[0], amount);
        await proxyConfs.request(blockNum, 2, reqHash,
            {value: confGas.mul(defaultPrice), gasPrice: defaultPrice});
        assert.equal(await proxyConfs.status(blockNum, 2, reqHash), 1,
            "Expected Pending status");

        // Compute the hash from the raw log.
        let confHash = logHash(tx.receipt.rawLogs[2]);
        assert.equal(confHash, reqHash, "Hashes don't match");

        // Confirm and claim the transfer.
        await proxyConfs.confirm(blockNum, 2, confHash, true, {gasPrice: defaultPrice});
        await proxy.claimTransfer(blockNum, 2, amount);
        assert.equal(await proxy.balanceOf(accounts[0]), amount, "Unexpected proxy balance");

        // Initiate transfer back to origin.
        tx = await proxy.originTransferTo(amount, accounts[1]);
        truffleAssert.eventEmitted(tx, 'GOSTTransfer', (ev) => {
            return ev.addr == accounts[1] && ev.amount == amount;
        });
        
        blockNum = tx.receipt.blockNumber;
        
        // Request confirmation.
        reqHash = await custody.transferEventHash(accounts[1], amount);
        await originConfs.request(blockNum, 0, reqHash,
            {value: confGas.mul(defaultPrice), gasPrice: defaultPrice});
        assert.equal(await originConfs.status(blockNum, 0, reqHash), 1,
            "Expected Pending status");
        
        // Compute the hash from the raw log.
        confHash = logHash(tx.receipt.rawLogs[1]);
        assert.equal(confHash, reqHash, "Hashes don't match");
        
        // Confirm and claim the transfer.
        await originConfs.confirm(blockNum, 0, confHash, true, {gasPrice: defaultPrice});
        await custody.claimTransfer(blockNum, 0, amount, {from: accounts[1]});
        assert.equal(await token.balanceOf(accounts[1]), amount, "Unexpected token balance");
        assert.equal(await token.balanceOf(custody.address), 0, "Unexpected token balance");
    })

    //TODO more tests
})
