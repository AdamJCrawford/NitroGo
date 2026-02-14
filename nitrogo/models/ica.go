// Derived from Citrix ADC Nitro Go SDK (https://github.com/netscaler/adc-nitro-go)
// Originally licensed under the Apache License, Version 2.0
// Modifications by Adam Crawford, 2026

package models

type Icapolicyglobalbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Activepolicy           int32  `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Icapolicyvpnvserverbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Icapolicyvserverbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Activepolicy           int32  `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Icaaccessprofile struct {
	Name                         string `json:"name,omitempty"`
	Connectclientlptports        string `json:"connectclientlptports,omitempty"`
	Clientaudioredirection       string `json:"clientaudioredirection,omitempty"`
	Localremotedatasharing       string `json:"localremotedatasharing,omitempty"`
	Clientclipboardredirection   string `json:"clientclipboardredirection,omitempty"`
	Clientcomportredirection     string `json:"clientcomportredirection,omitempty"`
	Clientdriveredirection       string `json:"clientdriveredirection,omitempty"`
	Clientprinterredirection     string `json:"clientprinterredirection,omitempty"`
	Multistream                  string `json:"multistream,omitempty"`
	Clientusbdriveredirection    string `json:"clientusbdriveredirection,omitempty"`
	Clienttwaindeviceredirection string `json:"clienttwaindeviceredirection,omitempty"`
	Wiaredirection               string `json:"wiaredirection,omitempty"`
	Draganddrop                  string `json:"draganddrop,omitempty"`
	Smartcardredirection         string `json:"smartcardredirection,omitempty"`
	Fido2redirection             string `json:"fido2redirection,omitempty"`
	Refcnt                       string `json:"refcnt,omitempty"`
	Builtin                      string `json:"builtin,omitempty"`
	Feature                      string `json:"feature,omitempty"`
	Isdefault                    string `json:"isdefault,omitempty"`
	Nextgenapiresource           string `json:"_nextgenapiresource,omitempty"`
}

type Icaaction struct {
	Name               string `json:"name,omitempty"`
	Accessprofilename  string `json:"accessprofilename,omitempty"`
	Latencyprofilename string `json:"latencyprofilename,omitempty"`
	Newname            string `json:"newname,omitempty"`
	Hits               string `json:"hits,omitempty"`
	Referencecount     string `json:"referencecount,omitempty"`
	Undefhits          string `json:"undefhits,omitempty"`
	Builtin            string `json:"builtin,omitempty"`
	Feature            string `json:"feature,omitempty"`
	Isdefault          string `json:"isdefault,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Icaglobalicapolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Type                   string `json:"type,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Numpol                 int    `json:"numpol,omitempty"`
	Flowtype               int    `json:"flowtype,omitempty"`
	Globalbindtype         string `json:"globalbindtype,omitempty"`
}

type Icaglobalpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Type                   string `json:"type,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Numpol                 uint32 `json:"numpol,omitempty"`
	Flowtype               uint32 `json:"flowtype,omitempty"`
	Globalbindtype         string `json:"globalbindtype,omitempty"`
}

type Icaparameter struct {
	Enablesronhafailover string `json:"enablesronhafailover,omitempty"`
	Hdxinsightnonnsap    string `json:"hdxinsightnonnsap,omitempty"`
	Edtpmtuddf           string `json:"edtpmtuddf,omitempty"`
	Edtpmtuddftimeout    int    `json:"edtpmtuddftimeout,omitempty"`
	L7latencyfrequency   int    `json:"l7latencyfrequency,omitempty"`
	Edtlosstolerant      string `json:"edtlosstolerant,omitempty"`
	Edtpmtudrediscovery  string `json:"edtpmtudrediscovery,omitempty"`
	Dfpersistence        string `json:"dfpersistence,omitempty"`
	Builtin              string `json:"builtin,omitempty"`
	Nextgenapiresource   string `json:"_nextgenapiresource,omitempty"`
}

type Icapolicyicaglobalbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Icaglobalbinding struct {
}

type Icalatencyprofile struct {
	Name                     string `json:"name,omitempty"`
	L7latencymonitoring      string `json:"l7latencymonitoring,omitempty"`
	L7latencythresholdfactor int    `json:"l7latencythresholdfactor,omitempty"`
	L7latencywaittime        int    `json:"l7latencywaittime,omitempty"`
	L7latencynotifyinterval  int    `json:"l7latencynotifyinterval,omitempty"`
	L7latencymaxnotifycount  int    `json:"l7latencymaxnotifycount,omitempty"`
	Refcnt                   string `json:"refcnt,omitempty"`
	Builtin                  string `json:"builtin,omitempty"`
	Feature                  string `json:"feature,omitempty"`
	Isdefault                string `json:"isdefault,omitempty"`
	Nextgenapiresource       string `json:"_nextgenapiresource,omitempty"`
}

type Icapolicy struct {
	Name               string `json:"name,omitempty"`
	Rule               string `json:"rule,omitempty"`
	Action             string `json:"action,omitempty"`
	Comment            string `json:"comment,omitempty"`
	Logaction          string `json:"logaction,omitempty"`
	Newname            string `json:"newname,omitempty"`
	Hits               string `json:"hits,omitempty"`
	Undefhits          string `json:"undefhits,omitempty"`
	Builtin            string `json:"builtin,omitempty"`
	Feature            string `json:"feature,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Icapolicybinding struct {
	Name string `json:"name,omitempty"`
}

type Icapolicycrvserverbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
}
