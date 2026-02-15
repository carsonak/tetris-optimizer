# Tetris Optimizer

A Go program that assembles tetrominoes into the smallest possible square using a high-performance backtracking algorithm with heuristic optimizations.

## Quick Start

```bash
# Build the program
go build -o tetris-optimizer .

# Run with an input file
./tetris-optimizer tests/samples/sample00-04

```

## Input Format

Each tetromino must be represented as a 4×4 grid where:

* `#` represents a block
* `.` represents an empty cell
* Tetrominoes are separated by blank lines

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

* Exactly 4 blocks (`#`)
* Orthogonal connectivity (valid tetromino shape)
* 4×4 grid format
* Valid characters only (`#` or `.`)

Supports up to 26 tetrominoes (A-Z).

## Output

Prints the solution board with each tetromino labelled by a unique letter:

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
├── main.go                     # Entry point, CLI handling
├── parse_tetromino_stream.go   # Input file parsing and validation
├── solve.go                    # Hybrid backtracking solver (Heuristic + Fallback)
├── tetris/                     # Core data structures package
│   ├── piece.go                # Tetromino normalization and validation
│   ├── board.go                # Optimized board with contiguous memory
│   └── *_test.go               # Unit tests
├── tests/                      # Test suites
│   ├── run_tests.sh            # Advanced test runner script
│   ├── bad_examples/           # Invalid input test cases
│   ├── good_examples/          # Valid input test cases
│   └── samples/                # Benchmark samples
└── *_test.go                   # Unit tests

```

## Algorithm: The Hybrid Solver

The solver uses a dual-strategy approach to handle both "complex" and "trick" puzzles efficiently:

1. **Calculate Bounds**: Determine the theoretical minimum square size ().
2. **Iterate Sizes**: Start from the minimum size and increase until a solution is found.
3. **Strategy A (The Sprint)**:
    * **Heuristic**: Sort pieces by size (Largest/Widest first).
    This constrains the search space early, solving complex/dense puzzles instantly.
    * **Timeout**: A strict 500ms deadline is applied.
    If the solver gets stuck in a "bad root" branch (a heuristic trap), it aborts.

4. **Strategy B (The Fallback)**:
    * **Condition**: Runs only if Strategy A times out.
    * **Logic**: Reverts to the original input order and solves without a timeout.
    This handles puzzles where specific piece ordering is required to avoid dead ends.

5. **Backtracking**: Uses recursive depth-first search to place pieces.

**Complexity**: O(n! × size²)
**Optimizations**:

* **Hybrid Heuristic**: Solves hard cases in <0.5s while preventing worst-case freezes.
* **Contiguous Memory**: Board is allocated as a single flat array for cache locality.
* **Normalization**: Pieces are shifted to (0,0) to reduce coordinate math.

## Testing

The project includes a comprehensive test runner `tests/run_tests.sh`.

```bash
# Run unit tests
go test ./...

# Run the integration test suite, must build binary file first as shown in [#Quick Start]
# Flags: -s (samples), -e (examples), --colour (toggle color)
./tests/run_tests.sh -s ./tetris-optimizer

```

### Test File Conventions

Test files in `tests/good_examples` can end with `-NN` (e.g., `test-04`)
to assert that the solution contains exactly `NN` empty spaces.

## Error Handling

Errors are written to stderr with an `ERROR` prefix and the program exits with code 1.
Common errors include:

* Invalid file format
* Discontinuous tetromino shapes
* More than 26 tetrominoes

## Implementation Details

* **Solve Context (`solveCtx`)**: Manages deadlines and operation counting to
minimize syscall overhead (`time.Now()`) during recursion.
* **Memory Layout**: The board uses a 1D slice representation behind the scenes to
minimize pointer indirection and improve CPU cache hits.
* **Piece Normalization**: All pieces are pre-calculated to their top-left most position to
simplify collision checks.
