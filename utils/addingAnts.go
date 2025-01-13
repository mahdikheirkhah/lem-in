package utils

func MakeAntsQueue(paths [][]string, numberOfAnts int) []Solution {
	numberOfPaths := len(paths)

	ants := make([]Ant, numberOfAnts)
	initAnts(numberOfAnts, ants)

	solutions := make([]Solution, numberOfPaths)
	initSolutions(solutions)

	pathIndex := 0

	for antsIndex := 0; antsIndex < numberOfAnts; antsIndex++ {
		pathIndex = findSuitablePath(pathIndex, paths, solutions)

		ants[antsIndex].pathIndex = pathIndex
		solutions[pathIndex].ants = append(solutions[pathIndex].ants, ants[antsIndex])
	}

	return solutions
}

func findSuitablePath(lastPathIndex int, paths [][]string, solutions []Solution) int {
	cntr := 0
	for i := lastPathIndex; cntr < len(solutions); i++ {
		cntr++
		if i != len(solutions)-1 {
			if len(paths[i])+len(solutions[i].ants) > len(paths[i+1])+len(solutions[i+1].ants) {
				continue
			} else {
				return i
			}
		} else {
			if len(paths[i])+len(solutions[i].ants) > len(paths[0])+len(solutions[0].ants) {
				if lastPathIndex == 0 {
					return 0
				}
				i = -1
			} else {
				return i
			}
		}
	}
	return 0
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
