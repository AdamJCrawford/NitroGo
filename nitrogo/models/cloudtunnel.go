// Derived from Citrix ADC Nitro Go SDK (https://github.com/netscaler/adc-nitro-go)
// Originally licensed under the Apache License, Version 2.0
// Modifications by Adam Crawford, 2026

package models

type Cloudtunnelvserver struct {
	Name               string `json:"name,omitempty"`
	Servicetype        string `json:"servicetype,omitempty"`
	Listenpolicy       string `json:"listenpolicy,omitempty"`
	Listenpriority     int    `json:"listenpriority,omitempty"`
	State              string `json:"state,omitempty"`
	Effectivestate     string `json:"effectivestate,omitempty"`
	Type               string `json:"type,omitempty"`
	Ip                 string `json:"ip,omitempty"`
	Ipv46              string `json:"ipv46,omitempty"`
	Ippattern          string `json:"ippattern,omitempty"`
	Port               string `json:"port,omitempty"`
	Range              string `json:"range,omitempty"`
	Cachetype          string `json:"cachetype,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Cloudtunnelparameter struct {
	Controllerfqdn                 string `json:"controllerfqdn,omitempty"`
	Fqdn                           string `json:"fqdn,omitempty"`
	Resourcelocation               string `json:"resourcelocation,omitempty"`
	Subnetresourcelocationmappings string `json:"subnetresourcelocationmappings,omitempty"`
	Nextgenapiresource             string `json:"_nextgenapiresource,omitempty"`
}
