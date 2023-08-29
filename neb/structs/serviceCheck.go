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

// ServiceCheck service check structure
type ServiceCheck struct {
	Process
	Service
	CheckType      int
	CurrentAttempt int
	MaxAttempts    int
	StateType      int
	State          int
	Timeout        int
	StartTime      Timeval
	EndTime        Timeval
	EarlyTimeout   int
	ExecutionTime  float64
	Latency        float64
	ReturnCode     int
	Output         string
	LongOutput     string
	PerfData       string
	ObjectPtr      unsafe.Pointer
}

// CastServiceCheck tries to cast the pointer to an go struct
func CastServiceCheck(data unsafe.Pointer) ServiceCheck {
	st := *((*C.struct_nebstruct_service_check_struct)(data))
	return ServiceCheck{
		Process:        CastProcess(data),
		Service:        CastService(st.object_ptr),
		CheckType:      int(st.check_type),
		CurrentAttempt: int(st.current_attempt),
		MaxAttempts:    int(st.max_attempts),
		StateType:      int(st.state_type),
		State:          int(st.state),
		Timeout:        int(st.timeout),
		StartTime:      CastTimevalStruct(st.start_time),
		EndTime:        CastTimevalStruct(st.end_time),
		EarlyTimeout:   int(st.early_timeout),
		ExecutionTime:  float64(st.execution_time),
		Latency:        float64(st.latency),
		ReturnCode:     int(st.return_code),
		Output:         C.GoString(st.output),
		LongOutput:     C.GoString(st.long_output),
		PerfData:       C.GoString(st.perf_data),
		ObjectPtr:      st.object_ptr,
	}
}
