package main

import "testing"

func TestGetTetrominoes(t *testing.T) {
	t.Run("file not found", func(t *testing.T) {
		_, err := GetTetrominoes("does_not_exist")

		if err == nil {
			t.Fatal("Expected an error, got nil")
		}
	})
}
