package main

import (
	"fmt"
	"math"
)

func divide(divided int, divisor int) int {
	if divided == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	result := 0
	sign := -1
	if divided > 0 && divisor > 0 || divided < 0 && divisor < 0 {
		sign = 1
	}
	dvd, dvs := abs64(divided), abs64(divisor)
	for dvd >= dvs {
		temp := dvs
		m := 1
		for temp<<1 <= dvd {
			temp <<= 1
			m <<= 1
		}
		dvd -= temp
		result += m
	}
	return sign * result
}

func abs64(n int) int64 {
	ret := int64(n)
	if ret < 0 {
		return -ret
	}
	return ret
}

// func main() {
// 	divided := -7
// 	divisor := 3
// 	fmt.Println(divide(divided, divisor))
// }

func main() {
	x := 22 // 10110
	n := 3
	tar := pow(2, n-1)
	fmt.Println(x ^ tar) // 10010
}

func pow(x, n int) int {
	result := 1
	for n > 0 {
		result = result * x
		n--
	}
	return result
}
