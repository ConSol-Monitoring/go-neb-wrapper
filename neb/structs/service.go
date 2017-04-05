package structs

/*

#cgo nagios3 CFLAGS: -DNAGIOS3 -I. -I${SRCDIR}/../../libs
#cgo nagios3 LDFLAGS: -Wl,-unresolved-symbols=ignore-all

#cgo nagios4 CFLAGS: -DNAGIOS4 -I. -I${SRCDIR}/../../libs
#cgo nagios4 LDFLAGS: -Wl,-unresolved-symbols=ignore-all

#cgo naemon CFLAGS: -DNAEMON -I.
#cgo naemon pkg-config: naemon

#include "../dependencies.h"

//This block is needed due to different naming schemes in Nagios3 and Nagios4/Naemon
#if defined(NAGIOS3)
char* ServiceGetCommand(void* data) {
	return ((service *)data)->service_check_command;
}
#elif defined(NAGIOS4) || defined(NAEMON)
char* ServiceGetCommand(void* data) {
	return ((service *)data)->check_command;
}
#endif

*/
import "C"
import (
	"unsafe"
)

//ServiceList is a list of services
type ServiceList []Service

//GenMetaHostAndServiceList will create a MetaHostAndServiceList
func (service ServiceList) GenMetaHostAndServiceList() MetaHostAndServiceList {
	meta := MetaHostAndServiceList{}
	for _, s := range service {
		meta = append(meta, s.MetaHostAndService)
	}
	return meta
}

//Service represents a nagios service
type Service struct {
	MetaHostAndService
	HostName    string
	Description string
	//CheckCommand contains args
	CheckCommand string
	//Command is the pure pluginname
	Command string
}

//CastService tries to cast the pointer to an go struct
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
