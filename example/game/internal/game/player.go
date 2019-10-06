package game

import "github.com/silbinarywolf/swir/example/game/internal/game/input"

const (
	playerMoveSpeed = 4
)

var (
	player *Player = NewPlayer()
)

type Player struct {
	X, Y float64
}

func NewPlayer() *Player {
	self := new(Player)
	self.X = 32
	self.Y = 32
	return self
}

func (self *Player) Update() {
	if input.ButtonCheck(input.Left) {
		player.X -= playerMoveSpeed
	}
	if input.ButtonCheck(input.Right) {
		player.X += playerMoveSpeed
	}
	if input.ButtonCheck(input.Up) {
		player.Y -= playerMoveSpeed
	}
	if input.ButtonCheck(input.Down) {
		player.Y += playerMoveSpeed
	}
}
