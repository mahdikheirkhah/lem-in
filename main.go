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
}
