package data

import (
	"fmt"

	file_utils "github.com/srohatgi01/tpp-helper/utils"
)

type TxnHead struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Rank        int    `json:"rank"`
}

var TxnHeads map[string]TxnHead

func init() {
	err := file_utils.ReadJSONFile("config/txn_heads.json", &TxnHeads)
	if err != nil {
		fmt.Println("Failed to load txn_heads.json:", err)
	}
}

func GetTxnHead(key string) *TxnHead {
	head, ok := TxnHeads[key]
	if !ok {
		return nil
	}
	return &head
}
