# memory-text-search
A simple memory text search. You can load the content (strings) and ask for them

# Getting started

How to get the package?

* `go get -v github.com/fransafu/memory-text-search`

How to test the program?

* Download the project (git clone or zip file)
* Download the dependencies, run the command `go get ./...`
* Test the program with the following command line: `go test -v ./...`

# Example usage

If you want to make your own program, you can use as base this source code, you just need to copy and paste the following example, in your own file (create a `main.go` file).

If you want to create your own program, you can rely on it code, you just need to copy and paste the following example, in your own file (create a `main.go` file).

```go
package main

import (
	"fmt"

	trie "github.com/fransafu/memory-text-search/trie"
)

func main() {
	// Created empty data structure
	rootTrie := trie.NewTrie()

	// insert the word "bread" into the data structure
	rootTrie.Insert("bread")

	// Search the word in the data structure (it should have the word in memory)
	fmt.Println(rootTrie.Search("bread"))

	// Delete and check if the action generate an error (if the word doesn't exist)
	err := rootTrie.Delete("bread")
	fmt.Println(err)

	// Check if the word exists
	fmt.Println(rootTrie.Search("bread"))
}
```
