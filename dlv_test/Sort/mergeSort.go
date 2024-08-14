package main

func merge(arr []int, left, mid, right int) []int {
	sorted := make([]int, len(arr))
	copy(sorted, arr)
	i, j, k := left, mid+1, left
	for i <= mid && j <= right {
		if arr[i] <= arr[j] {
			sorted[k] = arr[i]
			k++
			i++
		} else {
			sorted[k] = arr[j]
			k++
			j++
		}
	}
	if i > mid {
		for l := j; l <= right; l++ {
			sorted[k] = arr[l]
			k++
		}
	} else {
		for l := i; l <= mid; l++ {
			sorted[k] = arr[l]
			k++
		}
	}
	return sorted
}
func MergeSort(arr []int, left, right int) []int {
	var mid int
	if left < right {
		mid = (left + right) / 2
		arr = MergeSort(arr, left, mid)
		arr = MergeSort(arr, mid+1, right)
		arr = merge(arr, left, mid, right)
	}
	return arr
}
