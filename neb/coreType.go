package neb

/*
#cgo nagios3 CFLAGS: -DNAGIOS3 -I. -I${SRCDIR}/../libs
#cgo nagios3 LDFLAGS: -Wl,-unresolved-symbols=ignore-all

#cgo nagios4 CFLAGS: -DNAGIOS4 -I. -I${SRCDIR}/../libs
#cgo nagios4 LDFLAGS: -Wl,-unresolved-symbols=ignore-all

#cgo naemon CFLAGS: -DNAEMON -I.
#cgo naemon pkg-config: naemon

#include "dependencies.h"

*/
import "C"

const (
	//CoreNagios3 is an constant which represents the nagios3 core
	CoreNagios3 = C.CORE_NAGIOS3
	//CoreNagios4 is an constant which represents the nagios4 core
	CoreNagios4 = C.CORE_NAGIOS4
	//CoreNaemon is an constant which represents the naemon core
	CoreNaemon = C.CORE_NAEMON
)

var coreType = C.CORE_TYPE

func IsCoreNagios3() bool {
	return coreType == CoreNagios3
}

func IsCoreNagios4() bool {
	return coreType == CoreNagios4
}

func IsCoreNaemon() bool {
	return coreType == CoreNaemon
}

func CoreToString() string {
	switch coreType {
	case CoreNagios3:
		return "Nagios3"
	case CoreNagios4:
		return "Nagios4"
	case CoreNaemon:
		return "Naemon"
	default:
		return "Unknown"
	}
}
