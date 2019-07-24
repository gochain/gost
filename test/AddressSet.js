const AddressSetMock = artifacts.require("AddressSetMock");
const AddressSet = artifacts.require("AddressSet");

contract("AddressSetMock", async accounts => {
    let set;

    beforeEach(async () => {
        let lib = await AddressSet.new();
        await AddressSetMock.link("AddressSet", lib.address);
        set = await AddressSetMock.new();
    })

    it("add", async () => {
        await set.add(accounts[0]);
        assert.equal(await set.length(), 1, "Length should be 1");
        assert.isTrue(await set.has(accounts[0]), "Set should contain address");

        try {
            await set.add(accounts[0]);
            assert.fail("Should not be able to add again");
        } catch(error) {
            assert(error.toString().includes("Already in set"), "Unexpected error");
        }
    })

    it("addAll", async () => {
        await set.addAll(accounts.slice(0,5));
        assert.equal(await set.length(), 5, "Length should be 5");
        for (let i=0;i<5;i++) {
            assert.isTrue(await set.has(accounts[i]), "Set should contain address");
        }

        try {
            await set.addAll(accounts.slice(0,5));
            assert.fail("Should not be able to addAll again");
        } catch(error) {
            assert(error.toString().includes("Already in set"), "Unexpected error");
        }
    })

    it("remove", async () => {
        await set.add(accounts[0]);
        await set.remove(accounts[0]);
        assert.equal(await set.length(), 0, "Length should be 0");
        assert.isFalse(await set.has(accounts[0]), "Set should not contain address");

        try {
            await set.remove(accounts[0]);
            assert.fail("Should not be able to remove non-existent");
        } catch(error) {
            assert(error.toString().includes("Not in set"), "Unexpected error");
        }
    })

    it("majority", async () => {
        const want = [1, 2,2, 3,3, 4,4, 5,5, 6,6, 7,7, 8,8, 9,9,
            10,10, 11,11, 12,12, 13,13, 14,14, 15,15, 16,16, 17,17, 18,18, 19,19,
            20,20, 21,21, 22,22, 23,23, 24,24, 25,25, 26];
        for (let i=0;i<50;i++) {
            await set.add(web3.utils.randomHex(20));
            assert.equal(await set.length(), i+1, "Unexpected length");
            assert.equal(await set.majority(), want[i], "Wrong majority");
        }
    })
})
