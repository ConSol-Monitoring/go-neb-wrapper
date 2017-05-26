package neb

/*

#cgo nagios3 CFLAGS: -DNAGIOS3 -I. -I${SRCDIR}/../libs
#cgo nagios3 LDFLAGS: -Wl,-unresolved-symbols=ignore-all

#cgo nagios4 CFLAGS: -DNAGIOS4 -I. -I${SRCDIR}/../libs
#cgo nagios4 LDFLAGS: -Wl,-unresolved-symbols=ignore-all

#cgo naemon CFLAGS: -DNAEMON -I.
#cgo naemon pkg-config: naemon

#include "dependencies.h"

#if defined(NAGIOS3)
void Log(int lvl, char* data) { write_to_all_logs(data, lvl); }
#elif defined(NAGIOS4)
void Log(int lvl, char* data) { write_to_all_logs(data, lvl); }
#elif defined(NAEMON)
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

const (
	//InfoMessage log level for messages
	InfoMessage = C.NSLOG_INFO_MESSAGE
)

//CoreLogPrefix change this to your custom prefix for the core log
var CoreLogPrefix = ""

func log(lvl int64, data string) {
	cs := C.CString(data)
	defer C.free(unsafe.Pointer(cs))
	C.Log(C.int(lvl), cs)
}

//CoreLog passes the given string to the core which logs in the configured file
func CoreLog(msg string) {
	log(InfoMessage, msg)
}

//CoreFLogWithoutPrefix will log to the corelog without prefix
func CoreFLogWithoutPrefix(format string, a ...interface{}) {
	CoreLog(fmt.Sprintf(format, a))
}

//CoreFLog will log to the corelog with prefix
func CoreFLog(format string, a ...interface{}) {
	if CoreLogPrefix == "" {
		CoreLogPrefix = fmt.Sprintf("[%s] ", Name)
	}
	CoreLog(fmt.Sprintf(CoreLogPrefix+format+"\n", a...))
}

//CoreDump dumps the object as string to the core log
func CoreDump(v interface{}) {
	spew.Config.Indent = "\t"
	spew.Config.MaxDepth = 20
	spew.Config.DisableMethods = true
	CoreFLog(spew.Sdump(v))
}

//Dump can be uses to dump any data to stderr
func Dump(v interface{}) {
	spew.Config.DisablePointerMethods = true
	spew.Config.Indent = "\t"
	spew.Config.MaxDepth = 20
	spew.Config.DisableMethods = true
	fmt.Fprintf(os.Stderr, spew.Sdump(v))
}
