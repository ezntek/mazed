package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/tek967/mazed/src/ball"
	"github.com/tek967/mazed/src/block"
	pl "github.com/tek967/rgbapalette"
)

const (
	width  int32 = 1156
	height int32 = 864
)

var theBall *ball.Ball = ball.New(rl.NewVector2(50, 50), 20, rl.Black)
var blocks = []block.Block{*block.New(rl.NewVector2(0, 300), rl.NewVector2(100, 100))}

func beforeLoopSetup() {
	rl.InitWindow(width, height, "wow game")
}

func draw() {
	theBall.Draw()
	for _, block := range blocks {
		block.Draw()
	}
}

func update() {
	theBall.Update(width, height, &blocks)

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
