package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z:=1.0
	for i:=0;i<20;i++ {
		z=z-((z*z-x)/2*z)
	}
	return z
}

func nSqrt(x float64) float64 {
	z:=1.0
	d:=(z*z-x)/2*z
	fmt.Println(d)
	for math.Abs(d)>0.0001 {
		z=z-d
		d=(z*z-x)/2*z
		// fmt.Println(d)
	}
	return z
}

func main() {
	fmt.Println(nSqrt(2))
}
