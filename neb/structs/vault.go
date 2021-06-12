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

// GetVaultMacroName returns the requested macro string
func GetVaultMacroName(data unsafe.Pointer) string {
	st := (*C.struct_nebstruct_vault_macro_struct)(data)
	return (C.GoString(st.macro_name))
}

// SetVaultMacroValue set the returned macro value
func SetVaultMacroValue(data unsafe.Pointer, val string) {
	st := (*C.struct_nebstruct_vault_macro_struct)(data)
	if st.value != nil {
		C.free(unsafe.Pointer(st.value))
	}
	st.value = C.CString(val)
	return
}
