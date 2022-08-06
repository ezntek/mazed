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

func (b *Ball) collisionWithBlock(block *block.Block) string {
	var toCheck = [4]string{"left", "right", "up", "down"}
	for _, check := range toCheck {
		if block.Collision(check, b.Position, b.Radius) {
			return check
		}
	}
	return ""
}

func (b *Ball) movement() {
	if b.direction.up {
		b.Velocity.Y -= 0.05
	} else if b.direction.down {
		b.Velocity.Y += 0.05
	} else if b.direction.left {
		b.Velocity.X -= 0.05
	} else if b.direction.right {
		b.Velocity.X += 0.05
	}
}

func (b *Ball) collision(sw, sh int32, blockList *[]block.Block) {
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

	for _, block := range *blockList {
		switch b.collisionWithBlock(&block) {
		case "left":
			if b.direction.left {
				b.Position.X = block.Position.X - b.Radius
				b.Velocity.X = 0
				b.direction.left = false
			}
		case "right":
			if b.direction.right {
				b.Position.X = block.Position.X + block.Rect.Width + b.Radius
				b.Velocity.X = 0
				b.direction.right = false
			}
		case "up":
			if b.direction.up {
				b.Velocity.Y = 0
				b.direction.right = false
				b.Position.Y = block.Position.Y - b.Radius
			}
		case "down":
			if b.direction.down {
				b.Velocity.Y = 0
				b.direction.down = false
				b.Position.Y = block.Position.Y + block.Rect.Height + b.Radius
			}
		}
	}
}

func (b *Ball) Update(sw, sh int32, blockList *[]block.Block) {
	b.collision(sw, sh, blockList)

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
