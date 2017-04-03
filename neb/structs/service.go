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
	"strings"
	"unsafe"
)

type Service struct {
	Description string
	HostName    string
	//CheckCommand contains args
	CheckCommand string
	//Command is the pure pluginname
	Command     string
	DisplayName string
}

//CastServiceCheck tries to cast the pointer to an go struct
func CastService(data unsafe.Pointer) Service {
	st := *((*C.service)(data))
	command := C.GoString(C.ServiceGetCommand(data))
	return Service{
		Description:  C.GoString(st.description),
		HostName:     C.GoString(st.host_name),
		CheckCommand: command,
		Command:      splitCommand(command),
		DisplayName:  C.GoString(st.display_name),
	}
}

func splitCommand(checkCommand string) string {
	if strings.Contains(checkCommand, "!") {
		return strings.Split(checkCommand, "!")[0]
	} else {
		return checkCommand
	}
}