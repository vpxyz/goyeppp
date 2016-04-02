package gy

/*
#include <yepMath.h>

#cgo pkg-config: yeppp

*/
import "C"

import (
	"unsafe"
)

// Log Computes natural logarithm on an array of double precision (64-bit) floating-point elements.
func Log(x []Yep64f) ([]Yep64f, YepStatus) {
	y := make([]Yep64f, len(x))

	status := YepStatus(C.yepMath_Log_V64f_V64f((*C.Yep64f)(unsafe.Pointer(&x[0])), (*C.Yep64f)(unsafe.Pointer(&y[0])), C.YepSize(len(x))))

	return y, status
}

// Exp Computes base-e exponent on an array of double precision (64-bit) floating-point elements
func Exp(x []Yep64f) ([]Yep64f, YepStatus) {
	y := make([]Yep64f, len(x))

	status := YepStatus(C.yepMath_Exp_V64f_V64f((*C.Yep64f)(unsafe.Pointer(&x[0])), (*C.Yep64f)(unsafe.Pointer(&y[0])), C.YepSize(len(x))))

	return y, status
}

// Sin Computes sine on an array of double precision (64-bit) floating-point elements.
func Sin(x []Yep64f) ([]Yep64f, YepStatus) {
	y := make([]Yep64f, len(x))

	status := YepStatus(C.yepMath_Sin_V64f_V64f((*C.Yep64f)(unsafe.Pointer(&x[0])), (*C.Yep64f)(unsafe.Pointer(&y[0])), C.YepSize(len(x))))

	return y, status
}

// Cos Computes cosine on an array of double precision (64-bit) floating-point elements.
func Cos(x []Yep64f) ([]Yep64f, YepStatus) {
	y := make([]Yep64f, len(x))

	status := YepStatus(C.yepMath_Cos_V64f_V64f((*C.Yep64f)(unsafe.Pointer(&x[0])), (*C.Yep64f)(unsafe.Pointer(&y[0])), C.YepSize(len(x))))

	return y, status
}

// Tan Computes tangent on an array of double precision (64-bit) floating-point elements.
func Tan(x []Yep64f) ([]Yep64f, YepStatus) {
	y := make([]Yep64f, len(x))

	status := YepStatus(C.yepMath_Tan_V64f_V64f((*C.Yep64f)(unsafe.Pointer(&x[0])), (*C.Yep64f)(unsafe.Pointer(&y[0])), C.YepSize(len(x))))

	return y, status
}

// EvalPolynomial32 Evaluates polynomial with single precision (32-bit) floating-point coefficients on an array of single precision (32-bit) floating-point elements.
func EvalPolynomial32(coef, x []Yep32f) ([]Yep32f, YepStatus) {
	y := make([]Yep32f, len(x))

	status := YepStatus(C.yepMath_EvaluatePolynomial_V32fV32f_V32f(
		(*C.Yep32f)(unsafe.Pointer(&coef[0])),
		(*C.Yep32f)(unsafe.Pointer(&x[0])),
		(*C.Yep32f)(unsafe.Pointer(&y[0])),
		C.YepSize(len(coef)),
		C.YepSize(len(x)),
	))

	return y, status
}

// EvalPolynomial64 Evaluates polynomial with single precision (32-bit) floating-point coefficients on an array of single precision (32-bit) floating-point elements.
func EvalPolynomial64(coef, x []Yep64f) ([]Yep64f, YepStatus) {
	y := make([]Yep64f, len(x))

	status := YepStatus(C.yepMath_EvaluatePolynomial_V64fV64f_V64f(
		(*C.Yep64f)(unsafe.Pointer(&coef[0])),
		(*C.Yep64f)(unsafe.Pointer(&x[0])),
		(*C.Yep64f)(unsafe.Pointer(&y[0])),
		C.YepSize(len(coef)),
		C.YepSize(len(x)),
	))

	return y, status
}
