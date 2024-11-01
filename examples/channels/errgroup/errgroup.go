package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
)

// START OMIT
func sum(nums []int, resultChannel chan<- int) {
	total := 0
	for _, v := range nums {
		total += v
	}
	resultChannel <- total
}

func doit() error {
	var g errgroup.Group
	resultChannel := make(chan int, 2)

	numSets := [][]int{{1, 2, 3}, {4, 5, 6}}

	for _, nums := range numSets {
		nums := nums // no longer needed after 1.23
		g.Go(func() error {
			sum(nums, resultChannel)
			return nil
		})
	}
	// First error will cancel all pending go routines
	if err := g.Wait(); err != nil {
		return err
	}
	close(resultChannel)

	totalSum := 0
	for result := range resultChannel {
		totalSum += result
	}
	fmt.Printf("sum=%d\n", totalSum)
	return nil
}

// END OMIT

func main() {
	err := doit()
	if err != nil {
		log.Fatal(err)
	}
}
