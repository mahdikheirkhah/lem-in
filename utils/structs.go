package utils

type Room struct {
	Name    string
	Coord_x int
	Coord_y int
}

type Tunnel struct {
	FromRoom Room
	ToRoom   Room
}

// Define a graph using an adjacency list
type Graph struct {
	Edges map[string][]string
}

// Add an edge to the graph
func (g *Graph) AddEdge(from, to string) {
	g.Edges[from] = append(g.Edges[from], to)
}
