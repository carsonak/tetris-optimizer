package tetris

import (
	"fmt"
	"slices"
)

// Board represents a square grid where tetrominoes are placed.
// Empty cells are ' ', occupied cells are marked with piece IDs (A-Z).
type Board struct {
	board [][]rune
	Size  int // Width and height of the square
}

// NewBoard creates a new empty square board of the given size.
func NewBoard(size uint) Board {
	board := Board{
		Size:  int(size),
		board: make([][]rune, size),
	}

	// Initialize each row with empty spaces
	for i := range board.Size {
		board.board[i] = slices.Repeat([]rune{' '}, board.Size)
	}

	return board
}

// canPlace checks if a piece fits at the given position without overlap or bounds issues.
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

// Place places a piece on the board if possible. Returns true on success, false on failure.
func (b Board) Place(tet Piece, x int, y int) bool {
	if !b.canPlace(tet, x, y) {
		return false
	}

	for _, p := range tet.Pos {
		b.board[y+p.Y][x+p.X] = tet.ID
	}

	return true
}

// Remove clears the piece from the board at the given position (for backtracking).
func (b Board) Remove(tet Piece, x int, y int) {
	for _, p := range tet.Pos {
		if b.board[y+p.Y][x+p.X] == tet.ID {
			b.board[y+p.Y][x+p.X] = ' '
		}
	}
}

// Print outputs the board to stdout.
func (b Board) Print() {
	for _, row := range b.board {
		for _, r := range row {
			fmt.Printf("%c", r)
		}

		fmt.Println()
	}
}
