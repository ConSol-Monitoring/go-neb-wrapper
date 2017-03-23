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
import (
	"unsafe"

	"github.com/davecgh/go-spew/spew"
	"fmt"
	"os"
)

func log(lvl int64, data string) {
	cs := C.CString(data)
	defer C.free(unsafe.Pointer(cs))
	C.Log(C.int(lvl), cs)
}

//CoreLog passes the given string to the core which logs in the configured file
func CoreLog(msg string) {
	log(NslogInfoMessage, msg)
}

//CoreDump dumps the object as string to the core log
func CoreDump(v interface{}) {
	spew.Config.Indent = "\t"
	spew.Config.MaxDepth = 20
	spew.Config.DisableMethods = true
	CoreLog(spew.Sdump(v))
}

//Dump can be uses to dump any data to stderr
func Dump(v interface{}) {
	spew.Config.DisablePointerMethods = true
	spew.Config.Indent = "\t"
	spew.Config.MaxDepth = 20
	spew.Config.DisableMethods = true
	fmt.Fprintf(os.Stderr, spew.Sdump(v))
}