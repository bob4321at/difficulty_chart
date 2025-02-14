package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

func setupWindow() {
	rl.SetConfigFlags(rl.FlagFullscreenMode)
	rl.SetConfigFlags(rl.FlagMsaa4xHint)
	rl.InitWindow(1920, 1080, "get gud nerd")

	rl.SetTargetFPS(60)
}

func update() {
	player.update()
	world.update()
}

var player Player
var world Worldo

func draw() {
	rl.BeginDrawing()

	rl.ClearBackground(rl.Beige)

	rl.BeginMode3D(player.cam.cam)

	player.draw()
	world.draw()

	rl.EndMode3D()

	rl.DrawFPS(10, 10)

	rl.EndDrawing()
}

func main() {
	setupWindow()

	player = newPlayer(rl.NewVector3(-2, 0, -2))

	world = newWorld()

	world.addPlatform()

	for !rl.WindowShouldClose() {
		draw()
		update()
	}
}
