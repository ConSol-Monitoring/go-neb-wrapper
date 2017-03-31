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

int hosts;
int services;

void count_hosts_services()
{
    hosts = 0;
    host *h = (host *)host_list;
    while (h) {
        hosts ++;
        h = h->next;
    }

    services = 0;
    service *s = (service *)service_list;
    while (s) {
        services ++;
        s = s->next;
    }
}

*/
import "C"

var (
	hosts    = -1
	services = -1
)

//TODO: Currently just working with Naemon

func GetHosts() int {
	if hosts == -1 {
		C.count_hosts_services()
		hosts = int(C.hosts)
	}
	return hosts
}
func GetServices() int {
	if services == -1 {
		C.count_hosts_services()
		services = int(C.services)
	}
	return services
}
