package tetro

import "errors"

type Point struct {
	X, Y int
}

type Tetromino struct {
	Pos    [4]Point // The relative coordinates of the 4 blocks
	Width  int      // Max X - Min X + 1
	Height int      // Max Y - Min Y + 1
	ID     rune     // The character to print (A, B, C...)
}

func Init(id rune, rawTetromino [4][4]rune) (Tetromino, error) {
	var tet Tetromino
	tetI := 0

	// Bottom left corner in 4x4 grid is the Origin.
	for row := len(rawTetromino) - 1; row >= 0; row-- {
		for col := len(rawTetromino[row]); col >= 0; col-- {
			if rawTetromino[row][col] == '#' {
				if tetI >= 4 {
					return Tetromino{}, errors.New("Tetromino should only have 4 pieces.")
				}

				tet.Pos[tetI] = Point{row, col}

				// if tetI > 0 {
				// 	currPos := tet.Pos[tetI]
				// 	prevPos := tet.Pos[tetI-1]

				//     if prevPos.X != currPos.X-1 && prevPos.Y != currPos.Y-1 {
				//         return  Tetromino{}, errors.New("")
				//     }
				// }

				tetI++
			}
		}
	}
	return tet, nil
}

// Helper to adjust tetromino position to the bottom left.
// This would be called after you parse the raw '#' positions
func (t *Tetromino) Normalize() {
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

	// Shift all points to bottom-left (0,0)
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
}
