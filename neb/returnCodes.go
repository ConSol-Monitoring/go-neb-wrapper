package neb

const (
	//Ok OK
	Ok = 0
	//Error Error
	Error = -1
	// True True
	True = 1
	//False False
	False = 0
	//errorNomem memory could not be allocated
	ErrorNomem = 100
	//errorNocallbackfunc no callback function was specified
	ErrorNocallbackfunc = 200
	// errorNocallbacklist callback list not initialized
	ErrorNocallbacklist = 201
	//errorCallbackbounds callback type was out of bounds
	ErrorCallbackbounds = 202
	//errorCallbacknotfound the callback could not be found
	ErrorCallbacknotfound = 203
	//errorNomodulehandle no module handle specified
	ErrorNomodulehandle = 204
	//errorBadmodulehandle bad module handle
	ErrorBadmodulehandle = 205
	//errorCallbackoverride module wants to override default Nagios handling of event
	ErrorCallbackoverride = 206
	//errorCallbackcancel module wants to cancel callbacks to other modules
	ErrorCallbackcancel = 207
	//errorNomodule no module was specified
	ErrorNomodule = 300
	//errorModinfobounds module info index was out of bounds
	ErrorModinfobounds = 400
)
