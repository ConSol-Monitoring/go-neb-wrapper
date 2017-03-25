package structs

/*

#cgo nagios3 CFLAGS: -DNAGIOS3 -I. -I${SRCDIR}/../../libs
#cgo nagios3 LDFLAGS: -Wl,-unresolved-symbols=ignore-all

#cgo nagios4 CFLAGS: -DNAGIOS4 -I. -I${SRCDIR}/../../libs
#cgo nagios4 LDFLAGS: -Wl,-unresolved-symbols=ignore-all

#cgo naemon CFLAGS: -DNAEMON -I.
#cgo naemon pkg-config: naemon

#include "../dependencies.h"

*/
import "C"
import (
	"unsafe"
)

//ProcessStruct will be returned by process data
type Process struct {
	Type      int
	Flags     int
	Attr      int
	Timestamp Timeval
}

//CastProcess tries to cast the pointer to an go struct
func CastProcess(data unsafe.Pointer) Process {
	st := *((*C.struct_nebstruct_process_struct)(data))
	return Process{
		Type:      int(st._type),
		Flags:     int(st.flags),
		Attr:      int(st.attr),
		Timestamp: CastTimevalStruct(st.timestamp),
	}
}
