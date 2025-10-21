func absInt(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func findClosest(x int, y int, z int) int {
	distX := absInt(x - z)
	distY := absInt(y - z)

	switch {
	case distX < distY:
		return 1
	case distX > distY:
		return 2
	default:
		return 0
	}
}