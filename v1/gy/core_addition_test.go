package gy

import (
	"math/rand"
	"testing"
)

func TestAddV8sV8sV8s(t *testing.T) {

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
	sum := make([]Yep8s, vLength)

	for i := 0; i < vLength; i++ {
		sum[i] = x[i] + y[i]
	}

	gyInit()

	// 4. now use AddV8sV8sV8s
	gySum, status := AddV8sV8sV8s(x, y)

	checkStatus(status)

	// 5. compare sum and gySum, print only differente value
	for i := 0; i < vLength; i++ {
		if sum[i] != gySum[i] {
			t.Errorf("Error: sum[%d] != gySum[%d], %d != %d\n", i, i, sum[i], gySum[i])
		}
	}

	gyRelease()

}

func BenchmarkNaiveAddV8sV8sV8s(b *testing.B) {
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
	sum := make([]Yep8s, vLength)

	b.ResetTimer()

	for j := 0; j < b.N; j++ {
		for i := 0; i < vLength; i++ {
			sum[i] = x[i] + y[i]
		}
	}
}

func BenchmarkAddV8sV8sV8s(b *testing.B) {
	// 1. generate two random Yep8s slice
	x := make([]Yep8s, vLength)
	y := make([]Yep8s, vLength)

	// 2. random fill
	r := rand.New(rand.NewSource(vLength))

	for i := 0; i < vLength; i++ {
		x[i] = Yep8s(r.Int())
		y[i] = Yep8s(r.Int())
	}

	b.ResetTimer()

	gyInit()
	for j := 0; j < b.N; j++ {
		_, _ = AddV8sV8sV8s(x, y)
	}

	gyRelease()
}

func TestAddV8sV8sV16s(t *testing.T) {

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
	sum := make([]Yep16s, vLength)

	for i := 0; i < vLength; i++ {
		sum[i] = Yep16s(x[i]) + Yep16s(y[i])
	}

	gyInit()

	// 4. now use AddV8sV8sV8s
	gySum, status := AddV8sV8sV16s(x, y)

	checkStatus(status)

	// 5. compare sum and gySum, print only differente value
	for i := 0; i < vLength; i++ {
		if sum[i] != gySum[i] {
			t.Errorf("Error: sum[%d] != gySum[%d], %d != %d\n", i, i, sum[i], gySum[i])
		}
	}

	gyRelease()

}
