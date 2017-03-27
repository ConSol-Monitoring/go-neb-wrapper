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

const (
	//HostcheckInitiate a check of the route to the host has been initiated
	HostcheckInitiate = C.NEBTYPE_HOSTCHECK_INITIATE
	//HostcheckProcessed the processed/final result of a host check
	HostcheckProcessed = C.NEBTYPE_HOSTCHECK_PROCESSED
)
