package utils

// ArrayContains takes a slice of strings and looks for an element in it.
// If found it returns its index, otherwise it will return -1 and false.
func ArrayContains(haystack []string, needle string) (int, bool) {
	for i, item := range haystack {
		if item == needle {
			return i, true
		}
	}

	return -1, false
}
