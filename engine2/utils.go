package engine2

import (
	"fmt"
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

func (self Position) Add(x, y float32) Position {
	return Position{X: self.X + x, Y: self.Y + y}
}

func (self Position) Subtract(x, y float32) Position {
	return Position{X: self.X - x, Y: self.Y - y}
}

func (self Position) GreaterThan(other Position) bool {
	return self.X >= other.X && self.Y >= other.Y
}

func (self Position) LessThan(other Position) bool {
	return self.X <= other.X && self.Y <= other.Y
}

func (self Position) Equals(other Position) bool {
	return self.X == other.X && self.Y == other.Y
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

type Collision struct {
	Start Position
	End   Position
}

func (self *Collision) CollidesWith(other Collision) bool {
	var leeway float64 = 5

	self_start := self.Start
	self_end := self.End
	self_points := [4]Position{
		Position{X: self_start.X, Y: self_start.Y},
		Position{X: self_start.X, Y: self_end.Y},
		Position{X: self_end.X, Y: self_start.Y},
		Position{X: self_end.X, Y: self_end.Y},
	}

	other_start := other.Start
	other_end := other.End
	other_points := [4]Position{
		Position{X: other_start.X, Y: other_start.Y},
		Position{X: other_start.X, Y: other_end.Y},
		Position{X: other_end.X, Y: other_start.Y},
		Position{X: other_end.X, Y: other_end.Y},
	}

	for _, p1 := range self_points {
		for _, p2 := range other_points {
			if math.Abs(float64(p1.X-p2.X)) <= leeway && math.Abs(float64(p1.Y-p2.Y)) <= leeway {
				fmt.Println(p1, p2)
				fmt.Println(math.Abs(float64(p1.X-p2.X)), math.Abs(float64(p1.Y-p2.Y)))
				return true
			}
		}
	}

	return false
}

// func (self *Collision) CollidesWith(other Collision) bool {

// 	self_start := self.Start
// 	self_end := self.End
// 	// Checks for a Rectangular hitbox
// 	self_start_line := [2]Position{
// 		Position{X: self_start.X, Y: self_start.Y},
// 		Position{X: self_start.X, Y: self_end.Y},
// 	}

// 	self_end_line := [2]Position{
// 		Position{X: self_end.X, Y: self_start.Y},
// 		Position{X: self_end.X, Y: self_end.Y},
// 	}

// 	other_start := other.Start
// 	other_end := other.End
// 	// Checks for a Rectangular hitbox
// 	other_start_points := [2]Position{
// 		Position{X: other_start.X, Y: other_start.Y},
// 		Position{X: other_start.X, Y: other_end.Y},
// 	}

// 	other_end_points := [2]Position{
// 		Position{X: other_end.X, Y: other_start.Y},
// 		Position{X: other_end.X, Y: other_end.Y},
// 	}

// 	fmt.Println("#############################################")
// 	fmt.Println("# Collision Debugging #")
// 	fmt.Println("#############################################")

// 	fmt.Println(self, other)
// 	fmt.Println("#############################################")

// 	fmt.Println(self_start_line)
// 	fmt.Println(other_end_points)
// 	fmt.Println("#############################################")

// 	fmt.Println(self_end_line)
// 	fmt.Println(other_start_points)
// 	fmt.Println("#############################################")
// 	fmt.Println()
// 	fmt.Println()

// 	// if self.Start.GreaterThan(other.Start) && self.Start.LessThan(other.end) {
// 	// 	return true
// 	// } else if

// 	// for _, pos := range other_start_points {
// 	// 	if pos.GreaterThan(self_end_line[0]) && pos.LessThan(self_end_line[1]) {
// 	// 		fmt.Println("ONE")
// 	// 		fmt.Println(self_end_line)
// 	// 		fmt.Println(pos)
// 	// 		return true
// 	// 	}
// 	// }

// 	// for _, pos := range other_end_points {
// 	// 	if pos.GreaterThan(self_start_line[0]) && pos.LessThan(self_start_line[1]) {
// 	// 		fmt.Println("TWO")
// 	// 		fmt.Println(self_start_line)
// 	// 		fmt.Println(other_end_points)
// 	// 		return true
// 	// 	}
// 	// }

// 	return false

// }

// drawCircle draws a circle for the specified radius, rotation angle, and the specified number of sides
func gl_drawCircle(radius float64, sides int) {
	gl.Begin(gl.TRIANGLE_FAN)
	for a := 0.0; a < 2*math.Pi; a += (2 * math.Pi / float64(70)) {
		gl.Vertex2d(math.Sin(a)*radius, math.Cos(a)*radius)
	}
	gl.Vertex3f(0, 0, 0)
	gl.End()

}
