// abs: calculates the absolute value of an integer
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func findPermutationDifference(s string, t string) int {
	tCharIndices := make(map[rune]int)

	for i, v := range t {
		tCharIndices[v] = i
	}

	var totalDiff int

	for i, letter := range s {
		diff := i - tCharIndices[letter] // Calculate the difference once
		totalDiff += abs(diff)
	}
	return totalDiff
}