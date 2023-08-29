package neb

/* This file contains the glue for calling naemon functions from
 * the golang neb module.
 */

/*
#cgo CFLAGS: -DNAEMON -I.
#cgo pkg-config: naemon

#include "neb_wrapper.h"

extern nebmodule *neb_handle;

void RegisterCallback(int type, void*callback) {
	neb_register_callback(type, neb_handle, 0, callback);
}

void DeregisterCallback(int type, void*callback) {
	neb_deregister_callback(type, callback);
}

*/
import "C"

// RegisterGenericCallback capsules the c function neb_register_callback and passes the generic_callback handler
func RegisterGenericCallback(callbacktype int64) {
	C.RegisterCallback(C.int(callbacktype), C.generic_callback)
}

// DeregisterGenericCallback capsules the c function neb_deregister_callback and passes the generic_callback handler
func DeregisterGenericCallback(callbacktype int64) {
	C.DeregisterCallback(C.int(callbacktype), C.generic_callback)
}
