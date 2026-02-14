// Derived from Citrix ADC Nitro Go SDK (https://github.com/netscaler/adc-nitro-go)
// Originally licensed under the Apache License, Version 2.0
// Modifications by Adam Crawford, 2026

package models

type Dospolicy struct {
	Name          string `json:"name,omitempty"`
	Qdepth        int    `json:"qdepth,omitempty"`
	Cltdetectrate int    `json:"cltdetectrate,omitempty"`
}
