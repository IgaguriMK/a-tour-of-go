package main

import "fmt"

const (
	DiffLimit = 1e-10
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z, zOld := 1.0, 0.1

	for z/zOld > (1+DiffLimit) || (1-DiffLimit) > z/zOld {
		zOld = z
		z = z - (z*z-x)/(2*z)
	}

	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
