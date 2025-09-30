func maxFreqSum(s string) int {
	counts := make(map[rune]int)
	for _, char := range s {
		counts[char]++
	}
	var maxVowelFreq, maxConsonantFreq int
	isVowel := map[rune]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}
	for char, freq := range counts {
		_, isVowelFound := isVowel[char]
		if isVowelFound {
			if freq > maxVowelFreq {
				maxVowelFreq = freq
			}
		} else {
			if freq > maxConsonantFreq {
				maxConsonantFreq = freq
			}
		}
	}
	return maxVowelFreq + maxConsonantFreq
}