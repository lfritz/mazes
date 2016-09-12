package graph

import (
	"github.com/lfritz/mazes/grid"
	"testing"
)

func TestDepthFirstSearch(t *testing.T) {
	type Case struct {
		g               MazeGraph
		connected, loop bool
	}
	var cases []Case

	// example maze is connected without a loop
	m := grid.Example()
	cases = append(cases, Case{m, true, false})

	// add a wall so it isn't connected anymore
	m = grid.Example()
	*m.WallLeftOf(2, 3) = true
	cases = append(cases, Case{m, false, false})

	// remove a wall to create a loop
	m = grid.Example()
	*m.WallAbove(2, 2) = false
	cases = append(cases, Case{m, true, true})

	// do both
	m = grid.Example()
	*m.WallLeftOf(2, 3) = true
	*m.WallAbove(2, 2) = false
	cases = append(cases, Case{m, false, true})

	for _, c := range cases {
		connected, loop := DepthFirstSearch(c.g)
		if connected != c.connected || loop != c.loop {
			t.Errorf("DepthFirstSearch returned %v,%v, want %v,%v",
				connected, loop, c.connected, c.loop)
		}
	}
}
