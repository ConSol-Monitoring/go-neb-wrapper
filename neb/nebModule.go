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

var Name = ""
var Title = ""
var Author = ""
var Version = ""
var License = ""
var Desc = ""

//export Neb_Module_Init
func Neb_Module_Init(flags int, args string) C.int {
	handle := unsafe.Pointer(C.neb_handle)
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
	return C.NEB_OK
}

func setModuleInfo(handle unsafe.Pointer, infoType C.int, value string) {
	cValue := C.CString(value)
	C.neb_set_module_info(handle, infoType, cValue)
	C.free(unsafe.Pointer(cValue))
}

//export Neb_Module_Deinit
func Neb_Module_Deinit(flags, reason int) C.int {
	Log(C.NSLOG_INFO_MESSAGE, fmt.Sprintf("[%s] deinitializing\n", Name))
	deinitCallbacks()
	return C.NEB_OK
}

func Dumper_Callback(callbacktype int, data unsafe.Pointer) int {
	Dump("Dumper_Callback:")
	Dump(callbacktype)
	Dump(data)
	return (C.NEB_OK)
}
