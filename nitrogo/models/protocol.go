// Derived from Citrix ADC Nitro Go SDK (https://github.com/netscaler/adc-nitro-go)
// Originally licensed under the Apache License, Version 2.0
// Modifications by Adam Crawford, 2026

package models

type Protocolhttpband struct {
	Reqbandsize        int    `json:"reqbandsize,omitempty"`
	Respbandsize       int    `json:"respbandsize,omitempty"`
	Type               string `json:"type,omitempty"`
	Nodeid             int    `json:"nodeid,omitempty"`
	Bandrange          string `json:"bandrange,omitempty"`
	Numberofbands      string `json:"numberofbands,omitempty"`
	Totalbandsize      string `json:"totalbandsize,omitempty"`
	Avgbandsize        string `json:"avgbandsize,omitempty"`
	Avgbandsizenew     string `json:"avgbandsizenew,omitempty"`
	Banddata           string `json:"banddata,omitempty"`
	Banddatanew        string `json:"banddatanew,omitempty"`
	Accesscount        string `json:"accesscount,omitempty"`
	Accessratio        string `json:"accessratio,omitempty"`
	Accessrationew     string `json:"accessrationew,omitempty"`
	Totals             string `json:"totals,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}
