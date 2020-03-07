package diameter

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
    if root == nil {
        return 0
    }
    max := 1
	_ = helper(root,&max)
	return max-1
}

func helper(root *TreeNode,max *int) int {
	if root == nil {
		return 0
	}
	left := helper(root.Left,max)
	right := helper(root.Right,max)
	temp := left + right + 1
	if temp > *max {
		*max = temp
	}
	return getMax(left+1, right+1)
}

func getMax(nums ...int) int {
	max := math.MinInt32
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}