package main

import "testing"

func TestHello(t *testing.T) {
	got := hello("world")
	want := "hello world"
	if got == want {
		t.Errorf("got %s, want %s", got, want)
	}
}
