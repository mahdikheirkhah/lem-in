package utils

import (
	"LemIn/errorHandler"
	"errors"
)

// ExtractAllPaths extracts all paths from start to end
func ExtractAllPaths(graph Graph, start, end Room, rooms []Room) [][]Room {
	var allPaths [][]Room
	var currentPath []Room

	// DFS to find all paths from start to end
	var dfs func(node Room, visited map[string]bool)
	dfs = func(node Room, visited map[string]bool) {
		if node.Name == end.Name {
			pathCopy := append([]Room(nil), currentPath...)
			allPaths = append(allPaths, pathCopy)
			return
		}

		visited[node.Name] = true
		for _, neighborName := range graph.Edges[node.Name] {
			if !visited[neighborName] {
				neighborIndex := FindRoom(neighborName, rooms)
				neighbor := rooms[neighborIndex] // Assuming Room can be reconstructed from name
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

	if len(allPaths) < 1 {
		errorHandler.CheckError(errors.New("ERROR: invalid data format, no path found"), true)
		return nil
	}

	return allPaths
}

func FilterNonIntersectingGroups(allPaths [][]Room) [][][]Room {
	var result [][][]Room

	// Helper function to check if a path can be added to the group
	canAddPath := func(group [][]Room, path []Room) bool {
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

	// Backtracking function to build groups
	var backtrack func(start int, group [][]Room)
	backtrack = func(start int, group [][]Room) {
		// Add the current group to the result
		if len(group) > 0 {
			groupCopy := make([][]Room, len(group))
			copy(groupCopy, group)
			result = append(result, groupCopy)
		}

		// Iterate over remaining paths
		for i := start; i < len(allPaths); i++ {
			if canAddPath(group, allPaths[i]) {
				// Add path to the group
				group = append(group, allPaths[i])
				// Recur to add more paths
				backtrack(i+1, group)
				// Backtrack by removing the last path
				group = group[:len(group)-1]
			}
		}
	}

	// Start backtracking from the first path
	backtrack(0, [][]Room{})
	return result
}

func RemoveSmallerGroups(groups [][][]Room) [][][]Room {
	// Helper function to check if all paths in group1 exist in group2
	containsAllPaths := func(group1, group2 [][]Room) bool {
		pathSet := make(map[string]bool)
		for _, path := range group2 {
			var pathKey string
			for _, room := range path {
				pathKey += room.Name + ","
			}
			pathSet[pathKey] = true
		}

		for _, path := range group1 {
			var pathKey string
			for _, room := range path {
				pathKey += room.Name + ","
			}
			if !pathSet[pathKey] {
				return false
			}
		}
		return true
	}

	// Filter groups
	var result [][][]Room
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

func FindBestPathGroup(groups [][][]Room, ants int) [][]string {
	bestGroup := [][]Room{}
	minTime := int(^uint(0) >> 1) // Initialize to max int

	for _, group := range groups {
		pathLengths := make([]int, len(group))
		for i, path := range group {
			pathLengths[i] = len(path)
		}

		// Assign ants to paths using a greedy strategy
		numAnts := assignAntsToPaths(pathLengths, ants)

		// Calculate the time for this group which is equal to the longest time between all paths of group
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

	// Here we found bestGroup but we need names of the rooms in paths of group so we do this
	var bestPathGroupNames [][]string
	for _, path := range bestGroup {
		var pathNames []string
		for roomIndex, room := range path {
			if roomIndex != 0 {
				pathNames = append(pathNames, room.Name)
			}
		}
		bestPathGroupNames = append(bestPathGroupNames, pathNames)
	}

	return bestPathGroupNames
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
