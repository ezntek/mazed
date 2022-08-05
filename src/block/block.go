package block

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Block struct {
	Rect     rl.Rectangle
	Position rl.Vector2
}

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
