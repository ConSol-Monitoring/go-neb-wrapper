package neb

/*

#cgo CFLAGS: -DNAEMON -I.
#cgo pkg-config: naemon

#include "dependencies.h"

*/
import "C"

const (
	// ServicecheckInitiate a check of the route to the host has been initiated
	ServicecheckInitiate = C.NEBTYPE_SERVICECHECK_INITIATE
	// ServicecheckProcessed the processed/final result of a service check
	ServicecheckProcessed = C.NEBTYPE_SERVICECHECK_PROCESSED

	// HostcheckInitiate a check of the route to the host has been initiated
	HostcheckInitiate = C.NEBTYPE_HOSTCHECK_INITIATE
	// HostcheckProcessed the processed/final result of a host check
	HostcheckProcessed = C.NEBTYPE_HOSTCHECK_PROCESSED

	NotificationStart = C.NEBTYPE_NOTIFICATION_START
	NotificationEnd   = C.NEBTYPE_NOTIFICATION_END

	ContactNotificationStart = C.NEBTYPE_CONTACTNOTIFICATION_START
	ContactNotificationEnd   = C.NEBTYPE_CONTACTNOTIFICATION_END
)
