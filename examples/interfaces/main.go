package main

import (
	"log"

	"github.com/xebia/go-training/examples/interfaces/datastore"
	"github.com/xebia/go-training/examples/interfaces/patient"
)

// START OMIT
func main() {
	// Inject Storer into business logic service
	patientService := patient.NewService(datastore.NewSimplisticDatastore())

	p := patient.Patient{UID: "patient-12345", FullName: "Sjoerd Sjoerdsma", Allergies: []string{"pinda"}}

	// Initialize with data
	err := patientService.Create(p) // uses Datastorer.Put // HL
	if err != nil {
		log.Fatalf("Error creating patient: %s", err)
	}

	// Adjust patient
	err = patientService.MarkAllergicToAntiBiotics(p.UID) // uses Datastorer.Get and Put // HL
	if err != nil {
		log.Fatalf("MarkAllergicToAntiBiotics error: %s", err)
	}
}

// END OMIT
