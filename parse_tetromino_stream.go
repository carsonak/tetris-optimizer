// Package main contains file parsing for tetromino input.
package main

import (
	"bufio"
	"errors"

	"tetris-optimizer/tetris"
)

// ParseTetrominoStream reads tetrominoes from a scanner (4 rows × 4 cols, separated by blanks).
func ParseTetrominoStream(scanner *bufio.Scanner) (pieces []tetris.RawPiece, err error) {
	if scanner == nil {
		return nil, errors.New("scanner should not be nil")
	}

	var current tetris.RawPiece
	rowCount := 0

	for scanner.Scan() {
		line := scanner.Text()

		// Allow back-to-back tetrominoes without a blank separator.
		if rowCount == 4 {
			pieces = append(pieces, current)
			current = tetris.RawPiece{}
			rowCount = 0

			if len(line) != 0 {
				return nil, errors.New("invalid file format; Tetrominoes should be separated by blank lines")
			}

			continue
		}

		if len(line) == 0 {
			if rowCount == 0 {
				continue // Allow for several blank lines between tetrominoes.
			}

			return nil, errors.New("invalid file format; Tetromino should have 4 rows")
		}

		if len(line) != 4 {
			return nil, errors.New("invalid file format; Tetromino should have 4 columns")
		}

		copy(current[rowCount][:], []rune(line))
		rowCount++
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	// Add final tetromino if present
	if rowCount == 4 {
		pieces = append(pieces, current)
	} else if rowCount > 0 {
		return nil, errors.New("invalid file format; Tetromino should have 4 rows")
	}

	return pieces, nil
}
