package ball

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/tek967/mazed/src/block"
)

type movingInDirection struct {
	left, right, up, down bool
}

type Ball struct {
	direction movingInDirection
	Position  rl.Vector2
	Velocity  rl.Vector2
	Radius    float32
	Color     color.RGBA
	isHidden  bool
}

func (b *Ball) Draw() {
	if !b.isHidden {
		rl.DrawCircle(int32(b.Position.X), int32(b.Position.Y), b.Radius, b.Color)
	}
}

func (b *Ball) movement() {
	if b.direction.up {
		b.Velocity.Y -= 0.25
	} else if b.direction.down {
		b.Velocity.Y += 0.25
	} else if b.direction.left {
		b.Velocity.X -= 0.25
	} else if b.direction.right {
		b.Velocity.X += 0.25
	}
}

func (b *Ball) collision(sw, sh int32) {
	if b.direction.left || b.direction.right {
		if b.Position.X > float32(sw)-b.Radius {
			b.Position.X = float32(sw) - b.Radius
			b.Velocity.X = 0
			b.direction.right = false
		} else if b.Position.X < b.Radius {
			b.Position.X = b.Radius
			b.Velocity.X = 0
			b.direction.left = false
		}
	}

	if b.direction.up || b.direction.down {
		if b.Position.Y > float32(sh)-b.Radius {
			b.Position.Y = float32(sh) - b.Radius
			b.Velocity.Y = 0
			b.direction.down = false
		} else if b.Position.Y < b.Radius {
			b.Position.Y = b.Radius
			b.Velocity.Y = 0
			b.direction.up = false
		}
	}
}

func (b *Ball) Update(sw, sh int32, blockList *[]block.Block) {
	b.collision(sw, sh)

	if !b.direction.up && !b.direction.down && !b.direction.left && !b.direction.right {
		if rl.IsKeyPressed(rl.KeyUp) {
			b.direction.up = true
		}
		if rl.IsKeyPressed(rl.KeyDown) {
			b.direction.down = true
		}
		if rl.IsKeyPressed(rl.KeyLeft) {
			b.direction.left = true
		}
		if rl.IsKeyPressed(rl.KeyRight) {
			b.direction.right = true
		}
	}

	b.movement()
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
