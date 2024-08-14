package main

import (
	"fmt"
	"log"
	"math"
	"time"
)

const INF = math.MaxInt32

type Heap struct {
	capacity int
	size     int
	elements []int
}

func CreateHeap(heapsize int) *Heap {
	newHeap := &Heap{
		capacity: heapsize,
		size:     0,
		elements: make([]int, heapsize+1),
	}
	newHeap.elements[0] = INF
	return newHeap
}

func Find(heap *Heap, value int) bool {
	for i := 1; i <= heap.size; i++ {
		if heap.elements[i] == value {
			return true
		}
	}
	return false
}

func isFull(heap *Heap) bool {
	if heap.capacity > heap.size {
		return false
	} else {
		return true
	}
}

func Insert(heap *Heap, value int) {
	if isFull(heap) {
		fmt.Printf("Insert error : heap is full\n")
	} else if Find(heap, value) {
		fmt.Printf("Insert error : %d is already in heap\n", value)
	}
	var i int
	heap.size++
	for i = heap.size; heap.elements[i/2] < value; i /= 2 {
		heap.elements[i] = heap.elements[i/2]
	}

	heap.elements[i] = value
	// fmt.Printf("Insert %d", value)
}

func DeleteMax(heap *Heap) int {
	var i, child int
	var max, last int
	if heap.size == 0 {
		fmt.Printf("Delete Error : heap is empty\n")
		return -INF
	}
	max = heap.elements[1]
	last = heap.elements[heap.size]
	heap.size--
	for i = 1; i*2 <= heap.size; i = child {
		child = i * 2
		if child < heap.size && heap.elements[child+1] > heap.elements[child] {
			child++
		}
		if last < heap.elements[child] {
			heap.elements[i] = heap.elements[child]
		} else {
			break
		}
	}
	heap.elements[i] = last
	return max
}

func PrintHeap(heap *Heap) {
	if heap.size == 0 {
		fmt.Printf("Print error : heap is empty\n")
		return
	}
	for i := 1; i <= heap.size; i++ {
		fmt.Printf("%d ", heap.elements[i])
	}
	fmt.Println()
}

func main() {
	start := time.Now()
	maxHeap := CreateHeap(10000)
	// PrintHeap(maxHeap)
	// DeleteMax(maxHeap)
	// Insert(maxHeap, 1)
	// PrintHeap(maxHeap)
	// Insert(maxHeap, 2)
	// Insert(maxHeap, 3)
	// PrintHeap(maxHeap)
	// DeleteMax(maxHeap)
	for i := 0; i < 9000; i++ {
		Insert(maxHeap, 9000-i)
		if i%100 == 0 {
			DeleteMax(maxHeap)
		}
	}

	PrintHeap(maxHeap)
	fmt.Println()
	elapsed := time.Since(start)
	log.Printf("execution time : %s\n", elapsed)
	// DeleteMax(maxHeap)
	// PrintHeap(maxHeap)
}
