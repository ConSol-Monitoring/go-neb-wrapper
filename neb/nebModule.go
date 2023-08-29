package neb

/*

#cgo CFLAGS: -DNAEMON
#cgo pkg-config: naemon

#include "dependencies.h"

extern nebmodule *neb_handle;

*/
import "C"

import (
	"unsafe"
)

// Name is used for some default logs
var Name = "go_neb_broker"

// Title is used for the module info
var Title = ""

// Author is used for the module info
var Author = ""

// Version is used for the module info
var Version = ""

// License is used for the module info
var License = ""

// Desc is used for the module info
var Desc = ""

// NebModuleInitHook gives you access to the flags and args which the core passes to the brokers, when it's loading your code.
// This function will be called at the end of the init function
// return Ok if everything went well
var NebModuleInitHook func(flags int, args string) int

// NebModuleDeinitHook gives you access to the flags and reason which the core passes to the brokers, when it's unloading your code.
// This function will be called at the end of the deinit function
// return Ok if everything went well
var NebModuleDeinitHook func(flags, reason int) int

//export GoNebModuleInit
func GoNebModuleInit(flags int, args *C.char) C.int {
	handle := unsafe.Pointer(C.neb_handle)
	defer C.free(unsafe.Pointer(args))
	modinfoMapping := map[C.int]string{
		C.NEBMODULE_MODINFO_TITLE:   Title,
		C.NEBMODULE_MODINFO_AUTHOR:  Author,
		C.NEBMODULE_MODINFO_VERSION: Version,
		C.NEBMODULE_MODINFO_LICENSE: License,
		C.NEBMODULE_MODINFO_DESC:    Desc,
	}
	// write module info
	for infoType, value := range modinfoMapping {
		setModuleInfo(handle, infoType, value)
	}

	// set callbacks
	initCallbacks()

	// default returncode
	returnCode := Ok
	// load Init hook if needed
	if NebModuleInitHook != nil {
		returnCode = NebModuleInitHook(flags, C.GoString(args))
	}
	CoreFLog("finished Init with returncode: %d\n", returnCode)
	return C.int(returnCode)
}

func setModuleInfo(handle unsafe.Pointer, infoType C.int, value string) {
	cValue := C.CString(value)
	C.neb_set_module_info(handle, infoType, cValue)
	C.free(unsafe.Pointer(cValue))
}

//export GoNebModuleDeinit
func GoNebModuleDeinit(flags, reason int) C.int {
	// unload callbacks
	deinitCallbacks()
	// default returncode
	returnCode := Ok
	// load Init hook if needed
	if NebModuleInitHook != nil {
		returnCode = NebModuleDeinitHook(flags, reason)
	}
	CoreFLog("finished Deinit with returncode: %d\n", returnCode)
	return C.int(returnCode)
}
