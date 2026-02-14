// Derived from Citrix ADC Nitro Go SDK (https://github.com/netscaler/adc-nitro-go)
// Originally licensed under the Apache License, Version 2.0
// Modifications by Adam Crawford, 2026

package models

type Ulfdserver struct {
	Loggerip           string `json:"loggerip,omitempty"`
	State              string `json:"state,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}
