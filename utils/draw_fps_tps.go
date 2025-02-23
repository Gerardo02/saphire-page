package utils

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func DrawFPSTPS(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("FPS: %.3f", ebiten.ActualFPS()), 0, 0)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("TPS: %.3f", ebiten.ActualTPS()), 0, 20)
}
