package utils

import "testing"

func TestBST(t *testing.T) {
	list := []int{1, 3, -1, 1, 2, -1, 5, 6, 7}
	root := &TreeNode{}
	constructTree(root, list)
	displayTree(root)
}
