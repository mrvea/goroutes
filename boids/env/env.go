package env

import (
	"image/color"
)

const (
	ScreenWidth, ScreenHeight = 640, 360
	BoidCount                 = 500
	ViewRadius                = 13
	AdjRate                   = 0.015
)

var (
	Green = color.RGBA{R: 10, G: 255, B: 50, A: 255}

	BoidMap [ScreenWidth + 1][ScreenHeight + 1]int
)
