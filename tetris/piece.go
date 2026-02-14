// Package tetris contains the core data structures for representing tetrominoes
// and the game board, including validation and manipulation logic.
package tetris

import (
	"errors"
	"fmt"
)

// RawPiece is an unverified 4x4 tetromino with '#' for blocks, other chars for empty space.
type RawPiece [4][4]rune

// Point represents a 2D coordinate.
type Point struct {
	X, Y int
}

// Piece is a validated and normalized tetromino with relative block coordinates,
// Width/Height bounds, and an ID for printing (A-Z).
type Piece struct {
	Pos    [4]Point // Relative coordinates of the 4 blocks
	Width  int
	Height int
	ID     rune // Character to print (A, B, C, ...)
}

//////////////////// STATIC FUNCTIONS ////////////////////

// countNeighbors returns the number of orthogonally adjacent blocks (0-4).
// Used to validate that tetromino blocks are properly connected.
func countNeighbors(pos Point, tet RawPiece) int {
	neighbors := 0

	if pos.X > 0 && tet[pos.Y][pos.X-1] == '#' { // check left
		neighbors++
	}

	if pos.X < len(tet[pos.Y])-1 && tet[pos.Y][pos.X+1] == '#' { // check right
		neighbors++
	}

	if pos.Y > 0 && tet[pos.Y-1][pos.X] == '#' { // check down
		neighbors++
	}

	if pos.Y < len(tet)-1 && tet[pos.Y+1][pos.X] == '#' { // check up
		neighbors++
	}

	return neighbors
}

//////////////////// PRIVATE METHODS ////////////////////

// normalize shifts the piece so its top-left block is at (0, 0) and sets Width/Height.
func (t *Piece) normalize() {
	// Find minimum X and Y coordinates
	minX, minY := t.Pos[0].X, t.Pos[0].Y
	for _, p := range t.Pos {
		if p.X < minX {
			minX = p.X
		}
		if p.Y < minY {
			minY = p.Y
		}
	}

	maxX, maxY := 0, 0
	// Shift coordinates to start at (0, 0) and compute bounds
	for i := range t.Pos {
		t.Pos[i].X -= minX
		t.Pos[i].Y -= minY

		if t.Pos[i].X > maxX {
			maxX = t.Pos[i].X
		}

		if t.Pos[i].Y > maxY {
			maxY = t.Pos[i].Y
		}
	}

	t.Width = maxX + 1
	t.Height = maxY + 1
}

//////////////////// PUBLIC METHODS ////////////////////

// Init validates a RawPiece and returns a normalized Piece with the given ID.
// Each block must be connected to 1-3 neighbors (orthogonally).
// Returns an error if the piece is invalid (wrong count, disconnected blocks, etc).
func Init(rawTet RawPiece, id rune) (Piece, error) {
	var tet Piece
	blockCount := 0

	for y, row := range rawTet {
		for x, char := range row {
			if char == ' ' {
				continue
			}

			if char != '#' {
				return Piece{}, fmt.Errorf("unrecognised character '%c'", char)
			}

			if blockCount >= 4 {
				return Piece{}, errors.New("Piece should have 4 blocks.")
			}

			pos := Point{X: x, Y: y}
			if neighbors := countNeighbors(pos, rawTet); neighbors < 1 || neighbors > 3 {
				return Piece{}, errors.New("invalid Piece")
			}

			tet.Pos[blockCount] = pos
			blockCount++
		}
	}

	if blockCount != 4 {
		return Piece{}, errors.New("Piece should have 4 blocks")
	}

	tet.ID = id
	tet.normalize()
	return tet, nil
}
