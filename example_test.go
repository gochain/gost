package gost

import (
	"encoding/json"
	"fmt"

	"github.com/gochain/gochain/v3/core/types"
)

func Example() {
	jsonStr := `{
		"address": "0x054ab6c0d612285f5fd119a984ad133f7b730a72",
		"topics": [
			"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef",
			"0x00000000000000000000000030dd8ba783a0fc6a91491dc7af603f30346ccd78",
			"0x0000000000000000000000005ff264137b0fecf76d08f173dbd8610e27cdfc5f"
		],
		"data": "0x0000000000000000000000000000000000000000000000000000000000000190",
		"blockNumber": "0x7295c7",
		"transactionHash": "0x260fc133404c69bf7f161ccb63fffc9d00a74a8a96f8371f2ca19e6d1cdb5623",
		"transactionIndex": "0x0",
		"blockHash": "0xbd3436d487c6da58b50451b5923bea5e6d1ad89ea24d135a0eaef1d6bd5decfd",
		"logIndex": "0x0",
		"removed": false
	}`
	var l *types.Log
	err := json.Unmarshal([]byte(jsonStr), &l)
	if err != nil {
		fmt.Println(err)
		return
	}
	hash := HashLog(l)
	fmt.Println(hash.Hex())

	// Output:
	// 0x0d26878bb299d578ad23488e239a901f4c4338c7395426ece438d043a375f855
}
