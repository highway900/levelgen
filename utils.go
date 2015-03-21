package levelgen

// Return int if greater than v
// else return v
//
func bi(v int, n int) int {
	if v < n {
		return n
	}
	return v
}