package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup, results chan<- int) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)

	fmt.Printf("Worker %d done\n", id)
	results <- id
}

func main() {
	var wg sync.WaitGroup
	results := make(chan int, 5)

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg, results)
	}

	wg.Wait()
	close(results)

	for result := range results {
		fmt.Printf("Result from worker %d\n", result)
	}
}
