package kissduck

import (
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"path"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	laserBeam        *ebiten.Image
	duckImage        *ebiten.Image
	handImage        *ebiten.Image
	humanImage       *ebiten.Image
	firingHumanImage *ebiten.Image
)

const (
	screenWidth  = 640
	screenHeight = 480
	maxHandCount = 100
	speedScale   = 5
	handsize     = 15
	ducksize     = 15
)

type Game struct {
	D *Duck
	C *Canon
	H *HandList
	S *Score
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.D.Draw(screen)
	g.C.Draw(screen)
	// g.C.DrawScope(screen)
	g.H.Draw(screen)
	g.S.Draw(screen)
}

func (g *Game) Update() error {
	g.D.Move()
	g.C.RotateCannonIfKeyPressed()
	g.H.Move()
	g.H.AddHandIfSpacebarPressed(g.C.AimTheta)
	g.H.CheckIfTouchingDuck(g.D.X, g.D.Y, g.D.R+handsize)
	g.S.Update(g.H.HitDuckCount, g.H.FireCount)
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func loadImages() {
	img, err := getImageFromFilePath(path.Join("assets", "laser.png"))
	if err != nil {
		log.Fatal("failed on getting image from path ", err.Error())
	}
	laserBeam = ebiten.NewImageFromImage(img)
	img, err = getImageFromFilePath(path.Join("assets", "duck.png"))
	if err != nil {
		log.Fatal("failed getting duck from path ", err.Error())
	}
	duckImage = ebiten.NewImageFromImage(img)
	img, err = getImageFromFilePath(path.Join("assets", "hand.png"))
	if err != nil {
		log.Fatal("failed getting hand from path ", err.Error())
	}
	handImage = ebiten.NewImageFromImage(img)
	img, err = getImageFromFilePath(path.Join("assets", "human0.png"))
	if err != nil {
		log.Fatal("failed getting human from path ", err.Error())
	}
	humanImage = ebiten.NewImageFromImage(img)
	img, err = getImageFromFilePath(path.Join("assets", "human1.png"))
	if err != nil {
		log.Fatal("failed getting human1 from path ", err.Error())
	}
	firingHumanImage = ebiten.NewImageFromImage(img)
}

func RunGame() {
	// load image that will be the beam
	loadImages()
	g := &Game{
		D: &Duck{
			Location: Location{
				X:  screenWidth / 2,
				Y:  screenHeight / 4,
				Vx: -speedScale,
				Vy: -speedScale,
				R:  ducksize,
			},
			C: color.RGBA{
				R: 100,
				G: 100,
				B: 255,
				A: 255},
		},
		C: &Canon{
			Location: Location{
				X: screenWidth / 2,
				Y: screenHeight,
				R: 20,
			},
			AimTheta: 180,
		},
		H: &HandList{
			Hands:        make([]*Hand, maxHandCount),
			HitDuckCount: 0,
			LastHandInd:  0,
			FireCount:    0,
		},
		S: &Score{
			Points:    0,
			HighScore: 0,
		},
	}
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Duck Prototype")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
