package neb

/*

#cgo CFLAGS: -DNAEMON -I.
#cgo pkg-config: naemon

#include "dependencies.h"

void Log(int lvl, char* data) { nm_log(lvl, data); }

*/
import "C"

import (
	"fmt"
	"os"
	"unsafe"

	"github.com/davecgh/go-spew/spew"
)

const (
	// InfoMessage log level for messages
	InfoMessage = C.NSLOG_INFO_MESSAGE
)

// CoreLogPrefix change this to your custom prefix for the core log
var CoreLogPrefix = ""

func log(lvl int64, data string) {
	cs := C.CString(data)
	defer C.free(unsafe.Pointer(cs))
	C.Log(C.int(lvl), cs)
}

// CoreLog passes the given string to the core which logs in the configured file
func CoreLog(msg string) {
	log(InfoMessage, msg)
}

// CoreFLogWithoutPrefix will log to the corelog without prefix
func CoreFLogWithoutPrefix(format string, a ...interface{}) {
	CoreLog(fmt.Sprintf(format, a...))
}

// CoreFLog will log to the corelog with prefix
func CoreFLog(format string, a ...interface{}) {
	if CoreLogPrefix == "" {
		CoreLogPrefix = fmt.Sprintf("[%s] ", Name)
	}
	CoreLog(fmt.Sprintf(CoreLogPrefix+format+"\n", a...))
}

// CoreDump dumps the object as string to the core log
func CoreDump(v interface{}) {
	spew.Config.Indent = "\t"
	spew.Config.MaxDepth = 20
	spew.Config.DisableMethods = true
	CoreFLog(spew.Sdump(v))
}

// Dump can be uses to dump any data to stderr
func Dump(v interface{}) {
	spew.Config.DisablePointerMethods = true
	spew.Config.Indent = "\t"
	spew.Config.MaxDepth = 20
	spew.Config.DisableMethods = true
	fmt.Fprintf(os.Stderr, "%s", spew.Sdump(v))
}
