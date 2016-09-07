// Command-line program that generates a maze and saves it as a PNG file.
package main

import (
	"github.com/lfritz/maze/draw"
	"github.com/lfritz/maze/generate"
	"github.com/lfritz/maze/grid"
	"math/rand"
)

func main() {
	generated := grid.NewMaze(12, 12, true)
	*generated.WallAbove(0, 0) = false
	*generated.WallAbove(generated.Width()-1, generated.Height()) = false
	generate.Backtracking(generated, rand.New(rand.NewSource(0)))
	draw.ToPNG(generated, "grid-maze.png")
}
