// Derived from Citrix ADC Nitro Go SDK (https://github.com/netscaler/adc-nitro-go)
// Originally licensed under the Apache License, Version 2.0
// Modifications by Adam Crawford, 2026

package models

type Smppparam struct {
	Clientmode         string `json:"clientmode,omitempty"`
	Msgqueue           string `json:"msgqueue,omitempty"`
	Msgqueuesize       int    `json:"msgqueuesize,omitempty"`
	Addrton            int    `json:"addrton"`
	Addrnpi            int    `json:"addrnpi"`
	Addrrange          string `json:"addrrange,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Smppuser struct {
	Username           string `json:"username,omitempty"`
	Password           string `json:"password,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}
