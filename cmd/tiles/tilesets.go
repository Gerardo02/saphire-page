package tiles

import (
	"encoding/json"
	"image"
	"log"
	"os"
	"strings"

	"github.com/gerardo02/saphire-page/types"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type ITileset interface {
	Img(id int) *ebiten.Image
	GetClass() string
}

type TilesetJSON struct {
	Class       string `json:"class"`
	Path        string `json:"image"`
	ImageWidth  int    `json:"imagewidth"`
	ImageHeight int    `json:"imageheight"`
	TileWidth   int    `json:"tilewidth"`
	TileHeight  int    `json:"tileheight"`
}

type Tileset struct {
	class    string
	img      *ebiten.Image
	gid      int
	imgSize  types.Vector2D[int]
	tileSize types.Vector2D[int]
}

func (t *Tileset) Img(id int) *ebiten.Image {
	// tileID - tilesetOffset = id relative to tileset
	id -= t.gid
	tileSizeX := t.tileSize.X
	tileSizeY := t.tileSize.Y

	srcX := id % (t.imgSize.X / t.tileSize.X)
	srcY := id / (t.imgSize.Y / t.tileSize.Y)

	srcX *= tileSizeX
	srcY *= tileSizeY

	return t.img.SubImage(
		image.Rect(srcX, srcY, srcX+tileSizeX, srcY+tileSizeY),
	).(*ebiten.Image)
}

func (t *Tileset) GetClass() string {
	return t.class
}

func InitTileset(filePath string, gid int) (ITileset, error) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	contents, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	tilesetJSON := TilesetJSON{}

	err = json.Unmarshal(contents, &tilesetJSON)
	if err != nil {
		return nil, err
	}

	newPath := strings.Replace(tilesetJSON.Path, "../../", dir+"/assets/", -1)

	img, _, err := ebitenutil.NewImageFromFile(newPath)
	if err != nil {
		return nil, err
	}

	tileset := Tileset{
		class:    tilesetJSON.Class,
		img:      img,
		gid:      gid,
		imgSize:  types.Vector2D[int]{X: tilesetJSON.ImageWidth, Y: tilesetJSON.ImageHeight},
		tileSize: types.Vector2D[int]{X: tilesetJSON.TileWidth, Y: tilesetJSON.TileHeight},
	}

	return &tileset, nil
}
