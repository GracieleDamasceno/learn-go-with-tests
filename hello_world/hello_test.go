package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("say hello to someone", func(t *testing.T) {
		got := Hello("Joe", "")
		want := "Hello Joe"

		assertStrings(got, want, t)
	})
	t.Run("say hello world when name is empty", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello world"

		assertStrings(got, want, t)
	})
	t.Run("say hello in spanish", func(t *testing.T) {
		got := Hello("Bobby", "spanish")
		want := "Hola Bobby"

		assertStrings(got, want, t)
	})
	t.Run("say hello in korean", func(t *testing.T) {
		got := Hello("Jennie", "korean")
		want := "안녕하세요 Jennie"

		assertStrings(got, want, t)
	})
}

func assertStrings(got string, want string, t testing.TB) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
