package goyeppp

import (
	"log"
)

const (
	vLength = 1024 * 1024 // test vector length
)

func checkStatus(status YepStatus) {
	log.Printf("Status: %s\n", status.Name())
}

func gyInit() {
	status, error := Init()

	if error != nil {
		log.Fatalf("Error %v\n", error)
	}

	log.Printf("Init status: %v\n", status.Name())

}

func gyRelease() {
	status, error := Release()

	if error != nil {
		log.Fatalf("Error: %v\n", error)
	}
	log.Printf("Release status: %v\n", status.Name())
}
