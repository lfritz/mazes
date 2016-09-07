// Package graph contains definitions to view a maze as an undirected graph.
package graph

// MazeGraph is an interface to treat a maze as an undirected graph, where the
// cells of the maze correspond to nodes and neighboring cells are connected by
// edges. The nodes are numbered from 0 to N() - 1.
type MazeGraph interface {
	// N returns the number of nodes / cells.
	N() int

	// Neighbors returns the neighboring cells of cell i.
	Neighbors(i int) []int

	// AccessibleNeighbors returns those neighbors of node i where there's no
	// wall between i and the neighbor node.
	AccessibleNeighbors(i int) []int

	// Wall returns a pointer to the bool that determines if there's a wall
	// between cells i and j.
	Wall(i, j int) *bool
}

// DepthFirstSearch searches through the graph and determines if it's connected
// and if there are any loops.
func DepthFirstSearch(g MazeGraph) (connected, loop bool) {
	visited := make([]bool, g.N())
	var search func(node, parent int)
	search = func(node, parent int) {
		visited[node] = true
		for _, n := range g.AccessibleNeighbors(node) {
			if n == parent {
				continue
			}
			if visited[n] {
				loop = true
			} else {
				search(n, node)
			}
		}
	}
	search(0, 0)

	connected = true
	for _, v := range visited {
		if !v {
			connected = false
			break
		}
	}

	return
}
