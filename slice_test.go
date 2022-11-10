package slices

import (
	"fmt"
	"testing"
)

func TestCountNumber(t *testing.T) {
	var a []int = []int{1, 1, 2, 2, 2, 3}
	fmt.Println(CountNumber(2, a))
}
