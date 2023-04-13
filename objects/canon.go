package kissduck

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Canon struct {
	Location
	AimTheta int
}

func (c *Canon) RotateCannonIfKeyPressed() {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		c.AimTheta = c.AimTheta - speedScale
		if c.AimTheta >= 360 {
			c.AimTheta = 360
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		c.AimTheta = c.AimTheta + speedScale
		if c.AimTheta <= 0 {
			c.AimTheta = 0
		}
	}
}

func (c *Canon) Draw(screen *ebiten.Image) {
	im := humanImage
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		im = firingHumanImage
	}
	s := laserBeam.Bounds().Size()
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(s.X)/2, -float64(s.Y)/2)
	op.GeoM.Rotate(math.Pi)
	op.GeoM.Rotate(float64(c.AimTheta) * 2 * math.Pi / 360)
	op.GeoM.Scale(.1, .1)
	op.GeoM.Translate(screenWidth/2, screenHeight)

	screen.DrawImage(im, op)
}

// func (c *Canon) DrawCanon(screen *ebiten.Image) {
// 	vector.DrawFilledCircle(screen, float32(c.X), float32(c.Y), float32(c.R), color.RGBA{R: 255, G: 0, B: 10, A: 255}, false)
// }

// func (c *Canon) DrawScope(screen *ebiten.Image) {
// 	op := &ebiten.DrawImageOptions{}
// 	s := laserBeam.Bounds().Size()

// 	// Move the image's center to the screen's upper-left corner.
// 	// This is a preparation for rotating. When geometry matrices are applied,
// 	// the origin point is the upper-left corner.
// 	op.GeoM.Translate(-float64(s.X)/2, -float64(s.Y)/2)

// 	// Rotate the image. As a result, the anchor point of this rotate is
// 	// the center of the image.
// 	op.GeoM.Rotate(float64(c.AimTheta) * 2 * math.Pi / 360)
// 	op.GeoM.Scale(0.1, 0.1)
// 	// Move the image to the bottom center of screen.
// 	op.GeoM.Translate(screenWidth/2, screenHeight)

// 	screen.DrawImage(laserBeam, op)
// }
