package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Arun", "")
		want := "Hello, Arun"
		assertMessage(t, got, want)
	})

	t.Run("'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertMessage(t, got, want)
	})

	t.Run("in spanish", func(t *testing.T) {
		got := Hello("World", "Spanish")
		want := "Hola, World"
		assertMessage(t, got, want)
	})

	t.Run("in french", func(t *testing.T) {
		got := Hello("World", "French")
		want := "Bonjour, World"
		assertMessage(t, got, want)
	})

}

func assertMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
