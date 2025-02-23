package main

import (
	"log"

	"github.com/gerardo02/saphire-page/cmd/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("El juego indie")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	game := game.InitGame()

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
