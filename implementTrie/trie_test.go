package implementTrie

import (
	"fmt"
	"os"
	"testing"
)

func TestTrie(t *testing.T) {
	trie := Constructor()
	trie.Insert("apple")
	res1 := trie.Search("app")
	res2 := trie.StartsWith("app")
	fmt.Fprintf(os.Stdout, "res1: %v, res2: %v", res1, res2)
}
