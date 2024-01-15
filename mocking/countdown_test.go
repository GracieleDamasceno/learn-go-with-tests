package mocking

import (
	"bytes"
	"reflect"
	"testing"
)

func TestCountdown(t *testing.T) {
	t.Run("countdown from 3 and Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Countdown(buffer, &SpyCountdownOperations{})

		got := buffer.String()
		want := `3
2
1
Go!`

		assertStrings(got, want, t)
	})

	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &SpyCountdownOperations{}
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		assertDeepCopy(want, spySleepPrinter, t)
	})
}

func assertDeepCopy(want []string, got *SpyCountdownOperations, t *testing.T) {
	t.Helper()
	if !reflect.DeepEqual(want, got.Calls) {
		t.Errorf("wanted %v got %v", want, got.Calls)
	}
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
