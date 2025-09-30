func minOperations(boxes string) []int {
	res := make([]int, len(boxes))
	for i := 0; i < len(boxes); i++ {
		if boxes[i]-'0' == 1 {
			for idx, _ := range res {
				diff := i - idx
				if diff < 0 {
					diff = -diff
				}
				res[idx] += diff
			}
		}
	}
	return res
}