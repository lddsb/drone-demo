package main

import "testing"

func TestOutput(t *testing.T) {
	if 1 == 0 {
		t.Fatal("something error here!")
	}
}
