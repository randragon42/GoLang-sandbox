package sort

// Run by calling QuickSort(sliceToBeSorted, 0, len(arrayToBeSorted) - 1)
func QuickSort(numbers []int, low int, high int) {
	if low < high {
		pivot := partition(numbers, low, high)
		QuickSort(numbers, low, pivot - 1)
		QuickSort(numbers, pivot + 1, high)
	}
}

func partition(numbers []int, low int, high int) int {
	pivot := numbers[high]
	i := low - 1
	for j := low; j < high; j++ {
		if numbers[j] < pivot {
			i = i + 1
			Swap(numbers, i, j)
		}
	}
	Swap(numbers, i + 1, high)
	return i + 1
}

func Swap(numbers []int, index1 int, index2 int) {
	temp := numbers[index1]
	numbers[index1] = numbers[index2]
	numbers[index2] = temp
}
