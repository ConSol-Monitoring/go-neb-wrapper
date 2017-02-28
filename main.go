package main

/*

#include "naemon/naemon.h"

extern nebmodule *neb_handle;

*/
import "C"
import (
	"fmt"
	"unsafe"
)

// Build contains the current git commit id
// compile passing -ldflags "-X main.Build <build sha1>" to set the id.
var Build string

const (
	// VERSION contains the actual module version
	VERSION = "1.0.0"
	// NAME defines the name of this project
	NAME = "go-naemon-broker-module"
)

func main() {}

//export Neb_Module_Init
func Neb_Module_Init(flags int, args string) C.int {
	Log(C.NSLOG_INFO_MESSAGE, fmt.Sprintf("[%s] initializing version %s (Build: %s)\n", NAME, VERSION, Build))

	C.neb_set_module_info(unsafe.Pointer(C.neb_handle), C.NEBMODULE_MODINFO_TITLE, C.CString("Go Example NEB Module"))
	C.neb_set_module_info(unsafe.Pointer(C.neb_handle), C.NEBMODULE_MODINFO_AUTHOR, C.CString("Sven Nierlein"))
	C.neb_set_module_info(unsafe.Pointer(C.neb_handle), C.NEBMODULE_MODINFO_TITLE, C.CString("Copyright (c) 2017 Sven Nierlein"))
	C.neb_set_module_info(unsafe.Pointer(C.neb_handle), C.NEBMODULE_MODINFO_VERSION, C.CString(VERSION))
	C.neb_set_module_info(unsafe.Pointer(C.neb_handle), C.NEBMODULE_MODINFO_LICENSE, C.CString("GPL v3"))
	C.neb_set_module_info(unsafe.Pointer(C.neb_handle), C.NEBMODULE_MODINFO_DESC, C.CString("This module just shows that it is possible to write neb modules in golang."))

	RegisterCallback(C.NEBCALLBACK_PROCESS_DATA, Process_Data_Callback)

	return C.NEB_OK
}

//export Neb_Module_Deinit
func Neb_Module_Deinit(flags, reason int) C.int {
	Log(C.NSLOG_INFO_MESSAGE, fmt.Sprintf("[%s] deinitializing\n", NAME))

	return C.NEB_OK
}

//export Process_Data_Callback
func Process_Data_Callback(callbacktype int, data unsafe.Pointer) int {
	Dump("Process_Data_Callback:")
	Dump(callbacktype)
	Dump((*C.struct_nebstruct_process_struct)(data))
	return (C.NEB_OK)
}
