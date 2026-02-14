// Derived from Citrix ADC Nitro Go SDK (https://github.com/netscaler/adc-nitro-go)
// Originally licensed under the Apache License, Version 2.0
// Modifications by Adam Crawford, 2026

package models

type Appqoepolicybinding struct {
	Name string `json:"name,omitempty"`
}

type Appqoepolicylbvserverbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Bindpriority           int    `json:"bindpriority,omitempty"`
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Appqoepolicyvserverbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Bindpriority           uint32 `json:"bindpriority,omitempty"`
	Activepolicy           uint32 `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Appqoeaction struct {
	Name               string `json:"name,omitempty"`
	Priority           string `json:"priority,omitempty"`
	Respondwith        string `json:"respondwith,omitempty"`
	Customfile         string `json:"customfile,omitempty"`
	Altcontentsvcname  string `json:"altcontentsvcname,omitempty"`
	Altcontentpath     string `json:"altcontentpath,omitempty"`
	Polqdepth          int    `json:"polqdepth"`
	Priqdepth          int    `json:"priqdepth"`
	Maxconn            int    `json:"maxconn,omitempty"`
	Delay              int    `json:"delay,omitempty"`
	Dostrigexpression  string `json:"dostrigexpression,omitempty"`
	Dosaction          string `json:"dosaction,omitempty"`
	Tcpprofile         string `json:"tcpprofile,omitempty"`
	Retryonreset       string `json:"retryonreset,omitempty"`
	Retryontimeout     int    `json:"retryontimeout,omitempty"`
	Numretries         int    `json:"numretries"`
	Hits               string `json:"hits,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Appqoecustomresp struct {
	Src                string `json:"src,omitempty"`
	Name               string `json:"name,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Appqoeparameter struct {
	Sessionlife         int    `json:"sessionlife,omitempty"`
	Avgwaitingclient    int    `json:"avgwaitingclient"`
	Maxaltrespbandwidth int    `json:"maxaltrespbandwidth,omitempty"`
	Dosattackthresh     int    `json:"dosattackthresh"`
	Nextgenapiresource  string `json:"_nextgenapiresource,omitempty"`
}

type Appqoepolicy struct {
	Name               string `json:"name,omitempty"`
	Rule               string `json:"rule,omitempty"`
	Action             string `json:"action,omitempty"`
	Hits               string `json:"hits,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}
