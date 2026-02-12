package main

import (
	"math"

	"tetris-optimiser/tetris"
)

func minimumBoardSize(tetrominoCount int) int {
	cellCount := float64(tetrominoCount * 4)
	root := math.Sqrt(cellCount)
	ceil := math.Ceil(root)

	return int(ceil)
}

func maximumBoardSize(tetrominoCount int) int {
	cellCount := float64(tetrominoCount * 16)
	root := math.Sqrt(cellCount)
	ceil := math.Ceil(root)

	return int(ceil)
}

func solve(board *tetris.Board, tetrominoes []tetris.Piece) bool {
	if len(tetrominoes) < 1 {
		return true
	}

	tet := tetrominoes[0]

	for y := 0; y < board.Size-tet.Height; y++ {
		for x := 0; x < board.Size-tet.Width; x++ {
			if board.Place(tet, x, y) {
				solve(board, tetrominoes[1:])
			}
		}
	}

	return false
}

func FindSmallestSquare(tetrominoes []tetris.Piece) tetris.Board {
	tetCount := len(tetrominoes)
	minSize := minimumBoardSize(tetCount)
	maxSize := maximumBoardSize(tetCount)
	board := tetris.NewBoard(uint(minSize))

	for size := minSize + 1; size < maxSize; size++ {
		if solve(&board, tetrominoes) {
			return board
		}

		board = tetris.NewBoard(uint(size))
	}

	return tetris.Board{}
}
