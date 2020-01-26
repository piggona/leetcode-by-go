package symmetricTree

import (
	"algorithms/utils"
)

func isSymmetricBad(root *utils.TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left == root.Right && root.Left == nil {
		return true
	}
	if root.Left.Val != root.Right.Val {
		return false
	}
	lStack := []*utils.TreeNode{root.Left}
	rStack := []*utils.TreeNode{root.Right}
	for len(lStack) != 0 && len(rStack) != 0 {
		// 出队列
		ltemp := lStack[0]
		lStack = lStack[1:]
		rtemp := rStack[0]
		rStack = rStack[1:]
		// 取出栈元素的左右节点
		// 左&右
		lleft := ltemp.Left
		rright := rtemp.Right
		if (lleft != rright || lleft != nil) && (lleft.Val != rright.Val) {
			return false
		}
		if lleft == nil || rright == nil {
			return false
		}

		if ltemp.Left != rtemp.Right || ltemp.Left != nil {
			lStack = append(lStack, ltemp.Left)
			rStack = append(rStack, rtemp.Right)
		}
		// 右&左
		rleft := ltemp.Right
		lright := rtemp.Left
		if (rleft != lright || rleft != nil) && (rleft.Val != lright.Val) {
			return false
		}
		if rleft == nil || lright == nil {
			return false
		}
		if ltemp.Right != rtemp.Left || ltemp.Right != nil {
			lStack = append(lStack, ltemp.Right)
			rStack = append(rStack, rtemp.Left)
		}
	}
	return true

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return helper(root.Left, root.Right)
}

func helper(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if (left == nil && right != nil) || (right == nil && left != nil) || (right.Val != left.Val) {
		return false
	}
	return helper(left.Left, right.Right) && helper(left.Right, right.Left)
}
