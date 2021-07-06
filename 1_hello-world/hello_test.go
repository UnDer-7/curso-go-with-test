package hello_world

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, wat string) {
		t.Helper()
		if got != wat {
			t.Errorf("got %q wat %q", got, wat)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Frank", "French")
		want := "Bonjour, Frank"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Portuguese", func(t *testing.T) {
		got := Hello("MC Poze", "Portuguese")
		want := "Eae, MC Poze"

		assertCorrectMessage(t, got, want)
	})
}