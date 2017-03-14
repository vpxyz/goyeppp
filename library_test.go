package goyeppp

import (
	"log"
	"testing"
)

func TestLibrary(t *testing.T) {

	status, error := Init()

	if error != nil {
		log.Fatalf("Error %v\n", error)
	}

	log.Printf("Init status: %v\n", status.Name())

	var cm YepCpuMicroarchitecture
	status, error = cm.CPUMicroarchitecture()

	if error != nil {
		log.Fatalf("Error: %v\n", error)
	}

	log.Printf("Microarch: %v, status: %v\n", cm, status.Name())

	status, error = Release()

	if error != nil {
		log.Fatalf("Error: %v\n", error)
	}

	log.Printf("Release status: %v\n", status.Name())

}
