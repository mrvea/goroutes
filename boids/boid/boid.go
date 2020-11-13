package boid

import (
	"time"

	"github.com/mrvea/goroutine/boids/env"
)

//Boid struct node to represent a boid
type Boid struct {
	position Vector2d
	velocity Vector2d
	id       int
}

//NewBoid makes a new boid with random position and velocity
func NewBoid(bid int) *Boid {
	b := Boid{
		position: NewRandVector(),
		velocity: NewRandVector(),
		id:       bid,
	}
	return &b
}

//Start initilizes a boid in an infinite loop for a goroutine
func (b *Boid) Start() {
	for {
		b.moveOne()
		time.Sleep(5 * time.Millisecond)
	}
}

func (b *Boid) moveOne() {
	b.position = b.position.Add(b.velocity)
	next := b.position.Add(b.velocity)
	if next.X >= env.ScreenWidth || next.X < 0 {
		b.bounceOffX()
	}
	if next.Y >= env.ScreenHeight || next.Y < 0 {
		b.bounceOffY()
	}
}

func (b *Boid) bounceOffX() {
	b.velocity.X *= -1
}
func (b *Boid) bounceOffY() {
	b.velocity.Y *= -1
}

//PositionXInt return position X as an int
func (b *Boid) PositionXInt() int {
	return int(b.position.X)
}

//PositionYInt return position Y as an int
func (b *Boid) PositionYInt() int {
	return int(b.position.Y)
}
