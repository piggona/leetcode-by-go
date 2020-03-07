package zigzag

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	stack := []*TreeNode{root}
	result := [][]int{}
	var in bool
	for len(stack) != 0 {
		temp := stack
		tr := []int{}
		stack = []*TreeNode{}
		for len(temp) != 0 {
			sel := temp[len(temp)-1]
			temp = temp[:len(temp)-1]
			tr = append(tr, sel.Val)
			if !in {
				if sel.Left != nil {
					stack = append(stack, sel.Left)
				}
				if sel.Right != nil {
					stack = append(stack, sel.Right)
				}
			} else {
				if sel.Right != nil {
					stack = append(stack, sel.Right)
				}
				if sel.Left != nil {
					stack = append(stack, sel.Left)
				}
			}
		}
		in = !in
		result = append(result, tr)
	}
	return result
}
