package engine2

import (
	"fmt"
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
)

const (
	BALL_RADIUS = 25
	GRAVITY     = -2
	MOVESPEED   = 0.37
)

type World struct {
	Level *Level
}

func NewWorld() *World {
	level := NewLevel()
	level.AddPlayer(NewPosition(600, 600), NewColor(0.1, 0.4, 0.3, 0.9), GRAVITY, BALL_RADIUS, MOVESPEED)
	return &World{
		Level: level,
	}
}

func (self *World) Update() {
	self.Level.Update()
}

func (self *World) Draw() {
	gl.Clear(gl.COLOR_BUFFER_BIT)
	gl.Enable(gl.BLEND)
	gl.Enable(gl.POINT_SMOOTH)
	gl.Enable(gl.LINE_SMOOTH)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	gl.LoadIdentity()

	self.Level.Draw()
}

func (self *World) KeyCallback(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	player := self.Level.Player

	switch key {
	case glfw.KeyW:
		if action == glfw.Press {
			fmt.Printf("W Pressed!\n")
			player.Jump()
		}

	case glfw.KeyA:
		// fmt.Printf("A Pressed!\n")
		if action == glfw.Release {
			player.StopMoving()
		} else if action == glfw.Press {
			player.MoveLeft()
		}

	case glfw.KeyD:
		if action == glfw.Release {
			player.StopMoving()
		} else if action == glfw.Press {
			player.MoveRight()
		}
	}

}
