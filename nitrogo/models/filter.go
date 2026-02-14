// Derived from Citrix ADC Nitro Go SDK (https://github.com/netscaler/adc-nitro-go)
// Originally licensed under the Apache License, Version 2.0
// Modifications by Adam Crawford, 2026

package models

type Filterpolicyfilterglobalbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Filterprebodyinjection struct {
	Prebody   string `json:"prebody,omitempty"`
	Systemiid string `json:"systemiid,omitempty"`
}

type Filterglobalbinding struct {
}

type Filterglobalfilterpolicybinding struct {
	Policyname string `json:"policyname,omitempty"`
	Priority   int    `json:"priority,omitempty"`
	State      string `json:"state,omitempty"`
}

type Filterpolicy struct {
	Name      string `json:"name,omitempty"`
	Rule      string `json:"rule,omitempty"`
	Reqaction string `json:"reqaction,omitempty"`
	Resaction string `json:"resaction,omitempty"`
	Hits      string `json:"hits,omitempty"`
}

type Filterpolicylbvserverbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Filterhtmlinjectionparameter struct {
	Rate          int    `json:"rate,omitempty"`
	Frequency     int    `json:"frequency,omitempty"`
	Strict        string `json:"strict,omitempty"`
	Htmlsearchlen int    `json:"htmlsearchlen,omitempty"`
	Builtin       string `json:"builtin,omitempty"`
	Feature       string `json:"feature,omitempty"`
}

type Filterhtmlinjectionvariable struct {
	Variable string `json:"variable,omitempty"`
	Value    string `json:"value,omitempty"`
	Builtin  string `json:"builtin,omitempty"`
	Feature  string `json:"feature,omitempty"`
	Type     string `json:"type,omitempty"`
}

type Filteraction struct {
	Name        string `json:"name,omitempty"`
	Qual        string `json:"qual,omitempty"`
	Servicename string `json:"servicename,omitempty"`
	Value       string `json:"value,omitempty"`
	Respcode    int    `json:"respcode,omitempty"`
	Page        string `json:"page,omitempty"`
	Isdefault   string `json:"isdefault,omitempty"`
	Builtin     string `json:"builtin,omitempty"`
	Feature     string `json:"feature,omitempty"`
}

type Filterpolicyglobalbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     uint32 `json:"priority,omitempty"`
	Activepolicy uint32 `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Filterpolicyvserverbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     uint32 `json:"priority,omitempty"`
	Activepolicy uint32 `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Filterpostbodyinjection struct {
	Postbody  string `json:"postbody,omitempty"`
	Systemiid string `json:"systemiid,omitempty"`
}

type Filterglobalpolicybinding struct {
	Policyname string `json:"policyname,omitempty"`
	Priority   uint32 `json:"priority,omitempty"`
	State      string `json:"state,omitempty"`
}

type Filterpolicybinding struct {
	Name string `json:"name,omitempty"`
}

type Filterpolicycrvserverbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Filterpolicycsvserverbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}
