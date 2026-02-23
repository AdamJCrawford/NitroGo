package models

// dps configuration structs
type DPSParameter struct {
	Builtin            []string `json:"builtin,omitempty"`
	CustomerID         string   `json:"customerid,omitempty"`
	Deployment         string   `json:"deployment,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	NextGenAPIResource string   `json:"_nextgenapiresource,omitempty"`
	ServiceURL         string   `json:"serviceurl,omitempty"`
}
