package main

import (
	"container/heap"
	"fmt"
)

type hNum []int

func (h hNum) Len() int {
	return len(h)
}

func (h hNum) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h hNum) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h *hNum) Push(x interface{}) {
	if num, ok := x.(int); ok {
		*h = append(*h, num)
	}
}

func (h *hNum) Pop() interface{} {
	n := len(*h)
	temp := (*h)[n-1]
	*h = (*h)[:n-1]
	return temp
}

func firstMissingPositive(nums []int) int {
	if len(nums) == 0 {
		return 1
	}
	h := hNum(nums)
	heap.Init(&h)
	temp := 0
	former := 0
	for h.Len() > 0 {
		temp = heap.Pop(&h).(int)
		if temp > 0 && former > 0 && temp-former > 1 {
			return former + 1
		}
		if former <= 0 && temp > 1 {
			return 1
		}
		former = temp
	}
	if temp >= 0 {
		return temp + 1
	}
	return 1
}

func main() {
	nums := []int{-5}
	fmt.Println(firstMissingPositive(nums))
}
