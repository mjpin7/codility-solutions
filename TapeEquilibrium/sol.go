/**
*
*  TapeEquilibrium. Link: https://app.codility.com/programmers/lessons/3-time_complexity/tape_equilibrium/
*
*/

package solution

import "math"

func Solution(A []int) int {
    var min = math.MaxFloat64
    var currentMin = math.MaxFloat64
    var totalSum float64
    var firstSum float64
    var secondSum float64

    for _, v := range A {
        totalSum += float64(v)
    }

    for _, v := range A {
        firstSum += float64(v)
        secondSum = totalSum - firstSum
        currentMin = math.Abs(firstSum-secondSum)
        min = math.Min(currentMin, min)
    }

    return int(min)
}