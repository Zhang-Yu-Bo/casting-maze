package main

import (
	"log"
	"maze/core"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	height := 480
	width := height * 2

	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle("Hello, World!")

	game := core.NewGame(width, height)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
