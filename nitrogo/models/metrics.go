// Derived from Citrix ADC Nitro Go SDK (https://github.com/netscaler/adc-nitro-go)
// Originally licensed under the Apache License, Version 2.0
// Modifications by Adam Crawford, 2026

package models

type Metricsprofileuservserverbinding struct {
	Entityname string `json:"entityname,omitempty"`
	Entitytype string `json:"entitytype,omitempty"`
	Name       string `json:"name,omitempty"`
}

type Metricsprofile struct {
	Name                   string `json:"name,omitempty"`
	Collector              string `json:"collector,omitempty"`
	Outputmode             string `json:"outputmode,omitempty"`
	Metrics                string `json:"metrics,omitempty"`
	Servemode              string `json:"servemode,omitempty"`
	Schemafile             string `json:"schemafile,omitempty"`
	Metricsexportfrequency int    `json:"metricsexportfrequency,omitempty"`
	Metricsauthtoken       string `json:"metricsauthtoken,omitempty"`
	Metricsendpointurl     string `json:"metricsendpointurl,omitempty"`
	Refcnt                 string `json:"refcnt,omitempty"`
	Nextgenapiresource     string `json:"_nextgenapiresource,omitempty"`
}

type Metricsprofilecrvserverbinding struct {
	Entityname string `json:"entityname,omitempty"`
	Entitytype string `json:"entitytype,omitempty"`
	Name       string `json:"name,omitempty"`
}

type Metricsprofilecsvserverbinding struct {
	Entityname string `json:"entityname,omitempty"`
	Entitytype string `json:"entitytype,omitempty"`
	Name       string `json:"name,omitempty"`
}

type Metricsprofilegslbvserverbinding struct {
	Entityname string `json:"entityname,omitempty"`
	Entitytype string `json:"entitytype,omitempty"`
	Name       string `json:"name,omitempty"`
}

type Metricsprofilevpnvserverbinding struct {
	Entityname string `json:"entityname,omitempty"`
	Entitytype string `json:"entitytype,omitempty"`
	Name       string `json:"name,omitempty"`
}

type Metricsprofileauthenticationvserverbinding struct {
	Entityname string `json:"entityname,omitempty"`
	Entitytype string `json:"entitytype,omitempty"`
	Name       string `json:"name,omitempty"`
}

type Metricsprofilebinding struct {
	Name string `json:"name,omitempty"`
}

type Metricsprofilelbvserverbinding struct {
	Entityname string `json:"entityname,omitempty"`
	Entitytype string `json:"entitytype,omitempty"`
	Name       string `json:"name,omitempty"`
}

type Metricsprofileservicebinding struct {
	Entityname string `json:"entityname,omitempty"`
	Entitytype string `json:"entitytype,omitempty"`
	Name       string `json:"name,omitempty"`
}

type Metricsprofileservicegroupbinding struct {
	Entityname string `json:"entityname,omitempty"`
	Entitytype string `json:"entitytype,omitempty"`
	Name       string `json:"name,omitempty"`
}
