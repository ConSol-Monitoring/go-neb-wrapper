package checkTypes

/*

#cgo nagios3 CFLAGS: -DNAGIOS3 -I. -I${SRCDIR}/../libs
#cgo nagios3 LDFLAGS: -Wl,-unresolved-symbols=ignore-all

#cgo nagios4 CFLAGS: -DNAGIOS4 -I. -I${SRCDIR}/../libs
#cgo nagios4 LDFLAGS: -Wl,-unresolved-symbols=ignore-all

#cgo naemon CFLAGS: -DNAEMON -I.
#cgo naemon pkg-config: naemon

#include "../dependencies.h"

*/
import "C"

const (
	//Active for passive checks
	Active = C.CHECK_TYPE_ACTIVE
	//Passive for active checks
	Passive = C.CHECK_TYPE_PASSIVE
	//Parent (active) check for the benefit of dependent object
	Parent = C.CHECK_TYPE_PARENT
	//File from spool files (yuck)
	File = C.CHECK_TYPE_FILE
	//Other for modules to use
	Other = C.CHECK_TYPE_OTHER
)

var checkTypeMapping = map[int]string{
	Active:  "Active",
	Passive: "Passive",
	Parent:  "Parent",
	File:    "File",
	Other:   "Other",
}

func CheckTypeToString(checkType int) string {
	if str, ok := checkTypeMapping[checkType]; ok {
		return str
	}
	return "Unknown CheckType"
}
