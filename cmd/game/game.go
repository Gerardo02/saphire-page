package game

import (
	"image/color"
	"log"
	"os"

	"github.com/gerardo02/saphire-page/cmd/camera"
	"github.com/gerardo02/saphire-page/cmd/player"
	"github.com/gerardo02/saphire-page/cmd/tiles"
	"github.com/gerardo02/saphire-page/defs"
	"github.com/gerardo02/saphire-page/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Player   *player.Player
	Tilemap  *tiles.Tilemap
	Tilesets map[string]tiles.ITileset
}

func (g *Game) Update() error {
	err := g.Player.Update()
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{120, 180, 255, 255})

	g.Tilemap.Draw(screen, g.Tilesets, g.Player.Camera)

	g.Player.Draw(screen)

	utils.DrawFPSTPS(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	// return ebiten.WindowSize()
	return defs.SCREEN_WIDTH, defs.SCREEN_HEIGHT
}

func InitGame() *Game {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	tilemap := tiles.InitTilemap(dir + "/assets/maps/tilemaps/level1.json")

	return &Game{
		Player:   player.InitPlayer(camera.InitCamera()),
		Tilemap:  tilemap,
		Tilesets: tilemap.GenerateTilesets(),
	}
}
