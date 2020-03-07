package verifypost

import "sort"

func verifyPostorder(postorder []int) bool {
	if len(postorder) == 1 || len(postorder) == 0 {
		return true
	}
	root := postorder[len(postorder)-1]
	var i int
	i = 0
	for i < len(postorder)-1 && postorder[i] <= root {
		i++
	}
	left := postorder[:i]
	right := postorder[i:(len(postorder) - 1)]
	l := oneSide(left, root)
	r := oneSide(right, root)
	return l && r && verifyPostorder(left) && verifyPostorder(right)

}

func oneSide(nums []int, target int) bool {
	tnums := make([]int, len(nums))
	copy(tnums, nums)
	if len(nums) <= 0 {
		return true
	}
	sort.Ints(tnums)
	if target > tnums[0] && target < tnums[len(tnums)-1] {
		return false
	}
	return true
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
