func minOperations(boxes string) []int {
	res := make([]int, len(boxes))
	var leftBalls int
	var leftCost int
	var rightBalls int
	var rightCost int
	for i := 0; i < len(boxes); i++ {
		res[i] += leftCost
		if boxes[i]-'0' == 1 {
			leftBalls += 1
		}
		leftCost += leftBalls
	}
	for i := len(boxes) - 1; i >= 0; i-- {
		res[i] += rightCost
		if boxes[i]-'0' == 1 {
			rightBalls += 1
		}
		rightCost += rightBalls
	}
	return res
}