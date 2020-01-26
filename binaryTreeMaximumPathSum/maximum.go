package binaryTreeMaximumPathSum

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSumX(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max := math.MinInt32
	rmax := math.MinInt32
	helperX(root, &rmax, &max)
	return max
}

func helperX(root *TreeNode, max *int, fmax *int) {
	val := root.Val
	temp := math.MinInt32
	lmax := math.MinInt32
	rmax := math.MinInt32
	if root.Right == nil && root.Left == nil {
		return
	}
	if root.Left != nil {
		left := root.Left.Val
		if left+val < left {
			temp = left
		} else {
			temp = left + val
			root.Left.Val = left + val
		}
		helperX(root.Left, &lmax, fmax)
	}
	if root.Right != nil {
		right := root.Right.Val
		if right+val < right && right > temp {
			temp = right
		}
		if right+val >= right {
			if right+val > temp {
				temp = right + val
			}
			root.Right.Val = right + val
		}
		helperX(root.Right, &rmax, fmax)
	}
	if *max < temp {
		*max = temp
	}
	if root.Val >= 0 {
		ftemp := rmax + lmax - root.Val
		if ftemp > *fmax {
			*fmax = ftemp
		}
	} else {
		*fmax = *max
	}
	return
}

func maxPathSum(root *TreeNode) int {
	res := math.MinInt32
	helper(root, &res)
	return res
}

func helper(node *TreeNode, res *int) int {
	if node == nil {
		return 0
	}
	left := max(0, helper(node.Left, res))
	right := max(0, helper(node.Right, res))
	*res = max(*res, left+right+node.Val)
	return max(left, right) + node.Val
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
