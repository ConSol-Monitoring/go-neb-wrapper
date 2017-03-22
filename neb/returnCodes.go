package neb

const (
	//NebOk OK
	NebOk = 0
	//NebError Error
	NebError = -1
	// NebTrue True
	NebTrue = 1
	//NebFalse False
	NebFalse = 0
	//NeberrorNomem memory could not be allocated
	NeberrorNomem = 100
	//NeberrorNocallbackfunc no callback function was specified
	NeberrorNocallbackfunc = 200
	// NeberrorNocallbacklist callback list not initialized
	NeberrorNocallbacklist = 201
	//NeberrorCallbackbounds callback type was out of bounds
	NeberrorCallbackbounds = 202
	//NeberrorCallbacknotfound the callback could not be found
	NeberrorCallbacknotfound = 203
	//NeberrorNomodulehandle no module handle specified
	NeberrorNomodulehandle = 204
	//NeberrorBadmodulehandle bad module handle
	NeberrorBadmodulehandle = 205
	//NeberrorCallbackoverride module wants to override default Nagios handling of event
	NeberrorCallbackoverride = 206
	//NeberrorCallbackcancel module wants to cancel callbacks to other modules
	NeberrorCallbackcancel = 207
	//NeberrorNomodule no module was specified
	NeberrorNomodule = 300
	//NeberrorModinfobounds module info index was out of bounds
	NeberrorModinfobounds = 400
)
