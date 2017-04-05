package serviceStates

import "fmt"

const (
	Ok = iota
	Warning
	Critical
	Unknown
)

var stateServiceStates = map[int]string{
	Ok:       "Ok",
	Warning:  "Warning",
	Critical: "Critical",
	Unknown:  "Unknown",
}

func ServiceStatesToString(state int) string {
	if str, ok := stateServiceStates[state]; ok {
		return str
	}
	return fmt.Sprintf("Unknown ServiceState: %d", state)
}
