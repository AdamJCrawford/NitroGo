// Derived from Citrix ADC Nitro Go SDK (https://github.com/netscaler/adc-nitro-go)
// Originally licensed under the Apache License, Version 2.0
// Modifications by Adam Crawford, 2026

package models

type Kafkacluster struct {
	Name               string `json:"name,omitempty"`
	Activesvc          string `json:"activesvc,omitempty"`
	Totalsvc           string `json:"totalsvc,omitempty"`
	Topicname          string `json:"topicname,omitempty"`
	Numtopics          string `json:"numtopics,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Kafkaclusterbinding struct {
	Name string `json:"name,omitempty"`
}

type Kafkaclusterservicegroupbinding struct {
	Servicegroupname string `json:"servicegroupname,omitempty"`
	Name             string `json:"name,omitempty"`
}
