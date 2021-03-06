package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Joseph", "")
		want := "Hello, Joseph"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello in French", func(t *testing.T) {
		got := Hello("George", "French")
		want := "Bonjour, George"

		assertCorrectMessage(t, got, want)
	})
	t.Run("say hello in Tagalog", func(t *testing.T) {
		got := Hello("Luciano", "Tagalog")
		want := "Kamusta, Luciano"

		assertCorrectMessage(t, got, want)
	})
}
