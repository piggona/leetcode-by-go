package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(root *TreeNode, low, high int64) bool {
	if root == nil {
		return true
	}

	if int64(root.Val) <= low || int64(root.Val) >= high {
		return false
	}

	return helper(root.Left, low, int64(root.Val)) && helper(root.Right, int64(root.Val), high)
}

func main() {
	return
}
