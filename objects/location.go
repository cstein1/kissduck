package kissduck

type Location struct {
	X, Y, Vx, Vy, R float32
}

func (l *Location) Move(maxHeight, minHeight, maxWidth, minWidth int) {
	l.calculateVelocity(maxHeight, minHeight, maxWidth, minWidth)
	l.X += l.Vx
	l.Y += l.Vy
}

func (l *Location) calculateVelocity(maxHeight, minHeight, maxWidth, minWidth int) {
	if l.X+l.R >= float32(maxWidth) {
		l.Vx = -MAXSPEED
	}
	if l.X-l.R <= float32(minWidth) {
		l.Vx = MAXSPEED
	}
	if l.Y+l.R >= float32(maxHeight) {
		l.Vy = -MAXSPEED
	}
	if l.Y-l.R <= float32(minHeight) {
		l.Vy = MAXSPEED
	}
}
