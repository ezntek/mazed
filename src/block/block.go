package block

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Block struct {
	Rect     rl.Rectangle
	Position rl.Vector2
}

func (b *Block) Collision(side string, pos rl.Vector2, rad float32) bool {
	if !rl.CheckCollisionCircleRec(pos, rad, b.Rect) {
		return false
	}
	switch side {
	case "up":
		if rl.CheckCollisionCircleRec(pos, rad, b.Rect) && pos.Y+rad < b.Rect.Y && pos.X > b.Rect.X && b.Rect.X+b.Rect.Width < pos.X {
			fmt.Println("BLOCK UP")
			return true
		}
	case "down":
		if rl.CheckCollisionCircleRec(pos, rad, b.Rect) && pos.Y-rad > b.Rect.Y+b.Rect.Height && pos.X > b.Rect.X && b.Rect.X+b.Rect.Width < pos.X {
			fmt.Println("BLOCK DOWN")
			return true
		}
	case "left":
		if rl.CheckCollisionCircleRec(pos, rad, b.Rect) && pos.X < b.Rect.X && pos.Y > b.Rect.Y && b.Rect.Y+b.Rect.Height < pos.Y {
			fmt.Println("BLOCK LEFT")
			return true
		}
	case "right":
		if rl.CheckCollisionCircleRec(pos, rad, b.Rect) && pos.X > b.Rect.X+b.Rect.Width && pos.Y > b.Rect.Y && b.Rect.Y+b.Rect.Height < pos.Y {
			fmt.Println("BLOCK RIGHT")
			return true
		}
	}
return false}

func (b *Block) Draw() {
	b.Rect.X, b.Rect.Y = b.Position.X, b.Position.Y
	b.Position.X, b.Position.Y = b.Rect.X, b.Rect.Y
	rl.DrawRectangleRec(b.Rect, rl.Gray)
}
func New(pos rl.Vector2, size rl.Vector2) *Block {
	return &Block{
		Position: pos,
		Rect:     rl.NewRectangle(pos.X, pos.Y, size.X, size.Y),
	}
}
