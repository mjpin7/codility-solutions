package solution

func Solution(A []int, K int) []int {
	// write your code in Go 1.4
	s := make([]int, len(A))
	for i, n := range A {
		s[(i+K)%len(A)] = n
	}

	return s
}
