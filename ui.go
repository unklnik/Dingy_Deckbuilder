package main

import (
	z "github.com/gen2brain/raylib-go/raylib"
)

var ()

type BUTTON struct {
	r            z.Rectangle
	onoff        bool
	c, cON, cOFF z.Color
	lineW        float32
	im           IM
}

// MARK: DRAW
func dCLOSEtopRight(siz, offset float32, c z.Color, value bool) bool {
	r := z.NewRectangle(float32(SCRW)-(siz+offset), offset, siz, siz)
	if cMSREC(r) {
		dIMcolor(UI[0], r, z.Maroon)
	} else {
		dIMcolor(UI[0], r, c)
	}
	if cMSREC(r) && MSL2 {
		value = !value
	}
	return value
}
func dCLOSE(x, y, siz float32, c z.Color) bool {
	close := false
	r := z.NewRectangle(x, y, siz, siz)
	if cMSREC(r) {
		dIMcolor(UI[0], r, c)
		if MSL {
			close = true
		}
	} else {
		dIMcolor(UI[0], r, z.Maroon)
	}
	return close
}
func dBUTTON(b BUTTON, value bool) bool {

	if cMSREC(b.r) {
		if MSL {
			value = !value
		}
		r2 := qResizeRecOffsetLRGR(b.r, b.r.Width/8)
		dREC(r2, z.Magenta)
	}
	b.onoff = value
	dRECLINE(b.r, b.lineW, b.c)
	r2 := qResizeRecOffsetSMLR(b.r, b.r.Width/8)
	if b.onoff {
		dREC(r2, b.cON)
	} else {
		dREC(r2, b.cOFF)
	}
	return value
}

// MARK: MAKE

func mBUTTON(x, y, siz, lineW float32, c, cON, cOFF z.Color) BUTTON {
	b := BUTTON{}
	b.r = z.NewRectangle(x, y, siz, siz)
	b.c = c
	b.cON = cON
	b.cOFF = cOFF
	b.lineW = lineW
	return b
}
