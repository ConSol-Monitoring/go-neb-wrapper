package structs

/*

#cgo CFLAGS: -DNAEMON -I.
#cgo pkg-config: naemon

#include "../dependencies.h"

char* ServiceGetCommand(void* data) {
	return ((service *)data)->check_command;
}

*/
import "C"

import (
	"unsafe"
)

// ServiceList is a list of services
type ServiceList []Service

// GenMetaHostAndServiceList will create a MetaHostAndServiceList
func (service ServiceList) GenMetaHostAndServiceList() MetaHostAndServiceList {
	meta := MetaHostAndServiceList{}
	for _, s := range service {
		meta = append(meta, s.MetaHostAndService)
	}
	return meta
}

// Service represents a naemon service
type Service struct {
	MetaHostAndService
	HostName    string
	Description string
	// CheckCommand contains args
	CheckCommand string
	// Command is the pure pluginname
	Command string
}

// CastService tries to cast the pointer to an go struct
func CastService(data unsafe.Pointer) Service {
	st := *((*C.service)(data))
	command := C.GoString(C.ServiceGetCommand(data))
	return Service{
		MetaHostAndService: CastMetaHostAndService(data, MetaService),
		HostName:           C.GoString(st.host_name),
		Description:        C.GoString(st.description),
		CheckCommand:       command,
		Command:            splitCommand(command),
	}
}
