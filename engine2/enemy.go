package engine2

import (
	"github.com/go-gl/gl/v2.1/gl"
)

type Enemy struct {
	rotation  float32
	position  Position
	color     Color
	direction Direction

	velocity_x float32
	velocity_y float32

	move_speed float32
	radius     float32
}

func NewEnemy(pos Position, color Color, radius, move_speed float32) *Enemy {
	direction := NewDirection()

	// direction.MoveEast()

	return &Enemy{
		rotation:   0,
		position:   pos,
		direction:  direction,
		color:      color,
		radius:     radius,
		move_speed: move_speed,
	}
}

func (self Enemy) GetPosition() Position {
	return self.position
}

func (self *Enemy) Update() {
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
	// self.velocity_y += self.gravity

	//Updated constant to be 3.14 (pi) so it rotates exactly as the ball does, could be changed slightly
	self.rotation += 3.1415 * (-self.velocity_x)

	// Normalise position for now
	if self.position.Y < 100 {
		self.position.Y = 100
	}

}

func (self Enemy) Draw() {
	gl.PushMatrix()

	gl.Color4f(self.color.Red, self.color.Green, self.color.Blue, self.color.A)

	//Draw Enemy
	gl.PushMatrix()
	rot := self.rotation
	pos_x := self.position.X
	pos_y := self.position.Y

	gl.Translatef(pos_x, pos_y, 0.0)
	gl.Rotatef(float32(rot), 0, 0, 1)
	gl_drawCircle(float64(self.radius), 20)
	gl.PopMatrix()

	collision_points := self.GetCollision()

	gl.PushMatrix()
	gl.Begin(gl.LINES)
	gl.Color3f(.8, .8, .8)
	gl.Vertex3f(collision_points.Start.X, collision_points.Start.Y, 0)
	gl.Vertex3f(collision_points.Start.X, collision_points.End.Y, 0)

	gl.Vertex3f(collision_points.Start.X, collision_points.End.Y, 0)
	gl.Vertex3f(collision_points.End.X, collision_points.End.Y, 0)

	gl.Vertex3f(collision_points.End.X, collision_points.End.Y, 0)
	gl.Vertex3f(collision_points.End.X, collision_points.Start.Y, 0)

	gl.Vertex3f(collision_points.Start.X, collision_points.Start.Y, 0)
	gl.End()

	//Second Pop
	gl.PopMatrix()
}

func (self Enemy) GetCollision() Collision {
	length := self.radius
	return Collision{
		Start: self.position.Subtract(length, length),
		End:   self.position.Add(length, length),
	}
}

func (self *Enemy) StopMoving() {
	self.direction.Reset()
}

func (self *Enemy) MoveUp() {
	self.direction.MoveNorth()
}

func (self *Enemy) MoveDown() {
	self.direction.MoveSouth()
}

func (self *Enemy) MoveRight() {
	self.direction.MoveEast()
}

func (self *Enemy) MoveLeft() {
	self.direction.MoveWest()
}
