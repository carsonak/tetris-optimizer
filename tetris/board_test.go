package tetris

import "testing"

func TestBoardPlaceRemove(t *testing.T) {
	board := NewBoard(4)
	piece, err := Init(makeRaw(
		"##..",
		"##..",
		"....",
		"....",
	), 'A')
	if err != nil {
		t.Fatalf("unexpected init error: %v", err)
	}

	if !board.Place(piece, 0, 0) {
		t.Fatal("expected initial placement to succeed")
	}

	if board.Place(piece, 1, 1) {
		t.Fatal("expected overlapping placement to fail")
	}

	board.Remove(piece, 0, 0)

	if !board.Place(piece, 2, 2) {
		t.Fatal("expected placement after removal to succeed")
	}
}

func TestBoardCanPlaceBounds(t *testing.T) {
	board := NewBoard(4)
	piece, err := Init(makeRaw(
		"##..",
		"##..",
		"....",
		"....",
	), 'A')
	if err != nil {
		t.Fatalf("unexpected init error: %v", err)
	}

	if board.canPlace(piece, 3, 3) {
		t.Fatal("expected out-of-bounds placement to fail")
	}

	if !board.canPlace(piece, 2, 2) {
		t.Fatal("expected in-bounds placement to succeed")
	}
}
