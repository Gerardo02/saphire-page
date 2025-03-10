package camera

import (
	"math"

	"github.com/gerardo02/saphire-page/types"
)

type Camera struct {
	types.Vector2D[float64]
}

func InitCamera() *Camera {
	return &Camera{
		types.Vector2D[float64]{X: 0, Y: 0},
	}
}

func (c *Camera) FollowTarget(targetX, targetY, screenWidth, screenHeight float64) {
	c.X = -targetX + screenWidth/2.0
	c.Y = -targetY + screenHeight/2.0
}

func (c *Camera) Constraint(
	tilemapWidthPixels, tilemapHeightPixels, screenWidth, screenHeight float64,
) {
	c.X = math.Min(c.X, 0.0)
	c.Y = math.Min(c.Y, 0.0)

	c.X = math.Max(c.X, screenWidth-tilemapWidthPixels)
	c.Y = math.Max(c.Y, screenHeight-tilemapHeightPixels)
}
