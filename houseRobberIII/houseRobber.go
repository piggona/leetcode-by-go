package HouseRobberIII

import (
	"algorithms/utils"
	"math"
)

type (
	TreeNode = utils.TreeNode
)

func rob(root *TreeNode) int {
	return robber(root, true)
}

func robber(root *TreeNode, available bool) int {
	if root == nil {
		return 0
	}
	left := root.Left
	right := root.Right
	if available {
		//使用当前节点
		res1 := robber(left, false) + robber(right, false) + root.Val
		// 不使用当前节点
		res2 := robber(left, true) + robber(right, true)
		return getMax(res1, res2)
	} else {
		return robber(left, true) + robber(right, true)
	}
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

func robBetter(root *TreeNode) int {
    if root == nil {
        return 0
    }

    v := calc(root)
    return max(v.include, v.exclude)
}

type value struct {
    include int
    exclude int
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func calc(root *TreeNode) value {
    if root == nil {
        return value{}
    }
    
    left := calc(root.Left)
    right := calc(root.Right)
    v := value{}
    v.include = root.Val + left.exclude + right.exclude
    v.exclude = max(left.include, left.exclude) + max(right.include, right.exclude)
    return v
}
