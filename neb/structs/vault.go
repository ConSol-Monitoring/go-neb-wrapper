package structs

/*

#cgo CFLAGS: -DNAEMON -I.
#cgo pkg-config: naemon

#include "../dependencies.h"

void NaemonFree(void* data) {
	nm_free(data);
}

*/
import "C"

import (
	"unsafe"
)

// GetVaultMacroName returns macro name from pointer
func GetVaultMacroName(data unsafe.Pointer) string {
	st := *((*C.nebstruct_vault_macro_data)(data))
	return C.GoString(st.macro_name)
}

// SetVaultMacroValue sets vault macro value
func SetVaultMacroValue(data unsafe.Pointer, value string) {
	st := *((*C.nebstruct_vault_macro_data)(data))
	C.NaemonFree(unsafe.Pointer(st.value))
	st.value = C.CString(value)
}
