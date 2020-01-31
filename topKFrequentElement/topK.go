package topK

import "container/heap"

type statistic struct {
	keys    []int
	hashmap map[int][]int
	count   int
}

func (c *statistic) Len() int {
	return c.count
}

func (c *statistic) Less(a, b int) bool {
	i := c.keys[a]
	j := c.keys[b]
	left := c.hashmap[i]
	right := c.hashmap[j]
	if left[0] > right[0] {
		return true
	}
	return false
}

func (c *statistic) Swap(a, b int) {
	i := c.keys[a]
	j := c.keys[b]
	tleft := c.hashmap[i]
	tright := c.hashmap[j]
	left := tleft[1]
	right := tright[1]
	c.keys[left], c.keys[right] = c.keys[right], c.keys[left]
	c.hashmap[i][1], c.hashmap[j][1] = right, left
}

func (c *statistic) Push(b interface{}) {
	x := b.(int)
	if pos, ok := c.hashmap[x]; ok {
		pos[0]++
	}
	c.keys = append(c.keys, x)
	c.hashmap[x] = []int{1, len(c.keys) - 1}
	c.count++
}

func (c *statistic) Pop() interface{} {
	res := c.keys[len(c.keys)-1]
	c.keys = c.keys[:(len(c.keys) - 1)]
	delete(c.hashmap, res)
	c.count--
	return res
}

func topKFrequent(nums []int, k int) []int {
	keys := []int{}
	hashmap := make(map[int][]int)
	count := 0
	for i := 0; i < len(nums); i++ {
		if val, ok := hashmap[nums[i]]; ok {
			val[0]++
		} else {
			keys = append(keys, nums[i])
			hashmap[nums[i]] = []int{1, len(keys) - 1}
			count++
		}
	}
	body := statistic{
		keys:    keys,
		hashmap: hashmap,
		count:   count,
	}
	heap.Init(&body)
	res := []int{}
	for i := 0; i < k; i++ {
		temp := heap.Pop(&body)
		res = append(res, temp.(int))
	}
	return res
}
