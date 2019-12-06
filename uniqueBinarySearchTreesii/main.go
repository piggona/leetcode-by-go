package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// func generateTreesBad(n int) []*TreeNode {
// 	num := 1
// 	depth := 1
// 	root := &TreeNode{}
// 	stack := []*TreeNode{root}
// 	result := []*TreeNode{root}
// 	for i := 0; i < 2; i++ {
// 		if i == 0 {
// 			// 1.赋值弹出
// 			temp := stack[len(stack)-1]
// 			stack = stack[:len(stack)-1]
// 			temp.Val = num
// 			// 1.1推入右边元素
// 			// 继续递归
// 		} else {
// 			// 2.推入左边元素
// 		}
// 	}
// 	for num <= n {

// 	}
// }

// func treeRecurBad(num, n int, stack []*TreeNode, root *TreeNode) []*TreeNode {
// 	for i := 0; i < 2; i++ {
// 		if i == 0 {
// 			// 1.赋值弹出
// 			temp := stack[len(stack)-1]
// 			stack = stack[:len(stack)-1]
// 			temp.Val = num
// 			if len(stack) == 0 && num == n {
// 				return root
// 			} else if num == n {
// 				break
// 			} else if len(stack) == 0 {
// 				break
// 			}

// 			// 1.1推入右边元素
// 			troot = &TreeNode{}
// 			right := &TreeNode{}
// 			temp.Right = right
// 			tlist := treeRecur(num+1, n, append(stack, right))
// 			// 继续递归
// 		} else {
// 			// 2.推入左边元素
// 		}
// 	}
// 	return nil
// }

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	return generator(1, n)
}

func generator(left, right int) []*TreeNode {
	if left > right {
		return []*TreeNode{nil}
	}

	list := make([]*TreeNode, 0)
	for i := left; i <= right; i++ {
		leftList := generator(left, i-1)
		rightList := generator(i+1, right)
		for _, l := range leftList {
			for _, r := range rightList {
				list = append(list, &TreeNode{
					Val:   i,
					Left:  l,
					Right: r,
				})
			}
		}
	}
	return list
}

func main() {
	return
}
