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

type RGBA256i struct {
	M struct {
		Pix    []uint8
		Stride int
		Rect   image.Rectangle
	}
}

// NewRGBA256i returns a new RGBA256i with the given bounds.
func NewRGBA256i(r image.Rectangle) *RGBA256i {
	return new(RGBA256i).Init(make([]uint8, 32*r.Dx()*r.Dy()), 32*r.Dx(), r)
}

func (p *RGBA256i) Init(pix []uint8, stride int, rect image.Rectangle) *RGBA256i {
	*p = RGBA256i{
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

func (p *RGBA256i) Pix() []byte           { return p.M.Pix }
func (p *RGBA256i) Stride() int           { return p.M.Stride }
func (p *RGBA256i) Rect() image.Rectangle { return p.M.Rect }
func (p *RGBA256i) Channels() int         { return 4 }
func (p *RGBA256i) Depth() reflect.Kind   { return reflect.Int64 }

func (p *RGBA256i) ColorModel() color.Model { return colorRGBA256iModel }

func (p *RGBA256i) Bounds() image.Rectangle { return p.M.Rect }

func (p *RGBA256i) At(x, y int) color.Color {
	return p.RGBA256iAt(x, y)
}

func (p *RGBA256i) RGBA256iAt(x, y int) colorRGBA256i {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return colorRGBA256i{}
	}
	i := p.PixOffset(x, y)
	return pRGBA256iAt(p.M.Pix[i:])
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *RGBA256i) PixOffset(x, y int) int {
	return (y-p.M.Rect.Min.Y)*p.M.Stride + (x-p.M.Rect.Min.X)*32
}

func (p *RGBA256i) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	c1 := colorRGBA256iModel.Convert(c).(colorRGBA256i)
	pSetRGBA256i(p.M.Pix[i:], c1)
	return
}

func (p *RGBA256i) SetRGBA256i(x, y int, c colorRGBA256i) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	pSetRGBA256i(p.M.Pix[i:], c)
	return
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *RGBA256i) SubImage(r image.Rectangle) image.Image {
	r = r.Intersect(p.M.Rect)
	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to be inside
	// either r1 or r2 if the intersection is empty. Without explicitly checking for
	// this, the Pix[i:] expression below can panic.
	if r.Empty() {
		return &RGBA256i{}
	}
	i := p.PixOffset(r.Min.X, r.Min.Y)
	return new(RGBA256i).Init(
		p.M.Pix[i:],
		p.M.Stride,
		r,
	)
}

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *RGBA256i) Opaque() bool {
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