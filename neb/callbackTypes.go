package neb

/*
#include "naemon/naemon.h"
*/
import "C"

const (
	ProcessData                   = C.NEBCALLBACK_PROCESS_DATA
	TimedEventData                = C.NEBCALLBACK_TIMED_EVENT_DATA
	LogData                       = C.NEBCALLBACK_LOG_DATA
	SystemCommandData             = C.NEBCALLBACK_SYSTEM_COMMAND_DATA
	EventHandlerData              = C.NEBCALLBACK_EVENT_HANDLER_DATA
	NotificationData              = C.NEBCALLBACK_NOTIFICATION_DATA
	ServiceCheckData              = C.NEBCALLBACK_SERVICE_CHECK_DATA
	HostCheckData                 = C.NEBCALLBACK_HOST_CHECK_DATA
	CommentData                   = C.NEBCALLBACK_COMMENT_DATA
	DowntimeData                  = C.NEBCALLBACK_DOWNTIME_DATA
	FlappingData                  = C.NEBCALLBACK_FLAPPING_DATA
	ProgramStatusData             = C.NEBCALLBACK_PROGRAM_STATUS_DATA
	HostStatusData                = C.NEBCALLBACK_HOST_STATUS_DATA
	ServiceStatusData             = C.NEBCALLBACK_SERVICE_STATUS_DATA
	AdaptiveProgramData           = C.NEBCALLBACK_ADAPTIVE_PROGRAM_DATA
	AdaptiveHostData              = C.NEBCALLBACK_ADAPTIVE_HOST_DATA
	AdaptiveServiceData           = C.NEBCALLBACK_ADAPTIVE_SERVICE_DATA
	ExternalCommandData           = C.NEBCALLBACK_EXTERNAL_COMMAND_DATA
	AggregatedStatusData          = C.NEBCALLBACK_AGGREGATED_STATUS_DATA
	RetentionData                 = C.NEBCALLBACK_RETENTION_DATA
	ContactNotificationData       = C.NEBCALLBACK_CONTACT_NOTIFICATION_DATA
	ContactNotificationMethodData = C.NEBCALLBACK_CONTACT_NOTIFICATION_METHOD_DATA
	AcknowledgementData           = C.NEBCALLBACK_ACKNOWLEDGEMENT_DATA
	StateChangeData               = C.NEBCALLBACK_STATE_CHANGE_DATA
	ContactStatusData             = C.NEBCALLBACK_CONTACT_STATUS_DATA
	AdaptiveContactData           = C.NEBCALLBACK_ADAPTIVE_CONTACT_DATA
)
