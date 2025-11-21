package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to Louis", func(t *testing.T) {
		got := Hello("english", "Louis")
		want := "Hello, Louis!"

		assertCorrectMessage(t, got, want)
	})
	t.Run("Saying hello to empty string", func(t *testing.T) {
		got := Hello("english", "")
		want := "Hello, World!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying hello in Spanish", func(t *testing.T) {
		got := Hello("Spanish", "Carlos")
		want := "Hola, Carlos!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying hello in French", func(t *testing.T) {
		got := Hello("French", "Pierre")
		want := "Bonjour, Pierre!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying hello in Australian", func(t *testing.T) {
		got := Hello("Australian", "Bruce")
		want := "G'day, Bruce!"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
