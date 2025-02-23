package tiles

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/gerardo02/saphire-page/cmd/camera"
	"github.com/hajimehoshi/ebiten/v2"
)

type tilesetMeta struct {
	Gid  int    `json:"firstgid"`
	Path string `json:"source"`
}

type Layer struct {
	Data   []int  `json:"data"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Name   string `json:"name"`
}

type Tilemap struct {
	Layers       []Layer       `json:"layers"`
	TilesetsMeta []tilesetMeta `json:"tilesets"`
}

func (t *Tilemap) GenerateTilesets() map[string]ITileset {
	tilesets := make(map[string]ITileset)

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	for _, tilesetData := range t.TilesetsMeta {
		newPath := strings.Replace(tilesetData.Path, "../", dir+"/assets/maps/", -1)

		tileset, err := InitTileset(newPath, tilesetData.Gid)
		if err != nil {
			log.Fatal(err)
		}

		tilesets[tileset.GetClass()] = tileset
	}

	return tilesets
}

func (t *Tilemap) Draw(screen *ebiten.Image, tilesets map[string]ITileset, cam *camera.Camera) {
	opts := ebiten.DrawImageOptions{}
	for _, layer := range t.Layers {
		for j, tileID := range layer.Data {
			if tileID == 0 {
				continue
			}

			x := j % layer.Width
			y := j / layer.Width

			x *= 16
			y *= 16

			img := tilesets[layer.Name].Img(tileID)
			opts.GeoM.Translate(float64(x), float64(y))
			opts.GeoM.Translate(0.0, -(float64(img.Bounds().Dy()) + 16.0))
			opts.GeoM.Translate(cam.X, cam.Y)

			screen.DrawImage(img, &opts)

			opts.GeoM.Reset()
		}
	}
}

func InitTilemap(filePath string) *Tilemap {
	jsonFile, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	tilemap := Tilemap{}

	err = json.Unmarshal(jsonFile, &tilemap)
	if err != nil {
		log.Fatal(err)
	}

	return &tilemap
}
