package solution

import "sort"

func Solution(A []int) int {
    sort.Ints(A)
    max := 1

    for _, v := range A {
        if v == max {
            max += 1
        }

        if v > max {
            break;
        }
    }

    return max
}