package gy

/*
#include <yepCore.h>

#cgo pkg-config: yeppp

*/
import "C"

import (
	"unsafe"
)

// NegateV8sV8s Negates elements in signed 8-bit integer array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func NegateV8sV8s(x []Yep8s) ([]Yep8s, YepStatus) {
	y := make([]Yep8s, len(x))
	status := YepStatus(
		C.yepCore_Negate_V8s_V8s(
			(*C.Yep8s)(unsafe.Pointer(&x[0])),
			(*C.Yep8s)(unsafe.Pointer(&y[0])),
			C.YepSize(len(x)),
		),
	)

	return y, status
}

// NegateV16sV16s Negates elements in signed 16-bit integer array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func NegateV16sV16s(x []Yep16s) ([]Yep16s, YepStatus) {
	y := make([]Yep16s, len(x))
	status := YepStatus(
		C.yepCore_Negate_V16s_V16s(
			(*C.Yep16s)(unsafe.Pointer(&x[0])),
			(*C.Yep16s)(unsafe.Pointer(&y[0])),
			C.YepSize(len(x)),
		),
	)

	return y, status
}

// NegateV32sV32s Negates elements in signed 32-bit integer array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func NegateV32sV32s(x []Yep32s) ([]Yep32s, YepStatus) {
	y := make([]Yep32s, len(x))
	status := YepStatus(
		C.yepCore_Negate_V32s_V32s(
			(*C.Yep32s)(unsafe.Pointer(&x[0])),
			(*C.Yep32s)(unsafe.Pointer(&y[0])),
			C.YepSize(len(x)),
		),
	)

	return y, status
}

// NegateV64sV64s Negates elements in signed 64-bit integer array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func NegateV64sV64s(x []Yep64s) ([]Yep64s, YepStatus) {
	y := make([]Yep64s, len(x))
	status := YepStatus(
		C.yepCore_Negate_V64s_V64s(
			(*C.Yep64s)(unsafe.Pointer(&x[0])),
			(*C.Yep64s)(unsafe.Pointer(&y[0])),
			C.YepSize(len(x)),
		),
	)

	return y, status
}

// NegateV32fV32f Negates elements in single precision (32-bit) floating-point array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func NegateV32fV32f(x []Yep32f) ([]Yep32f, YepStatus) {
	y := make([]Yep32f, len(x))
	status := YepStatus(
		C.yepCore_Negate_V32f_V32f(
			(*C.Yep32f)(unsafe.Pointer(&x[0])),
			(*C.Yep32f)(unsafe.Pointer(&y[0])),
			C.YepSize(len(x)),
		),
	)

	return y, status
}

// NegateV64fV64f Negates elements in single precision (64-bit) floating-point array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func NegateV64fV64f(x []Yep64f) ([]Yep64f, YepStatus) {
	y := make([]Yep64f, len(x))
	status := YepStatus(
		C.yepCore_Negate_V64f_V64f(
			(*C.Yep64f)(unsafe.Pointer(&x[0])),
			(*C.Yep64f)(unsafe.Pointer(&y[0])),
			C.YepSize(len(x)),
		),
	)

	return y, status
}

// NegateIV8sIV8s Negates elements in signed 8-bit integer array and writes the results to the same array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func NegateIV8sIV8s(x []Yep8s) YepStatus {
	status := YepStatus(
		C.yepCore_Negate_IV8s_IV8s(
			(*C.Yep8s)(unsafe.Pointer(&x[0])),
			C.YepSize(len(x)),
		),
	)
	return status
}

// NegateIV16sIV16s Negates elements in signed 16-bit integer array and writes the results to the same array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func NegateIV16sIV16s(x []Yep16s) YepStatus {
	status := YepStatus(
		C.yepCore_Negate_IV16s_IV16s(
			(*C.Yep16s)(unsafe.Pointer(&x[0])),
			C.YepSize(len(x)),
		),
	)
	return status
}

// NegateIV32sIV32s Negates elements in signed 32-bit integer array and writes the results to the same array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func NegateIV32sIV32s(x []Yep32s) YepStatus {
	status := YepStatus(
		C.yepCore_Negate_IV32s_IV32s(
			(*C.Yep32s)(unsafe.Pointer(&x[0])),
			C.YepSize(len(x)),
		),
	)
	return status
}

// NegateIV64sIV64s Negates elements in signed 64-bit integer array and writes the results to the same array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func NegateIV64sIV64s(x []Yep64s) YepStatus {
	status := YepStatus(
		C.yepCore_Negate_IV64s_IV64s(
			(*C.Yep64s)(unsafe.Pointer(&x[0])),
			C.YepSize(len(x)),
		),
	)
	return status
}

// NegateIV32fIV32f Negates elements in single precision (32-bit) floating-point array and writes the results to the same array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func NegateIV32fIV32f(x []Yep32f) YepStatus {
	status := YepStatus(
		C.yepCore_Negate_IV32f_IV32f(
			(*C.Yep32f)(unsafe.Pointer(&x[0])),
			C.YepSize(len(x)),
		),
	)
	return status
}

// NegateIV64fIV64f Negates elements in double precision (64-bit) floating-point array and writes the results to the same array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func NegateIV64fIV64f(x []Yep64f) YepStatus {
	status := YepStatus(
		C.yepCore_Negate_IV64f_IV64f(
			(*C.Yep64f)(unsafe.Pointer(&x[0])),
			C.YepSize(len(x)),
		),
	)
	return status
}

// SumV32fS32f Computes the sum of single precision (32-bit) floating-point array elements.
func SumV32fS32f(v []Yep32f) (Yep32f, YepStatus) {
	var sum Yep32f
	status := YepStatus(
		C.yepCore_Sum_V32f_S32f(
			(*C.Yep32f)(unsafe.Pointer(&v[0])),
			(*C.Yep32f)(unsafe.Pointer(&sum)),
			C.YepSize(len(v)),
		),
	)

	return sum, status
}

// SumV64fS64f Computes the sum of single precision (64-bit) floating-point array elements.
func SumV64fS64f(v []Yep64f) (Yep64f, YepStatus) {
	var sum Yep64f
	status := YepStatus(
		C.yepCore_Sum_V64f_S64f(
			(*C.Yep64f)(unsafe.Pointer(&v[0])),
			(*C.Yep64f)(unsafe.Pointer(&sum)),
			C.YepSize(len(v)),
		),
	)

	return sum, status
}

// SumAbsV32fS32f Computes the sum of absolute values of single precision (32-bit) floating-point array elements.
func SumAbsV32fS32f(v []Yep32f) (Yep32f, YepStatus) {
	var sumAbs Yep32f
	status := YepStatus(
		C.yepCore_SumAbs_V32f_S32f(
			(*C.Yep32f)(unsafe.Pointer(&v[0])),
			(*C.Yep32f)(unsafe.Pointer(&sumAbs)),
			C.YepSize(len(v)),
		),
	)

	return sumAbs, status
}

// SumAbsV64fS64f Computes the sum of absolute values of double precision (64-bit) floating-point array elements.
func SumAbsV64fS64f(v []Yep64f) (Yep64f, YepStatus) {
	var sumAbs Yep64f
	status := YepStatus(
		C.yepCore_SumAbs_V64f_S64f(
			(*C.Yep64f)(unsafe.Pointer(&v[0])),
			(*C.Yep64f)(unsafe.Pointer(&sumAbs)),
			C.YepSize(len(v)),
		),
	)

	return sumAbs, status
}

// SumSquaresV32fS32f Computes the sum of squares of single precision (32-bit) floating-point array elements.
func SumSquaresV32fS32f(v []Yep32f) (Yep32f, YepStatus) {
	var sumSquares Yep32f
	status := YepStatus(
		C.yepCore_SumSquares_V32f_S32f(
			(*C.Yep32f)(unsafe.Pointer(&v[0])),
			(*C.Yep32f)(unsafe.Pointer(&sumSquares)),
			C.YepSize(len(v)),
		),
	)

	return sumSquares, status
}

// SumSquaresV64fS64f Computes the sum of squares of double precision (64-bit) floating-point array elements.
func SumSquaresV64fS64f(v []Yep64f) (Yep64f, YepStatus) {
	var sumSquares Yep64f
	status := YepStatus(
		C.yepCore_SumAbs_V64f_S64f(
			(*C.Yep64f)(unsafe.Pointer(&v[0])),
			(*C.Yep64f)(unsafe.Pointer(&sumSquares)),
			C.YepSize(len(v)),
		),
	)

	return sumSquares, status
}

// DotProductV32fV32fS32f Computes the dot product of two vectors of single precision (32-bit) floating-point elements. x and y must have same lenght
func DotProductV32fV32fS32f(x, y []Yep32f) (Yep32f, YepStatus) {
	if len(x) != len(y) {
		return 0, YepStatusInvalidArgument
	}
	var dp Yep32f
	status := YepStatus(
		C.yepCore_DotProduct_V32fV32f_S32f(
			(*C.Yep32f)(unsafe.Pointer(&x[0])),
			(*C.Yep32f)(unsafe.Pointer(&y[0])),
			(*C.Yep32f)(unsafe.Pointer(&dp)),
			C.YepSize(len(x)),
		),
	)

	return dp, status
}

// DotProductV64fV64fS64f Computes the dot product of two vectors of double precision (64-bit) floating-point elements. x and y must have same lenght
func DotProductV64fV64fS64f(x, y []Yep64f) (Yep64f, YepStatus) {
	if len(x) != len(y) {
		return 0, YepStatusInvalidArgument
	}
	var dp Yep64f
	status := YepStatus(
		C.yepCore_DotProduct_V64fV64f_S64f(
			(*C.Yep64f)(unsafe.Pointer(&x[0])),
			(*C.Yep64f)(unsafe.Pointer(&y[0])),
			(*C.Yep64f)(unsafe.Pointer(&dp)),
			C.YepSize(len(x)),
		),
	)

	return dp, status
}
