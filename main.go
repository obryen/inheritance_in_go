package main

import "fmt"

type Position struct {
	x float64
	y float64
}

type SpecialPosition struct {
	Position
}

func (p *Position) Move(x float64, y float64) {
	p.x += x
	p.y += y
}

func (p *SpecialPosition) MoveSpecial(x float64, y float64) {
	p.x *= x
	p.y *= y
}

func (p *Position) Teleport(x float64, y float64) {
	p.x = x
	p.y = y
}

type Player struct {
	*Position
}

func NewPlayer() *Player {
	return &Player{
		Position: &Position{},
	}
}

type Enemy struct {
	*SpecialPosition
	fire int64
}

func NewEnemy() *Enemy {
	return &Enemy{
		SpecialPosition: &SpecialPosition{},
	}
}

func main() {
	player := NewPlayer()
	player.Move(1.1, 3)
	fmt.Println("player position is", player.Position)
	player.Teleport(6, 7)
	fmt.Println("player has teleported to", player.Position)

	enemy := NewEnemy()
	enemy.Move(5, 6)
	enemy.fire = 000
	enemy.MoveSpecial(5, 5)
	fmt.Println("enemy position is", *&enemy.Position)
}
