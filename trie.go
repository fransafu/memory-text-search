package trie

type Trie struct {
	letter      rune
	children    []*Trie
	isEndOfWord bool
}
