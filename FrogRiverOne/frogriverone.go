/**
*
*  FrogRiverOne. Link: https://app.codility.com/programmers/lessons/4-counting_elements/frog_river_one/
*
*/

package solution

func Solution(X int, A []int) int {
    var total int
    var sum int
	// Slice to keep track of what was found
    found := make([]int, X)
	// Get the summation of (1, X)
    for i := 1; i <= X; i++{
        total+=i
    }

	/**
	* Idea here:
	* Each time you loop, check if you have seen this number before
	* If you havent, add it to the total sum and mark it as found
	* If you have hit every number between (1, X) then the total sum would equal the summation of (1, X) so return the index
	* If the total sum has not hit the summation, then not all of the positions have been found
	*/
    for i, v := range A {
        if found[v-1] == 0 {
            sum += v
            found[v-1]++

            if sum == total {
                return i
            }
        }
    }

    return -1
}