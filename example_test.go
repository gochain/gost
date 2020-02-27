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
			"0x00000000000000000000000071afc3a89bec7cbf060fbcde177cb6d62a1e2449"
		],
		"data": "0x000000000000000000000000000000000000000000027b46536c66c8e3000000",
		"blockNumber": "0x726167",
		"transactionHash": "0x6cf2b8c41d784708da5e3acc731cc1ca248afbf25450d868a3d2f3b916924593",
		"transactionIndex": "0x0",
		"blockHash": "0x6157ecaa8a055d1485d4daaf8634bcd7871230a16a4e124ad9b487be2577046b",
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
	// 0x3ccb06033e68bc14d4eff5f5f8d8e08a0d467054a3f5d453e17130e7f212a028
}
