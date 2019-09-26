package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isomorphism(s string) bool {
	if len(s) == 0 {
		return true
	}
	str := strings.Split(s, ";")
	first, second := []byte(str[0]), []byte(str[1])

	name := 1
	con := make(map[byte]byte)
	con1 := make(map[byte]byte)
	for i := 0; i < len(first); i++ {
		if con[first[i]] == 0 {
			first[i] = byte(name)
			con[first[i]] = byte(name)
		} else {
			first[i] = con[first[i]]
		}
		if con1[second[i]] == 0 {
			second[i] = byte(name)
			con1[second[i]] = byte(name)
		} else {
			second[i] = con1[second[i]]
		}
		name++
	}
	return string(first) == string(second)
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	s := "ababa;ststs"
	for input.Scan() {
		s = input.Text()
	}
	fmt.Println(isomorphism(s))
}
