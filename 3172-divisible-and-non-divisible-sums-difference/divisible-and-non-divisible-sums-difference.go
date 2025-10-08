func differenceOfSums(n int, m int) int {
    var divCount int
    var indivCount int
	for i := 1; i <= n; i++ {
		if i%m == 0 {
            divCount += i
		} else {
            indivCount += i
        }
	}
    return indivCount - divCount
}