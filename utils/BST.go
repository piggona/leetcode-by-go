package utils

import (
	"fmt"
	"os"
)

func constructTree(root *TreeNode, list []int) {
	root.Val = list[0]
	stack := []*TreeNode{root}
	for i := 1; i < len(list); i += 2 {
		top := stack[0]
		if len(stack) == 0 {
			stack = []*TreeNode{}
		} else {
			stack = stack[1:]
		}
		if list[i] != -1 {
			top.Left = &TreeNode{Val: list[i]}
			stack = append(stack, top.Left)
		}
		if list[i+1] != -1 {
			top.Right = &TreeNode{Val: list[i+1]}
			stack = append(stack, top.Right)
		}
	}
}

func displayTree(root *TreeNode) {
	stack := []*TreeNode{root}
	for len(stack) != 0 {
		top := stack[0]
		if len(stack) == 0 {
			stack = []*TreeNode{}
		} else {
			stack = stack[1:]
		}
		// fmt.Fprintf(os.Stdout, "%v ", top.Val)
		if top != nil {
			fmt.Fprintf(os.Stdout, "%v ", top.Val)
			stack = append(stack, top.Left)
			stack = append(stack, top.Right)
		} else {
			fmt.Fprintf(os.Stdout, "%v ", top)
		}
		// if top.Left != nil {
		// 	stack = append(stack, top.Left)
		// }
		// if top.Right != nil {
		// 	stack = append(stack, top.Right)
		// }
	}
}
