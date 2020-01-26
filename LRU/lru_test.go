package LRU

import (
	"fmt"
	"testing"
)

func TestLRU(t *testing.T) {
	lru := Constructor(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	res := lru.Get(1)
	fmt.Println(res)
}
