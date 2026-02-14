// Derived from Citrix ADC Nitro Go SDK (https://github.com/netscaler/adc-nitro-go)
// Originally licensed under the Apache License, Version 2.0
// Modifications by Adam Crawford, 2026

package models

type Apiprofileapispecbinding struct {
	Apispec string `json:"apispec,omitempty"`
	Name    string `json:"name,omitempty"`
}

type Apiprofilebinding struct {
	Name string `json:"name,omitempty"`
}

type Apispec struct {
	Name               string `json:"name,omitempty"`
	File               string `json:"file,omitempty"`
	Type               string `json:"type,omitempty"`
	Skipvalidation     string `json:"skipvalidation,omitempty"`
	Encrypted          bool   `json:"encrypted,omitempty"`
	Builtin            string `json:"builtin,omitempty"`
	Ready              string `json:"ready,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Apispecbinding struct {
	Name string `json:"name,omitempty"`
}

type Apispecspecendpointbinding struct {
	Apiname     string `json:"apiname,omitempty"`
	Apiservice  string `json:"apiservice,omitempty"`
	Httpmethod  string `json:"httpmethod,omitempty"`
	Httpurlpath string `json:"httpurlpath,omitempty"`
	Name        string `json:"name,omitempty"`
}

type Apispecfile struct {
	Name               string `json:"name,omitempty"`
	Src                string `json:"src,omitempty"`
	Overwrite          bool   `json:"overwrite,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Apiprofile struct {
	Name               string `json:"name,omitempty"`
	Apivisibility      string `json:"apivisibility,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}
