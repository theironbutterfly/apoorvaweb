package main

import "testing"

func TestAddition(t *testing.T) {
    result := 2 + 3

    if result != 5 {
        t.Errorf("Expected 5, got %d", result)
    }
}
