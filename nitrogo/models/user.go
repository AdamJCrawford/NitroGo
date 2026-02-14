// Derived from Citrix ADC Nitro Go SDK (https://github.com/netscaler/adc-nitro-go)
// Originally licensed under the Apache License, Version 2.0
// Modifications by Adam Crawford, 2026

package models

type Userprotocol struct {
	Name               string `json:"name,omitempty"`
	Transport          string `json:"transport,omitempty"`
	Extension          string `json:"extension,omitempty"`
	Comment            string `json:"comment,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Uservserver struct {
	Name                      string `json:"name,omitempty"`
	Userprotocol              string `json:"userprotocol,omitempty"`
	Ipaddress                 string `json:"ipaddress,omitempty"`
	Port                      int    `json:"port,omitempty"`
	Defaultlb                 string `json:"defaultlb,omitempty"`
	Params                    string `json:"Params,omitempty"`
	Comment                   string `json:"comment,omitempty"`
	State                     string `json:"state,omitempty"`
	Curstate                  string `json:"curstate,omitempty"`
	Value                     string `json:"value,omitempty"`
	Statechangetimesec        string `json:"statechangetimesec,omitempty"`
	Statechangetimemsec       string `json:"statechangetimemsec,omitempty"`
	Tickssincelaststatechange string `json:"tickssincelaststatechange,omitempty"`
	Nodefaultbindings         string `json:"nodefaultbindings,omitempty"`
	Nextgenapiresource        string `json:"_nextgenapiresource,omitempty"`
}
