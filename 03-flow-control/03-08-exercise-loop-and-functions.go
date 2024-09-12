package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	var a float64
	for z := 1.0; z < (x + 0.1); z += 0.1{
		a = z - ((z*z - x) / (2*z))
	}
	return a
}

func main() {
	fmt.Println(Sqrt(2))
}
