package utils_test

import (
	"LemIn/errorHandler"
	"LemIn/fileHandler"
	"LemIn/utils"
	"bytes"
	"log"
	"os"
	"regexp"
	"strings"
	"testing"
)

func TestMakeRoom(t *testing.T) {
	calledExit := false

	// Mock os.Exit to prevent the program from exiting during tests
	errorHandler.ExitFunc = func(code int) {
		calledExit = true
	}

	// Restore original os.Exit after tests
	defer func() {
		errorHandler.ExitFunc = os.Exit
	}()

	tests := []struct {
		name          string
		input         string
		expectedRoom  utils.Room
		expectedError string
	}{
		{
			name:  "Valid Room 1",
			input: "1 10 12",
			expectedRoom: utils.Room{
				Name:    "1",
				Coord_x: 10,
				Coord_y: 12,
			},
		},
		{
			name:  "Valid Room 2",
			input: "mahdi 21 3",
			expectedRoom: utils.Room{
				Name:    "mahdi",
				Coord_x: 21,
				Coord_y: 3,
			},
		},
		{
			name:          "Invalid Room 1",
			input:         "10 12", // Missing room name
			expectedError: "ERROR: invalid data format, invalid room format",
		},
		{
			name:          "Invalid Room 2",
			input:         "mahdi10 12", // Invalid room name format
			expectedError: "ERROR: invalid data format, invalid room format",
		},
		{
			name:          "Invalid Room 3",
			input:         "#mahdi 10 12", // Room name starts with #
			expectedError: "ERROR: invalid data format, invalid room format",
		},
		{
			name:          "Invalid Room 4",
			input:         "Lmahdi 10 12", // Room name starts with L
			expectedError: "ERROR: invalid data format, invalid room format",
		},
		{
			name:          "Invalid Room 5",
			input:         "mahdi 10.5 12", // Invalid x-coordinate (float)
			expectedError: "ERROR: invalid data format, invalid room format",
		},
		{
			name:          "Invalid Room 6",
			input:         "mahdi 10 12.2", // Invalid y-coordinate (float)
			expectedError: "ERROR: invalid data format, invalid room format",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Capture console output
			var buf bytes.Buffer
			log.SetOutput(&buf)

			// Reset the calledExit flag
			calledExit = false

			defer func() {
				log.SetOutput(os.Stderr) // Restore original output
			}()

			if test.expectedError == "" {
				// Test valid rooms
				gotRoom := utils.MakeRoom(test.input)
				if gotRoom != test.expectedRoom {
					t.Errorf("Expected %v but got %v", test.expectedRoom, gotRoom)
				}
			} else {
				// Test invalid rooms
				utils.MakeRoom(test.input)

				// Check if exit was called
				if !calledExit {
					t.Errorf("Expected program to exit, but it did not")
				}

				// Validate captured error message
				output := buf.String()
				t.Logf("Captured Output: '%s'", output) // Log the captured output for debugging

				if !strings.Contains(output, test.expectedError) {
					t.Errorf("Expected output to contain '%s', got '%s'", test.expectedError, output)
				}
			}
		})
	}
}

func TestMakeTunnel(t *testing.T) {
	calledExit := false

	// Mock os.Exit to prevent the program from exiting during tests
	errorHandler.ExitFunc = func(code int) {
		calledExit = true
	}

	// Restore original os.Exit after tests
	defer func() {
		errorHandler.ExitFunc = os.Exit
	}()
	rooms := []utils.Room{
		{
			Name:    "1",
			Coord_x: 23,
			Coord_y: 3,
		},
		{
			Name:    "2",
			Coord_x: 16,
			Coord_y: 7,
		},
		{
			Name:    "0",
			Coord_x: 9,
			Coord_y: 5,
		},
		{
			Name:    "3",
			Coord_x: 16,
			Coord_y: 3,
		}, {
			Name:    "4",
			Coord_x: 16,
			Coord_y: 5,
		}, {
			Name:    "5",
			Coord_x: 9,
			Coord_y: 3,
		},
		{
			Name:    "6",
			Coord_x: 1,
			Coord_y: 5,
		},
		{
			Name:    "7",
			Coord_x: 4,
			Coord_y: 8,
		},
	}
	tests := []struct {
		name  string
		input string
		// rooms         []utils.Room
		expectedTunnel utils.Tunnel
		expectedError  string
	}{
		{
			name:  "Valid tunnel 1",
			input: "1-3",
			expectedTunnel: utils.Tunnel{
				FromRoom: rooms[0],
				ToRoom:   rooms[3],
			},
		},
		{
			name:  "Valid tunnel 2",
			input: "7-2",
			expectedTunnel: utils.Tunnel{
				FromRoom: rooms[7],
				ToRoom:   rooms[1],
			},
		},
		{
			name:          "Invalid Tunnel 1",
			input:         "1 4",
			expectedError: "ERROR: invalid data format, invalid tunnel format",
		},
		{
			name:          "Invalid Tunnel 2",
			input:         "1-12",
			expectedError: "ERROR: invalid data format, invalid tunnel format",
		},
		{
			name:          "Invalid Tunnel 3",
			input:         "1-12-",
			expectedError: "ERROR: invalid data format, invalid tunnel format",
		},
		{
			name:          "Invalid Tunnel 4",
			input:         "112-",
			expectedError: "ERROR: invalid data format, invalid tunnel format",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Capture console output
			var buf bytes.Buffer
			log.SetOutput(&buf)

			// Reset the calledExit flag
			calledExit = false

			defer func() {
				log.SetOutput(os.Stderr) // Restore original output
			}()

			if test.expectedError == "" {
				// Test valid Tunnels
				gotTunnel := utils.MakeTunnel(test.input, rooms)
				if gotTunnel != test.expectedTunnel {
					t.Errorf("Expected %v but got %v", test.expectedTunnel, gotTunnel)
				}
			} else {
				// Test invalid Tunnels
				utils.MakeTunnel(test.input, rooms)

				// Check if exit was called
				if !calledExit {
					t.Errorf("Expected program to exit, but it did not")
				}

				// Validate captured error message
				output := buf.String()
				t.Logf("Captured Output: '%s'", output) // Log the captured output for debugging

				if !strings.Contains(output, test.expectedError) {
					t.Errorf("Expected output to contain '%s', got '%s'", test.expectedError, output)
				}
			}
		})
	}
}

func TestReadAll(t *testing.T) {
	calledExit := false

	// Mock os.Exit to prevent the program from exiting during tests
	errorHandler.ExitFunc = func(code int) {
		calledExit = true
	}

	// Restore original os.Exit after tests
	defer func() {
		errorHandler.ExitFunc = os.Exit
	}()

	tests := []struct {
		name           string
		fileName       string
		expectedOutput []string
		expectedError  string
	}{
		{
			name:     "Valid file",
			fileName: "ValidFile.txt",
			expectedOutput: []string{
				"100",
				"##start",
				"richard 0 6",
				"gilfoyle 6 3",
				"erlich 9 6",
				"dinish 6 9",
				"jimYoung 11 7",
				"#jdjdj",
				"##end",
				"peter 14 6",
				"richard-dinish",
				"dinish-jimYoung",
				"richard-gilfoyle",
				"gilfoyle-peter",
				"gilfoyle-erlich",
				"richard-erlich",
				"erlich-jimYoung",
				"jimYoung-peter",
			},
		},
		{
			name:          "Invalid file",
			fileName:      "invalidFile.txt",
			expectedError: "Error open invalidFile.txt: no such file or directory",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Capture console output
			var buf bytes.Buffer
			log.SetOutput(&buf)

			// Reset the calledExit flag
			calledExit = false

			defer func() {
				log.SetOutput(os.Stderr) // Restore original output
			}()

			if test.expectedError == "" {
				// Test valid rooms
				Output := fileHandler.ReadAll(test.fileName)
				if len(Output) != len(test.expectedOutput) {
					t.Errorf("Wrong outPut")
				}
				for i := 0; i < len(Output); i++ {
					if Output[i] != test.expectedOutput[i] {
						t.Errorf("Expected %v but got %v", test.expectedOutput[i], Output[i])
					}
				}
			} else {
				// Test invalid rooms
				fileHandler.ReadAll(test.fileName)
				// Check if exit was called
				if !calledExit {
					t.Errorf("Expected program to exit, but it did not")
				}

				// Validate captured error message
				output := buf.String()
				t.Logf("Captured Output: '%s'", output) // Log the captured output for debugging

				if !strings.Contains(output, test.expectedError) {
					t.Errorf("Expected output to contain '%s', got '%s'", test.expectedError, output)
				}
			}
		})
	}
}

func TestCheckContant(t *testing.T) {
	calledExit := false

	// Mock os.Exit to prevent the program from exiting during tests
	errorHandler.ExitFunc = func(code int) {
		calledExit = true
	}

	// Restore original os.Exit after tests
	defer func() {
		errorHandler.ExitFunc = os.Exit
	}()

	tests := []struct {
		name                 string
		testFileName         string
		expectedNumberOfAnts int
		expectedRooms        []utils.Room
		expectedTunnels      []utils.Tunnel
		expectedStart        utils.Room
		expectedEnd          utils.Room
		expectedError        string
	}{
		{
			name:                 "Valid input 1",
			testFileName:         "test1.txt",
			expectedNumberOfAnts: 4,
			expectedRooms: []utils.Room{
				{
					Name:        "0",
					Coord_x:     0,
					Coord_y:     3,
					IsStart:     true,
					IsEnd:       false,
					AddedInPath: false,
				},
				{
					Name:        "1",
					Coord_x:     8,
					Coord_y:     3,
					IsStart:     false,
					IsEnd:       true,
					AddedInPath: false,
				},
				{
					Name:        "2",
					Coord_x:     2,
					Coord_y:     5,
					IsStart:     false,
					IsEnd:       false,
					AddedInPath: false,
				},
				{
					Name:        "3",
					Coord_x:     4,
					Coord_y:     0,
					IsStart:     false,
					IsEnd:       false,
					AddedInPath: false,
				},
			},
			expectedTunnels: []utils.Tunnel{
				{
					FromRoom: utils.Room{
						Name:        "0",
						Coord_x:     0,
						Coord_y:     3,
						IsStart:     true,
						IsEnd:       false,
						AddedInPath: false,
					},
					ToRoom: utils.Room{
						Name:        "2",
						Coord_x:     2,
						Coord_y:     5,
						IsStart:     false,
						IsEnd:       false,
						AddedInPath: false,
					},
				},
				{
					FromRoom: utils.Room{
						Name:        "2",
						Coord_x:     2,
						Coord_y:     5,
						IsStart:     false,
						IsEnd:       false,
						AddedInPath: false,
					},
					ToRoom: utils.Room{
						Name:        "3",
						Coord_x:     4,
						Coord_y:     0,
						IsStart:     false,
						IsEnd:       false,
						AddedInPath: false,
					},
				},
				{
					FromRoom: utils.Room{
						Name:        "3",
						Coord_x:     4,
						Coord_y:     0,
						IsStart:     false,
						IsEnd:       false,
						AddedInPath: false,
					},
					ToRoom: utils.Room{
						Name:        "1",
						Coord_x:     8,
						Coord_y:     3,
						IsStart:     false,
						IsEnd:       true,
						AddedInPath: false,
					},
				},
			},
			expectedStart: utils.Room{
				Name:        "0",
				Coord_x:     0,
				Coord_y:     3,
				IsStart:     true,
				IsEnd:       false,
				AddedInPath: false,
			},
			expectedEnd: utils.Room{
				Name:        "1",
				Coord_x:     8,
				Coord_y:     3,
				IsStart:     false,
				IsEnd:       true,
				AddedInPath: false,
			},
		},
		{
			name:          "Invalid input 1",
			testFileName:  "test2.txt",
			expectedError: "ERROR: invalid data format, invalid number of Ants",
		},
		{
			name:          "Invalid input 2",
			testFileName:  "test3.txt",
			expectedError: "ERROR: invalid data format, no start room found",
		},
		{
			name:          "Invalid input 3",
			testFileName:  "test4.txt",
			expectedError: "ERROR: invalid data format, no end room found",
		},
		{
			name:          "Invalid input 4",
			testFileName:  "test5.txt",
			expectedError: "ERROR: invalid data format",
		},
		{
			name:          "Invalid input 5",
			testFileName:  "test6.txt",
			expectedError: "ERROR: invalid data format",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Capture console output
			var buf bytes.Buffer
			log.SetOutput(&buf)

			// Reset the calledExit flag
			calledExit = false

			defer func() {
				log.SetOutput(os.Stderr) // Restore original output
			}()
			fileContent := fileHandler.ReadAll(test.testFileName)
			if test.expectedError == "" {
				// Test valid rooms
				numberOfAnts, rooms, tunnels := utils.CheckContent(fileContent)
				_, start := utils.FindStart(rooms)
				_, end := utils.FindEnd(rooms)
				if numberOfAnts != test.expectedNumberOfAnts {
					t.Errorf("Expected %v but got %v", test.expectedNumberOfAnts, numberOfAnts)
				}

				for i := 0; i < len(rooms); i++ {
					if rooms[i] != test.expectedRooms[i] {
						t.Errorf("Expected %v but got %v", test.expectedRooms[i], rooms[i])
					}
				}
				for i := 0; i < len(tunnels); i++ {
					if tunnels[i] != test.expectedTunnels[i] {
						t.Errorf("Expected %v but got %v", test.expectedTunnels[i], tunnels[i])
					}
				}
				if start != test.expectedStart {
					t.Errorf("Expected %v but got %v", test.expectedStart, start)
				}
				if end != test.expectedEnd {
					t.Errorf("Expected %v but got %v", test.expectedEnd, end)
				}
			} else {
				// Test invalid rooms
				utils.CheckContent(fileContent)
				// Check if exit was called
				if !calledExit {
					t.Errorf("Expected program to exit, but it did not")
				}

				// Validate captured error message
				output := buf.String()
				t.Logf("Captured Output: '%s'", output) // Log the captured output for debugging

				if !strings.Contains(output, test.expectedError) {
					t.Errorf("Expected output to contain '%s', got '%s'", test.expectedError, output)
				}
			}
		})
	}
}

func TestExtractAllPaths(t *testing.T) {
	calledExit := false

	// Mock os.Exit to prevent the program from exiting during tests
	errorHandler.ExitFunc = func(code int) {
		calledExit = true
	}

	// Restore original os.Exit after tests
	defer func() {
		errorHandler.ExitFunc = os.Exit
	}()

	tests := []struct {
		name           string
		graph          utils.Graph
		start          utils.Room
		end            utils.Room
		rooms          []utils.Room
		expectedError  string
		expectedOutput [][]utils.Room
	}{
		{
			name: "Valid test",
			graph: utils.Graph{
				Vertices: 4,
				Edges: map[string][]string{
					"0": {"2"},
					"1": {"3"},
					"2": {"0", "3"},
					"3": {"2", "1"},
				},
			},
			start: utils.Room{Name: "0", Coord_x: 0, Coord_y: 3, IsStart: true, IsEnd: false, AddedInPath: false},
			end:   utils.Room{Name: "1", Coord_x: 8, Coord_y: 3, IsStart: false, IsEnd: true, AddedInPath: false},
			rooms: []utils.Room{
				{Name: "0", Coord_x: 0, Coord_y: 3, IsStart: true, IsEnd: false, AddedInPath: false},
				{Name: "1", Coord_x: 8, Coord_y: 3, IsStart: false, IsEnd: true, AddedInPath: false},
				{Name: "2", Coord_x: 2, Coord_y: 5, IsStart: false, IsEnd: false, AddedInPath: false},
				{Name: "3", Coord_x: 4, Coord_y: 0, IsStart: false, IsEnd: false, AddedInPath: false},
			},
			expectedOutput: [][]utils.Room{
				{
					{Name: "0", Coord_x: 0, Coord_y: 3, IsStart: true, IsEnd: false, AddedInPath: false},
					{Name: "2", Coord_x: 2, Coord_y: 5, IsStart: false, IsEnd: false, AddedInPath: false},
					{Name: "3", Coord_x: 4, Coord_y: 0, IsStart: false, IsEnd: false, AddedInPath: false},
					{Name: "1", Coord_x: 8, Coord_y: 3, IsStart: false, IsEnd: true, AddedInPath: false},
				},
			},
		},
		{
			name: "Invalid test",
			graph: utils.Graph{
				Vertices: 17,
				Edges: map[string][]string{
					"0":  {"1", "5", "9", "10"},
					"1":  {"0", "2", "11", "15"},
					"10": {"0", "16"},
					"11": {"1", "12"},
					"12": {"11", "13"},
					"13": {"12", "14"},
					"14": {"13", "15"},
					"15": {"14", "1"},
					"16": {"10", "7"},
					"2":  {"1", "3", "9"},
					"3":  {"2", "3", "3"},
					"5":  {"0", "6"},
				},
			},
			start: utils.Room{Name: "0", Coord_x: 2, Coord_y: 0, IsStart: true, IsEnd: false, AddedInPath: false},
			end:   utils.Room{Name: "4", Coord_x: 23, Coord_y: 0, IsStart: false, IsEnd: true, AddedInPath: false},
			rooms: []utils.Room{
				{Name: "0", Coord_x: 2, Coord_y: 0, IsStart: true, IsEnd: false, AddedInPath: false},
				{Name: "4", Coord_x: 23, Coord_y: 0, IsStart: false, IsEnd: true, AddedInPath: false},
				{Name: "1", Coord_x: 7, Coord_y: 0, IsStart: false, IsEnd: false, AddedInPath: false},
				{Name: "2", Coord_x: 13, Coord_y: 0, IsStart: false, IsEnd: false, AddedInPath: false},
				{Name: "3", Coord_x: 18, Coord_y: 0, IsStart: false, IsEnd: false, AddedInPath: false},
				{Name: "5", Coord_x: 7, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
				{Name: "6", Coord_x: 10, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
				{Name: "7", Coord_x: 13, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
				{Name: "8", Coord_x: 16, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
				{Name: "9", Coord_x: 7, Coord_y: 2, IsStart: false, IsEnd: false, AddedInPath: false},
				{Name: "10", Coord_x: 7, Coord_y: 4, IsStart: false, IsEnd: false, AddedInPath: false},
				{Name: "11", Coord_x: 13, Coord_y: 2, IsStart: false, IsEnd: false, AddedInPath: false},
				{Name: "12", Coord_x: 15, Coord_y: 2, IsStart: false, IsEnd: false, AddedInPath: false},
				{Name: "13", Coord_x: 17, Coord_y: 2, IsStart: false, IsEnd: false, AddedInPath: false},
				{Name: "14", Coord_x: 19, Coord_y: 2, IsStart: false, IsEnd: false, AddedInPath: false},
				{Name: "15", Coord_x: 21, Coord_y: 2, IsStart: false, IsEnd: false, AddedInPath: false},
				{Name: "16", Coord_x: 9, Coord_y: 4, IsStart: false, IsEnd: false, AddedInPath: false},
			},
			expectedError: "Error ERROR: invalid data format, no path found",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Capture console output
			var buf bytes.Buffer
			log.SetOutput(&buf)

			// Reset the calledExit flag
			calledExit = false

			defer func() {
				log.SetOutput(os.Stderr) // Restore original output
			}()

			if test.expectedError == "" {
				// Test valid rooms
				Output := utils.ExtractAllPaths(test.graph, test.start, test.end, test.rooms)
				if len(Output) != len(test.expectedOutput) {
					t.Errorf("Wrong outPut")
				}
				for i := 0; i < len(Output); i++ {
					for j := 0; j < len(Output[i]); j++ {
						if Output[i][j] != test.expectedOutput[i][j] {
							t.Errorf("Expected %v but got %v", test.expectedOutput[i][j], Output[i][j])
						}
					}
				}
			} else {
				// Test invalid rooms
				utils.ExtractAllPaths(test.graph, test.start, test.end, test.rooms)
				// Check if exit was called
				if !calledExit {
					t.Errorf("Expected program to exit, but it did not")
				}

				// Validate captured error message
				output := buf.String()
				t.Logf("Captured Output: '%s'", output) // Log the captured output for debugging

				if !strings.Contains(output, test.expectedError) {
					t.Errorf("Expected output to contain '%s', got '%s'", test.expectedError, output)
				}
			}
		})
	}
}

func TestFilterNonIntersectingGroups(t *testing.T) {
	calledExit := false

	// Mock os.Exit to prevent the program from exiting during tests
	errorHandler.ExitFunc = func(code int) {
		calledExit = true
	}

	// Restore original os.Exit after tests
	defer func() {
		errorHandler.ExitFunc = os.Exit
	}()

	tests := []struct {
		name           string
		allPaths       [][]utils.Room
		expectedError  string
		expectedOutput [][][]utils.Room
	}{
		{
			name: "Valid test",
			allPaths: [][]utils.Room{
				{
					{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
					{Name: "h", Coord_x: 4, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
					{Name: "n", Coord_x: 6, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
					{Name: "e", Coord_x: 8, Coord_y: 4, IsStart: false, IsEnd: false, AddedInPath: false},
					{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
				},
				{
					{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
					{Name: "h", Coord_x: 4, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
					{Name: "n", Coord_x: 6, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
					{Name: "m", Coord_x: 8, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
					{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
				},
				{
					{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
					{Name: "t", Coord_x: 1, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
					{Name: "E", Coord_x: 5, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
					{Name: "a", Coord_x: 8, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
					{Name: "m", Coord_x: 8, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
					{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
				},
				{
					{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
					{Name: "h", Coord_x: 4, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
					{Name: "A", Coord_x: 5, Coord_y: 2, IsStart: false, IsEnd: false, AddedInPath: false},
					{Name: "c", Coord_x: 8, Coord_y: 1, IsStart: false, IsEnd: false, AddedInPath: false},
					{Name: "k", Coord_x: 11, Coord_y: 2, IsStart: false, IsEnd: false, AddedInPath: false},
					{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
				},
				{
					{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
					{Name: "0", Coord_x: 4, Coord_y: 8, IsStart: false, IsEnd: false, AddedInPath: false},
					{Name: "o", Coord_x: 6, Coord_y: 8, IsStart: false, IsEnd: false, AddedInPath: false},
					{Name: "n", Coord_x: 6, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
					{Name: "e", Coord_x: 8, Coord_y: 4, IsStart: false, IsEnd: false, AddedInPath: false},
					{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
				},
				{
					{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
					{Name: "0", Coord_x: 4, Coord_y: 8, IsStart: false, IsEnd: false, AddedInPath: false},
					{Name: "o", Coord_x: 6, Coord_y: 8, IsStart: false, IsEnd: false, AddedInPath: false},
					{Name: "n", Coord_x: 6, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
					{Name: "m", Coord_x: 8, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
					{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
				},
				{
					{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
					{Name: "t", Coord_x: 1, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
					{Name: "E", Coord_x: 5, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
					{Name: "a", Coord_x: 8, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
					{Name: "m", Coord_x: 8, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
					{Name: "n", Coord_x: 6, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
					{Name: "e", Coord_x: 8, Coord_y: 4, IsStart: false, IsEnd: false, AddedInPath: false},
					{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
				},
				{
					{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
					{Name: "0", Coord_x: 4, Coord_y: 8, IsStart: false, IsEnd: false, AddedInPath: false},
					{Name: "o", Coord_x: 6, Coord_y: 8, IsStart: false, IsEnd: false, AddedInPath: false},
					{Name: "n", Coord_x: 6, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
					{Name: "h", Coord_x: 4, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
					{Name: "A", Coord_x: 5, Coord_y: 2, IsStart: false, IsEnd: false, AddedInPath: false},
					{Name: "c", Coord_x: 8, Coord_y: 1, IsStart: false, IsEnd: false, AddedInPath: false},
					{Name: "k", Coord_x: 11, Coord_y: 2, IsStart: false, IsEnd: false, AddedInPath: false},
					{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
				},
				{
					{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
					{Name: "t", Coord_x: 1, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
					{Name: "E", Coord_x: 5, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
					{Name: "a", Coord_x: 8, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
					{Name: "m", Coord_x: 8, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
					{Name: "n", Coord_x: 6, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
					{Name: "h", Coord_x: 4, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
					{Name: "A", Coord_x: 5, Coord_y: 2, IsStart: false, IsEnd: false, AddedInPath: false},
					{Name: "c", Coord_x: 8, Coord_y: 1, IsStart: false, IsEnd: false, AddedInPath: false},
					{Name: "k", Coord_x: 11, Coord_y: 2, IsStart: false, IsEnd: false, AddedInPath: false},
					{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
				},
			},
			expectedOutput: [][][]utils.Room{
				{
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "h", Coord_x: 4, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "n", Coord_x: 6, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "e", Coord_x: 8, Coord_y: 4, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
				},
				{
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "h", Coord_x: 4, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "n", Coord_x: 6, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "e", Coord_x: 8, Coord_y: 4, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "t", Coord_x: 1, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "E", Coord_x: 5, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "a", Coord_x: 8, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "m", Coord_x: 8, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
				},
				{
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "h", Coord_x: 4, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "n", Coord_x: 6, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "m", Coord_x: 8, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
				},
				{
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "t", Coord_x: 1, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "E", Coord_x: 5, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "a", Coord_x: 8, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "m", Coord_x: 8, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
				},
				{
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "t", Coord_x: 1, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "E", Coord_x: 5, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "a", Coord_x: 8, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "m", Coord_x: 8, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "h", Coord_x: 4, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "A", Coord_x: 5, Coord_y: 2, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "c", Coord_x: 8, Coord_y: 1, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "k", Coord_x: 11, Coord_y: 2, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
				},
				{
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "t", Coord_x: 1, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "E", Coord_x: 5, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "a", Coord_x: 8, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "m", Coord_x: 8, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "h", Coord_x: 4, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "A", Coord_x: 5, Coord_y: 2, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "c", Coord_x: 8, Coord_y: 1, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "k", Coord_x: 11, Coord_y: 2, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "0", Coord_x: 4, Coord_y: 8, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "o", Coord_x: 6, Coord_y: 8, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "n", Coord_x: 6, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "e", Coord_x: 8, Coord_y: 4, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
				},
				{
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "t", Coord_x: 1, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "E", Coord_x: 5, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "a", Coord_x: 8, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "m", Coord_x: 8, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "0", Coord_x: 4, Coord_y: 8, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "o", Coord_x: 6, Coord_y: 8, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "n", Coord_x: 6, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "e", Coord_x: 8, Coord_y: 4, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
				},
				{
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "t", Coord_x: 1, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "E", Coord_x: 5, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "a", Coord_x: 8, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "m", Coord_x: 8, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "0", Coord_x: 4, Coord_y: 8, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "o", Coord_x: 6, Coord_y: 8, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "n", Coord_x: 6, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "h", Coord_x: 4, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "A", Coord_x: 5, Coord_y: 2, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "c", Coord_x: 8, Coord_y: 1, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "k", Coord_x: 11, Coord_y: 2, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
				},
				{
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "h", Coord_x: 4, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "A", Coord_x: 5, Coord_y: 2, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "c", Coord_x: 8, Coord_y: 1, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "k", Coord_x: 11, Coord_y: 2, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
				},
				{
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "h", Coord_x: 4, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "A", Coord_x: 5, Coord_y: 2, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "c", Coord_x: 8, Coord_y: 1, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "k", Coord_x: 11, Coord_y: 2, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "0", Coord_x: 4, Coord_y: 8, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "o", Coord_x: 6, Coord_y: 8, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "n", Coord_x: 6, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "e", Coord_x: 8, Coord_y: 4, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
				},
				{
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "h", Coord_x: 4, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "A", Coord_x: 5, Coord_y: 2, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "c", Coord_x: 8, Coord_y: 1, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "k", Coord_x: 11, Coord_y: 2, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "0", Coord_x: 4, Coord_y: 8, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "o", Coord_x: 6, Coord_y: 8, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "n", Coord_x: 6, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "m", Coord_x: 8, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
				},
				{
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "h", Coord_x: 4, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "A", Coord_x: 5, Coord_y: 2, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "c", Coord_x: 8, Coord_y: 1, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "k", Coord_x: 11, Coord_y: 2, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "t", Coord_x: 1, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "E", Coord_x: 5, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "a", Coord_x: 8, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "m", Coord_x: 8, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "n", Coord_x: 6, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "e", Coord_x: 8, Coord_y: 4, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
				},
				{
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "0", Coord_x: 4, Coord_y: 8, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "o", Coord_x: 6, Coord_y: 8, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "n", Coord_x: 6, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "e", Coord_x: 8, Coord_y: 4, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
				},
				{
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "0", Coord_x: 4, Coord_y: 8, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "o", Coord_x: 6, Coord_y: 8, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "n", Coord_x: 6, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "m", Coord_x: 8, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
				},
				{
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "t", Coord_x: 1, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "E", Coord_x: 5, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "a", Coord_x: 8, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "m", Coord_x: 8, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "n", Coord_x: 6, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "e", Coord_x: 8, Coord_y: 4, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
				},
				{
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "0", Coord_x: 4, Coord_y: 8, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "o", Coord_x: 6, Coord_y: 8, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "n", Coord_x: 6, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "h", Coord_x: 4, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "A", Coord_x: 5, Coord_y: 2, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "c", Coord_x: 8, Coord_y: 1, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "k", Coord_x: 11, Coord_y: 2, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
				},
				{
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "t", Coord_x: 1, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "E", Coord_x: 5, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "a", Coord_x: 8, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "m", Coord_x: 8, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "n", Coord_x: 6, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "h", Coord_x: 4, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "A", Coord_x: 5, Coord_y: 2, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "c", Coord_x: 8, Coord_y: 1, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "k", Coord_x: 11, Coord_y: 2, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Capture console output
			var buf bytes.Buffer
			log.SetOutput(&buf)

			// Reset the calledExit flag
			calledExit = false

			defer func() {
				log.SetOutput(os.Stderr) // Restore original output
			}()

			if test.expectedError == "" {
				// Test valid rooms
				Output := utils.FilterNonIntersectingGroups(test.allPaths)
				if len(Output) != len(test.expectedOutput) {
					t.Errorf("Wrong outPut")
				}
				for i := 0; i < len(Output); i++ {
					for j := 0; j < len(Output[i]); j++ {
						for k := 0; k < len(Output[i][j]); k++ {
							if Output[i][j][k] != test.expectedOutput[i][j][k] {
								t.Errorf("Expected %v but got %v", test.expectedOutput[i][j][k], Output[i][j][k])
							}
						}

					}
				}
			} else {
				// Test invalid rooms
				utils.FilterNonIntersectingGroups(test.allPaths)
				// Check if exit was called
				if !calledExit {
					t.Errorf("Expected program to exit, but it did not")
				}

				// Validate captured error message
				output := buf.String()
				t.Logf("Captured Output: '%s'", output) // Log the captured output for debugging

				if !strings.Contains(output, test.expectedError) {
					t.Errorf("Expected output to contain '%s', got '%s'", test.expectedError, output)
				}
			}
		})
	}
}

func TestRemoveSmallerGroups(t *testing.T) {
	calledExit := false

	// Mock os.Exit to prevent the program from exiting during tests
	errorHandler.ExitFunc = func(code int) {
		calledExit = true
	}

	// Restore original os.Exit after tests
	defer func() {
		errorHandler.ExitFunc = os.Exit
	}()

	tests := []struct {
		name                  string
		nonIntersectingGroups [][][]utils.Room
		expectedError         string
		expectedOutput        [][][]utils.Room
	}{
		{
			name: "Valid test",
			nonIntersectingGroups: [][][]utils.Room{
				{
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "h", Coord_x: 4, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "n", Coord_x: 6, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "e", Coord_x: 8, Coord_y: 4, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
				},
				{
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "h", Coord_x: 4, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "n", Coord_x: 6, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "e", Coord_x: 8, Coord_y: 4, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "t", Coord_x: 1, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "E", Coord_x: 5, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "a", Coord_x: 8, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "m", Coord_x: 8, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
				},
				{
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "h", Coord_x: 4, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "n", Coord_x: 6, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "m", Coord_x: 8, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
				},
				{
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "t", Coord_x: 1, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "E", Coord_x: 5, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "a", Coord_x: 8, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "m", Coord_x: 8, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
				},
				{
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "t", Coord_x: 1, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "E", Coord_x: 5, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "a", Coord_x: 8, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "m", Coord_x: 8, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "h", Coord_x: 4, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "A", Coord_x: 5, Coord_y: 2, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "c", Coord_x: 8, Coord_y: 1, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "k", Coord_x: 11, Coord_y: 2, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
				},
				{
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "t", Coord_x: 1, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "E", Coord_x: 5, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "a", Coord_x: 8, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "m", Coord_x: 8, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "h", Coord_x: 4, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "A", Coord_x: 5, Coord_y: 2, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "c", Coord_x: 8, Coord_y: 1, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "k", Coord_x: 11, Coord_y: 2, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "0", Coord_x: 4, Coord_y: 8, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "o", Coord_x: 6, Coord_y: 8, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "n", Coord_x: 6, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "e", Coord_x: 8, Coord_y: 4, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
				},
				{
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "t", Coord_x: 1, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "E", Coord_x: 5, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "a", Coord_x: 8, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "m", Coord_x: 8, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "0", Coord_x: 4, Coord_y: 8, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "o", Coord_x: 6, Coord_y: 8, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "n", Coord_x: 6, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "e", Coord_x: 8, Coord_y: 4, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
				},
				{
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "t", Coord_x: 1, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "E", Coord_x: 5, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "a", Coord_x: 8, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "m", Coord_x: 8, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "0", Coord_x: 4, Coord_y: 8, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "o", Coord_x: 6, Coord_y: 8, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "n", Coord_x: 6, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "h", Coord_x: 4, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "A", Coord_x: 5, Coord_y: 2, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "c", Coord_x: 8, Coord_y: 1, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "k", Coord_x: 11, Coord_y: 2, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
				},
				{
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "h", Coord_x: 4, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "A", Coord_x: 5, Coord_y: 2, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "c", Coord_x: 8, Coord_y: 1, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "k", Coord_x: 11, Coord_y: 2, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
				},
				{
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "h", Coord_x: 4, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "A", Coord_x: 5, Coord_y: 2, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "c", Coord_x: 8, Coord_y: 1, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "k", Coord_x: 11, Coord_y: 2, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "0", Coord_x: 4, Coord_y: 8, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "o", Coord_x: 6, Coord_y: 8, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "n", Coord_x: 6, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "e", Coord_x: 8, Coord_y: 4, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
				},
				{
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "h", Coord_x: 4, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "A", Coord_x: 5, Coord_y: 2, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "c", Coord_x: 8, Coord_y: 1, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "k", Coord_x: 11, Coord_y: 2, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "0", Coord_x: 4, Coord_y: 8, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "o", Coord_x: 6, Coord_y: 8, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "n", Coord_x: 6, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "m", Coord_x: 8, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
				},
				{
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "h", Coord_x: 4, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "A", Coord_x: 5, Coord_y: 2, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "c", Coord_x: 8, Coord_y: 1, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "k", Coord_x: 11, Coord_y: 2, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "t", Coord_x: 1, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "E", Coord_x: 5, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "a", Coord_x: 8, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "m", Coord_x: 8, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "n", Coord_x: 6, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "e", Coord_x: 8, Coord_y: 4, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
				},
				{
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "0", Coord_x: 4, Coord_y: 8, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "o", Coord_x: 6, Coord_y: 8, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "n", Coord_x: 6, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "e", Coord_x: 8, Coord_y: 4, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
				},
				{
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "0", Coord_x: 4, Coord_y: 8, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "o", Coord_x: 6, Coord_y: 8, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "n", Coord_x: 6, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "m", Coord_x: 8, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
				},
				{
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "t", Coord_x: 1, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "E", Coord_x: 5, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "a", Coord_x: 8, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "m", Coord_x: 8, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "n", Coord_x: 6, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "e", Coord_x: 8, Coord_y: 4, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
				},
				{
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "0", Coord_x: 4, Coord_y: 8, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "o", Coord_x: 6, Coord_y: 8, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "n", Coord_x: 6, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "h", Coord_x: 4, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "A", Coord_x: 5, Coord_y: 2, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "c", Coord_x: 8, Coord_y: 1, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "k", Coord_x: 11, Coord_y: 2, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
				},
				{
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "t", Coord_x: 1, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "E", Coord_x: 5, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "a", Coord_x: 8, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "m", Coord_x: 8, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "n", Coord_x: 6, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "h", Coord_x: 4, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "A", Coord_x: 5, Coord_y: 2, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "c", Coord_x: 8, Coord_y: 1, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "k", Coord_x: 11, Coord_y: 2, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
				},
			},
			expectedOutput: [][][]utils.Room{
				{
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "h", Coord_x: 4, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "n", Coord_x: 6, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "e", Coord_x: 8, Coord_y: 4, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "t", Coord_x: 1, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "E", Coord_x: 5, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "a", Coord_x: 8, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "m", Coord_x: 8, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
				},
				{
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "h", Coord_x: 4, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "n", Coord_x: 6, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "m", Coord_x: 8, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
				},
				{
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "t", Coord_x: 1, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "E", Coord_x: 5, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "a", Coord_x: 8, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "m", Coord_x: 8, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "h", Coord_x: 4, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "A", Coord_x: 5, Coord_y: 2, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "c", Coord_x: 8, Coord_y: 1, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "k", Coord_x: 11, Coord_y: 2, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "0", Coord_x: 4, Coord_y: 8, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "o", Coord_x: 6, Coord_y: 8, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "n", Coord_x: 6, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "e", Coord_x: 8, Coord_y: 4, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
				},
				{
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "t", Coord_x: 1, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "E", Coord_x: 5, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "a", Coord_x: 8, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "m", Coord_x: 8, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "0", Coord_x: 4, Coord_y: 8, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "o", Coord_x: 6, Coord_y: 8, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "n", Coord_x: 6, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "h", Coord_x: 4, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "A", Coord_x: 5, Coord_y: 2, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "c", Coord_x: 8, Coord_y: 1, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "k", Coord_x: 11, Coord_y: 2, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
				},
				{
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "h", Coord_x: 4, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "A", Coord_x: 5, Coord_y: 2, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "c", Coord_x: 8, Coord_y: 1, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "k", Coord_x: 11, Coord_y: 2, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "0", Coord_x: 4, Coord_y: 8, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "o", Coord_x: 6, Coord_y: 8, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "n", Coord_x: 6, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "m", Coord_x: 8, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
				},
				{
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "h", Coord_x: 4, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "A", Coord_x: 5, Coord_y: 2, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "c", Coord_x: 8, Coord_y: 1, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "k", Coord_x: 11, Coord_y: 2, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "t", Coord_x: 1, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "E", Coord_x: 5, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "a", Coord_x: 8, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "m", Coord_x: 8, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "n", Coord_x: 6, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "e", Coord_x: 8, Coord_y: 4, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
				},
				{
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "t", Coord_x: 1, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "E", Coord_x: 5, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "a", Coord_x: 8, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "m", Coord_x: 8, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "n", Coord_x: 6, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "h", Coord_x: 4, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "A", Coord_x: 5, Coord_y: 2, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "c", Coord_x: 8, Coord_y: 1, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "k", Coord_x: 11, Coord_y: 2, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Capture console output
			var buf bytes.Buffer
			log.SetOutput(&buf)

			// Reset the calledExit flag
			calledExit = false

			defer func() {
				log.SetOutput(os.Stderr) // Restore original output
			}()

			if test.expectedError == "" {
				// Test valid rooms
				Output := utils.RemoveSmallerGroups(test.nonIntersectingGroups)
				if len(Output) != len(test.expectedOutput) {
					t.Errorf("Wrong outPut")
				}
				for i := 0; i < len(Output); i++ {
					for j := 0; j < len(Output[i]); j++ {
						for k := 0; k < len(Output[i][j]); k++ {
							if Output[i][j][k] != test.expectedOutput[i][j][k] {
								t.Errorf("Expected %v but got %v", test.expectedOutput[i][j][k], Output[i][j][k])
							}
						}

					}
				}
			} else {
				// Test invalid rooms
				utils.RemoveSmallerGroups(test.nonIntersectingGroups)
				// Check if exit was called
				if !calledExit {
					t.Errorf("Expected program to exit, but it did not")
				}

				// Validate captured error message
				output := buf.String()
				t.Logf("Captured Output: '%s'", output) // Log the captured output for debugging

				if !strings.Contains(output, test.expectedError) {
					t.Errorf("Expected output to contain '%s', got '%s'", test.expectedError, output)
				}
			}
		})
	}
}

func TestFindBestPathGroup(t *testing.T) {
	calledExit := false

	// Mock os.Exit to prevent the program from exiting during tests
	errorHandler.ExitFunc = func(code int) {
		calledExit = true
	}

	// Restore original os.Exit after tests
	defer func() {
		errorHandler.ExitFunc = os.Exit
	}()

	tests := []struct {
		name           string
		filteredGroups [][][]utils.Room
		numberOfAnts   int
		expectedError  string
		expectedOutput [][]string
	}{
		{
			name: "Valid test",
			filteredGroups: [][][]utils.Room{
				{
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "h", Coord_x: 4, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "n", Coord_x: 6, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "e", Coord_x: 8, Coord_y: 4, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "t", Coord_x: 1, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "E", Coord_x: 5, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "a", Coord_x: 8, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "m", Coord_x: 8, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
				},
				{
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "h", Coord_x: 4, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "n", Coord_x: 6, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "m", Coord_x: 8, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
				},
				{
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "t", Coord_x: 1, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "E", Coord_x: 5, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "a", Coord_x: 8, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "m", Coord_x: 8, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "h", Coord_x: 4, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "A", Coord_x: 5, Coord_y: 2, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "c", Coord_x: 8, Coord_y: 1, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "k", Coord_x: 11, Coord_y: 2, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "0", Coord_x: 4, Coord_y: 8, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "o", Coord_x: 6, Coord_y: 8, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "n", Coord_x: 6, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "e", Coord_x: 8, Coord_y: 4, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
				},
				{
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "t", Coord_x: 1, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "E", Coord_x: 5, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "a", Coord_x: 8, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "m", Coord_x: 8, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "0", Coord_x: 4, Coord_y: 8, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "o", Coord_x: 6, Coord_y: 8, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "n", Coord_x: 6, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "h", Coord_x: 4, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "A", Coord_x: 5, Coord_y: 2, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "c", Coord_x: 8, Coord_y: 1, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "k", Coord_x: 11, Coord_y: 2, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
				},
				{
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "h", Coord_x: 4, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "A", Coord_x: 5, Coord_y: 2, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "c", Coord_x: 8, Coord_y: 1, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "k", Coord_x: 11, Coord_y: 2, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "0", Coord_x: 4, Coord_y: 8, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "o", Coord_x: 6, Coord_y: 8, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "n", Coord_x: 6, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "m", Coord_x: 8, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
				},
				{
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "h", Coord_x: 4, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "A", Coord_x: 5, Coord_y: 2, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "c", Coord_x: 8, Coord_y: 1, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "k", Coord_x: 11, Coord_y: 2, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "t", Coord_x: 1, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "E", Coord_x: 5, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "a", Coord_x: 8, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "m", Coord_x: 8, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "n", Coord_x: 6, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "e", Coord_x: 8, Coord_y: 4, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
				},
				{
					{
						{Name: "start", Coord_x: 1, Coord_y: 6, IsStart: true, IsEnd: false, AddedInPath: false},
						{Name: "t", Coord_x: 1, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "E", Coord_x: 5, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "a", Coord_x: 8, Coord_y: 9, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "m", Coord_x: 8, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "n", Coord_x: 6, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "h", Coord_x: 4, Coord_y: 6, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "A", Coord_x: 5, Coord_y: 2, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "c", Coord_x: 8, Coord_y: 1, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "k", Coord_x: 11, Coord_y: 2, IsStart: false, IsEnd: false, AddedInPath: false},
						{Name: "end", Coord_x: 11, Coord_y: 6, IsStart: false, IsEnd: true, AddedInPath: false},
					},
				},
			},
			numberOfAnts: 10,
			expectedOutput: [][]string{
				{
					"t",
					"E",
					"a",
					"m",
					"end",
				},
				{
					"h",
					"A",
					"c",
					"k",
					"end",
				},
				{
					"0",
					"o",
					"n",
					"e",
					"end",
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Capture console output
			var buf bytes.Buffer
			log.SetOutput(&buf)

			// Reset the calledExit flag
			calledExit = false

			defer func() {
				log.SetOutput(os.Stderr) // Restore original output
			}()

			if test.expectedError == "" {
				// Test valid rooms
				Output := utils.FindBestPathGroup(test.filteredGroups, test.numberOfAnts)
				if len(Output) != len(test.expectedOutput) {
					t.Errorf("Wrong outPut")
				}
				for i := 0; i < len(Output); i++ {
					for j := 0; j < len(Output[i]); j++ {
						if Output[i][j] != test.expectedOutput[i][j] {
							t.Errorf("Expected %v but got %v", test.expectedOutput[i][j], Output[i][j])
						}
					}

				}
			} else {
				// Test invalid rooms
				utils.FindBestPathGroup(test.filteredGroups, test.numberOfAnts)
				// Check if exit was called
				if !calledExit {
					t.Errorf("Expected program to exit, but it did not")
				}

				// Validate captured error message
				output := buf.String()
				t.Logf("Captured Output: '%s'", output) // Log the captured output for debugging

				if !strings.Contains(output, test.expectedError) {
					t.Errorf("Expected output to contain '%s', got '%s'", test.expectedError, output)
				}
			}
		})
	}
}

func TestMakeAntsQueue(t *testing.T) {
	calledExit := false

	// Mock os.Exit to prevent the program from exiting during tests
	errorHandler.ExitFunc = func(code int) {
		calledExit = true
	}

	// Restore original os.Exit after tests
	defer func() {
		errorHandler.ExitFunc = os.Exit
	}()
	tests := []struct {
		name               string
		bestPathGroupNames [][]string
		numberOfAnts       int
		expectedError      string
		expectedOutput     []utils.Solution
	}{
		{
			name:         "Valid test",
			numberOfAnts: 10,
			bestPathGroupNames: [][]string{
				{
					"t",
					"E",
					"a",
					"m",
					"end",
				},
				{
					"h",
					"A",
					"c",
					"k",
					"end",
				},
				{
					"0",
					"o",
					"n",
					"e",
					"end",
				},
			},
			expectedOutput: []utils.Solution{{
				PathIndex: 0,
				Ants: []utils.Ant{
					{Id: 1, PathIndex: 0, CurrentRoomName: "", HasReachedTheEnd: false},
					{Id: 3, PathIndex: 0, CurrentRoomName: "", HasReachedTheEnd: false},
					{Id: 6, PathIndex: 0, CurrentRoomName: "", HasReachedTheEnd: false},
					{Id: 9, PathIndex: 0, CurrentRoomName: "", HasReachedTheEnd: false},
				},
			},
				{
					PathIndex: 1,
					Ants: []utils.Ant{
						{Id: 2, PathIndex: 1, CurrentRoomName: "", HasReachedTheEnd: false},
						{Id: 5, PathIndex: 1, CurrentRoomName: "", HasReachedTheEnd: false},
						{Id: 8, PathIndex: 1, CurrentRoomName: "", HasReachedTheEnd: false},
					},
				},
				{
					PathIndex: 2,
					Ants: []utils.Ant{
						{Id: 4, PathIndex: 2, CurrentRoomName: "", HasReachedTheEnd: false},
						{Id: 7, PathIndex: 2, CurrentRoomName: "", HasReachedTheEnd: false},
						{Id: 10, PathIndex: 2, CurrentRoomName: "", HasReachedTheEnd: false},
					},
				}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Capture console output
			var buf bytes.Buffer
			log.SetOutput(&buf)

			// Reset the calledExit flag
			calledExit = false

			defer func() {
				log.SetOutput(os.Stderr) // Restore original output
			}()

			if test.expectedError == "" {
				// Test valid rooms
				Output := utils.MakeAntsQueue(test.bestPathGroupNames, test.numberOfAnts)
				if len(Output) != len(test.expectedOutput) {
					t.Errorf("Wrong outPut")
				}
				for i := 0; i < len(Output); i++ {
					if Output[i].PathIndex != test.expectedOutput[i].PathIndex {
						t.Errorf("Expected %v but got %v", test.expectedOutput[i].PathIndex, Output[i].PathIndex)
					}
					for j := 0; j < len(Output[i].Ants); j++ {
						if Output[i].Ants[j] != test.expectedOutput[i].Ants[j] {
							t.Errorf("Expected %v but got %v", test.expectedOutput[i].Ants[j], Output[i].Ants[j])
						}
					}

				}
			} else {
				utils.MakeAntsQueue(test.bestPathGroupNames, test.numberOfAnts)
				// Check if exit was called
				if !calledExit {
					t.Errorf("Expected program to exit, but it did not")
				}

				// Validate captured error message
				output := buf.String()
				t.Logf("Captured Output: '%s'", output) // Log the captured output for debugging

				if !strings.Contains(output, test.expectedError) {
					t.Errorf("Expected output to contain '%s', got '%s'", test.expectedError, output)
				}
			}
		})
	}
}

// StripANSI removes ANSI escape sequences from a string.
func StripANSI(input string) string {
	re := regexp.MustCompile(`\x1b\[[0-9;]*[a-zA-Z]`)
	return re.ReplaceAllString(input, "")
}

func TestLem_in(t *testing.T) {
	// Prepare a temporary file with test data
	tests := []struct {
		name           string
		fileName       string
		expectedError  string
		expectedOutput string
	}{
		{
			name:     "Valid Test1",
			fileName: "../examples/example00.txt",
			expectedOutput: `4
##start
0 0 3
2 2 5
3 4 0
##end
1 8 3
0-2
2-3
3-1

turn 1: L1-2 
turn 2: L1-3 L2-2 
turn 3: L1-1 L2-3 L3-2 
turn 4: L2-1 L3-3 L4-2 
turn 5: L3-1 L4-3 
turn 6: L4-1 
`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.expectedError == "" {
				// Redirect stdout to capture printed output
				oldStdout := os.Stdout
				r, w, _ := os.Pipe()
				os.Stdout = w

				// Call the function with the test file
				utils.Lem_in(test.fileName)

				// Restore stdout and capture output
				w.Close()
				os.Stdout = oldStdout
				var buf bytes.Buffer
				buf.ReadFrom(r)
				output := buf.String()

				cleanOutput := StripANSI(output)
				cleanExpected := StripANSI(test.expectedOutput)

				// Assert the expected output (adjust as per your implementation's expected output)
				if cleanOutput != cleanExpected {
					t.Errorf("Unexpected output.\nGot:\n%s\nExpected:\n%s", cleanOutput, cleanExpected)
				}
			} else {

			}
		})
	}

}
