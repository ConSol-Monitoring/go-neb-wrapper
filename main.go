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
	neb.Title = "GO GO Neb Wrapper!"
	neb.Desc = "This is an example how to use the go neb wrapper"
	neb.License = "GPL v3"
	neb.Version = fmt.Sprintf("1.0.0 - %s", Build)
	neb.Author = "Philip Griesbacher / Sven Nierlein"
	exampleCallback := func(int, unsafe.Pointer) int {
		fmt.Println("Example Callback")
		return neb.NebOk
	}

	neb.AddCallback(naemon.NebcallbackProcessData, exampleCallback)
}
func main() {}
