package tetris

import (
	"reflect"
	"strings"
	"testing"
)

// Helper to create a Raw 4x4 grid from strings
// periods are replaced with spaces
// extra rows are ignored, missing rows are filled with spaces
// Usage: makeRaw("....", "####", "....", "....")
func makeRaw(rows ...string) RawPiece {
	var grid RawPiece

	for y, rowStr := range rows {
		if y > 3 {
			break
		}

		rowStr = strings.ReplaceAll(rowStr, ".", " ")
		copy(grid[y][:], []rune(rowStr))
	}

	for y := len(rows); y < 4; y++ {
		copy(grid[y][:], []rune("    "))
	}

	return grid
}

// Helper to create the expected Piece struct manually
func makePiece(id rune, w, h int, coords ...Point) Piece {
	var pos [4]Point

	copy(pos[:], coords)
	return Piece{
		ID:     id,
		Width:  w,
		Height: h,
		Pos:    pos,
	}
}

func TestInit(t *testing.T) {
	testData := []struct {
		name      string
		raw       RawPiece
		tetID     rune
		want      Piece
		expectErr bool
		errMsg    string
	}{
		{
			name: "Valid I-Shape (Horizontal)",
			raw: makeRaw(
				"####",
				"....",
				"....",
				"....",
			),
			tetID: 'A',
			want: makePiece('A', 4, 1,
				Point{0, -1}, Point{1, -1}, Point{2, -1}, Point{3, -1},
			),
			expectErr: false,
		},
		{
			name: "Valid O-Shape",
			raw: makeRaw(
				"##..",
				"##..",
				"....",
				"....",
			),
			tetID: 'B',
			want: makePiece('B', 2, 2,
				Point{0, -2}, Point{1, -2}, Point{0, -1}, Point{1, -1},
			),
			expectErr: false,
		},
		{
			name: "Valid L-Shape (Shifted in grid)",
			raw: makeRaw(
				".#..",
				".#..",
				".##.",
				"....",
			),
			tetID: 'C',
			want: makePiece('C', 2, 3,
				Point{0, -3}, Point{0, -2}, Point{0, -1}, Point{1, -1},
			),
			expectErr: false,
		},
		{
			name: "Invalid Character",
			raw: makeRaw(
				"##..",
				"##..",
				"x...", // 'x' is invalid
				"....",
			),
			expectErr: true,
			errMsg:    "unrecognised character 'x'",
		},
		{
			name: "Invalid Count (Too few)",
			raw: makeRaw(
				"###.",
				"....",
				"....",
				"....",
			),
			expectErr: true,
			errMsg:    "Piece should have 4 blocks",
		},
		{
			name: "Invalid Count (Too many)",
			raw: makeRaw(
				"####",
				"#...",
				"....",
				"....",
			),
			expectErr: true,
			errMsg:    "Piece should have 4 blocks.",
		},
		{
			name: "Invalid Connectivity (Isolated Block)",
			raw: makeRaw(
				"###.",
				"....",
				"#...",
				"....",
			),
			expectErr: true,
			errMsg:    "invalid Piece",
		},
	}

	for _, test := range testData {
		t.Run(test.name, func(t *testing.T) {
			got, err := Init(test.raw, test.tetID)

			if test.expectErr {
				if err == nil {
					t.Error("expected error but got nil")
				} else if err.Error() != test.errMsg {
					t.Errorf("expected error %q, got %q", test.errMsg, err.Error())
				}

				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Init() mismatch:\nGot:  %+v\nWant: %+v", got, test.want)
			}
		})
	}
}
