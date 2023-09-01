package main

import (
	"errors"
	"testing"
)

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word you were looking for")

func (dic Dictionary) Search(key string) (string, error) {
	value, ok := dic[key]
	if !ok {
		return "", ErrNotFound
	}

	return value, nil
}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	got, _ := dictionary.Search("test")
	want := "this is just a test"

	assertError(t, got, want)

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")

		assertError(t, got.Error(), ErrNotFound.Error())
	})
}

func assertError(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	dictionary.Add("test", "this is just a test")

	want := "this is just a test"
	got, err := dictionary.Search("test")
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if want != got {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func (dic Dictionary) Add(key string, value string) {
	dic[key] = value
}
