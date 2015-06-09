package engine2

import ()

type Level struct {
	Player *Player
}

func NewLevel() *Level {
	return &Level{}

}

func (self *Level) AddPlayer(pos Position, color Color, gravity, radius, move_speed float32) {
	self.Player = NewPlayer(pos, color, gravity, radius, move_speed)
}

func (self *Level) Draw() {
	self.Player.Draw()
}

func (self *Level) Update() {
	self.Player.Update()
}
