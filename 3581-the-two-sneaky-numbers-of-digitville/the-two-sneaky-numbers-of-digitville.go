func getSneakyNumbers(nums []int) []int {
    counts := make(map[int]int)
    for _, num := range nums {
        counts[num]++ 
    }
    
    res := make([]int, 2)
    var idx int
    for num, count := range counts {
        if count >= 2 {
            res[idx] = num
            idx++
        }
    }   
    return res
}