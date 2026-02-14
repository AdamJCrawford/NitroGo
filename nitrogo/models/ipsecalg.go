// Derived from Citrix ADC Nitro Go SDK (https://github.com/netscaler/adc-nitro-go)
// Originally licensed under the Apache License, Version 2.0
// Modifications by Adam Crawford, 2026

package models

type Ipsecalgprofile struct {
	Name               string `json:"name,omitempty"`
	Ikesessiontimeout  int    `json:"ikesessiontimeout,omitempty"`
	Espsessiontimeout  int    `json:"espsessiontimeout,omitempty"`
	Espgatetimeout     int    `json:"espgatetimeout,omitempty"`
	Connfailover       string `json:"connfailover,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Ipsecalgsession struct {
	Sourceipalg        string `json:"sourceip_alg,omitempty"`
	Natipalg           string `json:"natip_alg,omitempty"`
	Destipalg          string `json:"destip_alg,omitempty"`
	Sourceip           string `json:"sourceip,omitempty"`
	Natip              string `json:"natip,omitempty"`
	Destip             string `json:"destip,omitempty"`
	Spiin              string `json:"spiin,omitempty"`
	Spiout             string `json:"spiout,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}
