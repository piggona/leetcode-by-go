package main

import (
	"container/heap"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type Item struct {
	Val   int
	index int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Val < pq[j].Val // 这样的写法结果是从大到小排列的，Pop也要推出最后一个值（最小值）
	// return pq[j].Val < pq[i].Val // 这样的写法是从小到大排列
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) { // 因为这个操作包含对数组head的移动，所以要使用数组的指针
	// 因为要面对所有可能的输入值，所以使用interface{}类型的输入
	n := len(*pq)
	var item *Item
	if val, ok := x.(*Item); ok { // 如果允许所有类型输入，就一定要有这一步判断，将*Item类型断言给val
		item = val
		item.index = n
		*pq = append(*pq, item)
	}
}

func (pq *PriorityQueue) Pop() interface{} {
	n := len(*pq)
	item := (*pq)[n-1] // 对指针值（尤其是切片）做indexing时一定要有括号(*pq)
	item.index = -1
	*pq = (*pq)[:n-1]
	return item
}

func mergeKLists(lists []*ListNode) *ListNode {
	var arr []int
	for _, list := range lists {
		for list != nil {
			arr = append(arr, list.Val)
			list = list.Next
		}
	}
	pq := make(PriorityQueue, len(arr))
	for i, val := range arr {
		pq[i] = &Item{val, i}
	}

	heap.Init(&pq)
	res := new(ListNode)
	cur := res
	for len(pq) > 0 {
		cur.Next = &ListNode{heap.Pop(&pq).(*Item).Val, nil}
		// heap.Pop(&pq)返回的是一个interface{}类型的变量，如果我们要获取其Val属性(*Item才有Val属性)，
		// 就要使用断言将其转换为具体类型*Item
		cur = cur.Next
	}
	return res.Next
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
	for temp != nil {
		fmt.Printf("%d ", temp.Val)
		temp = temp.Next
	}
	fmt.Printf("\n")
}
