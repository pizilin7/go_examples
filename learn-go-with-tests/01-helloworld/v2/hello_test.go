package main

// go test v2/hello_test.go v2/hello.go

import "testing"

func TestHello(t *testing.T) {
	got := hello()
	want := "hello world"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}