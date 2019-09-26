package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	temp := &ListNode{}
	previous := &ListNode{}
	empty := &ListNode{}
	flag := 0

	flag = addNodes(result, l1, l2, flag)
	previous = result
	l1 = l1.Next
	l2 = l2.Next

	for _tk := false; _tk || (l1 != nil || l2 != nil); _tk = false {
		if l1 == nil {
			l1 = empty
		}
		if l2 == nil {
			l2 = empty
		}
		flag = addNodes(temp, l1, l2, flag)
		previous.Next = temp
		previous = temp
		temp = new(ListNode)
		l1 = l1.Next
		l2 = l2.Next
	}
	if flag != 0 {
		temp = new(ListNode)
		temp.Val = flag
		previous.Next = temp
	}
	return result

}

func addNodes(result *ListNode, l1 *ListNode, l2 *ListNode, flag int) int {
	result.Val = l1.Val + l2.Val + flag
	if result.Val > 9 {
		result.Val = result.Val % 10
		return 1
	}
	return 0
}

func sliceGetList(slice []int) *ListNode {
	result := &ListNode{}
	if len(slice) == 0 {
		return result
	}
	var previous *ListNode

	result.Val = slice[0]
	previous = result

	for i := 1; i < len(slice); i++ {
		temp := &ListNode{}
		temp.Val = slice[i]
		previous.Next = temp
		previous = temp
	}

	return result
}

func displayList(list *ListNode) {
	for _tk := false; _tk || list != nil; _tk = false {
		fmt.Printf("%d ,", list.Val)
		list = list.Next
	}
}
func main() {
	l1 := []int{5, 6, 1}
	l2 := []int{5, 7, 9}
	list1 := sliceGetList(l1)
	list2 := sliceGetList(l2)
	result := addTwoNumbers(list1, list2)
	displayList(result)
}
