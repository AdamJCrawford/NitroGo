// Derived from Citrix ADC Nitro Go SDK (https://github.com/netscaler/adc-nitro-go)
// Originally licensed under the Apache License, Version 2.0
// Modifications by Adam Crawford, 2026

package models

type Cmpparameter struct {
	Cmplevel                    string `json:"cmplevel,omitempty"`
	Quantumsize                 int    `json:"quantumsize,omitempty"`
	Servercmp                   string `json:"servercmp,omitempty"`
	Heurexpiry                  string `json:"heurexpiry,omitempty"`
	Heurexpirythres             int    `json:"heurexpirythres,omitempty"`
	Heurexpiryhistwt            int    `json:"heurexpiryhistwt,omitempty"`
	Minressize                  int    `json:"minressize,omitempty"`
	Cmpbypasspct                int    `json:"cmpbypasspct,omitempty"`
	Cmponpush                   string `json:"cmponpush,omitempty"`
	Policytype                  string `json:"policytype,omitempty"`
	Addvaryheader               string `json:"addvaryheader,omitempty"`
	Varyheadervalue             string `json:"varyheadervalue,omitempty"`
	Externalcache               string `json:"externalcache,omitempty"`
	Randomgzipfilename          string `json:"randomgzipfilename,omitempty"`
	Randomgzipfilenameminlength int    `json:"randomgzipfilenameminlength,omitempty"`
	Randomgzipfilenamemaxlength int    `json:"randomgzipfilenamemaxlength,omitempty"`
	Builtin                     string `json:"builtin,omitempty"`
	Feature                     string `json:"feature,omitempty"`
	Nextgenapiresource          string `json:"_nextgenapiresource,omitempty"`
}

type Cmppolicycmpglobalbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Cmppolicylbvserverbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Cmppolicylabelpolicybindingbinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Invokelabelname        string `json:"invoke_labelname,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
}

type Cmpglobalbinding struct {
}

type Cmppolicy struct {
	Name               string `json:"name,omitempty"`
	Rule               string `json:"rule,omitempty"`
	Resaction          string `json:"resaction,omitempty"`
	Newname            string `json:"newname,omitempty"`
	Reqaction          string `json:"reqaction,omitempty"`
	Hits               string `json:"hits,omitempty"`
	Txbytes            string `json:"txbytes,omitempty"`
	Rxbytes            string `json:"rxbytes,omitempty"`
	Clientttlb         string `json:"clientttlb,omitempty"`
	Clienttransactions string `json:"clienttransactions,omitempty"`
	Serverttlb         string `json:"serverttlb,omitempty"`
	Servertransactions string `json:"servertransactions,omitempty"`
	Description        string `json:"description,omitempty"`
	Builtin            string `json:"builtin,omitempty"`
	Feature            string `json:"feature,omitempty"`
	Isdefault          string `json:"isdefault,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Cmppolicycsvserverbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Cmppolicypolicylabelbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Activepolicy           uint32 `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Cmppolicylabel struct {
	Labelname              string `json:"labelname,omitempty"`
	Type                   string `json:"type,omitempty"`
	Newname                string `json:"newname,omitempty"`
	Numpol                 string `json:"numpol,omitempty"`
	Hits                   string `json:"hits,omitempty"`
	Priority               string `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Invokelabelname        string `json:"invoke_labelname,omitempty"`
	Flowtype               string `json:"flowtype,omitempty"`
	Description            string `json:"description,omitempty"`
	Nextgenapiresource     string `json:"_nextgenapiresource,omitempty"`
}

type Cmppolicylabelcmppolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Invokelabelname        string `json:"invoke_labelname,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
}

type Cmppolicylabelpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Invokelabelname        string `json:"invoke_labelname,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
}

type Cmpglobalcmppolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Type                   string `json:"type,omitempty"`
	Numpol                 int    `json:"numpol,omitempty"`
	Globalbindtype         string `json:"globalbindtype,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
}

type Cmppolicybinding struct {
	Name string `json:"name,omitempty"`
}

type Cmppolicycmppolicylabelbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Cmppolicycrvserverbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Cmppolicyglobalbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Activepolicy           uint32 `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Cmpaction struct {
	Name               string `json:"name,omitempty"`
	Cmptype            string `json:"cmptype,omitempty"`
	Addvaryheader      string `json:"addvaryheader,omitempty"`
	Varyheadervalue    string `json:"varyheadervalue,omitempty"`
	Deltatype          string `json:"deltatype,omitempty"`
	Newname            string `json:"newname,omitempty"`
	Builtin            string `json:"builtin,omitempty"`
	Feature            string `json:"feature,omitempty"`
	Isdefault          string `json:"isdefault,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Cmpglobalpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	State                  string `json:"state,omitempty"`
	Type                   string `json:"type,omitempty"`
	Numpol                 uint32 `json:"numpol,omitempty"`
	Policytype             string `json:"policytype,omitempty"`
	Globalbindtype         string `json:"globalbindtype,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
}

type Cmppolicyvserverbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Activepolicy           uint32 `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Cmppolicylabelbinding struct {
	Labelname string `json:"labelname,omitempty"`
}
