// Package grid defines a simple 2-D grid maze type.
package grid

import "fmt"

// A Maze is a maze with square cells in a rectangular grid.
type Maze struct {
	width  int
	height int
	walls  [][][]bool
}

func consistent(m *Maze) bool {
	if m.width < 1 || m.height < 1 {
		return false
	}
	if len(m.walls) != 2 {
		return false
	}

	for i, rows := range m.walls {
		expected_rows := m.width + i
		expected_cols := m.height + 1 - i
		if len(rows) != expected_rows {
			return false
		}
		for _, row := range rows {
			if len(row) != expected_cols {
				return false
			}
		}
	}
	return true
}

// NewMaze creates a new maze with the given width and height and all walls set
// or not set, depending on parameter setWalls.
func NewMaze(width, height int, setWalls bool) *Maze {
	walls := make([][][]bool, 2)
	for i := range walls {
		rows := width + i
		cols := height + 1 - i
		walls[i] = make([][]bool, rows)
		for x := range walls[i] {
			walls[i][x] = make([]bool, cols)
			if setWalls {
				for y := range walls[i][x] {
					walls[i][x][y] = true
				}
			}
		}
	}
	return &Maze{width, height, walls}
}

// Width returns the number of cells along the x axis.
func (m *Maze) Width() int {
	return m.width
}

// Height returns the number of cells along the y axis.
func (m *Maze) Height() int {
	return m.height
}

// WallAbove returns a pointer to the boolean for the wall above the cell at
// (x,y). For the walls below the last row, use y == Height().
func (m *Maze) WallAbove(x, y int) *bool {
	return &m.walls[0][x][y]
}

// WallLeftOf returns a pointer to the boolean for the wall left of the cell at
// (x,y). For the wall to the right of the last column, use x == Width().
func (m *Maze) WallLeftOf(x, y int) *bool {
	return &m.walls[1][x][y]
}

func (m *Maze) toIndex(x, y int) (i int) {
	return x + y*m.width
}

func (m *Maze) fromIndex(i int) (x, y int) {
	x = i % m.width
	y = i / m.width
	return
}

// N returns the number of cells.
func (m *Maze) N() int {
	return m.width * m.height
}

// Neighbors returns the neighboring cells of cell i.
func (m *Maze) Neighbors(i int) (result []int) {
	x, y := m.fromIndex(i)
	if y > 0 {
		result = append(result, m.toIndex(x, y-1))
	}
	if x > 0 {
		result = append(result, m.toIndex(x-1, y))
	}
	if x+1 < m.width {
		result = append(result, m.toIndex(x+1, y))
	}
	if y+1 < m.height {
		result = append(result, m.toIndex(x, y+1))
	}
	return
}

// AccessibleNeighbors returns those neighbors of node i where there's no wall
// between i and the neighbor node.
func (m *Maze) AccessibleNeighbors(i int) (result []int) {
	x, y := m.fromIndex(i)
	if y > 0 && !*m.WallAbove(x, y) {
		result = append(result, m.toIndex(x, y-1))
	}
	if x > 0 && !*m.WallLeftOf(x, y) {
		result = append(result, m.toIndex(x-1, y))
	}
	if x+1 < m.width && !*m.WallLeftOf(x+1, y) {
		result = append(result, m.toIndex(x+1, y))
	}
	if y+1 < m.height && !*m.WallAbove(x, y+1) {
		result = append(result, m.toIndex(x, y+1))
	}
	return
}

// Wall returns a pointer to the bool that determines if there's a wall between
// cells i and j.
func (m *Maze) Wall(i, j int) *bool {
	ix, iy := m.fromIndex(i)
	jx, jy := m.fromIndex(j)
	if iy == jy {
		if ix == jx+1 {
			return m.WallLeftOf(ix, iy)
		}
		if jx == ix+1 {
			return m.WallLeftOf(jx, jy)
		}
	}
	if ix == jx {
		if iy == jy+1 {
			return m.WallAbove(ix, iy)
		}
		if jy == iy+1 {
			return m.WallAbove(jx, jy)
		}
	}
	panic(fmt.Sprintf("Cannot get wall for nodes %v, %v", i, j))
}
