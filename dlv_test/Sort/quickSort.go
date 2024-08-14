package main

// divide and conquer

func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	pivot := arr[0]
	leftArr := make([]int, 0, len(arr))
	rightArr := make([]int, 0, len(arr))
	midArr := make([]int, 0, len(arr))
	for _, i := range arr {
		if i < pivot {
			leftArr = append(leftArr, i)
		} else if i == pivot {
			midArr = append(midArr, i)
		} else {
			rightArr = append(rightArr, i)
		}
	}
	leftArr = QuickSort(leftArr)
	rightArr = QuickSort(rightArr)

	leftArr = append(leftArr, midArr...)
	leftArr = append(leftArr, rightArr...)

	return leftArr
}
