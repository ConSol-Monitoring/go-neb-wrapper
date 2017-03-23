package structs

/*
#include "naemon/naemon.h"
#cgo pkg-config: naemon
#cgo CFLAGS: -I.
*/
import "C"
import (
	"unsafe"
)

//ServiceCheck service check structure
type ServiceCheck struct {
	Process
	ServiceDescription string
	HostName           string
	CheckType          int
	CurrentAttempt     int
	MaxAttempts        int
	StateType          int
	State              int
	Timeout            int
	CommandName        string
	CommandArgs        string
	CommandLine        string
	StartTime          Timeval
	EndTime            Timeval
	EarlyTimeout       int
	ExecutionTime      float64
	Latency            float64
	ReturnCode         int
	Output             string
	LongOutput         string
	PerfData           string
	ObjectPtr          unsafe.Pointer
}

//CastServiceCheck tries to cast the pointer to an go struct
func CastServiceCheck(data unsafe.Pointer) ServiceCheck {
	st := *((*C.struct_nebstruct_service_check_struct)(data))
	return ServiceCheck{
		Process:            CastProcess(data),
		ServiceDescription: C.GoString(st.service_description),
		HostName:           C.GoString(st.host_name),
		CheckType:          int(st.check_type),
		CurrentAttempt:     int(st.current_attempt),
		MaxAttempts:        int(st.max_attempts),
		StateType:          int(st.state_type),
		State:              int(st.state),
		Timeout:            int(st.timeout),
		CommandName:        C.GoString(st.command_name),
		CommandArgs:        C.GoString(st.command_args),
		CommandLine:        C.GoString(st.command_line),
		StartTime:          CastTimevalStruct(st.start_time),
		EndTime:            CastTimevalStruct(st.end_time),
		EarlyTimeout:       int(st.early_timeout),
		ExecutionTime:      float64(st.execution_time),
		Latency:            float64(st.latency),
		ReturnCode:         int(st.return_code),
		Output:             C.GoString(st.output),
		LongOutput:         C.GoString(st.long_output),
		PerfData:           C.GoString(st.perf_data),
		ObjectPtr:          st.object_ptr,
	}
}
