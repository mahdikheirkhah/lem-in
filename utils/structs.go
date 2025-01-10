package utils

import (
	"LemIn/errorHandler"
	"errors"
)

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
	id               int
	pathIndex        int
	currentRoomName  string
	hasReachedTheEnd bool
}

type Solution struct {
	pathIndex int
	ants      []Ant
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
		path = reverseSlice(path)
		allPath = append(allPath, path)
	}
	if len(allPath) == 0 {
		errorHandler.CheckError(errors.New("ERROR: invalid data format, no path found"), true)
	}
	return maxFlow, allPath
}

func reverseSlice(slice []string) []string {
	var result []string
	if len(slice) < 2 {
		errorHandler.CheckError(errors.New("invalid path"), true)
	}
	for i := len(slice) - 2; i > -1; i-- {
		result = append(result, slice[i])
	}
	return result
}

// DFS function to find all paths from start to end
func (g *Graph) Dfs(current, end Room, path []Room, allPaths *[][]Room, rooms []Room) {
	// Add current node to path
	path = append(path, current)

	// If we reach the destination node, add the current path to allPaths
	if current == end {
		*allPaths = append(*allPaths, append([]Room(nil), path...)) // Append a copy of the path
		return
	}

	// Explore all neighbors of the current node
	for _, neighbor := range g.Edges[current.Name] {
		neighborRoomIndex := FindRoom(neighbor, rooms)
		neighborRoom := rooms[neighborRoomIndex]
		// Avoid revisiting the same node in the current path
		if !contains(path, neighborRoom) {
			g.Dfs(neighborRoom, end, path, allPaths, rooms)
		}
	}
}

// Helper function to check if a node is in the path
func contains(path []Room, node Room) bool {
	for _, p := range path {
		if p == node {
			return true
		}
	}
	return false
}

func (g *Graph) bfsFindPathAvoidingEdges(start, end Room, usedEdges map[string]map[string]bool) ([]string, bool) {
	visited := make(map[string]bool)
	parent := make(map[string]string)
	queue := []string{start.Name}
	visited[start.Name] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, next := range g.Edges[current] {
			if visited[next] || (usedEdges[current] != nil && usedEdges[current][next]) {
				continue
			}

			visited[next] = true
			parent[next] = current
			queue = append(queue, next)

			if next == end.Name {
				// Construct the path
				var path []string
				for v := end.Name; v != ""; v = parent[v] {
					path = append(path, v)
				}
				return reverseSliceNew(path), true
			}
		}
	}
	return nil, false
}

func (g *Graph) FindAllNonIntersectingPaths(start, end Room) [][]string {
	usedEdges := make(map[string]map[string]bool)
	var allPaths [][]string

	for {
		path, found := g.bfsFindPathAvoidingEdges(start, end, usedEdges)
		if !found {
			break
		}

		// Mark edges along the path as used
		for i := 0; i < len(path)-1; i++ {
			u, v := path[i], path[i+1]
			if usedEdges[u] == nil {
				usedEdges[u] = make(map[string]bool)
			}
			if usedEdges[v] == nil {
				usedEdges[v] = make(map[string]bool)
			}
			usedEdges[u][v] = true
			usedEdges[v][u] = true // Mark both directions to make it undirected
		}

		allPaths = append(allPaths, path)
	}

	if len(allPaths) == 0 {
		errorHandler.CheckError(errors.New("ERROR: No paths found"), true)
	}
	return allPaths
}

func reverseSliceNew(slice []string) []string {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}

func (g *Graph) SimulateAntMovement(paths [][]string, antCount int) int {
	// Create an array to track ants on each path
	steps := 0
	antAssignments := make([]int, len(paths))

	// Assign ants to paths based on their lengths
	for i := 0; i < antCount; i++ {
		// Assign the next ant to the shortest available path
		shortestIndex := 0
		for j := 1; j < len(paths); j++ {
			if len(paths[j])+antAssignments[j] < len(paths[shortestIndex])+antAssignments[shortestIndex] {
				shortestIndex = j
			}
		}
		antAssignments[shortestIndex]++
	}

	// Simulate movement until all ants reach the destination
	for {
		allDone := true
		for i, path := range paths {
			if antAssignments[i] > 0 {
				antAssignments[i]--
				if len(path)-1+antAssignments[i] > steps {
					steps = len(path) - 1 + antAssignments[i]
				}
				allDone = false
			}
		}
		if allDone {
			break
		}
	}
	return steps
}
