package main

import (
	"bufio"
	"errors"
)

func ParseTetrominoStream(scanner *bufio.Scanner) (output [][4][4]rune, err error) {
	rowCount := 0
	rawTetromino := [4][4]rune{}

	for scanner.Scan() {
		if rowCount > 3 {
			output = append(output, rawTetromino)
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

		copy(rawTetromino[rowCount][:], []rune(line))
		rowCount++
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return
}
