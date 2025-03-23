package main

import "testing"

func TestHelloWorld(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Abhishek")
		want := "Hello, Abhishek!"

		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World!' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World!"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t *testing.T, got, want string) {
	// t.Helper() is needed to tell the test suite that this method is a helper.
	// By doing this when it fails the error message will show the name of the function
	// that failed.
	t.Helper()
	if got != want {
		// %q is for quoting strings
		t.Errorf("got %q want %q", got, want)
	}
}
