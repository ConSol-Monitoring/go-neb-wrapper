package structs

/*

#cgo CFLAGS: -DNAEMON -I.
#cgo pkg-config: naemon

#include "../dependencies.h"

char* HostGetCommand(void* data) {
	return ((host *)data)->check_command;
}

*/
import "C"

import (
	"unsafe"
)

// HostList is a list of hosts
type HostList []Host

// GenMetaHostAndServiceList will create a MetaHostAndServiceList
func (host HostList) GenMetaHostAndServiceList() MetaHostAndServiceList {
	meta := MetaHostAndServiceList{}
	for _, h := range host {
		meta = append(meta, h.MetaHostAndService)
	}
	return meta
}

// Host represents a naemon host
type Host struct {
	MetaHostAndService
	Name  string
	Alias string
	// CheckCommand contains args
	CheckCommand string
	// Command is the pure pluginname
	Command string
}

// CastHost tries to cast the pointer to an go struct
func CastHost(data unsafe.Pointer) Host {
	st := *((*C.host)(data))
	command := C.GoString(C.HostGetCommand(data))
	return Host{
		MetaHostAndService: CastMetaHostAndService(data, MetaService),
		Name:               C.GoString(st.name),
		Alias:              C.GoString(st.alias),
		CheckCommand:       command,
		Command:            splitCommand(command),
	}
}
