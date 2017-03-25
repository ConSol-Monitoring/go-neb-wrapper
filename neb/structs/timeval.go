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
