package main

import "fmt"

func sortColors(nums []int) {
	left := 0
	right := len(nums) - 1
	i := 0
	for i <= right {
		switch nums[i] {
		case 0:
			nums[i], nums[left] = nums[left], nums[i]
			i++
			left++
			break
		case 1:
			i++
			break
		case 2:
			nums[i], nums[right] = nums[right], nums[i]
			right--
			break
		}
	}
}

func sortColorsFinest(nums []int) {
	i := -1
	j := -1
	k := -1
	for p := 0; p < len(nums); p++ {
		switch nums[p] {
		case 0:
			i++
			j++
			k++
			nums[k] = 2
			nums[j] = 1
			nums[i] = 0
			break
		case 1:
			j++
			k++
			nums[k] = 2
			nums[j] = 1
			break
		case 2:
			k++
			nums[k] = 2
			break
		}
	}
}

func main() {
	input := []int{2, 0, 2, 0, 2, 1, 1, 0}
	sortColorsFinest(input)
	fmt.Println(input)
}
