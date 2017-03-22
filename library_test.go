package goyeppp

import (
	"testing"
)

func TestLibrary(t *testing.T) {

	status, error := Init()

	if error != nil {
		t.Fatalf("Error %v\n", error)
	}

	t.Logf("Init status: %v\n", status.Name())

	var cm YepCpuMicroarchitecture
	status, error = cm.CPUMicroarchitecture()

	if error != nil {
		t.Fatalf("Error: %v\n", error)
	}

	t.Logf("Microarch: %v, status: %v\n", cm, status.Name())

	status, error = Release()

	if error != nil {
		t.Fatalf("Error: %v\n", error)
	}

	t.Logf("Release status: %v\n", status.Name())

}
