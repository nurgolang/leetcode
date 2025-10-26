func alternatingSum(nums []int) int {
	var res int
    multiplier := 1
	for _, num := range nums {
        res += num * multiplier
        multiplier *= -1
	}
    return res
}