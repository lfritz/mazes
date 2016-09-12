package draw

import (
	"fmt"
	"github.com/ajstarks/svgo"
	"github.com/lfritz/mazes/grid"
	"io"
)

type Line struct {
	from, to int
}

func horizontalLines(m *grid.Maze, y int) []Line {
	var lines []Line
	width := m.Width()
	for x := 0; x < width; {
		for ; x < width && !*m.WallAbove(x, y); x++ {
		}
		start := x
		for ; x < width && *m.WallAbove(x, y); x++ {
		}
		end := x
		if end > start {
			lines = append(lines, Line{start, end})
		}
	}
	return lines
}

func verticalLines(m *grid.Maze, x int) []Line {
	var lines []Line
	height := m.Height()
	for y := 0; y < height; {
		for ; y < height && !*m.WallLeftOf(x, y); y++ {
		}
		start := y
		for ; y < height && *m.WallLeftOf(x, y); y++ {
		}
		end := y
		if end > start {
			lines = append(lines, Line{start, end})
		}
	}
	return lines
}

// Draw m as SVG.
func ToSVG(m *grid.Maze, w io.Writer) {
	const margin = 10
	const cell = 20

	canvas := svg.New(w)
	canvas.Start(2*margin+m.Width()*cell,
		2*margin+m.Height()*cell)

	canvas.Gtransform(fmt.Sprintf("translate(%d, %d)", margin, margin))
	canvas.Gstyle("stroke:black;stroke-linecap:round")

	for row := 0; row <= m.Height(); row++ {
		y := row * cell
		for _, l := range horizontalLines(m, row) {
			canvas.Line(l.from*cell, y, l.to*cell, y)
		}
	}

	for col := 0; col <= m.Width(); col++ {
		x := col * cell
		for _, l := range verticalLines(m, col) {
			canvas.Line(x, l.from*cell, x, l.to*cell)
		}
	}

	canvas.Gend()
	canvas.Gend()

	canvas.End()
}
