package utils

import (
	"LemIn/errorHandler"
	"errors"
	"strconv"
	"strings"
)

func CheckContent(fileContent []string) (int, []Room, []Tunnel, Room, Room) {
	var numberOfAnts int
	var rooms []Room
	var tunnels []Tunnel
	var start Room
	var end Room
	var err error

	if len(fileContent) < 6 {
		errorHandler.CheckError(errors.New("ERROR: invalid data format"), true)
		return -1,nil,nil,Room{},Room{}
	}
	fileContent, start, end = ExtractComments(fileContent)
	if fileContent == nil {
		return -1,nil,nil,Room{},Room{}
	}
	numberOfAnts, err = strconv.Atoi(fileContent[0])
	if err != nil || numberOfAnts < 1 {
		errorHandler.CheckError(errors.New("ERROR: invalid data format, invalid number of Ants"), true)
		return -1,nil,nil,Room{},Room{}
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
				return -1,nil,nil,Room{},Room{}
			}
		}
		rooms = append(rooms, MakeRoom(fileContent[i]))
	}
	if len(rooms) == 0 {
		errorHandler.CheckError(errors.New("ERROR: invalid data format"), true)
		return -1,nil,nil,Room{},Room{}
	}
	// Tunnels should be after the defination of rooms
	for i := index; i < size; i++ {
		if !IsTunnel(fileContent[i]) {
			errorHandler.CheckError(errors.New("ERROR: invalid data format"), true)
			return -1,nil,nil,Room{},Room{}
		}
		tunnels = append(tunnels, MakeTunnel(fileContent[i], rooms))
	}
	if len(tunnels) == 0 {
		errorHandler.CheckError(errors.New("ERROR: invalid data format"), true)
		return -1,nil,nil,Room{},Room{}
	}
	return numberOfAnts, rooms, tunnels, start, end
}

func ExtractComments(fileContent []string) ([]string, Room, Room) {
	var modifiedContent []string
	var start Room
	var end Room
	size := len(fileContent)
	startFlag := false
	endFlag := false
	for i := 0; i < size; i++ {
		if strings.ToLower(fileContent[i]) == "##start" {
			startFlag = true
			if i == size-1 {
				errorHandler.CheckError(errors.New("ERROR: invalid data format, no start room found"), true)
				return nil, Room{}, Room{}
			}
			start = MakeRoom(fileContent[i+1])
		} else if strings.ToLower(fileContent[i]) == "##end" {
			endFlag = true
			if i == size-1 {
				errorHandler.CheckError(errors.New("ERROR: invalid data format, no end room found"), true)
				return nil, Room{}, Room{}
			}
			end = MakeRoom(fileContent[i+1])
		} else if !strings.HasPrefix(fileContent[i], "#") {
			modifiedContent = append(modifiedContent, fileContent[i])
		}
	}
	if !startFlag{
		errorHandler.CheckError(errors.New("ERROR: invalid data format, no start room found"), true)
		return nil, Room{}, Room{}
	} else if !endFlag {
		errorHandler.CheckError(errors.New("ERROR: invalid data format, no end room found"), true)
		return nil, Room{}, Room{}
	}

	return modifiedContent, start, end
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
