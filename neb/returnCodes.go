package neb

/*

#cgo CFLAGS: -DNAEMON
#cgo pkg-config: naemon

#include "dependencies.h"

*/
import "C"

const (
	// Ok OK
	Ok = C.NEB_OK
	// Error Error
	Error = C.NEB_ERROR
	// True True
	True = C.NEB_TRUE
	// False False
	False = C.NEB_FALSE
	// errorNomem memory could not be allocated
	ErrorNomem = C.NEBERROR_NOMEM
	// errorNocallbackfunc no callback function was specified
	ErrorNocallbackfunc = C.NEBERROR_NOCALLBACKFUNC
	// errorNocallbacklist callback list not initialized
	ErrorNocallbacklist = C.NEBERROR_NOCALLBACKLIST
	// errorCallbackbounds callback type was out of bounds
	ErrorCallbackbounds = C.NEBERROR_CALLBACKBOUNDS
	// errorCallbacknotfound the callback could not be found
	ErrorCallbacknotfound = C.NEBERROR_CALLBACKNOTFOUND
	// errorNomodulehandle no module handle specified
	ErrorNomodulehandle = C.NEBERROR_NOMODULEHANDLE
	// errorBadmodulehandle bad module handle
	ErrorBadmodulehandle = C.NEBERROR_BADMODULEHANDLE
	// errorCallbackoverride module wants to override default Naemon handling of event
	ErrorCallbackoverride = C.NEBERROR_CALLBACKOVERRIDE
	// errorCallbackcancel module wants to cancel callbacks to other modules
	ErrorCallbackcancel = C.NEBERROR_CALLBACKCANCEL
	// errorNomodule no module was specified
	ErrorNomodule = C.NEBERROR_NOMODULE
	// errorModinfobounds module info index was out of bounds
	ErrorModinfobounds = C.NEBERROR_MODINFOBOUNDS
)
