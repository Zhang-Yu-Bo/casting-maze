package core

import (
	"math"
	"maze/entity"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) move(width, height int) {
	clonePos := g.CurrentPos

	// move
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.CurrentPos.Z -= float64(MoveSpeed)
		g.UpdateRay = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.CurrentPos.X -= float64(MoveSpeed)
		g.UpdateRay = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.CurrentPos.Z += float64(MoveSpeed)
		g.UpdateRay = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.CurrentPos.X += float64(MoveSpeed)
		g.UpdateRay = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		g.CurrentDir -= float64(RotateSpeed)
		g.UpdateRay = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyE) {
		g.CurrentDir += float64(RotateSpeed)
		g.UpdateRay = true
	}

	// collision
	edgeWidth, edgeHeight := width/MazeWidth, height/MazeHeight
	roomX, roomY := PosToRoom(int(g.CurrentPos.X), int(g.CurrentPos.Z), edgeWidth, edgeHeight)
	if MazeInfo[roomX][roomY] == 1 {
		g.CurrentPos = clonePos
	}

	// clamp
	g.CurrentDir = Clamp(g.CurrentDir, 0, 360, true)

	// update ray
	if g.UpdateRay {
		g.CastingRays.Clear()

		step := FOV / float64(NumOfRays)
		beginAngle := g.CurrentDir - FOV/2
		wg := new(sync.WaitGroup)
		wg.Add(NumOfRays)
		for i := 0; i < NumOfRays; i++ {
			go func(index int) {
				angle := beginAngle + float64(index)*step
				ray := entity.NewRay(g.CurrentPos.X, g.CurrentPos.Z, angle, float64(g.ScreenWidth))
				cloest := math.Inf(1)
				for j := 0; j < len(g.Edges); j++ {
					point, in := entity.Intersection(ray, &g.Edges[j])
					if in {
						tmpDis := point.Distance(&g.CurrentPos)
						if tmpDis < cloest {
							ray.End = point
							ray.Color = g.EdgeColors[j]
						}
					}
				}
				g.CastingRays.Add(*ray)
				wg.Done()
			}(i)
		}
		wg.Wait()
	}
}

func Clamp(val, min, max float64, circle bool) float64 {
	if val < min {
		if circle {
			return max
		}
		return min
	}
	if val > max {
		if circle {
			return min
		}
		return max
	}
	return val
}

func PosToRoom(x, y, width, height int) (int, int) {
	return x / width, y / height
}
