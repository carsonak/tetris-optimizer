package tetris

import (
	"reflect"
	"testing"
)

// Helper to create a Raw 4x4 grid from strings
func makeRaw(t *testing.T, rows ...string) RawPiece {
	t.Helper()
	if len(rows) != 4 {
		panic("invalid: number of rows for RawPiece")
	}

	var grid RawPiece

	for y, rowStr := range rows {
		copy(grid[y][:], []rune(rowStr))
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
			raw: makeRaw(t,
				"####",
				"....",
				"....",
				"....",
			),
			tetID: 'A',
			want: makePiece('A', 4, 1,
				Point{0, 0}, Point{1, 0}, Point{2, 0}, Point{3, 0},
			),
			expectErr: false,
		},
		{
			name: "Valid O-Shape",
			raw: makeRaw(t,
				"##..",
				"##..",
				"....",
				"....",
			),
			tetID: 'B',
			want: makePiece('B', 2, 2,
				Point{0, 0}, Point{1, 0}, Point{0, 1}, Point{1, 1},
			),
			expectErr: false,
		},
		{
			name: "Valid L-Shape (Shifted in grid)",
			raw: makeRaw(t,
				".#..",
				".#..",
				".##.",
				"....",
			),
			tetID: 'C',
			want: makePiece('C', 2, 3,
				Point{0, 0}, Point{0, 1}, Point{0, 2}, Point{1, 2},
			),
			expectErr: false,
		},
		{
			name: "Valid T-Shape (Shifted in grid)",
			raw: makeRaw(t,
				"....",
				".###",
				"..#.",
				"....",
			),
			tetID: 'D',
			want: makePiece('D', 3, 2,
				Point{0, 0}, Point{1, 0}, Point{2, 0}, Point{1, 1},
			),
			expectErr: false,
		},
		{
			name: "Invalid: character",
			raw: makeRaw(t,
				"##..",
				"##..",
				"x...", // 'x' is invalid
				"....",
			),
			expectErr: true,
			errMsg:    "unrecognised character 'x'",
		},
		{
			name: "Invalid: Count (Too few)",
			raw: makeRaw(t,
				"###.",
				"....",
				"....",
				"....",
			),
			expectErr: true,
			errMsg:    "tetromino should have 4 blocks",
		},
		{
			name: "Invalid: Count (Too many)",
			raw: makeRaw(t,
				"####",
				"#...",
				"....",
				"....",
			),
			expectErr: true,
			errMsg:    "tetromino should have 4 blocks.",
		},
		{
			name: "Invalid: Connectivity (Isolated Block)",
			raw: makeRaw(t,
				"###.",
				"....",
				"#...",
				"....",
			),
			expectErr: true,
			errMsg:    "invalid tetromino",
		},
		{
			name: "Invalid: Connectivity (2 floating blocks)",
			raw: makeRaw(t,
				"##..",
				"....",
				"...#",
				"...#",
			),
			expectErr: true,
			errMsg:    "invalid tetromino",
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
