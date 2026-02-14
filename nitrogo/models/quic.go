// Derived from Citrix ADC Nitro Go SDK (https://github.com/netscaler/adc-nitro-go)
// Originally licensed under the Apache License, Version 2.0
// Modifications by Adam Crawford, 2026

package models

type Quicprofile struct {
	Name                           string `json:"name,omitempty"`
	Ackdelayexponent               int    `json:"ackdelayexponent,omitempty"`
	Activeconnectionidlimit        int    `json:"activeconnectionidlimit,omitempty"`
	Activeconnectionmigration      string `json:"activeconnectionmigration,omitempty"`
	Congestionctrlalgorithm        string `json:"congestionctrlalgorithm,omitempty"`
	Initialmaxdata                 int    `json:"initialmaxdata,omitempty"`
	Initialmaxstreamdatabidilocal  int    `json:"initialmaxstreamdatabidilocal,omitempty"`
	Initialmaxstreamdatabidiremote int    `json:"initialmaxstreamdatabidiremote,omitempty"`
	Initialmaxstreamdatauni        int    `json:"initialmaxstreamdatauni,omitempty"`
	Initialmaxstreamsbidi          int    `json:"initialmaxstreamsbidi,omitempty"`
	Initialmaxstreamsuni           int    `json:"initialmaxstreamsuni,omitempty"`
	Maxackdelay                    int    `json:"maxackdelay,omitempty"`
	Maxidletimeout                 int    `json:"maxidletimeout,omitempty"`
	Maxudpdatagramsperburst        int    `json:"maxudpdatagramsperburst,omitempty"`
	Maxudppayloadsize              int    `json:"maxudppayloadsize,omitempty"`
	Newtokenvalidityperiod         int    `json:"newtokenvalidityperiod,omitempty"`
	Retrytokenvalidityperiod       int    `json:"retrytokenvalidityperiod,omitempty"`
	Statelessaddressvalidation     string `json:"statelessaddressvalidation,omitempty"`
	Refcnt                         string `json:"refcnt,omitempty"`
	Builtin                        string `json:"builtin,omitempty"`
	Feature                        string `json:"feature,omitempty"`
	Nextgenapiresource             string `json:"_nextgenapiresource,omitempty"`
}

type Quicparam struct {
	Quicsecrettimeout  int    `json:"quicsecrettimeout,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}
