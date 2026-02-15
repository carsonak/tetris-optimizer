package tetris

import (
	"slices"
	"strings"
)

// Board is a square grid for placing tetrominoes.
type Board struct {
	Size  int // Width and height of the square board
	board [][]byte
}

// NewBoard creates a new empty square board.
func NewBoard(size uint) Board {
	b := Board{
		Size:  int(size),
		board: make([][]byte, size),
	}

	// OPTIMISATION: Allocating all the board memory in one continuous block
	// improves cache locality
	backingMem := slices.Repeat([]byte{'.'}, b.Size*b.Size)

	for i := range b.Size {
		b.board[i] = backingMem[i*b.Size : (i+1)*b.Size]
	}

	return b
}

// CanPlace checks if a piece fits at the given position.
func (b *Board) CanPlace(tet Piece, x, y int) bool {
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

// Place places a piece on the board.
func (b *Board) Place(tet Piece, x int, y int) {
	for _, p := range tet.Pos {
		b.board[y+p.Y][x+p.X] = tet.ID
	}
}

// Remove clears a piece from the board (used for backtracking).
func (b *Board) Remove(tet Piece, x, y int) {
	for _, p := range tet.Pos {
		b.board[y+p.Y][x+p.X] = '.'
	}
}

// ToString returns a string representation of the board.
func (b Board) ToString() string {
	var str strings.Builder

	for _, row := range b.board {
		for _, r := range row {
			str.WriteByte(r)
		}

		str.WriteRune('\n')
	}

	return str.String()
}
