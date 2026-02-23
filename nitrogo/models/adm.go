package models

// adm configuration structs
type ADMParameter struct {
	ADMServiceConnect  string `json:"admserviceconnect,omitempty"`
	LowTouchOnBoard    string `json:"lowtouchonboard,omitempty"`
	NextGenAPIResource string `json:"_nextgenapiresource,omitempty"`
}
