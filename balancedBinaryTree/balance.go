package balanceBinaryTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	_, ok := helper(root)
	return ok
}

func helper(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	left, ok := helper(root.Left)
	if !ok {
		return 0, false
	}
	right, ok := helper(root.Right)
	if !ok {
		return 0, false
	}
	if abs(left-right) > 1 {
		return 0, false
	}
	return getMax(left, right) + 1, true

}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
