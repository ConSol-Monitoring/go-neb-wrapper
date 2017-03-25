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

type HostStatus struct {
	Process
	ObjectPtr unsafe.Pointer
}

//CastHostStatus tries to cast the pointer to an go struct
func CastHostStatus(data unsafe.Pointer) HostStatus {
	st := *((*C.struct_nebstruct_host_status_struct)(data))
	return HostStatus{
		Process: Process{
			Type:      int(st._type),
			Flags:     int(st.flags),
			Attr:      int(st.attr),
			Timestamp: CastTimevalStruct(st.timestamp),
		},
		ObjectPtr: st.object_ptr,
	}
}
