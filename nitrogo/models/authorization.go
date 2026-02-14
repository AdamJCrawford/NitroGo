// Derived from Citrix ADC Nitro Go SDK (https://github.com/netscaler/adc-nitro-go)
// Originally licensed under the Apache License, Version 2.0
// Modifications by Adam Crawford, 2026

package models

type Authorizationpolicyaaauserbinding struct {
	Boundto  string `json:"boundto,omitempty"`
	Priority int    `json:"priority,omitempty"`
	Name     string `json:"name,omitempty"`
}

type Authorizationpolicypolicylabelbinding struct {
	Boundto  string `json:"boundto,omitempty"`
	Priority uint32 `json:"priority,omitempty"`
	Name     string `json:"name,omitempty"`
}

type Authorizationpolicylabelauthorizationpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Invokelabelname        string `json:"invoke_labelname,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
}

type Authorizationpolicy struct {
	Name               string `json:"name,omitempty"`
	Rule               string `json:"rule,omitempty"`
	Action             string `json:"action,omitempty"`
	Newname            string `json:"newname,omitempty"`
	Activepolicy       string `json:"activepolicy,omitempty"`
	Expressiontype     string `json:"expressiontype,omitempty"`
	Hits               string `json:"hits,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Authorizationpolicybinding struct {
	Name string `json:"name,omitempty"`
}

type Authorizationpolicycsvserverbinding struct {
	Boundto  string `json:"boundto,omitempty"`
	Priority int    `json:"priority,omitempty"`
	Name     string `json:"name,omitempty"`
}

type Authorizationpolicyuserbinding struct {
	Boundto  string `json:"boundto,omitempty"`
	Priority uint32 `json:"priority,omitempty"`
	Name     string `json:"name,omitempty"`
}

type Authorizationaction struct {
	Name               string `json:"name,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Authorizationpolicyauthorizationpolicylabelbinding struct {
	Boundto  string `json:"boundto,omitempty"`
	Priority int    `json:"priority,omitempty"`
	Name     string `json:"name,omitempty"`
}

type Authorizationpolicylbvserverbinding struct {
	Boundto  string `json:"boundto,omitempty"`
	Priority int    `json:"priority,omitempty"`
	Name     string `json:"name,omitempty"`
}

type Authorizationpolicyvserverbinding struct {
	Boundto  string `json:"boundto,omitempty"`
	Priority uint32 `json:"priority,omitempty"`
	Name     string `json:"name,omitempty"`
}

type Authorizationpolicylabel struct {
	Labelname              string `json:"labelname,omitempty"`
	Newname                string `json:"newname,omitempty"`
	Numpol                 string `json:"numpol,omitempty"`
	Hits                   string `json:"hits,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               string `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Invokelabelname        string `json:"invoke_labelname,omitempty"`
	Flowtype               string `json:"flowtype,omitempty"`
	Description            string `json:"description,omitempty"`
	Nextgenapiresource     string `json:"_nextgenapiresource,omitempty"`
}

type Authorizationpolicylabelbinding struct {
	Labelname string `json:"labelname,omitempty"`
}

type Authorizationpolicylabelpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Invokelabelname        string `json:"invoke_labelname,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
}

type Authorizationpolicyaaagroupbinding struct {
	Boundto  string `json:"boundto,omitempty"`
	Priority int    `json:"priority,omitempty"`
	Name     string `json:"name,omitempty"`
}

type Authorizationpolicygroupbinding struct {
	Boundto  string `json:"boundto,omitempty"`
	Priority uint32 `json:"priority,omitempty"`
	Name     string `json:"name,omitempty"`
}
