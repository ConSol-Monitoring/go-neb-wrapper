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
	"unsafe"
)

//ContactNotificationCheck notification check structure
type ContactNotificationCheck struct {
	Process
	NotificationType   int
	StartTime          Timeval
	EndTime            Timeval
	HostName           string
	ServiceDescription string
	ReasonType         int
	State              int
	Output             string
	AckAuthor          string
	AckData            string
	Escalated          int
	ObjectPtr          unsafe.Pointer
	ContactName        string
}

//CastContactNotificationCheck tries to cast the pointer to an go struct
func CastContactNotificationCheck(data unsafe.Pointer) ContactNotificationCheck {
	st := *((*C.struct_nebstruct_contact_notification_struct)(data))
	return ContactNotificationCheck{
		Process:            CastProcess(data),
		NotificationType:   int(st.notification_type),
		StartTime:          CastTimevalStruct(st.start_time),
		EndTime:            CastTimevalStruct(st.end_time),
		HostName:           C.GoString(st.host_name),
		ServiceDescription: C.GoString(st.service_description),
		ReasonType:         int(st.reason_type),
		State:              int(st.state),
		Output:             C.GoString(st.output),
		AckAuthor:          C.GoString(st.ack_author),
		AckData:            C.GoString(st.ack_data),
		Escalated:          int(st.escalated),
		ObjectPtr:          st.object_ptr,
		ContactName:        C.GoString(st.contact_name),
	}
}
