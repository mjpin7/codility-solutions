/**
*
* StoneWall. Link: https://app.codility.com/programmers/lessons/7-stacks_and_queues/stone_wall/
*
*/

package solution

type Stack []int

func (s *Stack) isEmpty() bool {
    return len(*s) == 0
}

func (s *Stack) push(n int) {
    *s = append(*s, n)
}

func (s *Stack) pop() (int, bool){
    if s.isEmpty() {
        return 0, true
    } else {
        index := len(*s) - 1
        e := (*s)[index]
        *s = (*s)[:index]
        return e, false
    }
}

func Solution(H []int) int {

    var s Stack
    var num int

    for i := 0; i < len(H); i++ {
        for !s.isEmpty() && s[len(s) - 1] > H[i] {
            s.pop()
        }

        if !s.isEmpty() && s[len(s) - 1] == H[i] {
            continue
        } else {
            s.push(H[i])
            num++
        }
    }
    
    return num
}