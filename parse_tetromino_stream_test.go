package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
	"tetris-optimiser/tetromino"
)

func TestParseTetrominoStream(t *testing.T) {
	t.Run("nil stream", func(t *testing.T) {
		_, err := ParseTetrominoStream(nil)

		if err == nil {
			t.Error("nil stream should return an error")
		}
	})

	testsData := []struct {
		name     string
		input    *bufio.Scanner
		err      error
		expected []tetromino.Raw
	}{
		{
			name: "", input: bufio.NewScanner(strings.NewReader(strings.Repeat("1234\n", 4) + "\n")),
			err: nil, expected: []tetromino.Raw{{{'1', '2', '3', '4'}, {'1', '2', '3', '4'}, {'1', '2', '3', '4'}, {'1', '2', '3', '4'}}},
		},
	}

	for _, test := range testsData {
		t.Run(test.name, func(t *testing.T) {
			output, err := ParseTetrominoStream(test.input)

			if err != test.err {
				t.Errorf("expected error: %q\ngot error: %q", test.err, err)
				return
			}

			if !reflect.DeepEqual(output, test.expected) {
				t.Errorf("expected: %+v\ngot: %+v", test.expected, output)
			}
		})
	}
}
