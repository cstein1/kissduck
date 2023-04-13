package kissduck

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type HandList struct {
	Hands                []*Hand
	LastHandInd          int
	HitDuckCount         int
	SpaceBarAcknowledged bool
	FireCount            int
}

type Hand struct {
	Location
	IsTouchingDuck bool
	ExitedArena    bool
	C              color.Color
}

func (h *Hand) Move() {
	h.Location.Move(screenHeight, -100, screenWidth, 0)
	h.ExitedArena = h.Location.Y < 0
}

func (hs *HandList) Move() {
	for _, hand := range hs.Hands {
		if hand != nil {
			hand.Move()
		}
	}
}

func (h *Hand) Draw(screen *ebiten.Image) {
	// vector.DrawFilledCircle(screen, float32(h.X), float32(h.Y), float32(h.R), h.C, false)
	op := &ebiten.DrawImageOptions{}
	s := handImage.Bounds().Size()
	op.GeoM.Translate(-float64(s.X)/2, -float64(s.Y)/2)
	scale := 3
	op.GeoM.Scale(float64(h.R*float32(scale)/float32(s.X)), float64(h.R*float32(scale)/float32(s.Y)))
	op.GeoM.Translate(float64(h.X), float64(h.Y))
	screen.DrawImage(handImage, op)
}

func (h *Hand) CheckIfTouchingDuck(duckX, duckY, hitbox float32) bool {
	return (h.X > float32(duckX-hitbox) && h.X < float32(duckX+hitbox)) && (h.Y > float32(duckY-hitbox) && h.Y < float32(duckY+hitbox))
}

func (hs *HandList) CheckIfTouchingDuck(duckX, duckY, hitbox float32) {
	for _, hand := range hs.Hands {
		if hand == nil {
			continue
		}
		hand.IsTouchingDuck = hand.CheckIfTouchingDuck(duckX, duckY, hitbox)
		if hand.IsTouchingDuck {
			hs.HitDuckCount++
		}
	}
}

func (hs *HandList) Draw(screen *ebiten.Image) {
	for ind, hand := range hs.Hands {
		if hand == nil {
			continue
		}
		if !hand.IsTouchingDuck && !hand.ExitedArena {
			hand.Draw(screen)
		} else {
			hs.Hands[ind] = nil
		}
	}
}

func (hs *HandList) AddHandIfSpacebarPressed(aimTheta int) {
	if ebiten.IsKeyPressed(ebiten.KeySpace) && !hs.SpaceBarAcknowledged {
		hs.SpaceBarAcknowledged = true
		hs.FireCount++
		hs.Hands[hs.LastHandInd%maxHandCount] = &Hand{
			Location: Location{
				X:  screenWidth / 2,
				Y:  screenHeight,
				Vx: float32(aimTheta-180) / speedScale,
				Vy: -speedScale,
				R:  handsize,
			},
			C: color.RGBA{
				R: 100,
				G: 255,
				B: 100,
				A: 255,
			},
			IsTouchingDuck: false,
			ExitedArena:    false,
		}
		hs.LastHandInd++
	}
	hs.SpaceBarAcknowledged = ebiten.IsKeyPressed(ebiten.KeySpace)
}
