// build with debug functions

package neb

import (
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
)

//Dump can be uses to dump any data to stderr
func Dump(v interface{}) {
	spew.Config.Indent = "\t"
	spew.Config.MaxDepth = 20
	spew.Config.DisableMethods = true
	fmt.Fprintf(os.Stderr, spew.Sdump(v))
}
