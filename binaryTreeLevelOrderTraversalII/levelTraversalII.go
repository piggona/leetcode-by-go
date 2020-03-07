package levelOrderII

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	temp := []*TreeNode{root}
	result := [][]int{}
	helper(temp, &result)
	return result

}

func helper(stack []*TreeNode, result *[][]int) {
	if len(stack) == 0 {
		return
	}
	temp := stack
	res := []int{}
	stack = []*TreeNode{}
	for _,val := range temp {
		res = append(res, val.Val)
		left := val.Left
		right := val.Right
		if left != nil {
			stack = append(stack, left)
		}
		if right != nil {
			stack = append(stack,right)
		}
	}
	helper(stack, result)
	*result = append(*result, res)
}
