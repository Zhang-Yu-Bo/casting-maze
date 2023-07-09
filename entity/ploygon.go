package entity

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Ploygon struct {
	Vertices []ebiten.Vertex
	Indices  []uint16
}

func NewPolygon(vertices []Vertex, colors []color.Color) *Ploygon {
	N := len(vertices)
	if N != len(colors) || N < 3 {
		panic("len(vertices) != len(colors) || len(vertices) < 3")
	}

	p := new(Ploygon)
	p.Vertices = make([]ebiten.Vertex, N)
	p.Indices = make([]uint16, 3*(N-2))

	for i := 0; i < N; i++ {
		if i < N-2 {
			p.Indices[3*i] = 0
			p.Indices[3*i+1] = uint16(i + 1)
			p.Indices[3*i+2] = uint16(i + 2)
		}

		r, g, b, a := colors[i].RGBA()
		r >>= 8
		g >>= 8
		b >>= 8
		a >>= 8
		p.Vertices[i] = ebiten.Vertex{
			DstX:   float32(vertices[i].X),
			DstY:   float32(vertices[i].Y),
			ColorR: float32(r) / 255,
			ColorG: float32(g) / 255,
			ColorB: float32(b) / 255,
			ColorA: float32(a) / 255,
		}
	}
	return p
}
