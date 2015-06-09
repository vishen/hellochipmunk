package engine2

import (
	"github.com/go-gl/gl/v2.1/gl"
)

type Player struct {
	rotation  float32
	position  Position
	color     Color
	direction Direction

	velocity_x float32
	velocity_y float32

	gravity    float32
	move_speed float32
	radius     float32
}

func NewPlayer(pos Position, color Color, gravity, radius, move_speed float32) *Player {
	direction := NewDirection()

	// direction.MoveEast()

	return &Player{
		rotation:   0,
		position:   pos,
		direction:  direction,
		color:      color,
		gravity:    gravity,
		radius:     radius,
		move_speed: move_speed,
	}
}

func (self *Player) Update() {
	if self.direction.East {
		self.velocity_x += self.move_speed
	} else if self.direction.West {
		self.velocity_x -= self.move_speed
	}

	// Apply X velocity
	self.position.X += self.velocity_x

	//Then apply friction which naturally slows down object
	self.velocity_x *= 0.98

	//Apply y Velocity (Jump or gravity)
	self.position.Y += self.velocity_y

	//Then apply gravity to the velocity, this is done after to ensure initial jump surge is not affected untill next frame
	self.velocity_y += self.gravity

	//Updated constant to be 3.14 (pi) so it rotates exactly as the ball does, could be changed slightly
	self.rotation += 3.1415 * (-self.velocity_x)

	// Normalise position for now
	if self.position.Y < 100 {
		self.position.Y = 100
	}

}

func (self Player) Draw() {
	gl.PushMatrix()

	gl.Color4f(self.color.Red, self.color.Green, self.color.Blue, self.color.A)

	//Draw Player
	gl.PushMatrix()
	rot := self.rotation
	pos_x := self.position.X
	pos_y := self.position.Y

	gl.Translatef(pos_x, pos_y, 0.0)
	gl.Rotatef(float32(rot), 0, 0, 1)
	gl_drawCircle(float64(self.radius), 20)
	gl.PopMatrix()

	//Second Pop
	gl.PopMatrix()
}

func (self *Player) StopMoving() {
	self.direction.Reset()
}

func (self *Player) Jump() {
	self.velocity_y = 50
	// self.direction.MoveNorth()
}

func (self *Player) MoveRight() {
	self.direction.MoveEast()
}

func (self *Player) MoveLeft() {
	self.direction.MoveWest()
}
