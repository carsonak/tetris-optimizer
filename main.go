// Package main handles command-line parsing, file I/O, and orchestrates the puzzle solving process.
package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"tetris-optimizer/tetris"
)

// initTetrominoPieces converts raw tetrominoes to validated pieces with unique IDs (A-Z).
// Returns an error if any raw piece is invalid or if there are more than 26 pieces.
func initTetrominoPieces(rawTetrominoes []tetris.RawPiece) ([]tetris.Piece, error) {
	var pieces []tetris.Piece

	for i, raw := range rawTetrominoes {
		id := rune('A' + i)

		if id > 'Z' {
			return nil, errors.New("ERROR: cannot process more than 26 tetrominoes")
		}

		piece, err := tetris.Init(raw, id)
		if err != nil {
			return nil, err
		}

		pieces = append(pieces, piece)
	}

	return pieces, nil
}

// main expects exactly one argument: path to a tetromino file.
// Parses the file, validates tetrominoes, finds the smallest square, and prints the solution.
func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "ERROR: USAGE: %s tetromino_file", os.Args[0])
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
		os.Exit(1)
	}

	defer file.Close()
	rawTetrominoes, err := ParseTetrominoStream(bufio.NewScanner(file))
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
		os.Exit(1)
	}

	tetrominoes, err := initTetrominoPieces(rawTetrominoes)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
	}

	FindSmallestSquare(tetrominoes).Print()
}
