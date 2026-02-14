package tetris

import (
	"errors"
	"fmt"
)

// Unverified tetromino on a 4x4 grid
type RawPiece [4][4]rune

type Point struct {
	X, Y int
}

// Internal representation of a tetromino piece
type Piece struct {
	Pos    [4]Point // The relative coordinates of the 4 blocks
	Width  int
	Height int
	ID     rune // The character to print (A, B, C...)
}

//////////////////// STATIC FUNCTIONS ////////////////////

// Count blocks around another block in a tetromino.
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

// Adjust tetromino position to the top-left of the grid.
func (t *Piece) normalize() {
	minX, minY := t.Pos[0].X, t.Pos[0].Y

	// Find offsets from the X and Y axes.
	for _, p := range t.Pos {
		if p.X < minX {
			minX = p.X
		}

		if p.Y < minY {
			minY = p.Y
		}
	}

	maxX, maxY := 0, 0

	// Shift edges of the tetromino onto the axes
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


// Strip unnecessary information from a raw tetromino.
func Init(rawTet RawPiece, id rune) (Piece, error) {
	var tet Piece
	tetroBlocks := 0

	// Assume the 4x4 grid is the $th Quadrant of a cartesian plane,
	// top-left corner of the grid is the origin.
	for y, row := range rawTet {
		for x, char := range row {
			switch char {
			case ' ':
				continue
			case '#':
				if tetroBlocks >= 4 {
					return Piece{}, errors.New("Piece should have 4 blocks.")
				}

				// Although we are using the 4th quadrant, we leave Y positive to make
				// calculations easier simpler.
				pos := Point{X: x, Y: y}
				neighbors := countNeighbors(pos, rawTet)

				if neighbors < 1 || neighbors > 3 {
					return Piece{}, errors.New("invalid Piece")
				}

				tet.Pos[tetroBlocks] = pos
				tetroBlocks++
			default:
				return Piece{}, fmt.Errorf("unrecognised character '%c'", char)
			}
		}
	}

	if tetroBlocks != 4 {
		return Piece{}, errors.New("Piece should have 4 blocks")
	}

	tet.ID = id
	tet.normalize()
	return tet, nil
}
