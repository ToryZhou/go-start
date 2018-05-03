package insertion

import (
	"testing"
)

func TestSort(t *testing.T) {
	numbers := []int{1, 5, 4, 3, 2, 2, 3, 4, 5, 2}
	sort(numbers)
	for i := 0; i < len(numbers)-2; i++ {
		if numbers[i] > numbers[i+1] {
			t.Error()
		}
	}
}
