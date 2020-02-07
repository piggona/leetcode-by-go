package pathSum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	return helper(root, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}

func helper(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	cur := sum - root.Val
	if cur == 0 {
		return helper(root.Left, cur) + helper(root.Right, cur) + 1
	}
	return helper(root.Left, cur) + helper(root.Right, cur)
}
