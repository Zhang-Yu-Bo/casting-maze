package core

import (
	"fmt"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func (g *Game) drawMaze(screen *ebiten.Image) {
	sWidth := g.ScreenWidth / 2
	sHeight := g.ScreenHeight
	startX := float32(sWidth)
	beginAngle := g.CurrentDir - FOV/2

	numOfRays := g.CastingRays.Len()
	stepWidth := float32(sWidth) / float32(numOfRays)
	for i := 0; i < numOfRays; i++ {
		ray, ok := g.CastingRays.Get(i)
		if !ok {
			break
		}

		pTh := float32(math.Abs(ray.Angle-beginAngle)/FOV) * float32(numOfRays)
		distance := ray.Length() * math.Cos(math.Abs(ray.Angle-g.CurrentDir)*math.Pi/180)
		depth := Clamp(distance, 0, g.RayLength, false)
		wHeight := Weight(float64(sHeight), depth, g.RayLength)

		postX := startX + pTh*stepWidth
		postY := (float64(sHeight) - wHeight) / 2
		vector.DrawFilledRect(screen, postX, float32(postY), stepWidth, float32(wHeight), WeightColor(ray.Color, depth, g.RayLength), true)
	}
}

func Weight(val, p, max float64) float64 {
	return val * (1 - p/max)
}

func WeightColor(c color.Color, p, max float64) color.Color {
	r, g, b, _ := c.RGBA()
	r, g, b = r>>8, g>>8, b>>8
	weight := 1 - (p / max)
	return color.RGBA{
		uint8(float64(r) * weight * weight),
		uint8(float64(g) * weight * weight),
		uint8(float64(b) * weight * weight),
		255,
	}
}

func (g *Game) drawMinimap(screen *ebiten.Image) {
	sWidth := g.ScreenWidth / 2
	sHeight := g.ScreenHeight

	for i := 0; i < MazeWidth; i++ {
		for j := 0; j < MazeHeight; j++ {
			if MazeInfo[i][j] == 1 {
				vector.DrawFilledRect(screen, float32(i*g.EdgeWidth), float32(j*g.EdgeHeight), float32(g.EdgeWidth), float32(g.EdgeHeight), color.White, false)
			}
		}
	}

	lineColor := color.RGBA{128, 128, 128, 255}
	lineWidth := float32(5)
	for i := 0; i <= MazeWidth; i++ {
		vector.StrokeLine(screen, float32(i*g.EdgeWidth), 0, float32(i*g.EdgeWidth), float32(sHeight), lineWidth, lineColor, false)
	}
	for j := 0; j <= MazeHeight; j++ {
		vector.StrokeLine(screen, 0, float32(j*g.EdgeHeight), float32(sWidth), float32(j*g.EdgeHeight), lineWidth, lineColor, false)
	}

	// draw person
	personRange := float32(10)
	personColor := color.RGBA{100, 225, 190, 255}
	vector.DrawFilledCircle(screen, float32(g.CurrentPos.X), float32(g.CurrentPos.Z), personRange, personColor, true)

	// draw direction
	dirWidth := float32(2)
	dirColor := color.RGBA{247, 246, 136, 255}
	numOfRays := g.CastingRays.Len()
	for i := 0; i < numOfRays; i++ {
		ray, ok := g.CastingRays.Get(i)
		if !ok {
			break
		}

		vector.StrokeLine(screen, float32(ray.Start.X), float32(ray.Start.Y), float32(ray.End.X), float32(ray.End.Y), dirWidth, dirColor, true)
		vector.StrokeCircle(screen, float32(ray.End.X), float32(ray.End.Y), 5, dirWidth, dirColor, true)
	}
}

func (g *Game) drawFPS(screen *ebiten.Image) {
	fps := ebiten.ActualFPS()
	fpsStr := fmt.Sprintf("FPS: %0.2f", fps)
	ebitenutil.DebugPrint(screen, fpsStr)
}

func (g *Game) drawSky(screen *ebiten.Image) {
	sWidth := g.ScreenWidth / 2
	sHeight := g.ScreenHeight / 2

	// draw sky
	skyColor := color.RGBA{135, 206, 235, 255}
	vector.DrawFilledRect(screen, float32(sWidth), 0, float32(sWidth), float32(sHeight), skyColor, false)
}

func (g *Game) drawGround(screen *ebiten.Image) {
	sWidth := g.ScreenWidth / 2
	sHeight := g.ScreenHeight / 2

	// draw ground
	groundColor := color.RGBA{34, 139, 34, 255}
	vector.DrawFilledRect(screen, float32(sWidth), float32(sHeight), float32(sWidth), float32(sHeight), groundColor, false)
}
