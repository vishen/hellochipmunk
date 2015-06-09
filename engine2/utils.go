package engine2

import (
	"github.com/go-gl/gl/v2.1/gl"
	"math"
)

type Color struct {
	Red   float32
	Green float32
	Blue  float32
	A     float32
}

func NewColor(r, g, b, a float32) Color {
	return Color{Red: r, Green: g, Blue: b, A: a}
}

type Position struct {
	X float32
	Y float32
}

func NewPosition(x, y float32) Position {
	return Position{X: x, Y: y}
}

type Direction struct {
	North bool
	South bool
	East  bool
	West  bool
}

func NewDirection() Direction {
	return Direction{}
}

func (self *Direction) Reset() {
	self.North = false
	self.South = false
	self.East = false
	self.West = false
}

func (self *Direction) MoveNorth() {
	self.North = true
	self.South = false
}

func (self *Direction) MoveSouth() {
	self.North = false
	self.South = true
}

func (self *Direction) MoveEast() {
	self.East = true
	self.West = false
}

func (self *Direction) MoveWest() {
	self.East = false
	self.West = true
}

// drawCircle draws a circle for the specified radius, rotation angle, and the specified number of sides
func gl_drawCircle(radius float64, sides int) {
	gl.Begin(gl.TRIANGLE_FAN)
	for a := 0.0; a < 2*math.Pi; a += (2 * math.Pi / float64(70)) {
		gl.Vertex2d(math.Sin(a)*radius, math.Cos(a)*radius)
	}
	gl.Vertex3f(0, 0, 0)
	gl.End()

}
