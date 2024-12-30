package errorHandler

import (
	"log"
	"os"
)

func CheckError(err error, ExitFlag bool) {
	if err != nil {
		log.Println("Error", err)
		if ExitFlag {
			os.Exit(1)
		}
	}
}
