package trie_test

import (
	"testing"

	"github.com/fransafu/memory-text-search/trie"
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

func TestDeleteWord(t *testing.T) {
	rootTrie := trie.NewTrie()
	rootTrie.Insert("bread")

	err := rootTrie.Delete("bread")
	if err != nil {
		t.FailNow()
	}
}
