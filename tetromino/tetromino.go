package tetromino

import (
	"errors"
	"fmt"
)

type Raw [4][4]rune

type Point struct {
	X, Y int
}

type Piece struct {
	Pos    [4]Point // The relative coordinates of the 4 blocks
	Width  int
	Height int
	ID     rune // The character to print (A, B, C...)
}

func countNeighbors(pos Point, tet Raw) int {
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

func Init(rawTet Raw, id rune) (Piece, error) {
	var tet Piece
	tetroBlocks := 0

	// Bottom left corner in 4x4 grid is the Origin.
	for y, row := range rawTet {
		for x, char := range row {
			switch char {
			case ' ':
				continue
			case '#':
				if tetroBlocks >= 4 {
					return Piece{}, errors.New("Piece should have 4 blocks.")
				}

				pos := Point{x, y}
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

	return tet, nil
}

// Helper to adjust tetromino position to the bottom left.
// This would be called after you parse the raw '#' positions
func (t *Piece) Normalize() {
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

	// Shift tetromino onto the axes
	for i := range t.Pos {
		t.Pos[i].X -= minX
		t.Pos[i].Y -= minY

		// Track max dimensions for Width/Height calculation
		if t.Pos[i].X > maxX {
			maxX = t.Pos[i].X
		}

		if t.Pos[i].Y > maxY {
			maxY = t.Pos[i].Y
		}
	}

	t.Width = maxX + 1
	t.Height = maxY + 1

	// Shift Piece to 4th Quadrant
	for i := range t.Pos {
		t.Pos[i].Y -= t.Height
	}
}
