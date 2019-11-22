package main

import (
	"fmt"
	"math"
)

func largestRectangleAreaFinest(heights []int) int {
	area := 0
	// modify "heights" to ascending like: 4,3,2,4 ==> 2,2,2,4
	for i := 1; i < len(heights); i++ {
		if heights[i] >= heights[i-1] {
			continue
		}
		for j := i - 1; j >= 0; j-- {
			if heights[j] <= heights[i] {
				break
			}
			// calculate area before we modify heights[j]
			if curArea := heights[j] * (i - j); curArea > area {
				area = curArea
			}
			heights[j] = heights[i]
		}
	}
	// calculate max Area for new "heights"
	for i, h := range heights {
		if curArea := (len(heights) - i) * h; curArea > area {
			area = curArea
		}
	}
	return area
}

func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	return divideRec(heights, 0, len(heights)-1)
}

func divideRec(height []int, left, right int) int {
	if left >= right {
		return height[left]
	}
	val, pos := getMin(height, left, right)
	max := val * (right - left + 1)
	var rpos int
	var lpos int
	if pos-1 < left {
		rpos = pos + 1
	} else {
		rpos = pos
	}
	if right < pos+1 {
		lpos = pos - 1
	} else {
		lpos = pos
	}
	candidates := getMax(divideRec(height, left, rpos-1), divideRec(height, lpos+1, right))
	return getMax(max, candidates)
}

func getMax(i, j int) int {
	if i < j {
		return j
	}
	return i
}

func getMin(height []int, left, right int) (int, int) {
	min := math.MaxInt32
	pos := 0
	for i := left; i <= right; i++ {
		if height[i] < min {
			min = height[i]
			pos = i
		}
	}
	return min, pos
}

func largestRectangleAreaStack(height []int) int {
	height = append(height, 0)
	stack := []int{}
	i := 0
	max := 0
	for i < len(height) {
		if len(stack) == 0 || height[i] > height[stack[len(stack)-1]] {
			stack = append(stack, i)
			i++
		} else {
			temp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			var val int
			if len(stack) == 0 {
				val = height[temp] * i
			} else {
				val = height[temp] * (i - stack[len(stack)-1] - 1)
			}
			max = getMax(max, val)
		}
	}
	return max
}

func main() {
	heights := []int{2, 1, 6, 5, 2, 4}
	fmt.Println(largestRectangleAreaStack(heights))
}
