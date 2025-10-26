func alternatingSum(nums []int) int {
	sums := make([]int, 2)
	for i, num := range nums {
        sums[i%2] += num
	}
    return sums[0] - sums[1]
}