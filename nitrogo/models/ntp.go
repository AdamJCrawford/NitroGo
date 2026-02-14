// Derived from Citrix ADC Nitro Go SDK (https://github.com/netscaler/adc-nitro-go)
// Originally licensed under the Apache License, Version 2.0
// Modifications by Adam Crawford, 2026

package models

type Ntpparam struct {
	Authentication     string `json:"authentication,omitempty"`
	Trustedkey         []int  `json:"trustedkey,omitempty"`
	Autokeylogsec      int    `json:"autokeylogsec"`
	Revokelogsec       int    `json:"revokelogsec"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Ntpserver struct {
	Serverip           string `json:"serverip,omitempty"`
	Servername         string `json:"servername,omitempty"`
	Minpoll            int    `json:"minpoll,omitempty"`
	Maxpoll            int    `json:"maxpoll,omitempty"`
	Autokey            bool   `json:"autokey,omitempty"`
	Key                int    `json:"key,omitempty"`
	Preferredntpserver string `json:"preferredntpserver,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Ntpstatus struct {
	Response           string `json:"response,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Ntpsync struct {
	State              string `json:"state,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}
