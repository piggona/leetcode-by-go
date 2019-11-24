package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head.Next == nil {
		return head
	}
	var ls, left, right, rs *ListNode
	cur := head
	for cur != nil {
		if cur.Val < x {
			if left != nil {
				left.Next = cur
			} else {
				ls = cur
			}
			left = cur
		} else {
			if right != nil {
				right.Next = cur
			} else {
				rs = cur
			}
			right = cur
		}
		cur = cur.Next
	}
	if left == nil || right == nil {
		return head
	}
	left.Next = rs
	right.Next = nil
	return ls
}

func main() {
	input := []int{5, 1, 4, 3, 2, 5, 2}
	x := 3
	head := getListChain(input)
	displayListChain(partition(head, x))
}

func getListChain(l []int) *ListNode {
	if len(l) == 0 {
		return nil
	}
	head := &ListNode{l[0], nil}
	former := head
	for i := 1; i < len(l); i++ {
		temp := &ListNode{l[i], nil}
		former.Next = temp
		former = temp
	}
	return head
}

func displayListChain(l *ListNode) {
	var temp *ListNode
	temp = l
	for temp != nil {
		fmt.Printf("%d ", temp.Val)
		temp = temp.Next
	}
	fmt.Printf("\n")
}
