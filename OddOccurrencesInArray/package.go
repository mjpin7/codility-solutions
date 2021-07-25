/**
*
* OddOccurrencesInArray. Link: https://app.codility.com/programmers/lessons/2-arrays/odd_occurrences_in_array/
*
*/

package solution

func Solution(A []int) int {
    var max int
    var num int

    for i, v := range A {
        if i == 0 || v > max {
            max = v
        }
    }
    foundVals := make([]int, max+1)

    for _, v := range A {
        foundVals[v]++
    }

    for i, v := range foundVals {
        if v % 2 != 0 {
            num = i
            break;
        }
    }

    return num
}