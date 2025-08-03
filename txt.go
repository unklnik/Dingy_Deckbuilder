package main

import (
	z "github.com/gen2brain/raylib-go/raylib"
)

var (
	FONT1, FONT2 z.Font
)

func mFONTS() {
	FONT1 = z.LoadFont("data/Rubik-Medium.ttf")
	FONT2 = z.LoadFont("data/Rancho-Regular.ttf")
}

// UTILS
func qTXTlenFont1(t string, siz float32) z.Vector2 {
	return z.MeasureTextEx(FONT1, t, float32(FONT1.BaseSize)*siz, 0)
}
func qTXTlenFont2(t string, siz float32) z.Vector2 {
	return z.MeasureTextEx(FONT2, t, float32(FONT1.BaseSize)*siz, 0)
}
func qTXTheight(f z.Font, siz float32) float32 {
	v2 := z.MeasureTextEx(f, "Text", float32(f.BaseSize)*siz, 0)
	return v2.Y
}

// DRAW
func dTXTfont1XYbgrec(t string, x, y, siz float32, cTXT, cBG z.Color, offsetBGREC float32) {
	v2 := qTXTlenFont1(t, siz)
	r := z.NewRectangle(x-offsetBGREC, y-offsetBGREC/2, v2.X+offsetBGREC*2, v2.Y+offsetBGREC)
	dREC(r, cBG)
	dTXTfont1XY(t, x, y, siz, cTXT)
}
func dTXTfont2XYbgrec(t string, x, y, siz float32, cTXT, cBG z.Color, offsetBGREC float32) {
	v2 := qTXTlenFont2(t, siz)
	r := z.NewRectangle(x-offsetBGREC, y-offsetBGREC/2, v2.X+offsetBGREC*2, v2.Y+offsetBGREC)
	dREC(r, cBG)
	dTXTfont2XY(t, x, y, siz, cTXT)
}
func dTXTcntXFont2(t string, xMiddle, y, siz float32, c z.Color) {
	v2 := qTXTlenFont2(t, siz)
	x := xMiddle - v2.X/2
	dTXTfont2XY(t, x, y, siz, c)
}
func dTXTcntRecYToffsetFont2(t string, r z.Rectangle, siz, offset float32, c z.Color) {
	v2 := qTXTlenFont2(t, siz)
	x := (r.X + r.Width/2) - v2.X/2
	y := r.Y - offset - qTXTheight(FONT2, siz)
	dTXTfont2XY(t, x, y, siz, c)
}
func dTXTcntRecYBoffsetFont2(t string, r z.Rectangle, siz, offset float32, c z.Color) {
	v2 := qTXTlenFont2(t, siz)
	x := (r.X + r.Width/2) - v2.X/2
	y := r.Y + r.Height + offset
	dTXTfont2XY(t, x, y, siz, c)
}
func dTXTcntRecWfont1(t string, r z.Rectangle, offsetY, siz float32, c z.Color) {
	v2 := qTXTlenFont1(t, siz)
	x := (r.X + r.Width/2) - v2.X/2
	y := r.Y + offsetY
	dTXTfont1XY(t, x, y, siz, c)
}
func dTXTcntRecWfont2(t string, r z.Rectangle, offsetY, siz float32, c z.Color) {
	v2 := qTXTlenFont2(t, siz)
	x := (r.X + r.Width/2) - v2.X/2
	y := r.Y + offsetY
	dTXTfont2XY(t, x, y, siz, c)
}
func dTXTcntRecFont2(t string, r z.Rectangle, siz float32, c z.Color) {
	v2 := qTXTlenFont2(t, siz)
	x := (r.X + r.Width/2) - v2.X/2
	y := (r.Y + r.Height/2) - v2.Y/2
	dTXTfont2XY(t, x, y, siz, c)
}

func dTXTcntRecFont2SHADOW(t string, r z.Rectangle, siz float32, c, cShadow z.Color, offsetShadow float32) {
	v2 := qTXTlenFont2(t, siz)
	x := (r.X + r.Width/2) - v2.X/2
	y := (r.Y + r.Height/2) - v2.Y/2
	x2 := x - offsetShadow
	y2 := y + offsetShadow
	dTXTfont2XY(t, x2, y2, siz, cShadow)
	dTXTfont2XY(t, x, y, siz, c)
}
func dTXTcntRecFont1(t string, r z.Rectangle, siz float32, c z.Color) {
	v2 := qTXTlenFont1(t, siz)
	x := (r.X + r.Width/2) - v2.X/2
	y := (r.Y + r.Height/2) - v2.Y/2
	dTXTfont1XY(t, x, y, siz, c)
}
func dTXTcntRecWfont2SHADOW(t string, r z.Rectangle, offsetY, offsetShadow, siz float32, c, cShadow z.Color) {
	w := qTXTlenFont2(t, siz)
	x := (r.X + r.Width/2) - w.X/2
	y := r.Y + offsetY
	x2 := x - offsetShadow
	y2 := y + offsetShadow
	dTXTfont2XY(t, x2, y2, siz, CA(cShadow, 150))
	dTXTfont2XY(t, x, y, siz, c)
}
func dTXTfont1XY(t string, x, y, siz float32, c z.Color) {
	z.DrawTextEx(FONT1, t, z.NewVector2(x, y), float32(FONT1.BaseSize)*siz, 0, c)
}
func dTXTfont2XY(t string, x, y, siz float32, c z.Color) {
	z.DrawTextEx(FONT2, t, z.NewVector2(x, y), float32(FONT2.BaseSize)*siz, 0, c)
}
func dTXTfont2XYSHADOW(t string, x, y, siz float32, c, cShadow z.Color, offsetShadow float32) {
	z.DrawTextEx(FONT2, t, z.NewVector2(x-offsetShadow, y+offsetShadow), float32(FONT2.BaseSize)*siz, 0, cShadow)
	z.DrawTextEx(FONT2, t, z.NewVector2(x, y), float32(FONT2.BaseSize)*siz, 0, c)
}
func dTXTDEFF32(t string, x, y float32) {
	z.DrawText(t, int32(x), int32(y), 20, z.White)
}
func dTXTDEF10(t string, x, y int32) {
	z.DrawText(t, x, y, 10, z.White)
}
func dTXTDEF10C(t string, x, y int32, c z.Color) {
	z.DrawText(t, x, y, 10, c)
}
func dTXTDEF20(t string, x, y int32) {
	z.DrawText(t, x, y, 20, z.White)
}
func dTXTDEF20C(t string, x, y int32, c z.Color) {
	z.DrawText(t, x, y, 20, c)
}
