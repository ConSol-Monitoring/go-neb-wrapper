package structs

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
	"strings"
	"unsafe"
)

//MetaHostAndServiceList is a list of MetaHostAndService
type MetaHostAndServiceList []MetaHostAndService

//MetaHostAndService contains data for host and service
type MetaHostAndService struct {
	DisplayName              string
	ChecksEnabled            int
	CheckType                int
	IsFlapping               int
	CurrentState             int //0: up, 1: down, 2: unreachable. See stateTypes
	ScheduledDowntimeDepth   int
	PendingFlexDowntimeDepth int
}

//MetaType is needed to differentiate the void pointer in the constructor
type MetaType int

const (
	MetaHost MetaType = iota
	MetaService
)

//CastMetaHostAndService tries to cast the pointer to an go struct
func CastMetaHostAndService(data unsafe.Pointer, typ MetaType) MetaHostAndService {
	switch typ {
	case MetaHost:
		st := *((*C.host)(data))
		return MetaHostAndService{
			DisplayName:              C.GoString(st.display_name),
			ChecksEnabled:            int(st.checks_enabled),
			CheckType:                int(st.check_type),
			IsFlapping:               int(st.is_flapping),
			CurrentState:             int(st.current_state),
			ScheduledDowntimeDepth:   int(st.scheduled_downtime_depth),
			PendingFlexDowntimeDepth: int(st.pending_flex_downtime),
		}
	case MetaService:
		st := *((*C.service)(data))
		return MetaHostAndService{
			DisplayName:              C.GoString(st.display_name),
			ChecksEnabled:            int(st.checks_enabled),
			CheckType:                int(st.check_type),
			IsFlapping:               int(st.is_flapping),
			CurrentState:             int(st.current_state),
			ScheduledDowntimeDepth:   int(st.scheduled_downtime_depth),
			PendingFlexDowntimeDepth: int(st.pending_flex_downtime),
		}
	default:
		panic(fmt.Sprintf("The given type is not allowed: %d", typ))
	}
}

func splitCommand(checkCommand string) string {
	if strings.Contains(checkCommand, "!") {
		return strings.Split(checkCommand, "!")[0]
	} else {
		return checkCommand
	}
}
