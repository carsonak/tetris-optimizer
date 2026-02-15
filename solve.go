// Package main contains the backtracking solver for the tetromino packing problem.
package main

import (
	"math"
	"slices"
	"time"

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

// solveCtx holds the state for the timeout mechanism.
type solveCtx struct {
	deadline time.Time
	timedOut bool
	ops      int // Operation counter to reduce syscall overhead
}

// solve recursively places pieces using backtracking with an optional timeout.
func solve(board *tetris.Board, pieces []tetris.Piece, ctx *solveCtx) bool {
	// OPTIMIZATION: Check timeout every 1024 iterations.
	// time.Now() is a syscall; calling it every recursion is too slow.
	if ctx != nil {
		ctx.ops++
		if ctx.ops&1023 == 0 {
			if ctx.timedOut || time.Now().After(ctx.deadline) {
				ctx.timedOut = true
				return false
			}
		}
	}

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
			if solve(board, remaining, ctx) {
				return true
			}
			board.Remove(current, x, y)

			// OPTIMIZATION: If a timeout occurred deeper in the recursion,
			// break this loop immediately to unwind the stack fast.
			if ctx != nil && ctx.timedOut {
				return false
			}
		}
	}

	return false
}

// FindSmallestSquare finds the smallest square that fits all tetrominoes.
// Pieces are sorted using a widest first heuristic to trim decision branches.
// As the heuristic is not optimal for all cases, a hard timeout is used and
// the algorithm falls back onto the default ordering.
func FindSmallestSquare(tetrominoes []tetris.Piece) tetris.Board {
	tetCount := len(tetrominoes)
	minSize := minimumBoardSize(tetCount)
	maxSize := maximumBoardSize(tetCount)

	// OPTIMIZATION: Sort pieces to place the largest/hardest ones first.
	// This drastically reduces the branching factor of the recursion in some cases.
	// WARNING: This will also cripple performance of certain cases.
	sortedPieces := make([]tetris.Piece, len(tetrominoes))

	copy(sortedPieces, tetrominoes)
	slices.SortFunc(sortedPieces, func(a, b tetris.Piece) int {
		maxA := max(a.Width, a.Height)
		maxB := max(b.Width, b.Height)
		// The subtraction is reversed to cause items to be sorted in descending order.
		return maxB - maxA
	})

	useSorted := true

	for size := minSize; size <= maxSize; size++ {
		board := tetris.NewBoard(uint(size))

		if useSorted {
			ctx := &solveCtx{deadline: time.Now().Add(500 * time.Millisecond)}

			if solve(&board, sortedPieces, ctx) {
				return board
			}

			if !ctx.timedOut {
				continue
			}

			// TIMEOUT DETECTED: The sorting heuristic is a trap for this puzzle.
			// Disable it for all larger board sizes to avoid wasting 250ms/loop.
			useSorted = false
			// Reset the board as it may be in an incomplete state.
			board = tetris.NewBoard(uint(size))
		}

		// Fallback (Original Input Order)
		// passing 'nil' as context effectively disables the timeout checks inside solve
		if solve(&board, tetrominoes, nil) {
			return board
		}
	}

	return tetris.Board{}
}
