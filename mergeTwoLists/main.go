package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	if l1 == nil || l2 == nil {
		if l1 == nil {
			return l2
		} else {
			return l1
		}
	}
	if l1.Val < l2.Val {
		head = l1
		l1 = l1.Next
	} else {
		head = l2
		l2 = l2.Next
	}
	temp := head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			temp.Next = l1
			l1 = l1.Next
		} else {
			temp.Next = l2
			l2 = l2.Next
		}
		temp = temp.Next
	}
	if l1 == nil {
		temp.Next = l2
	} else {
		temp.Next = l1
	}
	return head
}

func main() {
	node1 := []int{}
	node2 := []int{}
	l1 := getListChain(node1)
	l2 := getListChain(node2)
	head := mergeTwoLists(l1, l2)
	displayListChain(head)

}

func getListChain(l []int) *ListNode {
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
