package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"time"

	"github.com/xebia/go-training/examples/slowapi"
)

func doit(numTasks int) (int, error) {
	g := errgroup.Group{}
	allResults := make([]int, numTasks)

	for i := 0; i < numTasks; i++ {

		i := i // pre 1.23: Capture loop variable. No longer needed after 1.23 🙌

		g.Go(func() error {
			allResults[i] = slowapi.Sum(i, i)
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return 0, err
	}

	return calculateSum(allResults), nil

}

func calculateSum(allResults []int) int {
	sum := 0
	for _, v := range allResults {
		sum += v
	}
	return sum
}

func main() {
	start := time.Now()
	const taskCount = 10000
	sum, err := doit(taskCount)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Got sum %d\n", sum)
	fmt.Printf(time.Now().Sub(start).String())

}
