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

		ants[antsIndex].PathIndex = pathIndex
		solutions[pathIndex].Ants = append(solutions[pathIndex].Ants, ants[antsIndex])
	}

	return solutions
}

func findSuitablePath(lastPathIndex int, paths [][]string, solutions []Solution) int {
	cntr := 0
	for i := 0; cntr < len(solutions); i++ {
		cntr++
		if i != len(solutions)-1 {
			if len(paths[i])+len(solutions[i].Ants) <= len(paths[i+1])+len(solutions[i+1].Ants) {
				return i
			} else {
				continue
			}
		} else {
			if len(paths[i])+len(solutions[i].Ants) <= len(paths[0])+len(solutions[0].Ants) {
				return i
			} else {
				i = -1
			}
		}
	}
	return 0
}

func initAnts(numberOfAnts int, ants []Ant) {
	for i := 0; i < numberOfAnts; i++ {
		ants[i] = Ant{Id: i + 1, PathIndex: -1, CurrentRoomName: ""}
	}
}

func initSolutions(solutions []Solution) {
	for i := 0; i < len(solutions); i++ {
		solutions[i].PathIndex = i
	}
}
