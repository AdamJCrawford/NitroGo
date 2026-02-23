package models

// api configuration structs
type APISpecFile struct {
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	Overwrite          bool    `json:"overwrite,omitempty"`
	Src                string  `json:"src,omitempty"`
}

type APIProfileAPISpecBinding struct {
	APISpec string `json:"apispec,omitempty"`
	Name    string `json:"name,omitempty"`
}

type APISpec struct {
	Builtin            []string `json:"builtin,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Encrypted          bool     `json:"encrypted,omitempty"`
	File               string   `json:"file,omitempty"`
	Name               string   `json:"name,omitempty"`
	NextGenAPIResource string   `json:"_nextgenapiresource,omitempty"`
	NSAppVersion       string   `json:"nsappversion,omitempty"`
	Ready              string   `json:"ready,omitempty"`
	SkipValidation     string   `json:"skipvalidation,omitempty"`
	TypeField          string   `json:"type,omitempty"`
}

type APISpecSpecEndpointBinding struct {
	APIName     string `json:"apiname,omitempty"`
	APIService  string `json:"apiservice,omitempty"`
	HTTPMethod  string `json:"httpmethod,omitempty"`
	HTTPURLPath string `json:"httpurlpath,omitempty"`
	Name        string `json:"name,omitempty"`
}

type APIProfile struct {
	APIVisibility      string  `json:"apivisibility,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
}

type APIProfileBinding struct {
	APIProfileAPISpecBinding []interface{} `json:"apiprofile_apispec_binding,omitempty"`
	Name                     string        `json:"name,omitempty"`
}

type APISpecBinding struct {
	APISpecSpecEndpointBinding []interface{} `json:"apispec_specendpoint_binding,omitempty"`
	Name                       string        `json:"name,omitempty"`
}
