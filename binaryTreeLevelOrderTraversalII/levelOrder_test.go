package levelOrderII

import "testing"

func NewTree(input []int) *TreeNode {
	if len(input) == 0 {
		return nil
	}
	if len(input) == 1 {
		return &TreeNode{
			Val: input[0],
		}
	}
	left := 0
	right := len(input) - 1
	mid := (right - left) >> 1
	return &TreeNode{
		Val:   input[mid],
		Left:  NewTree(input[:mid]),
		Right: NewTree(input[mid+1:]),
	}
}

func TestLevelOrder(t *testing.T) {
	input := []int{2, 5, 9, 11, 17, 29, 35, 38}
	res := NewTree(input)
	result := levelOrderBottom(res)
	t.Errorf("%v", result)
}
