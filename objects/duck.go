package kissduck

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	MAXSPEED = 10
)

type Duck struct {
	Location
	C color.Color
}

func (d *Duck) Move() {
	d.Location.Move(screenHeight/2, 0, screenWidth, 0)
}

func (d *Duck) Draw(screen *ebiten.Image) {
	// vector.DrawFilledCircle(screen, float32(d.X), float32(d.Y), float32(d.R), d.C, false)
	op := &ebiten.DrawImageOptions{}
	s := duckImage.Bounds().Size()
	op.GeoM.Translate(-float64(s.X)/2, -float64(s.Y)/2)
	scale := 3
	op.GeoM.Scale(float64(d.R*float32(scale)/float32(s.X)), float64(d.R*float32(scale)/float32(s.Y)))
	op.GeoM.Translate(float64(d.X), float64(d.Y))
	screen.DrawImage(duckImage, op)
}
