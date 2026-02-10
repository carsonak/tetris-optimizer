package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"tetris-optimiser/tetromino"
)

func initTetrominoPieces(rawTetrominoes []tetromino.Raw) ([]tetromino.Piece, error) {
	var tetrominoes []tetromino.Piece

	for i, t := range rawTetrominoes {
		id := rune('A' + i)

		if id > 'Z' { // Since IDs are only uppercase letters, number of tetrominoes is capped at 26
			return nil, errors.New("ERROR: cannot process more than 26 tetrominoes")
		}

		piece, err := tetromino.Init(t, id)
		if err != nil {
			return nil, err
		}

		tetrominoes = append(tetrominoes, piece)
	}

	return tetrominoes, nil
}

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

	_, err = initTetrominoPieces(rawTetrominoes)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
	}
}
