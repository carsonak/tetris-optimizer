package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"

	"tetris-optimiser/tetromino"
)

func TestParseTetrominoStream(t *testing.T) {
	makeRaw := func(char rune) tetromino.Raw {
		var tet tetromino.Raw

		for y := range 4 {
			for x := range 4 {
				tet[y][x] = char
			}
		}

		return tet
	}

	makeBlock := func(r rune) string {
		s := string([]rune{r})
		line := strings.Repeat(s, 4) + "\n"

		return strings.Repeat(line, 4)
	}

	tests := []struct {
		name        string
		input       string
		expected    []tetromino.Raw
		expectError bool
		expectedMsg string
	}{
		{
			name:        "Valid: Single tetromino with trailing newline",
			input:       makeBlock('1'),
			expected:    []tetromino.Raw{makeRaw('1')},
			expectError: false,
		},
		{
			name:        "Valid: EOF after 4 lines",
			input:       strings.TrimSuffix(makeBlock('2'), "\n"),
			expected:    []tetromino.Raw{makeRaw('2')},
			expectError: false,
		},
		{
			name:        "Valid: Multiple tetrominoes",
			input:       makeBlock('A') + makeBlock('B'),
			expected:    []tetromino.Raw{makeRaw('A'), makeRaw('B')},
			expectError: false,
		},
		{
			name:        "Valid: Multiple empty line separators",
			input:       makeBlock('A') + "\n\n\n\n" + makeBlock('B'),
			expected:    []tetromino.Raw{makeRaw('A'), makeRaw('B')},
			expectError: false,
		},
		{
			name:        "Invalid: Columns < 4",
			input:       "123\n" + strings.Repeat("1234\n", 3),
			expectError: true,
			expectedMsg: "invalid file format; Tetromino should have 4 columns",
		},
		{
			name:        "Invalid: Columns > 4",
			input:       "12345\n" + strings.Repeat("1234\n", 3),
			expectError: true,
			expectedMsg: "invalid file format; Tetromino should have 4 columns",
		},
		{
			name:        "Invalid: EOF after 3 lines",
			input:       strings.Repeat("1234\n", 3),
			expectError: true,
			expectedMsg: "invalid file format; Tetromino should have 4 rows",
		},
		{
			name:        "Invalid: Rows < 4",
			input:       strings.Repeat("1234\n", 2) + "\n1234\n",
			expectError: true,
			expectedMsg: "invalid file format; Tetromino should have 4 rows",
		},
		{
			name:        "Invalid: Rows > 4",
			input:       makeBlock('1') + "1234\n\n",
			expectError: true,
			expectedMsg: "invalid file format; Tetromino should have 4 rows",
		},
	}

	t.Run("nil stream", func(t *testing.T) {
		_, err := ParseTetrominoStream(nil)
		if err == nil {
			t.Error("expected error for nil stream, got nil")
		}
	})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scanner := bufio.NewScanner(strings.NewReader(tt.input))
			output, err := ParseTetrominoStream(scanner)

			if tt.expectError {
				if err == nil {
					t.Errorf("expected error %q, but got nil", tt.expectedMsg)
				} else if err.Error() != tt.expectedMsg && tt.expectedMsg != "" {
					t.Errorf("expected error %q, got %q", tt.expectedMsg, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if !reflect.DeepEqual(output, tt.expected) {
					t.Errorf("expected output:\n%+q\ngot:\n%+q", tt.expected, output)
				}
			}
		})
	}
}
