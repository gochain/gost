.PHONY: deps generate test

deps:
	rm -rf lib/oz
	git clone --depth 1 --branch master https://github.com/OpenZeppelin/openzeppelin-solidity lib/oz

generate:
	cd contracts && web3 contract build Confirmations.sol && cd ..
	rm -rf out
	mkdir out
	mv contracts/*.abi out
	mv contracts/*.bin out
	abigen --lang go --abi ./out/Confirmations.abi --bin ./out/Confirmations.bin --pkg gost --type Confirmations --out confirmations.go

test:
	truffle test
