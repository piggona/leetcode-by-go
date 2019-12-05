package utils

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getTree(input []int) *TreeNode {
	result := &TreeNode{Val: input[0]}
	prev := result
	direction := 0
	for i := 1; i < len(input); i++ {
		if input[i] == -1 {
			direction = 1
		} else {
			temp := &TreeNode{}
			temp.Val = input[i]
			if direction == 0 {
				prev.Left = temp
			} else {
				prev.Right = temp
			}
			prev = temp
			direction = 0
		}
	}
	return result
}
