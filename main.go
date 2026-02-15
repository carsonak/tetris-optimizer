// Package main handles CLI, file I/O, and puzzle solving orchestration.
package main

import (
	"bufio"
	"fmt"
	"os"

	"tetris-optimizer/tetris"
)

// initTetrominoPieces converts raw tetrominoes to validated pieces with IDs A-Z.
func initTetrominoPieces(rawTetrominoes []tetris.RawPiece) ([]tetris.Piece, error) {
	idLimit := int('Z'-'A') + 1

	if len(rawTetrominoes) > idLimit {
		return nil, fmt.Errorf("cannot process more than %d tetrominoes", idLimit)
	}

	var pieces []tetris.Piece
	id := byte('A')

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

// main parses input file, validates tetrominoes, and prints the solution.
func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "ERROR: USAGE: %s tetromino_file\n", os.Args[0])
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
