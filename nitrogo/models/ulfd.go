package models

// ulfd configuration structs
type ULFDServer struct {
	Count              float64 `json:"__count,omitempty"`
	LoggerIP           string  `json:"loggerip,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	State              string  `json:"state,omitempty"`
}
