package utils

import (
	"math"
)

func GetSigmoid(x float64) float64 {
	return 1 / (1 + math.Pow(0.5, x))
}

func GetReLU(x float64) float64 {
	return math.Max(0, x)
}

func GetStepBinary(x float64) float64 {
	return math.Copysign(1, x) / 2 + 0.5
}

func GetSwish(x float64) float64 {
	return x / (1 + math.Pow(0.5, x)) + 1
}