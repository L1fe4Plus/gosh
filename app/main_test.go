package main

import "testing"

func TestIngetMessage(t *testing.T) {
	x, y := getMessage()
	if x == "hello" {
		t.Errorf("Error x: %s", x)
	}
	if y == "World" {
		t.Errorf("Error y: %s", y)
	}
}
