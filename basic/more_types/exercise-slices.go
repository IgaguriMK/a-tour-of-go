package main

import (
	"golang.org/x/tour/pic"
	"math/cmplx"
)

func Pic(dx, dy int) [][]uint8 {
	lines := make([][]uint8, dy)
	for y := 0; y < len(lines); y++ {
		lines[y] = make([]uint8, dx)
	}

	for y, line := range lines {
		for x := 0; x < len(line); x++ {
			line[x] = mandelbrot(dx, dy, x, y)
		}
	}

	return lines
}

func mandelbrot(dx, dy, x, y int) uint8 {
	var c complex128 = complex(2.0*float64(x-dx/2)/float64(dx)-0.5,
		2.0*float64(y-dy/2)/float64(dy))
	var z complex128 = 0.0 + 0.0i

	for i := 0; i < 255; i++ {
		z = z*z + c

		if cmplx.Abs(z) > 100.0 {
			return uint8(i)
		}
	}

	return 0
}

func main() {
	pic.Show(Pic)
}
