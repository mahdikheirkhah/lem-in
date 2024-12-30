package utils

import (
	"LemIn/errorHandler"
	"errors"
	"os"
)

func ReadFromCommandLine() string {
	args := os.Args[1:]
	if len(args) != 1 {
		errorHandler.CheckError(errors.New("not enough argumnts"), true)
	}
	return args[0]
}
