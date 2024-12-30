package utils

func CreateGraph(tunnels []Tunnel) Graph {
	graph := Graph{Edges: make(map[string][]string)}
	for _, tunnel := range tunnels {
		graph.AddEdge(tunnel.FromRoom.Name, tunnel.ToRoom.Name)
	}
	return graph
}
