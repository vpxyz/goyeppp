package goyeppp

/*
#include <yepCore.h>

#cgo pkg-config: yeppp

*/
import "C"

import (
	"unsafe"
)

// MultiplyV8sV8sV8s Multiples corresponding elements in two signed 8-bit integer arrays, producing an array of signed 8-bit integer elements.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MultiplyV8sV8sV8s(x, y []Yep8s) ([]Yep8s, YepStatus) {
	product := make([]Yep8s, len(x))
	status := YepStatus(
		C.yepCore_Multiply_V8sV8s_V8s(
			(*C.Yep8s)(unsafe.Pointer(&x[0])),
			(*C.Yep8s)(unsafe.Pointer(&y[0])),
			(*C.Yep8s)(unsafe.Pointer(&product[0])),
			C.YepSize(len(product)),
		),
	)

	return product, status
}

// MultiplyV8sV8sV16s Multiples corresponding elements in two signed 8-bit integer arrays, producing an array of signed 16-bit integer elements.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MultiplyV8sV8sV16s(x, y []Yep8s) ([]Yep16s, YepStatus) {
	product := make([]Yep16s, len(x))
	status := YepStatus(
		C.yepCore_Multiply_V8sV8s_V16s(
			(*C.Yep8s)(unsafe.Pointer(&x[0])),
			(*C.Yep8s)(unsafe.Pointer(&y[0])),
			(*C.Yep16s)(unsafe.Pointer(&product[0])),
			C.YepSize(len(product)),
		),
	)

	return product, status
}

// MultiplyV8uV8uV16u Multiples corresponding elements in two unsigned 8-bit integer arrays, producing an array of unsigned 16-bit integer elements.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MultiplyV8uV8uV16u(x, y []Yep8u) ([]Yep16u, YepStatus) {
	product := make([]Yep16u, len(x))
	status := YepStatus(
		C.yepCore_Multiply_V8uV8u_V16u(
			(*C.Yep8u)(unsafe.Pointer(&x[0])),
			(*C.Yep8u)(unsafe.Pointer(&y[0])),
			(*C.Yep16u)(unsafe.Pointer(&product[0])),
			C.YepSize(len(product)),
		),
	)

	return product, status
}

// MultiplyV16sV16sV16s Multiples corresponding elements in two signed 16-bit integer arrays, producing an array of signed 16-bit integer elements.
func MultiplyV16sV16sV16s(x, y []Yep16s) ([]Yep16s, YepStatus) {
	product := make([]Yep16s, len(x))
	status := YepStatus(
		C.yepCore_Multiply_V16sV16s_V16s(
			(*C.Yep16s)(unsafe.Pointer(&x[0])),
			(*C.Yep16s)(unsafe.Pointer(&y[0])),
			(*C.Yep16s)(unsafe.Pointer(&product[0])),
			C.YepSize(len(product)),
		),
	)

	return product, status
}

// MultiplyV16sV16sV32s Multiples corresponding elements in two signed 16-bit integer arrays, producing an array of signed 32-bit integer elements.
func MultiplyV16sV16sV32s(x, y []Yep16s) ([]Yep32s, YepStatus) {
	product := make([]Yep32s, len(x))
	status := YepStatus(
		C.yepCore_Multiply_V16sV16s_V32s(
			(*C.Yep16s)(unsafe.Pointer(&x[0])),
			(*C.Yep16s)(unsafe.Pointer(&y[0])),
			(*C.Yep32s)(unsafe.Pointer(&product[0])),
			C.YepSize(len(product)),
		),
	)

	return product, status
}

// MultiplyV16uV16uV32u Multiples corresponding elements in two unsigned 16-bit integer arrays, producing an array of unsigned 32-bit integer elements.
func MultiplyV16uV16uV32u(x, y []Yep16u) ([]Yep32u, YepStatus) {
	product := make([]Yep32u, len(x))
	status := YepStatus(
		C.yepCore_Multiply_V16uV16u_V32u(
			(*C.Yep16u)(unsafe.Pointer(&x[0])),
			(*C.Yep16u)(unsafe.Pointer(&y[0])),
			(*C.Yep32u)(unsafe.Pointer(&product[0])),
			C.YepSize(len(product)),
		),
	)

	return product, status
}

// MultiplyV32sV32sV32s Multiples corresponding elements in two signed 32-bit integer arrays, producing an array of signed 32-bit integer elements.
func MultiplyV32sV32sV32s(x, y []Yep32s) ([]Yep32s, YepStatus) {
	product := make([]Yep32s, len(x))
	status := YepStatus(
		C.yepCore_Multiply_V32sV32s_V32s(
			(*C.Yep32s)(unsafe.Pointer(&x[0])),
			(*C.Yep32s)(unsafe.Pointer(&y[0])),
			(*C.Yep32s)(unsafe.Pointer(&product[0])),
			C.YepSize(len(product)),
		),
	)

	return product, status
}

// MultiplyV32sV32sV64s Multiples corresponding elements in two signed 32-bit integer arrays, producing an array of signed 64-bit integer elements.
func MultiplyV32sV32sV64s(x, y []Yep32s) ([]Yep64s, YepStatus) {
	product := make([]Yep64s, len(x))
	status := YepStatus(
		C.yepCore_Multiply_V32sV32s_V64s(
			(*C.Yep32s)(unsafe.Pointer(&x[0])),
			(*C.Yep32s)(unsafe.Pointer(&y[0])),
			(*C.Yep64s)(unsafe.Pointer(&product[0])),
			C.YepSize(len(product)),
		),
	)

	return product, status
}

// MultiplyV32uV32uV64u Multiples corresponding elements in two unsigned 32-bit integer arrays, producing an array of unsigned 64-bit integer elements.
func MultiplyV32uV32uV64u(x, y []Yep32u) ([]Yep64u, YepStatus) {
	product := make([]Yep64u, len(x))
	status := YepStatus(
		C.yepCore_Multiply_V32uV32u_V64u(
			(*C.Yep32u)(unsafe.Pointer(&x[0])),
			(*C.Yep32u)(unsafe.Pointer(&y[0])),
			(*C.Yep64u)(unsafe.Pointer(&product[0])),
			C.YepSize(len(product)),
		),
	)

	return product, status
}

// MultiplyV64sV64sV64s Multiples corresponding elements in two signed 64-bit integer arrays, producing an array of signed 64-bit integer elements.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MultiplyV64sV64sV64s(x, y []Yep64s) ([]Yep64s, YepStatus) {
	product := make([]Yep64s, len(x))
	status := YepStatus(
		C.yepCore_Multiply_V64sV64s_V64s(
			(*C.Yep64s)(unsafe.Pointer(&x[0])),
			(*C.Yep64s)(unsafe.Pointer(&y[0])),
			(*C.Yep64s)(unsafe.Pointer(&product[0])),
			C.YepSize(len(product)),
		),
	)

	return product, status
}

// MultiplyV32fV32fV32f Multiples corresponding elements in two single precision (32-bit) floating-point arrays, producing an array of single precision (32-bit) floating-point elements.
func MultiplyV32fV32fV32f(x, y []Yep32f) ([]Yep32f, YepStatus) {
	product := make([]Yep32f, len(x))
	status := YepStatus(
		C.yepCore_Multiply_V32fV32f_V32f(
			(*C.Yep32f)(unsafe.Pointer(&x[0])),
			(*C.Yep32f)(unsafe.Pointer(&y[0])),
			(*C.Yep32f)(unsafe.Pointer(&product[0])),
			C.YepSize(len(product)),
		),
	)

	return product, status
}

// MultiplyV64fV64fV64f Multiples corresponding elements in two double precision (64-bit) floating-point arrays, producing an array of double precision (64-bit) floating-point elements.
func MultiplyV64fV64fV64f(x, y []Yep64f) ([]Yep64f, YepStatus) {
	product := make([]Yep64f, len(x))
	status := YepStatus(
		C.yepCore_Multiply_V64fV64f_V64f(
			(*C.Yep64f)(unsafe.Pointer(&x[0])),
			(*C.Yep64f)(unsafe.Pointer(&y[0])),
			(*C.Yep64f)(unsafe.Pointer(&product[0])),
			C.YepSize(len(product)),
		),
	)

	return product, status
}

// MultiplyV8sS8sV8s Multiplies signed 8-bit integer array elements by a constant. Produces an array of signed 8-bit integer elements.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MultiplyV8sS8sV8s(x []Yep8s, y Yep8s) ([]Yep8s, YepStatus) {
	product := make([]Yep8s, len(x))
	status := YepStatus(
		C.yepCore_Multiply_V8sS8s_V8s(
			(*C.Yep8s)(unsafe.Pointer(&x[0])),
			(C.Yep8s)(y),
			(*C.Yep8s)(unsafe.Pointer(&product[0])),
			C.YepSize(len(product)),
		),
	)

	return product, status
}

// MultiplyV8sS8sV16s Multiplies signed 8-bit integer array elements by a constant. Produces an array of signed 16-bit integer elements.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MultiplyV8sS8sV16s(x []Yep8s, y Yep8s) ([]Yep16s, YepStatus) {
	product := make([]Yep16s, len(x))
	status := YepStatus(
		C.yepCore_Multiply_V8sS8s_V16s(
			(*C.Yep8s)(unsafe.Pointer(&x[0])),
			(C.Yep8s)(y),
			(*C.Yep16s)(unsafe.Pointer(&product[0])),
			C.YepSize(len(product)),
		),
	)

	return product, status
}

// MultiplyV8uS8uV16u Multiplies unsigned 8-bit integer array elements by a constant. Produces an array of unsigned 16-bit integer elements.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MultiplyV8uS8uV16u(x []Yep8u, y Yep8u) ([]Yep16u, YepStatus) {
	product := make([]Yep16u, len(x))
	status := YepStatus(
		C.yepCore_Multiply_V8uS8u_V16u(
			(*C.Yep8u)(unsafe.Pointer(&x[0])),
			(C.Yep8u)(y),
			(*C.Yep16u)(unsafe.Pointer(&product[0])),
			C.YepSize(len(product)),
		),
	)

	return product, status
}

// MultiplyV16sS16sV16s Multiplies signed 16-bit integer array elements by a constant. Produces an array of signed 16-bit integer elements.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MultiplyV16sS16sV16s(x []Yep16s, y Yep16s) ([]Yep16s, YepStatus) {
	product := make([]Yep16s, len(x))
	status := YepStatus(
		C.yepCore_Multiply_V16sS16s_V16s(
			(*C.Yep16s)(unsafe.Pointer(&x[0])),
			(C.Yep16s)(y),
			(*C.Yep16s)(unsafe.Pointer(&product[0])),
			C.YepSize(len(product)),
		),
	)

	return product, status
}

// MultiplyV16sS16sV32s Multiplies signed 16-bit integer array elements by a constant. Produces an array of signed 32-bit integer elements.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MultiplyV16sS16sV32s(x []Yep16s, y Yep16s) ([]Yep32s, YepStatus) {
	product := make([]Yep32s, len(x))
	status := YepStatus(
		C.yepCore_Multiply_V16sS16s_V32s(
			(*C.Yep16s)(unsafe.Pointer(&x[0])),
			(C.Yep16s)(y),
			(*C.Yep32s)(unsafe.Pointer(&product[0])),
			C.YepSize(len(product)),
		),
	)

	return product, status
}

// MultiplyV16uS16uV32u Multiplies unsigned 16-bit integer array elements by a constant. Produces an array of unsigned 32-bit integer elements.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MultiplyV16uS16uV32u(x []Yep16u, y Yep16u) ([]Yep32u, YepStatus) {
	product := make([]Yep32u, len(x))
	status := YepStatus(
		C.yepCore_Multiply_V16uS16u_V32u(
			(*C.Yep16u)(unsafe.Pointer(&x[0])),
			(C.Yep16u)(y),
			(*C.Yep32u)(unsafe.Pointer(&product[0])),
			C.YepSize(len(product)),
		),
	)

	return product, status
}

// MultiplyV32sS32sV32s Multiplies signed 32-bit integer array elements by a constant. Produces an array of signed 32-bit integer elements.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MultiplyV32sS32sV32s(x []Yep32s, y Yep32s) ([]Yep32s, YepStatus) {
	product := make([]Yep32s, len(x))
	status := YepStatus(
		C.yepCore_Multiply_V32sS32s_V32s(
			(*C.Yep32s)(unsafe.Pointer(&x[0])),
			(C.Yep32s)(y),
			(*C.Yep32s)(unsafe.Pointer(&product[0])),
			C.YepSize(len(product)),
		),
	)

	return product, status
}

// MultiplyV32sS32sV64s Multiplies signed 32-bit integer array elements by a constant. Produces an array of signed 64-bit integer elements.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MultiplyV32sS32sV64s(x []Yep32s, y Yep32s) ([]Yep64s, YepStatus) {
	product := make([]Yep64s, len(x))
	status := YepStatus(
		C.yepCore_Multiply_V32sS32s_V64s(
			(*C.Yep32s)(unsafe.Pointer(&x[0])),
			(C.Yep32s)(y),
			(*C.Yep64s)(unsafe.Pointer(&product[0])),
			C.YepSize(len(product)),
		),
	)

	return product, status
}

// MultiplyV32uS32uV64u Multiplies unsigned 32-bit integer array elements by a constant. Produces an array of unsigned 64-bit integer elements.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MultiplyV32uS32uV64u(x []Yep32u, y Yep32u) ([]Yep64u, YepStatus) {
	product := make([]Yep64u, len(x))
	status := YepStatus(
		C.yepCore_Multiply_V32uS32u_V64u(
			(*C.Yep32u)(unsafe.Pointer(&x[0])),
			(C.Yep32u)(y),
			(*C.Yep64u)(unsafe.Pointer(&product[0])),
			C.YepSize(len(product)),
		),
	)

	return product, status
}

// MultiplyV64sS64sV64s Multiplies signed 64-bit integer array elements by a constant. Produces an array of signed 64-bit integer elements.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MultiplyV64sS64sV64s(x []Yep64s, y Yep64s) ([]Yep64s, YepStatus) {
	product := make([]Yep64s, len(x))
	status := YepStatus(
		C.yepCore_Multiply_V64sS64s_V64s(
			(*C.Yep64s)(unsafe.Pointer(&x[0])),
			(C.Yep64s)(y),
			(*C.Yep64s)(unsafe.Pointer(&product[0])),
			C.YepSize(len(product)),
		),
	)

	return product, status
}

// MultiplyV64fS64fV64f Multiplies double precision (64-bit) floating-point array elements by a constant. Produces an array of double precision (64-bit) floating-point elements.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MultiplyV64fS64fV64f(x []Yep64f, y Yep64f) ([]Yep64f, YepStatus) {
	product := make([]Yep64f, len(x))
	status := YepStatus(
		C.yepCore_Multiply_V64fS64f_V64f(
			(*C.Yep64f)(unsafe.Pointer(&x[0])),
			(C.Yep64f)(y),
			(*C.Yep64f)(unsafe.Pointer(&product[0])),
			C.YepSize(len(product)),
		),
	)

	return product, status
}

// MultiplyIV8sV8sIV8s Multiplies corresponding elements in two signed 8-bit integer arrays and writes the result to the first array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MultiplyIV8sV8sIV8s(x, y []Yep8s) YepStatus {
	status := YepStatus(
		C.yepCore_Multiply_IV8sV8s_IV8s(
			(*C.Yep8s)(unsafe.Pointer(&x[0])),
			(*C.Yep8s)(unsafe.Pointer(&y[0])),
			C.YepSize(len(x)),
		),
	)
	return status
}

// MultiplyIV16sV16sIV16s Multiplies corresponding elements in two signed 16-bit integer arrays and writes the result to the first array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MultiplyIV16sV16sIV16s(x, y []Yep16s) YepStatus {
	status := YepStatus(
		C.yepCore_Multiply_IV16sV16s_IV16s(
			(*C.Yep16s)(unsafe.Pointer(&x[0])),
			(*C.Yep16s)(unsafe.Pointer(&y[0])),
			C.YepSize(len(x)),
		),
	)
	return status
}

// MultiplyIV32sV32sIV32s Multiplies corresponding elements in two signed 32-bit integer arrays and writes the result to the first array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MultiplyIV32sV32sIV32s(x, y []Yep32s) YepStatus {
	status := YepStatus(
		C.yepCore_Multiply_IV32sV32s_IV32s(
			(*C.Yep32s)(unsafe.Pointer(&x[0])),
			(*C.Yep32s)(unsafe.Pointer(&y[0])),
			C.YepSize(len(x)),
		),
	)
	return status
}

// MultiplyIV64sV64sIV64s Multiplies corresponding elements in two signed 64-bit integer arrays and writes the result to the first array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MultiplyIV64sV64sIV64s(x, y []Yep64s) YepStatus {
	status := YepStatus(
		C.yepCore_Multiply_IV64sV64s_IV64s(
			(*C.Yep64s)(unsafe.Pointer(&x[0])),
			(*C.Yep64s)(unsafe.Pointer(&y[0])),
			C.YepSize(len(x)),
		),
	)
	return status
}

// MultiplyIV32fV32fIV32f Multiplies corresponding elements in two single precision (32-bit) floating-point arrays and writes the result to the first array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MultiplyIV32fV32fIV32f(x, y []Yep32f) YepStatus {
	status := YepStatus(
		C.yepCore_Multiply_IV32fV32f_IV32f(
			(*C.Yep32f)(unsafe.Pointer(&x[0])),
			(*C.Yep32f)(unsafe.Pointer(&y[0])),
			C.YepSize(len(x)),
		),
	)
	return status
}

// MultiplyIV64fV64fIV64f Multiplies corresponding elements in two double precision (64-bit) floating-point arrays and writes the result to the first array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MultiplyIV64fV64fIV64f(x, y []Yep64f) YepStatus {
	status := YepStatus(
		C.yepCore_Multiply_IV64fV64f_IV64f(
			(*C.Yep64f)(unsafe.Pointer(&x[0])),
			(*C.Yep64f)(unsafe.Pointer(&y[0])),
			C.YepSize(len(x)),
		),
	)
	return status
}

// MultiplyIV8sS8sIV8s Multiplies signed 8-bit integer array elements by a constant and writes the result to the same array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MultiplyIV8sS8sIV8s(x []Yep8s, y Yep8s) YepStatus {
	status := YepStatus(
		C.yepCore_Multiply_IV8sS8s_IV8s(
			(*C.Yep8s)(unsafe.Pointer(&x[0])),
			(C.Yep8s)(y),
			C.YepSize(len(x)),
		),
	)
	return status
}

// MultiplyIV16sS16sIV16s Multiplies signed 16-bit integer array elements by a constant and writes the result to the same array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MultiplyIV16sS16sIV16s(x []Yep16s, y Yep16s) YepStatus {
	status := YepStatus(
		C.yepCore_Multiply_IV16sS16s_IV16s(
			(*C.Yep16s)(unsafe.Pointer(&x[0])),
			(C.Yep16s)(y),
			C.YepSize(len(x)),
		),
	)
	return status
}

// MultiplyIV32sS32sIV32s Multiplies signed 32-bit integer array elements by a constant and writes the result to the same array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MultiplyIV32sS32sIV32s(x []Yep32s, y Yep32s) YepStatus {
	status := YepStatus(
		C.yepCore_Multiply_IV32sS32s_IV32s(
			(*C.Yep32s)(unsafe.Pointer(&x[0])),
			(C.Yep32s)(y),
			C.YepSize(len(x)),
		),
	)
	return status
}

// MultiplyIV64sS64sIV64s Multiplies signed 64-bit integer array elements by a constant and writes the result to the same array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MultiplyIV64sS64sIV64s(x []Yep64s, y Yep64s) YepStatus {
	status := YepStatus(
		C.yepCore_Multiply_IV64sS64s_IV64s(
			(*C.Yep64s)(unsafe.Pointer(&x[0])),
			(C.Yep64s)(y),
			C.YepSize(len(x)),
		),
	)
	return status
}

// MultiplyIV32fS32fIV32f Multiplies single precision (32-bit) floating-point array elements by a constant and writes the result to the same array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MultiplyIV32fS32fIV32f(x []Yep32f, y Yep32f) YepStatus {
	status := YepStatus(
		C.yepCore_Multiply_IV32fS32f_IV32f(
			(*C.Yep32f)(unsafe.Pointer(&x[0])),
			(C.Yep32f)(y),
			C.YepSize(len(x)),
		),
	)
	return status
}

// MultiplyIV64fS64fIV64f Multiplies double precision (64-bit) floating-point array elements by a constant and writes the result to the same array.
// Warning This version of Yeppp! does not include optimized implementations for this function
func MultiplyIV64fS64fIV64f(x []Yep64f, y Yep64f) YepStatus {
	status := YepStatus(
		C.yepCore_Multiply_IV64fS64f_IV64f(
			(*C.Yep64f)(unsafe.Pointer(&x[0])),
			(C.Yep64f)(y),
			C.YepSize(len(x)),
		),
	)
	return status
}
