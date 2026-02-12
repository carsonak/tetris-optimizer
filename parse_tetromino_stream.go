package main

import (
	"bufio"
	"errors"

	"tetris-optimiser/tetris"
)

func ParseTetrominoStream(scanner *bufio.Scanner) (output []tetris.RawPiece, err error) {
	if scanner == nil {
		return nil, errors.New("scanner should not be nil")
	}

	rowCount := 0
	tet := tetris.RawPiece{}

	for scanner.Scan() {
		if rowCount > 3 {
			output = append(output, tet)
			rowCount = 0
		}

		line := scanner.Text()

		if len(line) < 1 {
			if rowCount == 0 {
				continue
			}

			return nil, errors.New("invalid file format; Tetromino should have 4 rows")
		}

		if len(line) != 4 {
			return nil, errors.New("invalid file format; Tetromino should have 4 columns")
		}

		copy(tet[rowCount][:], []rune(line))
		rowCount++
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	if rowCount == 4 {
		output = append(output, tet)
		rowCount = 0
	}

	if rowCount != 0 {
		return nil, errors.New("invalid file format; Tetromino should have 4 rows")
	}

	return
}
