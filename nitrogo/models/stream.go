// Derived from Citrix ADC Nitro Go SDK (https://github.com/netscaler/adc-nitro-go)
// Originally licensed under the Apache License, Version 2.0
// Modifications by Adam Crawford, 2026

package models

type Streamidentifiersessionbinding struct {
	Name string `json:"name,omitempty"`
}

type Streamidentifierstreamsessionbinding struct {
	Name             string `json:"name,omitempty"`
	Analyticsprofile string `json:"analyticsprofile,omitempty"`
}

type Streamselector struct {
	Name               string   `json:"name,omitempty"`
	Rule               []string `json:"rule,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
}

type Streamsession struct {
	Name string `json:"name,omitempty"`
}

type Streamidentifier struct {
	Name                    string `json:"name,omitempty"`
	Selectorname            string `json:"selectorname,omitempty"`
	Interval                int    `json:"interval,omitempty"`
	Samplecount             int    `json:"samplecount,omitempty"`
	Sort                    string `json:"sort,omitempty"`
	Snmptrap                string `json:"snmptrap,omitempty"`
	Appflowlog              string `json:"appflowlog,omitempty"`
	Trackackonlypackets     string `json:"trackackonlypackets,omitempty"`
	Tracktransactions       string `json:"tracktransactions,omitempty"`
	Maxtransactionthreshold int    `json:"maxtransactionthreshold,omitempty"`
	Mintransactionthreshold int    `json:"mintransactionthreshold,omitempty"`
	Acceptancethreshold     string `json:"acceptancethreshold,omitempty"`
	Breachthreshold         int    `json:"breachthreshold,omitempty"`
	Log                     string `json:"log,omitempty"`
	Loginterval             int    `json:"loginterval,omitempty"`
	Loglimit                int    `json:"loglimit,omitempty"`
	Rule                    string `json:"rule,omitempty"`
	Nextgenapiresource      string `json:"_nextgenapiresource,omitempty"`
}

type Streamidentifieranalyticsprofilebinding struct {
	Analyticsprofile string `json:"analyticsprofile,omitempty"`
	Name             string `json:"name,omitempty"`
}

type Streamidentifierbinding struct {
	Name string `json:"name,omitempty"`
}
