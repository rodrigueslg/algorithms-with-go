package module01

// NumInList will return true if num is in the
// list slice, and false otherwise.
func NumInList(list []int, num int) bool {

	if len(list) == 0 {
		return false
	}

	for _, x := range list {
		if x == num {
			return true
		}
	}

	return false
}
