package trie_test

import (
	"testing"
)

func TestAddNewWord(t *testing.T) {
	rootTrie := trie.NewTrie()
	rootTrie.Insert("bread")
}

func TestAddAndSearch(t *testing.T) {
	rootTrie := trie.NewTrie()
	rootTrie.Insert("bread")
	rootTrie.Insert("water")

	if rootTrie.Search("bread") == nil {
		t.FailNow()
	}

	if rootTrie.Search("water") == nil {
		t.FailNow()
	}
}
