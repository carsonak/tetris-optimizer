// Package main contains the file parsing functionality for extracting tetrominoes.
package main

import (
	"bufio"
	"errors"

	"tetris-optimizer/tetris"
)

// ParseTetrominoStream reads and validates tetrominoes from a file scanner.
// Each tetromino must be 4x4, with '#' for blocks and '.' or ' ' for empty space.
// Tetrominoes are separated by blank lines.
func ParseTetrominoStream(scanner *bufio.Scanner) ([]tetris.RawPiece, error) {
	if scanner == nil {
		return nil, errors.New("scanner should not be nil")
	}

	var pieces []tetris.RawPiece
	var current tetris.RawPiece
	rowCount := 0

	for scanner.Scan() {
		line := scanner.Text()

		// Allow back-to-back tetrominoes without a blank separator.
		if rowCount == 4 {
			pieces = append(pieces, current)
			current = tetris.RawPiece{}
			rowCount = 0
		}

		// Blank line signals end of current tetromino
		if len(line) == 0 {
			if rowCount == 0 {
				continue // Skip leading/trailing blank lines
			}
			return nil, errors.New("invalid file format; Tetromino should have 4 rows")
		}

		// Validate row dimensions
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
