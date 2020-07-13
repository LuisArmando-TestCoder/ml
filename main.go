package main

import (
	"fmt"

	"github.com/LuisArmando-TestCoder/ml/utils"
)

func main() {
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
}