package solution

import "sort"

func Solution(A []int) int {
    sort.Ints(A)
    var num int
    var lastNum int

    for i, v := range A {
        if i == 0 || v != lastNum {
            num++
        }

        lastNum = v
    }

    return num
}