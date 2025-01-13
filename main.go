package main

import (
	"LemIn/errorHandler"
	"LemIn/fileHandler"
	"LemIn/utils"
	"errors"
	"fmt"
	"sort"
)

type PathSlice [][]utils.Room

func (p PathSlice) Len() int           { return len(p) }
func (p PathSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PathSlice) Less(i, j int) bool { return len(p[i]) < len(p[j]) }

func main() {
	fileName := utils.ReadFromCommandLine()
	fileContent := fileHandler.ReadAll(fileName)
	for i := 0; i < len(fileContent); i++ {
		fmt.Println("\"", fileContent[i], "\"", ",")
	}
	numberOfAnts, rooms, tunnels := utils.CheckContent(fileContent)
	graph := utils.CreateGraph(tunnels)
	graph.Vertices = len(rooms)
	for i := 0; i < len(rooms); i++ {
		fmt.Println("{")
		fmt.Println("Name:", rooms[i].Name, ",")
		fmt.Println("Coord_x:", rooms[i].Coord_x, ",")
		fmt.Println("Coord_y:", rooms[i].Coord_y, ",")
		fmt.Println("IsStart:", rooms[i].IsStart, ",")
		fmt.Println("IsEnd:", rooms[i].IsEnd, ",")
		fmt.Println("},")
	}
	for i := 0; i < len(tunnels); i++ {
		fmt.Println("{")
		fmt.Println("FromRoom: utils.utils.Room{")
		fmt.Println("Name:\"", tunnels[i].FromRoom.Name, "\",")
		fmt.Println("Coord_x:", tunnels[i].FromRoom.Coord_x, ",")
		fmt.Println("Coord_y:", tunnels[i].FromRoom.Coord_y, ",")
		fmt.Println("},")
		fmt.Println("ToRoom: utils.utils.Room{")
		fmt.Println("Name:\"", tunnels[i].ToRoom.Name, "\",")
		fmt.Println("Coord_x:", tunnels[i].ToRoom.Coord_x, ",")
		fmt.Println("Coord_y:", tunnels[i].ToRoom.Coord_y, ",")
		fmt.Println("},")
		fmt.Println("},")
	}
	fmt.Println("numberOfAnts: ", numberOfAnts)
	fmt.Println("Graph Edges:\n", graph.Edges)
	fmt.Println("Graph Vertices:", graph.Vertices)

	_, startRoom := utils.FindStart(rooms)
	_, endRoom := utils.FindEnd(rooms)

	// allPaths := graph.FindAllNonIntersectingPaths(startRoom, endRoom)

	// fmt.Println("all paths are: ", allPaths)
	// fmt.Println("number of min movements are: ", graph.SimulateAntMovement(allPaths, numberOfAnts))

	// var allPaths [][]utils.Room

	// // Find all paths from 'start' to 'end'
	// graph.Dfs(startRoom, endRoom, []utils.Room{}, &allPaths, rooms)

	// sort.Sort(PathSlice(allPaths))

	// // Print all found paths
	// fmt.Println("All possible paths:")

	// for i, path := range allPaths {
	// 	fmt.Println("path number: ", i+1, " path len: ", len(path))
	// 	for _, v := range path {
	// 		fmt.Print(v.Name, "->")
	// 	}
	// 	fmt.Println()
	// }

	// // allPaths2 := [][]utils.Room{
	// // 	{{Name: "start"}, {Name: "0"}, {Name: "o"}, {Name: "n"}, {Name: "e"}, {Name: "end"}},
	// // 	{{Name: "start"}, {Name: "0"}, {Name: "n"}, {Name: "e"}, {Name: "end"}},
	// // 	{{Name: "start"}, {Name: "h"}, {Name: "n"}, {Name: "e"}, {Name: "end"}},
	// // 	{{Name: "start"}, {Name: "h"}, {Name: "a"}, {Name: "n"}, {Name: "e"}, {Name: "end"}},
	// // }

	// nonIntersectingGroups := groupNonIntersectingPaths(allPaths)
	// // Print the result
	// fmt.Println("Non-Intersecting Path Groups:")
	// for i, group := range nonIntersectingGroups {
	// 	fmt.Printf("Group %d:\n", i+1)
	// 	for _, path := range group {
	// 		for _, room := range path {
	// 			fmt.Print(room.Name, "->")
	// 		}
	// 		fmt.Println()
	// 	}
	// }

	// // maxFlow, paths := graph.EdmondsKarp(startRoom, endRoom, rooms)
	// // fmt.Println("max flow is: ", maxFlow)
	// // fmt.Println("path is: ", paths)
	// // solutions := utils.MakeAntsQueue(paths, numberOfAnts)
	// // fmt.Println(solutions)
	// // utils.MoveAnts(solutions, paths, rooms, numberOfAnts, endRoom)

	// Step 1: Extract all paths
	allPaths := ExtractAllPathGroups(graph, startRoom, endRoom)
	if len(allPaths) < 1 {
		errorHandler.CheckError(errors.New("ERROR: invalid data format, invalid number of paths"), true)
	}

	sort.Sort(PathSlice(allPaths))

	// Step 2: Filter non-intersecting groups
	nonIntersectingGroups := FilterNonIntersectingGroups(allPaths)

	filteredGroups := RemoveSmallerGroups(nonIntersectingGroups)

	// Print results
	fmt.Println("All Paths:")
	for pathIndex, path := range allPaths {
		fmt.Print("path number: ", pathIndex+1, " ")
		for _, room := range path {
			fmt.Print(room.Name, " ")
		}
		fmt.Println()
	}

	fmt.Println("\nNon-Intersecting Groups:")
	for groupIndex, group := range nonIntersectingGroups {
		fmt.Print("group number: ", groupIndex+1, " ")
		for _, path := range group {
			for _, room := range path {
				fmt.Print(room.Name, " ")
			}
			fmt.Print("| ")
		}
		fmt.Println()
	}

	fmt.Println("\nFilteredGroups Groups:")
	for groupIndex, group := range filteredGroups {
		fmt.Print("group number: ", groupIndex+1, " ")
		for _, path := range group {
			for _, room := range path {
				fmt.Print(room.Name, " ")
			}
			fmt.Print("| ")
		}
		fmt.Println()
	}

	bestPathGroup, turnsCount := FindBestPathGroup(filteredGroups, numberOfAnts)
	// fmt.Println(bestPath)
	var bestPathGroupNames [][]string
	fmt.Println("turnsCount is: ", turnsCount)
	for _, path := range bestPathGroup {
		var pathNames []string
		for roomIndex, room := range path {
			if roomIndex != 0 {
				fmt.Print(room.Name, " ")
				pathNames = append(pathNames, room.Name)
			}
		}
		bestPathGroupNames = append(bestPathGroupNames, pathNames)
		fmt.Print("| ")
	}
	fmt.Println()

	solutions := utils.MakeAntsQueue(bestPathGroupNames, numberOfAnts)
	fmt.Println(solutions)
	utils.MoveAnts(solutions, bestPathGroupNames, rooms, numberOfAnts, endRoom)
}

// Function to check if two paths intersect
func pathsIntersect(path1, path2 []utils.Room) bool {
	nodeSet := make(map[string]bool)

	// Add all nodes from path1 to the set
	for _, room := range path1 {
		nodeSet[room.Name] = true
	}

	// Check if any node in path2 is in the set
	for _, room := range path2 {
		if nodeSet[room.Name] {
			return true
		}
	}
	return false
}

// Function to group paths into sets of non-intersecting paths
func groupNonIntersectingPaths(allPaths [][]utils.Room) [][][]utils.Room {
	var result [][][]utils.Room

	// Iterate over all paths
	for _, path := range allPaths {

		// Try to place the path into an existing group
		for i, group := range result {
			intersects := false
			for _, existingPath := range group {
				if pathsIntersect(path[1:len(path)-1], existingPath[1:len(existingPath)-1]) {
					intersects = true
					break
				}
			}

			// If no intersections, add the path to this group
			if !intersects {
				result[i] = append(result[i], path)
				// placed = true
			}
		}

		// If the path couldn't be placed in any existing group, create a new group
		result = append(result, [][]utils.Room{path})
	}

	return result
}

// ExtractAllPathGroups extracts all paths from start to end
func ExtractAllPathGroups(graph utils.Graph, start, end utils.Room) [][]utils.Room {
	var allPaths [][]utils.Room
	var currentPath []utils.Room

	// DFS to find all paths from start to end
	var dfs func(node utils.Room, visited map[string]bool)
	dfs = func(node utils.Room, visited map[string]bool) {
		if node.Name == end.Name {
			pathCopy := append([]utils.Room(nil), currentPath...)
			allPaths = append(allPaths, pathCopy)
			return
		}

		visited[node.Name] = true
		for _, neighborName := range graph.Edges[node.Name] {
			if !visited[neighborName] {
				neighbor := utils.Room{Name: neighborName} // Assuming Room can be reconstructed from name
				currentPath = append(currentPath, neighbor)
				dfs(neighbor, visited)
				currentPath = currentPath[:len(currentPath)-1]
			}
		}
		visited[node.Name] = false
	}

	// Start DFS
	currentPath = append(currentPath, start)
	dfs(start, make(map[string]bool))

	return allPaths
}

// FilterNonIntersectingGroups filters out groups with intersecting paths
func FilterNonIntersectingGroups(allPaths [][]utils.Room) [][][]utils.Room {
	var result [][][]utils.Room

	// Helper function to check if a path can be added to the group
	canAddPath := func(group [][]utils.Room, path []utils.Room) bool {
		usedNodes := make(map[string]bool)
		for _, existingPath := range group {
			for i := 1; i < len(existingPath)-1; i++ { // Skip start and end nodes
				usedNodes[existingPath[i].Name] = true
			}
		}
		for i := 1; i < len(path)-1; i++ { // Skip start and end nodes
			if usedNodes[path[i].Name] {
				return false
			}
		}
		return true
	}

	// Generate all combinations of paths
	n := len(allPaths)
	for i := 1; i < (1 << n); i++ {
		var group [][]utils.Room
		valid := true
		for j := 0; j < n; j++ {
			if (i & (1 << j)) != 0 {
				if canAddPath(group, allPaths[j]) {
					group = append(group, allPaths[j])
				} else {
					valid = false
					break
				}
			}
		}
		if valid {
			result = append(result, group)
		}
	}

	return result
}

func RemoveSmallerGroups(groups [][][]utils.Room) [][][]utils.Room {
	// Helper function to check if all paths in group1 exist in group2
	containsAllPaths := func(group1, group2 [][]utils.Room) bool {
		pathSet := make(map[string]struct{})
		for _, path := range group2 {
			var pathKey string
			for _, room := range path {
				pathKey += room.Name + ","
			}
			pathSet[pathKey] = struct{}{}
		}

		for _, path := range group1 {
			var pathKey string
			for _, room := range path {
				pathKey += room.Name + ","
			}
			if _, exists := pathSet[pathKey]; !exists {
				return false
			}
		}
		return true
	}

	// Filter groups
	var result [][][]utils.Room
	for i := 0; i < len(groups); i++ {
		include := true
		for j := 0; j < len(groups); j++ {
			if i != j && len(groups[j]) > len(groups[i]) && containsAllPaths(groups[i], groups[j]) {
				include = false
				break
			}
		}
		if include {
			result = append(result, groups[i])
		}
	}

	return result
}

func FindBestPathGroup(groups [][][]utils.Room, ants int) ([][]utils.Room, int) {
	bestGroup := [][]utils.Room{}
	minTime := int(^uint(0) >> 1) // Initialize to max int

	for _, group := range groups {
		pathLengths := make([]int, len(group))
		for i, path := range group {
			pathLengths[i] = len(path)
		}

		// Assign ants to paths using a greedy strategy
		numAnts := assignAntsToPaths(pathLengths, ants)

		// Calculate the time for this group
		maxTime := 0
		for i, length := range pathLengths {
			time := length + numAnts[i] - 1
			if time > maxTime {
				maxTime = time
			}
		}

		// Update the best group if this one is faster
		if maxTime < minTime {
			minTime = maxTime
			bestGroup = group
		}
	}

	return bestGroup, minTime
}

// Helper function to assign ants to paths
func assignAntsToPaths(pathLengths []int, ants int) []int {
	numPaths := len(pathLengths)
	numAnts := make([]int, numPaths)

	for ants > 0 {
		// Find the path with the smallest "load" (length + assigned ants)
		minIndex := 0
		minLoad := pathLengths[0] + numAnts[0]
		for i := 1; i < numPaths; i++ {
			load := pathLengths[i] + numAnts[i]
			if load < minLoad {
				minIndex = i
				minLoad = load
			}
		}

		// Assign an ant to the chosen path
		numAnts[minIndex]++
		ants--
	}

	return numAnts
}
