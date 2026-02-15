package models

// dps configuration structs
type Dpsparameter struct {
	Builtin            []string `json:"builtin,omitempty"`
	Customerid         string   `json:"customerid,omitempty"`
	Deployment         string   `json:"deployment,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
	Serviceurl         string   `json:"serviceurl,omitempty"`
}
