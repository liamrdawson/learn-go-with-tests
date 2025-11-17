package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Say hello to provided name", func(t *testing.T) {
		got := Hello("Liam")
		want := "Hello, Liam"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Say 'Hello, World' when no name is provided", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	// Marking this function as a test helper means that the output
	// log will reference the line of the failing code rather than
	// where the failure is caught in our test.
	t.Helper()
	if got != want {
		t.Errorf("\ngot: %q\nwant: %q", got, want)
	}
}
