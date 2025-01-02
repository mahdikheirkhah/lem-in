package utils

func FindingPath(rooms []Room, graph Graph) [][]Room {
	var allPath [][]Room
	for _, room := range rooms {
		if room.IsStart {
			for _, node := range graph.Edges[room.Name] {
				var path []Room
				for _, room2 := range rooms {
					if room2.Name == node {
						path = append(path, room2)
					}
				}
				allPath = append(allPath, path)
			}
		}
	}

	return allPath
}
