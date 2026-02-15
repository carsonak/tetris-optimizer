// Package tetris contains core data structures and validation logic for tetrominoes and the board.
package tetris

import (
	"errors"
	"fmt"
)

// RawPiece is an unvalidated 4×4 tetromino grid.
type RawPiece [4][4]byte

// Point represents a 2D coordinate (0-indexed, origin at top-left).
type Point struct {
	X, Y int
}

// Piece is a validated, normalized tetromino with blocks, dimensions, and ID.
type Piece struct {
	Width  int
	Height int
	ID     byte     // Character to print (A, B, C, ...)
	Pos    [4]Point // Relative coordinates of the 4 blocks
}

//////////////////// STATIC FUNCTIONS ////////////////////

// countNeighbors returns the count of orthogonally adjacent blocks.
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

// normalize shifts the tetromino to start at (0,0) and calculates bounds.
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

// Init validates and normalizes a RawPiece (4 blocks, neighbour count 6 or 8).
func Init(rawTet RawPiece, id byte) (Piece, error) {
	var tet Piece
	neighbours := 0
	blockCount := 0

	for y, row := range rawTet {
		for x, char := range row {
			if char == '.' {
				continue
			}

			if char != '#' {
				return Piece{}, fmt.Errorf("unrecognised character '%c'", char)
			}

			if blockCount >= 4 {
				return Piece{}, errors.New("tetromino should have 4 blocks.")
			}

			pos := Point{X: x, Y: y}
			neighbours += countNeighbors(pos, rawTet)
			tet.Pos[blockCount] = pos
			blockCount++
		}
	}

	if blockCount != 4 {
		return Piece{}, errors.New("tetromino should have 4 blocks")
	}

	if neighbours != 6 && neighbours != 8 {
		return Piece{}, errors.New("invalid tetromino")
	}

	tet.ID = id
	tet.normalize()
	return tet, nil
}
