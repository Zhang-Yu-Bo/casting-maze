package entity

import (
	"image/color"
	"math"
	"sync"
)

type LineSeg struct {
	Start Vertex
	End   Vertex
	Color color.Color
	Angle float64
}

type LineSegs struct {
	Lines []LineSeg
	sync.RWMutex
}

func (l *LineSegs) Add(line LineSeg) {
	l.Lock()
	defer l.Unlock()
	l.Lines = append(l.Lines, line)
}

func (l *LineSegs) Get(index int) (LineSeg, bool) {
	l.RLock()
	defer l.RUnlock()

	if index < 0 || index >= len(l.Lines) {
		return LineSeg{}, false
	}
	return l.Lines[index], true
}

func (l *LineSegs) Len() int {
	l.RLock()
	defer l.RUnlock()
	return len(l.Lines)
}

func (l *LineSegs) Clear() {
	l.Lock()
	defer l.Unlock()
	l.Lines = []LineSeg{}
}

func (l *LineSeg) Length() float64 {
	return l.Start.Distance(&l.End)
}

func NewRay(x, y, dir, length float64) *LineSeg {
	toRadian := dir * math.Pi / 180
	return &LineSeg{
		Start: Vertex{
			X: x,
			Y: y,
			Z: 0,
			W: 1,
		},
		End: Vertex{
			X: x + length*math.Cos(toRadian),
			Y: y + length*math.Sin(toRadian),
			Z: 0,
			W: 1,
		},
		Angle: dir,
	}
}

func NewLineSeg(x1, y1, x2, y2 float64) *LineSeg {
	return &LineSeg{
		Start: Vertex{
			X: x1,
			Y: y1,
			Z: 0,
			W: 1,
		},
		End: Vertex{
			X: x2,
			Y: y2,
			Z: 0,
			W: 1,
		},
	}
}

func Intersection(line1, line2 *LineSeg) (Vertex, bool) {
	// https://en.wikipedia.org/wiki/Line%E2%80%93line_intersection#Given_two_points_on_each_line
	denom := (line1.Start.X-line1.End.X)*(line2.Start.Y-line2.End.Y) - (line1.Start.Y-line1.End.Y)*(line2.Start.X-line2.End.X)
	tNum := (line1.Start.X-line2.Start.X)*(line2.Start.Y-line2.End.Y) - (line1.Start.Y-line2.Start.Y)*(line2.Start.X-line2.End.X)
	uNum := -((line1.Start.X-line1.End.X)*(line1.Start.Y-line2.Start.Y) - (line1.Start.Y-line1.End.Y)*(line1.Start.X-line2.Start.X))

	if denom == 0 {
		return Vertex{}, false
	}

	t := tNum / denom
	if t > 1 || t < 0 {
		return Vertex{}, false
	}

	u := uNum / denom
	if u > 1 || u < 0 {
		return Vertex{}, false
	}

	x := line1.Start.X + t*(line1.End.X-line1.Start.X)
	y := line1.Start.Y + t*(line1.End.Y-line1.Start.Y)
	return Vertex{X: x, Y: y, W: 1}, true
}
