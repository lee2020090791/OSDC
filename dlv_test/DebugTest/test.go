package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Program started...")
	for i := 0; i < 1000; i++ {
		fmt.Printf("Number : %d\n", i)
		time.Sleep(time.Second)
	}
}
