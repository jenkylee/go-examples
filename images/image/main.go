package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct{}

func (i *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 200, 200)
}

func (i *Image) At(x, y int) color.Color {
	v := (x + y) / 2

	return color.RGBA{uint8(v), uint8(v), 0, 255}
}

func main() {
	m := &Image{}
	pic.ShowImage(m)
}
