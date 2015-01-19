package main

import (
    "code.google.com/p/go-tour/pic"
    "image/color"
    "image"
)

type Image struct {}

func (i Image) ColorModel() color.Model {
    return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
    return image.Rectangle{image.Point{0, 0}, image.Point{250, 150}}
}

func (i Image) At(x, y int) color.Color {
    return color.RGBA{uint8(x), uint8(y), 0xff, 0xff}
}

func main() {
    m := Image{}
    pic.ShowImage(m)
}
