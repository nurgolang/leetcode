func subsetXORSum(nums []int) int {
	N := len(nums)
	maxMask := 1 << N
	var totalSum int

	for mask := 0; mask < maxMask; mask++ {
		currentXOR := 0
		for i := 0; i < N; i++ {
            if mask & (1 << i) > 0 {
                currentXOR ^= nums[i]
            }
		}
        totalSum += currentXOR
	}
    return totalSum
}