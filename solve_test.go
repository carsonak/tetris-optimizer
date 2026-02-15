package main

import (
	"testing"

	"tetris-optimizer/tetris"
)

func TestMinimumBoardSize(t *testing.T) {
	testData := []struct {
		name     string
		count    int
		expected int
	}{
		{"1 tetromino", 1, 2},      // 1 piece = 4 cells, sqrt(4) = 2
		{"2 tetrominoes", 2, 3},    // 2 pieces = 8 cells, sqrt(8) ≈ 2.83, ceil = 3
		{"4 tetrominoes", 4, 4},    // 4 pieces = 16 cells, sqrt(16) = 4
		{"10 tetrominoes", 10, 7},  // 10 pieces = 40 cells, sqrt(40) ≈ 6.32, ceil = 7
		{"64 tetrominoes", 64, 16}, // 64 pieces = 256 cells, sqrt(256) = 16
	}

	for _, test := range testData {
		t.Run(test.name, func(t *testing.T) {
			got := minimumBoardSize(test.count)
			if got != test.expected {
				t.Errorf("minimumBoardSize(%d) = %d, want %d", test.count, got, test.expected)
			}
		})
	}
}

func TestMaximumBoardSize(t *testing.T) {
	testData := []struct {
		name  string
		count int
		min   int
	}{
		{"1 tetromino", 1, 4},      // cells: 1*16 = 16, sqrt(16) = 4
		{"2 tetrominoes", 2, 6},    // cells: 2*16 = 32, sqrt(32) ≈ 5.6568, 6
		{"4 tetrominoes", 4, 8},    // cells: 4*16 = 64, sqrt(64) = 8
		{"64 tetrominoes", 64, 32}, // cells: 64*16 = 1024, sqrt(1024) = 32
	}

	for _, test := range testData {
		t.Run(test.name, func(t *testing.T) {
			got := maximumBoardSize(test.count)
			if got != test.min {
				t.Errorf("maximumBoardSize(%d) = %d, want at least %d", test.count, got, test.min)
			}
		})
	}
}

func TestSolveEmptyList(t *testing.T) {
	board := tetris.NewBoard(4)

	if !solve(&board, []tetris.Piece{}) {
		t.Fatal("expected solve to succeed with empty piece list")
	}
}

func TestSolve(t *testing.T) {
	testData := []struct {
		name     string
		expected bool
		board    tetris.Board
		piece    tetris.Piece
	}{
		{
			name:     "Single piece",
			expected: true,
			board:    tetris.NewBoard(2),
			piece: tetris.Piece{
				Pos:    [4]tetris.Point{{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 0, Y: 1}, {X: 1, Y: 1}},
				Width:  2,
				Height: 2,
				ID:     'A',
			},
		},
		{
			name:     "Too small board",
			expected: false,
			board:    tetris.NewBoard(2),
			piece: tetris.Piece{
				Pos:    [4]tetris.Point{{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 2, Y: 0}, {X: 3, Y: 0}},
				Width:  4,
				Height: 1,
				ID:     'A',
			},
		},
	}

	for _, test := range testData {
		t.Run(test.name, func(t *testing.T) {
			output := solve(&test.board, []tetris.Piece{test.piece})
			if output != test.expected {
				t.Fatalf(
					"expected solve(board(%d), {piece(W: %d, H: %d)}) == '%v', got '%v'",
					test.board.Size, test.piece.Width, test.piece.Height, test.expected, output,
				)
			}
		})
	}
}

func TestFindSmallestSquare(t *testing.T) {
	// Create a simple 2x2 piece
	piece := tetris.Piece{
		Pos:    [4]tetris.Point{{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 0, Y: 1}, {X: 1, Y: 1}},
		Width:  2,
		Height: 2,
		ID:     'A',
	}

	board := FindSmallestSquare([]tetris.Piece{piece})

	if board.Size < 2 {
		t.Errorf("expected board size >= 2 for single 2x2 piece, got %d", board.Size)
	}

	if board.Size > 4 {
		t.Errorf("expected board size <= 4 for single piece, got %d", board.Size)
	}
}
