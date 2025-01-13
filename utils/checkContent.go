package utils

import (
	"LemIn/errorHandler"
	"errors"
	"strconv"
	"strings"
)

func CheckContent(fileContent []string) (int, []Room, []Tunnel) {
	var numberOfAnts int
	var rooms []Room
	var tunnels []Tunnel
	var err error

	if len(fileContent) < 6 {
		errorHandler.CheckError(errors.New("ERROR: invalid data format"), true)
		return -1, nil, nil
	}
	fileContent, rooms = ExtractComments(fileContent, rooms)
	if fileContent == nil {
		return -1, nil, nil
	}
	numberOfAnts, err = strconv.Atoi(fileContent[0])
	if err != nil || numberOfAnts < 1 {
		errorHandler.CheckError(errors.New("ERROR: invalid data format, invalid number of Ants"), true)
		return -1, nil, nil
	}
	size := len(fileContent)
	index := 1

	// fileContent[0] is for number of ants, So will be looped from index one.
	for i := 1; i < size; i++ {
		if !IsRoom(fileContent[i]) {
			if IsTunnel(fileContent[i]) {
				index = i
				break
			} else {
				errorHandler.CheckError(errors.New("ERROR: invalid data format"), true)
				return -1, nil, nil
			}
		}
		rooms = append(rooms, MakeRoom(fileContent[i]))
	}
	if len(rooms) == 0 {
		errorHandler.CheckError(errors.New("ERROR: invalid data format"), true)
		return -1, nil, nil
	}
	// Tunnels should be after the defination of rooms
	for i := index; i < size; i++ {
		if !IsTunnel(fileContent[i]) {
			errorHandler.CheckError(errors.New("ERROR: invalid data format"), true)
			return -1, nil, nil
		}
		tunnels = append(tunnels, MakeTunnel(fileContent[i], rooms))
	}
	if len(tunnels) == 0 {
		errorHandler.CheckError(errors.New("ERROR: invalid data format"), true)
		return -1, nil, nil
	}
	if !checkUniqueName(rooms) {
		errorHandler.CheckError(errors.New("ERROR: invalid data format, invalid room format"), true)
		return -1, nil, nil
	}
	return numberOfAnts, rooms, tunnels
}

func ExtractComments(fileContent []string, rooms []Room) ([]string, []Room) {
	var modifiedContent []string
	var start Room
	var end Room
	size := len(fileContent)
	startFlag := false
	endFlag := false
	for i := 0; i < size; i++ {
		if strings.ToLower(fileContent[i]) == "##start" {
			if startFlag {
				errorHandler.CheckError(errors.New("ERROR: invalid data format, more than one start room found"), true)
				return nil, []Room{}
			}
			startFlag = true
			if i == size-1 {
				errorHandler.CheckError(errors.New("ERROR: invalid data format, no start room found"), true)
				return nil, []Room{}
			}
			start = MakeRoom(fileContent[i+1])
			start.IsStart = true
			rooms = append(rooms, start)
			i++
		} else if strings.ToLower(fileContent[i]) == "##end" {
			if endFlag {
				errorHandler.CheckError(errors.New("ERROR: invalid data format, more than one end room found"), true)
				return nil, []Room{}
			}
			endFlag = true
			if i == size-1 {
				errorHandler.CheckError(errors.New("ERROR: invalid data format, no end room found"), true)
				return nil, []Room{}
			}
			end = MakeRoom(fileContent[i+1])
			end.IsEnd = true
			rooms = append(rooms, end)
			i++
		} else if !strings.HasPrefix(fileContent[i], "#") {
			modifiedContent = append(modifiedContent, fileContent[i])
		}
	}
	if !startFlag {
		errorHandler.CheckError(errors.New("ERROR: invalid data format, no start room found"), true)
		return nil, []Room{}
	} else if !endFlag {
		errorHandler.CheckError(errors.New("ERROR: invalid data format, no end room found"), true)
		return nil, []Room{}
	}

	return modifiedContent, rooms
}

func IsTunnel(line string) bool {
	splittedLine := strings.Split(line, "-")
	if len(splittedLine) != 2 {
		return false
	}
	if len(splittedLine[0]) == 0 || len(splittedLine[1]) == 0 {
		return false
	}
	return true
}

func IsRoom(line string) bool {
	if strings.Contains(line, "-") {
		return false
	}
	splittedLine := strings.Split(line, " ")
	return len(splittedLine) == 3
}

func checkUniqueName(rooms []Room) bool {

	for i := 0; i < len(rooms); i++ {
		for j := i + 1; j < len(rooms); j++ {
			if rooms[i].Name == rooms[j].Name {
				return false
			}
		}
	}
	return true
}
