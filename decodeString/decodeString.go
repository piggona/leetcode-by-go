package decode

import (
	"bytes"
	"strconv"
	"unicode"
)

type Stack struct {
	data []interface{}
	size int
}

func (s *Stack) Pop() interface{} {
	if s.size > 0 {
		ret := s.data[s.size-1]
		s.data = s.data[:s.size-1]
		s.size--
		return ret
	}
	return nil
}

func (s *Stack) Push(x interface{}) {
	s.size++
	s.data = append(s.data, x)
}

func NewStack() *Stack {
	return &Stack{
		data: make([]interface{}, 0),
		size: 0,
	}
}

// assume last byte must be ']' or letter
func decodeString(s string) string {
	si := []byte(s)
	times := NewStack()
	storage := make([][]byte, 1)
	storage[0] = []byte{}
	bracket := 0
	for i := 0; i < len(s); i++ {
		v := s[i]
		if v >= '0' && v <= '9' {
			start, end := i, i+1
			for s[end] >= '0' && s[end] <= '9' {
				i++
				end++
			}
			t, _ := strconv.Atoi(s[start:end])
			times.Push(t)
		} else if v == '[' {
			// if no numbers before '[', set times 0
			if i == 0 || !(s[i-1] >= '0' && s[i-1] <= '9') {
				times.Push(0)
			}
			bracket++
			if len(storage) < bracket+1 {
				storage = append(storage, []byte{})
			}
		} else if v == ']' {
			if rt := times.Pop(); rt != nil {
				rti := rt.(int)
				storage[bracket-1] = append(storage[bracket-1], bytes.Repeat(storage[bracket], rti)...)
				storage[bracket] = []byte{}
			}
			bracket--
		} else {
			start, end := i, i+1
			for end < len(s) && unicode.IsLetter(rune(s[end])) {
				i++
				end++
			}
			storage[bracket] = append(storage[bracket], si[start:end]...)
		}
	}
	return string(storage[0])
}
