func getConcatenation(nums []int) []int {
    res := append([]int{}, nums...)
    return append(res, nums...)
}