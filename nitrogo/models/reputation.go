// Derived from Citrix ADC Nitro Go SDK (https://github.com/netscaler/adc-nitro-go)
// Originally licensed under the Apache License, Version 2.0
// Modifications by Adam Crawford, 2026

package models

type Reputationsettings struct {
	Proxyserver        string `json:"proxyserver,omitempty"`
	Proxyport          int    `json:"proxyport,omitempty"`
	Proxyusername      string `json:"proxyusername,omitempty"`
	Proxypassword      string `json:"proxypassword,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}
