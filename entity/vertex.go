package entity

import "math"

type Vertex struct {
	X float64
	Y float64
	Z float64
	W float64
}

func NewVertex(x, y, z, w float64) *Vertex {
	return &Vertex{
		X: x,
		Y: y,
		Z: z,
		W: w,
	}
}

func (v *Vertex) Distance(v2 *Vertex) float64 {
	dx := v.X - v2.X
	dy := v.Y - v2.Y
	dz := v.Z - v2.Z
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}
