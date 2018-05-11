package main

import (
	//"fmt"
	//"sandbox/sort"
	"fmt"
	"sandbox/barren-land"
	"strings"
)

func main() {
	// var message string
	// message = "Hello Go World!"
	// fmt.Println(message)

	// numbers := []int{3, 4, 7, 1, 9, 5, 3, 6, 99, 17, 54, 42, 1000, -54}
	// fmt.Println(numbers)

	// fmt.Println("\nMerge Sort:")
	// sortedNumbers := sort.TopDownMergeSort(numbers)
	// fmt.Println(sortedNumbers)

	// fmt.Println("\nQuick Sort:")
	// sort.QuickSort(numbers, 0, len(numbers)-1)
	// fmt.Println(numbers)

	fertileLandSizes := barrenland.RunBarrenLandAnalysis(nil)
	fmt.Println(strings.Trim(fmt.Sprint(fertileLandSizes), "[]"))
}
