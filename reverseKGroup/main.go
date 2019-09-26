package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	result := &ListNode{}
	result.Next = head
	left := head
	former := result
	pre := &ListNode{}
	temp := head

	// 计算链表长度
	count := 0
	for temp != nil {
		temp = temp.Next
		count++
	}
	count = count / k

	for j := 0; j < count; j++ {
		// 倒置
		for i := 0; i < k-1; i++ {
			temp = left.Next
			pre = former.Next
			former.Next = temp
			left.Next = temp.Next
			temp.Next = pre
		}
		former = left
		left = left.Next
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
	result := reverseKGroup(list[0], 1)
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
