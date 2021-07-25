/**
*
* MaxCounters. Link: hhttps://app.codility.com/programmers/lessons/4-counting_elements/max_counters/
*
*/

package solution

func Solution(N int, A []int) []int {
    result := make([]int, N)
    var max int

    for _, v := range A {
        if 1 <= v && v <= N {
            result[v-1]++
            if result[v-1] > max {
                max = result[v-1]
            }
        } else if v == N + 1 {
            for i := range result {
                result[i] = max
            }
        }
    }

    return result
}

