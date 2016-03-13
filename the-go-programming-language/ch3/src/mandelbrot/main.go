package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, 2, 2
		width, height          = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {
			c1 := mandelbrot(complex(
				float64(px)/width*(xmax-xmin)+xmin,
				float64(py)/height*(ymax-ymin)+ymin,
			))
			c2 := mandelbrot(complex(
				(float64(px)+0.25)/width*(xmax-xmin)+xmin,
				(float64(py)+0.25)/height*(ymax-ymin)+ymin,
			))
			c3 := mandelbrot(complex(
				(float64(px)+0.5)/width*(xmax-xmin)+xmin,
				(float64(py)+0.5)/height*(ymax-ymin)+ymin,
			))
			c4 := mandelbrot(complex(
				(float64(px)+0.75)/width*(xmax-xmin)+xmin,
				(float64(py)+0.75)/height*(ymax-ymin)+ymin,
			))

			img.Set(px, py, averageColors(c1, c2, c3, c4))
		}
	}
	png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128) color.Color {
	const (
		iterations = 255
		contrast   = 15
	)

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{
				255 - contrast*n,
				255 - contrast*(uint8(math.Abs(float64(n-200)))),
				255 - contrast*(-n),
				255,
			}
		}
	}
	return color.Black
}

func averageColors(colors ...color.Color) color.Color {
	var r, g, b, a uint32
	for _, c := range colors {
		lr, lg, lb, la := c.RGBA()
		r += lr
		g += lg
		b += lb
		a += la
	}
	r /= uint32(len(colors))
	g /= uint32(len(colors))
	b /= uint32(len(colors))
	a /= uint32(len(colors))
	return color.RGBA{
		uint8(r >> 8),
		uint8(g >> 8),
		uint8(b >> 8),
		uint8(a >> 8),
	}
}
