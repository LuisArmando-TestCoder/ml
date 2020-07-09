package main

import (
	"fmt"

	"github.com/LuisArmando-TestCoder/ml/models"
)

func main() {
	var linearRegression models.LinearType
	linearRegression.UseCSVDataset("./data/housePrices.csv")
	linearRegression.TrainModelWithLabels("SqFt", "Price")
	fmt.Println(linearRegression.Regress(1790))
}