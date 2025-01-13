package utils

import (
	"LemIn/errorHandler"
	"errors"
	"strings"
)

func MakeTunnel(rowData string, rooms []Room) Tunnel {
	rowDataSplited := strings.Split(rowData, "-")

	if len(rowDataSplited) != 2 {
		errorHandler.CheckError(errors.New("ERROR: invalid data format, invalid tunnel format"), true)
		return Tunnel{}
	}

	firstRoomIndex := FindRoom(rowDataSplited[0], rooms)
	secondRoomIndex := FindRoom(rowDataSplited[1], rooms)

	if secondRoomIndex == -1 || firstRoomIndex == -1 {
		errorHandler.CheckError(errors.New("ERROR: invalid data format, invalid tunnel format"), true)
		return Tunnel{}
	}

	return Tunnel{
		FromRoom: rooms[firstRoomIndex],
		ToRoom:   rooms[secondRoomIndex],
	}
}

func FindRoom(roomName string, rooms []Room) int {
	for i, room := range rooms {
		if room.Name == roomName {
			return i
		}
	}

	return -1
}

func FindStart(rooms []Room) (int, Room) {
	for i, room := range rooms {
		if room.IsStart {
			return i, room
		}
	}

	return -1, Room{}
}

func FindEnd(rooms []Room) (int, Room) {
	for i, room := range rooms {
		if room.IsEnd {
			return i, room
		}
	}

	return -1, Room{}
}
