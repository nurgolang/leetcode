func numIdenticalPairs(nums []int) int {
    var cnt int
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
            if nums[i] == nums[j] {
                cnt++
            } 
		}
	}
    return cnt
}