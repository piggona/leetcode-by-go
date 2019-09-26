package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	former := &ListNode{}
	former.Next = head
	result := former
	front := head
	back := head.Next
	for back != nil && front != nil {
		former.Next = back
		front.Next, back.Next = back.Next, front
		former = front
		front = front.Next
		if front == nil {
			break
		}
		back = front.Next
	}
	return result.Next
}

func main() {
	l1 := []int{1, 2, 3, 4}
	// l2 := []int{1, 3, 4}
	// l3 := []int{2, 6}
	list := []*ListNode{}

	list = append(list, getListChain(l1))
	// list = append(list, getListChain(l2))
	// list = append(list, getListChain(l3))

	fmt.Println("题目:")
	for _, item := range list {
		displayListChain(item)
	}
	fmt.Println("答案:")
	result := swapPairs(list[0])
	displayListChain(result)
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
