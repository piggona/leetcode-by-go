package numIslands

import (
	"fmt"
	"testing"
)

func TestNumIsland(t *testing.T) {
	input := [][]byte{
		{'1', '1', '1'},
		{'0', '1', '0'},
		{'1', '1', '1'},
	}
	fmt.Println(numIslands(input))
}
