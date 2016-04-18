# GoYeppp!

GoYeppp is a bindings for [Yeppp!](http://www.yeppp.info/) high-performance SIMD-optimized mathematical library.

## Requirements

The Yeppp! library and [pkgconfig](http://pkgconfig.freedesktop.org).

Customize the pkgconfig config file (yeppp.pc) and then copy into the pkgconfig configurations directory (tipically /usr/lib64/pkgconfig or /usr/lib/pkgconfig)

## Install

This package is "go-gettable", just do:

    github.com/vpxyz/goyeppp


## Example

``` go
package main

import (
	"github.com/vpxyz/goyeppp/v1/gy"
	"log"
	"math/rand"
)

func main() {

	arraySize := 1024 * 1024 * 16
	x := make([]gy.Yep64f, arraySize)

	r := rand.New(rand.NewSource(5330))

	for i := 0; i < arraySize; i++ {
		x[i] = gy.Yep64f(r.Float64())
	}

	blockSize := 1024
	blockLength := 0

	_, _ = gy.Init()

	var entropy gy.Yep64f
	for i := 0; i < arraySize; i += blockSize {
		blockLength = arraySize - i
		if blockLength > blockSize {
			blockLength = blockSize
		}

		logP, _ := gy.Log(x[i : i+blockLength])

		dotProduct, _ := gy.DotProductV64fV64fS64f(x[i:i+blockLength], logP)
		entropy -= dotProduct

	}
	_, _ = gy.Release()

	log.Printf("Entropy: %f\n", entropy)
}
```

