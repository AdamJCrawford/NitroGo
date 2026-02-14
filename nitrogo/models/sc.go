// Derived from Citrix ADC Nitro Go SDK (https://github.com/netscaler/adc-nitro-go)
// Originally licensed under the Apache License, Version 2.0
// Modifications by Adam Crawford, 2026

package models

type Scpolicy struct {
	Name              string `json:"name,omitempty"`
	Url               string `json:"url,omitempty"`
	Rule              string `json:"rule,omitempty"`
	Delay             int    `json:"delay,omitempty"`
	Maxconn           int    `json:"maxconn,omitempty"`
	Action            string `json:"action,omitempty"`
	Altcontentsvcname string `json:"altcontentsvcname,omitempty"`
	Altcontentpath    string `json:"altcontentpath,omitempty"`
}

type Scparameter struct {
	Sessionlife int    `json:"sessionlife,omitempty"`
	Vsr         string `json:"vsr,omitempty"`
	Builtin     string `json:"builtin,omitempty"`
	Feature     string `json:"feature,omitempty"`
}
