package gy

/*
#include <yepRandom.h>

#cgo pkg-config: yeppp

*/
import "C"

import (
	"unsafe"
)

// WELL1024a random number generator state
type WELL1024a struct {
	State [32]Yep32u
	Index Yep32u
}

// InitWELL2014a Initializes the state of WELL1024a pseudo-random number generator with the default seed.
// Default seed is fixed, and never changes during execution of a program.
func (s *WELL1024a) InitWELL2014a() YepStatus {

	status := YepStatus(C.yepRandom_WELL1024a_Init((*C.struct_YepRandom_WELL1024a)(unsafe.Pointer(s))))

	return status
}

// InitWELL2014aSeed  Initializes the state of WELL1024a pseudo-random number generator with the specified seed.
func (s *WELL1024a) InitWELL2014aSeed(seed [32]Yep32u) YepStatus {
	status := YepStatus(
		C.yepRandom_WELL1024a_Init_V32u(
			(*C.struct_YepRandom_WELL1024a)(unsafe.Pointer(s)),
			(*C.Yep32u)(unsafe.Pointer(&seed[0])),
		),
	)
	return status

}

// GenerateDiscreteUniform8u Generates random 8-bit samples with WELL1024a pseudo-random number generator.
// Each 8-bit number is generated with the same probability.
func (s *WELL1024a) GenerateDiscreteUniform8u(samples []Yep8u) YepStatus {
	status := YepStatus(
		C.yepRandom_WELL1024a_GenerateDiscreteUniform__V8u(
			(*C.struct_YepRandom_WELL1024a)(unsafe.Pointer(s)),
			(*C.Yep8u)(unsafe.Pointer(&samples[0])),
			(C.YepSize)(len(samples)),
		),
	)

	return status
}

// GenerateDiscreteUniform16u Generates random 16-bit samples with WELL1024a pseudo-random number generator.
// Each 16-bit number is generated with the same probability.
func (s *WELL1024a) GenerateDiscreteUniform16u(samples []Yep16u) YepStatus {
	status := YepStatus(
		C.yepRandom_WELL1024a_GenerateDiscreteUniform__V16u(
			(*C.struct_YepRandom_WELL1024a)(unsafe.Pointer(s)),
			(*C.Yep16u)(unsafe.Pointer(&samples[0])),
			(C.YepSize)(len(samples)),
		),
	)

	return status
}

// GenerateDiscreteUniform32u Generates random 32-bit samples with WELL1024a pseudo-random number generator.
// Each 32-bit number is generated with the same probability.
func (s *WELL1024a) GenerateDiscreteUniform32u(samples []Yep32u) YepStatus {
	status := YepStatus(
		C.yepRandom_WELL1024a_GenerateDiscreteUniform__V32u(
			(*C.struct_YepRandom_WELL1024a)(unsafe.Pointer(s)),
			(*C.Yep32u)(unsafe.Pointer(&samples[0])),
			(C.YepSize)(len(samples)),
		),
	)

	return status
}

// GenerateDiscreteUniform64u Generates random 64-bit samples with WELL1024a pseudo-random number generator.
// Each 64-bit number is generated with the same probability.
func (s *WELL1024a) GenerateDiscreteUniform64u(samples []Yep64u) YepStatus {
	status := YepStatus(
		C.yepRandom_WELL1024a_GenerateDiscreteUniform__V64u(
			(*C.struct_YepRandom_WELL1024a)(unsafe.Pointer(s)),
			(*C.Yep64u)(unsafe.Pointer(&samples[0])),
			(C.YepSize)(len(samples)),
		),
	)

	return status
}

// GenerateDiscreteUniformS8sS8sV8s Generates random 8-bit signed integer samples in the specified range with WELL1024a pseudo-random number generator.
// All numbers between supportMin (inclusive) and supportMax (inclusive) are generated with the same probability.
func (s *WELL1024a) GenerateDiscreteUniformS8sS8sV8s(supportMin, supportMax Yep8s, samples []Yep8s) YepStatus {
	status := YepStatus(
		C.yepRandom_WELL1024a_GenerateDiscreteUniform_S8sS8s_V8s(
			(*C.struct_YepRandom_WELL1024a)(unsafe.Pointer(s)),
			(C.Yep8s)(supportMin),
			(C.Yep8s)(supportMax),
			(*C.Yep8s)(unsafe.Pointer(&samples[0])),
			(C.YepSize)(len(samples)),
		),
	)

	return status
}

// GenerateDiscreteUniformS16sS16sV16s Generates random 16-bit signed integer samples in the specified range with WELL1024a pseudo-random number generator.
// All numbers between supportMin (inclusive) and supportMax (inclusive) are generated with the same probability.
func (s *WELL1024a) GenerateDiscreteUniformS16sS16sV16s(supportMin, supportMax Yep16s, samples []Yep16s) YepStatus {
	status := YepStatus(
		C.yepRandom_WELL1024a_GenerateDiscreteUniform_S16sS16s_V16s(
			(*C.struct_YepRandom_WELL1024a)(unsafe.Pointer(s)),
			(C.Yep16s)(supportMin),
			(C.Yep16s)(supportMax),
			(*C.Yep16s)(unsafe.Pointer(&samples[0])),
			(C.YepSize)(len(samples)),
		),
	)

	return status
}

// GenerateDiscreteUniformS32sS32sV32s Generates random 32-bit signed integer samples in the specified range with WELL1024a pseudo-random number generator.
// All numbers between supportMin (inclusive) and supportMax (inclusive) are generated with the same probability.
func (s *WELL1024a) GenerateDiscreteUniformS32sS32sV32s(supportMin, supportMax Yep32s, samples []Yep32s) YepStatus {
	status := YepStatus(
		C.yepRandom_WELL1024a_GenerateDiscreteUniform_S32sS32s_V32s(
			(*C.struct_YepRandom_WELL1024a)(unsafe.Pointer(s)),
			(C.Yep32s)(supportMin),
			(C.Yep32s)(supportMax),
			(*C.Yep32s)(unsafe.Pointer(&samples[0])),
			(C.YepSize)(len(samples)),
		),
	)

	return status
}

// GenerateDiscreteUniformS64sS64sV64s Generates random 64-bit signed integer samples in the specified range with WELL1024a pseudo-random number generator.
// All numbers between supportMin (inclusive) and supportMax (inclusive) are generated with the same probability.
func (s *WELL1024a) GenerateDiscreteUniformS64sS64sV64s(supportMin, supportMax Yep64s, samples []Yep64s) YepStatus {
	status := YepStatus(
		C.yepRandom_WELL1024a_GenerateDiscreteUniform_S64sS64s_V64s(
			(*C.struct_YepRandom_WELL1024a)(unsafe.Pointer(s)),
			(C.Yep64s)(supportMin),
			(C.Yep64s)(supportMax),
			(*C.Yep64s)(unsafe.Pointer(&samples[0])),
			(C.YepSize)(len(samples)),
		),
	)

	return status
}

// GenerateDiscreteUniformS8uS8uV8u Generates random 8-bit unsigned integer samples in the specified range with WELL1024a pseudo-random number generator.
// All numbers between supportMin (inclusive) and supportMax (inclusive) are generated with the same probability.
func (s *WELL1024a) GenerateDiscreteUniformS8uS8uV8u(supportMin, supportMax Yep8u, samples []Yep8u) YepStatus {
	status := YepStatus(
		C.yepRandom_WELL1024a_GenerateDiscreteUniform_S8uS8u_V8u(
			(*C.struct_YepRandom_WELL1024a)(unsafe.Pointer(s)),
			(C.Yep8u)(supportMin),
			(C.Yep8u)(supportMax),
			(*C.Yep8u)(unsafe.Pointer(&samples[0])),
			(C.YepSize)(len(samples)),
		),
	)

	return status
}

// GenerateDiscreteUniformS16uS16uV16u Generates random 16-bit unsigned integer samples in the specified range with WELL1024a pseudo-random number generator.
// All numbers between supportMin (inclusive) and supportMax (inclusive) are generated with the same probability.
func (s *WELL1024a) GenerateDiscreteUniformS16uS16uV16u(supportMin, supportMax Yep16u, samples []Yep16u) YepStatus {
	status := YepStatus(
		C.yepRandom_WELL1024a_GenerateDiscreteUniform_S16uS16u_V16u(
			(*C.struct_YepRandom_WELL1024a)(unsafe.Pointer(s)),
			(C.Yep16u)(supportMin),
			(C.Yep16u)(supportMax),
			(*C.Yep16u)(unsafe.Pointer(&samples[0])),
			(C.YepSize)(len(samples)),
		),
	)

	return status
}

// GenerateDiscreteUniformS32uS32uV32u Generates random 32-bit unsigned integer samples in the specified range with WELL1024a pseudo-random number generator.
// All numbers between supportMin (inclusive) and supportMax (inclusive) are generated with the same probability.
func (s *WELL1024a) GenerateDiscreteUniformS32uS32uV32u(supportMin, supportMax Yep32u, samples []Yep32u) YepStatus {
	status := YepStatus(
		C.yepRandom_WELL1024a_GenerateDiscreteUniform_S32uS32u_V32u(
			(*C.struct_YepRandom_WELL1024a)(unsafe.Pointer(s)),
			(C.Yep32u)(supportMin),
			(C.Yep32u)(supportMax),
			(*C.Yep32u)(unsafe.Pointer(&samples[0])),
			(C.YepSize)(len(samples)),
		),
	)

	return status
}

// GenerateDiscreteUniformS64uS64uV64u Generates random 64-bit unsigned integer samples in the specified range with WELL1024a pseudo-random number generator.
// All numbers between supportMin (inclusive) and supportMax (inclusive) are generated with the same probability.
func (s *WELL1024a) GenerateDiscreteUniformS64uS64uV64u(supportMin, supportMax Yep64u, samples []Yep64u) YepStatus {
	status := YepStatus(
		C.yepRandom_WELL1024a_GenerateDiscreteUniform_S64uS64u_V64u(
			(*C.struct_YepRandom_WELL1024a)(unsafe.Pointer(s)),
			(C.Yep64u)(supportMin),
			(C.Yep64u)(supportMax),
			(*C.Yep64u)(unsafe.Pointer(&samples[0])),
			(C.YepSize)(len(samples)),
		),
	)

	return status
}

// GenerateUniformS32fS32fV32fAcc32 Generates random single precision floating-point samples in the specified range with WELL1024a pseudo-random number generator and 32-bit accuracy.
// All real numbers between supportMin (inclusive) and supportMax (inclusive) are generated with the same probability.
func (s *WELL1024a) GenerateUniformS32fS32fV32fAcc32(supportMin, supportMax Yep32f, samples []Yep32f) YepStatus {
	status := YepStatus(
		C.yepRandom_WELL1024a_GenerateUniform_S32fS32f_V32f_Acc32(
			(*C.struct_YepRandom_WELL1024a)(unsafe.Pointer(s)),
			(C.Yep32f)(supportMin),
			(C.Yep32f)(supportMax),
			(*C.Yep32f)(unsafe.Pointer(&samples[0])),
			(C.YepSize)(len(samples)),
		),
	)

	return status
}

// GenerateUniformS64fS64fV64fAcc32 Generates random double precision floating-point samples in the specified range with WELL1024a pseudo-random number generator and 32-bit accuracy.
// All real numbers between supportMin (inclusive) and supportMax (inclusive) are generated with the same probability.
func (s *WELL1024a) GenerateUniformS64fS64fV64fAcc32(supportMin, supportMax Yep64f, samples []Yep64f) YepStatus {
	status := YepStatus(
		C.yepRandom_WELL1024a_GenerateUniform_S64fS64f_V64f_Acc32(
			(*C.struct_YepRandom_WELL1024a)(unsafe.Pointer(s)),
			(C.Yep64f)(supportMin),
			(C.Yep64f)(supportMax),
			(*C.Yep64f)(unsafe.Pointer(&samples[0])),
			(C.YepSize)(len(samples)),
		),
	)

	return status
}

// GenerateUniformS64fS64fV64fAcc64 Generates random double precision floating-point samples in the specified range with WELL1024a pseudo-random number generator and 64-bit accuracy.
// All real numbers between supportMin (inclusive) and supportMax (inclusive) are generated with the same probability.
func (s *WELL1024a) GenerateUniformS64fS64fV64fAcc64(supportMin, supportMax Yep64f, samples []Yep64f) YepStatus {
	status := YepStatus(
		C.yepRandom_WELL1024a_GenerateUniform_S64fS64f_V64f_Acc64(
			(*C.struct_YepRandom_WELL1024a)(unsafe.Pointer(s)),
			(C.Yep64f)(supportMin),
			(C.Yep64f)(supportMax),
			(*C.Yep64f)(unsafe.Pointer(&samples[0])),
			(C.YepSize)(len(samples)),
		),
	)

	return status
}

// GenerateFPUniformS32fS32fV32f Generates random single precision floating-point samples in the specified range with WELL1024a pseudo-random number generator.
// All floating-point numbers between supportMin (inclusive) and supportMax (inclusive) are generated with the same probability.
func (s *WELL1024a) GenerateFPUniformS32fS32fV32f(supportMin, supportMax Yep32f, samples []Yep32f) YepStatus {
	status := YepStatus(
		C.yepRandom_WELL1024a_GenerateFPUniform_S32fS32f_V32f(
			(*C.struct_YepRandom_WELL1024a)(unsafe.Pointer(s)),
			(C.Yep32f)(supportMin),
			(C.Yep32f)(supportMax),
			(*C.Yep32f)(unsafe.Pointer(&samples[0])),
			(C.YepSize)(len(samples)),
		),
	)

	return status
}

// GenerateFPUniformS64fS64fV64f Generates random double precision floating-point samples in the specified range with WELL1024a pseudo-random number generator.
// All floating-point numbers between supportMin (inclusive) and supportMax (inclusive) are generated with the same probability.
func (s *WELL1024a) GenerateFPUniformS64fS64fV64f(supportMin, supportMax Yep64f, samples []Yep64f) YepStatus {
	status := YepStatus(
		C.yepRandom_WELL1024a_GenerateFPUniform_S64fS64f_V64f(
			(*C.struct_YepRandom_WELL1024a)(unsafe.Pointer(s)),
			(C.Yep64f)(supportMin),
			(C.Yep64f)(supportMax),
			(*C.Yep64f)(unsafe.Pointer(&samples[0])),
			(C.YepSize)(len(samples)),
		),
	)

	return status
}
