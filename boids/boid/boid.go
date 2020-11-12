package boid

import (
	"time"

	"github.com/mrvea/goroutine/boids/env"
)

type Boid struct {
	position Vector2d
	velocity Vector2d
	id       int
}

func NewBoid(bid int) *Boid {
	b := Boid{
		position: NewRandVector(),
		velocity: NewRandVector(),
		id:       bid,
	}
	return &b
}

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
