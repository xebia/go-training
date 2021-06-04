package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
)

// START OMIT
func printUsage() {
	fmt.Fprintf(os.Stderr, "\nUsage:\n")
	fmt.Fprintf(os.Stderr, " %s [flags]\n", path.Base(os.Args[0]))
	flag.PrintDefaults() // HL
	fmt.Fprintf(os.Stderr, "\n")
	os.Exit(1)
}

func main() {
	login := flag.String("login", "", "GitHub login of user") // HL
	once := flag.Bool("once", false, "Perform action once")   // HL
	reps := flag.Int("reps", 10, "Number of reps")            // HL
	flag.Parse()

	if *login == "" {
		printUsage()
	}

	for idx := 0; idx < *reps; idx++ {
		log.Printf("Looking up GitHub user: %s (once:%v)", *login, *once)
	}
}

// END OMIT
