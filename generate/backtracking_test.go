package generate

import (
	"github.com/lfritz/mazes/graph"
	"github.com/lfritz/mazes/grid"
	"math/rand"
	"reflect"
	"testing"
	"testing/quick"
)

func TestBacktracking(t *testing.T) {
	f := func(width, height int, r *rand.Rand) bool {
		m := grid.NewMaze(width, height, true)
		Backtracking(m, r)
		connected, loop := graph.DepthFirstSearch(m)
		return connected && !loop
	}
	config := quick.Config{
		// override the Values function so width and height are reasonable
		Values: func(vs []reflect.Value, r *rand.Rand) {
			vs[0] = reflect.ValueOf(r.Intn(100) + 1)
			vs[1] = reflect.ValueOf(r.Intn(100) + 1)
			vs[2] = reflect.ValueOf(r)
		},
	}
	if err := quick.Check(f, &config); err != nil {
		t.Error(err)
	}
}
