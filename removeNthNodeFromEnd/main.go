package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	cur := &ListNode{}
	left := &ListNode{}

	// 算法
	cur = head
	count := 0
	// 找到第n+1个元素，定位倒数n+1个元素，如果定位不到，那么就去掉第一个元素
	for count < n {
		if cur.Next == nil {
			return head.Next
		}
		cur = cur.Next
		count++
	}
	left = head
	for cur.Next != nil {
		cur = cur.Next
		left = left.Next
	}
	if n == 1 {
		left.Next = nil
	} else {
		temp := left.Next.Next
		left.Next = temp
	}
	return head
}

func main() {
	head := &ListNode{1, nil}
	former := head
	for i := 2; i <= 2; i++ {
		temp := &ListNode{i, nil}
		former.Next = temp
		former = temp
	}
	head = removeNthFromEnd(head, 2)
	for temp := head; temp != nil; temp = temp.Next {
		fmt.Println(temp.Val)
	}
	fmt.Println(strconv.Itoa(9))

}
