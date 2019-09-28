package main

import (
	"fmt"
	"strings"
)

func multiply(num1 string, num2 string) string {
	if len(num1) <= 0 || len(num2) <= 0 || num1 == "0" || num2 == "0" {
		return "0"
	}
	step := make([]int, len(num1)+len(num2))
	for i := 0; i < len(num1); i++ {
		for j := 0; j < len(num2); j++ {
			v1 := num1[len(num1)-i-1] - '0'
			v2 := num2[len(num2)-j-1] - '0'
			temp := v1 * v2
			v1 = temp / 10
			v2 = temp % 10
			step[i+j] += int(v2)
			step[i+j+1] += int(v1)
		}
	}
	for i, num := range step {
		if num > 9 {
			v1 := num / 10
			v2 := num % 10
			step[i] = v2
			step[i+1] += v1
		}
	}
	result := strings.Builder{}
	i := len(num1) + len(num2) - 1
	for step[i] == 0 && i >= 0 {
		i--
	}
	for i >= 0 {
		result.WriteByte(byte(step[i]) + '0')
		i--
	}
	return result.String()
}

func main() {
	num1 := "2"
	num2 := "3"
	fmt.Println(multiply(num1, num2))
}
