package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"

	"tetris-optimizer/tetris"
)

func makeRaw(char rune) tetris.RawPiece {
	var tet tetris.RawPiece

	for y := range 4 {
		for x := range 4 {
			tet[y][x] = char
		}
	}

	return tet
}

func makeBlock(r rune) string {
	s := string([]rune{r})
	line := strings.Repeat(s, 4) + "\n"

	return strings.Repeat(line, 4)
}

func TestParseTetrominoStream(t *testing.T) {
	testData := []struct {
		name        string
		input       string
		expected    []tetris.RawPiece
		expectError bool
		expectedMsg string
	}{
		{
			name:        "Valid: Single tetris with trailing newline",
			input:       makeBlock('1'),
			expected:    []tetris.RawPiece{makeRaw('1')},
			expectError: false,
		},
		{
			name:        "Valid: EOF after 4 lines",
			input:       strings.TrimSuffix(makeBlock('2'), "\n"),
			expected:    []tetris.RawPiece{makeRaw('2')},
			expectError: false,
		},
		{
			name:        "Valid: Multiple empty line separators",
			input:       makeBlock('A') + "\n\n\n\n" + makeBlock('B'),
			expected:    []tetris.RawPiece{makeRaw('A'), makeRaw('B')},
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
			input:       makeBlock('1') + "1234\n",
			expectError: true,
			expectedMsg: "invalid file format; Tetrominoes should be separated by blank lines",
		},
		{
			name: "Invalid: no separator between tetrominoes",
			input: makeBlock('1') + makeBlock('2'),
			expectError: true,
			expectedMsg: "invalid file format; Tetrominoes should be separated by blank lines",
		},
	}

	t.Run("nil stream", func(t *testing.T) {
		_, err := ParseTetrominoStream(nil)
		if err == nil {
			t.Error("expected error for nil stream, got nil")
		}
	})

	for _, test := range testData {
		t.Run(test.name, func(t *testing.T) {
			scanner := bufio.NewScanner(strings.NewReader(test.input))
			output, err := ParseTetrominoStream(scanner)

			if test.expectError {
				if err == nil {
					t.Errorf("expected error %q, but got nil", test.expectedMsg)
				} else if err.Error() != test.expectedMsg && test.expectedMsg != "" {
					t.Errorf("expected error %q, got %q", test.expectedMsg, err.Error())
				}

				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if !reflect.DeepEqual(output, test.expected) {
				t.Errorf("expected output:\n%+q\ngot:\n%+q", test.expected, output)
			}
		})
	}
}
