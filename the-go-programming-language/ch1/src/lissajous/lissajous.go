package main

import (
	"image"
	"image/color/palette"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(w io.Writer) {
	const (
		nFrames    = 64
		delay      = 4
		size       = 200
		cycles     = 2
		resolution = 0.001
	)
	anim := gif.GIF{LoopCount: nFrames}
	freq := rand.Float64() * 3.0
	phase := 0.0
	for i := 0; i < nFrames; i++ {
		frame := lissajousFrame(size, cycles, resolution, freq, phase)
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, frame)
	}
	gif.EncodeAll(w, &anim)
}

func lissajousFrame(size int, cycles, resolution, freq, phase float64) *image.Paletted {
	rect := image.Rect(0, 0, size*2+1, size*2+1)
	img := image.NewPaletted(rect, palette.WebSafe)
	for t := 0.0; t < cycles*2*math.Pi; t += resolution {
		x := math.Sin(t)
		y := math.Sin(t*freq + phase)
		img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), uint8(t*5+150))
	}
	return img
}
