package main

import (
	"hello/internal"
	"testing"
)

func TestAdd(t *testing.T) {
	if adder.Add(1, 2) != 3 { // passing test
		t.Fatal("Add(1,2) did not return 4")
	}
	if adder.Add(1, 3) != 3 { // failing test
		t.Fatal("Add(1,3) did not return 3")
	}
}
