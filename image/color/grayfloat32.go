package color

import (
	"image/color"
)

// GrayFloat32 represents a 32-bit grayscale color.
type GrayFloat32 struct {
	Y float32
}

func (c GrayFloat32) RGBA() (r, g, b, a uint32) {
	y := uint32(c.Y)
	return y, y, y, 0xffff
}

var GrayFloat32Model color.Model = color.ModelFunc(grayFloat32Model)

func grayFloat32Model(c color.Color) color.Color {
	if _, ok := c.(GrayFloat32); ok {
		return c
	}
	r, g, b, _ := c.RGBA()
	y := (299*r + 587*g + 114*b + 500) / 1000
	return GrayFloat32{float32(y)}
}
