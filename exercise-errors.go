package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("ErrNegativeSqrt: cannot Sqrt negative number: %.0f", e)
}

func Sqrt(x float64) (float64, error) {
	// return error if z is negative
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z := 1.0

	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}

	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
