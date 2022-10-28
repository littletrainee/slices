package slices

// check slice is contains element.
func ContainsElement(element any, targetslice []any) bool {
	for _, v := range targetslice {
		if v == element {
			return true
		}
	}
	return false
}

// find target element's index of element.
func FindIndexOfElement(element any, targetslice ...any) int {
	for i, v := range targetslice {
		if v == element {
			return i
		}
	}
	return -1
}
