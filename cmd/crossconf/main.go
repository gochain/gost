package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/gochain/gochain/v3/crypto"

	"github.com/gochain/gochain/v3/accounts/abi/bind"

	"github.com/gochain/gost"

	"github.com/gochain/gochain/v3/common"
	"github.com/gochain/gochain/v3/goclient"
)

const KEY_ENV = "WEB3_PRIVATE_KEY"

//TODO worth it so that we can calculate the amount from gas price appropriately
func main() {
	//TODO request <block> <log> <hash>
	// flags: -price
	//TODO status <block> <log> <hash>

	key, ok := os.LookupEnv(KEY_ENV)
	if !ok {
		log.Fatalf("Missing env var: %s", KEY_ENV)
	}
	pk, err := crypto.HexToECDSA(strings.TrimPrefix(key, "0x"))
	if err != nil {
		log.Fatalf("Failed to parse private key hex: %v", err)
	}

	client, err := goclient.Dial("https://ropsten-rpc.linkpool.io")
	if err != nil {
		log.Fatal(err)
	}
	if err := fixedRequest(client, pk); err != nil {
		log.Fatal(err)
	}
}

//TODO generalized then remove
func fixedRequest(client *goclient.Client, key *ecdsa.PrivateKey) error {
	opts := bind.NewKeyedTransactor(key)
	opts.Context = context.Background()
	opts.GasLimit = 4000000

	confs, err := gost.NewConfirmations(common.HexToAddress("0x4bE1f85654021E6AcA62E88a7E2124D23B129493"), client)
	if err != nil {
		return err
	}
	gasPrice, err := client.SuggestGasPrice(opts.Context)
	if err != nil {
		return fmt.Errorf("failed to suggest gas price: %v", err)
	}
	opts.Nonce = big.NewInt(7) //TODO optional from flag
	opts.GasPrice = gasPrice   //TODO optional from flag

	gas, err := confs.TotalConfirmGas(&bind.CallOpts{})
	if err != nil {
		return fmt.Errorf("failed to get confirm gas: %v", err)
	}
	//TODO push this in to a top level helper which caches the totalConfirmGas (maybe a struct extends confs)
	opts.Value = new(big.Int).Mul(gas, opts.GasPrice)

	block := big.NewInt(7496039)
	logIdx := big.NewInt(0)
	hash := common.HexToHash("0x3ccb06033e68bc14d4eff5f5f8d8e08a0d467054a3f5d453e17130e7f212a028")

	tx, err := confs.Request(opts, block, logIdx, hash)
	if err != nil {
		return err
	}
	log.Println(tx.Hash().Hex())
	return nil
}
