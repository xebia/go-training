package main

import "fmt"

// START OMIT
// This function will be running in the background  // HL
func sum(nums []int, resultChannel chan<- int) {
	total := 0
	for _, v := range nums {
		total += v
	}
	resultChannel <- total
}

func doit() {
	// Create result channel // HL
	resultChannel := make(chan int) // construct blocking channel
	defer close(resultChannel)      // knows when all writers are done, so can close

	// Divide the work over multiple go-routines that run in background // HL
	go sum([]int{1, 2, 3}, resultChannel) // 1 + 2 + 3 = 6
	go sum([]int{4, 5, 6}, resultChannel) // 4 + 5 + 6 = 15

	// Wait for all background tasks to have completed 	// HL
	x := <-resultChannel
	y := <-resultChannel

	// Continue with result // HL
	fmt.Printf("sum=%d", x+y) // order undefined
}

// END OMIT

func main() {
	doit()
}
