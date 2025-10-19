func transformArray(nums []int) []int {
	var cnt int
	for _, val := range nums {
		if val%2 != 0 {
			cnt++
		}
	}
    res := make([]int, len(nums))

    for i:= len(nums) - cnt; i < len(nums); i++ {
        res[i] = 1
    }
	return res
}