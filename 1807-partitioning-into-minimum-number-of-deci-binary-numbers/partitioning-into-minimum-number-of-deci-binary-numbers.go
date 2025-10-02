func minPartitions(n string) int {
	var minNumber int // minNumber would be a max digit
	for _, val := range n {
		if int(val) - '0' > minNumber {
			minNumber = int(val) - '0'
		}
	}
	return minNumber
}