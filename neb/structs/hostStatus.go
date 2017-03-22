package structs

/*
#include "naemon/naemon.h"
#include <stdlib.h>
#cgo pkg-config: naemon
#cgo CFLAGS: -I.
*/
import "C"
import (
	"unsafe"
)

type HostStatus struct {
	Type      int
	Flags     int
	Attr      int
	Timestamp Timeval
	ObjectPtr unsafe.Pointer
}

//CastHostStatus tries to cast the pointer to an go struct
func CastHostStatus(data unsafe.Pointer) HostStatus {
	st := *((*C.struct_nebstruct_host_status_struct)(data))
	return HostStatus{
		Type:  int(st._type),
		Flags: int(st.flags),
		Attr:  int(st.attr),
		Timestamp: Timeval{
			TvSec:  int(st.timestamp.tv_sec),
			TvUsec: int(st.timestamp.tv_usec),
		},
		ObjectPtr: st.object_ptr,
	}
}
