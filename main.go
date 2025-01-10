package main

import (
	"LemIn/fileHandler"
	"LemIn/utils"
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

	var allPaths [][]utils.Room

	// Find all paths from 'start' to 'end'
	graph.Dfs(startRoom, endRoom, []utils.Room{}, &allPaths, rooms)

	sort.Sort(PathSlice(allPaths))

	// Print all found paths
	fmt.Println("All possible paths:")

	for i, path := range allPaths {
		fmt.Println("path number: ", i+1, " path len: ", len(path))
		for _, v := range path {
			fmt.Print(v.Name, "->")
		}
		fmt.Println()
	}

	// allPaths2 := [][]utils.Room{
	// 	{{Name: "start"}, {Name: "0"}, {Name: "o"}, {Name: "n"}, {Name: "e"}, {Name: "end"}},
	// 	{{Name: "start"}, {Name: "0"}, {Name: "n"}, {Name: "e"}, {Name: "end"}},
	// 	{{Name: "start"}, {Name: "h"}, {Name: "n"}, {Name: "e"}, {Name: "end"}},
	// 	{{Name: "start"}, {Name: "h"}, {Name: "a"}, {Name: "n"}, {Name: "e"}, {Name: "end"}},
	// }

	nonIntersectingGroups := groupNonIntersectingPaths(allPaths)
	// Print the result
	fmt.Println("Non-Intersecting Path Groups:")
	for i, group := range nonIntersectingGroups {
		fmt.Printf("Group %d:\n", i+1)
		for _, path := range group {
			for _, room := range path {
				fmt.Print(room.Name, "->")
			}
			fmt.Println()
		}
	}

	// maxFlow, paths := graph.EdmondsKarp(startRoom, endRoom, rooms)
	// fmt.Println("max flow is: ", maxFlow)
	// fmt.Println("path is: ", paths)
	// solutions := utils.MakeAntsQueue(paths, numberOfAnts)
	// fmt.Println(solutions)
	// utils.MoveAnts(solutions, paths, rooms, numberOfAnts, endRoom)
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
