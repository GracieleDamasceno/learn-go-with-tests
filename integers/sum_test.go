package integers

import "testing"

func TestSum(t *testing.T) {
	sum := Add(2, 2)
	want := 4

	assertIntegers(sum, want, t)
}

func assertIntegers(got int, want int, t testing.TB) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
