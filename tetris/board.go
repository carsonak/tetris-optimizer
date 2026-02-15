package tetris

import (
	"slices"
	"strings"
)

// Represents a square grid where tetrominoes are placed.
// Empty cells are ' ', occupied cells are marked with piece IDs (A-Z).
type Board struct {
	board [][]rune
	Size  int // Width and height of the square
}

// Create a new empty square board of the given size.
func NewBoard(size uint) Board {
	board := Board{
		Size:  int(size),
		board: make([][]rune, size),
	}

	for i := range board.Size {
		board.board[i] = slices.Repeat([]rune{' '}, board.Size)
	}

	return board
}

// Check if a piece fits at the given position without overlap or bounds issues.
func (b Board) canPlace(tet Piece, x int, y int) bool {
	if x >= b.Size || y >= b.Size {
		return false
	}

	if x+tet.Width > b.Size || y+tet.Height > b.Size {
		return false
	}

	// Check that all cells where the piece would occupy are empty
	for _, p := range tet.Pos {
		if b.board[y+p.Y][x+p.X] != ' ' {
			return false
		}
	}

	return true
}

// Insert a piece on the board if possible.
// Returns true on success, false on failure.
func (b Board) Place(tet Piece, x int, y int) bool {
	if !b.canPlace(tet, x, y) {
		return false
	}

	for _, p := range tet.Pos {
		b.board[y+p.Y][x+p.X] = tet.ID
	}

	return true
}

// Clear a piece from the board at the given position.
func (b Board) Remove(tet Piece, x int, y int) {
	for _, p := range tet.Pos {
		if b.board[y+p.Y][x+p.X] == tet.ID {
			b.board[y+p.Y][x+p.X] = ' '
		}
	}
}

// Return a string representation of the board.
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
