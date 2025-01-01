package utils

import (
	"LemIn/errorHandler"
	"errors"
	"strconv"
	"strings"
)

func MakeRoom(rowData string) Room {
	rowDataSplited := strings.Split(rowData, " ")
	if len(rowDataSplited) != 3 {
		errorHandler.CheckError(errors.New("ERROR: invalid data format, invalid room format"), true)
		return Room{}
	}
	roomName := rowDataSplited[0]
	if strings.HasPrefix(roomName, "#") || strings.HasPrefix(roomName, "L") {
		errorHandler.CheckError(errors.New("ERROR: invalid data format, invalid room format"), true)
		return Room{}
	}

	coord_x, err_x := strconv.Atoi(rowDataSplited[1])
	coord_y, err_y := strconv.Atoi(rowDataSplited[2])
	if err_x != nil || err_y != nil {
		errorHandler.CheckError(errors.New("ERROR: invalid data format, invalid room format"), true)
		return Room{}
	}
	return Room{
		Name:    roomName,
		Coord_x: coord_x,
		Coord_y: coord_y,
	}
}
