package neb

/*

#cgo CFLAGS: -DNAEMON -I.
#cgo pkg-config: naemon

#include "dependencies.h"

extern struct host *host_list;
extern struct service *service_list;

*/
import "C"

import (
	"unsafe"

	"github.com/ConSol-Monitoring/go-neb-wrapper/neb/structs"
)

func GetHosts() structs.HostList {
	hostList := structs.HostList{}
	hostPointer := unsafe.Pointer(C.host_list)
	for hostPointer != nil {
		hostList = append(hostList, structs.CastHost(hostPointer))
		host := *((*C.host)(hostPointer))
		hostPointer = unsafe.Pointer(host.next)
	}
	return hostList
}

func GetServices() structs.ServiceList {
	serviceList := structs.ServiceList{}
	servicePointer := unsafe.Pointer(C.service_list)
	for servicePointer != nil {
		serviceList = append(serviceList, structs.CastService(servicePointer))
		service := *((*C.service)(servicePointer))
		servicePointer = unsafe.Pointer(service.next)
	}
	return serviceList
}
