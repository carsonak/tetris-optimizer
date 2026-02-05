package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "ERROR: %s tetromino_file", os.Args[0])
		os.Exit(1)
	}

	
}
