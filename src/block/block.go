package block

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/tek967/mazed/src/ball"
)

type Block struct {
	rect     rl.Rectangle
	Position rl.Vector2
}

func (b *Block) Draw() {
	b.rect.X, b.rect.Y = b.Position.X, b.Position.Y
	rl.DrawRectangleRec(b.rect, rl.Gray)
}

func (b *Block) CollisionWithBall(ball *ball.Ball) string {
	if rl.CheckCollisionCircleRec(ball.Position, ball.Radius, b.rect) {
		// do some magic to find out which side
		if ball.Position.X > b.Position.X && ball.Position.X < b.Position.X+b.rect.Width { // within X range
			if ball.Position.Y < b.Position.Y {
				return "above"
			} else if ball.Position.Y > b.Position.Y+b.rect.Height {
				return "below"
			}
		} else if ball.Position.Y > b.Position.Y && b.Position.Y < b.Position.Y+b.rect.Height { // within Y range
			if ball.Position.X < b.Position.X {
				return "left"
			} else if ball.Position.X > b.Position.X+b.rect.Width {
				return "right"
			}
		}
	} else {
		return "false"
	}
	return ""
}

func New(pos rl.Vector2, size rl.Vector2) *Block {
	return &Block{
		Position: pos,
		rect:     rl.NewRectangle(pos.X, pos.Y, size.X, size.Y),
	}
}
