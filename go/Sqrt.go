package main

import (
	"fmt"
	"math"
)

// Sqrt returns the square root of x.
func Sqrt(x float64) (z float64) {

	if x <= 0 || x == 1 || math.IsNaN(x) || math.IsInf(x, 1) {
		return x
	}

	z = x / 2

	for i := 0; i <= 35; i++ {
		// 2z is the derivative of z²
		cacl := z - (z*z-x)/(2*z)

		if cacl == z {
			break
		}

		z = cacl
	}

	return
}

func main() {
	fmt.Println(Sqrt(1<<64 - 1))
	fmt.Println(math.Sqrt(1<<64 - 1))
}
