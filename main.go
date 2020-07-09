package main

import (
	"fmt"

	"github.com/LuisArmando-TestCoder/ml/utils"
)

func main() {
	records := utils.GetCSVRecords("./data/housePrices.csv")
	fmt.Println(records["Home"][0])
	fmt.Println(records["Price"][0])
}