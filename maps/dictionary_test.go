package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is only a test"}
	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is only a test"
		AssertStrings(t, got, want)
	})
	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("word")
		AssertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	dictionary.Add("test", "this is only a test")
	got, _ := dictionary.Search("test")
	want := "this is only a test"
	AssertStrings(t, got, want)
}

func AssertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got: %v but wanted: %v", got, want)
	}
}
func AssertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("Got: %q but wanted: %q", got, want)
	}

}
