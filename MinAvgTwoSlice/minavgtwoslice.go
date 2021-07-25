package solution

func Solution(A []int) int {
    minIndex := 0
    var minAvg float32 = float32(A[0] + A[1] / 2)
    var avg1 float32
    var avg2 float32

    for i := 0; i < len(A) - 2; i++ {
        avg1 = float32(A[i] + A[i + 1] / 2)
        avg2 = float32(A[i] + A[i + 1] + A[i + 2] / 3)
        var curAvg float32

        if avg1 < avg2 {
            curAvg = avg1
        } else {
            curAvg = avg2
        }

        if curAvg <  minAvg {
            minIndex = i
            minAvg = curAvg
        }
    }

    if float32(A[len(A) - 2] + A[len(A) - 1] / 2) < minAvg {
        minIndex = len(A) - 2
    }

    return minIndex
}

