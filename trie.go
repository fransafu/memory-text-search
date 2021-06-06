package trie

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

func (trie *Trie) Insert(word string) *Trie {
	return nil
}

func (trie *Trie) Search(query string) *Trie {
	return nil
}

func (trie *Trie) Delete(word string) error {
	return nil
}
