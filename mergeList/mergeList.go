package mergeList

func merge(A []int, m int, B []int, n int) {
	a := m - 1
	b := n - 1
	for i := len(A) - 1; i >= 0; i-- {
		if a >= 0 && b >= 0 {
			if A[a] > B[b] {
				A[i] = A[a]
				a--
			} else if A[a] <= B[b] {
				A[i] = B[b]
				b--
			}
		} else if a >= 0 || b >= 0 {
			if b >= 0 {
				A[i] = B[b]
				b--
			} else {
				A[i] = A[a]
				a--
			}
		}
	}
}
