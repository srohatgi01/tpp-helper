package data

import (
	"fmt"

	file_utils "github.com/srohatgi01/tpp-helper/utils"
)

type Lamination struct {
	Price  float64 `json:"price"`
	Name   string  `json:"name"`
	Seller string  `json:"seller"`
}

var LaminationMap map[string]Lamination

func init() {
	err := file_utils.ReadJSONFile("config/lamination.json", &LaminationMap)
	if err != nil {
		fmt.Println("Failed to load lamination.json:", err)
	}
}

func GetLamination(key string) *Lamination {
	lamination, ok := LaminationMap[key]
	if !ok {
		return nil
	}
	return &lamination
}
