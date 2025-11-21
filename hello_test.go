package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Louis")
	want := "Hello, Louis!"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
