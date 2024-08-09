package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dict := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dict.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("word not found", func(t *testing.T) {
		_, err := dict.Search("whatever")
		if err == nil {
			t.Fatal("expected an error, didn't get one")
		}
		assertError(t, err, ErrWordNotFound)
	})
}

func TestAddWord(t *testing.T) {
	dict := Dictionary{}

	t.Run("word does not exist", func(t *testing.T) {
		word := "test"
		def := "this is just a test"
		err := dict.Add(word, def)

		assertError(t, err, nil)
		assertDefinition(t, dict, word, def)
	})

	t.Run("word already exists", func(t *testing.T) {
		word := "test"
		def := "this is just a test"
		err := dict.Add(word, def)

		assertError(t, err, ErrWordExistsAlready)
	})
}

func TestUpdateWord(t *testing.T) {
	word := "test"
	def := "this is just a test"
	newDef := "new definition"

	t.Run("existing word", func(t *testing.T) {
		dict := Dictionary{word: def}

		dict.Update(word, newDef)

		assertDefinition(t, dict, word, newDef)
	})

	t.Run("word does not exist", func(t *testing.T) {
		dict := Dictionary{}
		err := dict.Update(word, newDef)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDeleteWord(t *testing.T) {
	word := "test"
	dict := Dictionary{word: "test definition"}

	dict.Delete(word)

	_, err := dict.Search(word)

	assertError(t, err, ErrWordNotFound)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q, wanted error %q", got, want)
	}
}

func assertDefinition(t testing.TB, dict Dictionary, word, definition string) {
	t.Helper()
	got, err := dict.Search(word)

	if err != nil {
		t.Fatal("got an unexpected error:", err)
	}

	assertStrings(t, got, definition)
}
