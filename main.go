package main

import (
	"fmt"
	"unsafe"

	"github.com/ConSol/go-neb-wrapper/neb"
	"github.com/ConSol/go-neb-wrapper/neb/naemon"
	"github.com/ConSol/go-neb-wrapper/neb/nlog"
)

// Build contains the current git commit id
// compile passing -ldflags "-X main.Build <build sha1>" to set the id.
var Build string

func genericCallback(callbackType int, data unsafe.Pointer, returnChannel chan int) {
	nlog.CoreLog(fmt.Sprintf("[%s] Generic Callback for %d\n", neb.Name, callbackType))
	switch callbackType {
	case naemon.ProcessData:
		neb.Dump(neb.CastProcessStruct(data))
	default:
		fmt.Println(fmt.Sprintf("[%s] Unkown CallbackType: %d\n", neb.Name, callbackType))
	}
	returnChannel <- neb.Ok
	close(returnChannel)
}

//This is an example main file, which should demonstrate how to use the library.
func init() {
	// just some information about your plugin
	neb.Title = "GO GO Neb Wrapper!"
	neb.Name = neb.Title
	neb.Desc = "This is an example how to use the go neb wrapper"
	neb.License = "GPL v3"
	neb.Version = fmt.Sprintf("1.0.0 - %s", Build)
	neb.Author = "Philip Griesbacher / Sven Nierlein"


	// this functions will be called every time a ProcessData event is triggered
	exampleCallback1 := func(callbackType int, data unsafe.Pointer, returnChannel chan int) {
		fmt.Printf("Example Callback1 for %d\n", callbackType)
		nlog.CoreLog(fmt.Sprintf("[%s] Example Callback1 logged for %d\n", neb.Name, callbackType))
		returnChannel <- neb.Ok
		close(returnChannel)
	}
	exampleCallback2 := func(callbackType int, data unsafe.Pointer, returnChannel chan int) {
		nlog.CoreLog(fmt.Sprintf("[%s] Example Callback2 logged for %d\n", neb.Name, callbackType))
		returnChannel <- neb.Ok
		close(returnChannel)
	}
	//There can be multiple of them
	neb.AddCallback(naemon.ProcessData, exampleCallback1)
	neb.AddCallback(naemon.ProcessData, exampleCallback2)


	neb.AddCallback(naemon.ProcessData, genericCallback)
	neb.AddCallback(naemon.ProcessData, genericCallback)

	//Init Hook Example
	neb.NebModuleInitHook = func(flags int, args string) int {
		fmt.Printf("Loading %s\n", neb.Title)
		fmt.Printf("Init flags: %d\n", flags)
		fmt.Printf("Init args: %s\n", args)
		return neb.Ok
	}

	//Deinit Hook Example
	neb.NebModuleDeinitHook = func(flags, reason int) int {
		fmt.Printf("Unloading %s\n", neb.Title)
		fmt.Printf("Deinit flags: %d\n", flags)
		fmt.Printf("Deinit reason: %d\n", reason)
		return neb.Ok
	}

}

//DON'T USE MAIN, IT WILL NEVER BE CALLED! USE CALLBACKS.
func main() {}
