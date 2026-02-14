// Derived from Citrix ADC Nitro Go SDK (https://github.com/netscaler/adc-nitro-go)
// Originally licensed under the Apache License, Version 2.0
// Modifications by Adam Crawford, 2026

package models

type Transformglobalbinding struct {
}

type Transformpolicytransformglobalbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Transformpolicyvserverbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Activepolicy           int32  `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Transformpolicylabelpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Invokelabelname        string `json:"invoke_labelname,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
}

type Transformpolicylbvserverbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Transformpolicylabel struct {
	Labelname              string `json:"labelname,omitempty"`
	Policylabeltype        string `json:"policylabeltype,omitempty"`
	Newname                string `json:"newname,omitempty"`
	Numpol                 string `json:"numpol,omitempty"`
	Hits                   string `json:"hits,omitempty"`
	Priority               string `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Invokelabelname        string `json:"invoke_labelname,omitempty"`
	Description            string `json:"description,omitempty"`
	Nextgenapiresource     string `json:"_nextgenapiresource,omitempty"`
}

type Transformpolicylabeltransformpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Invokelabelname        string `json:"invoke_labelname,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
}

type Transformglobalpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Type                   string `json:"type,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Numpol                 uint32 `json:"numpol,omitempty"`
	Flowtype               uint32 `json:"flowtype,omitempty"`
	Globalbindtype         string `json:"globalbindtype,omitempty"`
}

type Transformglobaltransformpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Type                   string `json:"type,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Numpol                 int    `json:"numpol,omitempty"`
	Flowtype               int    `json:"flowtype,omitempty"`
	Globalbindtype         string `json:"globalbindtype,omitempty"`
}

type Transformpolicycsvserverbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Transformpolicyglobalbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Activepolicy           int32  `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Transformpolicypolicylabelbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Activepolicy           int32  `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Transformpolicytransformpolicylabelbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Transformpolicylabelpolicybindingbinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Invokelabelname        string `json:"invoke_labelname,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
}

type Transformprofileactionbinding struct {
	Actionname       string `json:"actionname,omitempty"`
	Priority         uint32 `json:"priority,omitempty"`
	State            string `json:"state,omitempty"`
	Profilename      string `json:"profilename,omitempty"`
	Requrlfrom       string `json:"requrlfrom,omitempty"`
	Requrlinto       string `json:"requrlinto,omitempty"`
	Resurlfrom       string `json:"resurlfrom,omitempty"`
	Resurlinto       string `json:"resurlinto,omitempty"`
	Cookiedomainfrom string `json:"cookiedomainfrom,omitempty"`
	Cookiedomaininto string `json:"cookiedomaininto,omitempty"`
	Actioncomment    string `json:"actioncomment,omitempty"`
	Name             string `json:"name,omitempty"`
}

type Transformaction struct {
	Name               string `json:"name,omitempty"`
	Profilename        string `json:"profilename,omitempty"`
	Priority           int    `json:"priority,omitempty"`
	State              string `json:"state,omitempty"`
	Requrlfrom         string `json:"requrlfrom,omitempty"`
	Requrlinto         string `json:"requrlinto,omitempty"`
	Resurlfrom         string `json:"resurlfrom,omitempty"`
	Resurlinto         string `json:"resurlinto,omitempty"`
	Cookiedomainfrom   string `json:"cookiedomainfrom,omitempty"`
	Cookiedomaininto   string `json:"cookiedomaininto,omitempty"`
	Comment            string `json:"comment,omitempty"`
	Continuematching   string `json:"continuematching,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Transformpolicy struct {
	Name               string `json:"name,omitempty"`
	Rule               string `json:"rule,omitempty"`
	Profilename        string `json:"profilename,omitempty"`
	Comment            string `json:"comment,omitempty"`
	Logaction          string `json:"logaction,omitempty"`
	Newname            string `json:"newname,omitempty"`
	Hits               string `json:"hits,omitempty"`
	Isdefault          string `json:"isdefault,omitempty"`
	Builtin            string `json:"builtin,omitempty"`
	Feature            string `json:"feature,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Transformpolicybinding struct {
	Name string `json:"name,omitempty"`
}

type Transformpolicylabelbinding struct {
	Labelname string `json:"labelname,omitempty"`
}

type Transformprofile struct {
	Name                           string `json:"name,omitempty"`
	Type                           string `json:"type,omitempty"`
	Onlytransformabsurlinbody      string `json:"onlytransformabsurlinbody,omitempty"`
	Comment                        string `json:"comment,omitempty"`
	Regexforfindingurlinjavascript string `json:"regexforfindingurlinjavascript,omitempty"`
	Regexforfindingurlincss        string `json:"regexforfindingurlincss,omitempty"`
	Regexforfindingurlinxcomponent string `json:"regexforfindingurlinxcomponent,omitempty"`
	Regexforfindingurlinxml        string `json:"regexforfindingurlinxml,omitempty"`
	Additionalreqheaderslist       string `json:"additionalreqheaderslist,omitempty"`
	Additionalrespheaderslist      string `json:"additionalrespheaderslist,omitempty"`
	Nextgenapiresource             string `json:"_nextgenapiresource,omitempty"`
}

type Transformprofilebinding struct {
	Name string `json:"name,omitempty"`
}

type Transformprofiletransformactionbinding struct {
	Actionname       string `json:"actionname,omitempty"`
	Priority         int    `json:"priority,omitempty"`
	State            string `json:"state,omitempty"`
	Profilename      string `json:"profilename,omitempty"`
	Requrlfrom       string `json:"requrlfrom,omitempty"`
	Requrlinto       string `json:"requrlinto,omitempty"`
	Resurlfrom       string `json:"resurlfrom,omitempty"`
	Resurlinto       string `json:"resurlinto,omitempty"`
	Cookiedomainfrom string `json:"cookiedomainfrom,omitempty"`
	Cookiedomaininto string `json:"cookiedomaininto,omitempty"`
	Actioncomment    string `json:"actioncomment,omitempty"`
	Name             string `json:"name,omitempty"`
}
