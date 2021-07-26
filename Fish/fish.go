/**
*
* Fish. Link: https://app.codility.com/programmers/lessons/7-stacks_and_queues/fish/
*
*/

package solution

type Fish struct {
    size int
    direction int
}

type Stack []Fish

func (s *Stack) isEmpty() bool {
    return len(*s) == 0
}

func (s *Stack) push(f Fish) {
    *s = append(*s, f)
}

func (s *Stack) pop() (Fish, bool){
    if s.isEmpty() {
        return Fish{}, true
    } else {
        index := len(*s) - 1
        e := (*s)[index]
        *s = (*s)[:index]
        return e, false
    }
}

func Solution(A []int, B []int) int {

    var s Stack
    num := len(A)


    /**
    * IDEA: 
    * Iterate through the lists. The "important fish" are going downstream
    * Every downstream fish encountered, push it onto the stack
    * Then whenever we find a fish going upstream, we simulate it "swimming through" all of the fish on the stack
    *
    */
    for i := 0; i < len(A); i++ {
        if B[i] == 1 {
            s.push(Fish{size: A[i], direction: B[i]})
        } else {
            for !s.isEmpty() {
                if s[len(s) - 1].size > A[i] {
                    num--
                    break
                } else {
                    s.pop()
                    num--
                }
            }
        }
    }
    
    return num
}
