package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}
	if root.Left != nil {
		temp := inorderTraversal(root.Left)
		result = append(result, temp...)
	}
	result = append(result, root.Val)
	if root.Right != nil {
		temp := inorderTraversal(root.Right)
		result = append(result, temp...)
	}
	return result
}

func inorderTraversalStack(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}
	stack := []*TreeNode{root}
	for len(stack) != 0 {
		top := stack[len(stack)-1]
		if top.Left != nil {
			//入栈
			stack = append(stack, top.Left)
			top.Left = nil
		} else {
			result = append(result, top.Val)
			stack = stack[:len(stack)-1]
			if top.Right != nil {
				stack = append(stack, top.Right)
			}
		}
	}
	return result
}

func main() {
	input := []int{1, -1, 2, 3}
	fmt.Println(inorderTraversalStack(getTree(input)))
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
