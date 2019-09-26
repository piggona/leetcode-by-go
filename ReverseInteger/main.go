package main

import "fmt"

func reverse(x int) int {
	temp := 0
	result := 0
	flag := 1
	if x < 0 {
		x = -x
		flag = -1
	}
	if x < 10 {
		return flag * x
	}

	// 算法
	for {
		temp = x % 10
		x = x / 10
		result = result*10 + temp
		if x == 0 {
			break
		}
	}
	return result * flag
}

func main() {
	i := 120
	fmt.Println(reverse(i))
}
