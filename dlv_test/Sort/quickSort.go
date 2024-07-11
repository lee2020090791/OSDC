package main

// divide and conquer
import (
	"fmt"
	"math/rand"
	"time"
)

func quickSort(arr []int) []int {
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
	leftArr = quickSort(leftArr)
	rightArr = quickSort(rightArr)

	leftArr = append(leftArr, midArr...)
	leftArr = append(leftArr, rightArr...)

	return leftArr
}

func RandArray(n int) []int {
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, n)
	for i := 0; i <= n-1; i++ {
		arr[i] = rand.Intn(n)
	}
	return arr
}

func main() {
	arr := RandArray(10)
	fmt.Println("Initial array is:", arr)
	fmt.Println()
	fmt.Println("Sorted array is:", quickSort(arr))
}
