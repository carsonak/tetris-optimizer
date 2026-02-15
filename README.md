# Tetris Optimizer

A Go program that assembles tetrominoes into the smallest possible square using backtracking.

## Quick Start

```bash
# Build the program
go build -o tetris-optimizer .

# Run with an input file
./tetris-optimizer <tetromino-file>

# Example
./tetris-optimizer tests/good_examples/goodexample00-00
```

## Input Format

Each tetromino must be represented as a 4×4 grid where:

- `#` represents a block
- `.` or space (` `) represents an empty cell
- Tetrominoes are separated by blank lines

**Valid Example:**

```text
#...
#...
#...
#...

....
....
..##
..##
```

### Input Validation

Each tetromino is validated for:

- Exactly 4 blocks (`#`)
- Orthogonal connectivity (valid tetromino shape)
- 4×4 grid format
- Valid characters only (`#`, `.`, space)

Supports up to 26 tetrominoes (A-Z).

## Output

Prints the solution board with each tetromino labeled by a unique letter:

```text
ABBBB.
ACCCEE
AFFCEE
A.FFGG
HHHDDG
.HDD.G
```

## Project Structure

```text
tetris-optimizer/
├── main.go                        # Entry point, CLI handling, orchestration
├── parse_tetromino_stream.go      # Input file parsing and validation
├── solve.go                       # Backtracking solver algorithm
├── tetris/                        # Core data structures package
│   ├── piece.go                   # Tetromino representation and validation
│   └── board.go                   # Game board operations
├── tests/                         # Test files and examples
│   ├── run_tests.sh              # Test runner script
│   ├── bad_examples/             # Invalid input test cases
│   ├── good_examples/            # Valid input test cases
│   └── samples/                  # Additional sample inputs
└── *_test.go                     # Unit tests
```

## Algorithm

1. **Calculate Bounds**: Min ⌈√(n×4)⌉, Max ⌈√(n×16)⌉
2. **Test Sizes**: Iterate from min to max
3. **Backtrack**: Try all positions for each piece
4. **Return**: First valid solution (guaranteed smallest)

**Complexity**: O(n! × size²) | **Optimizations**: Normalization, early termination

## Testing

```bash
go test ./...                  # Run unit tests
go test -cover ./...           # With coverage
cd tests && ./run_tests.sh     # Run example tests
```

## Error Handling

Errors written to stderr with "ERROR:" prefix, exit code 1.
Common issues: invalid file, invalid format/shape, too many pieces (>26).

## Implementation Details

**Piece Normalization**: Top-left block at (0,0) reduces search space.
**Neighbor Validation**: Total neighbor count of 6 or 8 ensures connected shape.
**Board Operations**: `Place()`, `Remove()`, `ToString()`
