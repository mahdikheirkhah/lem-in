package utils

import "fmt"

func MakeAntsQueue(paths [][]string, numberOfAnts int) []Solution {
	numberOfPaths := len(paths)

	ants := make([]Ant, numberOfAnts)
	initAnts(numberOfAnts, ants)

	fmt.Println("ants before update: ", ants)

	solutions := make([]Solution, numberOfPaths)
	initSolutions(solutions)

	pathIndex := 0
	// AntsNumberInEachPath(ants)
	for antsIndex := 0; antsIndex < numberOfAnts; antsIndex++ {
		pathIndex = findSuitablePath(pathIndex, paths, solutions)

		ants[antsIndex].pathIndex = pathIndex
		solutions[pathIndex].ants = append(solutions[pathIndex].ants, ants[antsIndex])
	}

	fmt.Println("ants after update: ", ants)

	return solutions
	// for i := 1; i < numberOfAnts+1; i++ {
	// 	findTheSuitablePath(paths, antsQueue)
	// }
}

func findSuitablePath(lastPathIndex int, paths [][]string, solutions []Solution) int {
	//minLen := len(paths[lastPathIndex]) + len(solutions[lastPathIndex].ants)
	cntr := 0
	for i := lastPathIndex; cntr < len(solutions) && i < len(solutions); i++ {
		cntr++
		// if len(paths[i])+len(solutions[i].ants) <= minLen {
		if i == len(solutions)-1 {
			if len(paths[i])+len(solutions[i].ants) > len(paths[0])+len(solutions[0].ants) {
				if lastPathIndex == 0 {
					return 0
				}
				i = -1
			} else {
				return i
			}
		} else {
			// if i == 1 {
			// 	fmt.Println("lastPathIndex", lastPathIndex)
			// 	fmt.Println("i is :", i)
			// 	fmt.Println("solutions is ", solutions)
			// 	fmt.Println("len(paths[i])+len(solutions[i].ants) is ", len(paths[i])+len(solutions[i].ants))
			// 	fmt.Println("len(paths[i+1])+len(solutions[i+1].ants)", len(paths[i+1])+len(solutions[i+1].ants))
			// }
			if len(paths[i])+len(solutions[i].ants) > len(paths[i+1])+len(solutions[i+1].ants) {
				continue
			} else {
				return i
			}
		}
	}
	return 1
}

func initAnts(numberOfAnts int, ants []Ant) {
	for i := 0; i < numberOfAnts; i++ {
		ants[i] = Ant{id: i + 1, pathIndex: -1, currentRoomName: ""}
	}
}

func initSolutions(solutions []Solution) {
	for i := 0; i < len(solutions); i++ {
		solutions[i].pathIndex = i
	}
}
