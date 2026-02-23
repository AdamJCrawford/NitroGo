package models

// router configuration structs
type RouterDynamicRouting struct {
	CommandString      string  `json:"commandstring,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	NodeID             int     `json:"nodeid,omitempty"`
	Output             string  `json:"output,omitempty"`
}
