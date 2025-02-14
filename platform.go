package main

import (
	"image/color"

	"github.com/gen2brain/raylib-go/raylib"
)

type Platform struct {
	pos  rl.Vector3 `json: pos`
	size rl.Vector3 `json: size`
	col  rl.Color   `json: col`
}

func (p *Platform) draw() {
	rl.DrawCube(p.pos, p.size.X, p.size.Y, p.size.Z, p.col)
}

func newPlatform(pos, size rl.Vector3, col color.RGBA) (p Platform) {
	p.pos = pos
	p.size = size
	p.col = col

	return p
}
