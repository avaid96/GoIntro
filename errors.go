package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	} else {
		z := 1.0
		d := (z*z - x) / 2 * z
		for math.Abs(d) > 0.0001 {
			z = z - d
			d = (z*z - x) / 2 * z
			// fmt.Println(d)
		}
		return z, nil
	}
}

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Can't sqrt the negative number: %v", float64(e))
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
