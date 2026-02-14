// Derived from Citrix ADC Nitro Go SDK (https://github.com/netscaler/adc-nitro-go)
// Originally licensed under the Apache License, Version 2.0
// Modifications by Adam Crawford, 2026

package models

type Autoscaleaction struct {
	Name                 string `json:"name,omitempty"`
	Type                 string `json:"type,omitempty"`
	Profilename          string `json:"profilename,omitempty"`
	Parameters           string `json:"parameters,omitempty"`
	Vmdestroygraceperiod int    `json:"vmdestroygraceperiod,omitempty"`
	Quiettime            int    `json:"quiettime,omitempty"`
	Vserver              string `json:"vserver,omitempty"`
	Nextgenapiresource   string `json:"_nextgenapiresource,omitempty"`
}

type Autoscalepolicy struct {
	Name               string `json:"name,omitempty"`
	Rule               string `json:"rule,omitempty"`
	Action             string `json:"action,omitempty"`
	Comment            string `json:"comment,omitempty"`
	Logaction          string `json:"logaction,omitempty"`
	Newname            string `json:"newname,omitempty"`
	Hits               string `json:"hits,omitempty"`
	Undefhits          string `json:"undefhits,omitempty"`
	Priority           string `json:"priority,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Autoscalepolicybinding struct {
	Name string `json:"name,omitempty"`
}

type Autoscalepolicynstimerbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Autoscalepolicytimerbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Autoscaleprofile struct {
	Name               string `json:"name,omitempty"`
	Type               string `json:"type,omitempty"`
	Url                string `json:"url,omitempty"`
	Apikey             string `json:"apikey,omitempty"`
	Sharedsecret       string `json:"sharedsecret,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}
