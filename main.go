package main

import (
	"fmt"
	"unsafe"

	"github.com/ConSol/go-neb-wrapper/neb"
	"github.com/ConSol/go-neb-wrapper/neb/naemon"
)

// Build contains the current git commit id
// compile passing -ldflags "-X main.Build <build sha1>" to set the id.
var Build string

//This is an example main file, which should demonstrate how to use the library.
func init() {
	// just some information about your plugin
	neb.Title = "GO GO Neb Wrapper!"
	neb.Name = neb.Title
	neb.Desc = "This is an example how to use the go neb wrapper"
	neb.License = "GPL v3"
	neb.Version = fmt.Sprintf("1.0.0 - %s", Build)
	neb.Author = "Philip Griesbacher / Sven Nierlein"

	// this function will be called every time a ProcessData event is triggered
	exampleCallback := func(int, unsafe.Pointer) int {
		fmt.Println("Example Callback")
		return neb.NebOk
	}
	neb.AddCallback(naemon.NebcallbackProcessData, exampleCallback)

	//Init Hook Example
	neb.NebModuleInitHook = func(flags int, args string) int {
		fmt.Printf("Loading %s\n", neb.Title)
		fmt.Printf("Init flags: %d\n", flags)
		fmt.Printf("Init args: %s\n", args)
		return neb.NebOk
	}

	//Deinit Hook Example
	neb.NebModuleDeinitHook = func(flags, reason int) int {
		fmt.Printf("Unloading %s\n", neb.Title)
		fmt.Printf("Deinit flags: %d\n", flags)
		fmt.Printf("Deinit reason: %d\n", reason)
		return neb.NebOk
	}

}
func main() {}
