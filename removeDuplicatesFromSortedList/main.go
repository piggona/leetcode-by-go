package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	cur := head
	next := cur.Next
	for next != nil {
		if cur.Val != next.Val {
			cur.Next = next
			cur = cur.Next
			next = next.Next
		} else {
			next = next.Next
		}
	}
	cur.Next = next
	return head
}

func main() {
	input := []int{1, 1, 1, 2, 3}
	head := getListChain(input)
	result := deleteDuplicates(head)
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
