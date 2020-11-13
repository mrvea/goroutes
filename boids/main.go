package main

import (
	"log"

	"github.com/mrvea/goroutine/boids/env"

	"github.com/mrvea/goroutine/boids/boid"

	"github.com/hajimehoshi/ebiten"
)

var boids = map[int]*boid.Boid{}

func update(screen *ebiten.Image) error {
	if !ebiten.IsDrawingSkipped() {
		for _, boid := range boids {
			screen.Set(boid.PositionXInt()+1, boid.PositionYInt(), env.Green)
			screen.Set(boid.PositionXInt()-1, boid.PositionYInt(), env.Green)
			screen.Set(boid.PositionXInt(), boid.PositionYInt()-1, env.Green)
			screen.Set(boid.PositionXInt(), boid.PositionYInt()+1, env.Green)
		}
	}
	return nil
}
func startBoid(bid int) {
	b := boid.NewBoid(bid)
	boids[bid] = b
	go b.Start()
}

func main() {
	for i := 0; i < env.BoidCount; i++ {
		startBoid(i)
	}
	if err := ebiten.Run(update, env.ScreenWidth, env.ScreenHeight, 2, "Boids in a box"); err != nil {
		log.Fatal(err)
	}
}
