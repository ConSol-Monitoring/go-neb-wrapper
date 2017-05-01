package checkTypes

/*

#cgo nagios3 CFLAGS: -DNAGIOS3 -I. -I${SRCDIR}/../../libs
#cgo nagios3 LDFLAGS: -Wl,-unresolved-symbols=ignore-all

#cgo nagios4 CFLAGS: -DNAGIOS4 -I. -I${SRCDIR}/../../libs
#cgo nagios4 LDFLAGS: -Wl,-unresolved-symbols=ignore-all

#cgo naemon CFLAGS: -DNAEMON -I.
#cgo naemon pkg-config: naemon

#include "../dependencies.h"

*/
import "C"
import (
	"fmt"
	"github.com/ConSol/go-neb-wrapper/neb"
)

var (
	//Active for passive checks
	Active int
	//Passive for active checks
	Passive int
	//Parent (active) check for the benefit of dependent object
	Parent int
	//File from spool files (yuck)
	File int
	//Other for modules to use
	Other int
)

var checkTypeMapping map[int]string

func init() {
	switch neb.GetCoreType() {
	case neb.CoreNagios3:
		Active = C.SERVICE_CHECK_ACTIVE
		Passive = C.SERVICE_CHECK_PASSIVE
		Parent = -1
		File = -1
		Other = -1
	case neb.CoreNaemon, neb.CoreNagios4:
		Active = 0
		Passive = 1
		Parent = 2
		File = 3
		Other = 4
	}

	checkTypeMapping = map[int]string{
		Active:  "Active",
		Passive: "Passive",
		Parent:  "Parent",
		File:    "File",
		Other:   "Other",
	}
}

func CheckTypeToString(checkType int) string {
	if str, ok := checkTypeMapping[checkType]; ok {
		return str
	}
	return fmt.Sprintf("Unknown CheckType: %d", checkType)
}
