package core

import (
	"image/color"
	"math"
	"math/rand"
	"maze/entity"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	ScreenWidth  int
	ScreenHeight int
	EdgeWidth    int
	EdgeHeight   int
	UpdateRay    bool
	RayLength    float64
	CastingRays  entity.LineSegs
	CurrentPos   entity.Vertex
	CurrentDir   float64
	Edges        []entity.LineSeg
	EdgeColors   []color.Color
}

func (g *Game) Update() error {
	g.move(g.ScreenWidth/2, g.ScreenHeight)
	g.UpdateRay = false
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.drawMinimap(screen)
	g.drawSky(screen)
	g.drawGround(screen)
	g.drawMaze(screen)
	g.drawFPS(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.ScreenWidth, g.ScreenHeight
}

func NewGame(width, height int) *Game {
	g := &Game{
		ScreenWidth:  width,
		ScreenHeight: height,
	}
	g.EdgeWidth, g.EdgeHeight = width/(MazeWidth*2), height/MazeHeight
	g.UpdateRay = true
	g.CurrentPos = *entity.NewVertex(100, 0, 100, 1)
	g.RayLength = float64(height) * math.Sqrt(2)
	g.CastingRays.Add(*entity.NewRay(g.CurrentPos.X, g.CurrentPos.Z, g.CurrentDir, g.RayLength))
	g.CurrentDir = 0

	g.Edges = make([]entity.LineSeg, 0, 4*MazeWidth*MazeHeight)
	g.EdgeColors = make([]color.Color, 0, 4*MazeWidth*MazeHeight)
	for i := 0; i < MazeWidth; i++ {
		for j := 0; j < MazeHeight; j++ {
			if MazeInfo[i][j] == 1 {
				x1, y1 := float64(i*g.EdgeWidth), float64(j*g.EdgeHeight)
				x2, y2 := float64((i+1)*g.EdgeWidth), float64((j+1)*g.EdgeHeight)
				c := color.RGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), 255}

				g.Edges = append(g.Edges, *entity.NewLineSeg(x1, y1, x2, y1))
				g.EdgeColors = append(g.EdgeColors, c)
				g.Edges = append(g.Edges, *entity.NewLineSeg(x2, y1, x2, y2))
				g.EdgeColors = append(g.EdgeColors, c)
				g.Edges = append(g.Edges, *entity.NewLineSeg(x2, y2, x1, y2))
				g.EdgeColors = append(g.EdgeColors, c)
				g.Edges = append(g.Edges, *entity.NewLineSeg(x1, y2, x1, y1))
				g.EdgeColors = append(g.EdgeColors, c)
			}
		}
	}

	return g
}
