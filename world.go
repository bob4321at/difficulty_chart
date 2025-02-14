package main

import (
	"fmt"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Worldo struct {
	platforms         []Platform
	selected_platform int
	editing           bool
}

func (w *Worldo) addPlatform() {
	w.platforms = append(w.platforms, newPlatform(rl.NewVector3(0, 0, 0), rl.NewVector3(1, 1, 1), rl.Black))
}

func (w *Worldo) update() {
	if w.selected_platform > len(w.platforms) {
		w.selected_platform = len(w.platforms) - 1
	}

	if rl.IsKeyPressed(rl.KeyTab) {
		w.editing = !w.editing
	}
	if w.editing {
		if rl.IsKeyPressed(rl.KeyR) {
			w.addPlatform()
			w.selected_platform += 1
		}

		if rl.IsKeyPressed(rl.KeyF) {
			if w.selected_platform+1 < len(w.platforms) {
				w.selected_platform += 1
			} else {
				w.selected_platform = 0
			}
		}

		p := &w.platforms[w.selected_platform]
		if rl.IsKeyDown(rl.KeyW) {
			p.pos.Z -= 0.1
		} else if rl.IsKeyDown(rl.KeyS) {
			p.pos.Z += 0.1
		}
		if rl.IsKeyDown(rl.KeyA) {
			p.pos.X -= 0.1
		} else if rl.IsKeyDown(rl.KeyD) {
			p.pos.X += 0.1
		}
		if rl.IsKeyDown(rl.KeyQ) {
			p.pos.Y -= 0.1
		} else if rl.IsKeyDown(rl.KeyE) {
			p.pos.Y += 0.1
		}

		if rl.IsKeyDown(rl.KeyI) {
			p.size.X -= 0.1
		} else if rl.IsKeyDown(rl.KeyO) {
			p.size.X += 0.1
		}
		if rl.IsKeyDown(rl.KeyU) {
			p.size.Z -= 0.1
		} else if rl.IsKeyDown(rl.KeyP) {
			p.size.Z += 0.1
		}
		if rl.IsKeyDown(rl.KeyMinus) {
			p.size.Y -= 0.1
		} else if rl.IsKeyDown(rl.KeyEqual) {
			p.size.Y += 0.1
		}

		if rl.IsKeyPressed(rl.KeyB) {
			file_info := ""
			file_info = "package main\n import \"github.com/gen2brain/raylib-go/raylib\"\n var level = []Platform{"
			for platform_index := 0; platform_index < len(w.platforms); platform_index++ {
				p := &w.platforms[platform_index]
				file_info = file_info + "newPlatform(rl.NewVector3(" + fmt.Sprintf("%v", p.pos.X) + ", " + fmt.Sprintf("%v", p.pos.Y) + ", " + fmt.Sprintf("%v", p.pos.Z) + "), rl.NewVector3(" + fmt.Sprintf("%v", p.size.X) + ", " + fmt.Sprintf("%v", p.size.Y) + ", " + fmt.Sprintf("%v", p.size.Z) + "), rl.Black),"
			}
			file_info = file_info + "}"

			os.Remove("./level.go")

			file, err := os.Create("level.go")
			if err != nil {
				panic(err)
			}
			file.Write([]byte(file_info))
		}
	}
}

func (w *Worldo) draw() {
	for platform_index := 0; platform_index < len(w.platforms); platform_index++ {
		p := &w.platforms[platform_index]
		p.draw()
	}
}

func newWorld() (w Worldo) {
	w.editing = false
	w.selected_platform = 0

	if level != nil {
		w.platforms = nil
		w.platforms = level
	}

	return w
}
