package player

import (
	"image"
	"log"
	"os"

	"github.com/gerardo02/saphire-page/cmd/camera"
	"github.com/gerardo02/saphire-page/cmd/sprites"
	"github.com/gerardo02/saphire-page/defs"
	"github.com/gerardo02/saphire-page/types"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Player struct {
	*sprites.Sprite

	Camera   *camera.Camera
	position types.Vector2D[float64]
	speed    types.Vector2D[float64]
	size     types.Vector2D[float64]
}

func (p *Player) Update() error {
	p.speed.X = 0
	p.speed.Y = 0

	p.updateMovement()

	p.position.X += p.speed.X
	p.position.Y += p.speed.Y

	p.Camera.FollowTarget(p.position.X+16, p.position.Y+32, defs.SCREEN_WIDTH, defs.SCREEN_HEIGHT)

	return nil
}

func (p *Player) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}

	opts.GeoM.Translate(p.position.X, p.position.Y)
	opts.GeoM.Translate(p.Camera.X, p.Camera.Y)

	screen.DrawImage(
		p.Sprite.Img.SubImage(
			image.Rect(0, 0, 16*2, 16*4),
		).(*ebiten.Image),
		opts,
	)

	opts.GeoM.Reset()
}

func InitPlayer(c *camera.Camera) *Player {
	// Load player sprite
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	sprite, _, err := ebitenutil.NewImageFromFile(dir + "/assets/sprites/player/player-cat.png")
	if err != nil {
		log.Fatal(err)
	}

	return &Player{
		position: types.Vector2D[float64]{X: 150, Y: 38 * 16},
		speed:    types.Vector2D[float64]{X: 0, Y: 0},
		size:     types.Vector2D[float64]{X: 16.0 * 2.0, Y: 16.0 * 4.0},

		Camera: c,
		Sprite: &sprites.Sprite{
			Img: sprite,
		},
	}
}

func (p *Player) updateMovement() {
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		p.speed.X = -2
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		p.speed.X = 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		p.speed.Y = -2
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		p.speed.Y = 2
	}
}
