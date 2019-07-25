generate:
	rm -rf lib/oz
	git clone --depth 1 --branch master https://github.com/OpenZeppelin/openzeppelin-solidity lib/oz

test:
	truffle test

.PHONY: generate test
