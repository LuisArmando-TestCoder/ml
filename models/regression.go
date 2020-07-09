package models

import (
	"fmt"
	"math"
	"strconv"

	"github.com/LuisArmando-TestCoder/ml/utils"
)

type TermsType struct {
	x float64
	y float64
	eachXPow2 float64
	xPow2 float64
	xy float64
	n float64
}

type TrainedType struct {
	slope float64
	intercept float64
}

type LinearType struct {
	Dataset map[string][]string
	Trained TrainedType
	t TermsType
}

func (l *LinearType) UseCSVDataset(datasetPath string) {
	l.Dataset = utils.GetCSVRecords(datasetPath)
}

func (l *LinearType) setTrainIntercept() {
	slopeNumerator := l.t.y * l.t.eachXPow2 - l.t.x * l.t.xy
	slopeDenominator := l.t.n * l.t.eachXPow2 - l.t.xPow2
	l.Trained.intercept = slopeNumerator / slopeDenominator
}

func (l *LinearType) setTrainSlope() {
	slopeNumerator := l.t.n * l.t.xy - l.t.x * l.t.y
	slopeDenominator := l.t.n * l.t.eachXPow2 - l.t.xPow2
	l.Trained.slope = slopeNumerator / slopeDenominator
}

func (l *LinearType) setTerms(label, resultLabel string) {
	l.t.x = getSummation(l.Dataset[label], 1)
	l.t.y = getSummation(l.Dataset[resultLabel], 1)
	l.t.eachXPow2 = getSummation(l.Dataset[label], 2)
	l.t.xPow2 = math.Pow(l.t.x, 2)
	l.t.xy = getSummation(getStringifiedProducts(l.Dataset[label], l.Dataset[resultLabel]), 1)
	l.t.n = float64(len(l.Dataset[label]))
}

func (l *LinearType) TrainModelWithLabels(label, resultLabel string) {
	l.setTerms(label, resultLabel)
	l.setTrainIntercept()
	l.setTrainSlope()
}

func (l *LinearType) Regress(x float64) float64 {
	y := l.Trained.intercept + l.Trained.slope * x 
	return y
}

func numerate(stringNumber string) float64 {
	if s, err := strconv.ParseFloat(stringNumber, 64); err == nil {
		return s
	}
	return 0.0
}

func getNumeratedList(list []string) []float64 {
	var numeratedList []float64
	for _, stringNumber := range list {
		numeratedList = append(numeratedList, numerate(stringNumber))
	}
	return numeratedList
}

func getStringifiedProducts(stringifiedNumbers1 []string, stringifiedNumbers2 []string) []string {
	var stringifiedProducts []string
	ns1 := getNumeratedList(stringifiedNumbers1)
	ns2 := getNumeratedList(stringifiedNumbers2)
	for i, n1 := range ns1 {
		n2 := ns2[i]
		stringifiedProduct := fmt.Sprint(n1 * n2)
		stringifiedProducts = append(stringifiedProducts, stringifiedProduct)
	}
	return stringifiedProducts
}

func getSummation(stringNumbers []string, power float64) float64 {
	summation := 0.0

	for _, n := range getNumeratedList(stringNumbers) {
		summation += math.Pow(n, power)
	}

	return summation
}