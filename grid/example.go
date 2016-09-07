package grid

// Example returns a connected maze without loops of size 6x5. See example.png
// in the source directory.
func Example() *Maze {
	return &Maze{
		width:  6,
		height: 5,
		walls: [][][]bool{
			[][]bool{
				[]bool{false, false, true, false, false, true},
				[]bool{true, true, true, true, true, true},
				[]bool{true, false, true, true, false, true},
				[]bool{true, false, true, true, true, true},
				[]bool{true, true, false, true, false, true},
				[]bool{true, false, false, true, false, false},
			},
			[][]bool{
				[]bool{true, true, true, true, true},
				[]bool{false, false, false, true, false},
				[]bool{false, true, false, false, false},
				[]bool{true, false, false, false, true},
				[]bool{false, true, false, false, false},
				[]bool{false, true, false, true, false},
				[]bool{true, true, true, true, true},
			},
		},
	}
}
