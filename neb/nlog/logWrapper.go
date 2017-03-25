package nlog

/* This file contains the glue for calling naemon functions from
 * the golang neb module.
 */

/*

#cgo nagios3 CFLAGS: -DNAGIOS3 -I. -I${SRCDIR}/../../libs
#cgo nagios3 LDFLAGS: -Wl,-unresolved-symbols=ignore-all

#cgo nagios4 CFLAGS: -DNAGIOS4 -I. -I${SRCDIR}/../../libs
#cgo nagios4 LDFLAGS: -Wl,-unresolved-symbols=ignore-all

#cgo naemon CFLAGS: -DNAEMON -I.
#cgo naemon pkg-config: naemon

#include "../dependencies.h"

#if defined(NAGIOS3)
#include "../../libs/nagios3/nebcallbacks.h"
void Log(int lvl, char* data) { }
#elif defined(NAGIOS4)
#include "../../libs/nagios4/nebcallbacks.h"
void Log(int lvl, char* data) { }
#elif defined(NAEMON)
#include "naemon/naemon.h"
void Log(int lvl, char* data) { nm_log(lvl, data); }
#endif

*/
import "C"
import (
	"unsafe"

	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
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
