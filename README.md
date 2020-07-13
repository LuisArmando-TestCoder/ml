# Machine Learning with Go

## Linear Regression
[https://www.statisticshowto.com/probability-and-statistics/regression-analysis/find-a-linear-regression-equation/](https://www.statisticshowto.com/probability-and-statistics/regression-analysis/find-a-linear-regression-equation/)
```go
var linearRegression models.LinearType
linearRegression.UseCSVDataset("./data/housePrices.csv")
linearRegression.TrainModelWithLabels("SqFt", "Price")
fmt.Println(linearRegression.Regress(1790))
```

## Basic Statistics: Mean, Median & Mode
[https://www.calculator.net/mean-median-mode-range-calculator.html](https://www.calculator.net/mean-median-mode-range-calculator.html)
```go
records := utils.GetCSVRecords("./data/housePrices.csv")
fmt.Println("GetRange", utils.GetRange(records, "Price"))
fmt.Println("GetMedian", utils.GetMedian(records, "Price"))
fmt.Println("GetMean", utils.GetMean(records, "Price"))
fmt.Println("GetMode", utils.GetMode(records, "Price"))
fmt.Println("GetLowest", utils.GetLowest(records, "Price"))
fmt.Println("GetHighest", utils.GetHighest(records, "Price"))
```