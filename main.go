package main

import (
	"fmt"

	"github.com/LuisArmando-TestCoder/ml/utils"
)

func main() {
	records := utils.GetCSVRecords("./data/housePrices.csv")
	fmt.Println("GetRange", utils.GetRange(records, "Price"))
	fmt.Println("GetMedian", utils.GetMedian(records, "Price"))
	fmt.Println("GetMean", utils.GetMean(records, "Price"))
	fmt.Println("GetMode", utils.GetMode(records, "Price"))
	fmt.Println("GetLowest", utils.GetLowest(records, "Price"))
	fmt.Println("GetHighest", utils.GetHighest(records, "Price"))
}
