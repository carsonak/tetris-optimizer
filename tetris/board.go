package tetris

import (
	"slices"
	"strings"
)

// Board is a square grid for placing tetrominoes.
type Board struct {
	board [][]rune
	Size  int // Width and height of the square board
}

// NewBoard creates a new empty square board.
func NewBoard(size uint) Board {
	board := Board{
		Size:  int(size),
		board: make([][]rune, size),
	}

	for i := range board.Size {
		board.board[i] = slices.Repeat([]rune{'.'}, board.Size)
	}

	return board
}

// canPlace checks if a piece fits at the given position.
func (b Board) canPlace(tet Piece, x int, y int) bool {
	if x >= b.Size || y >= b.Size {
		return false
	}

	if x+tet.Width > b.Size || y+tet.Height > b.Size {
		return false
	}

	// Check that all cells where the piece would occupy are empty
	for _, p := range tet.Pos {
		if b.board[y+p.Y][x+p.X] != '.' {
			return false
		}
	}

	return true
}

// Place places a piece on the board (returns false if placement fails).
func (b Board) Place(tet Piece, x int, y int) bool {
	if !b.canPlace(tet, x, y) {
		return false
	}

	for _, p := range tet.Pos {
		b.board[y+p.Y][x+p.X] = tet.ID
	}

	return true
}

// Remove clears a piece from the board (used for backtracking).
func (b Board) Remove(tet Piece, x int, y int) {
	for _, p := range tet.Pos {
		if b.board[y+p.Y][x+p.X] == tet.ID {
			b.board[y+p.Y][x+p.X] = '.'
		}
	}
}

// ToString returns a string representation of the board.
func (b Board) ToString() string {
	var str strings.Builder

	for _, row := range b.board {
		for _, r := range row {
			str.WriteRune(r)
		}

		str.WriteRune('\n')
	}

	return str.String()
}
