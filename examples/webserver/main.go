package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
)

// START OMIT

func main() {
	router := http.NewServeMux()

	uider := func() string {
		return uuid.New().String()
	}
	nower := func() time.Time {
		return time.Now()
	}

	store := newPatientStore(nower)
	webService := NewPatientService(uider, store)
	webService.RegisterEndpoints(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

// END OMIT
