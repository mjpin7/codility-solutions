/**
*
* BinaryGap. Link: https://app.codility.com/programmers/lessons/1-iterations/binary_gap/
*
*/

package solution

import (
    "strconv"
    "strings"
)


func Solution(N int) int {
    n := strconv.FormatInt(int64(N), 2)
    num := 0

	// Not sure I like this "infinite loop"
    for {
		// Get index of first and second "1" in string
        i := strings.Index(n, "1")
        j := strings.Index(n[i+1:], "1")

        if j == -1 {
            break;
        }

        if j - i > num {
            num = j - i
        }
        n = n[j+1:]
    }

    return num
}