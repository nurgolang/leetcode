func largestLocal(grid [][]int) [][]int {
	n := len(grid)
	m := n - 2
	maxLocal := make([][]int, m)
	for i := 0; i <= m-1; i++ {
		maxLocal[i] = make([]int, m)
	}
	for i := 0; i < m; i++ {
		for j := 0; j <= m-1; j++ {
			currentMax := grid[i][j]
			for k := i; k <= i+2; k++ {
				for l := j; l <= j+2; l++ {
					currentMax = max(currentMax, grid[k][l])
				}
			}
			maxLocal[i][j] = currentMax
		}
	}
	return maxLocal
}