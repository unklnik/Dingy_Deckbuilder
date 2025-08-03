package main

import z "github.com/gen2brain/raylib-go/raylib"

var ()

func SCAN(lineW, spc float32, c z.Color) {
	var x, y float32
	for y < float32(SCRH) {
		v1 := z.NewVector2(x, y)
		v2 := z.NewVector2(x+float32(SCRW), y)
		z.DrawLineEx(v1, v2, lineW, c)
		y += lineW + spc
	}
}
