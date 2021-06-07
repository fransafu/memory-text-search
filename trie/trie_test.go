package trie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddNewWord(t *testing.T) {
	rootTrie := NewTrie()
	_, err := rootTrie.Insert("bread")
	assert.Equal(t, err, nil)
}

func TestAddAndSearch(t *testing.T) {
	rootTrie := NewTrie()
	rootTrie.Insert("bread")
	rootTrie.Insert("water")

	if isFound, _ := rootTrie.Search("bread"); isFound == false {
		t.FailNow()
	}

	if isFound, _ := rootTrie.Search("water"); isFound == false {
		t.FailNow()
	}
}

func TestDeleteWord(t *testing.T) {
	rootTrie := NewTrie()
	rootTrie.Insert("bread")

	err := rootTrie.Delete("bread")
	if err != nil {
		t.FailNow()
	}
}
