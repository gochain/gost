.PHONY: deps generate test

deps:
	rm -rf lib/oz
	git clone --depth 1 --branch master https://github.com/OpenZeppelin/openzeppelin-solidity lib/oz

generate:
	abigen --lang go --sol ./contracts/Confirmations.sol --pkg gost --type Confirmations --out confirmations.go

test:
	go test ./...
	truffle test
