package solution

func Solution(A []int) int {
    N := len(A)
    found := make([]int, N)

    var total int
    var sum int
    for i := 1; i <= N; i++ {
        total += i
    }

    for _, v := range A {
        if v <= N && found[v-1] == 0 {
            found[v-1] = 1
            sum += v

            if sum == total {
                return 1
            }
        }
    }

    return 0
}