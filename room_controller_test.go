package main

import "testing"

func TestAdd(t *testing.T) {
	x := 5
	y := 10
	v := addInts(x, y)
	if v != 15 {
		t.Error("Expected 15, got ", v)
	}
}
