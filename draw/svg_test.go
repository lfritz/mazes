package draw

import (
	"github.com/lfritz/mazes/grid"
	"reflect"
	"testing"
)

func TestHorizontalLines(t *testing.T) {
	m := grid.Example()
	cases := [][]line{
		{{1, 6}},
		{{1, 2}, {4, 5}},
		{{0, 4}},
		{{1, 6}},
		{{1, 2}, {3, 4}},
		{{0, 5}},
	}
	for y, c := range cases {
		got := horizontalLines(m, y)
		if !reflect.DeepEqual(got, c) {
			t.Errorf("horizontalLines(m, %v) == %v, want %v", y, got, c)
		}
	}
}

func TestVerticalLines(t *testing.T) {
	m := grid.Example()
	cases := [][]line{
		{{0, 5}},
		{{3, 4}},
		{{1, 2}},
		{{0, 1}, {4, 5}},
		{{1, 2}},
		{{1, 2}, {3, 4}},
		{{0, 5}},
	}
	for x, c := range cases {
		got := verticalLines(m, x)
		if !reflect.DeepEqual(got, c) {
			t.Errorf("verticalLines(m, %v) == %v, want %v", x, got, c)
		}
	}
}
