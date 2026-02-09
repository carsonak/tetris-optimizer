package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "ERROR: USAGE: %s tetromino_file", os.Args[0])
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v", err)
		os.Exit(1)
	}

	defer file.Close()
	_, err = ParseTetrominoStream(bufio.NewScanner(file))
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v", err)
	}
}
