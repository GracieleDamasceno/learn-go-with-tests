package mocking

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}

	Countdown(buffer, spySleeper)

	got := buffer.String()
	want := `3
2
1
Go!`

	assertStrings(got, want, t)
	assertInteger(spySleeper.Calls, 3, t)
}

func assertStrings(got string, want string, t testing.TB) {
	t.Helper()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertInteger(got int, want int, t testing.TB) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}