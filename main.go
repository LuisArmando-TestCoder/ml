package main

import (
	"fmt"

	"github.com/LuisArmando-TestCoder/ml/utils"
)

func main() {	
	fmt.Println("GetSigmoid", utils.GetSigmoid(0.5))
	
	fmt.Println("GetReLU", utils.GetReLU(0.5))
	
	fmt.Println("GetStepBinary", utils.GetStepBinary(0.5))
	
	fmt.Println("GetSwish", utils.GetSwish(0.5))
}