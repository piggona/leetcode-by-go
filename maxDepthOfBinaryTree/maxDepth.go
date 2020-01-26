package maxDepthOfBinaryTree

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + getMax(maxDepth(root.Right), maxDepth(root.Left))
}

func getMax(max ...int) int {
	res := math.MinInt32
	for _, value := range max {
		if value > res {
			res = value
		}
	}
	return res
}

func getMin(nums ...int) int {
	res := math.MaxInt32
	for _, value := range nums {
		if value < res {
			res = value
		}
	}
	return res
}
