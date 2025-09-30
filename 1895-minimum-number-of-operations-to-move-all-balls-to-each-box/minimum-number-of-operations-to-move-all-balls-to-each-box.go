func minOperations(boxes string) []int {
    res := make([]int, len(boxes))
    for i:=0; i < len(boxes); i++ {
        if boxes[i] - '0' == 1 {
            for idx, _ := range res {
                if i >= idx {
                res[idx] += i - idx 
                } else {
                    res [idx] += idx - i
                }
            }
        }
    }
    return res
}