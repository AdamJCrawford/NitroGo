package models

// wasm configuration structs
type WasmModule struct {
	Count              float64 `json:"__count,omitempty"`
	ModuleFile         string  `json:"modulefile,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	SignatureFile      string  `json:"signaturefile,omitempty"`
}
