package main

import "fmt"

func mySqrt(x int) int {
	if x <= 1 {
		return x
	}
	return oneSqrt(x, x/2)
}

func oneSqrt(target, i int) int {
	cur := i * i
	next := (i + 1) * (i + 1)
	if cur <= target && next > target {
		return i
	}
	if next == target {
		return i + 1
	}
	if next < target {
		return oneSqrt(target, (target+i)/2)
	}
	return oneSqrt(target, i/2)
}

func mySqrtFinest(x int) int {
	if x == 1 {
		return 1
	}
	low, high := 0, x
	var mid int
	var sqr int

	for {
		mid = (low + high) / 2
		if mid == low {
			return mid
		}

		sqr = mid * mid
		if sqr == x {
			return mid
		}
		if sqr > x {
			high = mid
		}
		if sqr < x {
			low = mid
		}
	}
}

func main() {
	input := 1041
	fmt.Println(mySqrt(input))
}
