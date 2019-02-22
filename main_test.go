package main

import "testing"

func TestHelloWorld(t *testing.T) {
	if 1 == 0 {
		t.Fatal("Oh, no!")
	}
}
