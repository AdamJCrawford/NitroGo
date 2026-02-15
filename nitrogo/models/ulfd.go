package models

// ulfd configuration structs
type Ulfdserver struct {
	Count              float64 `json:"__count,omitempty"`
	Loggerip           string  `json:"loggerip,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	State              string  `json:"state,omitempty"`
}
