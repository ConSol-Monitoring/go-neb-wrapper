package main

/* This file contains the glue for calling naemon functions from
 * the golang neb module.
 */

/*
#cgo pkg-config: naemon
#cgo CFLAGS: -I.
#include "neb_wrapper.h"

void Log(int lvl, char* data) { nm_log(lvl, data); }

void RegisterCallback(int type, void*callback) {
	neb_register_callback(type, neb_handle, 0, callback);
}

*/
import "C"
import "unsafe"

func Log(lvl int64, data string) {
	cs := C.CString(data)
	defer C.free(unsafe.Pointer(cs))
	C.Log(C.int(lvl), cs)
}

func RegisterCallback(callbacktype int64, callback func(int, unsafe.Pointer) int) {
	// TODO: use function pointer from args instead of hardcoded
	C.RegisterCallback(C.int(callbacktype), C.process_data_callback)
}
