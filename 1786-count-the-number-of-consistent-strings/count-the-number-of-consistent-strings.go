func countConsistentStrings(allowed string, words []string) int {
	// create an empty map to store allowed characters
	allowedChars := make(map[rune]struct{})

	// iterate over each character (rune) in the allowed string
	for _, char := range allowed {
		allowedChars[char] = struct{}{}
	}

	var cnt int

wordLoop:
	for _, word := range words {
		for _, letter := range word {
			_, found := allowedChars[letter]
			if !found {
				continue wordLoop
			}
		}
        cnt++
	}
    return cnt
}