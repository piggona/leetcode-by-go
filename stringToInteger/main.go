package main

import "fmt"

func myAtoi(str string) int {
	result := 0
	flag := 1
	i := 0
	for {
		if i == len(str) {
			return 0
		}
		if str[i] == ' ' {
			i++
		} else {
			break
		}
	}
	if str[i] == '-' {
		flag = -1
		i++
	}
	if str[i] == '+' {
		flag = 1
		i++
	}
	for i < len(str) && int(str[i]) <= int('9') && int(str[i]) >= int('0') {
		result = result*10 + int(str[i]) - int('0')
		if result*flag >= 2147483648 {
			return 2147483647
		}
		if result*flag < -2147483648 {
			return -2147483648
		}
		i++
	}

	return result * flag
}

func main() {
	s := "9223372036854775808"
	fmt.Println(myAtoi(s))
}
