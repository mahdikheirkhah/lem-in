package utils_test

import (
	"LemIn/errorHandler"
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
