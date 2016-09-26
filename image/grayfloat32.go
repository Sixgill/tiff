package image

import (
	"image"

	"github.com/sixgill/tiff/image/color"
)

// implement a (gray) GrayFloat32 image type and color

// GrayFloat32 is an in-memory image whose At method returns GrayFloat32 values.
type GrayFloat32 struct {
	// Pix holds the image's pixels, as gray values in big-endian format. The pixel at
	// (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*2].
	Pix []uint8
	// Stride is the Pix stride (in bytes) between vertically adjacent pixels.
	Stride int
	// Rect is the image's bounds.
	Rect image.Rectangle
}

func (p *GrayFloat32) ColorModel() color.Model { return color.GrayFloat32Model }

func (p *GrayFloat32) Bounds() image.Rectangle { return p.Rect }

func (p *GrayFloat32) At(x, y int) color.Color {
	return p.GrayFloat32At(x, y)
}

func (p *GrayFloat32) GrayFloat32At(x, y int) color.GrayFloat32 {
	if !(Point{x, y}.In(p.Rect)) {
		return color.GrayFloat32{}
	}
	i := p.PixOffset(x, y)
	return color.GrayFloat32{uint16(p.Pix[i+0])<<8 | uint16(p.Pix[i+1])}
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *GrayFloat32) PixOffset(x, y int) int {
	return (y-p.Rect.Min.Y)*p.Stride + (x-p.Rect.Min.X)*2
}

func (p *GrayFloat32) Set(x, y int, c color.Color) {
	if !(Point{x, y}.In(p.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	c1 := color.GrayFloat32Model.Convert(c).(color.GrayFloat32)
	p.Pix[i+0] = uint8(c1.Y >> 8)
	p.Pix[i+1] = uint8(c1.Y)
}

func (p *GrayFloat32) SetGrayFloat32(x, y int, c color.GrayFloat32) {
	if !(Point{x, y}.In(p.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	p.Pix[i+0] = uint8(c.Y >> 8)
	p.Pix[i+1] = uint8(c.Y)
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *GrayFloat32) SubImage(r image.Rectangle) Image {
	r = r.Intersect(p.Rect)
	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to be inside
	// either r1 or r2 if the intersection is empty. Without explicitly checking for
	// this, the Pix[i:] expression below can panic.
	if r.Empty() {
		return &GrayFloat32{}
	}
	i := p.PixOffset(r.Min.X, r.Min.Y)
	return &GrayFloat32{
		Pix:    p.Pix[i:],
		Stride: p.Stride,
		Rect:   r,
	}
}

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *GrayFloat32) Opaque() bool {
	return true
}

// NewGrayFloat32 returns a new GrayFloat32 image with the given bounds.
func NewGrayFloat32(r image.Rectangle) *GrayFloat32 {
	w, h := r.Dx(), r.Dy()
	pix := make([]uint8, 2*w*h)
	return &GrayFloat32{pix, 2 * w, r}
}
