package main

import (
	"math"

	"github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	pos rl.Vector3
	vel rl.Vector3
	cam Camera
}

func (p *Player) update() {
	if !world.editing {
		p.vel.X = 0
		p.vel.Z = 0

		if rl.IsKeyDown(rl.KeyW) {
			p.vel.Z += -float32(math.Cos(deg2rad(float64(p.cam.rot.Y)))) / 10
			p.vel.X += -float32(math.Sin(deg2rad(float64(p.cam.rot.Y)))) / 10
		} else if rl.IsKeyDown(rl.KeyS) {
			p.vel.Z += float32(math.Cos(deg2rad(float64(p.cam.rot.Y)))) / 10
			p.vel.X += float32(math.Sin(deg2rad(float64(p.cam.rot.Y)))) / 10
		}

		if rl.IsKeyDown(rl.KeyA) {
			p.vel.Z += -float32(math.Cos(deg2rad(float64(p.cam.rot.Y+90)))) / 10
			p.vel.X += -float32(math.Sin(deg2rad(float64(p.cam.rot.Y+90)))) / 10
		} else if rl.IsKeyDown(rl.KeyD) {
			p.vel.Z += float32(math.Cos(deg2rad(float64(p.cam.rot.Y+90)))) / 10
			p.vel.X += float32(math.Sin(deg2rad(float64(p.cam.rot.Y+90)))) / 10
		}

		p.pos = rl.Vector3Add(p.pos, p.vel)
		p.cam.update(p)
	}
}
func (p *Player) draw() {
	rl.DrawCube(p.pos, 1, 2, 1, rl.Yellow)
}

func newPlayer(pos rl.Vector3) (p Player) {
	p.pos = pos
	p.cam = newCamera(pos)

	return p
}
