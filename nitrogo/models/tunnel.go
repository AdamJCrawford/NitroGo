// Derived from Citrix ADC Nitro Go SDK (https://github.com/netscaler/adc-nitro-go)
// Originally licensed under the Apache License, Version 2.0
// Modifications by Adam Crawford, 2026

package models

type Tunnelglobaltrafficpolicybinding struct {
	Policyname             string   `json:"policyname,omitempty"`
	Priority               uint32   `json:"priority,omitempty"`
	State                  string   `json:"state,omitempty"`
	Builtin                []string `json:"builtin,omitempty"`
	Feature                string   `json:"feature,omitempty"`
	Type                   string   `json:"type,omitempty"`
	Numpol                 uint32   `json:"numpol,omitempty"`
	Gotopriorityexpression string   `json:"gotopriorityexpression,omitempty"`
	Policytype             string   `json:"policytype,omitempty"`
	Globalbindtype         string   `json:"globalbindtype,omitempty"`
}

type Tunnelglobaltunneltrafficpolicybinding struct {
	Policyname             string   `json:"policyname,omitempty"`
	Priority               int      `json:"priority,omitempty"`
	State                  string   `json:"state,omitempty"`
	Builtin                []string `json:"builtin,omitempty"`
	Feature                string   `json:"feature,omitempty"`
	Type                   string   `json:"type,omitempty"`
	Numpol                 int      `json:"numpol,omitempty"`
	Gotopriorityexpression string   `json:"gotopriorityexpression,omitempty"`
	Policytype             string   `json:"policytype,omitempty"`
	Globalbindtype         string   `json:"globalbindtype,omitempty"`
}

type Tunneltrafficpolicy struct {
	Name               string `json:"name,omitempty"`
	Rule               string `json:"rule,omitempty"`
	Action             string `json:"action,omitempty"`
	Comment            string `json:"comment,omitempty"`
	Logaction          string `json:"logaction,omitempty"`
	Newname            string `json:"newname,omitempty"`
	Expressiontype     string `json:"expressiontype,omitempty"`
	Hits               string `json:"hits,omitempty"`
	Undefhits          string `json:"undefhits,omitempty"`
	Txbytes            string `json:"txbytes,omitempty"`
	Rxbytes            string `json:"rxbytes,omitempty"`
	Clientttlb         string `json:"clientttlb,omitempty"`
	Clienttransactions string `json:"clienttransactions,omitempty"`
	Serverttlb         string `json:"serverttlb,omitempty"`
	Servertransactions string `json:"servertransactions,omitempty"`
	Isdefault          string `json:"isdefault,omitempty"`
	Builtin            string `json:"builtin,omitempty"`
	Feature            string `json:"feature,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Tunneltrafficpolicybinding struct {
	Name string `json:"name,omitempty"`
}

type Tunneltrafficpolicyglobalbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Activepolicy           uint32 `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Tunneltrafficpolicytunnelglobalbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Tunnelglobalbinding struct {
}
