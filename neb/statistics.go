package neb

import "time"

type Statistics struct {
	OverallCallbackDuration   chan map[int]time.Duration
	RegisteredCallbacksByType chan map[int]int
}

//Statistics if you create an Statistics object and place it here, make sure you are reading the channels, otherwise the whole module will be slowed down!
//This can be changed at the beginning.
var Stats *Statistics
