package models

// router configuration structs
type Routerdynamicrouting struct {
	Commandstring      string  `json:"commandstring,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Nodeid             int     `json:"nodeid,omitempty"`
	Output             string  `json:"output,omitempty"`
}
