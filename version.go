package goyeppp

/*
#include <yepVersion.h>

#cgo pkg-config: yeppp

static int getMajorVersion(void)
{
        return YEP_MAJOR_VERSION;
}

static int getMinorVersion(void)
{
        return YEP_MINOR_VERSION;
}

static int getPatchVersion(void)
{
        return YEP_PATCH_VERSION;
}


static int getBuildVersion(void)
{
        return YEP_BUILD_VERSION;

}

static char *getReleaseName(void)
{
        return YEP_RELEASE_NAME;
}

static char *getFullName(void)
{
        return YEP_FULL_VERSION_STR;
}
*/
import "C"

// MajorVersion return the Yeppp major version number
func MajorVersion() int {
	return int(C.getMajorVersion())
}

// MinorVersion return the Yeppp minor version number
func MinorVersion() int {
	return int(C.getMinorVersion())
}

// PatchVersion return the Yeppp patch version
func PatchVersion() int {
	return int(C.getPatchVersion())
}

// BuildVersion return the Yeppp build version
func BuildVersion() int {
	return int(C.getBuildVersion())

}

// ReleaseName return the Yeppp release name
func ReleaseName() string {
	return C.GoString(C.getReleaseName())
}

// FullName return the Yeppp full name
func FullName() string {
	return C.GoString(C.getFullName())
}
