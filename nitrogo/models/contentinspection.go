// Derived from Citrix ADC Nitro Go SDK (https://github.com/netscaler/adc-nitro-go)
// Originally licensed under the Apache License, Version 2.0
// Modifications by Adam Crawford, 2026

package models

type Contentinspectionpolicy struct {
	Name               string `json:"name,omitempty"`
	Rule               string `json:"rule,omitempty"`
	Action             string `json:"action,omitempty"`
	Undefaction        string `json:"undefaction,omitempty"`
	Comment            string `json:"comment,omitempty"`
	Logaction          string `json:"logaction,omitempty"`
	Newname            string `json:"newname,omitempty"`
	Hits               string `json:"hits,omitempty"`
	Undefhits          string `json:"undefhits,omitempty"`
	Builtin            string `json:"builtin,omitempty"`
	Feature            string `json:"feature,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Contentinspectionpolicycontentinspectionglobalbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Contentinspectionpolicycsvserverbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Contentinspectionpolicylbvserverbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Contentinspectionpolicyvserverbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Activepolicy           int32  `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Contentinspectionpolicylabel struct {
	Labelname              string `json:"labelname,omitempty"`
	Type                   string `json:"type,omitempty"`
	Comment                string `json:"comment,omitempty"`
	Newname                string `json:"newname,omitempty"`
	Numpol                 string `json:"numpol,omitempty"`
	Hits                   string `json:"hits,omitempty"`
	Priority               string `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Invokelabelname        string `json:"invoke_labelname,omitempty"`
	Flowtype               string `json:"flowtype,omitempty"`
	Isdefault              string `json:"isdefault,omitempty"`
	Nextgenapiresource     string `json:"_nextgenapiresource,omitempty"`
}

type Contentinspectionpolicylabelbinding struct {
	Labelname string `json:"labelname,omitempty"`
}

type Contentinspectionglobalpolicybinding struct {
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

type Contentinspectionpolicycontentinspectionpolicylabelbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Contentinspectionpolicypolicylabelbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Activepolicy           int32  `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Contentinspectionprofile struct {
	Name               string `json:"name,omitempty"`
	Type               string `json:"type,omitempty"`
	Ingressinterface   string `json:"ingressinterface,omitempty"`
	Ingressvlan        int    `json:"ingressvlan,omitempty"`
	Egressinterface    string `json:"egressinterface,omitempty"`
	Iptunnel           string `json:"iptunnel,omitempty"`
	Egressvlan         int    `json:"egressvlan,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Contentinspectioncallout struct {
	Name               string `json:"name,omitempty"`
	Type               string `json:"type,omitempty"`
	Profilename        string `json:"profilename,omitempty"`
	Servername         string `json:"servername,omitempty"`
	Serverip           string `json:"serverip,omitempty"`
	Serverport         int    `json:"serverport,omitempty"`
	Returntype         string `json:"returntype,omitempty"`
	Resultexpr         string `json:"resultexpr,omitempty"`
	Comment            string `json:"comment,omitempty"`
	Hits               string `json:"hits,omitempty"`
	Undefhits          string `json:"undefhits,omitempty"`
	Undefreason        string `json:"undefreason,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Contentinspectionpolicybinding struct {
	Name string `json:"name,omitempty"`
}

type Contentinspectionpolicyglobalbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Activepolicy           int32  `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Contentinspectionpolicylabelcontentinspectionpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Invokelabelname        string `json:"invoke_labelname,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
}

type Contentinspectionaction struct {
	Name               string `json:"name,omitempty"`
	Type               string `json:"type,omitempty"`
	Servername         string `json:"servername,omitempty"`
	Serverip           string `json:"serverip,omitempty"`
	Serverport         int    `json:"serverport,omitempty"`
	Icapprofilename    string `json:"icapprofilename,omitempty"`
	Ifserverdown       string `json:"ifserverdown,omitempty"`
	Reqtimeout         string `json:"reqtimeout,omitempty"`
	Reqtimeoutaction   string `json:"reqtimeoutaction,omitempty"`
	Hits               string `json:"hits,omitempty"`
	Referencecount     string `json:"referencecount,omitempty"`
	Undefhits          string `json:"undefhits,omitempty"`
	Builtin            string `json:"builtin,omitempty"`
	Feature            string `json:"feature,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Contentinspectionglobalbinding struct {
}

type Contentinspectionglobalcontentinspectionpolicybinding struct {
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

type Contentinspectionparameter struct {
	Undefaction        string `json:"undefaction,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Contentinspectionpolicylabelpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Invokelabelname        string `json:"invoke_labelname,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
}

type Contentinspectionpolicylabelpolicybindingbinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Invokelabelname        string `json:"invoke_labelname,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
}
