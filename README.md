# Tetris Optimizer

A Go program that assembles tetrominoes into the smallest possible square using backtracking.

## Overview

Reads tetrominoes (4-block Tetris pieces) from a text file and finds the optimal square arrangement. Each piece is labeled with a letter (A, B, C, etc.).

## Quick Start

```bash
go build -o tetris-optimizer .
./tetris-optimizer <tetromino-file>
```

## Input Format

Each tetromino is a 4x4 grid with `#` for blocks and `.`/` ` for empty space. Tetrominoes are separated by blank lines:

```
#...
#...
#...
#...

....
....
..##
..##
```

## Output

Solution board with each tetromino labeled:

```
ABBBB
ACCCEE
AFFCEE
A.FFGG
HHHDDG
.HDD.G
```

## Algorithm

1. Calculate min/max board sizes based on piece count
2. Test each size from min to max
3. Use recursive backtracking to place pieces
4. Return first valid solution (guarantees smallest square)

**Complexity**: O(n! × size²) worst case, where n is piece count
**Optimizations**: Early termination, efficient bounds checking
