package models

// endpoint configuration structs
type Endpointinfo struct {
	Count              float64 `json:"__count,omitempty"`
	Endpointkind       string  `json:"endpointkind,omitempty"`
	Endpointlabelsjson string  `json:"endpointlabelsjson,omitempty"`
	Endpointmetadata   string  `json:"endpointmetadata,omitempty"`
	Endpointname       string  `json:"endpointname,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
}
