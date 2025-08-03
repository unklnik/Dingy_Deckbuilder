package main

import (
	z "github.com/gen2brain/raylib-go/raylib"
)

type GRID struct {
	numW, numH, numA     int
	blokSiz, W, H        float32
	cnt, tl, tr, bl, br  z.Vector2
	im, im2, bgIM, bgIM2 []IM
	r                    z.Rectangle
}
type CIRC struct {
	cnt z.Vector2
	rad float32
	c   z.Color
}

// MARK: ELEMENTS
func dWallsShadowInner(walls []IM, c z.Color, offset float32, a uint8) {
	xl := walls[0].rD.X
	yt := walls[0].rD.Y
	xr := xl
	yb := yt
	var v1, v2, v3, v4 z.Vector2
	for i := range walls {
		if walls[i].rD.X > xr {
			xr = walls[i].rD.X
		}
		if walls[i].rD.Y > yb {
			yb = walls[i].rD.Y
		}
	}
	v1 = z.NewVector2(xl, yt)
	v1.X += walls[0].rD.Width
	v1.Y += walls[0].rD.Height
	v2 = z.NewVector2(xr, yt)
	v2.Y += walls[0].rD.Height
	v3 = z.NewVector2(xr, yb)
	v4 = z.NewVector2(xl, yb)
	v4.X += walls[0].rD.Width
	w := v2.X - v1.X
	h := v3.Y - v1.Y
	z.DrawRectangleGradientV(int32(v1.X), int32(v1.Y), int32(w), int32(offset), CA(c, a), z.Blank)
	z.DrawRectangleGradientH(int32(v1.X), int32(v1.Y), int32(offset), int32(h), CA(c, a), z.Blank)
	z.DrawRectangleGradientV(int32(v4.X), int32(v4.Y-offset), int32(w), int32(offset), z.Blank, CA(c, a))
	z.DrawRectangleGradientH(int32(v2.X-offset), int32(v2.Y), int32(offset), int32(h), z.Blank, CA(c, a))
}
func mWallsRec(r z.Rectangle, blokSiz float32, im IM) []IM {
	var ims []IM
	x := r.X
	y := r.Y
	x -= blokSiz
	y -= blokSiz
	for x < r.X+r.Width+blokSiz {
		im.rD = R(x, y, blokSiz, blokSiz)
		ims = append(ims, im)
		im.rD.Y += r.Height + blokSiz
		ims = append(ims, im)
		x += blokSiz
	}
	x = r.X
	y = r.Y
	x -= blokSiz
	for y < r.Y+r.Height {
		im.rD = R(x, y, blokSiz, blokSiz)
		ims = append(ims, im)
		im.rD.X += r.Width + blokSiz
		ims = append(ims, im)
		y += blokSiz
	}
	return ims
}

// MARK: REC
func dLineRECsideCNTtoRECsideCNToffsetSHADOW(r, r2 z.Rectangle, side1234, sidetwo1234 int, offset1, offset2, offsetSHADOW, lineW float32, c, cSHADOW z.Color) {
	cnt1 := qRecCNT(r)
	cnt2 := qRecCNT(r2)

	switch side1234 {
	case 1:
		cnt1.Y -= r.Height / 2
		if offset1 > 0 {
			cnt1.Y -= offset1
		}
	case 2:
		cnt1.X += r.Width / 2
		if offset1 > 0 {
			cnt1.X += offset1
		}
	case 3:
		cnt1.Y += r.Height / 2
		if offset1 > 0 {
			cnt1.Y += offset1
		}
	case 4:
		cnt1.X -= r.Width / 2
		if offset1 > 0 {
			cnt1.X -= offset1
		}
	}
	switch sidetwo1234 {
	case 1:
		cnt2.Y -= r2.Height / 2
		if offset1 > 0 {
			cnt2.Y -= offset2
		}
	case 2:
		cnt2.X += r2.Width / 2
		if offset1 > 0 {
			cnt2.X += offset2
		}
	case 3:
		cnt2.Y += r2.Height / 2
		if offset1 > 0 {
			cnt2.Y += offset2
		}
	case 4:
		cnt2.X -= r2.Width / 2
		if offset1 > 0 {
			cnt2.X -= offset2
		}
	}

	cnt3 := cnt1
	cnt3.X -= offsetSHADOW
	cnt3.Y += offsetSHADOW
	cnt4 := cnt2
	cnt4.X -= offsetSHADOW
	cnt4.Y += offsetSHADOW
	z.DrawLineEx(cnt3, cnt4, lineW, cSHADOW)
	z.DrawLineEx(cnt1, cnt2, lineW, c)
}
func dLineRECsideCNTtoRECsideCNToffset(r, r2 z.Rectangle, side1234, sidetwo1234 int, offset1, offset2, lineW float32, c z.Color) {
	cnt1 := qRecCNT(r)
	cnt2 := qRecCNT(r2)

	switch side1234 {
	case 1:
		cnt1.Y -= r.Height / 2
		if offset1 > 0 {
			cnt1.Y -= offset1
		}
	case 2:
		cnt1.X += r.Width / 2
		if offset1 > 0 {
			cnt1.X += offset1
		}
	case 3:
		cnt1.Y += r.Height / 2
		if offset1 > 0 {
			cnt1.Y += offset1
		}
	case 4:
		cnt1.X -= r.Width / 2
		if offset1 > 0 {
			cnt1.X -= offset1
		}
	}
	switch sidetwo1234 {
	case 1:
		cnt2.Y -= r2.Height / 2
		if offset1 > 0 {
			cnt2.Y -= offset2
		}
	case 2:
		cnt2.X += r2.Width / 2
		if offset1 > 0 {
			cnt2.X += offset2
		}
	case 3:
		cnt2.Y += r2.Height / 2
		if offset1 > 0 {
			cnt2.Y += offset2
		}
	case 4:
		cnt2.X -= r2.Width / 2
		if offset1 > 0 {
			cnt2.X -= offset2
		}
	}

	z.DrawLineEx(cnt1, cnt2, lineW, c)
}
func dLineRECsideCNTtoRECsideCNT(r, r2 z.Rectangle, side1234, sidetwo1234 int, lineW float32, c z.Color) {
	cnt1 := qRecCNT(r)
	cnt2 := qRecCNT(r2)

	switch side1234 {
	case 1:
		cnt1.Y -= r.Height / 2
	case 2:
		cnt1.X += r.Width / 2
	case 3:
		cnt1.Y += r.Height / 2
	case 4:
		cnt1.X -= r.Width / 2
	}
	switch sidetwo1234 {
	case 1:
		cnt2.Y -= r2.Height / 2
	case 2:
		cnt2.X += r2.Width / 2
	case 3:
		cnt2.Y += r2.Height / 2
	case 4:
		cnt2.X -= r2.Width / 2
	}

	z.DrawLineEx(cnt1, cnt2, lineW, c)
}
func dLineRecV2toRecV2SHADOW(r, r2 z.Rectangle, point1234, pointtwo1234 int, offsetSHADOW, lineW float32, c, cSHADOW z.Color) {
	vp1 := qRecPointsREC(r)
	vp2 := qRecPointsREC(r2)
	var v1, v2 z.Vector2
	switch point1234 {
	case 1:
		v1 = vp1[0]
	case 2:
		v1 = vp1[1]
	case 3:
		v1 = vp1[2]
	case 4:
		v1 = vp1[3]
	}

	switch pointtwo1234 {
	case 1:
		v2 = vp2[0]
	case 2:
		v2 = vp2[1]
	case 3:
		v2 = vp2[2]
	case 4:
		v2 = vp2[3]
	}

	v3 := v1
	v3.X -= offsetSHADOW
	v3.Y += offsetSHADOW
	v4 := v2
	v4.X -= offsetSHADOW
	v4.Y += offsetSHADOW

	z.DrawLineEx(v3, v4, lineW, cSHADOW)
	z.DrawLineEx(v1, v2, lineW, c)
}
func dLineRecV2toRecV2SHADOWxplus(r, r2 z.Rectangle, point1234, pointtwo1234 int, offsetSHADOW, lineW float32, c, cSHADOW z.Color) {
	vp1 := qRecPointsREC(r)
	vp2 := qRecPointsREC(r2)
	var v1, v2 z.Vector2
	switch point1234 {
	case 1:
		v1 = vp1[0]
	case 2:
		v1 = vp1[1]
	case 3:
		v1 = vp1[2]
	case 4:
		v1 = vp1[3]
	}

	switch pointtwo1234 {
	case 1:
		v2 = vp2[0]
	case 2:
		v2 = vp2[1]
	case 3:
		v2 = vp2[2]
	case 4:
		v2 = vp2[3]
	}

	v3 := v1
	v3.X += offsetSHADOW
	v3.Y += offsetSHADOW
	v4 := v2
	v4.X += offsetSHADOW
	v4.Y += offsetSHADOW

	z.DrawLineEx(v3, v4, lineW, cSHADOW)
	z.DrawLineEx(v1, v2, lineW, c)
}
func dLineRecV2toRecV2(r, r2 z.Rectangle, point1234, pointtwo1234 int, lineW float32, c z.Color) {
	vp1 := qRecPointsREC(r)
	vp2 := qRecPointsREC(r2)
	var v1, v2 z.Vector2
	switch point1234 {
	case 1:
		v1 = vp1[0]
	case 2:
		v1 = vp1[1]
	case 3:
		v1 = vp1[2]
	case 4:
		v1 = vp1[3]
	}

	switch pointtwo1234 {
	case 1:
		v2 = vp2[0]
	case 2:
		v2 = vp2[1]
	case 3:
		v2 = vp2[2]
	case 4:
		v2 = vp2[3]
	}
	z.DrawLineEx(v1, v2, lineW, c)
}
func RECCNT(cnt z.Vector2, w, h float32) z.Rectangle {
	return z.NewRectangle(cnt.X-w/2, cnt.Y-h/2, w, h)
}
func R(x, y, w, h float32) z.Rectangle {
	return z.NewRectangle(x, y, w, h)
}
func RPoint(v z.Vector2, w, h float32) z.Rectangle {
	return z.NewRectangle(v.X, v.Y, w, h)
}
func dSQLineMouseOver(x, y, w float32, c1, c2 z.Color) (click, down bool) {
	click = false
	down = false
	r := R(x, y, w, w)
	if z.CheckCollisionPointRec(MS, r) {
		dRECLINE(r, 4, c2)
		if MSL {
			click = true
		}
		if MSLD {
			down = true
		}
	} else {
		dRECLINE(r, 4, c1)
	}
	return click, down
}
func dSQ(x, y, w float32, c z.Color) {
	r := z.NewRectangle(x, y, w, w)
	z.DrawRectangleRec(r, c)
}
func dRECXY(x, y, w, h float32, c z.Color) {
	r := z.NewRectangle(x, y, w, h)
	z.DrawRectangleRec(r, c)
}
func dREC(r z.Rectangle, c z.Color) {
	z.DrawRectangleRec(r, c)
}
func dRECLINE(r z.Rectangle, lineW float32, c z.Color) {
	z.DrawRectangleLinesEx(r, lineW, c)
}
func dRECSHADOWONLY(r z.Rectangle, offset float32, c z.Color, a uint8) {
	r.X -= offset
	r.Y += offset
	dREC(r, CA(c, a))
}

// MARK: GRID UP
func uGridAddBackIM2SingleRandom(g GRID, im IM, min, max float32, ro bool) GRID {
	siz := RF32(min, max)
	x := RF32(g.tl.X+UNIT*2, g.tr.X-siz-UNIT)
	y := RF32(g.tl.Y+UNIT*2, g.bl.Y-siz-UNIT)
	im.rD = z.NewRectangle(x, y, siz, siz)
	if ro {
		im.ro = RF32(0, 360)
	}
	g.bgIM2 = append(g.bgIM2, im)
	return g
}
func uGridAddBackIMRandom(g GRID, ims []IM, scatter1to1000 int, minSiz, maxSiz float32) GRID {
	for i := g.numW; i < g.numA-g.numW; i++ {
		if RINT(0, 1000) <= scatter1to1000 {
			im := ims[RINT(0, len(ims))]
			siz := RF32(minSiz, maxSiz)
			im.rD = RECCNT(g.im[i].cntRD, siz, siz)
			if !cRecIMSheetRecDrawCollis(im.rD, g.bgIM) {
				g.bgIM = append(g.bgIM, im)
				i += RINT(2, 5)
			}
		}
	}
	return g
}
func uGridIM2AlphaRANDOM(g GRID, min, max int) GRID {
	for i := range g.im2 {
		g.im2[i].a = RUINT8(min, max)
	}
	return g
}
func uGridAddIM(g GRID, im IM) {
	for i := range g.im {
		g.im[i].r = im.r
		g.im[i].tex = im.tex
	}
}

// MARK: GRID DRAW
func dGridBackIM2(g GRID) {
	if len(g.bgIM2) > 0 {
		for i := range len(g.bgIM2) {
			if g.bgIM2[i].ro != 0 {
				dIMrotating(g.bgIM2[i], g.bgIM2[i].rD, 0)
			} else {
				dIM(g.bgIM2[i], g.bgIM2[i].rD)
			}
		}
	}
}
func dGridBackIM(g GRID) {
	if len(g.bgIM) > 0 {
		for i := range g.bgIM {
			dIM(g.bgIM[i], g.bgIM[i].rD)
		}
	}
}
func dGridIM(g GRID) {
	for i := range g.im {
		dIM(g.im[i], g.im[i].rD)
	}
}
func dGridIM2(g GRID) {
	for i := range g.im2 {
		if g.im2[i].a != 0 {
			dIMA(g.im2[i], g.im2[i].rD, g.im2[i].a)
		} else {
			dIM(g.im2[i], g.im2[i].rD)
		}
	}
}
func dGridIM2color(g GRID, c z.Color) {
	for i := range g.im2 {
		if g.im2[i].a != 0 {
			dIMAcolor(g.im2[i], g.im2[i].rD, g.im2[i].a, c)
		} else {
			dIMcolor(g.im2[i], g.im2[i].rD, c)
		}
	}
}
func dGridLines(g GRID, lineW float32, c z.Color) {
	for i := range g.im {
		z.DrawRectangleLinesEx(g.im[i].rD, lineW, c)
	}
}

// MARK: GRID MAKE
func mGridCNT(cnt z.Vector2, numW, numH int, blokSiz float32) GRID {
	g := GRID{}
	g.cnt = cnt
	g.W = float32(numW) * blokSiz
	g.H = float32(numH) * blokSiz
	p := qRecPointsCNT(cnt, g.W, g.H)
	g.tl, g.tr, g.br, g.bl = p[0], p[1], p[2], p[3]
	g.r = z.NewRectangle(g.tl.X, g.tl.Y, g.W, g.H)
	g.numW = numW
	g.numH = numH
	g.numA = numH * numW

	x := g.tl.X
	ox := x
	y := g.tl.Y
	c := 0
	im := IM{}
	for range g.numA {
		im.rD = z.NewRectangle(x, y, blokSiz, blokSiz)
		im.cntRD = qRecCNT(im.rD)
		g.im = append(g.im, im)
		x += blokSiz
		c++
		if c == numW {
			c = 0
			y += blokSiz
			x = ox
		}
	}
	return g
}

// MARK: CIRC
func dCircCNT(cnt z.Vector2, rad float32, c z.Color) {
	z.DrawCircle(int32(cnt.X), int32(cnt.Y), rad, c)
}
