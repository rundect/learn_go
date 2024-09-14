package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	var result float64
	returning_result := 0.0
	for z := 1.0; z < (x + 0.1); z += 0.1{
		result = z - ((z*z - x) / (2*z))
		powed_result := math.Pow(result, 2)
		if powed_result > returning_result && powed_result < x{
			returning_result = result
		}
	}
	return returning_result
}

func main() {
	fmt.Println(Sqrt(2))
}
