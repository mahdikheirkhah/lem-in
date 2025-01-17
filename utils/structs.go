package utils

type Room struct {
	Name        string
	Coord_x     int
	Coord_y     int
	IsStart     bool
	IsEnd       bool
	AddedInPath bool
}

type Tunnel struct {
	FromRoom Room
	ToRoom   Room
}
type Ant struct {
	Id               int
	PathIndex        int
	CurrentRoomName  string
	HasReachedTheEnd bool
}

type Solution struct {
	PathIndex int
	Ants      []Ant
}

// Define a graph using an adjacency list
type Graph struct {
	Vertices int
	Edges    map[string][]string
}

// Add an edge to the graph
func (g *Graph) AddEdge(from, to string) {
	g.Edges[from] = append(g.Edges[from], to)
	g.Edges[to] = append(g.Edges[to], from) // For undirected flow
}
