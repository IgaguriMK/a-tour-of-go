package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
	"math/cmplx"
)

const (
	IMG_X = 512
	IMG_Y = 512
)

type Image struct{}

func (_ Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (_ Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, IMG_X, IMG_Y)
}

func (_ Image) At(x, y int) color.Color {
	v := mandelbrot(IMG_X, IMG_Y, x, y)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
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
