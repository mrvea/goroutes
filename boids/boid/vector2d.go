package boid

import (
	"math"
	"math/rand"

	"github.com/mrvea/goroutine/boids/env"
)

type Vector2d struct {
	X float64
	Y float64
}

func NewRandVector() Vector2d {
	return Vector2d{
		X: rand.Float64() * env.ScreenWidth,
		Y: rand.Float64() * env.ScreenHeight,
	}
}

func (v1 Vector2d) Add(v2 Vector2d) Vector2d {
	return Vector2d{X: v1.X + v2.X, Y: v1.Y + v2.Y}
}
func (v1 Vector2d) Sub(v2 Vector2d) Vector2d {
	return Vector2d{X: v1.X - v2.X, Y: v1.Y - v2.Y}
}
func (v1 Vector2d) Mul(v2 Vector2d) Vector2d {
	return Vector2d{X: v1.X * v2.X, Y: v1.Y * v2.Y}
}
func (v1 Vector2d) AddF(d float64) Vector2d {
	return Vector2d{X: v1.X + d, Y: v1.Y + d}
}

func (v1 Vector2d) MulF(d float64) Vector2d {
	return Vector2d{X: v1.X * d, Y: v1.Y * d}
}

func (v1 Vector2d) DivF(d float64) Vector2d {
	return Vector2d{X: v1.X / d, Y: v1.Y / d}
}

func (v1 Vector2d) limit(lower, upper float64) Vector2d {
	return Vector2d{
		X: math.Min(math.Max(v1.X, lower), upper),
		Y: math.Min(math.Max(v1.Y, lower), upper),
	}
}
