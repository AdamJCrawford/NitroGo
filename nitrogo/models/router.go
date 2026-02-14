// Derived from Citrix ADC Nitro Go SDK (https://github.com/netscaler/adc-nitro-go)
// Originally licensed under the Apache License, Version 2.0
// Modifications by Adam Crawford, 2026

package models

type Routerdynamicrouting struct {
	Commandstring      string `json:"commandstring,omitempty"`
	Nodeid             int    `json:"nodeid,omitempty"`
	Output             string `json:"output,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}
