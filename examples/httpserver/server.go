package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io"
	"log"
	"net/http"
)

type Server struct {
	router *http.ServeMux
}

func main() {
	s := &Server{
		router: http.NewServeMux(),
	}

	s.router.HandleFunc("POST /api/patient", s.createPatientHandler())

	addr := "localhost:8080"

	fmt.Printf("Start listening at %s", addr)
	err := http.ListenAndServe(addr, s.router)

	if err != nil {
		log.Fatal("could not start server")
	}
}

type CreatePatientRequest struct {
	Patient
}

type CreatePatientResponse struct {
	UID string `json:"uid,omitempty"`
	Patient
}

type Patient struct {
	FullName    string   `json:"fullName,omitempty"`
	AddressLine string   `json:"addressLine,omitempty"`
	Allergies   []string `json:"allergies,omitempty"`
}

func (s *Server) createPatientHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}
		body, err := io.ReadAll(r.Body)

		if err != nil {
			http.Error(w, "Unable to read request body", http.StatusBadRequest)
			return
		}
		defer func(body io.ReadCloser) {
			err := body.Close()
			if err != nil {
				// log error
			}
		}(r.Body)

		cpr := CreatePatientRequest{}
		err = json.Unmarshal(body, &cpr)
		if err != nil {
			http.Error(w, "Unable to read request body", http.StatusBadRequest)
			return
		}

		patient := cpr.Patient
		patientID, err := s.CreatePatient(patient)

		resp := CreatePatientResponse{
			UID:     patientID,
			Patient: patient,
		}

		b, err := json.Marshal(resp)

		_, err = w.Write(b)

		if err != nil {
			http.Error(w, fmt.Sprintf("could not write response: %s", err), http.StatusInternalServerError)
		}
	}

}

func (s *Server) CreatePatient(p Patient) (string, error) {
	uid := uuid.New().String()
	fmt.Printf("storing patient %+v, with uid %s", p, uid)

	return uid, nil
}
