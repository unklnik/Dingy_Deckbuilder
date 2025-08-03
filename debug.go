package main

import (
	"fmt"

	z "github.com/gen2brain/raylib-go/raylib"
)

var (
	debug, debug2, debugBOOL bool
	debugNUM                 int
)

func DEBUG() {

	debugBOOL = FULLSCREEN

	if !debug2 {
		dREC(z.NewRectangle(0, 0, 300, float32(SCRH)), CA(z.Maroon, 150))
	}
	var x, y, sizTXT float32 = 4, 4, 0.5
	h := qTXTheight(FONT1, sizTXT)
	dTXTfont1XY("mouseX "+fmt.Sprintf("%.0f", MS.X)+" mouseY "+fmt.Sprintf("%.0f", MS.Y), x, y, sizTXT, z.White)
	y += h
	dTXTfont1XY("debugNUM "+fmt.Sprint(debugNUM)+" debugBOOL "+fmt.Sprint(debugBOOL), x, y, sizTXT, z.White)
	y += h
	dTXTfont1XY("UNIT "+fmt.Sprint(UNIT)+" pl.drawnum "+fmt.Sprint(pl.drawNUM), x, y, sizTXT, z.White)
	y += h
	dTXTfont1XY("pllevmapNUM "+fmt.Sprint(pllevmapNUM)+" pl.drawnum "+fmt.Sprint(pl.drawNUM), x, y, sizTXT, z.White)
	y += h
}
