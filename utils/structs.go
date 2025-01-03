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

// BFS to find the shortest path
func (g *Graph) bfs(start, end Room, parent map[string]string, rooms []Room) bool {
	visited := make(map[string]bool, g.Vertices)
	queue := []string{start.Name}
	visited[start.Name] = true

	for len(queue) > 0 {
		current := queue[0]

		for _, next := range g.Edges[current] {
			roomIndex := FindRoom(next, rooms)
			if !visited[next] && !rooms[roomIndex].AddedInPath {
				queue = append(queue, next)
				visited[next] = true
				parent[next] = current
				if next == end.Name {
					return true
				}
			}
		}
		queue = queue[1:]
	}
	return false
}

func (g *Graph) EdmondsKarp(start, end Room, rooms []Room) (int, [][]string) {
	parent := make(map[string]string, g.Vertices)
	maxFlow := 0
	var allPath [][]string

	for g.bfs(start, end, parent, rooms) {
		var path []string
		path = append(path, end.Name)
		// Each path contributes 1 unit of flow
		maxFlow++

		// Reverse the edges along the path to simulate residual graph
		for v := end.Name; v != start.Name; v = parent[v] {
			u := parent[v]
			path = append(path, u)
			// Remove forward edge
			// g.removeEdge(u, v)
			roomIndex := FindRoom(u, rooms)
			rooms[roomIndex].AddedInPath = true
			// Add reverse edge
			// g.adj[v] = append(g.adj[v], u)
		}

		allPath = append(allPath, path)
	}

	return maxFlow, allPath
}
