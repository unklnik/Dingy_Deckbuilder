package main

import (
	z "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	z.InitWindow(SCRW, SCRH, "GAME NAME")
	z.SetConfigFlags(z.FlagMsaa4xHint | z.FlagVsyncHint)
	z.SetTraceLogLevel(z.LogError)
	z.SetBlendMode(z.BlendAlpha)
	z.SetTargetFPS(FPS)
	INITIAL()

	mLEV()

	for !z.WindowShouldClose() {

		z.BeginDrawing()

		z.ClearBackground(z.Black)
		z.BeginTextureMode(TEXRND)
		z.ClearBackground(z.Black)
		z.BeginMode2D(CAM)
		DRAWCAM()

		z.EndMode2D()
		DRAWNOCAMSHADER()
		z.EndTextureMode()

		z.BeginShaderMode(SHADER)
		z.DrawTextureRec(TEXRND.Texture, z.NewRectangle(0, 0, float32(TEXRND.Texture.Width), float32(-TEXRND.Texture.Height)), z.NewVector2(0, 0), z.White)
		z.EndShaderMode()

		DRAWNOCAMNOSHADER()

		z.EndDrawing()

		UP()
	}
	UNLOAD()
	z.CloseWindow()
}
