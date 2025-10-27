import "slices"

func maximumWealth(accounts [][]int) int {
	allWealths := make([]int, len(accounts))
	for i, customerAccounts := range accounts {
		wealth := 0
		for _, customerAccount := range customerAccounts {
			wealth += customerAccount
		}
		allWealths[i] = wealth
	}
	return slices.Max(allWealths)
}