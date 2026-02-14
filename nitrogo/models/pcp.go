// Derived from Citrix ADC Nitro Go SDK (https://github.com/netscaler/adc-nitro-go)
// Originally licensed under the Apache License, Version 2.0
// Modifications by Adam Crawford, 2026

package models

type Pcpprofile struct {
	Name               string `json:"name,omitempty"`
	Mapping            string `json:"mapping,omitempty"`
	Peer               string `json:"peer,omitempty"`
	Minmaplife         int    `json:"minmaplife,omitempty"`
	Maxmaplife         int    `json:"maxmaplife,omitempty"`
	Announcemulticount int    `json:"announcemulticount"`
	Thirdparty         string `json:"thirdparty,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Pcpserver struct {
	Name               string `json:"name,omitempty"`
	Ipaddress          string `json:"ipaddress,omitempty"`
	Port               int    `json:"port,omitempty"`
	Pcpprofile         string `json:"pcpprofile,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Pcpmap struct {
	Nattype            string `json:"nattype,omitempty"`
	Pcpsrcip           string `json:"pcpsrcip,omitempty"`
	Subscrip           string `json:"subscrip,omitempty"`
	Pcpsrcport         string `json:"pcpsrcport,omitempty"`
	Pcpdstip           string `json:"pcpdstip,omitempty"`
	Pcpdstport         string `json:"pcpdstport,omitempty"`
	Pcpnatip           string `json:"pcpnatip,omitempty"`
	Pcpnatport         string `json:"pcpnatport,omitempty"`
	Pcpprotocol        string `json:"pcpprotocol,omitempty"`
	Pcpaddr            string `json:"pcpaddr,omitempty"`
	Pcpnounce          string `json:"pcpnounce,omitempty"`
	Pcprefcnt          string `json:"pcprefcnt,omitempty"`
	Pcplifetime        string `json:"pcplifetime,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}
