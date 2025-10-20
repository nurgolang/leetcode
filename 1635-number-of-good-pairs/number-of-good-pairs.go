// Gauss formula solution
func numIdenticalPairs(nums []int) int {
	counts := make(map[int]int)
	for _, num := range nums {
       counts[num]++
	}

    var totalPairs int
    for _, freq := range counts{
        totalPairs += (freq * (freq - 1)) / 2
    }
    return totalPairs
}