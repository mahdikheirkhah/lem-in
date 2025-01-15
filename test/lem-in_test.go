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
		expectedTurns  int
		expectedOutput string
	}{
		{
			name:          "InValid Test1",
			fileName:      "../examples/badexample00.txt",
			expectedError: "Error ERROR: invalid data format, invalid number of Ants",
		},
		{
			name:          "InValid Test2",
			fileName:      "../examples/badexample01.txt",
			expectedError: "Error ERROR: invalid data format, no path found",
		},
		{
			name:          "InValid Test3",
			fileName:      "../examples/badexample02.txt",
			expectedError: "Error ERROR: invalid data format, more than one start room found",
		},
		{
			name:          "InValid Test4",
			fileName:      "../examples/badexample03.txt",
			expectedError: "Error ERROR: invalid data format, more than one end room found"},
		{
			name:          "InValid Test5",
			fileName:      "../examples/badexample04.txt",
			expectedError: "Error ERROR: invalid data format, no start room found",
		},
		{
			name:          "InValid Test6",
			fileName:      "../examples/badexample05.txt",
			expectedError: "Error ERROR: invalid data format, no end room found",
		},
		{
			name:          "InValid Test7",
			fileName:      "../examples/badexample06.txt",
			expectedError: "Error ERROR: invalid data format",
		},
		{
			name:          "InValid Test8",
			fileName:      "../examples/badexample07.txt",
			expectedError: "Error ERROR: invalid data format, invalid tunnel format",
		},
		{
			name:          "Valid Test1",
			fileName:      "../examples/example00.txt",
			expectedTurns: 6,
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
		{
			name:          "Valid Test2",
			fileName:      "../examples/example01.txt",
			expectedTurns: 8,
			expectedOutput: `10
##start
start 1 6
0 4 8
o 6 8
n 6 6
e 8 4
t 1 9
E 5 9
a 8 9
m 8 6
h 4 6
A 5 2
c 8 1
k 11 2
##end
end 11 6
start-t
n-e
a-m
A-c
0-o
E-a
k-end
start-h
o-n
m-end
t-E
start-0
h-A
e-end
c-k
n-m
h-n

turn 1: L1-t L2-h L4-0 
turn 2: L1-E L3-t L2-A L5-h L4-o L7-0 
turn 3: L1-a L3-E L6-t L2-c L5-A L8-h L4-n L7-o L10-0 
turn 4: L1-m L3-a L6-E L9-t L2-k L5-c L8-A L4-e L7-n L10-o 
turn 5: L1-end L3-m L6-a L9-E L2-end L5-k L8-c L4-end L7-e L10-n 
turn 6: L3-end L6-m L9-a L5-end L8-k L7-end L10-e 
turn 7: L6-end L9-m L8-end L10-end 
turn 8: L9-end 
`,
		},
		{
			name:          "Valid Test3",
			fileName:      "../examples/example02.txt",
			expectedTurns: 11,
			expectedOutput: `20
##start
0 2 0
1 4 1
2 6 0
##end
3 5 3
0-1
0-3
1-2
3-2

turn 1: L1-3 L4-1 
turn 2: L2-3 L4-2 L6-1 
turn 3: L3-3 L4-3 L6-2 L8-1 
turn 4: L5-3 L6-3 L8-2 L10-1 
turn 5: L7-3 L8-3 L10-2 L12-1 
turn 6: L9-3 L10-3 L12-2 L14-1 
turn 7: L11-3 L12-3 L14-2 L16-1 
turn 8: L13-3 L14-3 L16-2 L18-1 
turn 9: L15-3 L16-3 L18-2 L20-1 
turn 10: L17-3 L18-3 L20-2 
turn 11: L19-3 L20-3 
`,
		},
		{name: "Valid Test4",
			fileName:      "../examples/example03.txt",
			expectedTurns: 6,
			expectedOutput: `4
4 5 4
##start
0 1 4
1 3 6
##end
5 6 4
2 3 4
3 3 1
0-1
2-4
1-4
0-2
4-5
3-0
4-3

turn 1: L1-1 
turn 2: L1-4 L2-1 
turn 3: L1-5 L2-4 L3-1 
turn 4: L2-5 L3-4 L4-1 
turn 5: L3-5 L4-4 
turn 6: L4-5 
`},
		{
			name:          "Valid Test5",
			fileName:      "../examples/example04.txt",
			expectedTurns: 6,
			expectedOutput: `9
##start
richard 0 6
gilfoyle 6 3
erlich 9 6
dinish 6 9
jimYoung 11 7
##end
peter 14 6
richard-dinish
dinish-jimYoung
richard-gilfoyle
gilfoyle-peter
gilfoyle-erlich
richard-erlich
erlich-jimYoung
jimYoung-peter

turn 1: L1-gilfoyle L3-dinish 
turn 2: L1-peter L2-gilfoyle L3-jimYoung L5-dinish 
turn 3: L2-peter L4-gilfoyle L3-peter L5-jimYoung L7-dinish 
turn 4: L4-peter L6-gilfoyle L5-peter L7-jimYoung L9-dinish 
turn 5: L6-peter L8-gilfoyle L7-peter L9-jimYoung 
turn 6: L8-peter L9-peter 
`},
		{
			name:          "Valid Test6",
			fileName:      "../examples/example05.txt",
			expectedTurns: 8,
			expectedOutput: `9
#rooms
##start
start 0 3
##end
end 10 1
C0 1 0
C1 2 0
C2 3 0
C3 4 0
I4 5 0
I5 6 0
A0 1 2
A1 2 1
A2 4 1
B0 1 4
B1 2 4
E2 6 4
D1 6 3
D2 7 3
D3 8 3
H4 4 2
H3 5 2
F2 6 2
F3 7 2
F4 8 2
G0 1 5
G1 2 5
G2 3 5
G3 4 5
G4 6 5
H3-F2
H3-H4
H4-A2
start-G0
G0-G1
G1-G2
G2-G3
G3-G4
G4-D3
start-A0
A0-A1
A0-D1
A1-A2
A1-B1
A2-end
A2-C3
start-B0
B0-B1
B1-E2
start-C0
C0-C1
C1-C2
C2-C3
C3-I4
D1-D2
D1-F2
D2-E2
D2-D3
D2-F3
D3-end
F2-F3
F3-F4
F4-end
I4-I5
I5-end

turn 1: L1-A0 L4-B0 L8-C0 
turn 2: L1-A1 L2-A0 L4-B1 L6-B0 L8-C1 
turn 3: L1-A2 L2-A1 L3-A0 L4-E2 L6-B1 L9-B0 L8-C2 
turn 4: L1-end L2-A2 L3-A1 L5-A0 L4-D2 L6-E2 L9-B1 L8-C3 
turn 5: L2-end L3-A2 L5-A1 L7-A0 L4-D3 L6-D2 L9-E2 L8-I4 
turn 6: L3-end L5-A2 L7-A1 L4-end L6-D3 L9-D2 L8-I5 
turn 7: L5-end L7-A2 L6-end L9-D3 L8-end 
turn 8: L7-end L9-end 
`},

		{
			name:          "Valid Test7",
			fileName:      "../examples/example06.txt",
			expectedTurns: 52,
			expectedOutput: `100
##start
richard 0 6
gilfoyle 6 3
erlich 9 6
dinish 6 9
jimYoung 11 7
##end
peter 14 6
richard-dinish
dinish-jimYoung
richard-gilfoyle
gilfoyle-peter
gilfoyle-erlich
richard-erlich
erlich-jimYoung
jimYoung-peter

turn 1: L1-gilfoyle L3-dinish 
turn 2: L1-peter L2-gilfoyle L3-jimYoung L5-dinish 
turn 3: L2-peter L4-gilfoyle L3-peter L5-jimYoung L7-dinish 
turn 4: L4-peter L6-gilfoyle L5-peter L7-jimYoung L9-dinish 
turn 5: L6-peter L8-gilfoyle L7-peter L9-jimYoung L11-dinish 
turn 6: L8-peter L10-gilfoyle L9-peter L11-jimYoung L13-dinish 
turn 7: L10-peter L12-gilfoyle L11-peter L13-jimYoung L15-dinish 
turn 8: L12-peter L14-gilfoyle L13-peter L15-jimYoung L17-dinish 
turn 9: L14-peter L16-gilfoyle L15-peter L17-jimYoung L19-dinish 
turn 10: L16-peter L18-gilfoyle L17-peter L19-jimYoung L21-dinish 
turn 11: L18-peter L20-gilfoyle L19-peter L21-jimYoung L23-dinish 
turn 12: L20-peter L22-gilfoyle L21-peter L23-jimYoung L25-dinish 
turn 13: L22-peter L24-gilfoyle L23-peter L25-jimYoung L27-dinish 
turn 14: L24-peter L26-gilfoyle L25-peter L27-jimYoung L29-dinish 
turn 15: L26-peter L28-gilfoyle L27-peter L29-jimYoung L31-dinish 
turn 16: L28-peter L30-gilfoyle L29-peter L31-jimYoung L33-dinish 
turn 17: L30-peter L32-gilfoyle L31-peter L33-jimYoung L35-dinish 
turn 18: L32-peter L34-gilfoyle L33-peter L35-jimYoung L37-dinish 
turn 19: L34-peter L36-gilfoyle L35-peter L37-jimYoung L39-dinish 
turn 20: L36-peter L38-gilfoyle L37-peter L39-jimYoung L41-dinish 
turn 21: L38-peter L40-gilfoyle L39-peter L41-jimYoung L43-dinish 
turn 22: L40-peter L42-gilfoyle L41-peter L43-jimYoung L45-dinish 
turn 23: L42-peter L44-gilfoyle L43-peter L45-jimYoung L47-dinish 
turn 24: L44-peter L46-gilfoyle L45-peter L47-jimYoung L49-dinish 
turn 25: L46-peter L48-gilfoyle L47-peter L49-jimYoung L51-dinish 
turn 26: L48-peter L50-gilfoyle L49-peter L51-jimYoung L53-dinish 
turn 27: L50-peter L52-gilfoyle L51-peter L53-jimYoung L55-dinish 
turn 28: L52-peter L54-gilfoyle L53-peter L55-jimYoung L57-dinish 
turn 29: L54-peter L56-gilfoyle L55-peter L57-jimYoung L59-dinish 
turn 30: L56-peter L58-gilfoyle L57-peter L59-jimYoung L61-dinish 
turn 31: L58-peter L60-gilfoyle L59-peter L61-jimYoung L63-dinish 
turn 32: L60-peter L62-gilfoyle L61-peter L63-jimYoung L65-dinish 
turn 33: L62-peter L64-gilfoyle L63-peter L65-jimYoung L67-dinish 
turn 34: L64-peter L66-gilfoyle L65-peter L67-jimYoung L69-dinish 
turn 35: L66-peter L68-gilfoyle L67-peter L69-jimYoung L71-dinish 
turn 36: L68-peter L70-gilfoyle L69-peter L71-jimYoung L73-dinish 
turn 37: L70-peter L72-gilfoyle L71-peter L73-jimYoung L75-dinish 
turn 38: L72-peter L74-gilfoyle L73-peter L75-jimYoung L77-dinish 
turn 39: L74-peter L76-gilfoyle L75-peter L77-jimYoung L79-dinish 
turn 40: L76-peter L78-gilfoyle L77-peter L79-jimYoung L81-dinish 
turn 41: L78-peter L80-gilfoyle L79-peter L81-jimYoung L83-dinish 
turn 42: L80-peter L82-gilfoyle L81-peter L83-jimYoung L85-dinish 
turn 43: L82-peter L84-gilfoyle L83-peter L85-jimYoung L87-dinish 
turn 44: L84-peter L86-gilfoyle L85-peter L87-jimYoung L89-dinish 
turn 45: L86-peter L88-gilfoyle L87-peter L89-jimYoung L91-dinish 
turn 46: L88-peter L90-gilfoyle L89-peter L91-jimYoung L93-dinish 
turn 47: L90-peter L92-gilfoyle L91-peter L93-jimYoung L95-dinish 
turn 48: L92-peter L94-gilfoyle L93-peter L95-jimYoung L97-dinish 
turn 49: L94-peter L96-gilfoyle L95-peter L97-jimYoung L99-dinish 
turn 50: L96-peter L98-gilfoyle L97-peter L99-jimYoung 
turn 51: L98-peter L100-gilfoyle L99-peter 
turn 52: L100-peter 
`},
		{name: "Valid Test8",
			fileName:      "../examples/example07.txt",
			expectedTurns: 502,
			expectedOutput: `1000
##start
richard 0 6
gilfoyle 6 3
erlich 9 6
dinish 6 9
jimYoung 11 7
##end
peter 14 6
richard-dinish
dinish-jimYoung
richard-gilfoyle
gilfoyle-peter
gilfoyle-erlich
richard-erlich
erlich-jimYoung
jimYoung-peter

turn 1: L1-gilfoyle L3-dinish 
turn 2: L1-peter L2-gilfoyle L3-jimYoung L5-dinish 
turn 3: L2-peter L4-gilfoyle L3-peter L5-jimYoung L7-dinish 
turn 4: L4-peter L6-gilfoyle L5-peter L7-jimYoung L9-dinish 
turn 5: L6-peter L8-gilfoyle L7-peter L9-jimYoung L11-dinish 
turn 6: L8-peter L10-gilfoyle L9-peter L11-jimYoung L13-dinish 
turn 7: L10-peter L12-gilfoyle L11-peter L13-jimYoung L15-dinish 
turn 8: L12-peter L14-gilfoyle L13-peter L15-jimYoung L17-dinish 
turn 9: L14-peter L16-gilfoyle L15-peter L17-jimYoung L19-dinish 
turn 10: L16-peter L18-gilfoyle L17-peter L19-jimYoung L21-dinish 
turn 11: L18-peter L20-gilfoyle L19-peter L21-jimYoung L23-dinish 
turn 12: L20-peter L22-gilfoyle L21-peter L23-jimYoung L25-dinish 
turn 13: L22-peter L24-gilfoyle L23-peter L25-jimYoung L27-dinish 
turn 14: L24-peter L26-gilfoyle L25-peter L27-jimYoung L29-dinish 
turn 15: L26-peter L28-gilfoyle L27-peter L29-jimYoung L31-dinish 
turn 16: L28-peter L30-gilfoyle L29-peter L31-jimYoung L33-dinish 
turn 17: L30-peter L32-gilfoyle L31-peter L33-jimYoung L35-dinish 
turn 18: L32-peter L34-gilfoyle L33-peter L35-jimYoung L37-dinish 
turn 19: L34-peter L36-gilfoyle L35-peter L37-jimYoung L39-dinish 
turn 20: L36-peter L38-gilfoyle L37-peter L39-jimYoung L41-dinish 
turn 21: L38-peter L40-gilfoyle L39-peter L41-jimYoung L43-dinish 
turn 22: L40-peter L42-gilfoyle L41-peter L43-jimYoung L45-dinish 
turn 23: L42-peter L44-gilfoyle L43-peter L45-jimYoung L47-dinish 
turn 24: L44-peter L46-gilfoyle L45-peter L47-jimYoung L49-dinish 
turn 25: L46-peter L48-gilfoyle L47-peter L49-jimYoung L51-dinish 
turn 26: L48-peter L50-gilfoyle L49-peter L51-jimYoung L53-dinish 
turn 27: L50-peter L52-gilfoyle L51-peter L53-jimYoung L55-dinish 
turn 28: L52-peter L54-gilfoyle L53-peter L55-jimYoung L57-dinish 
turn 29: L54-peter L56-gilfoyle L55-peter L57-jimYoung L59-dinish 
turn 30: L56-peter L58-gilfoyle L57-peter L59-jimYoung L61-dinish 
turn 31: L58-peter L60-gilfoyle L59-peter L61-jimYoung L63-dinish 
turn 32: L60-peter L62-gilfoyle L61-peter L63-jimYoung L65-dinish 
turn 33: L62-peter L64-gilfoyle L63-peter L65-jimYoung L67-dinish 
turn 34: L64-peter L66-gilfoyle L65-peter L67-jimYoung L69-dinish 
turn 35: L66-peter L68-gilfoyle L67-peter L69-jimYoung L71-dinish 
turn 36: L68-peter L70-gilfoyle L69-peter L71-jimYoung L73-dinish 
turn 37: L70-peter L72-gilfoyle L71-peter L73-jimYoung L75-dinish 
turn 38: L72-peter L74-gilfoyle L73-peter L75-jimYoung L77-dinish 
turn 39: L74-peter L76-gilfoyle L75-peter L77-jimYoung L79-dinish 
turn 40: L76-peter L78-gilfoyle L77-peter L79-jimYoung L81-dinish 
turn 41: L78-peter L80-gilfoyle L79-peter L81-jimYoung L83-dinish 
turn 42: L80-peter L82-gilfoyle L81-peter L83-jimYoung L85-dinish 
turn 43: L82-peter L84-gilfoyle L83-peter L85-jimYoung L87-dinish 
turn 44: L84-peter L86-gilfoyle L85-peter L87-jimYoung L89-dinish 
turn 45: L86-peter L88-gilfoyle L87-peter L89-jimYoung L91-dinish 
turn 46: L88-peter L90-gilfoyle L89-peter L91-jimYoung L93-dinish 
turn 47: L90-peter L92-gilfoyle L91-peter L93-jimYoung L95-dinish 
turn 48: L92-peter L94-gilfoyle L93-peter L95-jimYoung L97-dinish 
turn 49: L94-peter L96-gilfoyle L95-peter L97-jimYoung L99-dinish 
turn 50: L96-peter L98-gilfoyle L97-peter L99-jimYoung L101-dinish 
turn 51: L98-peter L100-gilfoyle L99-peter L101-jimYoung L103-dinish 
turn 52: L100-peter L102-gilfoyle L101-peter L103-jimYoung L105-dinish 
turn 53: L102-peter L104-gilfoyle L103-peter L105-jimYoung L107-dinish 
turn 54: L104-peter L106-gilfoyle L105-peter L107-jimYoung L109-dinish 
turn 55: L106-peter L108-gilfoyle L107-peter L109-jimYoung L111-dinish 
turn 56: L108-peter L110-gilfoyle L109-peter L111-jimYoung L113-dinish 
turn 57: L110-peter L112-gilfoyle L111-peter L113-jimYoung L115-dinish 
turn 58: L112-peter L114-gilfoyle L113-peter L115-jimYoung L117-dinish 
turn 59: L114-peter L116-gilfoyle L115-peter L117-jimYoung L119-dinish 
turn 60: L116-peter L118-gilfoyle L117-peter L119-jimYoung L121-dinish 
turn 61: L118-peter L120-gilfoyle L119-peter L121-jimYoung L123-dinish 
turn 62: L120-peter L122-gilfoyle L121-peter L123-jimYoung L125-dinish 
turn 63: L122-peter L124-gilfoyle L123-peter L125-jimYoung L127-dinish 
turn 64: L124-peter L126-gilfoyle L125-peter L127-jimYoung L129-dinish 
turn 65: L126-peter L128-gilfoyle L127-peter L129-jimYoung L131-dinish 
turn 66: L128-peter L130-gilfoyle L129-peter L131-jimYoung L133-dinish 
turn 67: L130-peter L132-gilfoyle L131-peter L133-jimYoung L135-dinish 
turn 68: L132-peter L134-gilfoyle L133-peter L135-jimYoung L137-dinish 
turn 69: L134-peter L136-gilfoyle L135-peter L137-jimYoung L139-dinish 
turn 70: L136-peter L138-gilfoyle L137-peter L139-jimYoung L141-dinish 
turn 71: L138-peter L140-gilfoyle L139-peter L141-jimYoung L143-dinish 
turn 72: L140-peter L142-gilfoyle L141-peter L143-jimYoung L145-dinish 
turn 73: L142-peter L144-gilfoyle L143-peter L145-jimYoung L147-dinish 
turn 74: L144-peter L146-gilfoyle L145-peter L147-jimYoung L149-dinish 
turn 75: L146-peter L148-gilfoyle L147-peter L149-jimYoung L151-dinish 
turn 76: L148-peter L150-gilfoyle L149-peter L151-jimYoung L153-dinish 
turn 77: L150-peter L152-gilfoyle L151-peter L153-jimYoung L155-dinish 
turn 78: L152-peter L154-gilfoyle L153-peter L155-jimYoung L157-dinish 
turn 79: L154-peter L156-gilfoyle L155-peter L157-jimYoung L159-dinish 
turn 80: L156-peter L158-gilfoyle L157-peter L159-jimYoung L161-dinish 
turn 81: L158-peter L160-gilfoyle L159-peter L161-jimYoung L163-dinish 
turn 82: L160-peter L162-gilfoyle L161-peter L163-jimYoung L165-dinish 
turn 83: L162-peter L164-gilfoyle L163-peter L165-jimYoung L167-dinish 
turn 84: L164-peter L166-gilfoyle L165-peter L167-jimYoung L169-dinish 
turn 85: L166-peter L168-gilfoyle L167-peter L169-jimYoung L171-dinish 
turn 86: L168-peter L170-gilfoyle L169-peter L171-jimYoung L173-dinish 
turn 87: L170-peter L172-gilfoyle L171-peter L173-jimYoung L175-dinish 
turn 88: L172-peter L174-gilfoyle L173-peter L175-jimYoung L177-dinish 
turn 89: L174-peter L176-gilfoyle L175-peter L177-jimYoung L179-dinish 
turn 90: L176-peter L178-gilfoyle L177-peter L179-jimYoung L181-dinish 
turn 91: L178-peter L180-gilfoyle L179-peter L181-jimYoung L183-dinish 
turn 92: L180-peter L182-gilfoyle L181-peter L183-jimYoung L185-dinish 
turn 93: L182-peter L184-gilfoyle L183-peter L185-jimYoung L187-dinish 
turn 94: L184-peter L186-gilfoyle L185-peter L187-jimYoung L189-dinish 
turn 95: L186-peter L188-gilfoyle L187-peter L189-jimYoung L191-dinish 
turn 96: L188-peter L190-gilfoyle L189-peter L191-jimYoung L193-dinish 
turn 97: L190-peter L192-gilfoyle L191-peter L193-jimYoung L195-dinish 
turn 98: L192-peter L194-gilfoyle L193-peter L195-jimYoung L197-dinish 
turn 99: L194-peter L196-gilfoyle L195-peter L197-jimYoung L199-dinish 
turn 100: L196-peter L198-gilfoyle L197-peter L199-jimYoung L201-dinish 
turn 101: L198-peter L200-gilfoyle L199-peter L201-jimYoung L203-dinish 
turn 102: L200-peter L202-gilfoyle L201-peter L203-jimYoung L205-dinish 
turn 103: L202-peter L204-gilfoyle L203-peter L205-jimYoung L207-dinish 
turn 104: L204-peter L206-gilfoyle L205-peter L207-jimYoung L209-dinish 
turn 105: L206-peter L208-gilfoyle L207-peter L209-jimYoung L211-dinish 
turn 106: L208-peter L210-gilfoyle L209-peter L211-jimYoung L213-dinish 
turn 107: L210-peter L212-gilfoyle L211-peter L213-jimYoung L215-dinish 
turn 108: L212-peter L214-gilfoyle L213-peter L215-jimYoung L217-dinish 
turn 109: L214-peter L216-gilfoyle L215-peter L217-jimYoung L219-dinish 
turn 110: L216-peter L218-gilfoyle L217-peter L219-jimYoung L221-dinish 
turn 111: L218-peter L220-gilfoyle L219-peter L221-jimYoung L223-dinish 
turn 112: L220-peter L222-gilfoyle L221-peter L223-jimYoung L225-dinish 
turn 113: L222-peter L224-gilfoyle L223-peter L225-jimYoung L227-dinish 
turn 114: L224-peter L226-gilfoyle L225-peter L227-jimYoung L229-dinish 
turn 115: L226-peter L228-gilfoyle L227-peter L229-jimYoung L231-dinish 
turn 116: L228-peter L230-gilfoyle L229-peter L231-jimYoung L233-dinish 
turn 117: L230-peter L232-gilfoyle L231-peter L233-jimYoung L235-dinish 
turn 118: L232-peter L234-gilfoyle L233-peter L235-jimYoung L237-dinish 
turn 119: L234-peter L236-gilfoyle L235-peter L237-jimYoung L239-dinish 
turn 120: L236-peter L238-gilfoyle L237-peter L239-jimYoung L241-dinish 
turn 121: L238-peter L240-gilfoyle L239-peter L241-jimYoung L243-dinish 
turn 122: L240-peter L242-gilfoyle L241-peter L243-jimYoung L245-dinish 
turn 123: L242-peter L244-gilfoyle L243-peter L245-jimYoung L247-dinish 
turn 124: L244-peter L246-gilfoyle L245-peter L247-jimYoung L249-dinish 
turn 125: L246-peter L248-gilfoyle L247-peter L249-jimYoung L251-dinish 
turn 126: L248-peter L250-gilfoyle L249-peter L251-jimYoung L253-dinish 
turn 127: L250-peter L252-gilfoyle L251-peter L253-jimYoung L255-dinish 
turn 128: L252-peter L254-gilfoyle L253-peter L255-jimYoung L257-dinish 
turn 129: L254-peter L256-gilfoyle L255-peter L257-jimYoung L259-dinish 
turn 130: L256-peter L258-gilfoyle L257-peter L259-jimYoung L261-dinish 
turn 131: L258-peter L260-gilfoyle L259-peter L261-jimYoung L263-dinish 
turn 132: L260-peter L262-gilfoyle L261-peter L263-jimYoung L265-dinish 
turn 133: L262-peter L264-gilfoyle L263-peter L265-jimYoung L267-dinish 
turn 134: L264-peter L266-gilfoyle L265-peter L267-jimYoung L269-dinish 
turn 135: L266-peter L268-gilfoyle L267-peter L269-jimYoung L271-dinish 
turn 136: L268-peter L270-gilfoyle L269-peter L271-jimYoung L273-dinish 
turn 137: L270-peter L272-gilfoyle L271-peter L273-jimYoung L275-dinish 
turn 138: L272-peter L274-gilfoyle L273-peter L275-jimYoung L277-dinish 
turn 139: L274-peter L276-gilfoyle L275-peter L277-jimYoung L279-dinish 
turn 140: L276-peter L278-gilfoyle L277-peter L279-jimYoung L281-dinish 
turn 141: L278-peter L280-gilfoyle L279-peter L281-jimYoung L283-dinish 
turn 142: L280-peter L282-gilfoyle L281-peter L283-jimYoung L285-dinish 
turn 143: L282-peter L284-gilfoyle L283-peter L285-jimYoung L287-dinish 
turn 144: L284-peter L286-gilfoyle L285-peter L287-jimYoung L289-dinish 
turn 145: L286-peter L288-gilfoyle L287-peter L289-jimYoung L291-dinish 
turn 146: L288-peter L290-gilfoyle L289-peter L291-jimYoung L293-dinish 
turn 147: L290-peter L292-gilfoyle L291-peter L293-jimYoung L295-dinish 
turn 148: L292-peter L294-gilfoyle L293-peter L295-jimYoung L297-dinish 
turn 149: L294-peter L296-gilfoyle L295-peter L297-jimYoung L299-dinish 
turn 150: L296-peter L298-gilfoyle L297-peter L299-jimYoung L301-dinish 
turn 151: L298-peter L300-gilfoyle L299-peter L301-jimYoung L303-dinish 
turn 152: L300-peter L302-gilfoyle L301-peter L303-jimYoung L305-dinish 
turn 153: L302-peter L304-gilfoyle L303-peter L305-jimYoung L307-dinish 
turn 154: L304-peter L306-gilfoyle L305-peter L307-jimYoung L309-dinish 
turn 155: L306-peter L308-gilfoyle L307-peter L309-jimYoung L311-dinish 
turn 156: L308-peter L310-gilfoyle L309-peter L311-jimYoung L313-dinish 
turn 157: L310-peter L312-gilfoyle L311-peter L313-jimYoung L315-dinish 
turn 158: L312-peter L314-gilfoyle L313-peter L315-jimYoung L317-dinish 
turn 159: L314-peter L316-gilfoyle L315-peter L317-jimYoung L319-dinish 
turn 160: L316-peter L318-gilfoyle L317-peter L319-jimYoung L321-dinish 
turn 161: L318-peter L320-gilfoyle L319-peter L321-jimYoung L323-dinish 
turn 162: L320-peter L322-gilfoyle L321-peter L323-jimYoung L325-dinish 
turn 163: L322-peter L324-gilfoyle L323-peter L325-jimYoung L327-dinish 
turn 164: L324-peter L326-gilfoyle L325-peter L327-jimYoung L329-dinish 
turn 165: L326-peter L328-gilfoyle L327-peter L329-jimYoung L331-dinish 
turn 166: L328-peter L330-gilfoyle L329-peter L331-jimYoung L333-dinish 
turn 167: L330-peter L332-gilfoyle L331-peter L333-jimYoung L335-dinish 
turn 168: L332-peter L334-gilfoyle L333-peter L335-jimYoung L337-dinish 
turn 169: L334-peter L336-gilfoyle L335-peter L337-jimYoung L339-dinish 
turn 170: L336-peter L338-gilfoyle L337-peter L339-jimYoung L341-dinish 
turn 171: L338-peter L340-gilfoyle L339-peter L341-jimYoung L343-dinish 
turn 172: L340-peter L342-gilfoyle L341-peter L343-jimYoung L345-dinish 
turn 173: L342-peter L344-gilfoyle L343-peter L345-jimYoung L347-dinish 
turn 174: L344-peter L346-gilfoyle L345-peter L347-jimYoung L349-dinish 
turn 175: L346-peter L348-gilfoyle L347-peter L349-jimYoung L351-dinish 
turn 176: L348-peter L350-gilfoyle L349-peter L351-jimYoung L353-dinish 
turn 177: L350-peter L352-gilfoyle L351-peter L353-jimYoung L355-dinish 
turn 178: L352-peter L354-gilfoyle L353-peter L355-jimYoung L357-dinish 
turn 179: L354-peter L356-gilfoyle L355-peter L357-jimYoung L359-dinish 
turn 180: L356-peter L358-gilfoyle L357-peter L359-jimYoung L361-dinish 
turn 181: L358-peter L360-gilfoyle L359-peter L361-jimYoung L363-dinish 
turn 182: L360-peter L362-gilfoyle L361-peter L363-jimYoung L365-dinish 
turn 183: L362-peter L364-gilfoyle L363-peter L365-jimYoung L367-dinish 
turn 184: L364-peter L366-gilfoyle L365-peter L367-jimYoung L369-dinish 
turn 185: L366-peter L368-gilfoyle L367-peter L369-jimYoung L371-dinish 
turn 186: L368-peter L370-gilfoyle L369-peter L371-jimYoung L373-dinish 
turn 187: L370-peter L372-gilfoyle L371-peter L373-jimYoung L375-dinish 
turn 188: L372-peter L374-gilfoyle L373-peter L375-jimYoung L377-dinish 
turn 189: L374-peter L376-gilfoyle L375-peter L377-jimYoung L379-dinish 
turn 190: L376-peter L378-gilfoyle L377-peter L379-jimYoung L381-dinish 
turn 191: L378-peter L380-gilfoyle L379-peter L381-jimYoung L383-dinish 
turn 192: L380-peter L382-gilfoyle L381-peter L383-jimYoung L385-dinish 
turn 193: L382-peter L384-gilfoyle L383-peter L385-jimYoung L387-dinish 
turn 194: L384-peter L386-gilfoyle L385-peter L387-jimYoung L389-dinish 
turn 195: L386-peter L388-gilfoyle L387-peter L389-jimYoung L391-dinish 
turn 196: L388-peter L390-gilfoyle L389-peter L391-jimYoung L393-dinish 
turn 197: L390-peter L392-gilfoyle L391-peter L393-jimYoung L395-dinish 
turn 198: L392-peter L394-gilfoyle L393-peter L395-jimYoung L397-dinish 
turn 199: L394-peter L396-gilfoyle L395-peter L397-jimYoung L399-dinish 
turn 200: L396-peter L398-gilfoyle L397-peter L399-jimYoung L401-dinish 
turn 201: L398-peter L400-gilfoyle L399-peter L401-jimYoung L403-dinish 
turn 202: L400-peter L402-gilfoyle L401-peter L403-jimYoung L405-dinish 
turn 203: L402-peter L404-gilfoyle L403-peter L405-jimYoung L407-dinish 
turn 204: L404-peter L406-gilfoyle L405-peter L407-jimYoung L409-dinish 
turn 205: L406-peter L408-gilfoyle L407-peter L409-jimYoung L411-dinish 
turn 206: L408-peter L410-gilfoyle L409-peter L411-jimYoung L413-dinish 
turn 207: L410-peter L412-gilfoyle L411-peter L413-jimYoung L415-dinish 
turn 208: L412-peter L414-gilfoyle L413-peter L415-jimYoung L417-dinish 
turn 209: L414-peter L416-gilfoyle L415-peter L417-jimYoung L419-dinish 
turn 210: L416-peter L418-gilfoyle L417-peter L419-jimYoung L421-dinish 
turn 211: L418-peter L420-gilfoyle L419-peter L421-jimYoung L423-dinish 
turn 212: L420-peter L422-gilfoyle L421-peter L423-jimYoung L425-dinish 
turn 213: L422-peter L424-gilfoyle L423-peter L425-jimYoung L427-dinish 
turn 214: L424-peter L426-gilfoyle L425-peter L427-jimYoung L429-dinish 
turn 215: L426-peter L428-gilfoyle L427-peter L429-jimYoung L431-dinish 
turn 216: L428-peter L430-gilfoyle L429-peter L431-jimYoung L433-dinish 
turn 217: L430-peter L432-gilfoyle L431-peter L433-jimYoung L435-dinish 
turn 218: L432-peter L434-gilfoyle L433-peter L435-jimYoung L437-dinish 
turn 219: L434-peter L436-gilfoyle L435-peter L437-jimYoung L439-dinish 
turn 220: L436-peter L438-gilfoyle L437-peter L439-jimYoung L441-dinish 
turn 221: L438-peter L440-gilfoyle L439-peter L441-jimYoung L443-dinish 
turn 222: L440-peter L442-gilfoyle L441-peter L443-jimYoung L445-dinish 
turn 223: L442-peter L444-gilfoyle L443-peter L445-jimYoung L447-dinish 
turn 224: L444-peter L446-gilfoyle L445-peter L447-jimYoung L449-dinish 
turn 225: L446-peter L448-gilfoyle L447-peter L449-jimYoung L451-dinish 
turn 226: L448-peter L450-gilfoyle L449-peter L451-jimYoung L453-dinish 
turn 227: L450-peter L452-gilfoyle L451-peter L453-jimYoung L455-dinish 
turn 228: L452-peter L454-gilfoyle L453-peter L455-jimYoung L457-dinish 
turn 229: L454-peter L456-gilfoyle L455-peter L457-jimYoung L459-dinish 
turn 230: L456-peter L458-gilfoyle L457-peter L459-jimYoung L461-dinish 
turn 231: L458-peter L460-gilfoyle L459-peter L461-jimYoung L463-dinish 
turn 232: L460-peter L462-gilfoyle L461-peter L463-jimYoung L465-dinish 
turn 233: L462-peter L464-gilfoyle L463-peter L465-jimYoung L467-dinish 
turn 234: L464-peter L466-gilfoyle L465-peter L467-jimYoung L469-dinish 
turn 235: L466-peter L468-gilfoyle L467-peter L469-jimYoung L471-dinish 
turn 236: L468-peter L470-gilfoyle L469-peter L471-jimYoung L473-dinish 
turn 237: L470-peter L472-gilfoyle L471-peter L473-jimYoung L475-dinish 
turn 238: L472-peter L474-gilfoyle L473-peter L475-jimYoung L477-dinish 
turn 239: L474-peter L476-gilfoyle L475-peter L477-jimYoung L479-dinish 
turn 240: L476-peter L478-gilfoyle L477-peter L479-jimYoung L481-dinish 
turn 241: L478-peter L480-gilfoyle L479-peter L481-jimYoung L483-dinish 
turn 242: L480-peter L482-gilfoyle L481-peter L483-jimYoung L485-dinish 
turn 243: L482-peter L484-gilfoyle L483-peter L485-jimYoung L487-dinish 
turn 244: L484-peter L486-gilfoyle L485-peter L487-jimYoung L489-dinish 
turn 245: L486-peter L488-gilfoyle L487-peter L489-jimYoung L491-dinish 
turn 246: L488-peter L490-gilfoyle L489-peter L491-jimYoung L493-dinish 
turn 247: L490-peter L492-gilfoyle L491-peter L493-jimYoung L495-dinish 
turn 248: L492-peter L494-gilfoyle L493-peter L495-jimYoung L497-dinish 
turn 249: L494-peter L496-gilfoyle L495-peter L497-jimYoung L499-dinish 
turn 250: L496-peter L498-gilfoyle L497-peter L499-jimYoung L501-dinish 
turn 251: L498-peter L500-gilfoyle L499-peter L501-jimYoung L503-dinish 
turn 252: L500-peter L502-gilfoyle L501-peter L503-jimYoung L505-dinish 
turn 253: L502-peter L504-gilfoyle L503-peter L505-jimYoung L507-dinish 
turn 254: L504-peter L506-gilfoyle L505-peter L507-jimYoung L509-dinish 
turn 255: L506-peter L508-gilfoyle L507-peter L509-jimYoung L511-dinish 
turn 256: L508-peter L510-gilfoyle L509-peter L511-jimYoung L513-dinish 
turn 257: L510-peter L512-gilfoyle L511-peter L513-jimYoung L515-dinish 
turn 258: L512-peter L514-gilfoyle L513-peter L515-jimYoung L517-dinish 
turn 259: L514-peter L516-gilfoyle L515-peter L517-jimYoung L519-dinish 
turn 260: L516-peter L518-gilfoyle L517-peter L519-jimYoung L521-dinish 
turn 261: L518-peter L520-gilfoyle L519-peter L521-jimYoung L523-dinish 
turn 262: L520-peter L522-gilfoyle L521-peter L523-jimYoung L525-dinish 
turn 263: L522-peter L524-gilfoyle L523-peter L525-jimYoung L527-dinish 
turn 264: L524-peter L526-gilfoyle L525-peter L527-jimYoung L529-dinish 
turn 265: L526-peter L528-gilfoyle L527-peter L529-jimYoung L531-dinish 
turn 266: L528-peter L530-gilfoyle L529-peter L531-jimYoung L533-dinish 
turn 267: L530-peter L532-gilfoyle L531-peter L533-jimYoung L535-dinish 
turn 268: L532-peter L534-gilfoyle L533-peter L535-jimYoung L537-dinish 
turn 269: L534-peter L536-gilfoyle L535-peter L537-jimYoung L539-dinish 
turn 270: L536-peter L538-gilfoyle L537-peter L539-jimYoung L541-dinish 
turn 271: L538-peter L540-gilfoyle L539-peter L541-jimYoung L543-dinish 
turn 272: L540-peter L542-gilfoyle L541-peter L543-jimYoung L545-dinish 
turn 273: L542-peter L544-gilfoyle L543-peter L545-jimYoung L547-dinish 
turn 274: L544-peter L546-gilfoyle L545-peter L547-jimYoung L549-dinish 
turn 275: L546-peter L548-gilfoyle L547-peter L549-jimYoung L551-dinish 
turn 276: L548-peter L550-gilfoyle L549-peter L551-jimYoung L553-dinish 
turn 277: L550-peter L552-gilfoyle L551-peter L553-jimYoung L555-dinish 
turn 278: L552-peter L554-gilfoyle L553-peter L555-jimYoung L557-dinish 
turn 279: L554-peter L556-gilfoyle L555-peter L557-jimYoung L559-dinish 
turn 280: L556-peter L558-gilfoyle L557-peter L559-jimYoung L561-dinish 
turn 281: L558-peter L560-gilfoyle L559-peter L561-jimYoung L563-dinish 
turn 282: L560-peter L562-gilfoyle L561-peter L563-jimYoung L565-dinish 
turn 283: L562-peter L564-gilfoyle L563-peter L565-jimYoung L567-dinish 
turn 284: L564-peter L566-gilfoyle L565-peter L567-jimYoung L569-dinish 
turn 285: L566-peter L568-gilfoyle L567-peter L569-jimYoung L571-dinish 
turn 286: L568-peter L570-gilfoyle L569-peter L571-jimYoung L573-dinish 
turn 287: L570-peter L572-gilfoyle L571-peter L573-jimYoung L575-dinish 
turn 288: L572-peter L574-gilfoyle L573-peter L575-jimYoung L577-dinish 
turn 289: L574-peter L576-gilfoyle L575-peter L577-jimYoung L579-dinish 
turn 290: L576-peter L578-gilfoyle L577-peter L579-jimYoung L581-dinish 
turn 291: L578-peter L580-gilfoyle L579-peter L581-jimYoung L583-dinish 
turn 292: L580-peter L582-gilfoyle L581-peter L583-jimYoung L585-dinish 
turn 293: L582-peter L584-gilfoyle L583-peter L585-jimYoung L587-dinish 
turn 294: L584-peter L586-gilfoyle L585-peter L587-jimYoung L589-dinish 
turn 295: L586-peter L588-gilfoyle L587-peter L589-jimYoung L591-dinish 
turn 296: L588-peter L590-gilfoyle L589-peter L591-jimYoung L593-dinish 
turn 297: L590-peter L592-gilfoyle L591-peter L593-jimYoung L595-dinish 
turn 298: L592-peter L594-gilfoyle L593-peter L595-jimYoung L597-dinish 
turn 299: L594-peter L596-gilfoyle L595-peter L597-jimYoung L599-dinish 
turn 300: L596-peter L598-gilfoyle L597-peter L599-jimYoung L601-dinish 
turn 301: L598-peter L600-gilfoyle L599-peter L601-jimYoung L603-dinish 
turn 302: L600-peter L602-gilfoyle L601-peter L603-jimYoung L605-dinish 
turn 303: L602-peter L604-gilfoyle L603-peter L605-jimYoung L607-dinish 
turn 304: L604-peter L606-gilfoyle L605-peter L607-jimYoung L609-dinish 
turn 305: L606-peter L608-gilfoyle L607-peter L609-jimYoung L611-dinish 
turn 306: L608-peter L610-gilfoyle L609-peter L611-jimYoung L613-dinish 
turn 307: L610-peter L612-gilfoyle L611-peter L613-jimYoung L615-dinish 
turn 308: L612-peter L614-gilfoyle L613-peter L615-jimYoung L617-dinish 
turn 309: L614-peter L616-gilfoyle L615-peter L617-jimYoung L619-dinish 
turn 310: L616-peter L618-gilfoyle L617-peter L619-jimYoung L621-dinish 
turn 311: L618-peter L620-gilfoyle L619-peter L621-jimYoung L623-dinish 
turn 312: L620-peter L622-gilfoyle L621-peter L623-jimYoung L625-dinish 
turn 313: L622-peter L624-gilfoyle L623-peter L625-jimYoung L627-dinish 
turn 314: L624-peter L626-gilfoyle L625-peter L627-jimYoung L629-dinish 
turn 315: L626-peter L628-gilfoyle L627-peter L629-jimYoung L631-dinish 
turn 316: L628-peter L630-gilfoyle L629-peter L631-jimYoung L633-dinish 
turn 317: L630-peter L632-gilfoyle L631-peter L633-jimYoung L635-dinish 
turn 318: L632-peter L634-gilfoyle L633-peter L635-jimYoung L637-dinish 
turn 319: L634-peter L636-gilfoyle L635-peter L637-jimYoung L639-dinish 
turn 320: L636-peter L638-gilfoyle L637-peter L639-jimYoung L641-dinish 
turn 321: L638-peter L640-gilfoyle L639-peter L641-jimYoung L643-dinish 
turn 322: L640-peter L642-gilfoyle L641-peter L643-jimYoung L645-dinish 
turn 323: L642-peter L644-gilfoyle L643-peter L645-jimYoung L647-dinish 
turn 324: L644-peter L646-gilfoyle L645-peter L647-jimYoung L649-dinish 
turn 325: L646-peter L648-gilfoyle L647-peter L649-jimYoung L651-dinish 
turn 326: L648-peter L650-gilfoyle L649-peter L651-jimYoung L653-dinish 
turn 327: L650-peter L652-gilfoyle L651-peter L653-jimYoung L655-dinish 
turn 328: L652-peter L654-gilfoyle L653-peter L655-jimYoung L657-dinish 
turn 329: L654-peter L656-gilfoyle L655-peter L657-jimYoung L659-dinish 
turn 330: L656-peter L658-gilfoyle L657-peter L659-jimYoung L661-dinish 
turn 331: L658-peter L660-gilfoyle L659-peter L661-jimYoung L663-dinish 
turn 332: L660-peter L662-gilfoyle L661-peter L663-jimYoung L665-dinish 
turn 333: L662-peter L664-gilfoyle L663-peter L665-jimYoung L667-dinish 
turn 334: L664-peter L666-gilfoyle L665-peter L667-jimYoung L669-dinish 
turn 335: L666-peter L668-gilfoyle L667-peter L669-jimYoung L671-dinish 
turn 336: L668-peter L670-gilfoyle L669-peter L671-jimYoung L673-dinish 
turn 337: L670-peter L672-gilfoyle L671-peter L673-jimYoung L675-dinish 
turn 338: L672-peter L674-gilfoyle L673-peter L675-jimYoung L677-dinish 
turn 339: L674-peter L676-gilfoyle L675-peter L677-jimYoung L679-dinish 
turn 340: L676-peter L678-gilfoyle L677-peter L679-jimYoung L681-dinish 
turn 341: L678-peter L680-gilfoyle L679-peter L681-jimYoung L683-dinish 
turn 342: L680-peter L682-gilfoyle L681-peter L683-jimYoung L685-dinish 
turn 343: L682-peter L684-gilfoyle L683-peter L685-jimYoung L687-dinish 
turn 344: L684-peter L686-gilfoyle L685-peter L687-jimYoung L689-dinish 
turn 345: L686-peter L688-gilfoyle L687-peter L689-jimYoung L691-dinish 
turn 346: L688-peter L690-gilfoyle L689-peter L691-jimYoung L693-dinish 
turn 347: L690-peter L692-gilfoyle L691-peter L693-jimYoung L695-dinish 
turn 348: L692-peter L694-gilfoyle L693-peter L695-jimYoung L697-dinish 
turn 349: L694-peter L696-gilfoyle L695-peter L697-jimYoung L699-dinish 
turn 350: L696-peter L698-gilfoyle L697-peter L699-jimYoung L701-dinish 
turn 351: L698-peter L700-gilfoyle L699-peter L701-jimYoung L703-dinish 
turn 352: L700-peter L702-gilfoyle L701-peter L703-jimYoung L705-dinish 
turn 353: L702-peter L704-gilfoyle L703-peter L705-jimYoung L707-dinish 
turn 354: L704-peter L706-gilfoyle L705-peter L707-jimYoung L709-dinish 
turn 355: L706-peter L708-gilfoyle L707-peter L709-jimYoung L711-dinish 
turn 356: L708-peter L710-gilfoyle L709-peter L711-jimYoung L713-dinish 
turn 357: L710-peter L712-gilfoyle L711-peter L713-jimYoung L715-dinish 
turn 358: L712-peter L714-gilfoyle L713-peter L715-jimYoung L717-dinish 
turn 359: L714-peter L716-gilfoyle L715-peter L717-jimYoung L719-dinish 
turn 360: L716-peter L718-gilfoyle L717-peter L719-jimYoung L721-dinish 
turn 361: L718-peter L720-gilfoyle L719-peter L721-jimYoung L723-dinish 
turn 362: L720-peter L722-gilfoyle L721-peter L723-jimYoung L725-dinish 
turn 363: L722-peter L724-gilfoyle L723-peter L725-jimYoung L727-dinish 
turn 364: L724-peter L726-gilfoyle L725-peter L727-jimYoung L729-dinish 
turn 365: L726-peter L728-gilfoyle L727-peter L729-jimYoung L731-dinish 
turn 366: L728-peter L730-gilfoyle L729-peter L731-jimYoung L733-dinish 
turn 367: L730-peter L732-gilfoyle L731-peter L733-jimYoung L735-dinish 
turn 368: L732-peter L734-gilfoyle L733-peter L735-jimYoung L737-dinish 
turn 369: L734-peter L736-gilfoyle L735-peter L737-jimYoung L739-dinish 
turn 370: L736-peter L738-gilfoyle L737-peter L739-jimYoung L741-dinish 
turn 371: L738-peter L740-gilfoyle L739-peter L741-jimYoung L743-dinish 
turn 372: L740-peter L742-gilfoyle L741-peter L743-jimYoung L745-dinish 
turn 373: L742-peter L744-gilfoyle L743-peter L745-jimYoung L747-dinish 
turn 374: L744-peter L746-gilfoyle L745-peter L747-jimYoung L749-dinish 
turn 375: L746-peter L748-gilfoyle L747-peter L749-jimYoung L751-dinish 
turn 376: L748-peter L750-gilfoyle L749-peter L751-jimYoung L753-dinish 
turn 377: L750-peter L752-gilfoyle L751-peter L753-jimYoung L755-dinish 
turn 378: L752-peter L754-gilfoyle L753-peter L755-jimYoung L757-dinish 
turn 379: L754-peter L756-gilfoyle L755-peter L757-jimYoung L759-dinish 
turn 380: L756-peter L758-gilfoyle L757-peter L759-jimYoung L761-dinish 
turn 381: L758-peter L760-gilfoyle L759-peter L761-jimYoung L763-dinish 
turn 382: L760-peter L762-gilfoyle L761-peter L763-jimYoung L765-dinish 
turn 383: L762-peter L764-gilfoyle L763-peter L765-jimYoung L767-dinish 
turn 384: L764-peter L766-gilfoyle L765-peter L767-jimYoung L769-dinish 
turn 385: L766-peter L768-gilfoyle L767-peter L769-jimYoung L771-dinish 
turn 386: L768-peter L770-gilfoyle L769-peter L771-jimYoung L773-dinish 
turn 387: L770-peter L772-gilfoyle L771-peter L773-jimYoung L775-dinish 
turn 388: L772-peter L774-gilfoyle L773-peter L775-jimYoung L777-dinish 
turn 389: L774-peter L776-gilfoyle L775-peter L777-jimYoung L779-dinish 
turn 390: L776-peter L778-gilfoyle L777-peter L779-jimYoung L781-dinish 
turn 391: L778-peter L780-gilfoyle L779-peter L781-jimYoung L783-dinish 
turn 392: L780-peter L782-gilfoyle L781-peter L783-jimYoung L785-dinish 
turn 393: L782-peter L784-gilfoyle L783-peter L785-jimYoung L787-dinish 
turn 394: L784-peter L786-gilfoyle L785-peter L787-jimYoung L789-dinish 
turn 395: L786-peter L788-gilfoyle L787-peter L789-jimYoung L791-dinish 
turn 396: L788-peter L790-gilfoyle L789-peter L791-jimYoung L793-dinish 
turn 397: L790-peter L792-gilfoyle L791-peter L793-jimYoung L795-dinish 
turn 398: L792-peter L794-gilfoyle L793-peter L795-jimYoung L797-dinish 
turn 399: L794-peter L796-gilfoyle L795-peter L797-jimYoung L799-dinish 
turn 400: L796-peter L798-gilfoyle L797-peter L799-jimYoung L801-dinish 
turn 401: L798-peter L800-gilfoyle L799-peter L801-jimYoung L803-dinish 
turn 402: L800-peter L802-gilfoyle L801-peter L803-jimYoung L805-dinish 
turn 403: L802-peter L804-gilfoyle L803-peter L805-jimYoung L807-dinish 
turn 404: L804-peter L806-gilfoyle L805-peter L807-jimYoung L809-dinish 
turn 405: L806-peter L808-gilfoyle L807-peter L809-jimYoung L811-dinish 
turn 406: L808-peter L810-gilfoyle L809-peter L811-jimYoung L813-dinish 
turn 407: L810-peter L812-gilfoyle L811-peter L813-jimYoung L815-dinish 
turn 408: L812-peter L814-gilfoyle L813-peter L815-jimYoung L817-dinish 
turn 409: L814-peter L816-gilfoyle L815-peter L817-jimYoung L819-dinish 
turn 410: L816-peter L818-gilfoyle L817-peter L819-jimYoung L821-dinish 
turn 411: L818-peter L820-gilfoyle L819-peter L821-jimYoung L823-dinish 
turn 412: L820-peter L822-gilfoyle L821-peter L823-jimYoung L825-dinish 
turn 413: L822-peter L824-gilfoyle L823-peter L825-jimYoung L827-dinish 
turn 414: L824-peter L826-gilfoyle L825-peter L827-jimYoung L829-dinish 
turn 415: L826-peter L828-gilfoyle L827-peter L829-jimYoung L831-dinish 
turn 416: L828-peter L830-gilfoyle L829-peter L831-jimYoung L833-dinish 
turn 417: L830-peter L832-gilfoyle L831-peter L833-jimYoung L835-dinish 
turn 418: L832-peter L834-gilfoyle L833-peter L835-jimYoung L837-dinish 
turn 419: L834-peter L836-gilfoyle L835-peter L837-jimYoung L839-dinish 
turn 420: L836-peter L838-gilfoyle L837-peter L839-jimYoung L841-dinish 
turn 421: L838-peter L840-gilfoyle L839-peter L841-jimYoung L843-dinish 
turn 422: L840-peter L842-gilfoyle L841-peter L843-jimYoung L845-dinish 
turn 423: L842-peter L844-gilfoyle L843-peter L845-jimYoung L847-dinish 
turn 424: L844-peter L846-gilfoyle L845-peter L847-jimYoung L849-dinish 
turn 425: L846-peter L848-gilfoyle L847-peter L849-jimYoung L851-dinish 
turn 426: L848-peter L850-gilfoyle L849-peter L851-jimYoung L853-dinish 
turn 427: L850-peter L852-gilfoyle L851-peter L853-jimYoung L855-dinish 
turn 428: L852-peter L854-gilfoyle L853-peter L855-jimYoung L857-dinish 
turn 429: L854-peter L856-gilfoyle L855-peter L857-jimYoung L859-dinish 
turn 430: L856-peter L858-gilfoyle L857-peter L859-jimYoung L861-dinish 
turn 431: L858-peter L860-gilfoyle L859-peter L861-jimYoung L863-dinish 
turn 432: L860-peter L862-gilfoyle L861-peter L863-jimYoung L865-dinish 
turn 433: L862-peter L864-gilfoyle L863-peter L865-jimYoung L867-dinish 
turn 434: L864-peter L866-gilfoyle L865-peter L867-jimYoung L869-dinish 
turn 435: L866-peter L868-gilfoyle L867-peter L869-jimYoung L871-dinish 
turn 436: L868-peter L870-gilfoyle L869-peter L871-jimYoung L873-dinish 
turn 437: L870-peter L872-gilfoyle L871-peter L873-jimYoung L875-dinish 
turn 438: L872-peter L874-gilfoyle L873-peter L875-jimYoung L877-dinish 
turn 439: L874-peter L876-gilfoyle L875-peter L877-jimYoung L879-dinish 
turn 440: L876-peter L878-gilfoyle L877-peter L879-jimYoung L881-dinish 
turn 441: L878-peter L880-gilfoyle L879-peter L881-jimYoung L883-dinish 
turn 442: L880-peter L882-gilfoyle L881-peter L883-jimYoung L885-dinish 
turn 443: L882-peter L884-gilfoyle L883-peter L885-jimYoung L887-dinish 
turn 444: L884-peter L886-gilfoyle L885-peter L887-jimYoung L889-dinish 
turn 445: L886-peter L888-gilfoyle L887-peter L889-jimYoung L891-dinish 
turn 446: L888-peter L890-gilfoyle L889-peter L891-jimYoung L893-dinish 
turn 447: L890-peter L892-gilfoyle L891-peter L893-jimYoung L895-dinish 
turn 448: L892-peter L894-gilfoyle L893-peter L895-jimYoung L897-dinish 
turn 449: L894-peter L896-gilfoyle L895-peter L897-jimYoung L899-dinish 
turn 450: L896-peter L898-gilfoyle L897-peter L899-jimYoung L901-dinish 
turn 451: L898-peter L900-gilfoyle L899-peter L901-jimYoung L903-dinish 
turn 452: L900-peter L902-gilfoyle L901-peter L903-jimYoung L905-dinish 
turn 453: L902-peter L904-gilfoyle L903-peter L905-jimYoung L907-dinish 
turn 454: L904-peter L906-gilfoyle L905-peter L907-jimYoung L909-dinish 
turn 455: L906-peter L908-gilfoyle L907-peter L909-jimYoung L911-dinish 
turn 456: L908-peter L910-gilfoyle L909-peter L911-jimYoung L913-dinish 
turn 457: L910-peter L912-gilfoyle L911-peter L913-jimYoung L915-dinish 
turn 458: L912-peter L914-gilfoyle L913-peter L915-jimYoung L917-dinish 
turn 459: L914-peter L916-gilfoyle L915-peter L917-jimYoung L919-dinish 
turn 460: L916-peter L918-gilfoyle L917-peter L919-jimYoung L921-dinish 
turn 461: L918-peter L920-gilfoyle L919-peter L921-jimYoung L923-dinish 
turn 462: L920-peter L922-gilfoyle L921-peter L923-jimYoung L925-dinish 
turn 463: L922-peter L924-gilfoyle L923-peter L925-jimYoung L927-dinish 
turn 464: L924-peter L926-gilfoyle L925-peter L927-jimYoung L929-dinish 
turn 465: L926-peter L928-gilfoyle L927-peter L929-jimYoung L931-dinish 
turn 466: L928-peter L930-gilfoyle L929-peter L931-jimYoung L933-dinish 
turn 467: L930-peter L932-gilfoyle L931-peter L933-jimYoung L935-dinish 
turn 468: L932-peter L934-gilfoyle L933-peter L935-jimYoung L937-dinish 
turn 469: L934-peter L936-gilfoyle L935-peter L937-jimYoung L939-dinish 
turn 470: L936-peter L938-gilfoyle L937-peter L939-jimYoung L941-dinish 
turn 471: L938-peter L940-gilfoyle L939-peter L941-jimYoung L943-dinish 
turn 472: L940-peter L942-gilfoyle L941-peter L943-jimYoung L945-dinish 
turn 473: L942-peter L944-gilfoyle L943-peter L945-jimYoung L947-dinish 
turn 474: L944-peter L946-gilfoyle L945-peter L947-jimYoung L949-dinish 
turn 475: L946-peter L948-gilfoyle L947-peter L949-jimYoung L951-dinish 
turn 476: L948-peter L950-gilfoyle L949-peter L951-jimYoung L953-dinish 
turn 477: L950-peter L952-gilfoyle L951-peter L953-jimYoung L955-dinish 
turn 478: L952-peter L954-gilfoyle L953-peter L955-jimYoung L957-dinish 
turn 479: L954-peter L956-gilfoyle L955-peter L957-jimYoung L959-dinish 
turn 480: L956-peter L958-gilfoyle L957-peter L959-jimYoung L961-dinish 
turn 481: L958-peter L960-gilfoyle L959-peter L961-jimYoung L963-dinish 
turn 482: L960-peter L962-gilfoyle L961-peter L963-jimYoung L965-dinish 
turn 483: L962-peter L964-gilfoyle L963-peter L965-jimYoung L967-dinish 
turn 484: L964-peter L966-gilfoyle L965-peter L967-jimYoung L969-dinish 
turn 485: L966-peter L968-gilfoyle L967-peter L969-jimYoung L971-dinish 
turn 486: L968-peter L970-gilfoyle L969-peter L971-jimYoung L973-dinish 
turn 487: L970-peter L972-gilfoyle L971-peter L973-jimYoung L975-dinish 
turn 488: L972-peter L974-gilfoyle L973-peter L975-jimYoung L977-dinish 
turn 489: L974-peter L976-gilfoyle L975-peter L977-jimYoung L979-dinish 
turn 490: L976-peter L978-gilfoyle L977-peter L979-jimYoung L981-dinish 
turn 491: L978-peter L980-gilfoyle L979-peter L981-jimYoung L983-dinish 
turn 492: L980-peter L982-gilfoyle L981-peter L983-jimYoung L985-dinish 
turn 493: L982-peter L984-gilfoyle L983-peter L985-jimYoung L987-dinish 
turn 494: L984-peter L986-gilfoyle L985-peter L987-jimYoung L989-dinish 
turn 495: L986-peter L988-gilfoyle L987-peter L989-jimYoung L991-dinish 
turn 496: L988-peter L990-gilfoyle L989-peter L991-jimYoung L993-dinish 
turn 497: L990-peter L992-gilfoyle L991-peter L993-jimYoung L995-dinish 
turn 498: L992-peter L994-gilfoyle L993-peter L995-jimYoung L997-dinish 
turn 499: L994-peter L996-gilfoyle L995-peter L997-jimYoung L999-dinish 
turn 500: L996-peter L998-gilfoyle L997-peter L999-jimYoung 
turn 501: L998-peter L1000-gilfoyle L999-peter 
turn 502: L1000-peter 
`},
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
				stripedOutput := removeComments(cleanOutput)
				turns, indexs := countTurns(stripedOutput)
				if turns == -1 && indexs == nil {
					t.Errorf("Error compiling regex: \n")
				}
				if turns != test.expectedTurns {
					t.Errorf("Error expected this nummber of Turns:%v \n but Got:\n%v\n", test.expectedTurns, turns)
				}

				// for i := 0; i < len(stripedOutput); i++ {
				// 	fmt.Println(i, "is:", stripedOutput[i])
				// }
				// // fmt.Println(indexs)
				// // fmt.Println(turns)
				if hasConflicts(indexs, stripedOutput) {
					t.Errorf("There is house with more than one ants in it at the same time \n")
				}
			} else {
				calledExit := false
				// Mock os.Exit to prevent the program from exiting during tests
				errorHandler.ExitFunc = func(code int) {
					calledExit = true
				}

				var buf bytes.Buffer
				log.SetOutput(&buf)

				// Reset the calledExit flag
				calledExit = false

				defer func() {
					log.SetOutput(os.Stderr) // Restore original output
				}()

				// Restore original os.Exit after tests
				defer func() {
					errorHandler.ExitFunc = os.Exit
				}()
				utils.Lem_in(test.fileName)
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

func removeComments(input string) []string {
	inputSeperated := strings.Split(input, "\n")
	var seperatedlines []string
	for _, line := range inputSeperated {
		if line != "" {
			cleanLine := strings.TrimSpace(line)
			if !strings.HasPrefix(cleanLine, "#") {
				seperatedlines = append(seperatedlines, cleanLine)
			}
		}
	}
	return seperatedlines
}

func countTurns(input []string) (int, []int) {
	counter := 0
	var indexs []int
	pattern := `^turn \d+: L`

	// Compile the regex
	re, err := regexp.Compile(pattern)
	if err != nil {
		return -1, nil
	}

	for i := 0; i < len(input); i++ {
		if re.MatchString(input[i]) {
			counter++
			indexs = append(indexs, i)
		}
	}
	return counter, indexs
}

func hasConflicts(indexs []int, stripedOutput []string) bool {
	endRoomName := findEndRoomName(stripedOutput[indexs[len(indexs)-1]])
	for _, index := range indexs {
		roomNames := extractRooms(stripedOutput[index])
		// fmt.Println(i+1, "Turn: ", roomNames)
		if !checkRooms(roomNames, endRoomName) {
			return true
		}
	}
	return false
}

func extractRooms(input string) []string {
	var roomNames []string
	var roomName string
	for i, char := range input {
		if char == '-' {
			roomName = ""
			for j := i + 1; j < len(input); j++ {
				if input[j] == ' ' || j == len(input)-1 {
					if j == len(input)-1 && input[j] != ' ' {
						roomName += string(input[j])
					}
					roomNames = append(roomNames, roomName)
					break
				}
				roomName += string(input[j])
			}
		}
	}
	return roomNames
}

func checkRooms(roomNames []string, endRoomName string) bool {
	for i := 0; i < len(roomNames); i++ {
		for j := 0; j < len(roomNames); j++ {
			if j != i && roomNames[j] == roomNames[i] && roomNames[j] != endRoomName {
				return false
			}
		}
	}
	return true
}

func findEndRoomName(lastLine string) string {
	var roomName string
	for i, char := range lastLine {
		if char == '-' {
			roomName = ""
			for j := i + 1; j < len(lastLine); j++ {
				if lastLine[j] == ' ' || j == len(lastLine)-1 {
					if j == len(lastLine)-1 && lastLine[j] != ' ' {
						roomName += string(lastLine[j])
					}
					return roomName
				}
				roomName += string(lastLine[j])
			}
		}
	}
	return roomName
}
