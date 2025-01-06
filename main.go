package main

import (
	"LemIn/fileHandler"
	"LemIn/utils"
	"fmt"
)

func main() {
	fileName := utils.ReadFromCommandLine()
	fileContent := fileHandler.ReadAll(fileName)
	for i := 0; i < len(fileContent); i++ {
		fmt.Println("\"", fileContent[i], "\"", ",")
	}
	numberOfAnts, rooms, tunnels := utils.CheckContent(fileContent)
	graph := utils.CreateGraph(tunnels)
	graph.Vertices = len(rooms)
	for i := 0; i < len(rooms); i++ {
		fmt.Println("{")
		fmt.Println("Name:", rooms[i].Name, ",")
		fmt.Println("Coord_x:", rooms[i].Coord_x, ",")
		fmt.Println("Coord_y:", rooms[i].Coord_y, ",")
		fmt.Println("IsStart:", rooms[i].IsStart, ",")
		fmt.Println("IsEnd:", rooms[i].IsEnd, ",")
		fmt.Println("},")
	}
	for i := 0; i < len(tunnels); i++ {
		fmt.Println("{")
		fmt.Println("FromRoom: utils.Room{")
		fmt.Println("Name:\"", tunnels[i].FromRoom.Name, "\",")
		fmt.Println("Coord_x:", tunnels[i].FromRoom.Coord_x, ",")
		fmt.Println("Coord_y:", tunnels[i].FromRoom.Coord_y, ",")
		fmt.Println("},")
		fmt.Println("ToRoom: utils.Room{")
		fmt.Println("Name:\"", tunnels[i].ToRoom.Name, "\",")
		fmt.Println("Coord_x:", tunnels[i].ToRoom.Coord_x, ",")
		fmt.Println("Coord_y:", tunnels[i].ToRoom.Coord_y, ",")
		fmt.Println("},")
		fmt.Println("},")
	}
	fmt.Println("numberOfAnts: ", numberOfAnts)
	fmt.Println("Graph Edges:\n", graph.Edges)
	fmt.Println("Graph Vertices:", graph.Vertices)

	_, startRoom := utils.FindStart(rooms)
	_, endRoom := utils.FindEnd(rooms)
	maxFlow, paths := graph.EdmondsKarp(startRoom, endRoom, rooms)
	fmt.Println("max flow is: ", maxFlow)
	fmt.Println("path is: ", paths)
	utils.MakeAntsQueue(paths, numberOfAnts)
}
