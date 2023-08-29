package structs

/*

#cgo CFLAGS: -DNAEMON -I.
#cgo pkg-config: naemon

#include "../dependencies.h"

*/
import "C"

import (
	"unsafe"
)

const (
	// NotificationNormal notification reason
	NotificationNormal = C.NOTIFICATION_NORMAL
	// NotificationAcknowledgement notification reason
	NotificationAcknowledgement = C.NOTIFICATION_ACKNOWLEDGEMENT
	// NotificationFlappingstart notification reason
	NotificationFlappingstart = C.NOTIFICATION_FLAPPINGSTART
	// NotificationFlappingstop notification reason
	NotificationFlappingstop = C.NOTIFICATION_FLAPPINGSTOP
	// NotificationFlappingdisabled notification reason
	NotificationFlappingdisabled = C.NOTIFICATION_FLAPPINGDISABLED
	// NotificationDowntimestart notification reason
	NotificationDowntimestart = C.NOTIFICATION_DOWNTIMESTART
	// NotificationDowntimeend notification reason
	NotificationDowntimeend = C.NOTIFICATION_DOWNTIMEEND
	// NotificationDowntimecancelled notification reason
	NotificationDowntimecancelled = C.NOTIFICATION_DOWNTIMECANCELLED
	// NotificationCustom notification reason
	NotificationCustom = C.NOTIFICATION_CUSTOM

	// HostNotification notification type
	HostNotification = C.HOST_NOTIFICATION
	// ServiceNotification notification type
	ServiceNotification = C.SERVICE_NOTIFICATION
)

var notificationReasonMapping = map[int]string{
	NotificationNormal:            "NotificationNormal",
	NotificationAcknowledgement:   "NotificationAcknowledgement",
	NotificationFlappingstart:     "NotificationFlappingstart",
	NotificationFlappingstop:      "NotificationFlappingstop",
	NotificationFlappingdisabled:  "NotificationFlappingdisabled",
	NotificationDowntimestart:     "NotificationDowntimestart",
	NotificationDowntimeend:       "NotificationDowntimeend",
	NotificationDowntimecancelled: "NotificationDowntimecancelled",
	NotificationCustom:            "NotificationCustom",
}

var notificationTypeMapping = map[int]string{
	HostNotification:    "HostNotification",
	ServiceNotification: "ServiceNotification",
}

// NotificationCheck notification check structure
type NotificationCheck struct {
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
	ContactsNotified   int
	ObjectPtr          unsafe.Pointer
}

// CastNotificationCheck tries to cast the pointer to an go struct
func CastNotificationCheck(data unsafe.Pointer) NotificationCheck {
	st := *((*C.struct_nebstruct_notification_struct)(data))
	return NotificationCheck{
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
		ContactsNotified:   int(st.contacts_notified),
		ObjectPtr:          st.object_ptr,
	}
}

func CastNotificationReasonToString(typ int) string {
	if str, ok := notificationReasonMapping[typ]; ok {
		return str
	}
	return "Unknown Reason"
}

func CastNotificationTypeToString(typ int) string {
	if str, ok := notificationTypeMapping[typ]; ok {
		return str
	}
	return "Unknown Type"
}
