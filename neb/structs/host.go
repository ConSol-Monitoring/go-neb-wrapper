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
char* HostGetCommand(void* data) {
	return ((host *)data)->host_check_command;
}
#elif defined(NAGIOS4) || defined(NAEMON)
char* HostGetCommand(void* data) {
	return ((host *)data)->check_command;
}
#endif

*/
import "C"
import (
	"unsafe"
)

//Hostlist is a list of hosts
type Hostlist []Host

type Host struct {
	Name        string
	Alias       string
	DisplayName string
	//CheckCommand contains args
	CheckCommand string
	//Command is the pure pluginname
	Command       string
	ChecksEnabled int
	CheckType     int
	IsFlapping    int
}

//CastHost tries to cast the pointer to an go struct
func CastHost(data unsafe.Pointer) Host {
	st := *((*C.host)(data))
	command := C.GoString(C.HostGetCommand(data))
	return Host{
		Name:          C.GoString(st.name),
		Alias:         C.GoString(st.alias),
		DisplayName:   C.GoString(st.display_name),
		CheckCommand:  command,
		Command:       splitCommand(command),
		ChecksEnabled: int(st.checks_enabled),
		CheckType:     int(st.check_type),
		IsFlapping:    int(st.is_flapping),
	}
}
