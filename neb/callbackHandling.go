package neb

/*

#include "naemon/naemon.h"

extern nebmodule *neb_handle;

*/
import "C"
import "unsafe"

//Callback defines an function, which will be called by the core
type Callback func(int, unsafe.Pointer) int
type callbackMapping map[int][]Callback

var usedCallbackMapping = callbackMapping{}

//Generic_Callback this is a mapping function for C. Don't use it.
//export Generic_Callback
func Generic_Callback(callbackType int, data unsafe.Pointer) int {
	//TODO: parallel execution
	returnCode := C.NEB_OK
	if calls, ok := usedCallbackMapping[callbackType]; ok {
		for _, c := range calls {
			r := c(callbackType, data)
			if r != returnCode {
				return r
			}
		}
	}
	return returnCode
}

//AddCallback can be uses to register a function for a certain event
func AddCallback(callbackType int, callback Callback) {
	usedCallbackMapping[callbackType] = append(usedCallbackMapping[callbackType], callback)
}

func initCallbacks() {
	for callbackType := range usedCallbackMapping {
		RegisterGenericCallback(int64(callbackType))
	}
}

func deinitCallbacks() {
	for callbackType := range usedCallbackMapping {
		RegisterGenericCallback(int64(callbackType))
	}
}
