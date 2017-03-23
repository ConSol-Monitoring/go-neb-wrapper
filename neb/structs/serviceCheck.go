package structs

/*
#include "naemon/naemon.h"
#cgo pkg-config: naemon
#cgo CFLAGS: -I.
*/
import "C"
import (
	"unsafe"
	"github.com/ConSol/go-neb-wrapper/neb/nlog"
)
/*
(structs._Ctype_struct_nebstruct_service_check_struct) {
    _type: (structs._Ctype_int) 704,
    flags: (structs._Ctype_int) 0,
    attr: (structs._Ctype_int) 0,
    _: ([4]uint8) (len=4 cap=4) {
        00000000  fe 7f 00 00                                       |....|
    },
    timestamp: (structs._Ctype_struct_timeval) {
        tv_sec: (structs._Ctype___time_t) 1490254938,
        tv_usec: (structs._Ctype___suseconds_t) 346814
    },
    host_name: (*structs._Ctype_char)(0x2172670)(104),
    service_description: (*structs._Ctype_char)(0x2172b90)(99),
    check_type: (structs._Ctype_int) 0,
    current_attempt: (structs._Ctype_int) 1,
    max_attempts: (structs._Ctype_int) 3,
    state_type: (structs._Ctype_int) 1,
    state: (structs._Ctype_int) 0,
    timeout: (structs._Ctype_int) 0,
    command_name: (*structs._Ctype_char)(0x2288c80)(99),
    command_args: (*structs._Ctype_char)(0x2288c91)(49),
    command_line: (*structs._Ctype_char)(<nil>),
    start_time: (structs._Ctype_struct_timeval) {
        tv_sec: (structs._Ctype___time_t) 0,
        tv_usec: (structs._Ctype___suseconds_t) 0
    },
    end_time: (structs._Ctype_struct_timeval) {
        tv_sec: (structs._Ctype___time_t) 0,
        tv_usec: (structs._Ctype___suseconds_t) 0
    },
    early_timeout: (structs._Ctype_int) 0,
    _: ([4]uint8) (len=4 cap=4) {
        00000000  00 00 00 00                                       |....|
    },
    execution_time: (structs._Ctype_double) 0,
    latency: (structs._Ctype_double) 0.3468019962310791,
    return_code: (structs._Ctype_int) 0,
    _: ([4]uint8) (len=4 cap=4) {
        00000000  d1 7f 00 00                                       |....|
    },
    output: (*structs._Ctype_char)(0x2272310)(79),
    long_output: (*structs._Ctype_char)(0x2272360)(0),
    perf_data: (*structs._Ctype_char)(0x2271ec0)(108),
    check_result_ptr: (*structs._Ctype_struct_check_result)(<nil>),
    object_ptr: (unsafe.Pointer) 0x2172860
}
 */

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
	nlog.Dump(st)
	return ServiceCheck{
		Process:CastProcess(data),
		ServiceDescription:C.GoString(st.service_description),
		HostName:C.GoString(st.host_name),
		CheckType:int(st.check_type),
		CurrentAttempt:int(st.current_attempt),
		MaxAttempts:int(st.max_attempts),
		StateType:int(st.state_type),
		State:int(st.state),
		Timeout:int(st.timeout),
		CommandName:C.GoString(st.command_name),
		CommandArgs:C.GoString(st.command_args),
		CommandLine:C.GoString(st.command_line),
		StartTime:CastTimevalStruct(st.start_time),
		EndTime:CastTimevalStruct(st.end_time),
		EarlyTimeout:int(st.early_timeout),
		ExecutionTime:float64(st.execution_time),
		Latency:float64(st.latency),
		ReturnCode:int(st.return_code),
		Output:C.GoString(st.output),
		LongOutput:C.GoString(st.long_output),
		PerfData:C.GoString(st.perf_data),
		ObjectPtr : st.object_ptr,
	}
}