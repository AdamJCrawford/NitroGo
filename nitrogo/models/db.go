// Derived from Citrix ADC Nitro Go SDK (https://github.com/netscaler/adc-nitro-go)
// Originally licensed under the Apache License, Version 2.0
// Modifications by Adam Crawford, 2026

package models

type Dbdbprofile struct {
	Name                   string `json:"name,omitempty"`
	Interpretquery         string `json:"interpretquery,omitempty"`
	Stickiness             string `json:"stickiness,omitempty"`
	Kcdaccount             string `json:"kcdaccount,omitempty"`
	Conmultiplex           string `json:"conmultiplex,omitempty"`
	Enablecachingconmuxoff string `json:"enablecachingconmuxoff,omitempty"`
	Refcnt                 string `json:"refcnt,omitempty"`
	Nextgenapiresource     string `json:"_nextgenapiresource,omitempty"`
}

type Dbuser struct {
	Username           string `json:"username,omitempty"`
	Password           string `json:"password,omitempty"`
	Loggedin           bool   `json:"loggedin,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}
