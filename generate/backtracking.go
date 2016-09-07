// Package generate contains functions to randomly generate mazes.
package generate

import (
	"github.com/lfritz/maze/graph"
	"math/rand"
)

// Backtracking generates a maze with recursive backtracking. Parameter g
// should be a maze with all walls set.
func Backtracking(g graph.MazeGraph, r *rand.Rand) {
	n := g.N()
	current := rand.Intn(n - 1)
	visited := make([]bool, n)
	visited[current] = true
	var stack []int

	for {
		neighbors := g.Neighbors(current)
		var notVisited []int
		for _, n := range neighbors {
			if !visited[n] {
				notVisited = append(notVisited, n)
			}
		}

		if len(notVisited) > 0 {
			next := notVisited[r.Intn(len(notVisited))]
			*g.Wall(current, next) = false
			visited[next] = true
			stack = append(stack, current)
			current = next
		} else if len(stack) > 0 {
			next := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			current = next
		} else {
			break
		}
	}
}
