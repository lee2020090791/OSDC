package main

import (
	"fmt"
	"sync"
	// "time"
)

var gcount int

func main() {
	var wg sync.WaitGroup
	counter := 0
	goroutines := 100

	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter++
				fmt.Printf("%d ", i)
				gcount++
				// time.Sleep(time.Millisecond)
			}
			//fmt.Printf("Goroutine %d finished\n", id)
		}(i)
	}

	wg.Wait()

	fmt.Println("All goroutines finished")
	fmt.Println("Final counter value:", counter)
	fmt.Println("Final global counter value:", gcount)
}
