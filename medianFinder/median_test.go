package medianFinder

import "testing"

func TestMedian(t *testing.T) {
	m := Constructor()
	m.AddNum(1)
	m.AddNum(2)
	m.AddNum(3)
	res := m.FindMedian()
	t.Errorf("%v", res)
}
