package nlog

/* This file contains the glue for calling naemon functions from
 * the golang neb module.
 */

/*
#include "naemon/naemon.h"
#include <stdlib.h>
#cgo pkg-config: naemon
#cgo CFLAGS: -I.

void Log(int lvl, char* data) { nm_log(lvl, data); }

*/
import "C"
import "unsafe"

func log(lvl int64, data string) {
	cs := C.CString(data)
	defer C.free(unsafe.Pointer(cs))
	C.Log(C.int(lvl), cs)
}

func CoreLog(msg string) {
	log(NslogInfoMessage, msg)
}
