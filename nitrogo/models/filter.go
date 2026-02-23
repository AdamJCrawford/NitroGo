// Derived from Citrix ADC Nitro Go SDK (https://github.com/netscaler/adc-nitro-go)
// Originally licensed under the Apache License, Version 2.0
// Modifications by Adam Crawford, 2026

package models

type FilterPolicyFilterGlobalBinding struct {
	BoundTo      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	ActivePolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type FilterPreBodyInjection struct {
	PreBody   string `json:"prebody,omitempty"`
	SystemIID string `json:"systemiid,omitempty"`
}

type FilterGlobalBinding struct {
}

type FilterGlobalFilterPolicyBinding struct {
	PolicyName string `json:"policyname,omitempty"`
	Priority   int    `json:"priority,omitempty"`
	State      string `json:"state,omitempty"`
}

type FilterPolicy struct {
	Name      string `json:"name,omitempty"`
	Rule      string `json:"rule,omitempty"`
	ReqAction string `json:"reqaction,omitempty"`
	ResAction string `json:"resaction,omitempty"`
	Hits      string `json:"hits,omitempty"`
}

type FilterPolicyLBVServerBinding struct {
	BoundTo      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	ActivePolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type FilterHTMLInjectionParameter struct {
	Rate          int    `json:"rate,omitempty"`
	Frequency     int    `json:"frequency,omitempty"`
	Strict        string `json:"strict,omitempty"`
	HTMLSearchLen int    `json:"htmlsearchlen,omitempty"`
	Builtin       string `json:"builtin,omitempty"`
	Feature       string `json:"feature,omitempty"`
}

type FilterHTMLInjectionVariable struct {
	Variable string `json:"variable,omitempty"`
	Value    string `json:"value,omitempty"`
	Builtin  string `json:"builtin,omitempty"`
	Feature  string `json:"feature,omitempty"`
	Type     string `json:"type,omitempty"`
}

type FilterAction struct {
	Name        string `json:"name,omitempty"`
	Qual        string `json:"qual,omitempty"`
	ServiceName string `json:"servicename,omitempty"`
	Value       string `json:"value,omitempty"`
	RespCode    int    `json:"respcode,omitempty"`
	Page        string `json:"page,omitempty"`
	IsDefault   string `json:"isdefault,omitempty"`
	Builtin     string `json:"builtin,omitempty"`
	Feature     string `json:"feature,omitempty"`
}

type FilterPolicyGlobalBinding struct {
	BoundTo      string `json:"boundto,omitempty"`
	Priority     uint32 `json:"priority,omitempty"`
	ActivePolicy uint32 `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type FilterPolicyVServerBinding struct {
	BoundTo      string `json:"boundto,omitempty"`
	Priority     uint32 `json:"priority,omitempty"`
	ActivePolicy uint32 `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type FilterPostBodyInjection struct {
	PostBody  string `json:"postbody,omitempty"`
	SystemIID string `json:"systemiid,omitempty"`
}

type FilterGlobalPolicyBinding struct {
	PolicyName string `json:"policyname,omitempty"`
	Priority   uint32 `json:"priority,omitempty"`
	State      string `json:"state,omitempty"`
}

type FilterPolicyBinding struct {
	Name string `json:"name,omitempty"`
}

type FilterPolicyCRVServerBinding struct {
	BoundTo      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	ActivePolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type FilterPolicyCSVServerBinding struct {
	BoundTo      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	ActivePolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}
