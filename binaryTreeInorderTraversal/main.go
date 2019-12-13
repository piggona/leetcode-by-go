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

func inorderTraversalMorris(root *TreeNode) []int {
	result := []int{}
	cur := root
	for cur != nil {
		if cur.Left == nil {
			result = append(result, cur.Val)
			cur = cur.Right
		} else {
			temp := getFrontier(cur)
			if temp.Right == cur {
				temp.Right = nil
				result = append(result, cur.Val)
				cur = cur.Right
			} else if temp.Right == nil {
				temp.Right = cur
				cur = cur.Left
			}
		}
	}
	return result
}

func preorderTraversalMorris(root *TreeNode) []int {
	result := []int{}
	cur := root
	for cur != nil {
		if cur.Left == nil {
			result = append(result, cur.Val)
			cur = cur.Right
		} else {
			temp := getFrontier(cur)
			if temp.Right == cur {
				temp.Right = nil
				cur = cur.Right
			} else if temp.Right == nil {
				temp.Right = cur
				result = append(result, cur.Val)
				cur = cur.Left
			}
		}
	}
	return result
}

func postorderTraversalMorris(root *TreeNode) []int {
	result := []int{}
	dump := &TreeNode{Left: root}
	cur := dump
	for cur != nil {
		if cur.Left == nil {
			cur = cur.Right
		} else {
			temp := getFrontier(cur)
			if temp.Right == nil {
				temp.Right = cur
				cur = cur.Left
			} else if temp.Right == cur {
				reverse(cur.Left, temp)
				for temp != cur.Left {
					result = append(result, temp.Val)
					temp = temp.Right
				}
				result = append(result, temp.Val)
				temp.Right = nil
				cur = cur.Right
			}
		}
	}
	return result
}

func reverse(from, to *TreeNode) {
	if from == to {
		return
	}
	x, y := from, from.Right
	var z *TreeNode
	for x != to {
		z = y.Right
		y.Right = x
		x.Right = z
		x = y
		y = z
	}
}

func getFrontier(root *TreeNode) *TreeNode {
	result := root.Left
	for result.Right != nil && result.Right != root {
		result = result.Right
	}
	return result
}

func main() {
	input := []int{1, -1, 2, 3}
	fmt.Println(postorderTraversalMorris(getTree(input)))
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
