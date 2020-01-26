package main

import (
	"fmt"
	"testing"
)

func TestMaximumSub(t *testing.T) {
	input := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArrayOn(input))
}
