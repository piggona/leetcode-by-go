package main

import "fmt"

func romanToInt(s string) int {
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	integers := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	i := 0
	result := 0
	for i < len(s) {
		for j := 0; j < len(romans); j++ {
			if string(s[i]) == romans[j] {
				result += integers[j]
				i++
				break
			} else if i+2 < len(s) && s[i:i+2] == romans[j] {
				result += integers[j]
				i = i + 2
				break
			} else if i+2 == len(s) && s[i:] == romans[j] {
				result += integers[j]
				i = i + 2
				break
			}
		}
	}
	return result
}

func main() {
	s := "MCMXCIV"
	fmt.Println(romanToInt(s))
}
