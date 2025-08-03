package main

import z "github.com/gen2brain/raylib-go/raylib"

type OBJ struct {
	r, rD, rC, r2 z.Rectangle

	nm  string
	anm []ANIM
	lr  bool
	im  IM

	cnt, cnt2        z.Vector2
	sz, rad          float32
	state, statePREV int

	atk, action1, action2 bool
	off, onoff1, onoff2   bool

	atkSPD, atkT int

	velX, velY, velMAX, velORG, ro float32
}
