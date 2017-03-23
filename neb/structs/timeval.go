package structs

/*
#include "naemon/naemon.h"
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

//CastTimeval tries to cast the pointer to an go struct
func CastTimeval(data unsafe.Pointer) Timeval {
	st := *((*C.struct_timeval)(data))
	return CastTimevalStruct(st)
}

//CastTimevalStruct cast the C struct to an go struct
func CastTimevalStruct(data C.struct_timeval) Timeval {
	return Timeval{
		TvSec:  int(data.tv_sec),
		TvUsec: int(data.tv_usec),
	}
}