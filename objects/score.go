package kissduck

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Score struct {
	Points    int
	HighScore int
	Accuracy  int
}

func (s *Score) Draw(screen *ebiten.Image) {
	scoreMessage := fmt.Sprintf("Points: %d\nHigh Score: %d\nAccuracy: %d%%", s.Points, s.HighScore, s.Accuracy)
	ebitenutil.DebugPrintAt(screen, scoreMessage, screenWidth*3/4, screenHeight*3/4) // screenWidth*3/4, screenWidth*3/4)
}

func (s *Score) Update(scoreNum, numShots int) {
	s.Points = scoreNum
	if scoreNum > s.HighScore {
		s.HighScore = scoreNum
	}
	if numShots != 0 {
		s.Accuracy = int(100 * float32(scoreNum) / float32(numShots))
	} else {
		s.Accuracy = 100
	}
}
