package models

// wasm configuration structs
type Wasmmodule struct {
	Count              float64 `json:"__count,omitempty"`
	Modulefile         string  `json:"modulefile,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Signaturefile      string  `json:"signaturefile,omitempty"`
}
