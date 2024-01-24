package patient

import (
	"fmt"
	"github.com/MarcGrol/go-training/examples/interfaces/datastore"
)

type Patient struct {
	UID       string
	FullName  string
	Allergies []string
}

type PatientService struct {
	ds datastore.Datastorer
}

func NewService(datastore datastore.Datastorer) *PatientService {
	return &PatientService{
		ds: datastore,
	}
}

func (ps PatientService) Create(patient Patient) error {
	err := ps.ds.Put(patient.UID, patient)
	if err != nil {
		return fmt.Errorf("Technical error creating patient with uid:%s", err)
	}
	return nil
}

func (ps PatientService) MarkAllergicToAntiBiotics(patientUID string) error {
	opaque, exists, err := ps.ds.Get(patientUID)
	if err != nil {
		return fmt.Errorf("Technical error fetching patient: %s", err)
	}
	if !exists {
		return fmt.Errorf("Patient with uid %s does not exist", patientUID)
	}
	patient := opaque.(Patient)
	patient.Allergies = append(patient.Allergies, "antibiotics")
	err = ps.ds.Put(patientUID, patient)
	if err != nil {
		return fmt.Errorf("Technical error updating patient allergies: %s", err)
	}
	return nil
}
