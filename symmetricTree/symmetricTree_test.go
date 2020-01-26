package symmetricTree

import (
	"algorithms/utils"
	"fmt"
	"testing"
)

func TestIsSymmetric(t *testing.T) {
	res := utils.GetTree([]int{1, 2, 2, 3, 4, 4, 3})
	fmt.Println(isSymmetric(res))
}
