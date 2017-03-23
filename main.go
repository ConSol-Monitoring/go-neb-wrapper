package main

import (
	"fmt"
	"unsafe"

	"github.com/ConSol/go-neb-wrapper/neb"
	"github.com/ConSol/go-neb-wrapper/neb/nlog"
	"github.com/ConSol/go-neb-wrapper/neb/structs"
)

// Build contains the current git commit id
// compile passing -ldflags "-X main.Build <build sha1>" to set the id.
var Build string

//an example function how to handle multiple events at one function
func genericCallback(callbackType int, data unsafe.Pointer) int {
	nlog.CoreLog(fmt.Sprintf("[%s] Generic Callback for %d\n", neb.Name, callbackType))
	switch callbackType {
	case neb.ProcessData:
		nlog.Dump(structs.CastProcess(data))
	case neb.HostStatusData:
		nlog.Dump(structs.CastHostStatus(data))
	case neb.ServiceCheckData:
		nlog.Dump(structs.CastServiceCheck(data))
	case neb.HostCheckData:
		nlog.Dump(structs.CastHostCheck(data))
	default:
		fmt.Println(fmt.Sprintf("[%s] Unkown CallbackType: %d\n", neb.Name, callbackType))
	}
	return neb.Ok

}

//This is an example main file, which should demonstrate how to use the library.
func init() {
	// just some information about your plugin
	neb.Title = "GO GO Neb Wrapper! *\\o/*"
	neb.Name = neb.Title
	neb.Desc = "This is an example how to use the go neb wrapper"
	neb.License = "GPL v3"
	neb.Version = fmt.Sprintf("1.0.0 - %s", Build)
	neb.Author = "Philip Griesbacher / Sven Nierlein"

	// this functions will be called every time a ProcessData event is triggered
	exampleCallback1 := func(callbackType int, data unsafe.Pointer) int {
		fmt.Printf("Example Callback1 for %d\n", callbackType)
		nlog.CoreLog(fmt.Sprintf("[%s] Example Callback1 logged for %d\n", neb.Name, callbackType))
		return neb.Ok
	}
	exampleCallback2 := func(callbackType int, data unsafe.Pointer) int {
		nlog.CoreLog(fmt.Sprintf("[%s] Example Callback2 logged for %d\n", neb.Name, callbackType))
		return neb.Ok
	}
	//There can be multiple of them
	neb.AddCallback(neb.ProcessData, exampleCallback1)
	neb.AddCallback(neb.ProcessData, exampleCallback2)

	//And a lot more
	for _, t := range []int{neb.ProcessData, neb.HostStatusData, neb.ServiceCheckData, neb.HostCheckData} {
		neb.AddCallback(t, genericCallback)

	}

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
