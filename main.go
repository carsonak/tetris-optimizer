// Package main handles command-line parsing, file I/O, and orchestrates the puzzle solving process.
package main

import (
	"bufio"
	"fmt"
	"os"

	"tetris-optimizer/tetris"
)

// Convert raw tetrominoes to validated pieces with unique IDs (A-Z).
// Returns an error if any raw piece is invalid or if there are more than 26 pieces.
func initTetrominoPieces(rawTetrominoes []tetris.RawPiece) ([]tetris.Piece, error) {
	idLimit := int('Z'-'A') + 1

	if len(rawTetrominoes) > idLimit {
		return nil, fmt.Errorf("cannot process more than %d tetrominoes", idLimit)
	}

	var pieces []tetris.Piece
	id := 'A'

	for _, raw := range rawTetrominoes {
		p, err := tetris.Init(raw, id)
		if err != nil {
			return nil, err
		}

		pieces = append(pieces, p)
		id++
	}

	return pieces, nil
}

// Parses a tetromino file, validates tetrominoes,
// find the smallest square to fit all the tetrominoes
// and print the solution.
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
		os.Exit(1)
	}

	fmt.Print(FindSmallestSquare(tetrominoes).ToString())
}
