package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Arun")
	want := "Hello, Arun"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
