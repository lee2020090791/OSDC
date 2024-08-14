package main

func SelectionSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	for i := 0; i < len(arr)-1; i++ {
		minIdx := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		temp := arr[i]
		arr[i] = arr[minIdx]
		arr[minIdx] = temp
	}
	return arr
}
