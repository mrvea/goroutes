package main

import (
	"log"

	"github.com/mrvea/goroutine/boids/env"

	"github.com/mrvea/goroutine/boids/boid"

	"github.com/hajimehoshi/ebiten"
)

func update(screen *ebiten.Image) error {
	if !ebiten.IsDrawingSkipped() {
		for _, boid := range boid.Boids {
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
	boid.Boids[bid] = b
	env.BoidMap[b.PositionXInt()][b.PositionYInt()] = bid
	go b.Start()
}
func initBoidMap() {
	for i, row := range env.BoidMap {
		for j := range row {
			env.BoidMap[i][j] = -1
		}
	}
}

func main() {
	initBoidMap()
	for i := 0; i < env.BoidCount; i++ {
		startBoid(i)
	}
	if err := ebiten.Run(update, env.ScreenWidth, env.ScreenHeight, 2, "Boids in a box"); err != nil {
		log.Fatal(err)
	}
}
