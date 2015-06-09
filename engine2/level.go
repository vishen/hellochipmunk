package engine2

import (
	"fmt"
	"github.com/go-gl/gl/v2.1/gl"
)

type Level struct {
	enemy_spawn_interval int
	enemy_spawn_tick     int

	Player  *Player
	Enemies []*Enemy
}

func NewLevel() *Level {
	return &Level{enemy_spawn_interval: 200}

}

func (self *Level) AddPlayer(pos Position, color Color, gravity, radius, move_speed float32) {
	self.Player = NewPlayer(pos, color, gravity, radius, move_speed)
}

func (self *Level) SpawnEnemy() {
	pos := Position{100, 100}
	color := Color{0.5, 0.5, 0.5, 1}
	var radius float32 = 25
	var move_speed float32 = 0.1
	self.Enemies = append(self.Enemies, NewEnemy(pos, color, radius, move_speed))
}

func (self *Level) Draw() {

	gl.Begin(gl.LINES)
	gl.Color3f(.2, .5, .2)
	gl.Vertex3f(float32(0), float32(100), 0)
	gl.Vertex3f(float32(1200), float32(100), 0)
	gl.End()

	self.Player.Draw()
	for _, enemy := range self.Enemies {
		enemy.Draw()
	}
}

func (self *Level) Update() {
	self.Player.Update()

	player_collision := self.Player.GetCollision()
	player_position := self.Player.GetPosition()

	var enemy_position Position

	for i, enemy := range self.Enemies {
		enemy_position = enemy.GetPosition()
		if enemy_position.X > player_position.X {
			enemy.MoveLeft()
		} else if enemy_position.X < player_position.X {
			enemy.MoveRight()
		}

		if enemy_position.Y > player_position.Y {
			enemy.MoveDown()
		} else if enemy_position.Y < player_position.Y {
			enemy.MoveUp()
		}

		enemy.Update()
		// Check Collision
		if player_collision.CollidesWith(enemy.GetCollision()) {
			fmt.Println("COllides")
			// Not sure if the best way to do this ??
			self.Enemies = append(self.Enemies[:i], self.Enemies[i+1:]...)
			// panic()
		}

	}

	if self.enemy_spawn_tick >= self.enemy_spawn_interval {
		self.SpawnEnemy()
		self.enemy_spawn_tick = 0
	}

	self.enemy_spawn_tick += 1

}
