package main

import (
	"LemIn/utils"
)

func main() {
	fileName := utils.ReadFromCommandLine()
	utils.Lem_in(fileName)
}
