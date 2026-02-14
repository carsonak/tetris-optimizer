// Package main contains the backtracking algorithm that solves the tetromino packing problem.
package main

import (
	"math"

	"tetris-optimizer/tetris"
)

// minimumBoardSize returns the theoretical minimum board size (⌈√(count×4)⌉),
// assuming perfect packing with no wasted space.
func minimumBoardSize(tetrominoCount int) int {
	cellCount := float64(tetrominoCount * 4)
	root := math.Sqrt(cellCount)
	ceil := math.Ceil(root)

	return int(ceil)
}

// maximumBoardSize returns the upper search bound (⌈√(count×16)⌉),
// accounting for worst-case spacing between pieces.
func maximumBoardSize(tetrominoCount int) int {
	cellCount := float64(tetrominoCount * 16)
	root := math.Sqrt(cellCount)
	ceil := math.Ceil(root)

	return int(ceil)
}

// solve recursively places tetrominoes using backtracking.
// Returns true if all pieces fit on the board, false otherwise.
// The board is modified in place and restored via Remove on backtrack.
func solve(board tetris.Board, pieces []tetris.Piece) bool {
	if len(pieces) == 0 {
		return true
	}

	current := pieces[0]
	remaining := pieces[1:]

	// Try all valid positions for the current piece
	for y := 0; y <= board.Size-current.Height; y++ {
		for x := 0; x <= board.Size-current.Width; x++ {
			if !board.Place(current, x, y) {
				continue
			}

			if solve(board, remaining) {
				return true
			}

			board.Remove(current, x, y)
		}
	}

	return false
}

// FindSmallestSquare finds the smallest square that fits all tetrominoes.
// Tests board sizes from minimum to maximum, returning the first valid solution.
func FindSmallestSquare(tetrominoes []tetris.Piece) tetris.Board {
	tetCount := len(tetrominoes)
	minSize := minimumBoardSize(tetCount)
	maxSize := maximumBoardSize(tetCount)
	board := tetris.NewBoard(uint(minSize))

	for size := minSize + 1; size < maxSize; size++ {
		if solve(board, tetrominoes) {
			return board
		}

		board = tetris.NewBoard(uint(size))
	}

	return tetris.Board{}
}
