package goyeppp

/*
#include <yepCore.h>

#cgo pkg-config: yeppp

*/
import "C"

import (
	"unsafe"
)

// SubtractV8sV8sV8s Subtracts corresponding elements in two signed 8-bit integer arrays. Produces an array of signed 8-bit integer elements.
func SubtractV8sV8sV8s(x, y []Yep8s) ([]Yep8s, YepStatus) {
	diff := make([]Yep8s, len(x))
	status := YepStatus(
		C.yepCore_Subtract_V8sV8s_V8s(
			(*C.Yep8s)(unsafe.Pointer(&x[0])),
			(*C.Yep8s)(unsafe.Pointer(&y[0])),
			(*C.Yep8s)(unsafe.Pointer(&diff[0])),
			C.YepSize(len(diff)),
		),
	)

	return diff, status
}

// SubtractV8sV8sV16s Subtracts corresponding elements in two signed 8-bit integer arrays. Produces an array of signed 16-bit integer elements.
func SubtractV8sV8sV16s(x, y []Yep8s) ([]Yep16s, YepStatus) {
	diff := make([]Yep16s, len(x))
	status := YepStatus(
		C.yepCore_Subtract_V8sV8s_V16s(
			(*C.Yep8s)(unsafe.Pointer(&x[0])),
			(*C.Yep8s)(unsafe.Pointer(&y[0])),
			(*C.Yep16s)(unsafe.Pointer(&diff[0])),
			C.YepSize(len(diff)),
		),
	)

	return diff, status
}

// SubtractV8uV8uV16u Subtracts corresponding elements in two unsigned 8-bit integer arrays. Produces an array of unsigned 16-bit integer elements.
func SubtractV8uV8uV16u(x, y []Yep8u) ([]Yep16u, YepStatus) {
	diff := make([]Yep16u, len(x))
	status := YepStatus(
		C.yepCore_Subtract_V8uV8u_V16u(
			(*C.Yep8u)(unsafe.Pointer(&x[0])),
			(*C.Yep8u)(unsafe.Pointer(&y[0])),
			(*C.Yep16u)(unsafe.Pointer(&diff[0])),
			C.YepSize(len(diff)),
		),
	)

	return diff, status
}

// SubtractV16sV16sV16s Subtracts corresponding elements in two signed 16-bit integer arrays. Produces an array of signed 16-bit integer elements.
func SubtractV16sV16sV16s(x, y []Yep16s) ([]Yep16s, YepStatus) {
	diff := make([]Yep16s, len(x))
	status := YepStatus(
		C.yepCore_Subtract_V16sV16s_V16s(
			(*C.Yep16s)(unsafe.Pointer(&x[0])),
			(*C.Yep16s)(unsafe.Pointer(&y[0])),
			(*C.Yep16s)(unsafe.Pointer(&diff[0])),
			C.YepSize(len(diff)),
		),
	)

	return diff, status
}

// SubtractV16sV16sV32s Subtracts corresponding elements in two signed 16-bit integer arrays. Produces an array of signed 32-bit integer elements.
func SubtractV16sV16sV32s(x, y []Yep16s) ([]Yep32s, YepStatus) {
	diff := make([]Yep32s, len(x))
	status := YepStatus(
		C.yepCore_Subtract_V16sV16s_V32s(
			(*C.Yep16s)(unsafe.Pointer(&x[0])),
			(*C.Yep16s)(unsafe.Pointer(&y[0])),
			(*C.Yep32s)(unsafe.Pointer(&diff[0])),
			C.YepSize(len(diff)),
		),
	)

	return diff, status
}

// SubtractV16uV16uV32u Subtracts corresponding elements in two unsigned 16-bit integer arrays. Produces an array of unsigned 32-bit integer elements.
func SubtractV16uV16uV32u(x, y []Yep16u) ([]Yep32u, YepStatus) {
	diff := make([]Yep32u, len(x))
	status := YepStatus(
		C.yepCore_Subtract_V16uV16u_V32u(
			(*C.Yep16u)(unsafe.Pointer(&x[0])),
			(*C.Yep16u)(unsafe.Pointer(&y[0])),
			(*C.Yep32u)(unsafe.Pointer(&diff[0])),
			C.YepSize(len(diff)),
		),
	)

	return diff, status
}

// SubtractV32sV32sV32s Subtracts corresponding elements in two signed 32-bit integer arrays. Produces an array of signed 32-bit integer elements.
func SubtractV32sV32sV32s(x, y []Yep32s) ([]Yep32s, YepStatus) {
	diff := make([]Yep32s, len(x))
	status := YepStatus(
		C.yepCore_Subtract_V32sV32s_V32s(
			(*C.Yep32s)(unsafe.Pointer(&x[0])),
			(*C.Yep32s)(unsafe.Pointer(&y[0])),
			(*C.Yep32s)(unsafe.Pointer(&diff[0])),
			C.YepSize(len(diff)),
		),
	)

	return diff, status
}

// SubtractV32sV32sV64s Subtracts corresponding elements in two signed 32-bit integer arrays. Produces an array of signed 64-bit integer elements.
func SubtractV32sV32sV64s(x, y []Yep32s) ([]Yep64s, YepStatus) {
	diff := make([]Yep64s, len(x))
	status := YepStatus(
		C.yepCore_Subtract_V32sV32s_V64s(
			(*C.Yep32s)(unsafe.Pointer(&x[0])),
			(*C.Yep32s)(unsafe.Pointer(&y[0])),
			(*C.Yep64s)(unsafe.Pointer(&diff[0])),
			C.YepSize(len(diff)),
		),
	)

	return diff, status
}

// SubtractV32uV32uV64u Subtracts corresponding elements in two unsigned 32-bit integer arrays. Produces an array of unsigned 64-bit integer elements.
func SubtractV32uV32uV64u(x, y []Yep32u) ([]Yep64u, YepStatus) {
	diff := make([]Yep64u, len(x))
	status := YepStatus(
		C.yepCore_Subtract_V32uV32u_V64u(
			(*C.Yep32u)(unsafe.Pointer(&x[0])),
			(*C.Yep32u)(unsafe.Pointer(&y[0])),
			(*C.Yep64u)(unsafe.Pointer(&diff[0])),
			C.YepSize(len(diff)),
		),
	)

	return diff, status
}

// SubtractV64sV64sV64s Subtracts corresponding elements in two signed 64-bit integer arrays. Produces an array of signed 64-bit integer elements.
func SubtractV64sV64sV64s(x, y []Yep64s) ([]Yep64s, YepStatus) {
	diff := make([]Yep64s, len(x))
	status := YepStatus(
		C.yepCore_Subtract_V64sV64s_V64s(
			(*C.Yep64s)(unsafe.Pointer(&x[0])),
			(*C.Yep64s)(unsafe.Pointer(&y[0])),
			(*C.Yep64s)(unsafe.Pointer(&diff[0])),
			C.YepSize(len(diff)),
		),
	)

	return diff, status
}

// SubtractV32fV32fV32f Subtracts corresponding elements in two single precision (32-bit) floating-point arrays. Produces an array of single precision (32-bit) floating-point elements.
func SubtractV32fV32fV32f(x, y []Yep32f) ([]Yep32f, YepStatus) {
	diff := make([]Yep32f, len(x))
	status := YepStatus(
		C.yepCore_Subtract_V32fV32f_V32f(
			(*C.Yep32f)(unsafe.Pointer(&x[0])),
			(*C.Yep32f)(unsafe.Pointer(&y[0])),
			(*C.Yep32f)(unsafe.Pointer(&diff[0])),
			C.YepSize(len(diff)),
		),
	)

	return diff, status
}

// SubtractV64fV64fV64f Subtracts corresponding elements in two double precision (64-bit) floating-point arrays. Produces an array of double precision (64-bit) floating-point elements.
func SubtractV64fV64fV64f(x, y []Yep64f) ([]Yep64f, YepStatus) {
	diff := make([]Yep64f, len(x))
	status := YepStatus(
		C.yepCore_Subtract_V64fV64f_V64f(
			(*C.Yep64f)(unsafe.Pointer(&x[0])),
			(*C.Yep64f)(unsafe.Pointer(&y[0])),
			(*C.Yep64f)(unsafe.Pointer(&diff[0])),
			C.YepSize(len(diff)),
		),
	)

	return diff, status
}

// SubtractV8sS8sV8s Subtracts a constant from signed 8-bit integer array elements. Produces an array of signed 8-bit integer elements.
// This version of Yeppp! does not include optimized implementations for this functio.
func SubtractV8sS8sV8s(x []Yep8s, y Yep8s) ([]Yep8s, YepStatus) {
	diff := make([]Yep8s, len(x))
	status := YepStatus(
		C.yepCore_Subtract_V8sS8s_V8s(
			(*C.Yep8s)(unsafe.Pointer(&x[0])),
			(C.Yep8s)(y),
			(*C.Yep8s)(unsafe.Pointer(&diff[0])),
			C.YepSize(len(diff)),
		),
	)
	return diff, status
}

// SubtractV8sS8sV16s Subtracts a constant from signed 8-bit integer array elements. Produces an array of signed 16-bit integer elements.
// This version of Yeppp! does not include optimized implementations for this function
func SubtractV8sS8sV16s(x []Yep8s, y Yep8s) ([]Yep16s, YepStatus) {
	diff := make([]Yep16s, len(x))
	status := YepStatus(
		C.yepCore_Subtract_V8sS8s_V16s(
			(*C.Yep8s)(unsafe.Pointer(&x[0])),
			(C.Yep8s)(y),
			(*C.Yep16s)(unsafe.Pointer(&diff[0])),
			C.YepSize(len(diff)),
		),
	)
	return diff, status
}

// SubtractV8uS8uV16u Subtracts a constant from unsigned 8-bit integer array elements. Produces an array of unsigned 16-bit integer elements.
// Warning: This version of Yeppp! does not include optimized implementations for this function
func SubtractV8uS8uV16u(x []Yep8u, y Yep8u) ([]Yep16u, YepStatus) {
	diff := make([]Yep16u, len(x))
	status := YepStatus(
		C.yepCore_Subtract_V8uS8u_V16u(
			(*C.Yep8u)(unsafe.Pointer(&x[0])),
			(C.Yep8u)(y),
			(*C.Yep16u)(unsafe.Pointer(&diff[0])),
			C.YepSize(len(diff)),
		),
	)
	return diff, status
}

// SubtractV16sS16sV16s Subtracts a constant to signed 16-bit integer array elements. Produces an array of signed 16-bit integer elements.
// Warning: This version of Yeppp! does not include optimized implementations for this function
func SubtractV16sS16sV16s(x []Yep16s, y Yep16s) ([]Yep16s, YepStatus) {
	diff := make([]Yep16s, len(x))
	status := YepStatus(
		C.yepCore_Subtract_V16sS16s_V16s(
			(*C.Yep16s)(unsafe.Pointer(&x[0])),
			(C.Yep16s)(y),
			(*C.Yep16s)(unsafe.Pointer(&diff[0])),
			C.YepSize(len(diff)),
		),
	)
	return diff, status
}

// SubtractV16sS16sV32s Subtracts a constant from signed 16-bit integer array elements. Produces an array of signed 32-bit integer elements.
// Warning: This version of Yeppp! does not include optimized implementations for this function
func SubtractV16sS16sV32s(x []Yep16s, y Yep16s) ([]Yep32s, YepStatus) {
	diff := make([]Yep32s, len(x))
	status := YepStatus(
		C.yepCore_Subtract_V16sS16s_V32s(
			(*C.Yep16s)(unsafe.Pointer(&x[0])),
			(C.Yep16s)(y),
			(*C.Yep32s)(unsafe.Pointer(&diff[0])),
			C.YepSize(len(diff)),
		),
	)
	return diff, status
}

// SubtractV16uS16uV32u Subtracts a constant to unsigned 16-bit integer array elements. Produces an array of unsigned 32-bit integer elements.
// Warning: This version of Yeppp! does not include optimized implementations for this function
func SubtractV16uS16uV32u(x []Yep16u, y Yep16u) ([]Yep32u, YepStatus) {
	diff := make([]Yep32u, len(x))
	status := YepStatus(
		C.yepCore_Subtract_V16uS16u_V32u(
			(*C.Yep16u)(unsafe.Pointer(&x[0])),
			(C.Yep16u)(y),
			(*C.Yep32u)(unsafe.Pointer(&diff[0])),
			C.YepSize(len(diff)),
		),
	)
	return diff, status
}

// SubtractV32sS32sV32s Subtracts a constant to signed 32-bit integer array elements. Produces an array of signed 32-bit integer elements.
// Warning: This version of Yeppp! does not include optimized implementations for this function
func SubtractV32sS32sV32s(x []Yep32s, y Yep32s) ([]Yep32s, YepStatus) {
	diff := make([]Yep32s, len(x))
	status := YepStatus(
		C.yepCore_Subtract_V32sS32s_V32s(
			(*C.Yep32s)(unsafe.Pointer(&x[0])),
			(C.Yep32s)(y),
			(*C.Yep32s)(unsafe.Pointer(&diff[0])),
			C.YepSize(len(diff)),
		),
	)
	return diff, status
}

// SubtractV32uS32uV64u Subtracts a constant from unsigned 32-bit integer array elements. Produces an array of unsigned 64-bit integer elements.
// Warning: This version of Yeppp! does not include optimized implementations for this function
func SubtractV32uS32uV64u(x []Yep32u, y Yep32u) ([]Yep64u, YepStatus) {
	diff := make([]Yep64u, len(x))
	status := YepStatus(
		C.yepCore_Subtract_V32uS32u_V64u(
			(*C.Yep32u)(unsafe.Pointer(&x[0])),
			(C.Yep32u)(y),
			(*C.Yep64u)(unsafe.Pointer(&diff[0])),
			C.YepSize(len(diff)),
		),
	)

	return diff, status
}

// SubtractV32sS32sV64s Subtracts a constant from signed 32-bit integer array elements. Produces an array of signed 64-bit integer elements.
// Warning: This version of Yeppp! does not include optimized implementations for this function
func SubtractV32sS32sV64s(x []Yep32s, y Yep32s) ([]Yep64s, YepStatus) {
	diff := make([]Yep64s, len(x))
	status := YepStatus(
		C.yepCore_Subtract_V32sS32s_V64s(
			(*C.Yep32s)(unsafe.Pointer(&x[0])),
			(C.Yep32s)(y),
			(*C.Yep64s)(unsafe.Pointer(&diff[0])),
			C.YepSize(len(diff)),
		),
	)
	return diff, status
}

// SubtractV64sS64sV64s Subtracts a constant to signed 64-bit integer array elements. Produces an array of signed 64-bit integer elements.
// Warning This version of Yeppp! does not include optimized implementations for this function
func SubtractV64sS64sV64s(x []Yep64s, y Yep64s) ([]Yep64s, YepStatus) {
	diff := make([]Yep64s, len(x))
	status := YepStatus(
		C.yepCore_Subtract_V64sS64s_V64s(
			(*C.Yep64s)(unsafe.Pointer(&x[0])),
			(C.Yep64s)(y),
			(*C.Yep64s)(unsafe.Pointer(&diff[0])),
			C.YepSize(len(diff)),
		),
	)
	return diff, status
}

// SubtractS32fV32fV32f Subtracts single precision (32-bit) floating-point array elements from a constant. Produces an array of single precision (32-bit) floating-point elements.
// Warning This version of Yeppp! does not include optimized implementations for this function
func SubtractS32fV32fV32f(x Yep32f, y []Yep32f) ([]Yep32f, YepStatus) {
	diff := make([]Yep32f, len(y))
	status := YepStatus(
		C.yepCore_Subtract_S32fV32f_V32f(
			(C.Yep32f)(x),
			(*C.Yep32f)(unsafe.Pointer(&y[0])),
			(*C.Yep32f)(unsafe.Pointer(&diff[0])),
			C.YepSize(len(diff)),
		),
	)
	return diff, status
}

// SubtractS64fV64fV64f Subtracts double precision (64-bit) floating-point array elements from a constant. Produces an array of double precision (64-bit) floating-point elements.
// Warning This version of Yeppp! does not include optimized implementations for this function
func SubtractS64fV64fV64f(x Yep64f, y []Yep64f) ([]Yep64f, YepStatus) {
	diff := make([]Yep64f, len(y))
	status := YepStatus(
		C.yepCore_Subtract_S64fV64f_V64f(
			(C.Yep64f)(x),
			(*C.Yep64f)(unsafe.Pointer(&y[0])),
			(*C.Yep64f)(unsafe.Pointer(&diff[0])),
			C.YepSize(len(diff)),
		),
	)
	return diff, status
}

// SubtractV32fS32fV32f Subtracts a constant from single precision (32-bit) floating-point array elements. Produces an array of single precision (32-bit) floating-point elements.
// Warning This version of Yeppp! does not include optimized implementations for this function
func SubtractV32fS32fV32f(x []Yep32f, y Yep32f) ([]Yep32f, YepStatus) {
	diff := make([]Yep32f, len(x))
	status := YepStatus(
		C.yepCore_Subtract_V32fS32f_V32f(
			(*C.Yep32f)(unsafe.Pointer(&x[0])),
			(C.Yep32f)(y),
			(*C.Yep32f)(unsafe.Pointer(&diff[0])),
			C.YepSize(len(diff)),
		),
	)
	return diff, status
}

// SubtractV64fS64fV64f Subtracts a constant to double precision (64-bit) floating-point array elements. Produces an array of double precision (64-bit) floating-point elements.
// Warning This version of Yeppp! does not include optimized implementations for this function
func SubtractV64fS64fV64f(x []Yep64f, y Yep64f) ([]Yep64f, YepStatus) {
	diff := make([]Yep64f, len(x))
	status := YepStatus(
		C.yepCore_Subtract_V64fS64f_V64f(
			(*C.Yep64f)(unsafe.Pointer(&x[0])),
			(C.Yep64f)(y),
			(*C.Yep64f)(unsafe.Pointer(&diff[0])),
			C.YepSize(len(diff)),
		),
	)
	return diff, status
}

// SubtractIV8sV8sIV8s Subtracts corresponding elements in two signed 8-bit integer arrays and writes the result to the first array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func SubtractIV8sV8sIV8s(x, y []Yep8s) YepStatus {
	status := YepStatus(
		C.yepCore_Subtract_IV8sV8s_IV8s(
			(*C.Yep8s)(unsafe.Pointer(&x[0])),
			(*C.Yep8s)(unsafe.Pointer(&y[0])),
			C.YepSize(len(x)),
		),
	)
	return status
}

// SubtractIV16sV16sIV16s Subtracts corresponding elements in two signed 16-bit integer arrays and writes the result to the first array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func SubtractIV16sV16sIV16s(x, y []Yep16s) YepStatus {
	status := YepStatus(
		C.yepCore_Subtract_IV16sV16s_IV16s(
			(*C.Yep16s)(unsafe.Pointer(&x[0])),
			(*C.Yep16s)(unsafe.Pointer(&y[0])),
			C.YepSize(len(x)),
		),
	)
	return status
}

// SubtractIV32sV32sIV32s Subtracts corresponding elements in two signed 32-bit integer arrays and writes the result to the first array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func SubtractIV32sV32sIV32s(x, y []Yep32s) YepStatus {
	status := YepStatus(
		C.yepCore_Subtract_IV32sV32s_IV32s(
			(*C.Yep32s)(unsafe.Pointer(&x[0])),
			(*C.Yep32s)(unsafe.Pointer(&y[0])),
			C.YepSize(len(x)),
		),
	)
	return status
}

// SubtractIV64sV64sIV64s Subtracts corresponding elements in two signed 64-bit integer arrays and writes the result to the first array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func SubtractIV64sV64sIV64s(x, y []Yep64s) YepStatus {
	status := YepStatus(
		C.yepCore_Subtract_IV64sV64s_IV64s(
			(*C.Yep64s)(unsafe.Pointer(&x[0])),
			(*C.Yep64s)(unsafe.Pointer(&y[0])),
			C.YepSize(len(x)),
		),
	)
	return status
}

// SubtractIV32fV32fIV32f Subtracts corresponding elements in two single precision (32-bit) floating-point arrays and writes the result to the first array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func SubtractIV32fV32fIV32f(x, y []Yep32f) YepStatus {
	status := YepStatus(
		C.yepCore_Subtract_IV32fV32f_IV32f(
			(*C.Yep32f)(unsafe.Pointer(&x[0])),
			(*C.Yep32f)(unsafe.Pointer(&y[0])),
			C.YepSize(len(x)),
		),
	)
	return status
}

// SubtractIV64fV64fIV64f Subtracts corresponding elements in two double precision (64-bit) floating-point arrays and writes the result to the first array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func SubtractIV64fV64fIV64f(x, y []Yep64f) YepStatus {
	status := YepStatus(
		C.yepCore_Subtract_IV64fV64f_IV64f(
			(*C.Yep64f)(unsafe.Pointer(&x[0])),
			(*C.Yep64f)(unsafe.Pointer(&y[0])),
			C.YepSize(len(x)),
		),
	)
	return status
}

// SubtractIV8sS8sIV8s Subtracts corresponding elements in two signed 8-bit integer arrays and writes the result to the first array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func SubtractIV8sS8sIV8s(x []Yep8s, y Yep8s) YepStatus {
	status := YepStatus(
		C.yepCore_Subtract_IV8sS8s_IV8s(
			(*C.Yep8s)(unsafe.Pointer(&x[0])),
			(C.Yep8s)(y),
			C.YepSize(len(x)),
		),
	)
	return status
}

// SubtractIV16sS16sIV16s Subtracts a constant from signed 16-bit integer array elements and writes the result to the same array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func SubtractIV16sS16sIV16s(x []Yep16s, y Yep16s) YepStatus {
	status := YepStatus(
		C.yepCore_Subtract_IV16sS16s_IV16s(
			(*C.Yep16s)(unsafe.Pointer(&x[0])),
			(C.Yep16s)(y),
			C.YepSize(len(x)),
		),
	)
	return status
}

// SubtractIV32sS32sIV32s Subtracts a constant from signed 32-bit integer array elements and writes the result to the same array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func SubtractIV32sS32sIV32s(x []Yep32s, y Yep32s) YepStatus {
	status := YepStatus(
		C.yepCore_Subtract_IV32sS32s_IV32s(
			(*C.Yep32s)(unsafe.Pointer(&x[0])),
			(C.Yep32s)(y),
			C.YepSize(len(x)),
		),
	)
	return status
}

// SubtractIV64sS64sIV64s Subtracts a constant from signed 64-bit integer array elements and writes the result to the same array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func SubtractIV64sS64sIV64s(x []Yep64s, y Yep64s) YepStatus {
	status := YepStatus(
		C.yepCore_Subtract_IV64sS64s_IV64s(
			(*C.Yep64s)(unsafe.Pointer(&x[0])),
			(C.Yep64s)(y),
			C.YepSize(len(x)),
		),
	)
	return status
}

// SubtractIV32fS32fIV32f Subtracts a constant to single precision (32-bit) floating-point array elements and writes the result to the same array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func SubtractIV32fS32fIV32f(x []Yep32f, y Yep32f) YepStatus {
	status := YepStatus(
		C.yepCore_Subtract_IV32fS32f_IV32f(
			(*C.Yep32f)(unsafe.Pointer(&x[0])),
			(C.Yep32f)(y),
			C.YepSize(len(x)),
		),
	)
	return status
}

// SubtractIV64fS64fIV64f Subtracts a constant from double precision (64-bit) floating-point array elements and writes the result to the same array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func SubtractIV64fS64fIV64f(x []Yep64f, y Yep64f) YepStatus {
	status := YepStatus(
		C.yepCore_Subtract_IV64fS64f_IV64f(
			(*C.Yep64f)(unsafe.Pointer(&x[0])),
			(C.Yep64f)(y),
			C.YepSize(len(x)),
		),
	)
	return status
}

// SubtractV8sIV8sIV8s Subtracts corresponding elements in two signed 8-bit integer arrays and writes the result to the second array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func SubtractV8sIV8sIV8s(x, y []Yep8s) YepStatus {
	status := YepStatus(
		C.yepCore_Subtract_V8sIV8s_IV8s(
			(*C.Yep8s)(unsafe.Pointer(&x[0])),
			(*C.Yep8s)(unsafe.Pointer(&y[0])),
			C.YepSize(len(x)),
		),
	)
	return status
}

// SubtractV16sIV16sIV16s Subtracts corresponding elements in two signed 16-bit integer arrays and writes the result to the second array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func SubtractV16sIV16sIV16s(x, y []Yep16s) YepStatus {
	status := YepStatus(
		C.yepCore_Subtract_V16sIV16s_IV16s(
			(*C.Yep16s)(unsafe.Pointer(&x[0])),
			(*C.Yep16s)(unsafe.Pointer(&y[0])),
			C.YepSize(len(x)),
		),
	)
	return status
}

// SubtractV32sIV32sIV32s Subtracts corresponding elements in two signed 32-bit integer arrays and writes the result to the second array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func SubtractV32sIV32sIV32s(x, y []Yep32s) YepStatus {
	status := YepStatus(
		C.yepCore_Subtract_V32sIV32s_IV32s(
			(*C.Yep32s)(unsafe.Pointer(&x[0])),
			(*C.Yep32s)(unsafe.Pointer(&y[0])),
			C.YepSize(len(x)),
		),
	)
	return status
}

// SubtractV64sIV64sIV64s Subtracts corresponding elements in two signed 64-bit integer arrays and writes the result to the second array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func SubtractV64sIV64sIV64s(x, y []Yep64s) YepStatus {
	status := YepStatus(
		C.yepCore_Subtract_V64sIV64s_IV64s(
			(*C.Yep64s)(unsafe.Pointer(&x[0])),
			(*C.Yep64s)(unsafe.Pointer(&y[0])),
			C.YepSize(len(x)),
		),
	)
	return status
}

// SubtractV32fIV32fIV32f Subtracts corresponding elements in two single precision (32-bit) floating-point arrays and writes the result to the second array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func SubtractV32fIV32fIV32f(x, y []Yep32f) YepStatus {
	status := YepStatus(
		C.yepCore_Subtract_V32fIV32f_IV32f(
			(*C.Yep32f)(unsafe.Pointer(&x[0])),
			(*C.Yep32f)(unsafe.Pointer(&y[0])),
			C.YepSize(len(x)),
		),
	)
	return status
}

// SubtractV64fIV64fIV64f Subtracts corresponding elements in two double precision (64-bit) floating-point arrays and writes the result to the second array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func SubtractV64fIV64fIV64f(x, y []Yep64f) YepStatus {
	status := YepStatus(
		C.yepCore_Subtract_V64fIV64f_IV64f(
			(*C.Yep64f)(unsafe.Pointer(&x[0])),
			(*C.Yep64f)(unsafe.Pointer(&y[0])),
			C.YepSize(len(x)),
		),
	)
	return status
}

// SubtractS8sV8sV8s Subtracts signed 8-bit integer array elements from a constant. Produces an array of signed 8-bit integer elements.
// Warning This version of Yeppp! does not include optimized implementations for this function
func SubtractS8sV8sV8s(x Yep8s, y []Yep8s) ([]Yep8s, YepStatus) {
	diff := make([]Yep8s, len(y))

	status := YepStatus(
		C.yepCore_Subtract_S8sV8s_V8s(
			(C.Yep8s)(x),
			(*C.Yep8s)(unsafe.Pointer(&y[0])),
			(*C.Yep8s)(unsafe.Pointer(&diff[0])),
			C.YepSize(len(y)),
		),
	)

	return diff, status
}

// SubtractS8sV8sV16s Subtracts signed 8-bit integer array elements from a constant. Produces an array of signed 16-bit integer elements.
// Warning This version of Yeppp! does not include optimized implementations for this function
func SubtractS8sV8sV16s(x Yep8s, y []Yep8s) ([]Yep16s, YepStatus) {
	diff := make([]Yep16s, len(y))

	status := YepStatus(
		C.yepCore_Subtract_S8sV8s_V16s(
			(C.Yep8s)(x),
			(*C.Yep8s)(unsafe.Pointer(&y[0])),
			(*C.Yep16s)(unsafe.Pointer(&diff[0])),
			C.YepSize(len(y)),
		),
	)

	return diff, status
}

// SubtractS8uV8uV16u Subtracts unsigned 8-bit integer array elements from a constant. Produces an array of unsigned 16-bit integer elements.
// Warning This version of Yeppp! does not include optimized implementations for this function
func SubtractS8uV8uV16u(x Yep8u, y []Yep8u) ([]Yep16u, YepStatus) {
	diff := make([]Yep16u, len(y))

	status := YepStatus(
		C.yepCore_Subtract_S8uV8u_V16u(
			(C.Yep8u)(x),
			(*C.Yep8u)(unsafe.Pointer(&y[0])),
			(*C.Yep16u)(unsafe.Pointer(&diff[0])),
			C.YepSize(len(y)),
		),
	)

	return diff, status
}

// SubtractS16sV16sV16s Subtracts signed 16-bit integer array elements from a constant. Produces an array of signed 16-bit integer elements.
// Warning This version of Yeppp! does not include optimized implementations for this function
func SubtractS16sV16sV16s(x Yep16s, y []Yep16s) ([]Yep16s, YepStatus) {
	diff := make([]Yep16s, len(y))

	status := YepStatus(
		C.yepCore_Subtract_S16sV16s_V16s(
			(C.Yep16s)(x),
			(*C.Yep16s)(unsafe.Pointer(&y[0])),
			(*C.Yep16s)(unsafe.Pointer(&diff[0])),
			C.YepSize(len(y)),
		),
	)

	return diff, status
}

// SubtractS16sV16sV32s Subtracts signed 16-bit integer array elements from a constant. Produces an array of signed 32-bit integer elements.
// Warning This version of Yeppp! does not include optimized implementations for this function
func SubtractS16sV16sV32s(x Yep16s, y []Yep16s) ([]Yep32s, YepStatus) {
	diff := make([]Yep32s, len(y))

	status := YepStatus(
		C.yepCore_Subtract_S16sV16s_V32s(
			(C.Yep16s)(x),
			(*C.Yep16s)(unsafe.Pointer(&y[0])),
			(*C.Yep32s)(unsafe.Pointer(&diff[0])),
			C.YepSize(len(y)),
		),
	)

	return diff, status
}

// SubtractS16uV16uV32u Subtracts unsigned 16-bit integer array elements from a constant. Produces an array of unsigned 32-bit integer elements.
// Warning This version of Yeppp! does not include optimized implementations for this function
func SubtractS16uV16uV32u(x Yep16u, y []Yep16u) ([]Yep32u, YepStatus) {
	diff := make([]Yep32u, len(y))

	status := YepStatus(
		C.yepCore_Subtract_S16uV16u_V32u(
			(C.Yep16u)(x),
			(*C.Yep16u)(unsafe.Pointer(&y[0])),
			(*C.Yep32u)(unsafe.Pointer(&diff[0])),
			C.YepSize(len(y)),
		),
	)

	return diff, status
}

// SubtractS32sV32sV32s Subtracts signed 32-bit integer array elements from a constant. Produces an array of signed 32-bit integer elements.
// Warning This version of Yeppp! does not include optimized implementations for this function
func SubtractS32sV32sV32s(x Yep32s, y []Yep32s) ([]Yep32s, YepStatus) {
	diff := make([]Yep32s, len(y))

	status := YepStatus(
		C.yepCore_Subtract_S32sV32s_V32s(
			(C.Yep32s)(x),
			(*C.Yep32s)(unsafe.Pointer(&y[0])),
			(*C.Yep32s)(unsafe.Pointer(&diff[0])),
			C.YepSize(len(y)),
		),
	)

	return diff, status
}

// SubtractS32sV32sV64s Subtracts signed 32-bit integer array elements from a constant. Produces an array of signed 64-bit integer elements.
// Warning This version of Yeppp! does not include optimized implementations for this function
func SubtractS32sV32sV64s(x Yep32s, y []Yep32s) ([]Yep64s, YepStatus) {
	diff := make([]Yep64s, len(y))

	status := YepStatus(
		C.yepCore_Subtract_S32sV32s_V64s(
			(C.Yep32s)(x),
			(*C.Yep32s)(unsafe.Pointer(&y[0])),
			(*C.Yep64s)(unsafe.Pointer(&diff[0])),
			C.YepSize(len(y)),
		),
	)

	return diff, status
}

// SubtractS32uV32uV64u Subtracts unsigned 32-bit integer array elements from a constant. Produces an array of unsigned 64-bit integer elements.
// Warning This version of Yeppp! does not include optimized implementations for this function
func SubtractS32uV32uV64u(x Yep32u, y []Yep32u) ([]Yep64u, YepStatus) {
	diff := make([]Yep64u, len(y))

	status := YepStatus(
		C.yepCore_Subtract_S32uV32u_V64u(
			(C.Yep32u)(x),
			(*C.Yep32u)(unsafe.Pointer(&y[0])),
			(*C.Yep64u)(unsafe.Pointer(&diff[0])),
			C.YepSize(len(y)),
		),
	)

	return diff, status
}

// SubtractS64sV64sV64s Subtracts signed 64-bit integer array elements from a constant. Produces an array of signed 64-bit integer elements.
// Warning This version of Yeppp! does not include optimized implementations for this function
func SubtractS64sV64sV64s(x Yep64s, y []Yep64s) ([]Yep64s, YepStatus) {
	diff := make([]Yep64s, len(y))

	status := YepStatus(
		C.yepCore_Subtract_S64sV64s_V64s(
			(C.Yep64s)(x),
			(*C.Yep64s)(unsafe.Pointer(&y[0])),
			(*C.Yep64s)(unsafe.Pointer(&diff[0])),
			C.YepSize(len(y)),
		),
	)

	return diff, status
}

// SubtractS8sIV8sIV8s Subtracts signed 8-bit integer array elements from a constant and writes the result to the same array.
// warning This version of Yeppp! does not include optimized implementations for this function
func SubtractS8sIV8sIV8s(x Yep8s, y []Yep8s) YepStatus {
	status := YepStatus(
		C.yepCore_Subtract_S8sIV8s_IV8s(
			(C.Yep8s)(x),
			(*C.Yep8s)(unsafe.Pointer(&y[0])),
			C.YepSize(len(y)),
		),
	)
	return status
}

// SubtractS16sIV16sIV16s Subtracts signed 8-bit integer array elements from a constant and writes the result to the same array.
// warning This version of Yeppp! does not include optimized implementations for this function
func SubtractS16sIV16sIV16s(x Yep16s, y []Yep16s) YepStatus {
	status := YepStatus(
		C.yepCore_Subtract_S16sIV16s_IV16s(
			(C.Yep16s)(x),
			(*C.Yep16s)(unsafe.Pointer(&y[0])),
			C.YepSize(len(y)),
		),
	)
	return status
}

// SubtractS32sIV32sIV32s Subtracts signed 32-bit integer array elements from a constant and writes the result to the same array.
// warning This version of Yeppp! does not include optimized implementations for this function
func SubtractS32sIV32sIV32s(x Yep32s, y []Yep32s) YepStatus {
	status := YepStatus(
		C.yepCore_Subtract_S32sIV32s_IV32s(
			(C.Yep32s)(x),
			(*C.Yep32s)(unsafe.Pointer(&y[0])),
			C.YepSize(len(y)),
		),
	)
	return status
}

// SubtractS64sIV64sIV64s Subtracts signed 64-bit integer array elements from a constant and writes the result to the same array.
// warning This version of Yeppp! does not include optimized implementations for this function
func SubtractS64sIV64sIV64s(x Yep64s, y []Yep64s) YepStatus {
	status := YepStatus(
		C.yepCore_Subtract_S64sIV64s_IV64s(
			(C.Yep64s)(x),
			(*C.Yep64s)(unsafe.Pointer(&y[0])),
			C.YepSize(len(y)),
		),
	)
	return status
}

// SubtractS32fIV32fIV32f Subtracts single precision (32-bit) floating-point array elements from a constant and writes the result to the same array.
// warning This version of Yeppp! does not include optimized implementations for this function
func SubtractS32fIV32fIV32f(x Yep32f, y []Yep32f) YepStatus {
	status := YepStatus(
		C.yepCore_Subtract_S32fIV32f_IV32f(
			(C.Yep32f)(x),
			(*C.Yep32f)(unsafe.Pointer(&y[0])),
			C.YepSize(len(y)),
		),
	)
	return status
}

// SubtractS64fIV64fIV64f Subtracts double precision (64-bit) floating-point array elements from a constant and writes the result to the same array.
// warning This version of Yeppp! does not include optimized implementations for this function
func SubtractS64fIV64fIV64f(x Yep64f, y []Yep64f) YepStatus {
	status := YepStatus(
		C.yepCore_Subtract_S64fIV64f_IV64f(
			(C.Yep64f)(x),
			(*C.Yep64f)(unsafe.Pointer(&y[0])),
			C.YepSize(len(y)),
		),
	)
	return status
}
