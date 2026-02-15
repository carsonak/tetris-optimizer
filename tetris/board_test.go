package tetris

import (
	"testing"
)

func TestNewBoard(t *testing.T) {
	board := NewBoard(5)

	if board.Size != 5 {
		t.Errorf("expected size 5, got %d", board.Size)
	}

	// Verify all cells are empty
	for y := range board.Size {
		for x := range board.Size {
			if board.board[y][x] != '.' {
				t.Errorf("expected empty cell at (%d,%d), got %c", x, y, board.board[y][x])
			}
		}
	}
}

var OPiece = Piece{
	Pos:    [4]Point{{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 0, Y: 1}, {X: 1, Y: 1}},
	Width:  2,
	Height: 2,
	ID:     'A',
}

func TestRemove(t *testing.T) {
	board := NewBoard(4)

	board.Place(OPiece, 0, 0)
	board.Remove(OPiece, 0, 0)
	// After removal, should be able to place at same location
	if !board.CanPlace(OPiece, 0, 0) {
		t.Fatal("expected placement after removal to succeed")
	}
}

func TestCanPlace(t *testing.T) {
	board := NewBoard(4)
	testData := []struct {
		name     string
		x, y     int
		expected bool
	}{
		{"valid top-left", 0, 0, true},
		{"valid middle", 1, 1, true},
		{"valid bottom-right", 2, 2, true},
		{"invalid out-of-bounds x", 3, 0, false},
		{"invalid out-of-bounds y", 0, 3, false},
		{"invalid both out-of-bounds", 3, 3, false},
	}

	for _, test := range testData {
		t.Run(test.name, func(t *testing.T) {
			got := board.CanPlace(OPiece, test.x, test.y)
			if got != test.expected {
				t.Errorf("canPlace(%d, %d) = %v, want %v", test.x, test.y, got, test.expected)
			}
		})
	}
}

func TestToString(t *testing.T) {
	board := NewBoard(4)
	expected := "....\n" +
		".AA.\n" +
		".AA.\n" +
		"....\n"

	board.Place(OPiece, 1, 1)
	output := board.ToString()

	if output != expected {
		t.Fatalf("expected:\n%s\ngot:\n%s", expected, output)
	}
}
