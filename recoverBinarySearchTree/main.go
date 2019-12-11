package main

import (
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverTree(root *TreeNode) {
	list := inorderTraversal(root)
	sort.Ints(list)
	pos := 0
	constructTree(root, list, &pos)
}

func inorderTraversal(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}
	if root.Left != nil {
		temp := inorderTraversal(root.Left)
		result = append(result, temp...)
	}
	result = append(result, root.Val)
	if root.Right != nil {
		temp := inorderTraversal(root.Right)
		result = append(result, temp...)
	}
	return result
}

func constructTree(root *TreeNode, res []int, pos *int) {
	if root == nil {
		return
	}
	if root.Left != nil {
		constructTree(root.Left, res, pos)
	}
	root.Val = res[*pos]
	*pos = *pos + 1
	if root.Right != nil {
		constructTree(root.Right, res, pos)
	}
	return
}

func main() {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 3, Left: nil, Right: &TreeNode{Val: 2, Left: nil, Right: nil}}, Right: nil}
	recoverTree(root)
}
