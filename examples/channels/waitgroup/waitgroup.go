package main

import (
	"fmt"
	"sync"
)

// START OMIT
func sum(nums []int, resultChannel chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	total := 0
	for _, v := range nums {
		total += v
	}
	resultChannel <- total
}

func doit() {
	var wg sync.WaitGroup
	resultChannel := make(chan int, 2)

	wg.Add(2)

	go sum([]int{1, 2, 3}, resultChannel, &wg)
	go sum([]int{4, 5, 6}, resultChannel, &wg)

	// Wait for all goroutines to finish
	wg.Wait()
	close(resultChannel)

	totalSum := 0
	for result := range resultChannel {
		totalSum += result
	}
	fmt.Printf("sum=%d\n", totalSum)
}

// END OMIT

func main() {
	doit()
}
