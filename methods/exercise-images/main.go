package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct{}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 256, 256)
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) At(x, y int) color.Color {
	v := uint8((x * y) % 256)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
