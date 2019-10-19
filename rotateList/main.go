package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	cur := head
	kcur := head
	count := 0
	for cur != nil {
		count++
		cur = cur.Next
	}
	if count <= k {
		k = k % count
	}
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	cur = head
	for i := 0; i < k; i++ {
		cur = cur.Next
	}
	for cur != nil && cur.Next != nil {
		cur = cur.Next
		kcur = kcur.Next
	}
	curList := []*ListNode{kcur}
	for kcur.Next != nil {
		kcur = kcur.Next
		curList = append(curList, kcur)
	}
	node := k - 1
	for i := 0; i < k; i++ {
		kcur.Next = head
		head = kcur
		kcur = curList[node]
		node--
	}
	kcur.Next = nil
	return head
}

func main() {
	input := []int{1, 2}
	k := 2
	list := getListChain(input)
	result := rotateRight(list, k)
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
