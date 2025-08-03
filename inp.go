package main

import (
	z "github.com/gen2brain/raylib-go/raylib"
)

var (
	MS                                    z.Vector2
	MSL, MSR, MSM, MSLD, MSRD, MSMD, MSL2 bool
	kL, kR, kU, kD                        bool

	clickT int
	clickP = int(FPS) / 6

	optionsON bool

	cursorREC z.Rectangle
)

func INP() {
	//MOVE
	if z.IsKeyDown(z.KeyRight) || z.IsKeyDown(z.KeyD) {
		kR = true
	} else {
		kR = false
	}
	if z.IsKeyDown(z.KeyLeft) || z.IsKeyDown(z.KeyA) {
		kL = true
	} else {
		kL = false
	}
	if z.IsKeyDown(z.KeyUp) || z.IsKeyDown(z.KeyW) {
		kU = true
	} else {
		kU = false
	}
	if z.IsKeyDown(z.KeyDown) || z.IsKeyDown(z.KeyS) {
		kD = true
	} else {
		kD = false
	}

	//GAME
	if z.IsKeyPressed(z.KeyEscape) {
		optionsON = !optionsON
	}
	if z.IsKeyPressed(z.KeyF2) {
		RESTART()
	}

	//DEBUG
	if z.IsKeyPressed(z.KeyF1) {
		debug = !debug
	}
	if z.IsKeyPressed(z.KeyF2) && debug {
		debug2 = !debug2
	}
	if z.IsKeyPressed(z.KeyF3) {
		playerdeathONOFF = true
	}

	if z.IsKeyPressed(z.KeyF4) {
		MAPON = !MAPON
	}

	//MOUSE
	MS = z.GetMousePosition()
	if z.IsMouseButtonPressed(z.MouseButtonLeft) {
		MSL = true
	} else {
		MSL = false
	}
	if z.IsMouseButtonDown(z.MouseButtonLeft) {
		MSLD = true
	} else {
		MSLD = false
	}

	if z.IsMouseButtonPressed(z.MouseButtonLeft) && clickT == 0 {
		MSL2 = true
	} else {
		MSL2 = false
	}

	//CLICK TIMER
	if MSL && clickT == 0 {
		clickT = clickP
	}

}
