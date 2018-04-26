package sort

func TopDownMergeSort(numbers []int) (sorted []int){
	if len(numbers) <= 1 {
		return numbers
	}

	middleIndex := len(numbers)/2
	left := []int{}
	right := []int{}

	for i, number := range numbers {
		if i < middleIndex {
			left = append(left, number)
		} else {
			right = append(right, number)
		}
	}

	left = TopDownMergeSort(left)
	right = TopDownMergeSort(right)

	return merge(left, right)
}

func merge(left []int, right []int) (sorted []int){
	sorted = []int{}

	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
			sorted = append(sorted, left[0])
			left = left[1:]
		} else {
			sorted = append(sorted, right[0])
			right = right[1:]
		}
	}

	if len(left) > 0 {
		sorted = append(sorted, left...)
	}
	if len(right) > 0 {
		sorted = append(sorted, right...)
	}

	return
}