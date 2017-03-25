package neb

/*

#cgo nagios3 CFLAGS: -DNAGIOS3 -I. -I${SRCDIR}/../libs
#cgo nagios3 LDFLAGS: -Wl,-unresolved-symbols=ignore-all

#cgo nagios4 CFLAGS: -DNAGIOS4 -I. -I${SRCDIR}/../libs
#cgo nagios4 LDFLAGS: -Wl,-unresolved-symbols=ignore-all

#cgo naemon CFLAGS: -DNAEMON -I.
#cgo naemon pkg-config: naemon

#include "dependencies.h"

extern struct host *host_list;
extern struct service *service_list;

*/
import "C"
import (
	"github.com/ConSol/go-neb-wrapper/neb/nlog"
)

//TODO: Nagios3 host_struct / Nagios4 host. Same with service

func GetHosts() {
	currentHost := (*(*C.struct_host)(C.host_list))
	for {
		nlog.Dump(C.GoString(currentHost.name))
		nextPtr := (*C.struct_host)(currentHost.next)
		if nextPtr == nil {
			break
		}
		currentHost = (*nextPtr)
	}
}
func GetServices() {
	currentService := (*(*C.struct_service)(C.service_list))
	for {
		nlog.Dump(C.GoString(currentService.display_name))
		nextPtr := (*C.struct_service)(currentService.next)
		if nextPtr == nil {
			break
		}
		currentService = (*nextPtr)
	}
}
