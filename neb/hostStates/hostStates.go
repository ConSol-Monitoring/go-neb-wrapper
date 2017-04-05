package hostStates

import "fmt"

const (
	Up = iota
	Down
	Unreachable
)

var stateTypeMapping = map[int]string{
	Up:          "Up",
	Down:        "Down",
	Unreachable: "Unreachable",
}

//StateTypeToString returns an string represent for the const
func StateTypeToString(stateType int) string {
	if str, ok := stateTypeMapping[stateType]; ok {
		return str
	}
	return fmt.Sprintf("Unknown StateType: %d", stateType)
}
