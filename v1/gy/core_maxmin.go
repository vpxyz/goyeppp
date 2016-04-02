package gy

/*
#include <yepCore.h>

#cgo pkg-config: yeppp

*/
import "C"

import (
	"unsafe"
)

// MinV8sS8s Computes the minimum of signed 8-bit integer array elements.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MinV8sS8s(v []Yep8s) (Yep8s, YepStatus) {
	var minimum Yep8s
	status := YepStatus(
		C.yepCore_Min_V8s_S8s(
			(*C.Yep8s)(unsafe.Pointer(&v[0])),
			(*C.Yep8s)(unsafe.Pointer(&minimum)),
			C.YepSize(len(v)),
		),
	)

	return minimum, status
}

// MinV8uS8u Computes the minimum of unsigned 8-bit integer array elements.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MinV8uS8u(v []Yep8u) (Yep8u, YepStatus) {
	var minimum Yep8u
	status := YepStatus(
		C.yepCore_Min_V8u_S8u(
			(*C.Yep8u)(unsafe.Pointer(&v[0])),
			(*C.Yep8u)(unsafe.Pointer(&minimum)),
			C.YepSize(len(v)),
		),
	)

	return minimum, status
}

// MinV16sS16s Computes the minimum of signed 16-bit integer array elements.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MinV16sS16s(v []Yep16s) (Yep16s, YepStatus) {
	var minimum Yep16s
	status := YepStatus(
		C.yepCore_Min_V16s_S16s(
			(*C.Yep16s)(unsafe.Pointer(&v[0])),
			(*C.Yep16s)(unsafe.Pointer(&minimum)),
			C.YepSize(len(v)),
		),
	)

	return minimum, status
}

// MinV16uS16u Computes the minimum of unsigned 16-bit integer array elements.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MinV16uS16u(v []Yep16u) (Yep16u, YepStatus) {
	var minimum Yep16u
	status := YepStatus(
		C.yepCore_Min_V16u_S16u(
			(*C.Yep16u)(unsafe.Pointer(&v[0])),
			(*C.Yep16u)(unsafe.Pointer(&minimum)),
			C.YepSize(len(v)),
		),
	)

	return minimum, status
}

// MinV32sS32s Computes the minimum of signed 32-bit integer array elements.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MinV32sS32s(v []Yep32s) (Yep32s, YepStatus) {
	var minimum Yep32s
	status := YepStatus(
		C.yepCore_Min_V32s_S32s(
			(*C.Yep32s)(unsafe.Pointer(&v[0])),
			(*C.Yep32s)(unsafe.Pointer(&minimum)),
			C.YepSize(len(v)),
		),
	)

	return minimum, status
}

// MinV32uS32u Computes the minimum of unsigned 32-bit integer array elements.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MinV32uS32u(v []Yep32u) (Yep32u, YepStatus) {
	var minimum Yep32u
	status := YepStatus(
		C.yepCore_Min_V32u_S32u(
			(*C.Yep32u)(unsafe.Pointer(&v[0])),
			(*C.Yep32u)(unsafe.Pointer(&minimum)),
			C.YepSize(len(v)),
		),
	)

	return minimum, status
}

// MinV64sS64s Computes the minimum of signed 64-bit integer array elements.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MinV64sS64s(v []Yep64s) (Yep64s, YepStatus) {
	var minimum Yep64s
	status := YepStatus(
		C.yepCore_Min_V64s_S64s(
			(*C.Yep64s)(unsafe.Pointer(&v[0])),
			(*C.Yep64s)(unsafe.Pointer(&minimum)),
			C.YepSize(len(v)),
		),
	)

	return minimum, status
}

// MinV64uS64u Computes the minimum of unsigned 64-bit integer array elements.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MinV64uS64u(v []Yep64u) (Yep64u, YepStatus) {
	var minimum Yep64u
	status := YepStatus(
		C.yepCore_Min_V64u_S64u(
			(*C.Yep64u)(unsafe.Pointer(&v[0])),
			(*C.Yep64u)(unsafe.Pointer(&minimum)),
			C.YepSize(len(v)),
		),
	)

	return minimum, status
}

// MinV32fS32f Computes the minimum of single precision (32-bit) floating-point array elements.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MinV32fS32f(v []Yep32f) (Yep32f, YepStatus) {
	var minimum Yep32f
	status := YepStatus(
		C.yepCore_Min_V32f_S32f(
			(*C.Yep32f)(unsafe.Pointer(&v[0])),
			(*C.Yep32f)(unsafe.Pointer(&minimum)),
			C.YepSize(len(v)),
		),
	)

	return minimum, status
}

// MinV64fS64f Computes the minimum of single precision (64-bit) floating-point array elements.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MinV64fS64f(v []Yep64f) (Yep64f, YepStatus) {
	var minimum Yep64f
	status := YepStatus(
		C.yepCore_Min_V64f_S64f(
			(*C.Yep64f)(unsafe.Pointer(&v[0])),
			(*C.Yep64f)(unsafe.Pointer(&minimum)),
			C.YepSize(len(v)),
		),
	)

	return minimum, status
}

// MinV8sV8sV8s Computes pairwise minima of corresponding elements in two signed 8-bit integer arrays.
func MinV8sV8sV8s(x, y []Yep8s) ([]Yep8s, YepStatus) {
	minimum := make([]Yep8s, len(x))
	status := YepStatus(
		C.yepCore_Min_V8sV8s_V8s(
			(*C.Yep8s)(unsafe.Pointer(&x[0])),
			(*C.Yep8s)(unsafe.Pointer(&y[0])),
			(*C.Yep8s)(unsafe.Pointer(&minimum[0])),
			C.YepSize(len(x)),
		),
	)

	return minimum, status
}

// MinV8uV8uV8u Computes pairwise minima of corresponding elements in two unsigned 8-bit integer arrays.
func MinV8uV8uV8u(x, y []Yep8u) ([]Yep8u, YepStatus) {
	minimum := make([]Yep8u, len(x))
	status := YepStatus(
		C.yepCore_Min_V8uV8u_V8u(
			(*C.Yep8u)(unsafe.Pointer(&x[0])),
			(*C.Yep8u)(unsafe.Pointer(&y[0])),
			(*C.Yep8u)(unsafe.Pointer(&minimum[0])),
			C.YepSize(len(x)),
		),
	)

	return minimum, status
}

// MinV16sV16sV16s Computes pairwise minima of corresponding elements in two signed 16-bit integer arrays.
func MinV16sV16sV16s(x, y []Yep16s) ([]Yep16s, YepStatus) {
	minimum := make([]Yep16s, len(x))
	status := YepStatus(
		C.yepCore_Min_V16sV16s_V16s(
			(*C.Yep16s)(unsafe.Pointer(&x[0])),
			(*C.Yep16s)(unsafe.Pointer(&y[0])),
			(*C.Yep16s)(unsafe.Pointer(&minimum[0])),
			C.YepSize(len(x)),
		),
	)

	return minimum, status
}

// MinV16uV16uV16u Computes pairwise minima of corresponding elements in two unsigned 16-bit integer arrays.
func MinV16uV16uV16u(x, y []Yep16u) ([]Yep16u, YepStatus) {
	minimum := make([]Yep16u, len(x))
	status := YepStatus(
		C.yepCore_Min_V16uV16u_V16u(
			(*C.Yep16u)(unsafe.Pointer(&x[0])),
			(*C.Yep16u)(unsafe.Pointer(&y[0])),
			(*C.Yep16u)(unsafe.Pointer(&minimum[0])),
			C.YepSize(len(x)),
		),
	)

	return minimum, status
}

// MinV32sV32sV32s Computes pairwise minima of corresponding elements in two signed 32-bit integer arrays.
func MinV32sV32sV32s(x, y []Yep32s) ([]Yep32s, YepStatus) {
	minimum := make([]Yep32s, len(x))
	status := YepStatus(
		C.yepCore_Min_V32sV32s_V32s(
			(*C.Yep32s)(unsafe.Pointer(&x[0])),
			(*C.Yep32s)(unsafe.Pointer(&y[0])),
			(*C.Yep32s)(unsafe.Pointer(&minimum[0])),
			C.YepSize(len(x)),
		),
	)

	return minimum, status
}

// MinV32uV32uV32u Computes pairwise minima of corresponding elements in two unsigned 32-bit integer arrays.
func MinV32uV32uV32u(x, y []Yep32u) ([]Yep32u, YepStatus) {
	minimum := make([]Yep32u, len(x))
	status := YepStatus(
		C.yepCore_Min_V32uV32u_V32u(
			(*C.Yep32u)(unsafe.Pointer(&x[0])),
			(*C.Yep32u)(unsafe.Pointer(&y[0])),
			(*C.Yep32u)(unsafe.Pointer(&minimum[0])),
			C.YepSize(len(x)),
		),
	)

	return minimum, status
}

// MinV64sV32sV64s Computes pairwise minima of corresponding elements in two signed 64-bit integer arrays.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MinV64sV32sV64s(x []Yep64s, y []Yep32s) ([]Yep64s, YepStatus) {
	minimum := make([]Yep64s, len(x))
	status := YepStatus(
		C.yepCore_Min_V64sV32s_V64s(
			(*C.Yep64s)(unsafe.Pointer(&x[0])),
			(*C.Yep32s)(unsafe.Pointer(&y[0])),
			(*C.Yep64s)(unsafe.Pointer(&minimum[0])),
			C.YepSize(len(x)),
		),
	)

	return minimum, status
}

// MinV64uV32uV64u Computes pairwise minima of corresponding elements in two unsigned 64-bit integer arrays.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MinV64uV32uV64u(x []Yep64u, y []Yep32u) ([]Yep64u, YepStatus) {
	minimum := make([]Yep64u, len(x))
	status := YepStatus(
		C.yepCore_Min_V64uV32u_V64u(
			(*C.Yep64u)(unsafe.Pointer(&x[0])),
			(*C.Yep32u)(unsafe.Pointer(&y[0])),
			(*C.Yep64u)(unsafe.Pointer(&minimum[0])),
			C.YepSize(len(x)),
		),
	)

	return minimum, status
}

// MinV32fV32fV32f Computes pairwise minima of corresponding elements in two single precision (32-bit) floating-point arrays.
func MinV32fV32fV32f(x, y []Yep32f) ([]Yep32f, YepStatus) {
	minimum := make([]Yep32f, len(x))
	status := YepStatus(
		C.yepCore_Min_V32fV32f_V32f(
			(*C.Yep32f)(unsafe.Pointer(&x[0])),
			(*C.Yep32f)(unsafe.Pointer(&y[0])),
			(*C.Yep32f)(unsafe.Pointer(&minimum[0])),
			C.YepSize(len(x)),
		),
	)

	return minimum, status
}

// MinV64fV64fV64f Computes pairwise minima of corresponding elements in two double precision (64-bit) floating-point arrays.
func MinV64fV64fV64f(x, y []Yep64f) ([]Yep64f, YepStatus) {
	minimum := make([]Yep64f, len(x))
	status := YepStatus(
		C.yepCore_Min_V64fV64f_V64f(
			(*C.Yep64f)(unsafe.Pointer(&x[0])),
			(*C.Yep64f)(unsafe.Pointer(&y[0])),
			(*C.Yep64f)(unsafe.Pointer(&minimum[0])),
			C.YepSize(len(x)),
		),
	)

	return minimum, status
}

// MinV8sS8sV8s Computes pairwise minima of signed 8-bit integer array elements and a constant.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MinV8sS8sV8s(x []Yep8s, y Yep8s) ([]Yep8s, YepStatus) {
	minimum := make([]Yep8s, len(x))
	status := YepStatus(
		C.yepCore_Min_V8sS8s_V8s(
			(*C.Yep8s)(unsafe.Pointer(&x[0])),
			(C.Yep8s)(y),
			(*C.Yep8s)(unsafe.Pointer(&minimum[0])),
			C.YepSize(len(x)),
		),
	)
	return minimum, status
}

// MinV8uS8uV8u Computes pairwise minima of unsigned 8-bit integer array elements and a constant.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MinV8uS8uV8u(x []Yep8u, y Yep8u) ([]Yep8u, YepStatus) {
	minimum := make([]Yep8u, len(x))
	status := YepStatus(
		C.yepCore_Min_V8uS8u_V8u(
			(*C.Yep8u)(unsafe.Pointer(&x[0])),
			(C.Yep8u)(y),
			(*C.Yep8u)(unsafe.Pointer(&minimum[0])),
			C.YepSize(len(x)),
		),
	)
	return minimum, status
}

// MinV16sS16sV16s Computes pairwise minima of signed 16-bit integer array elements and a constant.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MinV16sS16sV16s(x []Yep16s, y Yep16s) ([]Yep16s, YepStatus) {
	minimum := make([]Yep16s, len(x))
	status := YepStatus(
		C.yepCore_Min_V16sS16s_V16s(
			(*C.Yep16s)(unsafe.Pointer(&x[0])),
			(C.Yep16s)(y),
			(*C.Yep16s)(unsafe.Pointer(&minimum[0])),
			C.YepSize(len(x)),
		),
	)
	return minimum, status
}

// MinV16uS16uV16u Computes pairwise minima of unsigned 16-bit integer array elements and a constant.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MinV16uS16uV16u(x []Yep16u, y Yep16u) ([]Yep16u, YepStatus) {
	minimum := make([]Yep16u, len(x))
	status := YepStatus(
		C.yepCore_Min_V16uS16u_V16u(
			(*C.Yep16u)(unsafe.Pointer(&x[0])),
			(C.Yep16u)(y),
			(*C.Yep16u)(unsafe.Pointer(&minimum[0])),
			C.YepSize(len(x)),
		),
	)
	return minimum, status
}

// MinV32sS32sV32s Computes pairwise minima of signed 32-bit integer array elements and a constant.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MinV32sS32sV32s(x []Yep32s, y Yep32s) ([]Yep32s, YepStatus) {
	minimum := make([]Yep32s, len(x))
	status := YepStatus(
		C.yepCore_Min_V32sS32s_V32s(
			(*C.Yep32s)(unsafe.Pointer(&x[0])),
			(C.Yep32s)(y),
			(*C.Yep32s)(unsafe.Pointer(&minimum[0])),
			C.YepSize(len(x)),
		),
	)
	return minimum, status
}

// MinV32uS32uV32u Computes pairwise minima of unsigned 32-bit integer array elements and a constant.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MinV32uS32uV32u(x []Yep32u, y Yep32u) ([]Yep32u, YepStatus) {
	minimum := make([]Yep32u, len(x))
	status := YepStatus(
		C.yepCore_Min_V32uS32u_V32u(
			(*C.Yep32u)(unsafe.Pointer(&x[0])),
			(C.Yep32u)(y),
			(*C.Yep32u)(unsafe.Pointer(&minimum[0])),
			C.YepSize(len(x)),
		),
	)
	return minimum, status
}

// MinV64sS32sV64s Computes pairwise minima of signed 64-bit integer array elements and a constant.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MinV64sS32sV64s(x []Yep64s, y Yep32s) ([]Yep64s, YepStatus) {
	minimum := make([]Yep64s, len(x))
	status := YepStatus(
		C.yepCore_Min_V64sS32s_V64s(
			(*C.Yep64s)(unsafe.Pointer(&x[0])),
			(C.Yep32s)(y),
			(*C.Yep64s)(unsafe.Pointer(&minimum[0])),
			C.YepSize(len(x)),
		),
	)
	return minimum, status
}

// MinV64uS32uV64u Computes pairwise minima of unsigned 64-bit integer array elements and a constant.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MinV64uS32uV64u(x []Yep64u, y Yep32u) ([]Yep64u, YepStatus) {
	minimum := make([]Yep64u, len(x))
	status := YepStatus(
		C.yepCore_Min_V64uS32u_V64u(
			(*C.Yep64u)(unsafe.Pointer(&x[0])),
			(C.Yep32u)(y),
			(*C.Yep64u)(unsafe.Pointer(&minimum[0])),
			C.YepSize(len(x)),
		),
	)
	return minimum, status
}

// MinV32fS32fV32f Computes pairwise minima of single precision (32-bit) floating-point array elements and a constant.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MinV32fS32fV32f(x []Yep32f, y Yep32f) ([]Yep32f, YepStatus) {
	minimum := make([]Yep32f, len(x))
	status := YepStatus(
		C.yepCore_Min_V32fS32f_V32f(
			(*C.Yep32f)(unsafe.Pointer(&x[0])),
			(C.Yep32f)(y),
			(*C.Yep32f)(unsafe.Pointer(&minimum[0])),
			C.YepSize(len(x)),
		),
	)
	return minimum, status
}

// MinV64fS64fV64f Computes pairwise minima of double precision (64-bit) floating-point array elements and a constant.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MinV64fS64fV64f(x []Yep64f, y Yep64f) ([]Yep64f, YepStatus) {
	minimum := make([]Yep64f, len(x))
	status := YepStatus(
		C.yepCore_Min_V64fS64f_V64f(
			(*C.Yep64f)(unsafe.Pointer(&x[0])),
			(C.Yep64f)(y),
			(*C.Yep64f)(unsafe.Pointer(&minimum[0])),
			C.YepSize(len(x)),
		),
	)
	return minimum, status
}

// MinIV8sV8sIV8s Computes pairwise minima of corresponding elements in two signed 8-bit integer arrays and writes the result to the first array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MinIV8sV8sIV8s(x, y []Yep8s) YepStatus {
	status := YepStatus(
		C.yepCore_Min_IV8sV8s_IV8s(
			(*C.Yep8s)(unsafe.Pointer(&x[0])),
			(*C.Yep8s)(unsafe.Pointer(&y[0])),
			C.YepSize(len(x)),
		),
	)

	return status
}

// MinIV8uV8uIV8u Computes pairwise minima of corresponding elements in two unsigned 8-bit integer arrays and writes the result to the first array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MinIV8uV8uIV8u(x, y []Yep8u) YepStatus {
	status := YepStatus(
		C.yepCore_Min_IV8uV8u_IV8u(
			(*C.Yep8u)(unsafe.Pointer(&x[0])),
			(*C.Yep8u)(unsafe.Pointer(&y[0])),
			C.YepSize(len(x)),
		),
	)

	return status
}

// MinIV16sV16sIV16s Computes pairwise minima of corresponding elements in two signed 16-bit integer arrays and writes the result to the first array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MinIV16sV16sIV16s(x, y []Yep16s) YepStatus {
	status := YepStatus(
		C.yepCore_Min_IV16sV16s_IV16s(
			(*C.Yep16s)(unsafe.Pointer(&x[0])),
			(*C.Yep16s)(unsafe.Pointer(&y[0])),
			C.YepSize(len(x)),
		),
	)

	return status
}

// MinIV16uV16uIV16u Computes pairwise minima of corresponding elements in two unsigned 16-bit integer arrays and writes the result to the first array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MinIV16uV16uIV16u(x, y []Yep16u) YepStatus {
	status := YepStatus(
		C.yepCore_Min_IV16uV16u_IV16u(
			(*C.Yep16u)(unsafe.Pointer(&x[0])),
			(*C.Yep16u)(unsafe.Pointer(&y[0])),
			C.YepSize(len(x)),
		),
	)

	return status
}

// MinIV32sV32sIV32s Computes pairwise minima of corresponding elements in two signed 32-bit integer arrays and writes the result to the first array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MinIV32sV32sIV32s(x, y []Yep32s) YepStatus {
	status := YepStatus(
		C.yepCore_Min_IV32sV32s_IV32s(
			(*C.Yep32s)(unsafe.Pointer(&x[0])),
			(*C.Yep32s)(unsafe.Pointer(&y[0])),
			C.YepSize(len(x)),
		),
	)

	return status
}

// MinIV32uV32uIV32u Computes pairwise minima of corresponding elements in two unsigned 32-bit integer arrays and writes the result to the first array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MinIV32uV32uIV32u(x, y []Yep32u) YepStatus {
	status := YepStatus(
		C.yepCore_Min_IV32uV32u_IV32u(
			(*C.Yep32u)(unsafe.Pointer(&x[0])),
			(*C.Yep32u)(unsafe.Pointer(&y[0])),
			C.YepSize(len(x)),
		),
	)

	return status
}

// MinIV64sV32sIV64s Computes pairwise minima of corresponding elements in two signed 64-bit integer arrays and writes the result to the first array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MinIV64sV32sIV64s(x []Yep64s, y []Yep32s) YepStatus {
	status := YepStatus(
		C.yepCore_Min_IV64sV32s_IV64s(
			(*C.Yep64s)(unsafe.Pointer(&x[0])),
			(*C.Yep32s)(unsafe.Pointer(&y[0])),
			C.YepSize(len(x)),
		),
	)

	return status
}

// MinIV64uV32uIV64u Computes pairwise minima of corresponding elements in two unsigned 64-bit integer arrays and writes the result to the first array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MinIV64uV32uIV64u(x []Yep64u, y []Yep32u) YepStatus {
	status := YepStatus(
		C.yepCore_Min_IV64uV32u_IV64u(
			(*C.Yep64u)(unsafe.Pointer(&x[0])),
			(*C.Yep32u)(unsafe.Pointer(&y[0])),
			C.YepSize(len(x)),
		),
	)

	return status
}

// MinIV32fV32fIV32f Computes pairwise minima of corresponding elements in two single precision (32-bit) floating-point arrays and writes the result to the first array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MinIV32fV32fIV32f(x, y []Yep32f) YepStatus {
	status := YepStatus(
		C.yepCore_Min_IV32fV32f_IV32f(
			(*C.Yep32f)(unsafe.Pointer(&x[0])),
			(*C.Yep32f)(unsafe.Pointer(&y[0])),
			C.YepSize(len(x)),
		),
	)

	return status
}

// MinIV64fV64fIV64f Computes pairwise minima of corresponding elements in two double precision (64-bit) floating-point arrays and writes the result to the first array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MinIV64fV64fIV64f(x, y []Yep64f) YepStatus {
	status := YepStatus(
		C.yepCore_Min_IV64fV64f_IV64f(
			(*C.Yep64f)(unsafe.Pointer(&x[0])),
			(*C.Yep64f)(unsafe.Pointer(&y[0])),
			C.YepSize(len(x)),
		),
	)

	return status
}

// MinIV8sS8sIV8s Computes pairwise minima of signed 8-bit integer array elements and a constant and writes the result to the same array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MinIV8sS8sIV8s(x []Yep8s, y Yep8s) YepStatus {
	status := YepStatus(
		C.yepCore_Min_IV8sS8s_IV8s(
			(*C.Yep8s)(unsafe.Pointer(&x[0])),
			(C.Yep8s)(y),
			C.YepSize(len(x)),
		),
	)
	return status
}

// MinIV8uS8uIV8u Computes pairwise minima of unsigned 8-bit integer array elements and a constant and writes the result to the same array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MinIV8uS8uIV8u(x []Yep8u, y Yep8u) YepStatus {
	status := YepStatus(
		C.yepCore_Min_IV8uS8u_IV8u(
			(*C.Yep8u)(unsafe.Pointer(&x[0])),
			(C.Yep8u)(y),
			C.YepSize(len(x)),
		),
	)

	return status
}

// MinIV16sS16sIV16s Computes pairwise minima of signed 16-bit integer array elements and a constant and writes the result to the same array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MinIV16sS16sIV16s(x []Yep16s, y Yep16s) YepStatus {
	status := YepStatus(
		C.yepCore_Min_IV16sS16s_IV16s(
			(*C.Yep16s)(unsafe.Pointer(&x[0])),
			(C.Yep16s)(y),
			C.YepSize(len(x)),
		),
	)
	return status
}

// MinIV16uS16uIV16u Computes pairwise minima of unsigned 16-bit integer array elements and a constant and writes the result to the same array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MinIV16uS16uIV16u(x []Yep16u, y Yep16u) YepStatus {
	status := YepStatus(
		C.yepCore_Min_IV16uS16u_IV16u(
			(*C.Yep16u)(unsafe.Pointer(&x[0])),
			(C.Yep16u)(y),
			C.YepSize(len(x)),
		),
	)

	return status
}

// MinIV32sS32sIV32s Computes pairwise minima of signed 32-bit integer array elements and a constant and writes the result to the same array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MinIV32sS32sIV32s(x []Yep32s, y Yep32s) YepStatus {
	status := YepStatus(
		C.yepCore_Min_IV32sS32s_IV32s(
			(*C.Yep32s)(unsafe.Pointer(&x[0])),
			(C.Yep32s)(y),
			C.YepSize(len(x)),
		),
	)
	return status
}

// MinIV32uS32uIV32u Computes pairwise minima of unsigned 32-bit integer array elements and a constant and writes the result to the same array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MinIV32uS32uIV32u(x []Yep32u, y Yep32u) YepStatus {
	status := YepStatus(
		C.yepCore_Min_IV32uS32u_IV32u(
			(*C.Yep32u)(unsafe.Pointer(&x[0])),
			(C.Yep32u)(y),
			C.YepSize(len(x)),
		),
	)

	return status
}

// MinIV64sS32sIV64s Computes pairwise minima of signed 64-bit integer array elements and a constant and writes the result to the same array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MinIV64sS32sIV64s(x []Yep64s, y Yep32s) YepStatus {
	status := YepStatus(
		C.yepCore_Min_IV64sS32s_IV64s(
			(*C.Yep64s)(unsafe.Pointer(&x[0])),
			(C.Yep32s)(y),
			C.YepSize(len(x)),
		),
	)
	return status
}

// MinIV64uS32uIV64u Computes pairwise minima of unsigned 64-bit integer array elements and a constant and writes the result to the same array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MinIV64uS32uIV64u(x []Yep64u, y Yep32u) YepStatus {
	status := YepStatus(
		C.yepCore_Min_IV64uS32u_IV64u(
			(*C.Yep64u)(unsafe.Pointer(&x[0])),
			(C.Yep32u)(y),
			C.YepSize(len(x)),
		),
	)
	return status
}

// MinIV32fS32fIV32f Computes pairwise minima of single precision (32-bit) floating-point array elements and a constant and writes the result to the same array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MinIV32fS32fIV32f(x []Yep32f, y Yep32f) YepStatus {
	status := YepStatus(
		C.yepCore_Min_IV32fS32f_IV32f(
			(*C.Yep32f)(unsafe.Pointer(&x[0])),
			(C.Yep32f)(y),
			C.YepSize(len(x)),
		),
	)

	return status
}

// MinIV64fS64fIV64f Computes pairwise minima of double precision (64-bit) floating-point array elements and a constant and writes the result to the same array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MinIV64fS64fIV64f(x []Yep64f, y Yep64f) YepStatus {
	status := YepStatus(
		C.yepCore_Min_IV64fS64f_IV64f(
			(*C.Yep64f)(unsafe.Pointer(&x[0])),
			(C.Yep64f)(y),
			C.YepSize(len(x)),
		),
	)

	return status
}

// MaxV8sS8s Computes the maximum of signed 8-bit integer array elements.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MaxV8sS8s(v []Yep8s) (Yep8s, YepStatus) {
	var maximum Yep8s
	status := YepStatus(
		C.yepCore_Max_V8s_S8s(
			(*C.Yep8s)(unsafe.Pointer(&v[0])),
			(*C.Yep8s)(unsafe.Pointer(&maximum)),
			C.YepSize(len(v)),
		),
	)

	return maximum, status
}

// MaxV8uS8u Computes the maximum of unsigned 8-bit integer array elements.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MaxV8uS8u(v []Yep8u) (Yep8u, YepStatus) {
	var maximum Yep8u
	status := YepStatus(
		C.yepCore_Max_V8u_S8u(
			(*C.Yep8u)(unsafe.Pointer(&v[0])),
			(*C.Yep8u)(unsafe.Pointer(&maximum)),
			C.YepSize(len(v)),
		),
	)

	return maximum, status
}

// MaxV16sS16s Computes the maximum of signed 16-bit integer array elements.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MaxV16sS16s(v []Yep16s) (Yep16s, YepStatus) {
	var maximum Yep16s
	status := YepStatus(
		C.yepCore_Max_V16s_S16s(
			(*C.Yep16s)(unsafe.Pointer(&v[0])),
			(*C.Yep16s)(unsafe.Pointer(&maximum)),
			C.YepSize(len(v)),
		),
	)

	return maximum, status
}

// MaxV16uS16u Computes the maximum of unsigned 16-bit integer array elements.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MaxV16uS16u(v []Yep16u) (Yep16u, YepStatus) {
	var maximum Yep16u
	status := YepStatus(
		C.yepCore_Max_V16u_S16u(
			(*C.Yep16u)(unsafe.Pointer(&v[0])),
			(*C.Yep16u)(unsafe.Pointer(&maximum)),
			C.YepSize(len(v)),
		),
	)

	return maximum, status
}

// MaxV32sS32s Computes the maximum of signed 32-bit integer array elements.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MaxV32sS32s(v []Yep32s) (Yep32s, YepStatus) {
	var maximum Yep32s
	status := YepStatus(
		C.yepCore_Max_V32s_S32s(
			(*C.Yep32s)(unsafe.Pointer(&v[0])),
			(*C.Yep32s)(unsafe.Pointer(&maximum)),
			C.YepSize(len(v)),
		),
	)

	return maximum, status
}

// MaxV32uS32u Computes the maximum of unsigned 32-bit integer array elements.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MaxV32uS32u(v []Yep32u) (Yep32u, YepStatus) {
	var maximum Yep32u
	status := YepStatus(
		C.yepCore_Max_V32u_S32u(
			(*C.Yep32u)(unsafe.Pointer(&v[0])),
			(*C.Yep32u)(unsafe.Pointer(&maximum)),
			C.YepSize(len(v)),
		),
	)

	return maximum, status
}

// MaxV64sS64s Computes the maximum of signed 64-bit integer array elements.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MaxV64sS64s(v []Yep64s) (Yep64s, YepStatus) {
	var maximum Yep64s
	status := YepStatus(
		C.yepCore_Max_V64s_S64s(
			(*C.Yep64s)(unsafe.Pointer(&v[0])),
			(*C.Yep64s)(unsafe.Pointer(&maximum)),
			C.YepSize(len(v)),
		),
	)

	return maximum, status
}

// MaxV64uS64u Computes the maximum of unsigned 64-bit integer array elements.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MaxV64uS64u(v []Yep64u) (Yep64u, YepStatus) {
	var maximum Yep64u
	status := YepStatus(
		C.yepCore_Max_V64u_S64u(
			(*C.Yep64u)(unsafe.Pointer(&v[0])),
			(*C.Yep64u)(unsafe.Pointer(&maximum)),
			C.YepSize(len(v)),
		),
	)

	return maximum, status
}

// MaxV32fS32f Computes the maximum of single precision (32-bit) floating-point array elements.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MaxV32fS32f(v []Yep32f) (Yep32f, YepStatus) {
	var maximum Yep32f
	status := YepStatus(
		C.yepCore_Max_V32f_S32f(
			(*C.Yep32f)(unsafe.Pointer(&v[0])),
			(*C.Yep32f)(unsafe.Pointer(&maximum)),
			C.YepSize(len(v)),
		),
	)

	return maximum, status
}

// MaxV64fS64f Computes the maximum of single precision (64-bit) floating-point array elements.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MaxV64fS64f(v []Yep64f) (Yep64f, YepStatus) {
	var maximum Yep64f
	status := YepStatus(
		C.yepCore_Max_V64f_S64f(
			(*C.Yep64f)(unsafe.Pointer(&v[0])),
			(*C.Yep64f)(unsafe.Pointer(&maximum)),
			C.YepSize(len(v)),
		),
	)

	return maximum, status
}

// MaxV8sV8sV8s Computes pairwise maxima of corresponding elements in two signed 8-bit integer arrays.
func MaxV8sV8sV8s(x, y []Yep8s) ([]Yep8s, YepStatus) {
	maximum := make([]Yep8s, len(x))
	status := YepStatus(
		C.yepCore_Max_V8sV8s_V8s(
			(*C.Yep8s)(unsafe.Pointer(&x[0])),
			(*C.Yep8s)(unsafe.Pointer(&y[0])),
			(*C.Yep8s)(unsafe.Pointer(&maximum[0])),
			C.YepSize(len(x)),
		),
	)

	return maximum, status
}

// MaxV8uV8uV8u Computes pairwise maxima of corresponding elements in two unsigned 8-bit integer arrays.
func MaxV8uV8uV8u(x, y []Yep8u) ([]Yep8u, YepStatus) {
	maximum := make([]Yep8u, len(x))
	status := YepStatus(
		C.yepCore_Max_V8uV8u_V8u(
			(*C.Yep8u)(unsafe.Pointer(&x[0])),
			(*C.Yep8u)(unsafe.Pointer(&y[0])),
			(*C.Yep8u)(unsafe.Pointer(&maximum[0])),
			C.YepSize(len(x)),
		),
	)

	return maximum, status
}

// MaxV16sV16sV16s Computes pairwise maxima of corresponding elements in two signed 16-bit integer arrays.
func MaxV16sV16sV16s(x, y []Yep16s) ([]Yep16s, YepStatus) {
	maximum := make([]Yep16s, len(x))
	status := YepStatus(
		C.yepCore_Max_V16sV16s_V16s(
			(*C.Yep16s)(unsafe.Pointer(&x[0])),
			(*C.Yep16s)(unsafe.Pointer(&y[0])),
			(*C.Yep16s)(unsafe.Pointer(&maximum[0])),
			C.YepSize(len(x)),
		),
	)

	return maximum, status
}

// MaxV16uV16uV16u Computes pairwise maxima of corresponding elements in two unsigned 16-bit integer arrays.
func MaxV16uV16uV16u(x, y []Yep16u) ([]Yep16u, YepStatus) {
	maximum := make([]Yep16u, len(x))
	status := YepStatus(
		C.yepCore_Max_V16uV16u_V16u(
			(*C.Yep16u)(unsafe.Pointer(&x[0])),
			(*C.Yep16u)(unsafe.Pointer(&y[0])),
			(*C.Yep16u)(unsafe.Pointer(&maximum[0])),
			C.YepSize(len(x)),
		),
	)

	return maximum, status
}

// MaxV32sV32sV32s Computes pairwise maxima of corresponding elements in two signed 32-bit integer arrays.
func MaxV32sV32sV32s(x, y []Yep32s) ([]Yep32s, YepStatus) {
	maximum := make([]Yep32s, len(x))
	status := YepStatus(
		C.yepCore_Max_V32sV32s_V32s(
			(*C.Yep32s)(unsafe.Pointer(&x[0])),
			(*C.Yep32s)(unsafe.Pointer(&y[0])),
			(*C.Yep32s)(unsafe.Pointer(&maximum[0])),
			C.YepSize(len(x)),
		),
	)

	return maximum, status
}

// MaxV32uV32uV32u Computes pairwise maxima of corresponding elements in two unsigned 32-bit integer arrays.
func MaxV32uV32uV32u(x, y []Yep32u) ([]Yep32u, YepStatus) {
	maximum := make([]Yep32u, len(x))
	status := YepStatus(
		C.yepCore_Max_V32uV32u_V32u(
			(*C.Yep32u)(unsafe.Pointer(&x[0])),
			(*C.Yep32u)(unsafe.Pointer(&y[0])),
			(*C.Yep32u)(unsafe.Pointer(&maximum[0])),
			C.YepSize(len(x)),
		),
	)

	return maximum, status
}

// MaxV64sV32sV64s Computes pairwise maxima of corresponding elements in two signed 64-bit integer arrays.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MaxV64sV32sV64s(x []Yep64s, y []Yep32s) ([]Yep64s, YepStatus) {
	maximum := make([]Yep64s, len(x))
	status := YepStatus(
		C.yepCore_Max_V64sV32s_V64s(
			(*C.Yep64s)(unsafe.Pointer(&x[0])),
			(*C.Yep32s)(unsafe.Pointer(&y[0])),
			(*C.Yep64s)(unsafe.Pointer(&maximum[0])),
			C.YepSize(len(x)),
		),
	)

	return maximum, status
}

// MaxV64uV32uV64u Computes pairwise maxima of corresponding elements in two unsigned 64-bit integer arrays.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MaxV64uV32uV64u(x []Yep64u, y []Yep32u) ([]Yep64u, YepStatus) {
	maximum := make([]Yep64u, len(x))
	status := YepStatus(
		C.yepCore_Max_V64uV32u_V64u(
			(*C.Yep64u)(unsafe.Pointer(&x[0])),
			(*C.Yep32u)(unsafe.Pointer(&y[0])),
			(*C.Yep64u)(unsafe.Pointer(&maximum[0])),
			C.YepSize(len(x)),
		),
	)

	return maximum, status
}

// MaxV32fV32fV32f Computes pairwise maxima of corresponding elements in two single precision (32-bit) floating-point arrays.
func MaxV32fV32fV32f(x, y []Yep32f) ([]Yep32f, YepStatus) {
	maximum := make([]Yep32f, len(x))
	status := YepStatus(
		C.yepCore_Max_V32fV32f_V32f(
			(*C.Yep32f)(unsafe.Pointer(&x[0])),
			(*C.Yep32f)(unsafe.Pointer(&y[0])),
			(*C.Yep32f)(unsafe.Pointer(&maximum[0])),
			C.YepSize(len(x)),
		),
	)

	return maximum, status
}

// MaxV64fV64fV64f Computes pairwise maxima of corresponding elements in two double precision (64-bit) floating-point arrays.
func MaxV64fV64fV64f(x, y []Yep64f) ([]Yep64f, YepStatus) {
	maximum := make([]Yep64f, len(x))
	status := YepStatus(
		C.yepCore_Max_V64fV64f_V64f(
			(*C.Yep64f)(unsafe.Pointer(&x[0])),
			(*C.Yep64f)(unsafe.Pointer(&y[0])),
			(*C.Yep64f)(unsafe.Pointer(&maximum[0])),
			C.YepSize(len(x)),
		),
	)

	return maximum, status
}

// MaxV8sS8sV8s Computes pairwise maxima of signed 8-bit integer array elements and a constant.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MaxV8sS8sV8s(x []Yep8s, y Yep8s) ([]Yep8s, YepStatus) {
	maximum := make([]Yep8s, len(x))
	status := YepStatus(
		C.yepCore_Max_V8sS8s_V8s(
			(*C.Yep8s)(unsafe.Pointer(&x[0])),
			(C.Yep8s)(y),
			(*C.Yep8s)(unsafe.Pointer(&maximum[0])),
			C.YepSize(len(x)),
		),
	)
	return maximum, status
}

// MaxV8uS8uV8u Computes pairwise maxima of unsigned 8-bit integer array elements and a constant.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MaxV8uS8uV8u(x []Yep8u, y Yep8u) ([]Yep8u, YepStatus) {
	maximum := make([]Yep8u, len(x))
	status := YepStatus(
		C.yepCore_Max_V8uS8u_V8u(
			(*C.Yep8u)(unsafe.Pointer(&x[0])),
			(C.Yep8u)(y),
			(*C.Yep8u)(unsafe.Pointer(&maximum[0])),
			C.YepSize(len(x)),
		),
	)
	return maximum, status
}

// MaxV16sS16sV16s Computes pairwise maxima of signed 16-bit integer array elements and a constant.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MaxV16sS16sV16s(x []Yep16s, y Yep16s) ([]Yep16s, YepStatus) {
	maximum := make([]Yep16s, len(x))
	status := YepStatus(
		C.yepCore_Max_V16sS16s_V16s(
			(*C.Yep16s)(unsafe.Pointer(&x[0])),
			(C.Yep16s)(y),
			(*C.Yep16s)(unsafe.Pointer(&maximum[0])),
			C.YepSize(len(x)),
		),
	)
	return maximum, status
}

// MaxV16uS16uV16u Computes pairwise maxima of unsigned 16-bit integer array elements and a constant.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MaxV16uS16uV16u(x []Yep16u, y Yep16u) ([]Yep16u, YepStatus) {
	maximum := make([]Yep16u, len(x))
	status := YepStatus(
		C.yepCore_Max_V16uS16u_V16u(
			(*C.Yep16u)(unsafe.Pointer(&x[0])),
			(C.Yep16u)(y),
			(*C.Yep16u)(unsafe.Pointer(&maximum[0])),
			C.YepSize(len(x)),
		),
	)
	return maximum, status
}

// MaxV32sS32sV32s Computes pairwise maxima of signed 32-bit integer array elements and a constant.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MaxV32sS32sV32s(x []Yep32s, y Yep32s) ([]Yep32s, YepStatus) {
	maximum := make([]Yep32s, len(x))
	status := YepStatus(
		C.yepCore_Max_V32sS32s_V32s(
			(*C.Yep32s)(unsafe.Pointer(&x[0])),
			(C.Yep32s)(y),
			(*C.Yep32s)(unsafe.Pointer(&maximum[0])),
			C.YepSize(len(x)),
		),
	)
	return maximum, status
}

// MaxV32uS32uV32u Computes pairwise maxima of unsigned 32-bit integer array elements and a constant.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MaxV32uS32uV32u(x []Yep32u, y Yep32u) ([]Yep32u, YepStatus) {
	maximum := make([]Yep32u, len(x))
	status := YepStatus(
		C.yepCore_Max_V32uS32u_V32u(
			(*C.Yep32u)(unsafe.Pointer(&x[0])),
			(C.Yep32u)(y),
			(*C.Yep32u)(unsafe.Pointer(&maximum[0])),
			C.YepSize(len(x)),
		),
	)
	return maximum, status
}

// MaxV64sS32sV64s Computes pairwise maxima of signed 64-bit integer array elements and a constant.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MaxV64sS32sV64s(x []Yep64s, y Yep32s) ([]Yep64s, YepStatus) {
	maximum := make([]Yep64s, len(x))
	status := YepStatus(
		C.yepCore_Max_V64sS32s_V64s(
			(*C.Yep64s)(unsafe.Pointer(&x[0])),
			(C.Yep32s)(y),
			(*C.Yep64s)(unsafe.Pointer(&maximum[0])),
			C.YepSize(len(x)),
		),
	)
	return maximum, status
}

// MaxV64uS32uV64u Computes pairwise maxima of unsigned 64-bit integer array elements and a constant.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MaxV64uS32uV64u(x []Yep64u, y Yep32u) ([]Yep64u, YepStatus) {
	maximum := make([]Yep64u, len(x))
	status := YepStatus(
		C.yepCore_Max_V64uS32u_V64u(
			(*C.Yep64u)(unsafe.Pointer(&x[0])),
			(C.Yep32u)(y),
			(*C.Yep64u)(unsafe.Pointer(&maximum[0])),
			C.YepSize(len(x)),
		),
	)
	return maximum, status
}

// MaxV32fS32fV32f Computes pairwise maxima of single precision (32-bit) floating-point array elements and a constant.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MaxV32fS32fV32f(x []Yep32f, y Yep32f) ([]Yep32f, YepStatus) {
	maximum := make([]Yep32f, len(x))
	status := YepStatus(
		C.yepCore_Max_V32fS32f_V32f(
			(*C.Yep32f)(unsafe.Pointer(&x[0])),
			(C.Yep32f)(y),
			(*C.Yep32f)(unsafe.Pointer(&maximum[0])),
			C.YepSize(len(x)),
		),
	)
	return maximum, status
}

// MaxV64fS64fV64f Computes pairwise maxima of double precision (64-bit) floating-point array elements and a constant.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MaxV64fS64fV64f(x []Yep64f, y Yep64f) ([]Yep64f, YepStatus) {
	maximum := make([]Yep64f, len(x))
	status := YepStatus(
		C.yepCore_Max_V64fS64f_V64f(
			(*C.Yep64f)(unsafe.Pointer(&x[0])),
			(C.Yep64f)(y),
			(*C.Yep64f)(unsafe.Pointer(&maximum[0])),
			C.YepSize(len(x)),
		),
	)
	return maximum, status
}

// MaxIV8sV8sIV8s Computes pairwise maxima of corresponding elements in two signed 8-bit integer arrays and writes the result to the first array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MaxIV8sV8sIV8s(x, y []Yep8s) YepStatus {
	status := YepStatus(
		C.yepCore_Max_IV8sV8s_IV8s(
			(*C.Yep8s)(unsafe.Pointer(&x[0])),
			(*C.Yep8s)(unsafe.Pointer(&y[0])),
			C.YepSize(len(x)),
		),
	)

	return status
}

// MaxIV8uV8uIV8u Computes pairwise maxima of corresponding elements in two unsigned 8-bit integer arrays and writes the result to the first array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MaxIV8uV8uIV8u(x, y []Yep8u) YepStatus {
	status := YepStatus(
		C.yepCore_Max_IV8uV8u_IV8u(
			(*C.Yep8u)(unsafe.Pointer(&x[0])),
			(*C.Yep8u)(unsafe.Pointer(&y[0])),
			C.YepSize(len(x)),
		),
	)

	return status
}

// MaxIV16sV16sIV16s Computes pairwise maxima of corresponding elements in two signed 16-bit integer arrays and writes the result to the first array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MaxIV16sV16sIV16s(x, y []Yep16s) YepStatus {
	status := YepStatus(
		C.yepCore_Max_IV16sV16s_IV16s(
			(*C.Yep16s)(unsafe.Pointer(&x[0])),
			(*C.Yep16s)(unsafe.Pointer(&y[0])),
			C.YepSize(len(x)),
		),
	)

	return status
}

// MaxIV16uV16uIV16u Computes pairwise maxima of corresponding elements in two unsigned 16-bit integer arrays and writes the result to the first array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MaxIV16uV16uIV16u(x, y []Yep16u) YepStatus {
	status := YepStatus(
		C.yepCore_Max_IV16uV16u_IV16u(
			(*C.Yep16u)(unsafe.Pointer(&x[0])),
			(*C.Yep16u)(unsafe.Pointer(&y[0])),
			C.YepSize(len(x)),
		),
	)

	return status
}

// MaxIV32sV32sIV32s Computes pairwise maxima of corresponding elements in two signed 32-bit integer arrays and writes the result to the first array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MaxIV32sV32sIV32s(x, y []Yep32s) YepStatus {
	status := YepStatus(
		C.yepCore_Max_IV32sV32s_IV32s(
			(*C.Yep32s)(unsafe.Pointer(&x[0])),
			(*C.Yep32s)(unsafe.Pointer(&y[0])),
			C.YepSize(len(x)),
		),
	)

	return status
}

// MaxIV32uV32uIV32u Computes pairwise maxima of corresponding elements in two unsigned 32-bit integer arrays and writes the result to the first array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MaxIV32uV32uIV32u(x, y []Yep32u) YepStatus {
	status := YepStatus(
		C.yepCore_Max_IV32uV32u_IV32u(
			(*C.Yep32u)(unsafe.Pointer(&x[0])),
			(*C.Yep32u)(unsafe.Pointer(&y[0])),
			C.YepSize(len(x)),
		),
	)

	return status
}

// MaxIV64sV32sIV64s Computes pairwise maxima of corresponding elements in two signed 64-bit integer arrays and writes the result to the first array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MaxIV64sV32sIV64s(x []Yep64s, y []Yep32s) YepStatus {
	status := YepStatus(
		C.yepCore_Max_IV64sV32s_IV64s(
			(*C.Yep64s)(unsafe.Pointer(&x[0])),
			(*C.Yep32s)(unsafe.Pointer(&y[0])),
			C.YepSize(len(x)),
		),
	)

	return status
}

// MaxIV64uV32uIV64u Computes pairwise maxima of corresponding elements in two unsigned 64-bit integer arrays and writes the result to the first array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MaxIV64uV32uIV64u(x []Yep64u, y []Yep32u) YepStatus {
	status := YepStatus(
		C.yepCore_Max_IV64uV32u_IV64u(
			(*C.Yep64u)(unsafe.Pointer(&x[0])),
			(*C.Yep32u)(unsafe.Pointer(&y[0])),
			C.YepSize(len(x)),
		),
	)

	return status
}

// MaxIV32fV32fIV32f Computes pairwise maxima of corresponding elements in two single precision (32-bit) floating-point arrays and writes the result to the first array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MaxIV32fV32fIV32f(x, y []Yep32f) YepStatus {
	status := YepStatus(
		C.yepCore_Max_IV32fV32f_IV32f(
			(*C.Yep32f)(unsafe.Pointer(&x[0])),
			(*C.Yep32f)(unsafe.Pointer(&y[0])),
			C.YepSize(len(x)),
		),
	)

	return status
}

// MaxIV64fV64fIV64f Computes pairwise maxima of corresponding elements in two double precision (64-bit) floating-point arrays and writes the result to the first array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MaxIV64fV64fIV64f(x, y []Yep64f) YepStatus {
	status := YepStatus(
		C.yepCore_Max_IV64fV64f_IV64f(
			(*C.Yep64f)(unsafe.Pointer(&x[0])),
			(*C.Yep64f)(unsafe.Pointer(&y[0])),
			C.YepSize(len(x)),
		),
	)

	return status
}

// MaxIV8sS8sIV8s Computes pairwise maxima of signed 8-bit integer array elements and a constant and writes the result to the same array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MaxIV8sS8sIV8s(x []Yep8s, y Yep8s) YepStatus {
	status := YepStatus(
		C.yepCore_Max_IV8sS8s_IV8s(
			(*C.Yep8s)(unsafe.Pointer(&x[0])),
			(C.Yep8s)(y),
			C.YepSize(len(x)),
		),
	)
	return status
}

// MaxIV8uS8uIV8u Computes pairwise maxima of unsigned 8-bit integer array elements and a constant and writes the result to the same array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MaxIV8uS8uIV8u(x []Yep8u, y Yep8u) YepStatus {
	status := YepStatus(
		C.yepCore_Max_IV8uS8u_IV8u(
			(*C.Yep8u)(unsafe.Pointer(&x[0])),
			(C.Yep8u)(y),
			C.YepSize(len(x)),
		),
	)

	return status
}

// MaxIV16sS16sIV16s Computes pairwise maxima of signed 16-bit integer array elements and a constant and writes the result to the same array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MaxIV16sS16sIV16s(x []Yep16s, y Yep16s) YepStatus {
	status := YepStatus(
		C.yepCore_Max_IV16sS16s_IV16s(
			(*C.Yep16s)(unsafe.Pointer(&x[0])),
			(C.Yep16s)(y),
			C.YepSize(len(x)),
		),
	)
	return status
}

// MaxIV16uS16uIV16u Computes pairwise maxima of unsigned 16-bit integer array elements and a constant and writes the result to the same array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MaxIV16uS16uIV16u(x []Yep16u, y Yep16u) YepStatus {
	status := YepStatus(
		C.yepCore_Max_IV16uS16u_IV16u(
			(*C.Yep16u)(unsafe.Pointer(&x[0])),
			(C.Yep16u)(y),
			C.YepSize(len(x)),
		),
	)

	return status
}

// MaxIV32sS32sIV32s Computes pairwise maxima of signed 32-bit integer array elements and a constant and writes the result to the same array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MaxIV32sS32sIV32s(x []Yep32s, y Yep32s) YepStatus {
	status := YepStatus(
		C.yepCore_Max_IV32sS32s_IV32s(
			(*C.Yep32s)(unsafe.Pointer(&x[0])),
			(C.Yep32s)(y),
			C.YepSize(len(x)),
		),
	)
	return status
}

// MaxIV32uS32uIV32u Computes pairwise maxima of unsigned 32-bit integer array elements and a constant and writes the result to the same array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MaxIV32uS32uIV32u(x []Yep32u, y Yep32u) YepStatus {
	status := YepStatus(
		C.yepCore_Max_IV32uS32u_IV32u(
			(*C.Yep32u)(unsafe.Pointer(&x[0])),
			(C.Yep32u)(y),
			C.YepSize(len(x)),
		),
	)

	return status
}

// MaxIV64sS32sIV64s Computes pairwise maxima of signed 64-bit integer array elements and a constant and writes the result to the same array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MaxIV64sS32sIV64s(x []Yep64s, y Yep32s) YepStatus {
	status := YepStatus(
		C.yepCore_Max_IV64sS32s_IV64s(
			(*C.Yep64s)(unsafe.Pointer(&x[0])),
			(C.Yep32s)(y),
			C.YepSize(len(x)),
		),
	)
	return status
}

// MaxIV64uS32uIV64u Computes pairwise maxima of unsigned 64-bit integer array elements and a constant and writes the result to the same array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MaxIV64uS32uIV64u(x []Yep64u, y Yep32u) YepStatus {
	status := YepStatus(
		C.yepCore_Max_IV64uS32u_IV64u(
			(*C.Yep64u)(unsafe.Pointer(&x[0])),
			(C.Yep32u)(y),
			C.YepSize(len(x)),
		),
	)
	return status
}

// MaxIV32fS32fIV32f Computes pairwise maxima of single precision (32-bit) floating-point array elements and a constant and writes the result to the same array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MaxIV32fS32fIV32f(x []Yep32f, y Yep32f) YepStatus {
	status := YepStatus(
		C.yepCore_Max_IV32fS32f_IV32f(
			(*C.Yep32f)(unsafe.Pointer(&x[0])),
			(C.Yep32f)(y),
			C.YepSize(len(x)),
		),
	)

	return status
}

// MaxIV64fS64fIV64f Computes pairwise maxima of double precision (64-bit) floating-point array elements and a constant and writes the result to the same array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MaxIV64fS64fIV64f(x []Yep64f, y Yep64f) YepStatus {
	status := YepStatus(
		C.yepCore_Max_IV64fS64f_IV64f(
			(*C.Yep64f)(unsafe.Pointer(&x[0])),
			(C.Yep64f)(y),
			C.YepSize(len(x)),
		),
	)

	return status
}
