// Derived from Citrix ADC Nitro Go SDK (https://github.com/netscaler/adc-nitro-go)
// Originally licensed under the Apache License, Version 2.0
// Modifications by Adam Crawford, 2026

package models

type SCPolicy struct {
	Name              string `json:"name,omitempty"`
	URL               string `json:"url,omitempty"`
	Rule              string `json:"rule,omitempty"`
	Delay             int    `json:"delay,omitempty"`
	MaxConn           int    `json:"maxconn,omitempty"`
	Action            string `json:"action,omitempty"`
	AltContentSvcName string `json:"altcontentsvcname,omitempty"`
	AltContentPath    string `json:"altcontentpath,omitempty"`
}

type SCParameter struct {
	SessionLife int    `json:"sessionlife,omitempty"`
	VSR         string `json:"vsr,omitempty"`
	Builtin     string `json:"builtin,omitempty"`
	Feature     string `json:"feature,omitempty"`
}
