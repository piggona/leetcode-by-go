package strtoint

import (
	"testing"
)

func TestStrToInt(t *testing.T) {
	input := "9223372036854775808"
	res := strToInt(input)
	t.Errorf("%v", res)
}
