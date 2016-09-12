// Command-line program that generates a maze and saves it as a PNG file.
package main

import (
	"fmt"
	"github.com/lfritz/mazes/draw"
	"github.com/lfritz/mazes/generate"
	"github.com/lfritz/mazes/grid"
	"math/rand"
	"os"
)

func main() {
	generated := grid.NewMaze(12, 12, true)
	*generated.WallAbove(0, 0) = false
	*generated.WallAbove(generated.Width()-1, generated.Height()) = false
	generate.Backtracking(generated, rand.New(rand.NewSource(0)))

	draw.ToPNG(generated, "grid-maze.png")

	svgFile, err := os.Create("grid-maze.svg")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err.Error())
		return
	}
	draw.ToSVG(generated, svgFile)
}
