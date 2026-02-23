package models

// endpoint configuration structs
type EndpointInfo struct {
	Count              float64 `json:"__count,omitempty"`
	EndpointKind       string  `json:"endpointkind,omitempty"`
	EndpointLabelsJSON string  `json:"endpointlabelsjson,omitempty"`
	EndpointMetadata   string  `json:"endpointmetadata,omitempty"`
	EndpointName       string  `json:"endpointname,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
}
