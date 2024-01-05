package iteration

import "testing"

func TestIteration(t *testing.T) {
	got := Iteration("a")
	want := "aaaaa"
	assertStrings(got, want, t)
}

func assertStrings(got string, want string, t testing.TB) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
