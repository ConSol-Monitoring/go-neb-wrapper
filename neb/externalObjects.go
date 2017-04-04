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
	"github.com/ConSol/go-neb-wrapper/neb/structs"
	"unsafe"
)

func GetHosts() structs.Hostlist {
	hostList := structs.Hostlist{}
	hostPointer := unsafe.Pointer(C.host_list)
	for hostPointer != nil {
		hostList = append(hostList, structs.CastHost(hostPointer))
		host := *((*C.host)(hostPointer))
		hostPointer = unsafe.Pointer(host.next)
	}
	return hostList
}

func GetServices() structs.Servicelist {
	serviceList := structs.Servicelist{}
	servicePointer := unsafe.Pointer(C.service_list)
	for servicePointer != nil {
		serviceList = append(serviceList, structs.CastService(servicePointer))
		service := *((*C.service)(servicePointer))
		servicePointer = unsafe.Pointer(service.next)
	}
	return serviceList
}
