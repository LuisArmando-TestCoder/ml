# Machine Learning with Go

## Linear Regression
[equation: finding a linear regression](https://www.statisticshowto.com/probability-and-statistics/regression-analysis/find-a-linear-regression-equation/)
```go
var linearRegression models.LinearType
linearRegression.UseCSVDataset("./data/housePrices.csv")
linearRegression.TrainModelWithLabels("SqFt", "Price")
fmt.Println(linearRegression.Regress(1790))
```

## Basic Statistics: Mean, Median & Mode
[mean, median, mode & range](https://www.calculator.net/mean-median-mode-range-calculator.html)
```go
records := utils.GetCSVRecords("./data/housePrices.csv")
fmt.Println("GetRange", utils.GetRange(records, "Price"))
fmt.Println("GetMedian", utils.GetMedian(records, "Price"))
fmt.Println("GetMean", utils.GetMean(records, "Price"))
fmt.Println("GetMode", utils.GetMode(records, "Price"))
fmt.Println("GetLowest", utils.GetLowest(records, "Price"))
fmt.Println("GetHighest", utils.GetHighest(records, "Price"))
```

## Vectors and Matrices
[GetMultipliedVectors](http://thejuniverse.org/PUBLIC/LinearAlgebra/MATH-232/Unit.1/Presentation.2/Section1B/multiplication.html)
[GetMultipliedMatrix](https://www.mathwarehouse.com/algebra/matrix/images/matrix-multiplication/how-to-multiply-2-matrices-demo.gif)
[GetSummedUpMatrix](https://economipedia.com/definiciones/suma-de-matrices.html)
```go
v1 := []float64{2,3,5}
v2 := []float64{3,3,6}

fmt.Println("GetMultipliedVectors", utils.GetMultipliedVectors(v1, v2))
fmt.Println("AddFloatingBiasToVector", utils.AddFloatingBiasToVector(3.0, v1))

m1 := [][]float64{
    {2, 3, 5},
    {5, 6, 8},
}
m2 := [][]float64{
    {5, 5, 6},
    {4, 4, 2},
}

fmt.Println("GetSummedUpMatrix", utils.GetSummedUpMatrix(m1, m2))
fmt.Println("GetMultipliedMatrix", utils.GetMultipliedMatrix(m1, m2))
```

## Activation Functions
[types of activation functions](https://missinglink.ai/guides/neural-network-concepts/7-types-neural-network-activation-functions-right/)
```go
fmt.Println("GetSigmoid", utils.GetSigmoid(0.5))
// GetSigmoid 0.585786437626905

fmt.Println("GetReLU", utils.GetReLU(0.5))
// GetReLU 0.5

fmt.Println("GetStepBinary", utils.GetStepBinary(0.5))
// GetStepBinary 1

fmt.Println("GetSwish", utils.GetSwish(0.5))
// GetSwish 1.2928932188134525
```