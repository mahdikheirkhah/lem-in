package utils

import (
	"LemIn/fileHandler"
	"fmt"
	"sort"
)

type PathSlice [][]Room

func (p PathSlice) Len() int           { return len(p) }
func (p PathSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PathSlice) Less(i, j int) bool { return len(p[i]) < len(p[j]) }

func Lem_in(fileName string) {
	fileContent := fileHandler.ReadAll(fileName)

	numberOfAnts, rooms, tunnels := CheckContent(fileContent)
	if numberOfAnts == -1 || rooms == nil || tunnels == nil {
		return
	}
	graph := CreateGraph(tunnels)
	graph.Vertices = len(rooms)

	_, startRoom := FindStart(rooms)
	_, endRoom := FindEnd(rooms)

	// Step 1: Extract all paths
	allPaths := ExtractAllPaths(graph, startRoom, endRoom, rooms)
	if allPaths == nil {
		return
	}

	sort.Sort(PathSlice(allPaths))

	// Step 2: Filter non-intersecting groups
	nonIntersectingGroups := FilterNonIntersectingGroups(allPaths)

	// Step 3: Remove smaller groups with in common members
	filteredGroups := RemoveSmallerGroups(nonIntersectingGroups)

	// Step 4: Find best group of paths
	bestPathGroupNames := FindBestPathGroup(filteredGroups, numberOfAnts)

	// Step 5: Assign ants to group of paths named solution
	solutions := MakeAntsQueue(bestPathGroupNames, numberOfAnts)

	// Step 6: Print file contents
	for i := 0; i < len(fileContent); i++ {
		fmt.Println(fileContent[i])
	}
	fmt.Println()

	// Step 7: Move ants in solution
	MoveAnts(solutions, bestPathGroupNames, rooms, numberOfAnts, endRoom)
}
