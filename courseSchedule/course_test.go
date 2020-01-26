package courseSchedule

import (
	"fmt"
	"testing"
)

func TestCourse(t *testing.T) {
	number := 3
	input := [][]int{
		{1, 0},
	}
	res := canFinish(number, input)
	fmt.Println(res)
}
