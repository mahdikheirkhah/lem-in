package utils

import "fmt"

func MoveAnts(solutions []Solution, pathsNames [][]string, rooms []Room, numberOfAnts int, end Room) {
	paths := changeTypeOfPaths(pathsNames, rooms)

	numberOfAntsNotReachedToEnd := numberOfAnts
	turnNumber := 0

	for numberOfAntsNotReachedToEnd > 0 {
		turnNumber++
		fmt.Print("turn ", turnNumber, ": ")

		for _, solution := range solutions {
			for antIndex, ant := range solution.Ants {
				stepForward(&solution.Ants[antIndex], solution, paths, end, &numberOfAntsNotReachedToEnd)
				if ant.CurrentRoomName == "" {
					break
				}
			}

		}
		fmt.Println()
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
	bgYellow := "\033[43m"
	reset := "\033[0m"

	antCurrentRoomName := ant.CurrentRoomName
	if antCurrentRoomName != end.Name {
		for i, v := range paths[solution.PathIndex] {
			if antCurrentRoomName == "" {
				ant.CurrentRoomName = paths[solution.PathIndex][0].Name
				break
			} else if v.Name == antCurrentRoomName {
				ant.CurrentRoomName = paths[solution.PathIndex][i+1].Name
				break
			}
		}
		if !ant.HasReachedTheEnd {
			if ant.CurrentRoomName == end.Name {
				fmt.Print(bgYellow, "L", ant.Id, "-", ant.CurrentRoomName, reset, " ")
				ant.HasReachedTheEnd = true
				*numberOfAntsNotReachedToEnd--
			} else {
				fmt.Print("L", ant.Id, "-", ant.CurrentRoomName, " ")
			}
		}
	}
}
