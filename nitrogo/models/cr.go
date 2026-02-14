// Derived from Citrix ADC Nitro Go SDK (https://github.com/netscaler/adc-nitro-go)
// Originally licensed under the Apache License, Version 2.0
// Modifications by Adam Crawford, 2026

package models

type Craction struct {
	Name               string `json:"name,omitempty"`
	Crtype             string `json:"crtype,omitempty"`
	Builtin            string `json:"builtin,omitempty"`
	Feature            string `json:"feature,omitempty"`
	Isdefault          string `json:"isdefault,omitempty"`
	Hits               string `json:"hits,omitempty"`
	Referencecount     string `json:"referencecount,omitempty"`
	Undefhits          string `json:"undefhits,omitempty"`
	Comment            string `json:"comment,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Crpolicybinding struct {
	Policyname string `json:"policyname,omitempty"`
}

type Crpolicycrvserverbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Hits                   int    `json:"hits,omitempty"`
	Bindhits               int    `json:"bindhits,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
}

type Crvserverappflowpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
	Targetvserver          string `json:"targetvserver,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Crvservercmppolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Inherited              string `json:"inherited,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
	Targetvserver          string `json:"targetvserver,omitempty"`
}

type Crvserverlbvserverbinding struct {
	Lbvserver string `json:"lbvserver,omitempty"`
	Hits      int    `json:"hits,omitempty"`
	Name      string `json:"name,omitempty"`
}

type Crvserverpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Targetvserver          string `json:"targetvserver,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Hits                   uint32 `json:"hits,omitempty"`
	Pipolicyhits           uint32 `json:"pipolicyhits,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Inherited              string `json:"inherited,omitempty"`
	Sc                     string `json:"sc,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Crvserverprofilebinding struct {
	Analyticsprofile string `json:"analyticsprofile,omitempty"`
	Name             string `json:"name,omitempty"`
}

type Crvservericapolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
	Targetvserver          string `json:"targetvserver,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
}

type Crvservermapbinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Targetvserver          string `json:"targetvserver,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
}

type Crvserverpolicymapbinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Targetvserver          string `json:"targetvserver,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
}

type Crvserverrewritepolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
	Targetvserver          string `json:"targetvserver,omitempty"`
}

type Crpolicy struct {
	Policyname         string `json:"policyname,omitempty"`
	Rule               string `json:"rule,omitempty"`
	Action             string `json:"action,omitempty"`
	Logaction          string `json:"logaction,omitempty"`
	Newname            string `json:"newname,omitempty"`
	Boundto            string `json:"boundto,omitempty"`
	Vstype             string `json:"vstype,omitempty"`
	Hits               string `json:"hits,omitempty"`
	Priority           string `json:"priority,omitempty"`
	Activepolicy       string `json:"activepolicy,omitempty"`
	Labelname          string `json:"labelname,omitempty"`
	Labeltype          string `json:"labeltype,omitempty"`
	Builtin            string `json:"builtin,omitempty"`
	Feature            string `json:"feature,omitempty"`
	Isdefault          string `json:"isdefault,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Crpolicyvserverbinding struct {
	Domain                 string `json:"domain,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Hits                   uint32 `json:"hits,omitempty"`
	Pihits                 uint32 `json:"pihits,omitempty"`
	Pipolicyhits           uint32 `json:"pipolicyhits,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
}

type Crvserver struct {
	Name                     string `json:"name,omitempty"`
	Td                       int    `json:"td,omitempty"`
	Servicetype              string `json:"servicetype,omitempty"`
	Ipv46                    string `json:"ipv46,omitempty"`
	Port                     int    `json:"port,omitempty"`
	Ipset                    string `json:"ipset,omitempty"`
	Range                    int    `json:"range,omitempty"`
	Cachetype                string `json:"cachetype,omitempty"`
	Redirect                 string `json:"redirect,omitempty"`
	Onpolicymatch            string `json:"onpolicymatch,omitempty"`
	Redirecturl              string `json:"redirecturl,omitempty"`
	Clttimeout               int    `json:"clttimeout,omitempty"`
	Precedence               string `json:"precedence,omitempty"`
	Arp                      string `json:"arp,omitempty"`
	Ghost                    string `json:"ghost,omitempty"`
	Map                      string `json:"map,omitempty"`
	Format                   string `json:"format,omitempty"`
	Via                      string `json:"via,omitempty"`
	Cachevserver             string `json:"cachevserver,omitempty"`
	Dnsvservername           string `json:"dnsvservername,omitempty"`
	Destinationvserver       string `json:"destinationvserver,omitempty"`
	Domain                   string `json:"domain,omitempty"`
	Sopersistencetimeout     int    `json:"sopersistencetimeout,omitempty"`
	Sothreshold              int    `json:"sothreshold,omitempty"`
	Reuse                    string `json:"reuse,omitempty"`
	State                    string `json:"state,omitempty"`
	Downstateflush           string `json:"downstateflush,omitempty"`
	Backupvserver            string `json:"backupvserver,omitempty"`
	Disableprimaryondown     string `json:"disableprimaryondown,omitempty"`
	L2conn                   string `json:"l2conn,omitempty"`
	Backendssl               string `json:"backendssl,omitempty"`
	Listenpolicy             string `json:"listenpolicy,omitempty"`
	Listenpriority           int    `json:"listenpriority,omitempty"`
	Tcpprofilename           string `json:"tcpprofilename,omitempty"`
	Httpprofilename          string `json:"httpprofilename,omitempty"`
	Comment                  string `json:"comment,omitempty"`
	Srcipexpr                string `json:"srcipexpr,omitempty"`
	Originusip               string `json:"originusip,omitempty"`
	Useportrange             string `json:"useportrange,omitempty"`
	Appflowlog               string `json:"appflowlog,omitempty"`
	Netprofile               string `json:"netprofile,omitempty"`
	Icmpvsrresponse          string `json:"icmpvsrresponse,omitempty"`
	Rhistate                 string `json:"rhistate,omitempty"`
	Useoriginipportforcache  string `json:"useoriginipportforcache,omitempty"`
	Tcpprobeport             int    `json:"tcpprobeport,omitempty"`
	Probeprotocol            string `json:"probeprotocol,omitempty"`
	Probesuccessresponsecode string `json:"probesuccessresponsecode,omitempty"`
	Probeport                int    `json:"probeport,omitempty"`
	Disallowserviceaccess    string `json:"disallowserviceaccess,omitempty"`
	Newname                  string `json:"newname,omitempty"`
	Ip                       string `json:"ip,omitempty"`
	Value                    string `json:"value,omitempty"`
	Ngname                   string `json:"ngname,omitempty"`
	Type                     string `json:"type,omitempty"`
	Curstate                 string `json:"curstate,omitempty"`
	Status                   string `json:"status,omitempty"`
	Authentication           string `json:"authentication,omitempty"`
	Homepage                 string `json:"homepage,omitempty"`
	Rule                     string `json:"rule,omitempty"`
	Policyname               string `json:"policyname,omitempty"`
	Pipolicyhits             string `json:"pipolicyhits,omitempty"`
	Servicename              string `json:"servicename,omitempty"`
	Weight                   string `json:"weight,omitempty"`
	Targetvserver            string `json:"targetvserver,omitempty"`
	Priority                 string `json:"priority,omitempty"`
	Somethod                 string `json:"somethod,omitempty"`
	Sopersistence            string `json:"sopersistence,omitempty"`
	Lbvserver                string `json:"lbvserver,omitempty"`
	Bindpoint                string `json:"bindpoint,omitempty"`
	Invoke                   string `json:"invoke,omitempty"`
	Labeltype                string `json:"labeltype,omitempty"`
	Labelname                string `json:"labelname,omitempty"`
	Gotopriorityexpression   string `json:"gotopriorityexpression,omitempty"`
	Nodefaultbindings        string `json:"nodefaultbindings,omitempty"`
	Nextgenapiresource       string `json:"_nextgenapiresource,omitempty"`
}

type Crvserveranalyticsprofilebinding struct {
	Analyticsprofile string `json:"analyticsprofile,omitempty"`
	Name             string `json:"name,omitempty"`
}

type Crvserverappfwpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
	Targetvserver          string `json:"targetvserver,omitempty"`
}

type Crvserverappqoepolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
	Targetvserver          string `json:"targetvserver,omitempty"`
}

type Crvserverbinding struct {
	Name string `json:"name,omitempty"`
}

type Crvservercachepolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
	Targetvserver          string `json:"targetvserver,omitempty"`
}

type Crvservercrpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Hits                   int    `json:"hits,omitempty"`
	Pipolicyhits           int    `json:"pipolicyhits,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Name                   string `json:"name,omitempty"`
	Targetvserver          string `json:"targetvserver,omitempty"`
}

type Crvservercspolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Targetvserver          string `json:"targetvserver,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Hits                   int    `json:"hits,omitempty"`
	Pipolicyhits           int    `json:"pipolicyhits,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Crvserverfeopolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Name                   string `json:"name,omitempty"`
	Targetvserver          string `json:"targetvserver,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
}

type Crvserverfilterpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Inherited              string `json:"inherited,omitempty"`
	Name                   string `json:"name,omitempty"`
	Targetvserver          string `json:"targetvserver,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
}

type Crvserverresponderpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
	Targetvserver          string `json:"targetvserver,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Crvserverspilloverpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Name                   string `json:"name,omitempty"`
	Targetvserver          string `json:"targetvserver,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
}

type Crvservervserverbinding struct {
	Lbvserver string `json:"lbvserver,omitempty"`
	Hits      uint32 `json:"hits,omitempty"`
	Name      string `json:"name,omitempty"`
}
