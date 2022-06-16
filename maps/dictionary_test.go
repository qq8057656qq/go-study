package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}
	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)
	})
	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")
		want := ErrNotFound
		assertError(t, got, want)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"
		err := dictionary.Add(word, definition)
		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")
		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"
		err := dictionary.Update(word, newDefinition)
		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}
		err := dictionary.Update(word, definition)
		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dictionary{word: "test definition"}
	dictionary.Delete(word)
	_, err := dictionary.Search(word)
	if err != ErrNotFound {
		t.Errorf("Expected '%s' to be deleted", word)
	}
}

func TestDictionaryErr(t *testing.T) {
	err := DictionaryErr("custom error")
	got := err.Error()
	want := "custom error"
	assertStrings(t, got, want)
}

func assertDefinition(t *testing.T, dictionary Dictionary, word string, definition string) {
	want := definition
	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added definition:", err)
	}
	if want != got {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func assertError(t *testing.T, got error, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error '%s' want '%s'", got, want)
	}
}

func assertStrings(t *testing.T, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got '%s' want '%s' given, '%s'", got, want, "test")
	}
}