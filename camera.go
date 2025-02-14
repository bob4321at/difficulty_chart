package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Camera struct {
	pos rl.Vector3
	rot rl.Vector3
	cam rl.Camera
}

func (c *Camera) update(p *Player) {
	c.pos.X = float32(math.Sin(deg2rad(float64(c.rot.Y)))) * 10
	c.pos.Z = float32(math.Cos(deg2rad(float64(c.rot.Y)))) * 10

	c.pos.Y = float32(math.Sin(deg2rad(float64(c.rot.X)))) * 10

	c.pos = rl.Vector3Add(c.pos, p.pos)

	if rl.IsKeyDown(rl.KeyH) {
		c.rot.Y -= 1
	} else if rl.IsKeyDown(rl.KeyL) {
		c.rot.Y += 1
	}

	if rl.IsKeyDown(rl.KeyJ) {
		c.rot.X -= 1
	} else if rl.IsKeyDown(rl.KeyK) {
		c.rot.X += 1
	}

	c.cam.Position = c.pos
	c.cam.Target = p.pos
}

func newCamera(pos rl.Vector3) (c Camera) {
	c.pos = pos
	c.rot = rl.NewVector3(0, 0, 0)
	c.cam = rl.NewCamera3D(pos, rl.NewVector3(0, 0, 0), rl.NewVector3(0, 1, 0), 66, rl.CameraPerspective)

	return c
}
