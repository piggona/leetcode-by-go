package main

import "fmt"

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	result := []int{1}
	str := con(result, n-1)
	return str
}

func con(result []int, remain int) string {
	if remain == 0 {
		str := []byte{}
		for _, num := range result {
			str = append(str, byte(num)+'0')
		}
		return string(str)
	}
	temp := []int{}
	count := 1
	for i := 0; i < len(result); i++ {
		if i+1 == len(result) || result[i] != result[i+1] {
			temp = append(temp, count, result[i])
			count = 1
		} else if result[i] == result[i+1] {
			count++
		}
	}
	return con(temp, remain-1)
}

func main() {
	n := 6
	fmt.Println(countAndSay(n))
}
