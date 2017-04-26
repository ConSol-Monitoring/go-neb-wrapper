package neb

/*

#cgo nagios3 CFLAGS: -DNAGIOS3 -I. -I${SRCDIR}/../libs
#cgo nagios3 LDFLAGS: -Wl,-unresolved-symbols=ignore-all

#cgo nagios4 CFLAGS: -DNAGIOS4 -I. -I${SRCDIR}/../libs
#cgo nagios4 LDFLAGS: -Wl,-unresolved-symbols=ignore-all

#cgo naemon CFLAGS: -DNAEMON -I.
#cgo naemon pkg-config: naemon

#include "dependencies.h"

*/
import "C"
import (
	"fmt"
	"sync"
	"time"
	"unsafe"
)

//Callback defines an function, which will be called by the core
//Return your result in the channel
type Callback func(int, unsafe.Pointer) int
type callbackMapping map[int][]Callback

var usedCallbackMapping = callbackMapping{}
var callbackMutex = sync.Mutex{}

//CallbackTimeout is the duration each callback has to return.
//This can be changed at the beginning.
var CallbackTimeout = time.Duration(10) * time.Millisecond

//Generic_Callback this is a mapping function for C, do not use it on your own.
//export Generic_Callback
func Generic_Callback(callbackType int, data unsafe.Pointer) int {
	startTime := time.Now()
	returnCode := Ok

	var callbacks []Callback
	var contains bool

	//Test if this call is interesting
	callbackMutex.Lock()
	callbacks, contains = usedCallbackMapping[callbackType]
	callbackMutex.Unlock()
	callbackAmount := len(callbacks)
	if !contains || callbackAmount == 0 {
		CoreFLog("We did not register for the callback %d", callbackType)
		return Error
	}

	switch GetCoreType() {
	case CoreNagios3, CoreNagios4:
		returnCode = serialCallbackHandling(callbackType, data, callbacks)
	case CoreNaemon:
		returnCode = concurrentCallbackHandling(callbackType, data, callbacks)
	default:
		CoreFLog("The coretype is not supported: %d", GetCoreType())
		returnCode = Error
	}

	if Stats != nil {
		select {
		case Stats.OverallCallbackDuration <- map[int]time.Duration{callbackType: time.Now().Sub(startTime)}:
		case <-time.After(CallbackTimeout):
			CoreFLog("Read your statstics data or don't set the global statistics object")
			return Ok
		}
	}
	return returnCode
}

//parallelCallbackHandling will execute all the callbacks in go routines and thous concurrent
func concurrentCallbackHandling(callbackType int, data unsafe.Pointer, callbacks []Callback) int {
	callbackAmount := len(callbacks)

	resultChannels := make([]chan int, callbackAmount)
	for i := range resultChannels {
		resultChannels[i] = make(chan int, 1)
	}

	//start all handlers for this callback
	for i, c := range callbacks {
		go func(result chan int, call Callback) {
			defer func() {
				if rec := recover(); rec != nil {
					CoreFLog("Cought panic: %s", fmt.Sprint(rec))
					result <- Error
				}
			}()
			result <- call(callbackType, data)
		}(resultChannels[i], c)
	}

	resultList := make([]int, callbackAmount)
	//Wait for all callbacks to signal that they are done and collect the returncodes
	var result int
	for i, c := range resultChannels {
		select {
		case result = <-c:
			resultList[i] = result
		case <-time.After(CallbackTimeout):
			return Error
		}
	}

	return testResults(resultList)
}

//serialCallbackHandling will call the callbacks one by one.
func serialCallbackHandling(callbackType int, data unsafe.Pointer, callbacks []Callback) int {
	defer func() {
		if rec := recover(); rec != nil {
			CoreFLog("Cought panic: %s", fmt.Sprint(rec))
		}
	}()
	callbackAmount := len(callbacks)
	resultList := make([]int, callbackAmount)
	for i, c := range callbacks {
		resultList[i] = c(callbackType, data)
	}
	return testResults(resultList)
}

//testResults test the returncodes if any is not OK return the error otherwise OK
func testResults(results []int) int {
	for _, r := range results {
		if r != Ok {
			return r
		}
	}
	return Ok
}

//AddCallback can be uses to register a function for a certain event
func AddCallback(callbackType int, callback Callback) {
	callbackMutex.Lock()
	usedCallbackMapping[callbackType] = append(usedCallbackMapping[callbackType], callback)
	if Stats != nil {
		select {
		case Stats.RegisteredCallbacksByType <- map[int]int{callbackType: len(usedCallbackMapping[callbackType])}:
		case <-time.After(CallbackTimeout):
			CoreFLog("Read your statstics data or don't set the global statistics object")
		}
	}
	callbackMutex.Unlock()
}

func initCallbacks() {
	callbackMutex.Lock()
	for callbackType := range usedCallbackMapping {
		RegisterGenericCallback(int64(callbackType))
	}
	callbackMutex.Unlock()
}

func deinitCallbacks() {
	callbackMutex.Lock()
	for callbackType := range usedCallbackMapping {
		DeregisterGenericCallback(int64(callbackType))
	}
	callbackMutex.Unlock()
}
