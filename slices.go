package slices

type Type interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}

// check slice is contains element.
func ContainsElement[T Type](element T, targetslice []T) bool {
	for _, v := range targetslice {
		if v == element {
			return true
		}
	}
	return false
}

// find target element's index of element.
func FindIndexOfElement[T Type](element T, targetslice []T) int {
	for i, v := range targetslice {
		if v == element {
			return i
		}
	}
	return -1
}
