package goyeppp

/*
#include <yepTypes.h>
*/
import "C"

// Integral Type

// Yep8u 8-bit unsigned integer type
type Yep8u uint8

// Yep16u 16-bit unsigned integer type
type Yep16u uint16

// Yep32u 32-bit unsigned integer type
type Yep32u uint32

// Yep64u 64-bit unsigned integer type
type Yep64u uint64

// Yep8s 8-bit signed integer type
type Yep8s int8

// Yep16s 16-bit signed integer type
type Yep16s int16

// Yep32s 32-bit signed integer type
type Yep32s int32

// Yep64s 64-bit signed integer type
type Yep64s int64

// Yep16f Half-precision (16-bit) IEEE floating point type.
type Yep16f uint16

// Yep32f Single-precision (32-bit) IEEE floating point type.
type Yep32f float32

// Yep64f Double-precision (64-bit) IEEE floating point type.
type Yep64f float64

// YepStatus Indicates success or failure of Yeppp! functions.
type YepStatus int

// YepBoolean boolean type
type YepBoolean bool

// Yep16fc Complex half-precision (16-bit) IEEE floating point type.
type Yep16fc complex64

// Yep32fc Complex single-precision (32-bit) IEEE floating point type.
type Yep32fc complex64

// Yep64fc Complex double-precision (64-bit) IEEE floating point type.
type Yep64fc complex64

// Yep32df Dual single-precision (32-bit) IEEE floating point type.
// A number of this type is represented as an unevaluated sum of two Yep32f numbers.
type Yep32df struct {
	high Yep32f
	low  Yep32f
}

// Yep64df Dual double-precision (64-bit) IEEE floating point type.
// A number of this type is represented as an unevaluated sum of two Yep64f numbers.
type Yep64df struct {
	high Yep64f
	low  Yep64f
}

// Yep128u 128-bit unsigned integer type.
type Yep128u struct {
	low  Yep64u
	high Yep64u
}

// Yep128s 128-bit signed integer type.
type Yep128s struct {
	low  Yep64u
	high Yep64s
}

// YepSize Unsigned integer type of pointer width.
type YepSize C.YepSize

const (
	// YepStatusOk operation finished successfully.
	YepStatusOk YepStatus = iota

	// YepStatusNullPointer function call failed because one of the pointer arguments is null.
	YepStatusNullPointer

	// YepStatusMisalignedPointer function call failed because one of the pointer arguments is not properly aligned.
	YepStatusMisalignedPointer

	// YepStatusInvalidArgument function call failed because one of the integer arguments has unsupported value.
	YepStatusInvalidArgument

	// YepStatusInvalidData function call failed because some of the data passed to the function has invalid format or values.
	YepStatusInvalidData

	// YepStatusInvalidState function call failed because one of the state objects passed is corrupted.
	YepStatusInvalidState

	// YepStatusUnsupportedHardware function call failed because the system hardware does not support the requested operation.
	YepStatusUnsupportedHardware

	// YepStatusUnsupportedSoftware function call failed because the operating system does not support the requested operation.
	YepStatusUnsupportedSoftware

	// YepStatusInsufficientBuffer function call failed because the provided output buffer is too small or exhausted.
	YepStatusInsufficientBuffer

	// YepStatusOutOfMemory function call failed because the library could not allocate the memory.
	YepStatusOutOfMemory

	// YepStatusSystemError function call failed because some of the system calls inside the function failed.
	YepStatusSystemError

	// YepStatusAccessDenied function call failed because access to the requested resource is not allowed for this user.
	YepStatusAccessDenied
)

// Name return the status name
func (s *YepStatus) Name() string {
	switch *s {
	case YepStatusOk:
		return "YepStatusOk"

	case YepStatusNullPointer:
		return "YepStatusNullPointer"

	case YepStatusMisalignedPointer:
		return "YepStatusMisalignedPointer"

	case YepStatusInvalidArgument:
		return "YepStatusInvalidArgument"

	case YepStatusInvalidData:
		return "YepStatusInvalidData"

	case YepStatusInvalidState:
		return "YepStatusInvalidState"

	case YepStatusUnsupportedHardware:
		return "YepStatusUnsupportedHardware"

	case YepStatusUnsupportedSoftware:
		return "YepStatusUnsupportedSoftware"

	case YepStatusInsufficientBuffer:
		return "YepStatusInsufficientBuffer"

	case YepStatusOutOfMemory:
		return "YepStatusOutOfMemory"

	case YepStatusSystemError:
		return "YepStatusSystemError"

	case YepStatusAccessDenied:
		return "YepStatusAccessDenied"
	}

	return "Not defined"
}
