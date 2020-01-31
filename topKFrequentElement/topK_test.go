package topK

import (
	"fmt"
	"testing"
)

func TestTopK(t *testing.T) {
	nums := []int{1}
	k := 1
	res := topKFrequent(nums, k)
	fmt.Println(res)
}
