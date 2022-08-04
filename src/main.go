package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/tek967/mazed/src/ball"
	pl "github.com/tek967/rgbapalette"
)

const (
	width  int32 = 1156
	height int32 = 864
)

var theBall *ball.Ball = ball.New(rl.NewVector2(50, 50), 20, rl.Black)

func beforeLoopSetup() {
	rl.InitWindow(width, height, "wow game")
}

func draw() {
	theBall.Draw()
}

func update() {
	theBall.Update(width, height)

	rl.BeginDrawing()
	rl.ClearBackground(pl.Palette["verylightgray"])
	draw()
	rl.EndDrawing()
}

func main() {
	beforeLoopSetup()
	for !rl.WindowShouldClose() {
		update()
	}
}
