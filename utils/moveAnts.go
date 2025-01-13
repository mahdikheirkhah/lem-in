package utils

import "fmt"

func MoveAnts(solutions []Solution, pathsNames [][]string, rooms []Room, numberOfAnts int, end Room) {
	paths := changeTypeOfPaths(pathsNames, rooms)

	numberOfAntsNotReachedToEnd := numberOfAnts
	turnNumber := 0

	// isFirstTurn := true
	// fmt.Println("paths in room are: ", paths)
	for numberOfAntsNotReachedToEnd > 0 {
		turnNumber++
		fmt.Print("turn ", turnNumber, ": ")
		// if isFirstTurn {
		// 	for _, solution := range solutions {
		// 		// fmt.Println("solution pathIndex", solution.pathIndex)
		// 		// fmt.Println("solution pathIndex position 0", paths[solution.pathIndex][0])
		// 		// fmt.Println("solution pathIndex position 0 Name", paths[solution.pathIndex][0].Name)
		// 		stepForward(&solution.ants[0], solution, paths, &numberOfAntsNotReachedToEnd, isFirstTurn)
		// 	}
		// 	isFirstTurn = false
		// } else {
		for _, solution := range solutions {
			for antIndex, ant := range solution.ants {
				stepForward(&solution.ants[antIndex], solution, paths, end, &numberOfAntsNotReachedToEnd)
				if ant.currentRoomName == "" {
					break
				}
			}

		}
		// }
		fmt.Println()
		// fmt.Println("number of ants: ", numberOfAntsNotReachedToEnd)
	}
}

func changeTypeOfPaths(paths [][]string, rooms []Room) [][]Room {
	var output = make([][]Room, len(paths))
	for pathIndex, path := range paths {
		var eachPath []Room
		for _, roomName := range path {
			roomIndex := FindRoom(roomName, rooms)
			eachPath = append(eachPath, rooms[roomIndex])
		}
		output[pathIndex] = eachPath
	}

	return output
}

func stepForward(ant *Ant, solution Solution, paths [][]Room, end Room, numberOfAntsNotReachedToEnd *int) {
	antCurrentRoomName := ant.currentRoomName
	if antCurrentRoomName != end.Name {
		for i, v := range paths[solution.pathIndex] {
			if antCurrentRoomName == "" {
				ant.currentRoomName = paths[solution.pathIndex][0].Name
				break
			} else if v.Name == antCurrentRoomName {
				ant.currentRoomName = paths[solution.pathIndex][i+1].Name
				break
			}
		}
		if !ant.hasReachedTheEnd {
			fmt.Print("L", ant.id, "-", ant.currentRoomName, " ")
			if ant.currentRoomName == end.Name {
				ant.hasReachedTheEnd = true
				*numberOfAntsNotReachedToEnd--
			}
		}
	}
}
