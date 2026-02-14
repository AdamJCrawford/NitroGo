// Derived from Citrix ADC Nitro Go SDK (https://github.com/netscaler/adc-nitro-go)
// Originally licensed under the Apache License, Version 2.0
// Modifications by Adam Crawford, 2026

package models

type Endpointinfo struct {
	Endpointkind       string `json:"endpointkind,omitempty"`
	Endpointname       string `json:"endpointname,omitempty"`
	Endpointmetadata   string `json:"endpointmetadata,omitempty"`
	Endpointlabelsjson string `json:"endpointlabelsjson,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}
