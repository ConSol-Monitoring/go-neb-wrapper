package main

import (
	"fmt"

	"github.com/ConSol/go-neb-wrapper/neb"
)

// Build contains the current git commit id
// compile passing -ldflags "-X main.Build <build sha1>" to set the id.
var Build string

//This is an example main file, which should demonstrate how to use the library.
func main() {
	neb.Title = "GO GO Neb Wrapper!"
	neb.Desc = "This is an example how to use the go neb wrapper"
	neb.License = "GPL v3"
	neb.Version = fmt.Sprintf("1.0.0 - %s", Build)
	neb.Author = "Philip Griesbacher / Sven Nierlein"
}
