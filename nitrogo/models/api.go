package models

// api configuration structs
type Apispecfile struct {
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Overwrite          bool    `json:"overwrite,omitempty"`
	Src                string  `json:"src,omitempty"`
}

type ApiprofileApispecBinding struct {
	Apispec string `json:"apispec,omitempty"`
	Name    string `json:"name,omitempty"`
}

type Apispec struct {
	Builtin            []string `json:"builtin,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Encrypted          bool     `json:"encrypted,omitempty"`
	File               string   `json:"file,omitempty"`
	Name               string   `json:"name,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
	Nsappversion       string   `json:"nsappversion,omitempty"`
	Ready              string   `json:"ready,omitempty"`
	Skipvalidation     string   `json:"skipvalidation,omitempty"`
	TypeField          string   `json:"type,omitempty"`
}

type ApispecSpecendpointBinding struct {
	Apiname     string `json:"apiname,omitempty"`
	Apiservice  string `json:"apiservice,omitempty"`
	Httpmethod  string `json:"httpmethod,omitempty"`
	Httpurlpath string `json:"httpurlpath,omitempty"`
	Name        string `json:"name,omitempty"`
}

type Apiprofile struct {
	Apivisibility      string  `json:"apivisibility,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
}

type ApiprofileBinding struct {
	ApiprofileApispecBinding []interface{} `json:"apiprofile_apispec_binding,omitempty"`
	Name                     string        `json:"name,omitempty"`
}

type ApispecBinding struct {
	ApispecSpecendpointBinding []interface{} `json:"apispec_specendpoint_binding,omitempty"`
	Name                       string        `json:"name,omitempty"`
}
