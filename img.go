package main

import (
	"fmt"
	"time"

	z "github.com/gen2brain/raylib-go/raylib"
)

var (
	TEX                      []z.Texture2D
	PATTERNS, TILES, ETC, UI []IM
)

type IM struct {
	tex       z.Texture2D
	r, rD, rC z.Rectangle
	a         uint8
	ro        float32
	cntRD     z.Vector2
	c         z.Color
}
type ANIM struct {
	startIMnum, endIMnum, frames, dFrame int
	ims                                  []IM
	fps                                  int
	timer                                time.Time
	off                                  bool
}

// MARK: ANIM UTILS
func RESETANIM(a ANIM) ANIM {
	a.timer = time.Time{}
	a.dFrame = 0
	a.off = false
	return a
}

// MARK: ANIM DRAW
func dAnimFrame(anm ANIM) ANIM {
	if anm.timer.IsZero() {
		anm.timer = time.Now()
	}
	if time.Since(anm.timer) >= time.Second/time.Duration(anm.fps) {
		anm.dFrame++
		if anm.dFrame > anm.frames {
			anm.dFrame = 0
		}
		anm.timer = time.Now()
	}
	return anm
}
func dAnimFrameONCE(anm ANIM) ANIM {
	if anm.timer.IsZero() {
		anm.timer = time.Now()
	}
	if time.Since(anm.timer) >= time.Second/time.Duration(anm.fps) && !anm.off {
		anm.dFrame++
		if anm.dFrame > anm.frames {
			anm.dFrame = 0
			anm.off = true
		}
		anm.timer = time.Now()
	}
	return anm
}
func dAnimRecLoop(anm ANIM, r z.Rectangle) ANIM {
	dIM(anm.ims[anm.dFrame], r)
	anm = dAnimFrame(anm)
	return anm
}
func dAnimRecLoopFlip(anm ANIM, r z.Rectangle) ANIM {
	dIMflip(anm.ims[anm.dFrame], r)
	anm = dAnimFrame(anm)
	return anm
}
func dAnimRecLoopColor(anm ANIM, r z.Rectangle, c z.Color) ANIM {
	dIMcolor(anm.ims[anm.dFrame], r, c)
	anm = dAnimFrame(anm)
	return anm
}
func dAnimRecLoopFlipColor(anm ANIM, r z.Rectangle, c z.Color) ANIM {
	dIMflipColor(anm.ims[anm.dFrame], r, c)
	anm = dAnimFrame(anm)
	return anm
}
func dAnimRecLoopShadow(anm ANIM, r z.Rectangle, cShadow z.Color, offset float32) ANIM {
	dIMshadow(anm.ims[anm.dFrame], r, cShadow, offset)
	anm = dAnimFrame(anm)
	return anm
}
func dAnimRecLoopShadowOUTLINE(anm ANIM, r z.Rectangle, cShadow z.Color, offsetShadow, offsetOutline float32) ANIM {
	r2 := qResizeRecOffsetLRGR(r, offsetOutline)
	dIMcolor(anm.ims[anm.dFrame], r2, z.Black)
	dIMshadow(anm.ims[anm.dFrame], r, cShadow, offsetShadow)
	anm = dAnimFrame(anm)
	return anm
}
func dAnimRecLoopShadowOUTLINEonce(anm ANIM, r z.Rectangle, cShadow z.Color, offsetShadow, offsetOutline float32) ANIM {
	r2 := qResizeRecOffsetLRGR(r, offsetOutline)
	dIMcolor(anm.ims[anm.dFrame], r2, z.Black)
	dIMshadow(anm.ims[anm.dFrame], r, cShadow, offsetShadow)
	anm = dAnimFrameONCE(anm)
	return anm
}
func dAnimRecLoopFlipShadowOUTLINE(anm ANIM, r z.Rectangle, cShadow z.Color, offsetShadow, offsetOutline float32) ANIM {
	r2 := qResizeRecOffsetLRGR(r, offsetOutline)
	dIMflipColor(anm.ims[anm.dFrame], r2, z.Black)
	dIMflipShadow(anm.ims[anm.dFrame], r, cShadow, offsetShadow)
	anm = dAnimFrame(anm)
	return anm
}
func dAnimRecLoopFlipShadowOUTLINEonce(anm ANIM, r z.Rectangle, cShadow z.Color, offsetShadow, offsetOutline float32) ANIM {
	r2 := qResizeRecOffsetLRGR(r, offsetOutline)
	dIMflipColor(anm.ims[anm.dFrame], r2, z.Black)
	dIMflipShadow(anm.ims[anm.dFrame], r, cShadow, offsetShadow)
	anm = dAnimFrameONCE(anm)
	return anm
}
func dAnimRecLoopFlipShadow(anm ANIM, r z.Rectangle, cShadow z.Color, offset float32) ANIM {
	dIMflipShadow(anm.ims[anm.dFrame], r, cShadow, offset)
	anm = dAnimFrame(anm)
	return anm
}
func dAnimRecFlipShadowONCE(anm ANIM, r z.Rectangle, cShadow z.Color, offset float32) ANIM {
	dIMflipShadow(anm.ims[anm.dFrame], r, cShadow, offset)
	anm = dAnimFrameONCE(anm)
	return anm
}
func dAnimRecShadowONCE(anm ANIM, r z.Rectangle, cShadow z.Color, offset float32) ANIM {
	dIMshadow(anm.ims[anm.dFrame], r, cShadow, offset)
	anm = dAnimFrameONCE(anm)
	return anm
}
func dAnimRecONCE(anm ANIM, r z.Rectangle) ANIM {
	dIM(anm.ims[anm.dFrame], r)
	anm = dAnimFrameONCE(anm)
	return anm
}
func dAnimRecONCEalpha(anm ANIM, r z.Rectangle, a uint8) ANIM {
	dIMA(anm.ims[anm.dFrame], r, a)
	anm = dAnimFrameONCE(anm)
	return anm
}
func dAnimRecColorONCE(anm ANIM, r z.Rectangle, c z.Color) ANIM {
	dIMcolor(anm.ims[anm.dFrame], r, c)
	anm = dAnimFrameONCE(anm)
	return anm
}
func dAnimXYzoom(anm ANIM, x, y, zoom float32) ANIM {
	r := z.NewRectangle(x, y, float32(anm.ims[anm.dFrame].r.Width)*zoom, float32(anm.ims[anm.dFrame].r.Height)*zoom)
	dIM(anm.ims[anm.dFrame], r)
	anm = dAnimFrame(anm)
	return anm
}

// MARK: ANIM MAKE
func mAnimSheetFiles1LINEH(path string, w, h int32, fps int) []ANIM {
	var ams []ANIM
	p := qImagePaths(path)
	for i := range p {
		ams = append(ams, mAnimSize1LINEH(p[i], w, h, fps))
	}
	return ams
}
func mAnimSize1LINEH(path string, w, h int32, fps int) ANIM {
	a := ANIM{}
	a.ims = mIMSheetSize1LINE(path, w, h)
	a.frames = len(a.ims) - 1
	a.fps = fps
	return a
}
func mAnimImSheet(ims []IM, fps int) ANIM {
	a := ANIM{}
	a.ims = ims
	a.frames = len(ims) - 1
	a.fps = fps
	return a
}
func mAnimXY(path string, x, y, siz, spc float32, numW, numH int, fps int) ANIM {
	a := ANIM{}
	a.ims = mIMSheetXY(path, x, y, siz, spc, numW, numH)
	a.frames = len(a.ims) - 1
	a.fps = fps
	return a
}
func mAnimXYWH(path string, x, y, w, h, spc float32, numW, numH int, fps int) ANIM {
	a := ANIM{}
	a.ims = mIMSheetXYWH(path, x, y, w, h, spc, numW, numH)
	a.frames = len(a.ims) - 1
	a.fps = fps
	return a
}
func mAnimXYWHsubtract(path string, x, y, w, h, spc float32, numW, numH, subtract int, fps int) ANIM {
	a := ANIM{}
	a.ims = mIMSheetXYWHsubtract(path, x, y, w, h, spc, numW, numH, subtract)
	a.frames = len(a.ims) - 1
	a.fps = fps
	return a
}

// MARK: IM
func dIMro(im IM, r z.Rectangle, angl float32) {
	z.DrawTexturePro(im.tex, im.r, DREC(r), qORIGIN(r), angl, z.White)
}
func dIMroColor(im IM, r z.Rectangle, angl float32, c z.Color) {
	z.DrawTexturePro(im.tex, im.r, DREC(r), qORIGIN(r), angl, c)
}
func dIMroSHADOW(im IM, r z.Rectangle, angl, offsetShadow float32, cShadow z.Color) {
	r2 := qShadowREC(r, offsetShadow)
	z.DrawTexturePro(im.tex, im.r, DREC(r2), qORIGIN(r2), angl, cShadow)
	z.DrawTexturePro(im.tex, im.r, DREC(r), qORIGIN(r), angl, z.White)
}
func dIMroColorSHADOW(im IM, r z.Rectangle, angl, offsetShadow float32, c, cShadow z.Color) {
	r2 := qShadowREC(r, offsetShadow)
	z.DrawTexturePro(im.tex, im.r, DREC(r2), qORIGIN(r2), angl, cShadow)
	z.DrawTexturePro(im.tex, im.r, DREC(r), qORIGIN(r), angl, c)
}
func dIMrotating(im IM, r z.Rectangle, roChange float32) IM {
	z.DrawTexturePro(im.tex, im.r, DREC(r), qORIGIN(r), im.ro, z.White)
	im.ro += roChange
	return im
}
func dIMrotatingColor(im IM, r z.Rectangle, roChange float32, c z.Color) IM {
	z.DrawTexturePro(im.tex, im.r, DREC(r), qORIGIN(r), im.ro, c)
	im.ro += roChange
	return im
}
func dIMrotatingColorSHADOW(im IM, r z.Rectangle, roChange, offsetShadow float32, c, cShadow z.Color) IM {
	r2 := r
	r2.X -= offsetShadow
	r2.Y += offsetShadow
	z.DrawTexturePro(im.tex, im.r, DREC(r2), qORIGIN(r2), im.ro, cShadow)
	z.DrawTexturePro(im.tex, im.r, DREC(r), qORIGIN(r), im.ro, c)
	im.ro += roChange
	return im
}
func dIM(im IM, r z.Rectangle) {
	z.DrawTexturePro(im.tex, im.r, r, z.Vector2Zero(), 0, z.White)
}
func dIMoutline(im IM, r z.Rectangle, offsetOutline float32, cOutline z.Color) {
	r2 := qResizeRecOffsetLRGR(r, offsetOutline)
	dIMcolor(im, r2, cOutline)
	dIM(im, r)
}
func dIMcolorOutline(im IM, r z.Rectangle, offsetOutline float32, c, cOutline z.Color) {
	r2 := qResizeRecOffsetLRGR(r, offsetOutline)
	dIMcolor(im, r2, cOutline)
	dIMcolor(im, r, c)
}
func dIMoriginColor(im IM, r z.Rectangle, c z.Color) {
	z.DrawTexturePro(im.tex, im.r, r, qORIGIN(r), 0, c)
}
func dIMorigin(im IM, r z.Rectangle) {
	z.DrawTexturePro(im.tex, im.r, r, qORIGIN(r), 0, z.White)
}
func dIMshadow(im IM, r z.Rectangle, cShadow z.Color, offset float32) {
	r2 := r
	r2.X -= offset
	r2.Y += offset
	z.DrawTexturePro(im.tex, im.r, r2, z.Vector2Zero(), 0, cShadow)
	z.DrawTexturePro(im.tex, im.r, r, z.Vector2Zero(), 0, z.White)
}
func dIMflip(im IM, r z.Rectangle) {
	im.r.Width = -im.r.Width
	z.DrawTexturePro(im.tex, im.r, r, z.Vector2Zero(), 0, z.White)
}
func dIMflipColor(im IM, r z.Rectangle, c z.Color) {
	im.r.Width = -im.r.Width
	z.DrawTexturePro(im.tex, im.r, r, z.Vector2Zero(), 0, c)
}
func dIMflipShadow(im IM, r z.Rectangle, cShadow z.Color, offset float32) {
	im.r.Width = -im.r.Width
	r2 := r
	r2.X += offset
	r2.Y += offset
	z.DrawTexturePro(im.tex, im.r, r2, z.Vector2Zero(), 0, cShadow)
	z.DrawTexturePro(im.tex, im.r, r, z.Vector2Zero(), 0, z.White)
}
func dIMA(im IM, r z.Rectangle, a uint8) {
	z.DrawTexturePro(im.tex, im.r, r, z.Vector2Zero(), 0, CA(z.White, a))
}
func dIMcolor(im IM, r z.Rectangle, c z.Color) {
	z.DrawTexturePro(im.tex, im.r, r, z.Vector2Zero(), 0, c)
}
func dIMAcolorSHADOW(im IM, r z.Rectangle, c, cShadow z.Color, offsetShadow float32, a uint8) {
	r2 := r
	r2.X -= offsetShadow
	r2.Y += offsetShadow
	z.DrawTexturePro(im.tex, im.r, r2, z.Vector2Zero(), 0, cShadow)
	z.DrawTexturePro(im.tex, im.r, r, z.Vector2Zero(), 0, CA(c, a))
}
func dIMcolorSHADOW(im IM, r z.Rectangle, c, cShadow z.Color, offsetShadow float32) {
	r2 := r
	r2.X -= offsetShadow
	r2.Y += offsetShadow
	z.DrawTexturePro(im.tex, im.r, r2, z.Vector2Zero(), 0, cShadow)
	z.DrawTexturePro(im.tex, im.r, r, z.Vector2Zero(), 0, c)
}
func dIMAcolor(im IM, r z.Rectangle, a uint8, c z.Color) {
	z.DrawTexturePro(im.tex, im.r, r, z.Vector2Zero(), 0, CA(c, a))
}
func dIMAXY(im IM, x, y, w, h, a float32) {
	r := R(x, y, w, h)
	z.DrawTexturePro(im.tex, im.r, r, z.Vector2Zero(), 0, z.Fade(z.White, a))
}
func dIMRECproportionalHcolor(im IM, r z.Rectangle, c z.Color) {
	w := qWidthProportional(im.r.Width, im.r.Height, r.Height)
	r2 := RECCNT(qRecCNT(r), w, r.Height)
	dIMcolor(im, r2, c)
}
func dIMRECproportionalHcolorSHADOW(im IM, r z.Rectangle, c, cShadow z.Color, offsetShadow float32) {
	w := qWidthProportional(im.r.Width, im.r.Height, r.Height)
	r2 := RECCNT(qRecCNT(r), w, r.Height)
	r3 := r2
	r3.X -= offsetShadow
	r3.Y += offsetShadow
	dIMcolor(im, r3, cShadow)
	dIMcolor(im, r2, c)
}

// MARK: IMSHEET UPDATE
func uIMSheetAlphaRANDOM(ims []IM, min, max int) []IM {
	for i := range ims {
		ims[i].a = RUINT8(min, max)
	}
	return ims
}

// MARK: IMSHEET DRAW
func dIMSheet(ims []IM, x, y, spc, zoom float32) {
	ox := x
	h := float32(0)
	for i := range ims {
		r := qResizeRecXYZERO(ims[i].r, zoom)
		if x+r.Width > float32(SCRW) {
			x = ox
			y += h + spc
			if debug {
				y += 10
			}
			h = 0
		}
		r.X = x
		r.Y = y
		dIM(ims[i], r)
		if debug {
			dTXTDEF10C(fmt.Sprint(i), r.ToInt32().X+2, r.ToInt32().Y+r.ToInt32().Height, z.Green)
		}

		if y+r.Height > float32(SCRH) {
			scrollONOFF = true
		} else {
			scrollONOFF = false
		}
		if h < r.Height {
			h = r.Height
		}
		x += r.Width + spc
		if x+r.Width > float32(SCRW) {
			x = ox
			y += h + spc
			if debug {
				y += 10
			}
			h = 0
		}

	}
}
func dIMSheetDrawRec(ims []IM) {
	for i := range ims {
		dIM(ims[i], ims[i].rD)
	}
}
func dIMSheetDrawRecColor(ims []IM, c z.Color) {
	for i := range ims {
		dIMcolor(ims[i], ims[i].rD, c)
	}
}
func dIMSheetDrawRecCustomColorAlpha(ims []IM, c z.Color) {
	for i := range ims {
		dIMAcolor(ims[i], ims[i].rD, ims[i].a, c)
	}
}
func dIMSheetDrawRecColorAlpha(ims []IM) {
	for i := range ims {
		dIMAcolor(ims[i], ims[i].rD, ims[i].a, ims[i].c)
	}
}

// MARK: IMSHEET MAKE
func mIMSheetSize1LINE(path string, w, h int32) []IM {
	var ims []IM
	TEX = append(TEX, z.LoadTexture(path))
	var x, y float32
	im := IM{}
	im.tex = TEX[len(TEX)-1]
	numW := im.tex.Width / w
	for numW > 0 {
		im.r = z.NewRectangle(x, y, float32(w), float32(h))
		ims = append(ims, im)
		x += float32(w)
		numW--
	}
	return ims
}
func mIMSheetFiles(path string) []IM {
	var ims []IM
	p := qImagePaths(path)
	im := IM{}
	for i := range p {
		t := z.LoadTexture(p[i])
		im.tex = t
		im.r = z.NewRectangle(0, 0, float32(t.Width), float32(t.Height))
		ims = append(ims, im)
	}
	return ims
}
func mIMSheetXY(path string, x, y, siz, spc float32, numW, numH int) []IM {
	var ims []IM
	TEX = append(TEX, z.LoadTexture(path))
	a := numW * numH
	im := IM{}
	im.tex = TEX[len(TEX)-1]
	ox := x
	c := 0
	for a > 0 {
		im.r = z.NewRectangle(x, y, siz, siz)
		ims = append(ims, im)
		x += siz + spc
		c++
		a--
		if c == numW {
			c = 0
			x = ox
			y += siz + spc
		}
	}
	return ims
}
func mIMSheetXYWH(path string, x, y, w, h, spc float32, numW, numH int) []IM {
	var ims []IM
	TEX = append(TEX, z.LoadTexture(path))
	a := numW * numH
	im := IM{}
	im.tex = TEX[len(TEX)-1]
	ox := x
	c := 0
	for a > 0 {
		im.r = z.NewRectangle(x, y, w, h)
		ims = append(ims, im)
		x += w + spc
		c++
		a--
		if c == numW {
			c = 0
			x = ox
			y += h + spc
		}
	}
	return ims
}
func mIMSheetXYWHsubtract(path string, x, y, w, h, spc float32, numW, numH, subtract int) []IM {
	var ims []IM
	TEX = append(TEX, z.LoadTexture(path))
	a := numW * numH
	a -= subtract
	im := IM{}
	im.tex = TEX[len(TEX)-1]
	ox := x
	c := 0
	for a > 0 {
		im.r = z.NewRectangle(x, y, w, h)
		ims = append(ims, im)
		x += w + spc
		c++
		a--
		if c == numW {
			c = 0
			x = ox
			y += h + spc
		}
	}
	return ims
}
