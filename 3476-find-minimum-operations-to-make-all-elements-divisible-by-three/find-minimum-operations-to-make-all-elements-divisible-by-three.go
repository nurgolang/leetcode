func minimumOperations(nums []int) int {
    var totalOperations int
	for _, num := range nums {
		remainder := num % 3
        totalOperations += min(remainder, 3 - remainder)
	}
    return totalOperations
}