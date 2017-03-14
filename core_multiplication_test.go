package gy

import (
	"math/rand"
	"testing"
)

func TestMultiplyV8sV8sV8s(t *testing.T) {
	// 1. generate two random Yep8s slice
	x := make([]Yep8s, vLength)
	y := make([]Yep8s, vLength)

	// 2. random fill
	r := rand.New(rand.NewSource(vLength))

	for i := 0; i < vLength; i++ {
		x[i] = Yep8s(r.Int())
		y[i] = Yep8s(r.Int())
	}

	// 3. perform mul
	mul := make([]Yep8s, vLength)

	for i := 0; i < vLength; i++ {
		mul[i] = x[i] * y[i]
	}

	gyInit()

	gyMul, status := MultiplyV8sV8sV8s(x, y)
	checkStatus(status)

	for i := 0; i < vLength; i++ {
		if mul[i] != gyMul[i] {
			t.Errorf("Error: mul[%d] != gyMul[%d], %d != %d\n", i, i, mul[i], gyMul[i])
		}
	}
	gyRelease()

}

func TestMultiplyV8sV8sV16s(t *testing.T) {
	// 1. generate two random Yep8s slice
	x := make([]Yep8s, vLength)
	y := make([]Yep8s, vLength)

	// 2. random fill
	r := rand.New(rand.NewSource(vLength))

	for i := 0; i < vLength; i++ {
		x[i] = Yep8s(r.Int())
		y[i] = Yep8s(r.Int())
	}

	// 3. perform mul
	mul := make([]Yep16s, vLength)

	for i := 0; i < vLength; i++ {
		mul[i] = Yep16s(x[i]) * Yep16s(y[i])
	}

	gyInit()

	gyMul, status := MultiplyV8sV8sV16s(x, y)
	checkStatus(status)

	for i := 0; i < vLength; i++ {
		if mul[i] != gyMul[i] {
			t.Errorf("Error: mul[%d] != gyMul[%d], %d != %d\n", i, i, mul[i], gyMul[i])
		}
	}
	gyRelease()

}

func BenchmarkNaifMultiplyV8sV8sV8s(b *testing.B) {
	// 1. generate two random Yep8s slice
	x := make([]Yep8s, vLength)
	y := make([]Yep8s, vLength)

	// 2. random fill
	r := rand.New(rand.NewSource(vLength))

	for i := 0; i < vLength; i++ {
		x[i] = Yep8s(r.Int())
		y[i] = Yep8s(r.Int())
	}

	// 3. perform sum
	mul := make([]Yep8s, vLength)

	b.ResetTimer()

	for j := 0; j < b.N; j++ {
		for i := 0; i < vLength; i++ {
			mul[i] = x[i] * y[i]
		}
	}
}

func BenchmarkMultiplyV8sV8sV8s(b *testing.B) {

	// 1. generate two random Yep8s slice
	x := make([]Yep8s, vLength)
	y := make([]Yep8s, vLength)

	// 2. random fill
	r := rand.New(rand.NewSource(vLength))

	for i := 0; i < vLength; i++ {
		x[i] = Yep8s(r.Int())
		y[i] = Yep8s(r.Int())
	}

	// 3. perform multi
	b.ResetTimer()

	gyInit()
	for j := 0; j < b.N; j++ {
		_, _ = MultiplyV8sV8sV8s(x, y)
	}

	gyRelease()
}
