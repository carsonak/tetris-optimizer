package main

import (
	"strings"
	"testing"

	"tetris-optimizer/tetris"
)

func makeRaw(rows ...string) tetris.RawPiece {
	var grid tetris.RawPiece

	for y := 0; y < 4; y++ {
		row := "...."
		if y < len(rows) {
			row = rows[y]
		}
		row = strings.ReplaceAll(row, ".", " ")
		copy(grid[y][:], []rune(row))
	}

	return grid
}

func mustInit(t *testing.T, raw tetris.RawPiece, id rune) tetris.Piece {
	t.Helper()
	piece, err := tetris.Init(raw, id)
	if err != nil {
		t.Fatalf("unexpected init error: %v", err)
	}

	return piece
}

func TestFindSmallestSquareSingleO(t *testing.T) {
	piece := mustInit(t, makeRaw(
		"##..",
		"##..",
		"....",
		"....",
	), 'A')

	board := FindSmallestSquare([]tetris.Piece{piece})
	if board.Size != 2 {
		t.Fatalf("expected board size 2, got %d", board.Size)
	}
}

func TestFindSmallestSquareSingleIHorizontal(t *testing.T) {
	piece := mustInit(t, makeRaw(
		"####",
		"....",
		"....",
		"....",
	), 'A')

	board := FindSmallestSquare([]tetris.Piece{piece})
	if board.Size != 4 {
		t.Fatalf("expected board size 4, got %d", board.Size)
	}
}

func TestFindSmallestSquareTwoO(t *testing.T) {
	pieceA := mustInit(t, makeRaw(
		"##..",
		"##..",
		"....",
		"....",
	), 'A')
	pieceB := mustInit(t, makeRaw(
		"##..",
		"##..",
		"....",
		"....",
	), 'B')

	board := FindSmallestSquare([]tetris.Piece{pieceA, pieceB})
	if board.Size != 4 {
		t.Fatalf("expected board size 4, got %d", board.Size)
	}
}

func TestSolvePlacesAllPieces(t *testing.T) {
	pieceA := mustInit(t, makeRaw(
		"##..",
		"##..",
		"....",
		"....",
	), 'A')
	pieceB := mustInit(t, makeRaw(
		"####",
		"....",
		"....",
		"....",
	), 'B')

	board := tetris.NewBoard(4)
	if !solve(board, []tetris.Piece{pieceA, pieceB}) {
		t.Fatal("expected solve to succeed for size 4 board")
	}
}

func TestInitTetrominoPiecesLimit(t *testing.T) {
	raw := makeRaw(
		"#...",
		"#...",
		"#...",
		"#...",
	)

	inputs := make([]tetris.RawPiece, 27)
	for i := range inputs {
		inputs[i] = raw
	}

	_, err := initTetrominoPieces(inputs)
	if err == nil {
		t.Fatal("expected error for more than 26 tetrominoes")
	}
}

func TestInitTetrominoPiecesAssignsIDs(t *testing.T) {
	raw := makeRaw(
		"#...",
		"#...",
		"#...",
		"#...",
	)

	pieces, err := initTetrominoPieces([]tetris.RawPiece{raw, raw})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(pieces) != 2 || pieces[0].ID != 'A' || pieces[1].ID != 'B' {
		t.Fatalf("unexpected IDs: %+v", pieces)
	}
}
