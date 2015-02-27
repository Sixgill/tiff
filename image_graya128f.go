// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Auto Generated By 'go generate', DONOT EDIT!!!

package tiff

import (
	"image"
	"image/color"
	"reflect"
)

type GrayA128f struct {
	M struct {
		Pix    []uint8
		Stride int
		Rect   image.Rectangle
	}
}

// NewGrayA128f returns a new GrayA128f with the given bounds.
func NewGrayA128f(r image.Rectangle) *GrayA128f {
	return new(GrayA128f).Init(make([]uint8, 16*r.Dx()*r.Dy()), 16*r.Dx(), r)
}

func (p *GrayA128f) Init(pix []uint8, stride int, rect image.Rectangle) *GrayA128f {
	*p = GrayA128f{
		M: struct {
			Pix    []uint8
			Stride int
			Rect   image.Rectangle
		}{
			Pix:    pix,
			Stride: stride,
			Rect:   rect,
		},
	}
	return p
}

func (p *GrayA128f) Pix() []byte           { return p.M.Pix }
func (p *GrayA128f) Stride() int           { return p.M.Stride }
func (p *GrayA128f) Rect() image.Rectangle { return p.M.Rect }
func (p *GrayA128f) Channels() int         { return 2 }
func (p *GrayA128f) Depth() reflect.Kind   { return reflect.Float64 }

func (p *GrayA128f) ColorModel() color.Model { return colorGrayA128fModel }

func (p *GrayA128f) Bounds() image.Rectangle { return p.M.Rect }

func (p *GrayA128f) At(x, y int) color.Color {
	return p.GrayA128fAt(x, y)
}

func (p *GrayA128f) GrayA128fAt(x, y int) colorGrayA128f {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return colorGrayA128f{}
	}
	i := p.PixOffset(x, y)
	return pGrayA128fAt(p.M.Pix[i:])
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *GrayA128f) PixOffset(x, y int) int {
	return (y-p.M.Rect.Min.Y)*p.M.Stride + (x-p.M.Rect.Min.X)*16
}

func (p *GrayA128f) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	c1 := colorGrayA128fModel.Convert(c).(colorGrayA128f)
	pSetGrayA128f(p.M.Pix[i:], c1)
	return
}

func (p *GrayA128f) SetGrayA128f(x, y int, c colorGrayA128f) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	pSetGrayA128f(p.M.Pix[i:], c)
	return
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *GrayA128f) SubImage(r image.Rectangle) image.Image {
	r = r.Intersect(p.M.Rect)
	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to be inside
	// either r1 or r2 if the intersection is empty. Without explicitly checking for
	// this, the Pix[i:] expression below can panic.
	if r.Empty() {
		return &GrayA128f{}
	}
	i := p.PixOffset(r.Min.X, r.Min.Y)
	return new(GrayA128f).Init(
		p.M.Pix[i:],
		p.M.Stride,
		r,
	)
}

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *GrayA128f) Opaque() bool {
	if p.M.Rect.Empty() {
		return true
	}
	for y := p.M.Rect.Min.Y; y < p.M.Rect.Max.Y; y++ {
		for x := p.M.Rect.Min.X; x < p.M.Rect.Max.X; x++ {
			if _, _, _, a := p.At(x, y).RGBA(); a == 0xFFFF {
				return false
			}
		}
	}
	return true
}