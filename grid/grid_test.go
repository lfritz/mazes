package grid

import (
	"reflect"
	"testing"
)

func TestNewMaze(t *testing.T) {
	width, height := 3, 4
	m := NewMaze(width, height, true)
	if m.width != width || m.height != height {
		t.Errorf("NewMaze returned maze with wrong size")
	}
	if !consistent(m) {
		t.Errorf("NewMaze returned inconsistent Maze struct")
	}
OuterLoop:
	for _, w := range m.walls {
		for _, x := range w {
			for _, y := range x {
				if !y {
					t.Errorf("NewSimpleMaze returned maze with missing wall")
					break OuterLoop
				}
			}
		}
	}
}

func TestWidth(t *testing.T) {
	m := Example()
	if m.Width() != 6 {
		t.Errorf("m.Width() == %v, want 6", m.Width())
	}
}

func TestHeight(t *testing.T) {
	m := Example()
	if m.Height() != 5 {
		t.Errorf("m.Height() == %v, want 6", m.Height())
	}
}

func TestWallAbove(t *testing.T) {
	m := Example()
	cases := []struct {
		x, y int
		want bool
	}{
		{0, 0, false},
		{1, 0, true},
		{4, 1, true},
		{4, 2, false},
	}
	for _, c := range cases {
		got := *m.WallAbove(c.x, c.y)
		if got != c.want {
			t.Errorf("m.WallAbove(%v, %v) == %v, want %v",
				c.x, c.y, got, c.want)
		}
	}
}

func TestWallLeftOf(t *testing.T) {
	m := Example()
	cases := []struct {
		x, y int
		want bool
	}{
		{0, 1, true},
		{1, 2, false},
		{4, 1, true},
		{4, 2, false},
	}
	for _, c := range cases {
		got := *m.WallLeftOf(c.x, c.y)
		if got != c.want {
			t.Errorf("m.WallLeftOf(%v, %v) == %v, want %v",
				c.x, c.y, got, c.want)
		}
	}
}

func TestToIndex(t *testing.T) {
	m := Example()
	got := m.toIndex(1, 2)
	if got != 13 {
		t.Errorf("m.toIndex(1, 2) == %v, want 13", got)
	}
}

func TestFromIndex(t *testing.T) {
	m := Example()
	x, y := m.fromIndex(13)
	if !(x == 1 && y == 2) {
		t.Errorf("m.fromIndex(13) == %v,%v, want 1,2", x, y)
	}
}

func TestN(t *testing.T) {
	m := Example()
	got := m.N()
	if got != 30 {
		t.Errorf("m.N() == %v, want 30", got)
	}
}

func TestNeighbors(t *testing.T) {
	m := Example()
	cases := []struct {
		i    int
		want []int
	}{
		{24, []int{18, 25}},
		{18, []int{12, 19, 24}},
		{13, []int{7, 12, 14, 19}},
	}
	for _, c := range cases {
		got := m.Neighbors(c.i)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("m.Neighbors(%v) == %v, want %v", c.i, got, c.want)
		}
	}
}

func TestAccessibleNeighbors(t *testing.T) {
	m := Example()
	cases := []struct {
		i    int
		want []int
	}{
		{10, []int{16}},
		{17, []int{11, 16}},
		{16, []int{10, 15, 17}},
	}
	for _, c := range cases {
		got := m.AccessibleNeighbors(c.i)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("m.AccessibleNeighbors(%v) == %v, want %v",
				c.i, got, c.want)
		}
	}
}

func TestWall(t *testing.T) {
	m := Example()
	cases := []struct {
		i, j int
		want bool
	}{
		// horizontal neighbors
		{0, 1, false},
		{22, 21, false},
		{2, 3, true},
		{23, 22, true},
		// vertical neighbors
		{5, 11, false},
		{26, 20, false},
		{6, 12, true},
		{27, 21, true},
	}
	for _, c := range cases {
		got := *m.Wall(c.i, c.j)
		if got != c.want {
			t.Errorf("m.Wall(%v, %v) == %v, want %v", c.i, c.j, got, c.want)
		}
	}
}
