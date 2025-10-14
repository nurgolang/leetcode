	// Guaranteed O(N) by pre-allocating the final 2*N capacity to avoid reallocations.
	func getConcatenation(nums []int) []int {
		res := make([]int, len(nums)*2)
		copy(res, nums)
		copy(res[len(nums):], nums)
		return res
	}