// This is a small tool to aid deployment of Confirmations contracts, since the abigen code supports libraries.

package main

import (
	"context"
	"crypto/ecdsa"
	"flag"
	"log"
	"os"
	"strings"

	"github.com/gochain/gost"

	"github.com/gochain/gochain/v3/accounts/abi/bind"
	"github.com/gochain/gochain/v3/common"
	"github.com/gochain/gochain/v3/crypto"
	"github.com/gochain/gochain/v3/goclient"
)

const KEY_ENV = "WEB3_PRIVATE_KEY"

var (
	rpc   = flag.String("rpc", "https://testnet-rpc.gochain.io", "RPC URL.")
	voter = flag.String("voter", "", "Initial voter hex address.")
)

func main() {
	flag.Parse()

	if !common.IsHexAddress(*voter) {
		log.Fatalf("Invalid voter hex address: %s", *voter)
	}
	key, ok := os.LookupEnv(KEY_ENV)
	if !ok {
		log.Fatalf("Missing env var: %s", KEY_ENV)
	}
	pk, err := crypto.HexToECDSA(strings.TrimPrefix(key, "0x"))
	if err != nil {
		log.Fatalf("Failed to parse private key hex: %v", err)
	}

	client, err := goclient.Dial(*rpc)
	if err != nil {
		log.Fatalf("Failed to dial rpc: %v", err)
	}

	addr, _, err := deployConfirmations(client, pk, common.HexToAddress(*voter))
	if err != nil {
		log.Fatalf("Failed to deploy contract: %v", err)
	}
	log.Printf("Contract deployed to: %s\n", addr.Hex())
}

func deployConfirmations(client *goclient.Client, key *ecdsa.PrivateKey, voter common.Address) (common.Address, *gost.Confirmations, error) {
	opts := bind.NewKeyedTransactor(key)
	opts.GasLimit = 4000000
	expAddr, tx, confs, err := gost.DeployConfirmations(opts, client, voter)
	if err != nil {
		return common.Address{}, nil, err
	}
	log.Printf("Waiting for tx: %s\n", tx.Hash())
	log.Printf("Expecting address: %s\n", expAddr.Hex())
	addr, err := bind.WaitDeployed(context.Background(), client, tx)
	if err != nil {
		return common.Address{}, nil, err
	}
	return addr, confs, nil
}
