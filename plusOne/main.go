package main

import "fmt"

func plusOne(digits []int) []int {
	pos := len(digits) - 1
	digits[pos] = digits[pos] + 1
	for pos >= 0 && digits[pos] > 9 {
		temp := digits[pos]
		digits[pos] = temp % 10
		pos--
		if pos < 0 {
			break
		}
		digits[pos] = digits[pos] + 1
	}
	if pos < 0 {
		result := []int{1}
		return append(result, digits...)
	}
	return digits

}

func main() {
	digits := []int{0}
	fmt.Println(plusOne(digits))
}
