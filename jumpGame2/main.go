package main

import "fmt"

func canJump(nums []int) bool {
	pos := 0
	base := 0
	for pos < len(nums)-1 {
		temp := pos + nums[pos]
		for i := base; i < pos; i++ {
			if i+nums[i] > temp {
				temp = i + nums[i]
				pos = i
			}
		}
		if pos < len(nums)-1 && nums[pos] == 0 {
			return false
		}
		base = pos
		pos = temp
	}
	return true
}

func main() {
	input := []int{3, 2, 1, 0, 4}
	fmt.Println(canJump(input))
}
