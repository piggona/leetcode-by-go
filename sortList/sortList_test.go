package sortList

import (
	"fmt"
	"testing"
)

func TestSortList(t *testing.T) {
	head1 := GetListChain([]int{2, 4})
	head2 := GetListChain([]int{1, 3})
	res := merge(head1, head2)
	DisplayListChain(res)
}

func GetListChain(l []int) *ListNode {
	head := &ListNode{l[0], nil}
	former := head
	for i := 1; i < len(l); i++ {
		temp := &ListNode{l[i], nil}
		former.Next = temp
		former = temp
	}
	return head
}

func DisplayListChain(l *ListNode) {
	var temp *ListNode
	temp = l
	for temp != nil {
		fmt.Printf("%d ", temp.Val)
		temp = temp.Next
	}
	fmt.Printf("\n")
}
