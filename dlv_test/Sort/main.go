package main

// divide and conquer
import (
	"fmt"
	"math/rand"
	"sync"

	// "net/http"
	// _ "net/http/pprof"
	"time"
)

func RandArray(n int) []int {
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, n)
	for i := 0; i <= n-1; i++ {
		arr[i] = rand.Intn(n)
	}
	return arr
}

func main() {

	// go func() {
	// 	http.ListenAndServe("0.0.0.0:6060", nil)
	// }()

	// original := RandArray(10)
	// arr := make([]int, len(original))
	// copy(arr, original)
	// fmt.Println("Initial array is:", arr)
	// fmt.Println()
	// fmt.Println("SelectionSorted array is:", SelectionSort(arr))
	// copy(arr, original)
	// fmt.Println("QuickSorted array is:", QuickSort(arr))
	// copy(arr, original)
	// fmt.Println("InsertionSorted array is:", InsertionSort(arr))
	// copy(arr, original)
	// fmt.Println("MergeSorted array is:", MergeSort(arr, 0, len(arr)-1))
	// a := 3
	// b := 0
	// c := a / b
	// fmt.Printf("%d", c)
	var wg sync.WaitGroup
	wg.Add(4)
	original := RandArray(25)
	arr := make([]int, len(original))
	for i := 0; i < 4; i++ {
		go func(id int) {
			defer wg.Done()
			copy(arr, original)
			if i == 0 {
				SelectionSort(arr)
				fmt.Println("SelectionSort:", arr)
			} else if i == 1 {
				InsertionSort(arr)
				fmt.Println("InsertionSort:", arr)
			} else if i == 2 {
				MergeSort(arr, 0, len(arr)-1)
				fmt.Println("MergeSort:", arr)
			} else {
				QuickSort(arr)
				fmt.Println("QuickSort:", arr)
			}
		}(i)
	}
	wg.Wait()
	fmt.Println("All goroutines finished")

}
