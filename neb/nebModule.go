package neb

/*

#include "naemon/naemon.h"

extern nebmodule *neb_handle;

*/
import "C"
import (
	"fmt"
	"unsafe"
)

//Name is used for some default logs
var Name = "go_neb_broker"

//Title is used for the module info
var Title = ""

//Author is used for the module info
var Author = ""

//Version is used for the module info
var Version = ""

//License is used for the module info
var License = ""

//Desc is used for the module info
var Desc = ""

//NebModuleInitHook gives you access to the flags and args which the core passes to the brokers, when it's loading your code.
//This function will be called at the end of the init function
//return NebOk if everything went well
var NebModuleInitHook func(flags int, args string) int

//NebModuleDeinitHook gives you access to the flags and reason which the core passes to the brokers, when it's unloading your code.
//This function will be called at the end of the deinit function
//return NebOk if everything went well
var NebModuleDeinitHook func(flags, reason int) int

//export Neb_Module_Init
func Neb_Module_Init(flags int, args *C.char) C.int {
	handle := unsafe.Pointer(C.neb_handle)
	defer C.free(unsafe.Pointer(args))
	modinfoMapping := map[C.int]string{
		C.NEBMODULE_MODINFO_TITLE:   Title,
		C.NEBMODULE_MODINFO_AUTHOR:  Author,
		C.NEBMODULE_MODINFO_VERSION: Version,
		C.NEBMODULE_MODINFO_LICENSE: License,
		C.NEBMODULE_MODINFO_DESC:    Desc,
	}
	for infoType, value := range modinfoMapping {
		setModuleInfo(handle, infoType, value)
	}
	initCallbacks()
	if NebModuleInitHook == nil {
		return NebOk
	} else {
		return C.int(NebModuleInitHook(flags, C.GoString(args)))
	}

}

func setModuleInfo(handle unsafe.Pointer, infoType C.int, value string) {
	cValue := C.CString(value)
	C.neb_set_module_info(handle, infoType, cValue)
	C.free(unsafe.Pointer(cValue))
}

//export Neb_Module_Deinit
func Neb_Module_Deinit(flags, reason int) C.int {
	Log(C.NSLOG_INFO_MESSAGE, fmt.Sprintf("[%s] deinitializing Callbacks\n", Name))
	deinitCallbacks()
	if NebModuleInitHook == nil {
		return NebOk
	} else {
		return C.int(NebModuleDeinitHook(flags, reason))
	}
}

func Dumper_Callback(callbacktype int, data unsafe.Pointer) int {
	Dump("Dumper_Callback:")
	Dump(callbacktype)
	Dump(data)
	return NebOk
}
