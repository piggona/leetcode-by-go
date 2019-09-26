package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	count, err := strconv.Atoi(input.Text())
	if err != nil {
		return
	}
	most := ""
	result := 0
	commodity := make(map[string]int)
	for i := 0; i < count; i++ {
		input.Scan()
		item := input.Text()
		commodity[item] = commodity[item] + 1
	}
	for item, count := range commodity {
		if count > result {
			most = item
			result = count
		}
	}
	fmt.Printf("%s %d", most, result)
}
