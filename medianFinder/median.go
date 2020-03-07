package medianFinder

import "container/heap"

type MedianFinder struct {
	leftHeap  *bigHeap
	rightHeap *smallHeap
	isLeft    bool
}

type bigHeap []int

type smallHeap []int

func (b *bigHeap) Len() int {
	return len(*b)
}

func (b *bigHeap) Swap(i, j int) {
	(*b)[i], (*b)[j] = (*b)[j], (*b)[i]
}

func (b *bigHeap) Less(i, j int) bool {
	if (*b)[i] <= (*b)[j] {
		return false
	}
	return true
}

func (b *bigHeap) Pop() interface{} {
	pos := len(*b) - 1
	top := (*b)[pos]
	(*b) = (*b)[:pos]
	return top
}

func (b *bigHeap) Push(x interface{}) {
	val := x.(int)
	(*b) = append((*b), val)
}

func (b *smallHeap) Len() int {
	return len(*b)
}

func (b *smallHeap) Swap(i, j int) {
	(*b)[i], (*b)[j] = (*b)[j], (*b)[i]
}

func (b *smallHeap) Less(i, j int) bool {
	if (*b)[i] >= (*b)[j] {
		return false
	}
	return true
}

func (b *smallHeap) Pop() interface{} {
	pos := len(*b) - 1
	top := (*b)[pos]
	(*b) = (*b)[:pos]
	return top
}

func (b *smallHeap) Push(x interface{}) {
	val := x.(int)
	(*b) = append(*b, val)
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		leftHeap:  &bigHeap{},
		rightHeap: &smallHeap{},
		isLeft:    true,
	}
}

func (this *MedianFinder) AddNum(num int) {
	isL := this.isLeft
	if isL {
		heap.Push(this.rightHeap, num)
		x := heap.Pop(this.rightHeap)
		heap.Push(this.leftHeap, x)
	} else {
		heap.Push(this.leftHeap, num)
		x := heap.Pop(this.leftHeap)
		heap.Push(this.rightHeap, x)
	}
	this.isLeft = !isL
}

func (this *MedianFinder) FindMedian() float64 {
	if this.isLeft {
		x1 := heap.Pop(this.leftHeap)
		x2 := heap.Pop(this.rightHeap)
		return (float64(x1.(int)) + float64(x2.(int))) / 2.0
	} else {
		x := heap.Pop(this.leftHeap)
		return float64(x.(int))
	}
}
