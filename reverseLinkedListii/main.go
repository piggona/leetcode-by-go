package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	top := &ListNode{Next: head}
	prev := top
	cur := head
	next := cur.Next
	i := 1
	for cur != nil && i < m {
		next = cur.Next.Next
		cur = cur.Next
		prev = prev.Next
		i++
	}
	for cur != nil && i < n {
		temp := prev.Next
		prev.Next = next
		cur.Next = next.Next
		next.Next = temp
		next = cur.Next
		i++
	}
	return head
}

func main() {
	input := []int{1, 2, 3, 4, 5}
	m := 2
	n := 4
	head := getListChain(input)
	displayListChain(reverseBetween(head, m, n))
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
