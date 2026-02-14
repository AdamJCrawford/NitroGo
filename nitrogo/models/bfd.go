// Derived from Citrix ADC Nitro Go SDK (https://github.com/netscaler/adc-nitro-go)
// Originally licensed under the Apache License, Version 2.0
// Modifications by Adam Crawford, 2026

package models

type Bfdsession struct {
	Localip                           string `json:"localip,omitempty"`
	Remoteip                          string `json:"remoteip,omitempty"`
	State                             string `json:"state,omitempty"`
	Localport                         string `json:"localport,omitempty"`
	Remoteport                        string `json:"remoteport,omitempty"`
	Minimumtransmitinterval           string `json:"minimumtransmitinterval,omitempty"`
	Negotiatedminimumtransmitinterval string `json:"negotiatedminimumtransmitinterval,omitempty"`
	Minimumreceiveinterval            string `json:"minimumreceiveinterval,omitempty"`
	Negotiatedminimumreceiveinterval  string `json:"negotiatedminimumreceiveinterval,omitempty"`
	Multiplier                        string `json:"multiplier,omitempty"`
	Remotemultiplier                  string `json:"remotemultiplier,omitempty"`
	Vlan                              string `json:"vlan,omitempty"`
	Localdiagnotic                    string `json:"localdiagnotic,omitempty"`
	Localdiscriminator                string `json:"localdiscriminator,omitempty"`
	Remotediscriminator               string `json:"remotediscriminator,omitempty"`
	Passive                           string `json:"passive,omitempty"`
	Multihop                          string `json:"multihop,omitempty"`
	Admindown                         string `json:"admindown,omitempty"`
	Originalownerpe                   string `json:"originalownerpe,omitempty"`
	Currentownerpe                    string `json:"currentownerpe,omitempty"`
	Ownernode                         string `json:"ownernode,omitempty"`
	Nextgenapiresource                string `json:"_nextgenapiresource,omitempty"`
}
