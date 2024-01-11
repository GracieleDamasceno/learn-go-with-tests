package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}
	t.Run("existing word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(got, want, t)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, error := dictionary.Search("unknown")
		want := "word not found"

		assertError(error, want, t)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		dictionary.Add("test", "this is just a test")

		want := "this is just a test"
		got, err := dictionary.Search("test")

		if err != nil {
			t.Fatal("expected to find word ", err)
		}

		assertStrings(got, want, t)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new definition")

		assertError(err, "word already exists", t)

		got, error := dictionary.Search(word)
		if error != nil {
			t.Fatal("expected to find word ", err)
		}

		assertStrings(got, definition, t)
	})
}

func assertStrings(got string, want string, t testing.TB) {
	t.Helper()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(got error, want string, t testing.TB) {
	if got == nil {
		t.Fatal("Expected an error, got none")
	}

	if got.Error() != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
