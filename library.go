// Package gy Yeppp! library binding
package gy

/*
#include <yepLibrary.h>

#cgo pkg-config: yeppp

*/
import "C"

import (
	"errors"
	"unsafe"
)

/////////////// CPU ///////////////////

// YepCpuVendor The company which designed the processor microarchitecture.
type YepCpuVendor int

const (
	// YepCpuVendorUnknown Processor vendor is not known to the library, or the library failed to get vendor information from the OS.
	YepCpuVendorUnknown YepCpuVendor = iota

	//x86/x86-64 Cpus

	// YepCpuVendorIntel Intel Corporation. Vendor of x86, x86-64, IA64, and ARM processor microarchitectures.
	YepCpuVendorIntel

	// YepCpuVendorAMD Advanced Micro Devices, Inc. Vendor of x86 and x86-64 processor microarchitectures.
	YepCpuVendorAMD

	// YepCpuVendorVIA VIA Technologies, Inc. Vendor of x86 and x86-64 processor microarchitectures.
	YepCpuVendorVIA

	// YepCpuVendorTransmeta Transmeta Corporation. Vendor of x86 processor microarchitectures.
	YepCpuVendorTransmeta

	// YepCpuVendorCyrix Cyrix Corporation. Vendor of x86 processor microarchitectures.
	YepCpuVendorCyrix

	// YepCpuVendorRise Rise Technology. Vendor of x86 processor microarchitectures.
	YepCpuVendorRise

	// YepCpuVendorNSC National Semiconductor. Vendor of x86 processor microarchitectures.
	YepCpuVendorNSC

	// YepCpuVendorSiS Silicon Integrated Systems. Vendor of x86 processor microarchitectures.
	YepCpuVendorSiS

	// YepCpuVendorNexGen NexGen. Vendor of x86 processor microarchitectures.
	YepCpuVendorNexGen

	// YepCpuVendorUMC United Microelectronics Corporation. Vendor of x86 processor microarchitectures.
	YepCpuVendorUMC

	// YepCpuVendorRDC RDC Semiconductor Co., Ltd. Vendor of x86 processor microarchitectures.
	YepCpuVendorRDC

	// YepCpuVendorDMP DM&P Electronics Inc. Vendor of x86 processor microarchitectures.
	YepCpuVendorDMP

	/* ARM Cpus */

	// YepCpuVendorARM ARM Holdings plc. Vendor of ARM processor microarchitectures.
	YepCpuVendorARM

	// YepCpuVendorMarvell Marvell Technology Group Ltd. Vendor of ARM processor microarchitectures.
	YepCpuVendorMarvell

	// YepCpuVendorQualcomm Qualcomm Incorporated. Vendor of ARM processor microarchitectures.
	YepCpuVendorQualcomm

	// YepCpuVendorDEC Digital Equipment Corporation. Vendor of ARM processor microarchitecture.
	YepCpuVendorDEC

	// YepCpuVendorTI Texas Instruments Inc. Vendor of ARM processor microarchitectures.
	YepCpuVendorTI

	// YepCpuVendorApple Apple Inc. Vendor of ARM processor microarchitectures.
	YepCpuVendorApple

	/* MIPS Cpus */

	// YepCpuVendorIngenic Ingenic Semiconductor. Vendor of MIPS processor microarchitectures.
	YepCpuVendorIngenic

	// YepCpuVendorICT Institute of Computing Technology of the Chinese Academy of Sciences. Vendor of MIPS processor microarchitectures.
	YepCpuVendorICT

	// YepCpuVendorMIPS MIPS Technologies, Inc. Vendor of MIPS processor microarchitectures.
	YepCpuVendorMIPS

	/* PowerPC Cpus */

	// YepCpuVendorIBM International Business Machines Corporation. Vendor of PowerPC processor microarchitectures.
	YepCpuVendorIBM

	// YepCpuVendorMotorola Motorola Inc. Vendor of PowerPC and ARM processor microarchitectures.
	YepCpuVendorMotorola

	// YepCpuVendorPASemi P. A. Semi. Vendor of PowerPC processor microarchitectures.
	YepCpuVendorPASemi

	/* SPARC Cpus */

	// YepCpuVendorSun Sun Microsystems, Inc. Vendor of SPARC processor microarchitectures.
	YepCpuVendorSun

	// YepCpuVendorOracle Oracle Corporation. Vendor of SPARC processor microarchitectures.
	YepCpuVendorOracle

	// YepCpuVendorFujitsu Fujitsu Limited. Vendor of SPARC processor microarchitectures. */
	YepCpuVendorFujitsu

	// YepCpuVendorMCST Moscow Center of SPARC Technologies CJSC. Vendor of SPARC processor microarchitectures. */
	YepCpuVendorMCST
)

// YepCpuArchitecture The basic instruction set architecture of the processor.
type YepCpuArchitecture int

const (
	// YepCpuArchitectureUnknown Instruction set architecture is not known to the library. */
	YepCpuArchitectureUnknown = iota

	// YepCpuArchitectureX86 x86 or x86-64 ISA.
	YepCpuArchitectureX86

	// YepCpuArchitectureARM ARM ISA.
	YepCpuArchitectureARM

	// YepCpuArchitectureMIPS MIPS ISA.
	YepCpuArchitectureMIPS

	// YepCpuArchitecturePowerPC PowerPC ISA.
	YepCpuArchitecturePowerPC

	// YepCpuArchitectureIA64 IA64 ISA.
	YepCpuArchitectureIA64

	// YepCpuArchitectureSPARC SPARC ISA.
	YepCpuArchitectureSPARC
)

// YepCpuMicroarchitecture Type of processor microarchitecture.
type YepCpuMicroarchitecture int

const (
	// YepCpuMicroarchitectureUnknown Microarchitecture is unknown, or the library failed to get information about the microarchitecture from OS
	YepCpuMicroarchitectureUnknown YepCpuMicroarchitecture = 0

	// YepCpuMicroarchitectureP5 Pentium and Pentium MMX microarchitecture.
	YepCpuMicroarchitectureP5 = (YepCpuArchitectureX86 << 24) + (YepCpuVendorIntel << 16) + 0x0001

	// YepCpuMicroarchitectureP6 Pentium Pro, Pentium II, and Pentium III.
	YepCpuMicroarchitectureP6 = (YepCpuArchitectureX86 << 24) + (YepCpuVendorIntel << 16) + 0x0002

	// YepCpuMicroarchitectureWillamette Pentium 4 with Willamette, Northwood, or Foster cores. */
	YepCpuMicroarchitectureWillamette = (YepCpuArchitectureX86 << 24) + (YepCpuVendorIntel << 16) + 0x0003

	// YepCpuMicroarchitecturePrescott Pentium 4 with Prescott and later cores.
	YepCpuMicroarchitecturePrescott = (YepCpuArchitectureX86 << 24) + (YepCpuVendorIntel << 16) + 0x0004

	// YepCpuMicroarchitectureDothan Pentium M.
	YepCpuMicroarchitectureDothan = (YepCpuArchitectureX86 << 24) + (YepCpuVendorIntel << 16) + 0x0005

	//YepCpuMicroarchitectureYonah Intel Core microarchitecture.
	YepCpuMicroarchitectureYonah = (YepCpuArchitectureX86 << 24) + (YepCpuVendorIntel << 16) + 0x0006

	//YepCpuMicroarchitectureConroe Intel Core 2 microarchitecture on 65 nm process.
	YepCpuMicroarchitectureConroe = (YepCpuArchitectureX86 << 24) + (YepCpuVendorIntel << 16) + 0x0007

	// YepCpuMicroarchitecturePenryn Intel Core 2 microarchitecture on 45 nm process.
	YepCpuMicroarchitecturePenryn = (YepCpuArchitectureX86 << 24) + (YepCpuVendorIntel << 16) + 0x0008

	// YepCpuMicroarchitectureBonnell Intel Atom on 45 nm process.
	YepCpuMicroarchitectureBonnell = (YepCpuArchitectureX86 << 24) + (YepCpuVendorIntel << 16) + 0x0009

	// YepCpuMicroarchitectureNehalem Intel Nehalem and Westmere microarchitectures (Core i3/i5/i7 1st gen).
	YepCpuMicroarchitectureNehalem = (YepCpuArchitectureX86 << 24) + (YepCpuVendorIntel << 16) + 0x000A

	//YepCpuMicroarchitectureSandyBridge Intel Sandy Bridge microarchitecture (Core i3/i5/i7 2nd gen).
	YepCpuMicroarchitectureSandyBridge = (YepCpuArchitectureX86 << 24) + (YepCpuVendorIntel << 16) + 0x000B

	// YepCpuMicroarchitectureSaltwell Intel Atom on 32 nm process.
	YepCpuMicroarchitectureSaltwell = (YepCpuArchitectureX86 << 24) + (YepCpuVendorIntel << 16) + 0x000C

	// YepCpuMicroarchitectureIvyBridge Intel Ivy Bridge microarchitecture (Core i3/i5/i7 3rd gen).
	YepCpuMicroarchitectureIvyBridge = (YepCpuArchitectureX86 << 24) + (YepCpuVendorIntel << 16) + 0x000D

	// YepCpuMicroarchitectureHaswell Intel Haswell microarchitecture (Core i3/i5/i7 4th gen).
	YepCpuMicroarchitectureHaswell = (YepCpuArchitectureX86 << 24) + (YepCpuVendorIntel << 16) + 0x000E

	// YepCpuMicroarchitectureSilvermont Intel Silvermont microarchitecture (22 nm out-of-order Atom).
	YepCpuMicroarchitectureSilvermont = (YepCpuArchitectureX86 << 24) + (YepCpuVendorIntel << 16) + 0x000F

	// YepCpuMicroarchitectureKnightsFerry Intel Knights Ferry HPC boards.
	YepCpuMicroarchitectureKnightsFerry = (YepCpuArchitectureX86 << 24) + (YepCpuVendorIntel << 16) + 0x0100

	// YepCpuMicroarchitectureKnightsCorner Intel Knights Corner HPC boards (aka Xeon Phi).
	YepCpuMicroarchitectureKnightsCorner = (YepCpuArchitectureX86 << 24) + (YepCpuVendorIntel << 16) + 0x0101

	// YepCpuMicroarchitectureK5 AMD K5.
	YepCpuMicroarchitectureK5 = (YepCpuArchitectureX86 << 24) + (YepCpuVendorAMD << 16) + 0x0001

	// YepCpuMicroarchitectureK6 AMD K6 and alike.
	YepCpuMicroarchitectureK6 = (YepCpuArchitectureX86 << 24) + (YepCpuVendorAMD << 16) + 0x0002

	// YepCpuMicroarchitectureK7 AMD Athlon and Duron.
	YepCpuMicroarchitectureK7 = (YepCpuArchitectureX86 << 24) + (YepCpuVendorAMD << 16) + 0x0003

	//YepCpuMicroarchitectureGeode AMD Geode GX and LX.
	YepCpuMicroarchitectureGeode = (YepCpuArchitectureX86 << 24) + (YepCpuVendorAMD << 16) + 0x0004

	// YepCpuMicroarchitectureK8 AMD Athlon 64, Opteron 64.
	YepCpuMicroarchitectureK8 = (YepCpuArchitectureX86 << 24) + (YepCpuVendorAMD << 16) + 0x0005

	// YepCpuMicroarchitectureK10 AMD K10 (Barcelona, Istambul, Magny-Cours).
	YepCpuMicroarchitectureK10 = (YepCpuArchitectureX86 << 24) + (YepCpuVendorAMD << 16) + 0x0006

	// YepCpuMicroarchitectureBobcat AMD Bobcat mobile microarchitecture.
	YepCpuMicroarchitectureBobcat = (YepCpuArchitectureX86 << 24) + (YepCpuVendorAMD << 16) + 0x0007

	// YepCpuMicroarchitectureBulldozer AMD Bulldozer microarchitecture (1st gen K15).
	YepCpuMicroarchitectureBulldozer = (YepCpuArchitectureX86 << 24) + (YepCpuVendorAMD << 16) + 0x0008

	// YepCpuMicroarchitecturePiledriver AMD Piledriver microarchitecture (2nd gen K15).
	YepCpuMicroarchitecturePiledriver = (YepCpuArchitectureX86 << 24) + (YepCpuVendorAMD << 16) + 0x0009

	// YepCpuMicroarchitectureJaguar AMD Jaguar mobile microarchitecture.
	YepCpuMicroarchitectureJaguar = (YepCpuArchitectureX86 << 24) + (YepCpuVendorAMD << 16) + 0x000A

	// YepCpuMicroarchitectureSteamroller AMD Steamroller microarchitecture (3rd gen K15).
	YepCpuMicroarchitectureSteamroller = (YepCpuArchitectureX86 << 24) + (YepCpuVendorAMD << 16) + 0x000B

	// YepCpuMicroarchitectureStrongARM DEC/Intel StrongARM processors.
	YepCpuMicroarchitectureStrongARM = (YepCpuArchitectureARM << 24) + (YepCpuVendorIntel << 16) + 0x0001

	// YepCpuMicroarchitectureXScale Intel/Marvell XScale processors.
	YepCpuMicroarchitectureXScale = (YepCpuArchitectureARM << 24) + (YepCpuVendorIntel << 16) + 0x0002

	// YepCpuMicroarchitectureARM7 ARM7 series.
	YepCpuMicroarchitectureARM7 = (YepCpuArchitectureARM << 24) + (YepCpuVendorARM << 16) + 0x0001

	// YepCpuMicroarchitectureARM9 ARM9 series.
	YepCpuMicroarchitectureARM9 = (YepCpuArchitectureARM << 24) + (YepCpuVendorARM << 16) + 0x0002

	// YepCpuMicroarchitectureARM11 ARM 1136, ARM 1156, ARM 1176, or ARM 11MPCore.
	YepCpuMicroarchitectureARM11 = (YepCpuArchitectureARM << 24) + (YepCpuVendorARM << 16) + 0x0003

	// YepCpuMicroarchitectureCortexA5 ARM Cortex-A5.
	YepCpuMicroarchitectureCortexA5 = (YepCpuArchitectureARM << 24) + (YepCpuVendorARM << 16) + 0x0004

	// YepCpuMicroarchitectureCortexA7 ARM Cortex-A7.
	YepCpuMicroarchitectureCortexA7 = (YepCpuArchitectureARM << 24) + (YepCpuVendorARM << 16) + 0x0005

	// YepCpuMicroarchitectureCortexA8 ARM Cortex-A8.
	YepCpuMicroarchitectureCortexA8 = (YepCpuArchitectureARM << 24) + (YepCpuVendorARM << 16) + 0x0006

	// YepCpuMicroarchitectureCortexA9 ARM Cortex-A9.
	YepCpuMicroarchitectureCortexA9 = (YepCpuArchitectureARM << 24) + (YepCpuVendorARM << 16) + 0x0007

	// YepCpuMicroarchitectureCortexA15 ARM Cortex-A15.
	YepCpuMicroarchitectureCortexA15 = (YepCpuArchitectureARM << 24) + (YepCpuVendorARM << 16) + 0x0008

	// YepCpuMicroarchitectureScorpion Qualcomm Scorpion.
	YepCpuMicroarchitectureScorpion = (YepCpuArchitectureARM << 24) + (YepCpuVendorQualcomm << 16) + 0x0001

	// YepCpuMicroarchitectureKrait Qualcomm Krait.
	YepCpuMicroarchitectureKrait = (YepCpuArchitectureARM << 24) + (YepCpuVendorQualcomm << 16) + 0x0002

	// YepCpuMicroarchitecturePJ1 Marvell Sheeva PJ1.
	YepCpuMicroarchitecturePJ1 = (YepCpuArchitectureARM << 24) + (YepCpuVendorMarvell << 16) + 0x0001

	// YepCpuMicroarchitecturePJ4 Marvell Sheeva PJ4.
	YepCpuMicroarchitecturePJ4 = (YepCpuArchitectureARM << 24) + (YepCpuVendorMarvell << 16) + 0x0002

	// YepCpuMicroarchitectureSwift Apple A6 and A6X processors.
	YepCpuMicroarchitectureSwift = (YepCpuArchitectureARM << 24) + (YepCpuVendorApple << 16) + 0x0001

	// YepCpuMicroarchitectureItanium Intel Itanium.
	YepCpuMicroarchitectureItanium = (YepCpuArchitectureIA64 << 24) + (YepCpuVendorIntel << 16) + 0x0001

	// YepCpuMicroarchitectureItanium2 Intel Itanium 2.
	YepCpuMicroarchitectureItanium2 = (YepCpuArchitectureIA64 << 24) + (YepCpuVendorIntel << 16) + 0x0002

	//YepCpuMicroarchitectureMIPS24K MIPS 24K.
	YepCpuMicroarchitectureMIPS24K = (YepCpuArchitectureMIPS << 24) + (YepCpuVendorMIPS << 16) + 0x0001

	// YepCpuMicroarchitectureMIPS34K MIPS 34K.
	YepCpuMicroarchitectureMIPS34K = (YepCpuArchitectureMIPS << 24) + (YepCpuVendorMIPS << 16) + 0x0002

	// YepCpuMicroarchitectureMIPS74K MIPS 74K.
	YepCpuMicroarchitectureMIPS74K = (YepCpuArchitectureMIPS << 24) + (YepCpuVendorMIPS << 16) + 0x0003

	// YepCpuMicroarchitectureXBurst Ingenic XBurst.
	YepCpuMicroarchitectureXBurst = (YepCpuArchitectureMIPS << 24) + (YepCpuVendorIngenic << 16) + 0x0001

	// YepCpuMicroarchitectureXBurst2 Ingenic XBurst 2.
	YepCpuMicroarchitectureXBurst2 = (YepCpuArchitectureMIPS << 24) + (YepCpuVendorIngenic << 16) + 0x0002
)

// Init Initialized the Yeppp library.
func Init() (YepStatus, error) {
	status := YepStatus(C.yepLibrary_Init())

	if status == YepStatusSystemError {
		return status, errors.New("An uncoverable error inside the OS kernel occurred during library initialization.")
	}

	return status, nil
}

// Release Deinitialized the Yeppp library and releases the consumed system resources.
func Release() (YepStatus, error) {
	status := YepStatus(C.yepLibrary_Release())

	if status == YepStatusSystemError {
		return status, errors.New("The library failed to release some of the resources due to a failed call to the OS kernel.")
	}

	return status, nil
}

// CPUIsaFeatures returns as information about the supported ISA extensions (excluding SIMD extensions) as a 64-bit mask where information about the supported ISA extensions will be stored.
func CPUIsaFeatures() (uint64, YepStatus, error) {
	var isaFeatures uint64

	status := YepStatus(C.yepLibrary_GetCpuIsaFeatures((*C.Yep64u)(unsafe.Pointer(&isaFeatures))))

	if status == YepStatusNullPointer {
		return isaFeatures, status, errors.New("The isaFeatures is null.")
	}

	return isaFeatures, status, nil

}

// CPUSimdFeatures returns information about the supported SIMD extensions as 64-bit mask where information about the supported SIMD extensions will be stored.
func CPUSimdFeatures() (uint64, YepStatus, error) {
	var smdFeatures uint64

	status := YepStatus(C.yepLibrary_GetCpuIsaFeatures((*C.Yep64u)(unsafe.Pointer(&smdFeatures))))

	if status == YepStatusNullPointer {
		return smdFeatures, status, errors.New("The smdFeatures is null.")
	}

	return smdFeatures, status, nil

}

// CPUSystemFeatures returns information about processor features other than ISA extensions, and OS features related to Cpu.
func CPUSystemFeatures() (uint64, YepStatus, error) {
	var cpuFeatures uint64

	status := YepStatus(C.yepLibrary_GetCpuSystemFeatures((*C.Yep64u)(unsafe.Pointer(&cpuFeatures))))

	if status == YepStatusNullPointer {
		return cpuFeatures, status, errors.New("The systemFeatures is null.")
	}

	return cpuFeatures, status, nil
}

// CPUMicroarchitecture returns the type of processor microarchitecture used.
func (cm *YepCpuMicroarchitecture) CPUMicroarchitecture() (YepStatus, error) {

	status := YepStatus(C.yepLibrary_GetCpuMicroarchitecture((*C.enum_YepCpuMicroarchitecture)(unsafe.Pointer(cm))))

	if status == YepStatusNullPointer {
		return status, errors.New("The microarchitecture is null.")
	}

	return status, nil
}

/////////////// Energy ///////////////////

// YepEnergyCounter Energy counter state.
type YepEnergyCounter struct {
	State [6]uint64
}

// YepEnergyCounterType Energy counter type.
type YepEnergyCounterType int

const (
	// YepEnergyCounterTypeRaplPackageEnergy Intel RAPL per-package energy counter.
	// This counter is supported on Intel Sandy Bridge and Ivy Bridge processors, and estimates the energy (in Joules) consumed by all chips in the Cpu package.
	YepEnergyCounterTypeRaplPackageEnergy = 1

	// YepEnergyCounterTypeRaplPowerPlane0Energy Intel RAPL power plane 0 energy counter.
	// This counter is supported on Intel Sandy Bridge and Ivy Bridge processors, and estimates the energy (in Joules) consumed by power plane 0 (includes Cpu cores and caches).
	YepEnergyCounterTypeRaplPowerPlane0Energy = 2

	// YepEnergyCounterTypeRaplPowerPlane1Energy Intel RAPL power plane 1 energy counter.
	// This counter is supported on Intel Sandy Bridge and Ivy Bridge processors, and estimates the energy (in Joules) consumed by power plane 1 (includes GPU cores).
	YepEnergyCounterTypeRaplPowerPlane1Energy = 3

	// YepEnergyCounterTypeRaplDRAMEnergy Intel RAPL DRAM energy counter.
	// This counter is supported on Intel Sandy Bridge E processors, and estimates the energy (in Joules) consumed by DRAM modules.
	// Motherboard support is required to use this counter.
	YepEnergyCounterTypeRaplDRAMEnergy = 4

	// YepEnergyCounterTypeRaplPackagePower Intel RAPL per-package power counter.
	// This counter is supported on Intel Sandy Bridge and Ivy Bridge processors, and estimates the average power (in Watts) consumed by all chips in the Cpu package. This counter is implemented as a combination of RAPL per-package energy counter and system timer.
	YepEnergyCounterTypeRaplPackagePower = 5

	// YepEnergyCounterTypeRaplPowerPlane0Power Intel RAPL power plane 0 power counter.
	// This counter is supported on Intel Sandy Bridge and Ivy Bridge processors, and estimates the average power (in Watts) consumed by power plane 0 (includes Cpu cores and caches).
	// This counter is implemented as a combination of RAPL power plane 0 energy counter and system timer. */
	YepEnergyCounterTypeRaplPowerPlane0Power = 6

	// YepEnergyCounterTypeRaplPowerPlane1Power Intel RAPL power plane 1 power counter.
	// This counter is supported on Intel Sandy Bridge and Ivy Bridge processors, and estimates the average power (in Watts) consumed by power plane 1 (includes GPU cores).
	// This counter is implemented as a combination of RAPL power plane 1 energy counter and system timer. */
	YepEnergyCounterTypeRaplPowerPlane1Power = 7

	// YepEnergyCounterTypeRaplDRAMPower Intel RAPL DRAM power counter.
	// This counter is supported on Intel Sandy Bridge E processors, and estimates the average power (in Watts) consumed by DRAM modules.
	// This counter is implemented as a combination of RAPL DRAM energy counter and system timer.
	// Motherboard support is required to use this counter.
	YepEnergyCounterTypeRaplDRAMPower = 8
)

// EnergyCounterAcquire Initializes the specified energy counter and starts energy measurements.
func (ec *YepEnergyCounter) EnergyCounterAcquire(ct YepEnergyCounterType) (YepStatus, error) {
	status := YepStatus(C.yepLibrary_GetEnergyCounterAcquire((C.enum_YepEnergyCounterType)(ct), (*C.struct_YepEnergyCounter)(unsafe.Pointer(ec))))

	switch status {
	case YepStatusNullPointer:
		return status, errors.New("The microarchitecture is null.")
	case YepStatusInvalidArgument:
		return status, errors.New("The ct type parameter does not specify a valid energy counter type.")
	case YepStatusUnsupportedHardware:
		return status, errors.New("The hardware does not support the requested energy counter type.")
	case YepStatusUnsupportedSoftware:
		return status, errors.New("The operating system does not provide access to the specified energy counter.")
	case YepStatusSystemError:
		return status, errors.New("An attempt to read the energy counter or release the OS resources failed inside the OS kernel.")
	case YepStatusAccessDenied:
		return status, errors.New("The user does not possess the required access rights to read the energy counter.")
	}

	return status, nil
}
