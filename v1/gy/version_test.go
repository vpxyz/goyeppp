package gy

import (
	"log"
	"testing"
)

func TestVersion(t *testing.T) {
	log.Printf("Yeppp major version: %v\n", MajorVersion())
	log.Printf("Yeppp minor version: %v\n", MinorVersion())
	log.Printf("Yeppp patch version: %v\n", PatchVersion())
	log.Printf("Yeppp build version: %v\n", BuildVersion())
	log.Printf("Yeppp release name: %v\n", ReleaseName())
	log.Printf("Yeppp full name: %v\n", FullName())
}
