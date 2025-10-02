import "strings"

func defangIPaddr(address string) string {
	var res strings.Builder
	res.Grow(len(address) + 6) // pre-allocate space for efficiency
	ipBytes := []byte(address)
	for _, part := range ipBytes {
		if part == '.' {
			res.WriteString("[.]")
		} else {
			res.WriteByte(part)
		}
	}
	return res.String()
}