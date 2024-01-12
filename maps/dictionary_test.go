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
		assertError(error, ErrorWordExists, t)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"

		dictionary := Dictionary{}
		dictionary.Add(word, definition)

		assertDefinition(dictionary, word, definition, t)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, "new definition")

		assertError(err, ErrorWordExists, t)
		assertDefinition(dictionary, word, definition, t)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		newDefinition := "this is a new definition"
		dictionary := Dictionary{word: definition}

		dictionary.Update(word, newDefinition)

		assertDefinition(dictionary, word, newDefinition, t)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"

		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertError(err, ErrorWordDoesNotExists, t)
	})
}

func TestDelete(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}

		dictionary.Delete(word)

		_, err := dictionary.Search(word)

		assertError(err, ErrorWordNotFound, t)
	})

}

func assertDefinition(dictionary Dictionary, word, definition string, t *testing.T) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("expected to find word ", err)
	}
	assertStrings(got, definition, t)
}

func assertStrings(got string, want string, t testing.TB) {
	t.Helper()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(got error, want error, t testing.TB) {
	t.Helper()

	if got == nil {
		t.Fatal("Expected an error, got none")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
