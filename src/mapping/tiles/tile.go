package tiles

import (
	"github.com/gen2brain/raylib-go/raylib"
	"anim"
)

type Tile interface {
	Draw()
	Update(float32)
	GetResistance() float32
	GetPos() rl.Vector2
	GetDimensions() (float32, float32)
}

type BaseTile struct {
	Pos rl.Vector2
	Resistance float32 // Deceleration multiplier when on tile
	texture *rl.Texture2D
	W float32
	H float32
}

func (t *BaseTile) Draw() {
	rl.DrawTexturePro(*t.texture, rl.NewRectangle(0, 0, t.W, t.H), rl.NewRectangle(t.Pos.X, t.Pos.Y, t.W, t.H), rl.NewVector2(0, 0), 0, rl.White)
}

func (t *BaseTile) Update(dt float32) {}   // So it works with interface

func (t *BaseTile) GetResistance() float32 {
	return t.Resistance
}

func (t *BaseTile) GetPos() rl.Vector2 {
	return t.Pos
}

func (t *BaseTile) GetDimensions() (float32, float32) {
	return t.W, t.H
}


type AnimatedTile struct {
	BaseTile
	animData anim.AnimationData
	halfH float32
	halfW float32
}

func (a *AnimatedTile) Draw() {
	rl.DrawTexturePro(*a.texture, rl.NewRectangle(a.W * float32(a.animData.CurrFrame), 0, a.W, a.H), rl.NewRectangle(a.Pos.X, a.Pos.Y, a.W, a.H), rl.NewVector2(a.halfW, a.halfH), 0, rl.White)
}

func (a *AnimatedTile) Update(dt float32) {
	a.animData.Update(dt)
}
