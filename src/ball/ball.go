package ball

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Ball struct {
	Position rl.Vector2
	Velocity rl.Vector2
	Radius   float32
	Color    color.RGBA
	isHidden bool
}

func (b *Ball) Draw() {
	if !b.isHidden {
		rl.DrawCircle(int32(b.Position.X), int32(b.Position.Y), b.Radius, b.Color)
	}
}

func (b *Ball) collision(sw, sh int32) {
	if b.Position.X-b.Radius > float32(sw) {
		b.Position.X = float32(sw) - b.Radius
	} else if b.Position.X < b.Radius {
		b.Position.X = b.Radius
	}

	if b.Position.Y-b.Radius > float32(sh) {
		b.Position.Y = float32(sh) - b.Radius
	} else if b.Position.Y < b.Radius {
		b.Position.Y = b.Radius
	}
}

func (b *Ball) Update(sw, sh int32) {
	b.collision(sw, sh)

	if rl.IsKeyDown(rl.KeyUp) {

	}
	if rl.IsKeyDown(rl.KeyDown) {

	}
	if rl.IsKeyDown(rl.KeyLeft) {

	}
	if rl.IsKeyDown(rl.KeyRight) {

	}

	b.Position.X += b.Velocity.X
	b.Position.Y += b.Velocity.Y
}

func New(pos rl.Vector2, rad float32, color color.RGBA) *Ball {
	return &Ball{
		Position: pos,
		Radius:   rad,
		Color:    color,
		isHidden: false,
	}
}
