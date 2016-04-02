package gy

/*
#include <yepCore.h>

#cgo pkg-config: yeppp

*/
import "C"

import (
	"unsafe"
)

// AddV8sV8sV8s Adds corresponding elements in two signed 8-bit integer arrays. Produces an array of signed 8-bit integer elements.
func AddV8sV8sV8s(x, y []Yep8s) ([]Yep8s, YepStatus) {
	sum := make([]Yep8s, len(x))
	status := YepStatus(
		C.yepCore_Add_V8sV8s_V8s(
			(*C.Yep8s)(unsafe.Pointer(&x[0])),
			(*C.Yep8s)(unsafe.Pointer(&y[0])),
			(*C.Yep8s)(unsafe.Pointer(&sum[0])),
			C.YepSize(len(sum)),
		),
	)

	return sum, status
}

// AddV8sV8sV16s Adds corresponding elements in two signed 8-bit integer arrays. Produces an array of signed 16-bit integer elements.
func AddV8sV8sV16s(x, y []Yep8s) ([]Yep16s, YepStatus) {
	sum := make([]Yep16s, len(x))
	status := YepStatus(
		C.yepCore_Add_V8sV8s_V16s(
			(*C.Yep8s)(unsafe.Pointer(&x[0])),
			(*C.Yep8s)(unsafe.Pointer(&y[0])),
			(*C.Yep16s)(unsafe.Pointer(&sum[0])),
			C.YepSize(len(sum)),
		),
	)

	return sum, status
}

// AddV8uV8uV16u Adds corresponding elements in two unsigned 8-bit integer arrays. Produces an array of unsigned 16-bit integer elements.
func AddV8uV8uV16u(x, y []Yep8u) ([]Yep16u, YepStatus) {
	sum := make([]Yep16u, len(x))
	status := YepStatus(
		C.yepCore_Add_V8uV8u_V16u(
			(*C.Yep8u)(unsafe.Pointer(&x[0])),
			(*C.Yep8u)(unsafe.Pointer(&y[0])),
			(*C.Yep16u)(unsafe.Pointer(&sum[0])),
			C.YepSize(len(sum)),
		),
	)

	return sum, status
}

// AddV16sV16sV16s Adds corresponding elements in two signed 16-bit integer arrays. Produces an array of signed 16-bit integer elements.
func AddV16sV16sV16s(x, y []Yep16s) ([]Yep16s, YepStatus) {
	sum := make([]Yep16s, len(x))
	status := YepStatus(
		C.yepCore_Add_V16sV16s_V16s(
			(*C.Yep16s)(unsafe.Pointer(&x[0])),
			(*C.Yep16s)(unsafe.Pointer(&y[0])),
			(*C.Yep16s)(unsafe.Pointer(&sum[0])),
			C.YepSize(len(sum)),
		),
	)

	return sum, status
}

// AddV16sV16sV32s Adds corresponding elements in two signed 16-bit integer arrays. Produces an array of signed 32-bit integer elements.
func AddV16sV16sV32s(x, y []Yep16s) ([]Yep32s, YepStatus) {
	sum := make([]Yep32s, len(x))
	status := YepStatus(
		C.yepCore_Add_V16sV16s_V32s(
			(*C.Yep16s)(unsafe.Pointer(&x[0])),
			(*C.Yep16s)(unsafe.Pointer(&y[0])),
			(*C.Yep32s)(unsafe.Pointer(&sum[0])),
			C.YepSize(len(sum)),
		),
	)

	return sum, status
}

// AddV16uV16uV32u Adds corresponding elements in two unsigned 16-bit integer arrays. Produces an array of unsigned 32-bit integer elements.
func AddV16uV16uV32u(x, y []Yep16u) ([]Yep32u, YepStatus) {
	sum := make([]Yep32u, len(x))
	status := YepStatus(
		C.yepCore_Add_V16uV16u_V32u(
			(*C.Yep16u)(unsafe.Pointer(&x[0])),
			(*C.Yep16u)(unsafe.Pointer(&y[0])),
			(*C.Yep32u)(unsafe.Pointer(&sum[0])),
			C.YepSize(len(sum)),
		),
	)

	return sum, status
}

// AddV32sV32sV32s Adds corresponding elements in two signed 32-bit integer arrays. Produces an array of signed 32-bit integer elements.
func AddV32sV32sV32s(x, y []Yep32s) ([]Yep32s, YepStatus) {
	sum := make([]Yep32s, len(x))
	status := YepStatus(
		C.yepCore_Add_V32sV32s_V32s(
			(*C.Yep32s)(unsafe.Pointer(&x[0])),
			(*C.Yep32s)(unsafe.Pointer(&y[0])),
			(*C.Yep32s)(unsafe.Pointer(&sum[0])),
			C.YepSize(len(sum)),
		),
	)

	return sum, status
}

// AddV32sV32sV64s Adds corresponding elements in two signed 32-bit integer arrays. Produces an array of signed 64-bit integer elements.
func AddV32sV32sV64s(x, y []Yep32s) ([]Yep64s, YepStatus) {
	sum := make([]Yep64s, len(x))
	status := YepStatus(
		C.yepCore_Add_V32sV32s_V64s(
			(*C.Yep32s)(unsafe.Pointer(&x[0])),
			(*C.Yep32s)(unsafe.Pointer(&y[0])),
			(*C.Yep64s)(unsafe.Pointer(&sum[0])),
			C.YepSize(len(sum)),
		),
	)

	return sum, status
}

// AddV32uV32uV64u Adds corresponding elements in two unsigned 32-bit integer arrays. Produces an array of unsigned 64-bit integer elements.
func AddV32uV32uV64u(x, y []Yep32u) ([]Yep64u, YepStatus) {
	sum := make([]Yep64u, len(x))
	status := YepStatus(
		C.yepCore_Add_V32uV32u_V64u(
			(*C.Yep32u)(unsafe.Pointer(&x[0])),
			(*C.Yep32u)(unsafe.Pointer(&y[0])),
			(*C.Yep64u)(unsafe.Pointer(&sum[0])),
			C.YepSize(len(sum)),
		),
	)

	return sum, status
}

// AddV64sV64sV64s Adds corresponding elements in two signed 64-bit integer arrays. Produces an array of signed 64-bit integer elements.
func AddV64sV64sV64s(x, y []Yep64s) ([]Yep64s, YepStatus) {
	sum := make([]Yep64s, len(x))
	status := YepStatus(
		C.yepCore_Add_V64sV64s_V64s(
			(*C.Yep64s)(unsafe.Pointer(&x[0])),
			(*C.Yep64s)(unsafe.Pointer(&y[0])),
			(*C.Yep64s)(unsafe.Pointer(&sum[0])),
			C.YepSize(len(sum)),
		),
	)

	return sum, status
}

// AddV32fV32fV32f Adds corresponding elements in two single precision (32-bit) floating-point arrays. Produces an array of single precision (32-bit) floating-point elements.
func AddV32fV32fV32f(x, y []Yep32f) ([]Yep32f, YepStatus) {
	sum := make([]Yep32f, len(x))
	status := YepStatus(
		C.yepCore_Add_V32fV32f_V32f(
			(*C.Yep32f)(unsafe.Pointer(&x[0])),
			(*C.Yep32f)(unsafe.Pointer(&y[0])),
			(*C.Yep32f)(unsafe.Pointer(&sum[0])),
			C.YepSize(len(sum)),
		),
	)

	return sum, status
}

// AddV64fV64fV64f Adds corresponding elements in two double precision (64-bit) floating-point arrays. Produces an array of double precision (64-bit) floating-point elements.
func AddV64fV64fV64f(x, y []Yep64f) ([]Yep64f, YepStatus) {
	sum := make([]Yep64f, len(x))
	status := YepStatus(
		C.yepCore_Add_V64fV64f_V64f(
			(*C.Yep64f)(unsafe.Pointer(&x[0])),
			(*C.Yep64f)(unsafe.Pointer(&y[0])),
			(*C.Yep64f)(unsafe.Pointer(&sum[0])),
			C.YepSize(len(sum)),
		),
	)

	return sum, status
}

// AddV8sS8sV8s Adds a constant to signed 8-bit integer array elements. Produces an array of signed 8-bit integer elements.
// This version of Yeppp! does not include optimized implementations for this functio.
func AddV8sS8sV8s(x []Yep8s, y Yep8s) ([]Yep8s, YepStatus) {
	sum := make([]Yep8s, len(x))
	status := YepStatus(
		C.yepCore_Add_V8sS8s_V8s(
			(*C.Yep8s)(unsafe.Pointer(&x[0])),
			(C.Yep8s)(y),
			(*C.Yep8s)(unsafe.Pointer(&sum[0])),
			C.YepSize(len(sum)),
		),
	)
	return sum, status
}

// AddV8sS8sV16s Adds a constant to signed 8-bit integer array elements. Produces an array of signed 16-bit integer elements.
func AddV8sS8sV16s(x []Yep8s, y Yep8s) ([]Yep16s, YepStatus) {
	sum := make([]Yep16s, len(x))
	status := YepStatus(
		C.yepCore_Add_V8sS8s_V16s(
			(*C.Yep8s)(unsafe.Pointer(&x[0])),
			(C.Yep8s)(y),
			(*C.Yep16s)(unsafe.Pointer(&sum[0])),
			C.YepSize(len(sum)),
		),
	)
	return sum, status
}

// AddV8uS8uV16u Adds a constant to unsigned 8-bit integer array elements. Produces an array of unsigned 16-bit integer elements.
func AddV8uS8uV16u(x []Yep8u, y Yep8u) ([]Yep16u, YepStatus) {
	sum := make([]Yep16u, len(x))
	status := YepStatus(
		C.yepCore_Add_V8uS8u_V16u(
			(*C.Yep8u)(unsafe.Pointer(&x[0])),
			(C.Yep8u)(y),
			(*C.Yep16u)(unsafe.Pointer(&sum[0])),
			C.YepSize(len(sum)),
		),
	)
	return sum, status
}

// AddV16sS16sV16s Adds a constant to signed 16-bit integer array elements. Produces an array of signed 16-bit integer elements.
// Warning This version of Yeppp! does not include optimized implementations for this function
func AddV16sS16sV16s(x []Yep16s, y Yep16s) ([]Yep16s, YepStatus) {
	sum := make([]Yep16s, len(x))
	status := YepStatus(
		C.yepCore_Add_V16sS16s_V16s(
			(*C.Yep16s)(unsafe.Pointer(&x[0])),
			(C.Yep16s)(y),
			(*C.Yep16s)(unsafe.Pointer(&sum[0])),
			C.YepSize(len(sum)),
		),
	)
	return sum, status
}

// AddV16sS16sV32s Adds a constant to signed 16-bit integer array elements. Produces an array of signed 32-bit integer elements.
func AddV16sS16sV32s(x []Yep16s, y Yep16s) ([]Yep32s, YepStatus) {
	sum := make([]Yep32s, len(x))
	status := YepStatus(
		C.yepCore_Add_V16sS16s_V32s(
			(*C.Yep16s)(unsafe.Pointer(&x[0])),
			(C.Yep16s)(y),
			(*C.Yep32s)(unsafe.Pointer(&sum[0])),
			C.YepSize(len(sum)),
		),
	)
	return sum, status
}

// AddV16uS16uV32u Adds a constant to unsigned 16-bit integer array elements. Produces an array of unsigned 32-bit integer elements.
func AddV16uS16uV32u(x []Yep16u, y Yep16u) ([]Yep32u, YepStatus) {
	sum := make([]Yep32u, len(x))
	status := YepStatus(
		C.yepCore_Add_V16uS16u_V32u(
			(*C.Yep16u)(unsafe.Pointer(&x[0])),
			(C.Yep16u)(y),
			(*C.Yep32u)(unsafe.Pointer(&sum[0])),
			C.YepSize(len(sum)),
		),
	)
	return sum, status
}

// AddV32sS32sV32s Adds a constant to signed 32-bit integer array elements. Produces an array of signed 32-bit integer elements.
// Warning This version of Yeppp! does not include optimized implementations for this function
func AddV32sS32sV32s(x []Yep32s, y Yep32s) ([]Yep32s, YepStatus) {
	sum := make([]Yep32s, len(x))
	status := YepStatus(
		C.yepCore_Add_V32sS32s_V32s(
			(*C.Yep32s)(unsafe.Pointer(&x[0])),
			(C.Yep32s)(y),
			(*C.Yep32s)(unsafe.Pointer(&sum[0])),
			C.YepSize(len(sum)),
		),
	)
	return sum, status
}

// AddV32uS32uV64u Adds a constant to unsigned 32-bit integer array elements. Produces an array of unsigned 64-bit integer elements.
func AddV32uS32uV64u(x []Yep32u, y Yep32u) ([]Yep64u, YepStatus) {
	sum := make([]Yep64u, len(x))
	status := YepStatus(
		C.yepCore_Add_V32uS32u_V64u(
			(*C.Yep32u)(unsafe.Pointer(&x[0])),
			(C.Yep32u)(y),
			(*C.Yep64u)(unsafe.Pointer(&sum[0])),
			C.YepSize(len(sum)),
		),
	)

	return sum, status
}

// AddV32sS32sV64s Adds a constant to signed 32-bit integer array elements. Produces an array of signed 64-bit integer elements.
func AddV32sS32sV64s(x []Yep32s, y Yep32s) ([]Yep64s, YepStatus) {
	sum := make([]Yep64s, len(x))
	status := YepStatus(
		C.yepCore_Add_V32sS32s_V64s(
			(*C.Yep32s)(unsafe.Pointer(&x[0])),
			(C.Yep32s)(y),
			(*C.Yep64s)(unsafe.Pointer(&sum[0])),
			C.YepSize(len(sum)),
		),
	)
	return sum, status
}

// AddV64sS64sV64s Adds a constant to signed 64-bit integer array elements. Produces an array of signed 64-bit integer elements.
// Warning This version of Yeppp! does not include optimized implementations for this function
func AddV64sS64sV64s(x []Yep64s, y Yep64s) ([]Yep64s, YepStatus) {
	sum := make([]Yep64s, len(x))
	status := YepStatus(
		C.yepCore_Add_V64sS64s_V64s(
			(*C.Yep64s)(unsafe.Pointer(&x[0])),
			(C.Yep64s)(y),
			(*C.Yep64s)(unsafe.Pointer(&sum[0])),
			C.YepSize(len(sum)),
		),
	)
	return sum, status
}

// AddV32fS32fV32f Adds a constant to single precision (32-bit) floating-point array elements. Produces an array of single precision (32-bit) floating-point elements.
// Warning This version of Yeppp! does not include optimized implementations for this function
func AddV32fS32fV32f(x []Yep32f, y Yep32f) ([]Yep32f, YepStatus) {
	sum := make([]Yep32f, len(x))
	status := YepStatus(
		C.yepCore_Add_V32fS32f_V32f(
			(*C.Yep32f)(unsafe.Pointer(&x[0])),
			(C.Yep32f)(y),
			(*C.Yep32f)(unsafe.Pointer(&sum[0])),
			C.YepSize(len(sum)),
		),
	)
	return sum, status
}

// AddV64fS64fV64f Adds a constant to double precision (64-bit) floating-point array elements. Produces an array of double precision (64-bit) floating-point elements.
// Warning This version of Yeppp! does not include optimized implementations for this function
func AddV64fS64fV64f(x []Yep64f, y Yep64f) ([]Yep64f, YepStatus) {
	sum := make([]Yep64f, len(x))
	status := YepStatus(
		C.yepCore_Add_V64fS64f_V64f(
			(*C.Yep64f)(unsafe.Pointer(&x[0])),
			(C.Yep64f)(y),
			(*C.Yep64f)(unsafe.Pointer(&sum[0])),
			C.YepSize(len(sum)),
		),
	)
	return sum, status
}

// AddIV8sV8sIV8s Adds corresponding elements in two signed 8-bit integer arrays and writes the result to the first array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func AddIV8sV8sIV8s(x, y []Yep8s) YepStatus {
	status := YepStatus(
		C.yepCore_Add_IV8sV8s_IV8s(
			(*C.Yep8s)(unsafe.Pointer(&x[0])),
			(*C.Yep8s)(unsafe.Pointer(&y[0])),
			C.YepSize(len(x)),
		),
	)
	return status
}

// AddIV16sV16sIV16s Adds corresponding elements in two signed 16-bit integer arrays and writes the result to the first array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func AddIV16sV16sIV16s(x, y []Yep16s) YepStatus {
	status := YepStatus(
		C.yepCore_Add_IV16sV16s_IV16s(
			(*C.Yep16s)(unsafe.Pointer(&x[0])),
			(*C.Yep16s)(unsafe.Pointer(&y[0])),
			C.YepSize(len(x)),
		),
	)
	return status
}

// AddIV32sV32sIV32s Adds corresponding elements in two signed 32-bit integer arrays and writes the result to the first array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func AddIV32sV32sIV32s(x, y []Yep32s) YepStatus {
	status := YepStatus(
		C.yepCore_Add_IV32sV32s_IV32s(
			(*C.Yep32s)(unsafe.Pointer(&x[0])),
			(*C.Yep32s)(unsafe.Pointer(&y[0])),
			C.YepSize(len(x)),
		),
	)
	return status
}

// AddIV64sV64sIV64s Adds corresponding elements in two signed 64-bit integer arrays and writes the result to the first array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func AddIV64sV64sIV64s(x, y []Yep64s) YepStatus {
	status := YepStatus(
		C.yepCore_Add_IV64sV64s_IV64s(
			(*C.Yep64s)(unsafe.Pointer(&x[0])),
			(*C.Yep64s)(unsafe.Pointer(&y[0])),
			C.YepSize(len(x)),
		),
	)
	return status
}

// AddIV32fV32fIV32f Adds corresponding elements in two single precision (32-bit) floating-point arrays and writes the result to the first array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func AddIV32fV32fIV32f(x, y []Yep32f) YepStatus {
	status := YepStatus(
		C.yepCore_Add_IV32fV32f_IV32f(
			(*C.Yep32f)(unsafe.Pointer(&x[0])),
			(*C.Yep32f)(unsafe.Pointer(&y[0])),
			C.YepSize(len(x)),
		),
	)
	return status
}

// AddIV64fV64fIV64f Adds corresponding elements in two double precision (64-bit) floating-point arrays and writes the result to the first array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func AddIV64fV64fIV64f(x, y []Yep64f) YepStatus {
	status := YepStatus(
		C.yepCore_Add_IV64fV64f_IV64f(
			(*C.Yep64f)(unsafe.Pointer(&x[0])),
			(*C.Yep64f)(unsafe.Pointer(&y[0])),
			C.YepSize(len(x)),
		),
	)
	return status
}

// AddIV8sS8sIV8s Adds a constant to signed 8-bit integer array elements and writes the result to the same array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func AddIV8sS8sIV8s(x []Yep8s, y Yep8s) YepStatus {
	status := YepStatus(
		C.yepCore_Add_IV8sS8s_IV8s(
			(*C.Yep8s)(unsafe.Pointer(&x[0])),
			(C.Yep8s)(y),
			C.YepSize(len(x)),
		),
	)
	return status
}

// AddIV16sS16sIV16s Adds a constant to signed 16-bit integer array elements and writes the result to the same array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func AddIV16sS16sIV16s(x []Yep16s, y Yep16s) YepStatus {
	status := YepStatus(
		C.yepCore_Add_IV16sS16s_IV16s(
			(*C.Yep16s)(unsafe.Pointer(&x[0])),
			(C.Yep16s)(y),
			C.YepSize(len(x)),
		),
	)
	return status
}

// AddIV32sS32sIV32s Adds a constant to signed 32-bit integer array elements and writes the result to the same array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func AddIV32sS32sIV32s(x []Yep32s, y Yep32s) YepStatus {
	status := YepStatus(
		C.yepCore_Add_IV32sS32s_IV32s(
			(*C.Yep32s)(unsafe.Pointer(&x[0])),
			(C.Yep32s)(y),
			C.YepSize(len(x)),
		),
	)
	return status
}

// AddIV64sS64sIV64s Adds a constant to signed 64-bit integer array elements and writes the result to the same array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func AddIV64sS64sIV64s(x []Yep64s, y Yep64s) YepStatus {
	status := YepStatus(
		C.yepCore_Add_IV64sS64s_IV64s(
			(*C.Yep64s)(unsafe.Pointer(&x[0])),
			(C.Yep64s)(y),
			C.YepSize(len(x)),
		),
	)
	return status
}

// AddIV32fS32fIV32f Adds a constant to single precision (32-bit) floating-point array elements and writes the result to the same array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func AddIV32fS32fIV32f(x []Yep32f, y Yep32f) YepStatus {
	status := YepStatus(
		C.yepCore_Add_IV32fS32f_IV32f(
			(*C.Yep32f)(unsafe.Pointer(&x[0])),
			(C.Yep32f)(y),
			C.YepSize(len(x)),
		),
	)
	return status
}

// AddIV64fS64fIV64f Adds a constant to double precision (64-bit) floating-point array elements and writes the result to the same array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func AddIV64fS64fIV64f(x []Yep64f, y Yep64f) YepStatus {
	status := YepStatus(
		C.yepCore_Add_IV64fS64f_IV64f(
			(*C.Yep64f)(unsafe.Pointer(&x[0])),
			(C.Yep64f)(y),
			C.YepSize(len(x)),
		),
	)
	return status
}
