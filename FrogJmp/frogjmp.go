package solution

import "math"

func Solution(X int, Y int, D int) int {
    v := float64(Y) - float64(X)

    return int(math.Ceil(v / float64(D)))
}