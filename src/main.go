package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	pl "github.com/tek967/rgbapalette"
)

const (
	width  int32 = 1156
	height int32 = 864
)

func draw() {
	rl.DrawText("it works!", 20, 20, 20, rl.Black)
}
func update() {
	rl.BeginDrawing()
	rl.ClearBackground(pl.Palette["verylightgray"])
	draw()
	rl.EndDrawing()
}

func main() {
	rl.InitWindow(width, height, "wow game")
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		update()
	}
}
