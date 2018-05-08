package sort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	numbers := []int{3, 4, 7, 1, 9, 5, 3, 6, 99, 17, 54, 42, 1000, -54}
	expectedSortedNumbers := []int{-54, 1, 3, 3, 4, 5, 6, 7, 9, 17, 42, 54, 99, 1000}

	QuickSort(numbers, 0, len(numbers)-1)
	if !reflect.DeepEqual(numbers, expectedSortedNumbers) {
		t.Errorf("Merge sort did not sort slice correctly. Got: %d, expected: %d", numbers, expectedSortedNumbers)
	}
}
