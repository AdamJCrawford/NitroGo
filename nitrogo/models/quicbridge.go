// Derived from Citrix ADC Nitro Go SDK (https://github.com/netscaler/adc-nitro-go)
// Originally licensed under the Apache License, Version 2.0
// Modifications by Adam Crawford, 2026

package models

type Quicbridgeprofile struct {
	Name               string `json:"name,omitempty"`
	Routingalgorithm   string `json:"routingalgorithm,omitempty"`
	Serveridlength     int    `json:"serveridlength,omitempty"`
	Refcnt             string `json:"refcnt,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}
