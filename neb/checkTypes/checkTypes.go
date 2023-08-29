package checkTypes

/*

#cgo CFLAGS: -DNAEMON -I.
#cgo pkg-config: naemon

#include "../dependencies.h"

*/
import "C"

import (
	"fmt"
)

var (
	// Active for passive checks
	Active int
	// Passive for active checks
	Passive int
	// Parent (active) check for the benefit of dependent object
	Parent int
	// File from spool files (yuck)
	File int
	// Other for modules to use
	Other int
)

var checkTypeMapping map[int]string

func init() {
	Active = 0
	Passive = 1
	Parent = 2
	File = 3
	Other = 4

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
