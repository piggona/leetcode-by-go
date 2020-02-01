package decode

import "testing"

func TestDecode(t *testing.T) {
	input := "abcsfs"
	if s := decodeString(input); s != "abcsfs" {
		t.Errorf("Wrong Answer: %v", s)
	}
}
