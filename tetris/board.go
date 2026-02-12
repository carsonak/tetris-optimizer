package tetris

import (
	"fmt"
	"slices"
)

type Board struct {
	board [][]rune
	Size  int
}

func NewBoard(size uint) Board {
	board := Board{
		Size:  int(size),
		board: make([][]rune, size),
	}

	for i := range board.board {
		board.board[i] = slices.Repeat([]rune{' '}, board.Size)
	}

	return board
}

func (b Board) canPlace(tet Piece, x int, y int) bool {
	if x >= b.Size || y >= b.Size {
		return false
	}

	if x+tet.Width >= b.Size || y+tet.Height >= b.Size {
		return false
	}

	for _, p := range tet.Pos {
		if b.board[x+p.X][y+p.Y] != ' ' {
			return false
		}
	}

	return true
}

func (b *Board) Place(tet Piece, x int, y int) bool {
	if !b.canPlace(tet, x, y) {
		return false
	}

	for _, p := range tet.Pos {
		b.board[x+p.X][y+p.Y] = tet.ID
	}

	return true
}

func (b Board) Print() {
	for _, row := range b.board {
		for _, r := range row {
			fmt.Print(r)
		}

		fmt.Println()
	}
}
