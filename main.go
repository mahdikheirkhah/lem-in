package main

import (
	"LemIn/fileHandler"
	"LemIn/utils"
	"fmt"
)

func main() {
	fileName := utils.ReadFromCommandLine()
	fileContent := fileHandler.ReadAll(fileName)
	numberOfAnts, rooms, tunnels, start, end := utils.CheckContent(fileContent)
	graph := utils.CreateGraph(tunnels)
	fmt.Println(numberOfAnts, rooms, start, end)
	fmt.Println("Graph\n:", graph.Edges)
}
