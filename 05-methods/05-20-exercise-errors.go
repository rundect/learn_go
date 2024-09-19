package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string{
	if e < 0 {
		return fmt.Sprintf("cannot Sqrt negative number: %v", e)
	}
	return fmt.Sprint(float64(e))
}

func Sqrt(x float64) (float64, error) {
	var result float64
	returning_result := 0.0
	for z := 1.0; z < (x + 0.1); z += 0.1{
		result = z - ((z*z - x) / (2*z))
		powed_result := math.Pow(result, 2)
		if powed_result > returning_result && powed_result < x{
			returning_result = result
		}
	}
	return 0, ErrNegativeSqrt(returning_result).Error()
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
