func shuffle(nums []int, n int) []int {
	res := make([]int, len(nums))
	for i := range n {
        k := 2*i
		res[k] = nums[i]
		res[k+1] = nums[i+n]
	}
	return res
}