package main

import (
	"log"

	"github.com/mrvea/goroutine/boids/boid"

	"github.com/hajimehoshi/ebiten"
)

var boids = map[int]*boid.Boid{}

func update(screen *ebiten.Image) error {
	if !ebiten.IsDrawingSkipped() {
		for _, boid := range boids {
			screen.Set(int(boid.position.x+1), int(boid.position.y+1), green)
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
	if err := ebiten.Run(update, screenWidth, screenHeight, 2, "Boids in a box"); err != nil {
		log.Fatal(err)
	}
}
