/**
*
* PermMissingElem. Link: https://app.codility.com/programmers/lessons/3-time_complexity/perm_missing_elem/
*
*/

package solution

import "sort"

func Solution(A []int) int {
    sort.Ints(A)
    var num int

    if len(A) == 0 {
        num = 1
    } else {
        if A[0] != 1 {
            num = 1
        } else {
            for i := 0; i < len(A); i++ {
                if i == len(A) - 1 || A[i+1] - A[i] > 1 {
                    num = A[i] + 1
                    break;
                }
            }
        }
    }

    return num
}