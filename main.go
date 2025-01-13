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

	numberOfAnts, rooms, tunnels := utils.CheckContent(fileContent)
	graph := utils.CreateGraph(tunnels)
	graph.Vertices = len(rooms)

	_, startRoom := utils.FindStart(rooms)
	_, endRoom := utils.FindEnd(rooms)

	// Step 1: Extract all paths
	allPaths := utils.ExtractAllPaths(graph, startRoom, endRoom, rooms)

	sort.Sort(PathSlice(allPaths))

	// Step 2: Filter non-intersecting groups
	nonIntersectingGroups := utils.FilterNonIntersectingGroups(allPaths)

	// Step 3: Remove smaller groups with in common members
	filteredGroups := utils.RemoveSmallerGroups(nonIntersectingGroups)

	// Step 4: Find best group of paths
	bestPathGroupNames := utils.FindBestPathGroup(filteredGroups, numberOfAnts)

	// Step 5: Assign ants to group of paths named solution
	solutions := utils.MakeAntsQueue(bestPathGroupNames, numberOfAnts)

	// Step 6: Print file contents
	for i := 0; i < len(fileContent); i++ {
		fmt.Println(fileContent[i])
	}
	fmt.Println()

	// Step 7: Move ants in solution
	utils.MoveAnts(solutions, bestPathGroupNames, rooms, numberOfAnts, endRoom)

	// for i := 0; i < len(rooms); i++ {
	// 	fmt.Println("{")
	// 	fmt.Println("Name:", rooms[i].Name, ",")
	// 	fmt.Println("Coord_x:", rooms[i].Coord_x, ",")
	// 	fmt.Println("Coord_y:", rooms[i].Coord_y, ",")
	// 	fmt.Println("IsStart:", rooms[i].IsStart, ",")
	// 	fmt.Println("IsEnd:", rooms[i].IsEnd, ",")
	// 	fmt.Println("},")
	// }
	// for i := 0; i < len(tunnels); i++ {
	// 	fmt.Println("{")
	// 	fmt.Println("FromRoom: utils.utils.Room{")
	// 	fmt.Println("Name:\"", tunnels[i].FromRoom.Name, "\",")
	// 	fmt.Println("Coord_x:", tunnels[i].FromRoom.Coord_x, ",")
	// 	fmt.Println("Coord_y:", tunnels[i].FromRoom.Coord_y, ",")
	// 	fmt.Println("},")
	// 	fmt.Println("ToRoom: utils.utils.Room{")
	// 	fmt.Println("Name:\"", tunnels[i].ToRoom.Name, "\",")
	// 	fmt.Println("Coord_x:", tunnels[i].ToRoom.Coord_x, ",")
	// 	fmt.Println("Coord_y:", tunnels[i].ToRoom.Coord_y, ",")
	// 	fmt.Println("},")
	// 	fmt.Println("},")
	// }
	// fmt.Println("numberOfAnts: ", numberOfAnts)
	// fmt.Println("Graph Edges:\n", graph.Edges)
	// fmt.Println("Graph Vertices:", graph.Vertices)

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

	// Print results
	// fmt.Println("All Paths:")
	// for pathIndex, path := range allPaths {
	// 	fmt.Print("path number: ", pathIndex+1, " ")
	// 	for _, room := range path {
	// 		fmt.Print(room.Name, " ")
	// 	}
	// 	fmt.Println()
	// }

	// fmt.Println("\nNon-Intersecting Groups:")
	// for groupIndex, group := range nonIntersectingGroups {
	// 	fmt.Print("group number: ", groupIndex+1, " ")
	// 	for _, path := range group {
	// 		for _, room := range path {
	// 			fmt.Print(room.Name, " ")
	// 		}
	// 		fmt.Print("| ")
	// 	}
	// 	fmt.Println()
	// }

	// fmt.Println("\nFilteredGroups Groups:")
	// for groupIndex, group := range filteredGroups {
	// 	fmt.Print("group number: ", groupIndex+1, " ")
	// 	for _, path := range group {
	// 		for _, room := range path {
	// 			fmt.Print(room.Name, " ")
	// 		}
	// 		fmt.Print("| ")
	// 	}
	// 	fmt.Println()
	// }

}
