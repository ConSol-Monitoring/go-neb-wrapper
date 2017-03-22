package neb

/*

#include "naemon/naemon.h"

extern nebmodule *neb_handle;

*/
import "C"
import "unsafe"

type Callback func(int, unsafe.Pointer) int
type CallbackMapping map[int][]Callback

var usedCallbackMapping = CallbackMapping{}

//export Generic_Callback
func Generic_Callback(callbackType int, data unsafe.Pointer) int {
	//TODO: parallel execution
	Dump("Generic")
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

func AddCallback(callbackType int, callback Callback) {
	usedCallbackMapping[callbackType] = append(usedCallbackMapping[callbackType], callback)
}

func InitCallbacks() {
	//TODO: change range
	for i := int64(0); i < 2; i++ {
		RegisterGenericCallback(i)
	}
}

func DeinitCallbacks() {
	//TODO: change range
	for i := int64(0); i < 2; i++ {
		DeregisterGenericCallback(i)
	}
}
