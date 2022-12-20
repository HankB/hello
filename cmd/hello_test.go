package main

import (
	"testing"
	"hello/internal"
)

func TestAdd(t *testing.T) {
	if lib.Add(1,2) != 3 {	// passing test
		t.Fatal("Add(1,2) did not return 4")
	}
	if lib.Add(1,3) != 3 { // failing test
		t.Fatal("Add(1,3) did not return 3")
	}
}