package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0

	for i := 0; i<10; i++ {
		z -= (z*z - x) / (2*z)
	}
	fmt.Printf("The square root of %.1f is %.1f\n", x, z)
	fmt.Printf("Lets check. %.1f * %.1f = %.1f\n", z, z, z*z)
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
