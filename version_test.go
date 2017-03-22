package goyeppp

import (
	"testing"
)

func TestVersion(t *testing.T) {
	t.Logf("Yeppp major version: %v\n", MajorVersion())
	t.Logf("Yeppp minor version: %v\n", MinorVersion())
	t.Logf("Yeppp patch version: %v\n", PatchVersion())
	t.Logf("Yeppp build version: %v\n", BuildVersion())
	t.Logf("Yeppp release name: %v\n", ReleaseName())
	t.Logf("Yeppp full name: %v\n", FullName())
}
