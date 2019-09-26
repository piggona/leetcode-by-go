package main

import (
	"fmt"
	"io"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	for {
		flag := 0
		for i, item := range lists {
			if item == nil {
				lists = deleteNth(lists, i)
				break
			}
			if len(lists)-1 == i {
				flag = 1
			}
		}
		if flag == 1 || len(lists) == 0 {
			break
		}
	}
	if len(lists) == 0 {
		return nil
	}

	result := &ListNode{}
	former := &ListNode{}
	head := result

	pos := minInList(lists)
	node := lists[pos]
	head = node
	former = node

	if node.Next == nil {
		lists = deleteNth(lists, pos)
	} else {
		lists[pos] = node.Next
	}

	for len(lists) != 0 {
		pos := minInList(lists)
		node := lists[pos]
		former.Next = node
		former = node

		if node.Next == nil {
			lists = deleteNth(lists, pos)
		} else {
			lists[pos] = node.Next
		}
	}
	return head
}

func deleteNth(list []*ListNode, n int) []*ListNode {
	list[n] = list[len(list)-1]
	list = list[:len(list)-1]
	return list
}

func minInList(list []*ListNode) int {
	result := 65535
	pos := 0
	for i, item := range list {
		if item.Val < result {
			result = item.Val
			pos = i
		}
	}
	return pos
}

func main() {
	l1 := []int{1, 4, 5}
	l2 := []int{1, 3, 4}
	l3 := []int{2, 6}
	list := []*ListNode{}

	list = append(list, getListChain(l1))
	list = append(list, getListChain(l2))
	list = append(list, getListChain(l3))

	fmt.Println("题目:")
	for _, item := range list {
		displayListChain(item)
	}
	fmt.Println("答案:")
	displayListChain(mergeKLists(list))
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
	var te io.Writer
	fmt.Println(te)
	for temp != nil {
		fmt.Printf("%d ", temp.Val)
		temp = temp.Next
	}
	fmt.Printf("\n")
}
