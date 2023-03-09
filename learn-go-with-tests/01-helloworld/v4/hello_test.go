package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Chris")
	want := "hello chris"

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
