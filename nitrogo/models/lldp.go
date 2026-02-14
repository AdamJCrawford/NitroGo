// Derived from Citrix ADC Nitro Go SDK (https://github.com/netscaler/adc-nitro-go)
// Originally licensed under the Apache License, Version 2.0
// Modifications by Adam Crawford, 2026

package models

type Lldpneighbors struct {
	Ifnum              string `json:"ifnum,omitempty"`
	Nodeid             int    `json:"nodeid,omitempty"`
	Chassisidsubtype   string `json:"chassisidsubtype,omitempty"`
	Chassisid          string `json:"chassisid,omitempty"`
	Portidsubtype      string `json:"portidsubtype,omitempty"`
	Portid             string `json:"portid,omitempty"`
	Ttl                string `json:"ttl,omitempty"`
	Portdescription    string `json:"portdescription,omitempty"`
	Sys                string `json:"sys,omitempty"`
	Sysdesc            string `json:"sysdesc,omitempty"`
	Mgmtaddresssubtype string `json:"mgmtaddresssubtype,omitempty"`
	Mgmtaddress        string `json:"mgmtaddress,omitempty"`
	Iftype             string `json:"iftype,omitempty"`
	Ifnumber           string `json:"ifnumber,omitempty"`
	Vlan               string `json:"vlan,omitempty"`
	Vlanid             string `json:"vlanid,omitempty"`
	Portprotosupported string `json:"portprotosupported,omitempty"`
	Portprotoenabled   string `json:"portprotoenabled,omitempty"`
	Portprotoid        string `json:"portprotoid,omitempty"`
	Portvlanid         string `json:"portvlanid,omitempty"`
	Protocolid         string `json:"protocolid,omitempty"`
	Linkaggrcapable    string `json:"linkaggrcapable,omitempty"`
	Linkaggrenabled    string `json:"linkaggrenabled,omitempty"`
	Linkaggrid         string `json:"linkaggrid,omitempty"`
	Flag               string `json:"flag,omitempty"`
	Syscapabilities    string `json:"syscapabilities,omitempty"`
	Syscapenabled      string `json:"syscapenabled,omitempty"`
	Autonegsupport     string `json:"autonegsupport,omitempty"`
	Autonegenabled     string `json:"autonegenabled,omitempty"`
	Autonegadvertised  string `json:"autonegadvertised,omitempty"`
	Autonegmautype     string `json:"autonegmautype,omitempty"`
	Mtu                string `json:"mtu,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Lldpparam struct {
	Holdtimetxmult     int    `json:"holdtimetxmult,omitempty"`
	Timer              int    `json:"timer,omitempty"`
	Mode               string `json:"mode,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}
