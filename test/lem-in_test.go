package utils_test

import (
	"LemIn/errorHandler"
	"LemIn/fileHandler"
	"LemIn/utils"
	"bytes"
	"log"
	"os"
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
