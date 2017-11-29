package main

import (
	"fmt"
	"math"
)

const (
	DiffLimit = 1e-10
)

func Sqrt(x float64) float64 {
	z, zOld := 1.0, 0.1

	for z/zOld > (1 + DiffLimit) || (1 - DiffLimit) > z/zOld  {
		zOld = z
		z = z - (z*z - x) / (2*z)
		fmt.Println("|", z)
	}

	return z
}

func main() {
	x := 1e10
	
	fmt.Println(Sqrt(x) - math.Sqrt(x))
}
