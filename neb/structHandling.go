package neb
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

//Timeval represents the c struct
type Timeval struct {
	TvSec  int
	TvUsec int
}

//ProcessStruct will be returned by process data
type ProcessStruct struct {
	Type      int
	Flags     int
	Attr      int
	Timestamp Timeval
}

//CastProcessStruct tries to cast the pointer to an go struct
func CastProcessStruct(data unsafe.Pointer) ProcessStruct {
	st := *((*C.struct_nebstruct_process_struct)(data))
	return ProcessStruct{
		Type:int(st._type),
		Flags:int(st.flags),
		Attr:int(st.attr),
		Timestamp:Timeval{
			TvSec:int(st.timestamp.tv_sec),
			TvUsec:int(st.timestamp.tv_usec),
		},
	}
}