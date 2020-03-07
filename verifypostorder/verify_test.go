package verifypost

import (
	"fmt"
	"testing"
)

func TestVerify(t *testing.T) {
	input := []int{3, 10, 6, 9, 2}
	res := verifyPostorder(input)
	fmt.Println(res)
}
