package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicatesOnce(head *ListNode) *ListNode {
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

func deleteDuplicatesI(head *ListNode) *ListNode {
	anchor := &ListNode{
		Val:  -1,
		Next: nil,
	}
	top := anchor
	anchor.Next = head
	pre := anchor
	cur := pre.Next
	next := cur.Next
	for next != nil {
		if pre.Val != cur.Val && cur.Val != next.Val {
			anchor.Next = cur
			anchor = cur
		}
		pre = pre.Next
		cur = cur.Next
		next = next.Next
	}
	if pre.Val != cur.Val {
		anchor.Next = cur
		anchor = cur
	} else {
		anchor.Next = nil
	}
	return top.Next

}
func deleteDuplicates(head *ListNode) *ListNode {
	dummyHead := ListNode{Next: head}
	prev := &dummyHead
	cursor := dummyHead.Next
	for cursor != nil && cursor.Next != nil {
		p := cursor
		for p != nil && p.Next != nil && p.Val == p.Next.Val {
			p = p.Next
		}
		if p != cursor {
			prev.Next = p.Next
			cursor = p.Next
		} else {
			prev = cursor
			cursor = cursor.Next
		}
	}
	return dummyHead.Next
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
