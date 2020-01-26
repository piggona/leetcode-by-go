package main

import (
	"fmt"
	"testing"
)

func TestGray(t *testing.T) {
	num := 4
	fmt.Println(grayToBinary(num))
	fmt.Println(binaryToGray(num))
}
