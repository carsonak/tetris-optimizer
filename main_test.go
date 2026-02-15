package main

import (
	"testing"
	"tetris-optimizer/tetris"
)

func makeRaws(raw tetris.RawPiece, count int) []tetris.RawPiece {
	raws := make([]tetris.RawPiece, count)

	for i := range raws {
		raws[i] = raw
	}

	return raws
}

func TestInitTetrominoPieces(t *testing.T) {
	// Simple I-piece
	var rawI tetris.RawPiece

	for i := range 4 {
		copy(rawI[i][:], []rune("#   "))
	}

	t.Run("At limit", func(t *testing.T) {
		pieces, err := initTetrominoPieces(makeRaws(rawI, 26))
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(pieces) != 26 {
			t.Fatalf("expected 26 pieces, got %d", len(pieces))
		}

		id := 'A'
		for _, p := range pieces {
			if p.ID != id {
				t.Fatalf("expected ID: %c, got: %c", id, p.ID)
			}

			id++
		}
	})

	t.Run("More than max limit", func(t *testing.T) {
		expectedMsg := "cannot process more than 26 tetrominoes"
		_, err := initTetrominoPieces(makeRaws(rawI, 27))
		if err == nil {
			t.Fatal("expected error but got nil")
		}

		if err.Error() != expectedMsg {
			t.Errorf("expected error %q, got %q", expectedMsg, err.Error())
		}
	})

}
