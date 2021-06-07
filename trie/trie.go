package trie

import (
	"errors"
)

type Trie struct {
	letter      rune
	children    []*Trie
	isEndOfWord bool
}

func NewTrie() *Trie {
	newTrie := &Trie{}
	newTrie.children = []*Trie{}
	return newTrie
}

func (trie *Trie) Insert(word string) (*Trie, error) {
	letters, node, index := []rune(word), trie, 0

	for index < len(word) {
		if exists, childNode := node.hasLetter(letters[index]); exists {
			node = childNode
		} else {
			node = node.addLetter(letters[index])
		}

		index++

		if index == len(word) {
			node.isEndOfWord = true
		}
	}

	return node, nil
}

func (trie *Trie) Search(query string) (bool, *Trie) {
	index, letters, currentNode, queryLength := 0, []rune(query), trie, len(query)

	for index < queryLength {
		if exists, node := currentNode.hasLetter(letters[index]); exists {
			currentNode = node
		} else {
			return false, nil
		}
		index++
	}

	return currentNode.isEndOfWord, currentNode
}

func (trie *Trie) Delete(word string) error {
	isFound, currencurrentNode := trie.Search(word)

	if !isFound {
		return errors.New("the word is not found")
	}

	currencurrentNode.isEndOfWord = false

	return nil
}

func (trie *Trie) hasLetter(letter rune) (bool, *Trie) {
	for _, childNode := range trie.children {
		if childNode.letter == letter {
			return true, childNode
		}
	}

	return false, nil
}

func (trie *Trie) addLetter(letter rune) *Trie {
	newChildrenNode := NewTrie()
	newChildrenNode.letter = letter
	trie.children = append(trie.children, newChildrenNode)
	return newChildrenNode
}
