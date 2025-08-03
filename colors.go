package main

import z "github.com/gen2brain/raylib-go/raylib"

func GREENdark1() z.Color {
	return z.NewColor(6, 64, 43, 255)
}

func CA(c z.Color, a uint8) z.Color {
	c.A = a
	return c
}
func CRGB(r, g, b uint8) z.Color {
	return z.NewColor(r, g, b, 255)
}
