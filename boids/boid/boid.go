package boid

import (
	"math"
	"time"

	"github.com/mrvea/goroutine/boids/env"
)

var (
	Boids [env.BoidCount]*Boid
)

//Boid struct node to represent a boid
type Boid struct {
	position Vector2d
	velocity Vector2d
	id       int
}

//NewBoid makes a new boid with random position and velocity
func NewBoid(bid int) *Boid {
	v := NewRandVector(2, 2)
	v.X, v.Y = v.X-1.0, v.Y-1.0
	b := Boid{
		position: NewRandVector(env.ScreenWidth, env.ScreenHeight),
		velocity: v,
		id:       bid,
	}
	return &b
}

//ID returns the id fo the boid
func (b *Boid) ID() int {
	return b.id
}

//Start initilizes a boid in an infinite loop for a goroutine
func (b *Boid) Start() {
	for {
		b.moveOne()
		time.Sleep(5 * time.Millisecond)
	}
}

func (b *Boid) calcAcceleration() Vector2d {
	upper, lower := b.position.AddF(env.ViewRadius), b.position.AddF(-env.ViewRadius)
	avgPosition, avgVelocity, separation := Vector2d{0, 0}, Vector2d{0, 0}, Vector2d{0, 0}
	count := 0.0
	leftBound := int(math.Max(lower.X, 0))
	rightBound := int(math.Min(upper.X, env.ScreenWidth))
	topBound := int(math.Min(upper.Y, env.ScreenHeight))
	bottomBound := int(math.Max(lower.Y, 0))
	env.RWlock.RLock()
	for i := leftBound; i <= rightBound; i++ {
		for j := bottomBound; j <= topBound; j++ {
			if otherBoidID := env.BoidMap[i][j]; otherBoidID != -1 && otherBoidID != b.id {
				if dist := Boids[otherBoidID].position.Distance(b.position); dist < env.ViewRadius {
					count++
					avgVelocity = avgVelocity.Add(Boids[otherBoidID].velocity)
					avgPosition = avgPosition.Add(Boids[otherBoidID].position)
					separation = separation.Add(b.position.Sub(Boids[otherBoidID].position).DivF(dist))
				}
			}
		}
	}
	env.RWlock.RUnlock()
	accel := Vector2d{
		X: b.borderBounce(b.position.X, env.ScreenWidth),
		Y: b.borderBounce(b.position.Y, env.ScreenHeight),
	}
	if count > 0 {
		avgPosition, avgVelocity = avgPosition.DivF(count), avgVelocity.DivF(count)
		accelAlignment := avgVelocity.Sub(b.velocity).MulF(env.AdjRate)
		accelCohision := avgPosition.Sub(b.position).MulF(env.AdjRate)
		accelSeparation := separation.MulF(env.AdjRate)
		accel = accel.Add(accelAlignment).Add(accelCohision).Add(accelSeparation)
	}
	return accel
}

func (b *Boid) borderBounce(pos, maxBound float64) float64 {
	val := 0.0
	if pos < env.ViewRadius {
		val = 1 / pos
	} else if pos > maxBound-env.ViewRadius {
		val = 1 / (pos - maxBound)
	}

	return val
}

func (b *Boid) moveOne() {
	accel := b.calcAcceleration()
	env.RWlock.Lock()
	b.velocity = b.velocity.Add(accel).limit(-1, 1)
	env.BoidMap[b.PositionXInt()][b.PositionYInt()] = -1
	b.position = b.position.Add(b.velocity)
	env.BoidMap[b.PositionXInt()][b.PositionYInt()] = b.id
	// next := b.position.Add(b.velocity)
	// if next.X >= env.ScreenWidth || next.X < 0 {
	// 	b.bounceOffX()
	// }
	// if next.Y >= env.ScreenHeight || next.Y < 0 {
	// 	b.bounceOffY()
	// }
	env.RWlock.Unlock()
}

func (b *Boid) bounceOffX() {
	b.velocity = Vector2d{X: -b.velocity.X, Y: b.velocity.Y}
}
func (b *Boid) bounceOffY() {
	b.velocity = Vector2d{X: b.velocity.X, Y: -b.velocity.Y}
}

//PositionXInt return position X as an int
func (b *Boid) PositionXInt() int {
	return int(b.position.X)
}

//PositionYInt return position Y as an int
func (b *Boid) PositionYInt() int {
	return int(b.position.Y)
}
