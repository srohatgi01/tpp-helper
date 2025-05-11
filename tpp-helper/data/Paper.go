package data

import (
	"fmt"

	file_utils "github.com/srohatgi01/tpp-helper/utils"
)

type Paper struct {
	PaperName string  `json:"paper_name"`
	Price     float64 `json:"price"`
}

var PaperMap map[string]Paper

func init() {
	err := file_utils.ReadJSONFile("config/paper.json", &PaperMap)
	if err != nil {
		fmt.Println("Failed to load paper.json:", err)
	}
}

func GetPaper(paper_key string) *Paper {
	paper, ok := PaperMap[paper_key]
	if !ok {
		return nil
	}
	return &paper
}

func GetPaperByName(paper_name string) *Paper {
	for _, paper := range PaperMap {
		if paper.PaperName == paper_name {
			return &paper
		}
	}
	return nil
}

func GetPaperList() []Paper {
	paperList := []Paper{}
	for _, paper := range PaperMap {
		paperList = append(paperList, paper)
	}
	return paperList
}

func GetPaperMap() map[string]Paper {
	return PaperMap
}

func GetPaperPriceMap() map[string]float64 {
	priceMap := make(map[string]float64)
	for _, paper := range PaperMap {
		priceMap[paper.PaperName] = paper.Price
	}
	return priceMap
}
