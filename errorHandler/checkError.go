package errorHandler

import (
	"log"
	"os"
)

var ExitFunc = os.Exit

func CheckError(err error, ExitFlag bool) {
	if err != nil {
		log.Println("Error", err)
		if ExitFlag {
			ExitFunc(1)
		}
	}
}
