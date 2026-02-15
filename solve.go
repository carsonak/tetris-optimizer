// Package main contains the backtracking solver for the tetromino packing problem.
package main

import (
	"math"
	"slices"

	"tetris-optimizer/tetris"
)

// minimumBoardSize returns the theoretical minimum size: ⌈√(count×4)⌉.
func minimumBoardSize(tetrominoCount int) int {
	cellCount := float64(tetrominoCount * 4)
	root := math.Sqrt(cellCount)
	ceil := math.Ceil(root)

	return int(ceil)
}

// maximumBoardSize returns the upper search bound: ⌈√(count×16)⌉.
func maximumBoardSize(tetrominoCount int) int {
	cellCount := float64(tetrominoCount * 16)
	root := math.Sqrt(cellCount)
	ceil := math.Ceil(root)

	return int(ceil)
}

// solve recursively places pieces using backtracking.
func solve(board *tetris.Board, pieces []tetris.Piece) bool {
	if len(pieces) == 0 {
		return true
	}

	current := pieces[0]
	remaining := pieces[1:]

	// Try all valid positions for the current piece
	for y := 0; y <= board.Size-current.Height; y++ {
		for x := 0; x <= board.Size-current.Width; x++ {
			if !board.CanPlace(current, x, y) {
				continue
			}

			board.Place(current, x, y)
			if solve(board, remaining) {
				return true
			}

			board.Remove(current, x, y)
		}
	}

	return false
}

// FindSmallestSquare finds the smallest square that fits all tetrominoes.
func FindSmallestSquare(tetrominoes []tetris.Piece) tetris.Board {
	tetCount := len(tetrominoes)
	minSize := minimumBoardSize(tetCount)
	maxSize := maximumBoardSize(tetCount)

	// OPTIMIZATION: Sort pieces to place the largest/hardest ones first.
	// This drastically reduces the branching factor of the recursion in some cases.
	// WARNING: This will also cripple performance of certain cases.
	slices.SortFunc(tetrominoes, func(a, b tetris.Piece) int {
		maxA := max(a.Width, a.Height)
		maxB := max(b.Width, b.Height)
		// The subtraction is reversed to cause items to be sorted in descending order.
		return maxB - maxA
	})

	for size := minSize; size <= maxSize; size++ {
		board := tetris.NewBoard(uint(size))

		if solve(&board, tetrominoes) {
			return board
		}
	}

	return tetris.Board{}
}
