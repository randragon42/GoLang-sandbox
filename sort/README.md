## Sort

### Merge Sort
From the Wikipedia page on (merge sort)[https://en.wikipedia.org/wiki/Merge_sort]

>Conceptually, a merge sort works as follows:
> 1. Divide the unsorted list into n sublists, each containing 1 element (a list of 1 element is considered sorted).
> 2. Repeatedly merge sublists to produce new sorted sublists until there is only 1 sublist remaining. This will be the sorted list.

### Quick Sort
From the Wikipedia page on (quicksort)[https://en.wikipedia.org/wiki/Quicksort]

>Quicksort is a divide and conquer algorithm. Quicksort first divides a large array into two smaller sub-arrays: the low elements and the high elements. Quicksort can then recursively sort the sub-arrays.
>
>The steps are:
>
> 1. Pick an element, called a pivot, from the array.
> 2. Partitioning: reorder the array so that all elements with values less than the pivot come before the pivot, while all elements with values greater than the pivot come after it (equal values can go either way). After this partitioning, the pivot is in its final position. This is called the partition operation.
> 3. Recursively apply the above steps to the sub-array of elements with smaller values and separately to the sub-array of elements with greater values.
>The base case of the recursion is arrays of size zero or one, which are in order by definition, so they never need to be sorted.
>
>The pivot selection and partitioning steps can be done in several different ways; the choice of specific implementation schemes greatly affects the algorithm's performance.

### Testing
Run ```go test``` in  the ```/sort``` directory to run all tests for both sorting algorithms.