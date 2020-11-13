package boid

import (
	"math"
	"math/rand"

	"github.com/mrvea/goroutine/boids/env"
)

//Vector2d a struct with coordinates X,Y
type Vector2d struct {
	X float64
	Y float64
}

//NewRandVector makes a random 2d vector
func NewRandVector() Vector2d {
	return Vector2d{
		X: rand.Float64() * env.ScreenWidth,
		Y: rand.Float64() * env.ScreenHeight,
	}
}

//Add adds a 2d vector to another 2d vector
func (v1 Vector2d) Add(v2 Vector2d) Vector2d {
	return Vector2d{X: v1.X + v2.X, Y: v1.Y + v2.Y}
}

//Sub subtracts a 2d vector from another 2d vector
func (v1 Vector2d) Sub(v2 Vector2d) Vector2d {
	return Vector2d{X: v1.X - v2.X, Y: v1.Y - v2.Y}
}

//Mul multiplys a 2d vector by another 2d vector
func (v1 Vector2d) Mul(v2 Vector2d) Vector2d {
	return Vector2d{X: v1.X * v2.X, Y: v1.Y * v2.Y}
}

// AddF adds a float to a 2d vector
func (v1 Vector2d) AddF(d float64) Vector2d {
	return Vector2d{X: v1.X + d, Y: v1.Y + d}
}

// MulF multiplys a 2d vector by a float
func (v1 Vector2d) MulF(d float64) Vector2d {
	return Vector2d{X: v1.X * d, Y: v1.Y * d}
}

// DivF divides a 2d vector by a float
func (v1 Vector2d) DivF(d float64) Vector2d {
	return Vector2d{X: v1.X / d, Y: v1.Y / d}
}

func (v1 Vector2d) limit(lower, upper float64) Vector2d {
	return Vector2d{
		X: math.Min(math.Max(v1.X, lower), upper),
		Y: math.Min(math.Max(v1.Y, lower), upper),
	}
}

//Distance calculates a length of the hypotenuse
func (v1 Vector2d) Distance(v2 Vector2d) float64 {
	return math.Sqrt(math.Pow(v1.X-v2.X, 2) + math.Pow(v1.Y-v2.Y, 2))
}
