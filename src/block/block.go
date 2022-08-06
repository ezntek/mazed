package block

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Block struct {
	Rect                                     rl.Rectangle
	leftRect, rightRect, topRect, bottomRect rl.Rectangle
	Position                                 rl.Vector2
}

func (b *Block) Collision(side string, pos rl.Vector2, rad float32) bool {
	switch side {
	case "left":
		if rl.CheckCollisionCircleRec(pos, rad, b.leftRect) {
			return true
		}
	case "right":
		if rl.CheckCollisionCircleRec(pos, rad, b.rightRect) {
			return true
		}
	case "up":
		if rl.CheckCollisionCircleRec(pos, rad, b.topRect) {
			return true
		}
	case "down":
		if rl.CheckCollisionCircleRec(pos, rad, b.bottomRect) {
			return true
		}
	}
	return false
}

func (b *Block) Draw() {
	b.Rect.X, b.Rect.Y = b.Position.X, b.Position.Y
	b.Position.X, b.Position.Y = b.Rect.X, b.Rect.Y
	rl.DrawRectangleRec(b.Rect, rl.Gray)
}

func New(pos rl.Vector2, size rl.Vector2) *Block {
	return &Block{
		Position:   pos,
		Rect:       rl.NewRectangle(pos.X, pos.Y, size.X, size.Y),
		topRect:    rl.NewRectangle(pos.X, pos.Y, size.X, size.Y/4),
		bottomRect: rl.NewRectangle(pos.X, pos.Y+(size.Y*0.75), size.X, size.Y/4),
		leftRect:   rl.NewRectangle(pos.X, pos.Y, size.X/4, size.Y),
		rightRect:  rl.NewRectangle(pos.X+(0.75*size.X), pos.Y, size.X/4, size.Y),
	}
}
