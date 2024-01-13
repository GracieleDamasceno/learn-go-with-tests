package dependencyinjection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "John")

	got := buffer.String()
	want := "Hello John"

	assertStrings(got, want, t)
}

func assertStrings(got string, want string, t testing.TB) {
	t.Helper()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
