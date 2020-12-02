package libadvent

// ChanToSlice reads in all values from a channel and returns a slice with the values
func ChanToSlice(c chan string) []string {
	vals := make([]string, 0)
	for line := range c {
		vals = append(vals, line)
	}
	return vals
}
