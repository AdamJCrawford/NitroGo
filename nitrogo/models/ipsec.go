// Derived from Citrix ADC Nitro Go SDK (https://github.com/netscaler/adc-nitro-go)
// Originally licensed under the Apache License, Version 2.0
// Modifications by Adam Crawford, 2026

package models

type Ipsecparameter struct {
	Ikeversion            string   `json:"ikeversion,omitempty"`
	Encalgo               []string `json:"encalgo,omitempty"`
	Hashalgo              []string `json:"hashalgo,omitempty"`
	Lifetime              int      `json:"lifetime,omitempty"`
	Livenesscheckinterval int      `json:"livenesscheckinterval,omitempty"`
	Replaywindowsize      int      `json:"replaywindowsize,omitempty"`
	Ikeretryinterval      int      `json:"ikeretryinterval,omitempty"`
	Perfectforwardsecrecy string   `json:"perfectforwardsecrecy,omitempty"`
	Retransmissiontime    int      `json:"retransmissiontime,omitempty"`
	Responderonly         string   `json:"responderonly,omitempty"`
	Nextgenapiresource    string   `json:"_nextgenapiresource,omitempty"`
}

type Ipsecprofile struct {
	Name                  string   `json:"name,omitempty"`
	Ikeversion            string   `json:"ikeversion,omitempty"`
	Encalgo               []string `json:"encalgo,omitempty"`
	Hashalgo              []string `json:"hashalgo,omitempty"`
	Lifetime              int      `json:"lifetime,omitempty"`
	Psk                   string   `json:"psk,omitempty"`
	Publickey             string   `json:"publickey,omitempty"`
	Privatekey            string   `json:"privatekey,omitempty"`
	Peerpublickey         string   `json:"peerpublickey,omitempty"`
	Livenesscheckinterval int      `json:"livenesscheckinterval,omitempty"`
	Replaywindowsize      int      `json:"replaywindowsize,omitempty"`
	Ikeretryinterval      int      `json:"ikeretryinterval,omitempty"`
	Retransmissiontime    int      `json:"retransmissiontime,omitempty"`
	Perfectforwardsecrecy string   `json:"perfectforwardsecrecy,omitempty"`
	Responderonly         string   `json:"responderonly,omitempty"`
	Builtin               string   `json:"builtin,omitempty"`
	Feature               string   `json:"feature,omitempty"`
	Nextgenapiresource    string   `json:"_nextgenapiresource,omitempty"`
}
